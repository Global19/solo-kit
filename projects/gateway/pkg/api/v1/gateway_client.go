package v1

import (
	"sort"

	"github.com/gogo/protobuf/proto"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TODO: modify as needed to populate additional fields
func NewGateway(namespace, name string) *Gateway {
	return &Gateway{
		Metadata: core.Metadata{
			Name:      name,
			Namespace: namespace,
		},
	}
}

func (r *Gateway) SetStatus(status core.Status) {
	r.Status = status
}

func (r *Gateway) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

type GatewayList []*Gateway
type GatewaysByNamespace map[string]GatewayList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list GatewayList) Find(namespace, name string) (*Gateway, error) {
	for _, gateway := range list {
		if gateway.Metadata.Name == name {
			if namespace == "" || gateway.Metadata.Namespace == namespace {
				return gateway, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find gateway %v.%v", namespace, name)
}

func (list GatewayList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, gateway := range list {
		ress = append(ress, gateway)
	}
	return ress
}

func (list GatewayList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, gateway := range list {
		ress = append(ress, gateway)
	}
	return ress
}

func (list GatewayList) Names() []string {
	var names []string
	for _, gateway := range list {
		names = append(names, gateway.Metadata.Name)
	}
	return names
}

func (list GatewayList) NamespacesDotNames() []string {
	var names []string
	for _, gateway := range list {
		names = append(names, gateway.Metadata.Namespace+"."+gateway.Metadata.Name)
	}
	return names
}

func (list GatewayList) Sort() {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Metadata.Less(list[j].Metadata)
	})
}

func (list GatewayList) Clone() GatewayList {
	var gatewayList GatewayList
	for _, gateway := range list {
		gatewayList = append(gatewayList, proto.Clone(gateway).(*Gateway))
	}
	return gatewayList
}

func (list GatewayList) ByNamespace() GatewaysByNamespace {
	byNamespace := make(GatewaysByNamespace)
	for _, gateway := range list {
		byNamespace.Add(gateway)
	}
	return byNamespace
}

func (byNamespace GatewaysByNamespace) Add(gateway ...*Gateway) {
	for _, item := range gateway {
		byNamespace[item.Metadata.Namespace] = append(byNamespace[item.Metadata.Namespace], item)
	}
}

func (byNamespace GatewaysByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace GatewaysByNamespace) List() GatewayList {
	var list GatewayList
	for _, gatewayList := range byNamespace {
		list = append(list, gatewayList...)
	}
	list.Sort()
	return list
}

func (byNamespace GatewaysByNamespace) Clone() GatewaysByNamespace {
	return byNamespace.List().Clone().ByNamespace()
}

var _ resources.Resource = &Gateway{}

type GatewayClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*Gateway, error)
	Write(resource *Gateway, opts clients.WriteOpts) (*Gateway, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (GatewayList, error)
	Watch(namespace string, opts clients.WatchOpts) (<-chan GatewayList, <-chan error, error)
}

type gatewayClient struct {
	rc clients.ResourceClient
}

func NewGatewayClient(rcFactory factory.ResourceClientFactory) (GatewayClient, error) {
	rc, err := rcFactory.NewResourceClient(factory.NewResourceClientParams{
		ResourceType: &Gateway{},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base Gateway resource client")
	}
	return &gatewayClient{
		rc: rc,
	}, nil
}

func (client *gatewayClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *gatewayClient) Register() error {
	return client.rc.Register()
}

func (client *gatewayClient) Read(namespace, name string, opts clients.ReadOpts) (*Gateway, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Gateway), nil
}

func (client *gatewayClient) Write(gateway *Gateway, opts clients.WriteOpts) (*Gateway, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(gateway, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*Gateway), nil
}

func (client *gatewayClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()
	return client.rc.Delete(namespace, name, opts)
}

func (client *gatewayClient) List(namespace string, opts clients.ListOpts) (GatewayList, error) {
	opts = opts.WithDefaults()
	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToGateway(resourceList), nil
}

func (client *gatewayClient) Watch(namespace string, opts clients.WatchOpts) (<-chan GatewayList, <-chan error, error) {
	opts = opts.WithDefaults()
	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	gatewaysChan := make(chan GatewayList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				gatewaysChan <- convertToGateway(resourceList)
			case <-opts.Ctx.Done():
				close(gatewaysChan)
				return
			}
		}
	}()
	return gatewaysChan, errs, nil
}

func convertToGateway(resources resources.ResourceList) GatewayList {
	var gatewayList GatewayList
	for _, resource := range resources {
		gatewayList = append(gatewayList, resource.(*Gateway))
	}
	return gatewayList
}

// Kubernetes Adapter for Gateway

func (o *Gateway) GetObjectKind() schema.ObjectKind {
	t := GatewayCrd.TypeMeta()
	return &t
}

func (o *Gateway) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*Gateway)
}

var GatewayCrd = crd.NewCrd("gateway.solo.io",
	"gateways",
	"gateway.solo.io",
	"v1",
	"Gateway",
	"gw",
	&Gateway{})
