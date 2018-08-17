package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
)

// Option to copy anything from the original to the desired before writing
type TransitionGatewayFunc func(original, desired *Gateway) error

type GatewayReconciler interface {
	Reconcile(namespace string, desiredResources []*Gateway, transition TransitionGatewayFunc, opts clients.ListOpts) error
}

func gatewaysToResources(list GatewayList) []resources.Resource {
	var resourceList []resources.Resource
	for _, gateway := range list {
		resourceList = append(resourceList, gateway)
	}
	return resourceList
}

func NewGatewayReconciler(client GatewayClient) GatewayReconciler {
	return &gatewayReconciler{
		base: reconcile.NewReconciler(client.BaseClient()),
	}
}

type gatewayReconciler struct {
	base reconcile.Reconciler
}

func (r *gatewayReconciler) Reconcile(namespace string, desiredResources []*Gateway, transition TransitionGatewayFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "gateway_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) error {
			return transition(original.(*Gateway), desired.(*Gateway))
		}
	}
	return r.base.Reconcile(namespace, gatewaysToResources(desiredResources), transitionResources, opts)
}