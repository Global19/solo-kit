package kubernetes

import (
	"sync"

	"github.com/solo-io/go-utils/errors"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/wrapper"
	"github.com/solo-io/solo-kit/pkg/multicluster/handler"
	"k8s.io/client-go/rest"
)

type PodMultiClusterClient interface {
	handler.ClusterHandler
	PodInterface
}

type podMultiClusterClient struct {
	clients       map[string]PodClient
	clientAccess  sync.RWMutex
	aggregator    wrapper.WatchAggregator
	factoryGetter factory.ResourceClientFactoryGetter
}

func NewPodMultiClusterClient(factoryGetter factory.ResourceClientFactoryGetter) PodMultiClusterClient {
	return NewPodMultiClusterClientWithWatchAggregator(nil, factoryGetter)
}

func NewPodMultiClusterClientWithWatchAggregator(aggregator wrapper.WatchAggregator, factoryGetter factory.ResourceClientFactoryGetter) PodMultiClusterClient {
	return &podMultiClusterClient{
		clients:       make(map[string]PodClient),
		clientAccess:  sync.RWMutex{},
		aggregator:    aggregator,
		factoryGetter: factoryGetter,
	}
}

func (c *podMultiClusterClient) interfaceFor(cluster string) (PodInterface, error) {
	c.clientAccess.RLock()
	defer c.clientAccess.RUnlock()
	if client, ok := c.clients[cluster]; ok {
		return client, nil
	}
	return nil, errors.Errorf("%v.%v client not found for cluster %v", "kubernetes", "Pod", cluster)
}

func (c *podMultiClusterClient) ClusterAdded(cluster string, restConfig *rest.Config) {
	client, err := NewPodClient(c.factoryGetter.ForCluster(cluster, restConfig))
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

func (c *podMultiClusterClient) ClusterRemoved(cluster string, restConfig *rest.Config) {
	c.clientAccess.Lock()
	defer c.clientAccess.Unlock()
	if client, ok := c.clients[cluster]; ok {
		delete(c.clients, cluster)
		if c.aggregator != nil {
			c.aggregator.RemoveWatch(client.BaseClient())
		}
	}
}

func (c *podMultiClusterClient) Read(namespace, name string, opts clients.ReadOpts) (*Pod, error) {
	clusterInterface, err := c.interfaceFor(opts.Cluster)
	if err != nil {
		return nil, err
	}
	return clusterInterface.Read(namespace, name, opts)
}

func (c *podMultiClusterClient) Write(pod *Pod, opts clients.WriteOpts) (*Pod, error) {
	clusterInterface, err := c.interfaceFor(pod.GetMetadata().Cluster)
	if err != nil {
		return nil, err
	}
	return clusterInterface.Write(pod, opts)
}

func (c *podMultiClusterClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	clusterInterface, err := c.interfaceFor(opts.Cluster)
	if err != nil {
		return err
	}
	return clusterInterface.Delete(namespace, name, opts)
}

func (c *podMultiClusterClient) List(namespace string, opts clients.ListOpts) (PodList, error) {
	clusterInterface, err := c.interfaceFor(opts.Cluster)
	if err != nil {
		return nil, err
	}
	return clusterInterface.List(namespace, opts)
}

func (c *podMultiClusterClient) Watch(namespace string, opts clients.WatchOpts) (<-chan PodList, <-chan error, error) {
	clusterInterface, err := c.interfaceFor(opts.Cluster)
	if err != nil {
		return nil, nil, err
	}
	return clusterInterface.Watch(namespace, opts)
}
