package server

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	envoy_types "github.com/envoyproxy/go-control-plane/pkg/cache/types"
	envoy_cache "github.com/envoyproxy/go-control-plane/pkg/cache/v2"
	envoy_resource "github.com/envoyproxy/go-control-plane/pkg/resource/v2"

	"github.com/Kong/kuma/pkg/core"
	"github.com/Kong/kuma/pkg/core/ca/issuer"
	mesh_core "github.com/Kong/kuma/pkg/core/resources/apis/mesh"
	core_manager "github.com/Kong/kuma/pkg/core/resources/manager"
	core_model "github.com/Kong/kuma/pkg/core/resources/model"
	core_store "github.com/Kong/kuma/pkg/core/resources/store"
	core_xds "github.com/Kong/kuma/pkg/core/xds"
	sds_auth "github.com/Kong/kuma/pkg/sds/auth"
	sds_provider "github.com/Kong/kuma/pkg/sds/provider"
	util_proto "github.com/Kong/kuma/pkg/util/proto"
)

// DataplaneReconciler keeps the state of the Cache for SDS consistent
// When Dataplane connects to the Control Plane, the Watchdog (separate goroutine) is started which on the defined interval
// execute DataplaneReconciler#Reconcile. It will then check if certs needs to be regenerated because Mesh CA was changed
// This follows the same pattern as XDS.
//
// Snapshot are versioned with UnixNano;NameOfTheCA pattern
type DataplaneReconciler struct {
	resManager       core_manager.ReadOnlyResourceManager
	meshCaProvider   sds_provider.SecretProvider
	identityProvider sds_provider.SecretProvider
	cache            envoy_cache.SnapshotCache
}

func (d *DataplaneReconciler) Reconcile(dataplaneId core_model.ResourceKey) error {
	proxyID := core_xds.FromResourceKey(dataplaneId).String()

	dataplane := &mesh_core.DataplaneResource{}
	if err := d.resManager.Get(context.Background(), dataplane, core_store.GetBy(dataplaneId)); err != nil {
		if core_store.IsResourceNotFound(err) {
			sdsServerLog.V(1).Info("Dataplane not found. Clearing the Snapshot.", "dataplaneId", dataplaneId)
			d.cache.ClearSnapshot(proxyID)
		}
		return err
	}

	meshRes := mesh_core.MeshResource{}
	if err := d.resManager.Get(context.Background(), &meshRes, core_store.GetByKey(dataplane.GetMeta().GetMesh(), dataplane.GetMeta().GetMesh())); err != nil {
		return errors.Wrap(err, "could not retrieve a mesh")
	}

	if !meshRes.MTLSEnabled() {
		sdsServerLog.V(1).Info("mTLS for Mesh disabled. Clearing the Snapshot.", "dataplaneId", dataplaneId)
		d.cache.ClearSnapshot(proxyID)
		return nil
	}

	generateSnapshot, reason, err := d.shouldGenerateSnapshot(proxyID, meshRes)
	if err != nil {
		return err
	}

	if generateSnapshot {
		sdsServerLog.Info("Generating the Snapshot.", "dataplaneId", dataplaneId, "reason", reason)
		snapshot, err := d.generateSnapshot(dataplane, meshRes)
		if err != nil {
			return err
		}
		if err := d.cache.SetSnapshot(proxyID, snapshot); err != nil {
			return err
		}
	}
	return nil
}

func (d *DataplaneReconciler) shouldGenerateSnapshot(proxyID string, meshRes mesh_core.MeshResource) (bool, string, error) {
	currentSnapshot, err := d.cache.GetSnapshot(proxyID)
	if err != nil {
		return true, "Snapshot does not exist", nil
	}

	parts := strings.Split(currentSnapshot.GetVersion(envoy_resource.SecretType), ";")
	if len(parts) != 2 {
		return false, "", errors.New(`invalid snapshot version format. Format should be "UnixNano-NameOfTheCA"`)
	}
	// generate snapshot if CA changed
	caName := parts[1]
	if caName != meshRes.GetEnabledCertificateAuthorityBackend().Name {
		return true, fmt.Sprintf("Enabled CA changed from %s to %s", caName, meshRes.GetEnabledCertificateAuthorityBackend().Name), nil
	}
	// generate snapshot if cert expired
	generationUnixNano, err := strconv.Atoi(parts[0])
	if err != nil {
		return false, "", errors.Wrap(err, `invalid snapshot version format. Format should be "UnixNano;NameOfTheCA"`)
	}
	expiration := issuer.DefaultWorkloadCertValidityPeriod
	if meshRes.GetEnabledCertificateAuthorityBackend().GetDpCert().GetRotation().GetExpiration() != nil {
		expiration = util_proto.ToDuration(*meshRes.GetEnabledCertificateAuthorityBackend().GetDpCert().GetRotation().GetExpiration())
	}
	generationTime := time.Unix(0, int64(generationUnixNano))
	expirationTime := generationTime.Add(expiration)
	if core.Now().After(generationTime.Add(expiration / 5 * 4)) { // regenerate cert after 4/5 of its lifetime
		reason := fmt.Sprintf("Certificate generated at %s will expire in %s", generationTime, expirationTime.Sub(core.Now()))
		return true, reason, nil
	}
	return false, "", nil
}

func (d *DataplaneReconciler) generateSnapshot(dataplane *mesh_core.DataplaneResource, meshRes mesh_core.MeshResource) (envoy_cache.Snapshot, error) {
	requestor := sds_auth.Identity{
		Service: dataplane.Spec.GetIdentifyingService(),
		Mesh:    dataplane.GetMeta().GetMesh(),
	}
	identitySecret, err := d.identityProvider.Get(context.Background(), IdentityCertResource, requestor)
	if err != nil {
		return envoy_cache.Snapshot{}, errors.Wrap(err, "could not get Dataplane cert pair")
	}

	requestor = sds_auth.Identity{
		Mesh: dataplane.GetMeta().GetMesh(),
	}
	caSecret, err := d.meshCaProvider.Get(context.Background(), MeshCaResource, requestor)
	if err != nil {
		return envoy_cache.Snapshot{}, errors.Wrap(err, "could not get mesh CA cert")
	}

	version := fmt.Sprintf("%d;%s", core.Now().UTC().UnixNano(), meshRes.GetEnabledCertificateAuthorityBackend().Name)
	snap := envoy_cache.Snapshot{
		Resources: [envoy_types.UnknownType]envoy_cache.Resources{},
	}
	snap.Resources[envoy_types.Secret] = envoy_cache.NewResources(version, []envoy_types.Resource{
		identitySecret.ToResource(IdentityCertResource),
		caSecret.ToResource(MeshCaResource),
	})
	return snap, nil
}
