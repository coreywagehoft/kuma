package api_server

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/emicklei/go-restful/v3"
	"github.com/pkg/errors"
	http_prometheus "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/emicklei/go-restful/otelrestful"

	"github.com/kumahq/kuma/pkg/api-server/authn"
	"github.com/kumahq/kuma/pkg/api-server/customization"
	api_server "github.com/kumahq/kuma/pkg/config/api-server"
	kuma_cp "github.com/kumahq/kuma/pkg/config/app/kuma-cp"
	config_core "github.com/kumahq/kuma/pkg/config/core"
	config_store "github.com/kumahq/kuma/pkg/config/core/resources/store"
	config_types "github.com/kumahq/kuma/pkg/config/types"
	"github.com/kumahq/kuma/pkg/core"
	resources_access "github.com/kumahq/kuma/pkg/core/resources/access"
	"github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	"github.com/kumahq/kuma/pkg/core/resources/manager"
	"github.com/kumahq/kuma/pkg/core/resources/model"
	"github.com/kumahq/kuma/pkg/core/resources/registry"
	"github.com/kumahq/kuma/pkg/core/runtime"
	"github.com/kumahq/kuma/pkg/dns/vips"
	"github.com/kumahq/kuma/pkg/envoy/admin"
	"github.com/kumahq/kuma/pkg/metrics"
	"github.com/kumahq/kuma/pkg/plugins/authn/api-server/certs"
	"github.com/kumahq/kuma/pkg/plugins/resources/k8s"
	"github.com/kumahq/kuma/pkg/tokens/builtin"
	tokens_server "github.com/kumahq/kuma/pkg/tokens/builtin/server"
	util_prometheus "github.com/kumahq/kuma/pkg/util/prometheus"
	"github.com/kumahq/kuma/pkg/version"
	xds_context "github.com/kumahq/kuma/pkg/xds/context"
	"github.com/kumahq/kuma/pkg/xds/server"
)

var log = core.Log.WithName("api-server")

type ApiServer struct {
	mux    *http.ServeMux
	config api_server.ApiServerConfig
}

func (a *ApiServer) NeedLeaderElection() bool {
	return false
}

func (a *ApiServer) Address() string {
	return net.JoinHostPort(a.config.HTTP.Interface, strconv.FormatUint(uint64(a.config.HTTP.Port), 10))
}

func (a *ApiServer) Config() api_server.ApiServerConfig {
	return a.config
}

func init() {
	// turn off escape & character so the link in "next" fields for resources is user friendly
	restful.NewEncoder = func(w io.Writer) *json.Encoder {
		encoder := json.NewEncoder(w)
		encoder.SetEscapeHTML(false)
		return encoder
	}
	restful.MarshalIndent = func(v interface{}, prefix, indent string) ([]byte, error) {
		var buf bytes.Buffer
		encoder := restful.NewEncoder(&buf)
		encoder.SetIndent(prefix, indent)
		if err := encoder.Encode(v); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}
}

func NewApiServer(
	resManager manager.ResourceManager,
	meshContextBuilder xds_context.MeshContextBuilder,
	wsManager customization.APIInstaller,
	defs []model.ResourceTypeDescriptor,
	cfg *kuma_cp.Config,
	metrics metrics.Metrics,
	getInstanceId func() string, getClusterId func() string,
	authenticator authn.Authenticator,
	access runtime.Access,
	envoyAdminClient admin.EnvoyAdminClient,
	tokenIssuers builtin.TokenIssuers,
	wsCustomize func(*restful.WebService) error,
) (*ApiServer, error) {
	serverConfig := cfg.ApiServer
	container := restful.NewContainer()

	promMiddleware := middleware.New(middleware.Config{
		Recorder: http_prometheus.NewRecorder(http_prometheus.Config{
			Registry: metrics,
			Prefix:   "api_server",
		}),
	})
	container.Filter(util_prometheus.MetricsHandler("", promMiddleware))

	// NOTE: This must come before any filters that make HTTP calls
	container.Filter(otelrestful.OTelFilter("api-server"))

	if cfg.ApiServer.Authn.LocalhostIsAdmin {
		container.Filter(authn.LocalhostAuthenticator)
	}
	container.Filter(authenticator)

	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{restful.HEADER_AccessControlAllowOrigin},
		AllowedDomains: serverConfig.CorsAllowedDomains,
		Container:      container,
	}
	container.Filter(cors.Filter)

	// We create a WebService and set up resources endpoints and index endpoint instead of creating WebService
	// for every resource like /meshes/{mesh}/traffic-permissions, /meshes/{mesh}/traffic-log etc.
	// because go-restful detects it as a clash (you cannot register 2 WebServices with path /meshes/)
	ws := new(restful.WebService)
	ws.
		Path(cfg.ApiServer.BasePath).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	addResourcesEndpoints(ws, defs, resManager, cfg, access.ResourceAccess)
	addPoliciesWsEndpoints(ws, cfg.Mode, cfg.ApiServer.ReadOnly, defs)
	addInspectEndpoints(ws, cfg, meshContextBuilder, resManager)
	addInspectEnvoyAdminEndpoints(ws, cfg, resManager, access.EnvoyAdminAccess, envoyAdminClient)
	addZoneEndpoints(ws, resManager)
	guiUrl := cfg.ApiServer.GUI.BasePath
	if cfg.ApiServer.GUI.RootUrl != "" {
		guiUrl = cfg.ApiServer.GUI.RootUrl
	}
	apiUrl := cfg.ApiServer.BasePath
	if cfg.ApiServer.RootUrl != "" {
		apiUrl = cfg.ApiServer.RootUrl
	}

	err := addConfigEndpoints(ws, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "could not create configuration webservice")
	}
	enableGUI := cfg.ApiServer.GUI.Enabled && cfg.Mode != config_core.Zone
	if err := addIndexWsEndpoints(ws, getInstanceId, getClusterId, enableGUI, guiUrl); err != nil {
		return nil, errors.Wrap(err, "could not create index webservice")
	}
	addWhoamiEndpoints(ws)

	ws.SetDynamicRoutes(true)
	if err := wsCustomize(ws); err != nil {
		return nil, errors.Wrap(err, "couldn't customize webservice")
	}

	container.Add(ws)

	path := cfg.ApiServer.BasePath
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	container.Add(tokens_server.NewWebservice(
		path+"tokens",
		tokenIssuers.DataplaneToken,
		tokenIssuers.ZoneIngressToken,
		tokenIssuers.ZoneToken,
		access.DataplaneTokenAccess,
		access.ZoneTokenAccess,
	))
	guiPath := cfg.ApiServer.GUI.BasePath
	if !strings.HasSuffix(guiPath, "/") {
		guiPath += "/"
	}
	basePath := guiPath
	if cfg.ApiServer.GUI.RootUrl != "" {
		u, err := url.Parse(cfg.ApiServer.GUI.RootUrl)
		if err != nil {
			return nil, errors.New("Gui.RootUrl is not a valid url")
		}
		basePath = u.Path
	}
	basePath = strings.TrimSuffix(basePath, "/")
	guiHandler, err := NewGuiHandler(guiPath, enableGUI, GuiConfig{
		BaseGuiPath: basePath,
		ApiUrl:      apiUrl,
		Version:     version.Build.Version,
		Product:     version.Product,
		BasedOnKuma: version.Build.BasedOnKuma,
		Mode:        cfg.Mode,
		Environment: cfg.Environment,
		ReadOnly:    cfg.ApiServer.ReadOnly,
	})
	if err != nil {
		return nil, err
	}
	container.Handle(guiPath, guiHandler)

	newApiServer := &ApiServer{
		mux:    container.ServeMux,
		config: *serverConfig,
	}

	wsManager.Install(container)

	return newApiServer, nil
}

func addResourcesEndpoints(ws *restful.WebService, defs []model.ResourceTypeDescriptor, resManager manager.ResourceManager, cfg *kuma_cp.Config, resourceAccess resources_access.ResourceAccess) {
	dpOverviewEndpoints := dataplaneOverviewEndpoints{
		resManager:     resManager,
		resourceAccess: resourceAccess,
	}
	dpOverviewEndpoints.addListEndpoint(ws, "/meshes/{mesh}")
	dpOverviewEndpoints.addFindEndpoint(ws, "/meshes/{mesh}")
	dpOverviewEndpoints.addListEndpoint(ws, "") // listing all resources in all meshes

	zoneOverviewEndpoints := zoneOverviewEndpoints{
		resManager:     resManager,
		resourceAccess: resourceAccess,
	}
	zoneOverviewEndpoints.addFindEndpoint(ws)
	zoneOverviewEndpoints.addListEndpoint(ws)

	zoneIngressOverviewEndpoints := zoneIngressOverviewEndpoints{
		resManager:     resManager,
		resourceAccess: resourceAccess,
	}
	zoneIngressOverviewEndpoints.addFindEndpoint(ws)
	zoneIngressOverviewEndpoints.addListEndpoint(ws)

	zoneEgressOverviewEndpoints := zoneEgressOverviewEndpoints{
		resManager:     resManager,
		resourceAccess: resourceAccess,
	}
	zoneEgressOverviewEndpoints.addFindEndpoint(ws)
	zoneEgressOverviewEndpoints.addListEndpoint(ws)

	globalInsightsEndpoints := globalInsightsEndpoints{
		resManager:     resManager,
		resourceAccess: resourceAccess,
	}
	globalInsightsEndpoints.addEndpoint(ws)

	var k8sMapper k8s.ResourceMapperFunc
	switch cfg.Store.Type {
	case config_store.KubernetesStore:
		k8sMapper = k8s.NewKubernetesMapper(k8s.NewSimpleKubeFactory())
	default:
		k8sMapper = k8s.NewInferenceMapper(cfg.Store.Kubernetes.SystemNamespace, k8s.NewSimpleKubeFactory())
	}
	for _, definition := range defs {
		defType := definition.Name
		if ShouldBeReadOnly(definition.KDSFlags, cfg) {
			definition.ReadOnly = true
		}
		endpoints := resourceEndpoints{
			k8sMapper:      k8sMapper,
			mode:           cfg.Mode,
			resManager:     resManager,
			descriptor:     definition,
			resourceAccess: resourceAccess,
		}
		if cfg.Mode == config_core.Zone && cfg.Multizone != nil && cfg.Multizone.Zone != nil {
			endpoints.zoneName = cfg.Multizone.Zone.Name
		}
		switch defType {
		case mesh.ServiceInsightType:
			// ServiceInsight is a bit different
			ep := serviceInsightEndpoints{
				resourceEndpoints: endpoints,
				addressPortGenerator: func(svc string) string {
					return fmt.Sprintf("%s.%s:%d", svc, cfg.DNSServer.Domain, cfg.DNSServer.ServiceVipPort)
				},
			}
			ep.addCreateOrUpdateEndpoint(ws, "/meshes/{mesh}/"+definition.WsPath)
			ep.addDeleteEndpoint(ws, "/meshes/{mesh}/"+definition.WsPath)
			ep.addFindEndpoint(ws, "/meshes/{mesh}/"+definition.WsPath)
			ep.addListEndpoint(ws, "/meshes/{mesh}/"+definition.WsPath)
			ep.addListEndpoint(ws, "/"+definition.WsPath) // listing all resources in all meshes
		default:
			switch definition.Scope {
			case model.ScopeMesh:
				endpoints.addCreateOrUpdateEndpoint(ws, "/meshes/{mesh}/"+definition.WsPath)
				endpoints.addDeleteEndpoint(ws, "/meshes/{mesh}/"+definition.WsPath)
				endpoints.addFindEndpoint(ws, "/meshes/{mesh}/"+definition.WsPath)
				endpoints.addListEndpoint(ws, "/meshes/{mesh}/"+definition.WsPath)
				endpoints.addListEndpoint(ws, "/"+definition.WsPath) // listing all resources in all meshes
			case model.ScopeGlobal:
				endpoints.addCreateOrUpdateEndpoint(ws, "/"+definition.WsPath)
				endpoints.addDeleteEndpoint(ws, "/"+definition.WsPath)
				endpoints.addFindEndpoint(ws, "/"+definition.WsPath)
				endpoints.addListEndpoint(ws, "/"+definition.WsPath)
			}
		}
	}
}

func ShouldBeReadOnly(kdsFlag model.KDSFlagType, cfg *kuma_cp.Config) bool {
	if cfg.ApiServer.ReadOnly {
		return true
	}
	if kdsFlag == model.KDSDisabled {
		return false
	}
	if cfg.Mode == config_core.Global && !kdsFlag.Has(model.ProvidedByGlobal) {
		return true
	}
	if cfg.Mode == config_core.Zone && !kdsFlag.Has(model.ProvidedByZone) {
		return true
	}
	return false
}

func (a *ApiServer) Start(stop <-chan struct{}) error {
	errChan := make(chan error)

	var httpServer, httpsServer *http.Server
	if a.config.HTTP.Enabled {
		httpServer = a.startHttpServer(errChan)
	}
	if a.config.HTTPS.Enabled {
		var err error
		httpsServer, err = a.startHttpsServer(errChan)
		if err != nil {
			return err
		}
	}
	select {
	case <-stop:
		log.Info("stopping down API Server")
		if httpServer != nil {
			return httpServer.Shutdown(context.Background())
		}
		if httpsServer != nil {
			return httpsServer.Shutdown(context.Background())
		}
	case err := <-errChan:
		return err
	}
	return nil
}

func (a *ApiServer) startHttpServer(errChan chan error) *http.Server {
	server := &http.Server{
		ReadHeaderTimeout: time.Second,
		Addr:              net.JoinHostPort(a.config.HTTP.Interface, strconv.FormatUint(uint64(a.config.HTTP.Port), 10)),
		Handler:           a.mux,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			switch err {
			case http.ErrServerClosed:
				log.Info("shutting down server")
			default:
				log.Error(err, "could not start an HTTP Server")
				errChan <- err
			}
		}
	}()
	log.Info("starting", "interface", a.config.HTTP.Interface, "port", a.config.HTTP.Port)
	return server
}

func (a *ApiServer) startHttpsServer(errChan chan error) (*http.Server, error) {
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12, // to pass gosec (in practice it's always set after.
	}
	var err error
	tlsConfig.MinVersion, err = config_types.TLSVersion(a.config.HTTPS.TlsMinVersion)
	if err != nil {
		return nil, err
	}
	tlsConfig.CipherSuites, err = config_types.TLSCiphers(a.config.HTTPS.TlsCipherSuites)
	if err != nil {
		return nil, err
	}

	err = configureMTLS(tlsConfig, a.config)
	if err != nil {
		return nil, err
	}

	server := &http.Server{
		ReadHeaderTimeout: time.Second,
		Addr:              net.JoinHostPort(a.config.HTTPS.Interface, strconv.FormatUint(uint64(a.config.HTTPS.Port), 10)),
		Handler:           a.mux,
		TLSConfig:         tlsConfig,
	}

	go func() {
		err := server.ListenAndServeTLS(a.config.HTTPS.TlsCertFile, a.config.HTTPS.TlsKeyFile)
		if err != nil {
			switch err {
			case http.ErrServerClosed:
				log.Info("shutting down server")
			default:
				log.Error(err, "could not start an HTTPS Server")
				errChan <- err
			}
		}
	}()
	log.Info("starting", "interface", a.config.HTTPS.Interface, "port", a.config.HTTPS.Port, "tls", true)
	return server, nil
}

func configureMTLS(tlsConfig *tls.Config, cfg api_server.ApiServerConfig) error {
	clientCertPool := x509.NewCertPool()
	if cfg.Auth.ClientCertsDir != "" {
		log.Info("loading client certificates")
		files, err := os.ReadDir(cfg.Auth.ClientCertsDir)
		if err != nil {
			return err
		}
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if !strings.HasSuffix(file.Name(), ".pem") && !strings.HasSuffix(file.Name(), ".crt") {
				log.Info("skipping file without .pem or .crt extension", "file", file.Name())
				continue
			}
			log.Info("adding client certificate", "file", file.Name())
			path := filepath.Join(cfg.Auth.ClientCertsDir, file.Name())
			caCert, err := os.ReadFile(path)
			if err != nil {
				return errors.Wrapf(err, "could not read certificate %q", path)
			}
			if !clientCertPool.AppendCertsFromPEM(caCert) {
				return errors.Errorf("failed to load PEM client certificate from %q", path)
			}
		}
	}
	if cfg.HTTPS.TlsCaFile != "" {
		file, err := os.ReadFile(cfg.HTTPS.TlsCaFile)
		if err != nil {
			return err
		}
		if !clientCertPool.AppendCertsFromPEM(file) {
			return errors.Errorf("failed to load PEM client certificate from %q", cfg.HTTPS.TlsCaFile)
		}
	}

	tlsConfig.ClientCAs = clientCertPool
	if cfg.HTTPS.RequireClientCert {
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
	} else if cfg.Authn.Type == certs.PluginName {
		tlsConfig.ClientAuth = tls.VerifyClientCertIfGiven // client certs are required only for some endpoints when using admin client cert
	}
	return nil
}

func SetupServer(rt runtime.Runtime) error {
	cfg := rt.Config()
	apiServer, err := NewApiServer(
		rt.ResourceManager(),
		xds_context.NewMeshContextBuilder(
			rt.ResourceManager(),
			server.MeshResourceTypes(server.HashMeshExcludedResources),
			net.LookupIP,
			cfg.Multizone.Zone.Name,
			vips.NewPersistence(rt.ResourceManager(), rt.ConfigManager(), cfg.Experimental.UseTagFirstVirtualOutboundModel),
			cfg.DNSServer.Domain,
			cfg.DNSServer.ServiceVipPort,
		),
		rt.APIInstaller(),
		registry.Global().ObjectDescriptors(model.HasWsEnabled()),
		&cfg,
		rt.Metrics(),
		rt.GetInstanceId,
		rt.GetClusterId,
		rt.APIServerAuthenticator(),
		rt.Access(),
		rt.EnvoyAdminClient(),
		rt.TokenIssuers(),
		rt.APIWebServiceCustomize(),
	)
	if err != nil {
		return err
	}
	return rt.Add(apiServer)
}
