package v1

import (
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
)

// Option to copy anything from the original to the desired before writing
type TransitionSchemaFunc func(original, desired *Schema) error

type SchemaReconciler interface {
	Reconcile(namespace string, desiredResources SchemaList, transition TransitionSchemaFunc, opts clients.ListOpts) error
}

func schemasToResources(list SchemaList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, schema := range list {
		resourceList = append(resourceList, schema)
	}
	return resourceList
}

func NewSchemaReconciler(client SchemaClient) SchemaReconciler {
	return &schemaReconciler{
		base: reconcile.NewReconciler(client.BaseClient()),
	}
}

type schemaReconciler struct {
	base reconcile.Reconciler
}

func (r *schemaReconciler) Reconcile(namespace string, desiredResources SchemaList, transition TransitionSchemaFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "schema_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) error {
			return transition(original.(*Schema), desired.(*Schema))
		}
	}
	return r.base.Reconcile(namespace, schemasToResources(desiredResources), transitionResources, opts)
}