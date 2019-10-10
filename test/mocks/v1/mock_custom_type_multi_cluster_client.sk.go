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

type MockCustomTypeMultiClusterClient interface {
	handler.ClusterHandler
	MockCustomTypeInterface
}

type mockCustomTypeMultiClusterClient struct {
	clients       map[string]MockCustomTypeClient
	clientAccess  sync.RWMutex
	aggregator    wrapper.WatchAggregator
	factoryGetter factory.ResourceClientFactoryGetter
}

func NewMockCustomTypeMultiClusterClient(factoryGetter factory.ResourceClientFactoryGetter) MockCustomTypeMultiClusterClient {
	return NewMockCustomTypeMultiClusterClientWithWatchAggregator(nil, factoryGetter)
}

func NewMockCustomTypeMultiClusterClientWithWatchAggregator(aggregator wrapper.WatchAggregator, factoryGetter factory.ResourceClientFactoryGetter) MockCustomTypeMultiClusterClient {
	return &mockCustomTypeMultiClusterClient{
		clients:       make(map[string]MockCustomTypeClient),
		clientAccess:  sync.RWMutex{},
		aggregator:    aggregator,
		factoryGetter: factoryGetter,
	}
}

func (c *mockCustomTypeMultiClusterClient) interfaceFor(cluster string) (MockCustomTypeInterface, error) {
	c.clientAccess.RLock()
	defer c.clientAccess.RUnlock()
	if client, ok := c.clients[cluster]; ok {
		return client, nil
	}
	return nil, errors.Errorf("%v.%v client not found for cluster %v", "v1", "MockCustomType", cluster)
}

func (c *mockCustomTypeMultiClusterClient) ClusterAdded(cluster string, restConfig *rest.Config) {
	client, err := NewMockCustomTypeClient(c.factoryGetter.ForCluster(cluster, restConfig))
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

func (c *mockCustomTypeMultiClusterClient) ClusterRemoved(cluster string, restConfig *rest.Config) {
	c.clientAccess.Lock()
	defer c.clientAccess.Unlock()
	if client, ok := c.clients[cluster]; ok {
		delete(c.clients, cluster)
		if c.aggregator != nil {
			c.aggregator.RemoveWatch(client.BaseClient())
		}
	}
}

func (c *mockCustomTypeMultiClusterClient) Read(namespace, name string, opts clients.ReadOpts) (*MockCustomType, error) {
	clusterInterface, err := c.interfaceFor(opts.Cluster)
	if err != nil {
		return nil, err
	}
	return clusterInterface.Read(namespace, name, opts)
}

func (c *mockCustomTypeMultiClusterClient) Write(mockCustomType *MockCustomType, opts clients.WriteOpts) (*MockCustomType, error) {
	clusterInterface, err := c.interfaceFor(mockCustomType.GetMetadata().Cluster)
	if err != nil {
		return nil, err
	}
	return clusterInterface.Write(mockCustomType, opts)
}

func (c *mockCustomTypeMultiClusterClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	clusterInterface, err := c.interfaceFor(opts.Cluster)
	if err != nil {
		return err
	}
	return clusterInterface.Delete(namespace, name, opts)
}

func (c *mockCustomTypeMultiClusterClient) List(namespace string, opts clients.ListOpts) (MockCustomTypeList, error) {
	clusterInterface, err := c.interfaceFor(opts.Cluster)
	if err != nil {
		return nil, err
	}
	return clusterInterface.List(namespace, opts)
}

func (c *mockCustomTypeMultiClusterClient) Watch(namespace string, opts clients.WatchOpts) (<-chan MockCustomTypeList, <-chan error, error) {
	clusterInterface, err := c.interfaceFor(opts.Cluster)
	if err != nil {
		return nil, nil, err
	}
	return clusterInterface.Watch(namespace, opts)
}
