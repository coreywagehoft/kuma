package server

import (
	"context"
	"time"

	envoy_core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/go-logr/logr"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	core_model "github.com/kumahq/kuma/pkg/core/resources/model"
	"github.com/kumahq/kuma/pkg/events"
	reconcile_v2 "github.com/kumahq/kuma/pkg/kds/v2/reconcile"
	core_metrics "github.com/kumahq/kuma/pkg/metrics"
	test_metrics "github.com/kumahq/kuma/pkg/test/metrics"
)

type staticReconciler struct {
	changedResTypes chan map[core_model.ResourceType]struct{}
}

func (s staticReconciler) Reconcile(ctx context.Context, node *envoy_core.Node, m map[core_model.ResourceType]struct{}) (error, bool) {
	s.changedResTypes <- m
	return nil, true
}

func (s staticReconciler) Clear(ctx context.Context, node *envoy_core.Node) error {
	return nil
}

var _ reconcile_v2.Reconciler = &staticReconciler{}

var _ = Describe("Event Based Watchdog", func() {
	var eventBus events.EventBus
	var metrics core_metrics.Metrics
	var reconciler *staticReconciler
	var stopCh chan struct{}
	var flushCh chan time.Time
	var fullResyncCh chan time.Time
	var watchdog *EventBasedWatchdog

	BeforeAll(func() {
		var err error
		metrics, err = core_metrics.NewMetrics("")
		Expect(err).ToNot(HaveOccurred())
		kdsMetrics, err := NewMetrics(metrics)
		Expect(err).ToNot(HaveOccurred())

		eventBus = events.NewEventBus()

		stopCh = make(chan struct{})
		flushCh = make(chan time.Time)
		fullResyncCh = make(chan time.Time)
		reconciler = &staticReconciler{
			changedResTypes: make(chan map[core_model.ResourceType]struct{}, 1),
		}
		watchdog = &EventBasedWatchdog{
			Ctx:        context.Background(),
			Node:       nil,
			Listener:   eventBus.Subscribe(),
			Reconciler: reconciler,
			ProvidedTypes: map[core_model.ResourceType]struct{}{
				mesh.TrafficPermissionType: {},
				mesh.TrafficLogType:        {},
				mesh.TrafficRouteType:      {},
			},
			Metrics: kdsMetrics,
			Log:     logr.Discard(),
			NewFlushTicker: func() *time.Ticker {
				return &time.Ticker{C: flushCh}
			},
			NewFullResyncTicker: func() *time.Ticker {
				return &time.Ticker{C: fullResyncCh}
			},
		}
		go func() {
			watchdog.Start(stopCh)
		}()
	})

	AfterAll(func() {
		close(stopCh)
	})

	It("should reconcile on the first flush", func() {
		// when
		flushCh <- time.Now()

		// then
		changedResTypes := <-reconciler.changedResTypes
		Expect(changedResTypes).To(Equal(watchdog.ProvidedTypes))
		Eventually(func(g Gomega) {
			metric := test_metrics.FindMetric(metrics, "kds_delta_generation", "reason", ReasonResync)
			g.Expect(metric).ToNot(BeNil())
			g.Expect(*metric.Summary.SampleCount).To(BeEquivalentTo(1))
		}, "10s", "50ms").Should(Succeed())
	})

	It("should reconcile on the events flush", func() {
		// when
		eventBus.Send(events.ResourceChangedEvent{
			Type: mesh.TrafficPermissionType,
		})
		eventBus.Send(events.ResourceChangedEvent{
			Type: mesh.TrafficLogType,
		})
		// Send is not blocking so there is no guarantee that we execute flush before watchdog consumed events
		time.Sleep(500 * time.Millisecond)
		flushCh <- time.Now()

		// then
		changedResTypes := <-reconciler.changedResTypes
		Expect(changedResTypes).To(HaveLen(2))
		Expect(changedResTypes).To(HaveKey(mesh.TrafficPermissionType))
		Expect(changedResTypes).To(HaveKey(mesh.TrafficLogType))
		Eventually(func(g Gomega) {
			metric := test_metrics.FindMetric(metrics, "kds_delta_generation", "reason", ReasonEvent)
			g.Expect(metric).ToNot(BeNil())
			g.Expect(*metric.Summary.SampleCount).To(BeEquivalentTo(1))
		}, "10s", "50ms").Should(Succeed())
	})

	It("should reconcile on the full resync", func() {
		// when
		fullResyncCh <- time.Now()
		flushCh <- time.Now()

		// then
		changedResTypes := <-reconciler.changedResTypes
		Expect(changedResTypes).To(Equal(watchdog.ProvidedTypes))
		Eventually(func(g Gomega) {
			metric := test_metrics.FindMetric(metrics, "kds_delta_generation", "reason", ReasonResync)
			g.Expect(metric).ToNot(BeNil())
			g.Expect(*metric.Summary.SampleCount).To(BeEquivalentTo(2))
		}, "10s", "50ms").Should(Succeed())
	})
}, Ordered)
