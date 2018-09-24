package syncer

import (
	"context"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reporter"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
	"github.com/solo-io/solo-kit/projects/gateway/pkg/api/v1"
	"github.com/solo-io/solo-kit/projects/gateway/pkg/propagator"
	"github.com/solo-io/solo-kit/projects/gateway/pkg/translator"
	gloov1 "github.com/solo-io/solo-kit/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-kit/projects/sqoop/pkg/todo"
)

type translatorSyncer struct {
	writeNamespace  string
	reporter        reporter.Reporter
	propagator      *propagator.Propagator
	writeErrs       chan error
	proxyReconciler gloov1.ProxyReconciler
}

func NewTranslatorSyncer(writeNamespace string, proxyClient gloov1.ProxyClient, reporter reporter.Reporter, propagator *propagator.Propagator, writeErrs chan error) v1.ApiSyncer {
	return &translatorSyncer{
		writeNamespace:  writeNamespace,
		reporter:        reporter,
		propagator:      propagator,
		writeErrs:       writeErrs,
		proxyReconciler: gloov1.NewProxyReconciler(proxyClient),
	}
}

func (s *translatorSyncer) Sync(ctx context.Context, snap *v1.ApiSnapshot) error {
	ctx = contextutils.WithLogger(ctx, "translatorSyncer")

	logger := contextutils.LoggerFrom(ctx)
	logger.Infof("begin sync %v (%v virtual services, %v gateways)", snap.Hash(),
		len(snap.VirtualServices), len(snap.Gateways))
	defer logger.Infof("end sync %v", snap.Hash())
	logger.Debugf("%v", snap)

	proxy, resourceErrs := translator.Translate(ctx, s.writeNamespace, snap)
	reporterErr := s.reporter.WriteReports(ctx, resourceErrs)
	if err := resourceErrs.Validate(); err != nil {
		logger.Warnf("gateway %v was rejected due to invalid config: %v\nxDS cache will not be updated.", err)
		return err
	}

	labels := map[string]string{
		"created_by": "gateway",
	}

	var desiredResources gloov1.ProxyList
	if proxy != nil {
		logger.Infof("creating proxy %v", proxy.Metadata.Ref())
		proxy.Metadata.Labels = labels
		desiredResources = gloov1.ProxyList{proxy}
	}

	if err := s.proxyReconciler.Reconcile(s.writeNamespace, desiredResources, TODO.TransitionFunction, clients.ListOpts{
		Ctx:      ctx,
		Selector: labels,
	}); err != nil {
		return err
	}

	// start propagating for new set of resources
	// TODO(ilackarms): reinstate propagator
	return reporterErr // s.propagator.PropagateStatuses(snap, proxy, clients.WatchOpts{Ctx: ctx})
}