package server

import (
	"context"
	"strings"
	"time"

	envoy_core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/go-logr/logr"
	"golang.org/x/exp/maps"

	"github.com/kumahq/kuma/pkg/core"
	"github.com/kumahq/kuma/pkg/core/resources/model"
	"github.com/kumahq/kuma/pkg/events"
	"github.com/kumahq/kuma/pkg/kds/v2/reconcile"
	"github.com/kumahq/kuma/pkg/multitenant"
	util_maps "github.com/kumahq/kuma/pkg/util/maps"
	util_watchdog "github.com/kumahq/kuma/pkg/util/watchdog"
)

type EventBasedWatchdog struct {
	Ctx                 context.Context
	Node                *envoy_core.Node
	Listener            events.Listener
	Reconciler          reconcile.Reconciler
	ProvidedTypes       map[model.ResourceType]struct{}
	Metrics             *Metrics
	Log                 logr.Logger
	NewFlushTicker      func() *time.Ticker
	NewFullResyncTicker func() *time.Ticker
}

var _ util_watchdog.Watchdog = &EventBasedWatchdog{}

func (e *EventBasedWatchdog) Start(stop <-chan struct{}) {
	tenantID, _ := multitenant.TenantFromCtx(e.Ctx)
	flushTicker := e.NewFlushTicker()
	defer flushTicker.Stop()
	fullResyncTicker := e.NewFullResyncTicker()
	defer fullResyncTicker.Stop()

	// for the first reconcile assign all types
	changedTypes := maps.Clone(e.ProvidedTypes)
	reasons := map[string]struct{}{
		ReasonResync: {},
	}

	for {
		select {
		case <-stop:
			if err := e.Reconciler.Clear(e.Ctx, e.Node); err != nil {
				e.Log.Error(err, "reconcile clear failed")
			}
			e.Listener.Close()
			return
		case <-flushTicker.C:
			if len(changedTypes) == 0 {
				continue
			}
			reason := strings.Join(util_maps.SortedKeys(reasons), "&")
			e.Log.V(1).Info("reconcile", "changedTypes", changedTypes, "reason", reason)
			start := core.Now()
			err, changed := e.Reconciler.Reconcile(e.Ctx, e.Node, changedTypes)
			if err != nil {
				e.Log.Error(err, "reconcile failed", "changedTypes", changedTypes, "reason", reason)
				e.Metrics.KdsGenerationsErrors.Inc()
			} else {
				result := ResultNoChanges
				if changed {
					result = ResultChanged
				}
				// we want to combine reason. One of the reasons we introduce this metric is to check if we need full resync
				// If we just keep a single reason, we might get into races where full resync ticker runs,
				// then listener, and we would lose information what triggered flush.
				e.Metrics.KdsGenerations.WithLabelValues(reason, result).Observe(float64(core.Now().Sub(start).Milliseconds()))
				changedTypes = map[model.ResourceType]struct{}{}
				reasons = map[string]struct{}{}
			}
		case <-fullResyncTicker.C:
			e.Log.V(1).Info("schedule full resync")
			changedTypes = maps.Clone(e.ProvidedTypes)
			reasons[ReasonResync] = struct{}{}
		case event := <-e.Listener.Recv():
			resChange, ok := event.(events.ResourceChangedEvent)
			if !ok {
				continue
			}
			if resChange.TenantID != tenantID {
				continue
			}
			if _, ok := e.ProvidedTypes[resChange.Type]; !ok {
				continue
			}
			e.Log.V(1).Info("schedule sync for type", "typ", resChange.Type)
			changedTypes[resChange.Type] = struct{}{}
			reasons[ReasonEvent] = struct{}{}
		}
	}
}
