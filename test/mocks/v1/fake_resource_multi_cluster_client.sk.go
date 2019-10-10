// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"sync"

	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/wrapper"
	"github.com/solo-io/solo-kit/pkg/multicluster/handler"
	"k8s.io/client-go/rest"
)

type FakeResourceMultiClusterClient interface {
	handler.ClusterHandler
	FakeResourceInterface
}

type fakeResourceMultiClusterClient struct {
	clients       map[string]FakeResourceClient
	clientAccess  sync.RWMutex
	aggregator    wrapper.WatchAggregator
	factoryGetter factory.ResourceClientFactoryGetter
}

func NewFakeResourceMultiClusterClient(factoryGetter factory.ResourceClientFactoryGetter) FakeResourceMultiClusterClient {
	return NewFakeResourceMultiClusterClientWithWatchAggregator(nil, factoryGetter)
}

func NewFakeResourceMultiClusterClientWithWatchAggregator(aggregator wrapper.WatchAggregator, factoryGetter factory.ResourceClientFactoryGetter) FakeResourceMultiClusterClient {
	return &fakeResourceMultiClusterClient{
		clients:       make(map[string]FakeResourceClient),
		clientAccess:  sync.RWMutex{},
		aggregator:    aggregator,
		factoryGetter: factoryGetter,
	}
}

func (c *fakeResourceMultiClusterClient) interfaceFor(cluster string) (FakeResourceInterface, error) {
	c.clientAccess.RLock()
	defer c.clientAccess.RUnlock()
	if client, ok := c.clients[cluster]; ok {
		return client, nil
	}
	return nil, errors.Errorf("%v.%v client not found for cluster %v", "v1", "FakeResource", cluster)
}

func (c *fakeResourceMultiClusterClient) ClusterAdded(cluster string, restConfig *rest.Config) {
	client, err := NewFakeResourceClient(c.factoryGetter.ForCluster(cluster, restConfig))
	if err != nil {
		return
	}
	if err := client.Register(); err != nil {
		return
	}
	c.clientAccess.Lock()
	defer c.clientAccess.Unlock()
	c.clients[cluster] = client
	if c.aggregator != nil {
		c.aggregator.AddWatch(client.BaseClient())
	}
}

func (c *fakeResourceMultiClusterClient) ClusterRemoved(cluster string, restConfig *rest.Config) {
	c.clientAccess.Lock()
	defer c.clientAccess.Unlock()
	if client, ok := c.clients[cluster]; ok {
		delete(c.clients, cluster)
		if c.aggregator != nil {
			c.aggregator.RemoveWatch(client.BaseClient())
		}
	}
}

func (c *fakeResourceMultiClusterClient) Read(namespace, name string, opts clients.ReadOpts) (*FakeResource, error) {
	clusterInterface, err := c.interfaceFor(opts.Cluster)
	if err != nil {
		return nil, err
	}
	return clusterInterface.Read(namespace, name, opts)
}

func (c *fakeResourceMultiClusterClient) Write(fakeResource *FakeResource, opts clients.WriteOpts) (*FakeResource, error) {
	clusterInterface, err := c.interfaceFor(fakeResource.GetMetadata().Cluster)
	if err != nil {
		return nil, err
	}
	return clusterInterface.Write(fakeResource, opts)
}

func (c *fakeResourceMultiClusterClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	clusterInterface, err := c.interfaceFor(opts.Cluster)
	if err != nil {
		return err
	}
	return clusterInterface.Delete(namespace, name, opts)
}

func (c *fakeResourceMultiClusterClient) List(namespace string, opts clients.ListOpts) (FakeResourceList, error) {
	clusterInterface, err := c.interfaceFor(opts.Cluster)
	if err != nil {
		return nil, err
	}
	return clusterInterface.List(namespace, opts)
}

func (c *fakeResourceMultiClusterClient) Watch(namespace string, opts clients.WatchOpts) (<-chan FakeResourceList, <-chan error, error) {
	clusterInterface, err := c.interfaceFor(opts.Cluster)
	if err != nil {
		return nil, nil, err
	}
	return clusterInterface.Watch(namespace, opts)
}
