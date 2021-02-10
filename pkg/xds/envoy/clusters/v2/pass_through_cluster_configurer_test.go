package clusters_test

import (
	"time"

	mesh_proto "github.com/kumahq/kuma/api/mesh/v1alpha1"
	mesh_core "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/protobuf/types/known/durationpb"

	util_proto "github.com/kumahq/kuma/pkg/util/proto"
	"github.com/kumahq/kuma/pkg/xds/envoy"
	"github.com/kumahq/kuma/pkg/xds/envoy/clusters"
)

var _ = Describe("PassThroughClusterConfigurer", func() {

	It("should generate proper Envoy config", func() {
		// given
		clusterName := "test:cluster"
		expected := `
        altStatName: test_cluster
        connectTimeout: 5s
        lbPolicy: CLUSTER_PROVIDED
        name: test:cluster
        type: ORIGINAL_DST`

		// when
		cluster, err := clusters.NewClusterBuilder(envoy.APIV2).
			Configure(clusters.PassThroughCluster(clusterName)).
			Configure(clusters.Timeout(mesh_core.ProtocolTCP, &mesh_proto.Timeout_Conf{ConnectTimeout: durationpb.New(5 * time.Second)})).
			Build()

		// then
		Expect(err).ToNot(HaveOccurred())

		actual, err := util_proto.ToYAML(cluster)
		Expect(err).ToNot(HaveOccurred())
		Expect(actual).To(MatchYAML(expected))
	})
})
