package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
)

// Option to copy anything from the original to the desired before writing
type TransitionProxyFunc func(original, desired *Proxy) error

type ProxyReconciler interface {
	Reconcile(namespace string, desiredResources []*Proxy, transition TransitionProxyFunc, opts clients.ListOpts) error
}

func proxysToResources(list ProxyList) []resources.Resource {
	var resourceList []resources.Resource
	for _, proxy := range list {
		resourceList = append(resourceList, proxy)
	}
	return resourceList
}

func NewProxyReconciler(client ProxyClient) ProxyReconciler {
	return &proxyReconciler{
		base: reconcile.NewReconciler(client.BaseClient()),
	}
}

type proxyReconciler struct {
	base reconcile.Reconciler
}

func (r *proxyReconciler) Reconcile(namespace string, desiredResources []*Proxy, transition TransitionProxyFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "proxy_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) error {
			return transition(original.(*Proxy), desired.(*Proxy))
		}
	}
	return r.base.Reconcile(namespace, proxysToResources(desiredResources), transitionResources, opts)
}