// Code generated by solo-kit. DO NOT EDIT.

// +build solokit

package v1

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/wrapper"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-kit/test/tests/typed"
)

var _ = Describe("ClusterResourceMultiClusterClient", func() {
	for _, test := range []typed.ResourceClientTester{
		&typed.KubeRcTester{Crd: ClusterResourceCrd},
	} {
		Context("multi cluster resource client backed by "+test.Description(), func() {
			var (
				client              ClusterResourceMultiClusterClient
				name1, name2, name3 = "foo" + helpers.RandString(3), "boo" + helpers.RandString(3), "goo" + helpers.RandString(3)
			)

			BeforeEach(func() {
				test.Setup("")
			})
			AfterEach(func() {
				client.Delete(name1, clients.DeleteOpts{})
				client.Delete(name2, clients.DeleteOpts{})
				client.Delete(name3, clients.DeleteOpts{})
			})
			It("CRUDs ClusterResources "+test.Description(), func() {
				client = NewClusterResourceMultiClusterClient(test)
				ClusterResourceMultiClusterClientTest(client, name1, name2, name3)
			})
			It("errors when no client exists for the given cluster "+test.Description(), func() {
				client = NewClusterResourceMultiClusterClient(test)
				ClusterResourceMultiClusterClientCrudErrorsTest(client)
			})
			It("populates an aggregated watch "+test.Description(), func() {
				watchAggregator := wrapper.NewWatchAggregator()
				client = NewClusterResourceMultiClusterClientWithWatchAggregator(watchAggregator, test)
				ClusterResourceMultiClusterClientWatchAggregationTest(client, watchAggregator)
			})
		})
	}
})

func ClusterResourceMultiClusterClientTest(client ClusterResourceMultiClusterClient, name1, name2, name3 string) {
	cfg, err := kubeutils.GetConfig("", "")
	Expect(err).NotTo(HaveOccurred())
	client.ClusterAdded("", cfg)

	name := name1
	input := NewClusterResource("", name)

	r1, err := client.Write(input, clients.WriteOpts{})
	Expect(err).NotTo(HaveOccurred())

	_, err = client.Write(input, clients.WriteOpts{})
	Expect(err).To(HaveOccurred())
	Expect(errors.IsExist(err)).To(BeTrue())

	Expect(r1).To(BeAssignableToTypeOf(&ClusterResource{}))
	Expect(r1.GetMetadata().Name).To(Equal(name))
	Expect(r1.GetMetadata().ResourceVersion).NotTo(Equal(input.GetMetadata().ResourceVersion))
	Expect(r1.GetMetadata().Ref()).To(Equal(input.GetMetadata().Ref()))
	Expect(r1.Status).To(Equal(input.Status))
	Expect(r1.BasicField).To(Equal(input.BasicField))

	_, err = client.Write(input, clients.WriteOpts{
		OverwriteExisting: true,
	})
	Expect(err).To(HaveOccurred())

	resources.UpdateMetadata(input, func(meta *core.Metadata) {
		meta.ResourceVersion = r1.GetMetadata().ResourceVersion
	})
	r1, err = client.Write(input, clients.WriteOpts{
		OverwriteExisting: true,
	})
	Expect(err).NotTo(HaveOccurred())
	read, err := client.Read(name, clients.ReadOpts{})
	Expect(err).NotTo(HaveOccurred())
	Expect(read).To(Equal(r1))

	name = name2
	input = &ClusterResource{}

	input.SetMetadata(core.Metadata{
		Name: name,
	})

	r2, err := client.Write(input, clients.WriteOpts{})
	Expect(err).NotTo(HaveOccurred())
	list, err := client.List(clients.ListOpts{})
	Expect(err).NotTo(HaveOccurred())
	Expect(list).To(ContainElement(r1))
	Expect(list).To(ContainElement(r2))
	err = client.Delete("adsfw", clients.DeleteOpts{})
	Expect(err).To(HaveOccurred())
	Expect(errors.IsNotExist(err)).To(BeTrue())
	err = client.Delete("adsfw", clients.DeleteOpts{
		IgnoreNotExist: true,
	})
	Expect(err).NotTo(HaveOccurred())
	err = client.Delete(r2.GetMetadata().Name, clients.DeleteOpts{})
	Expect(err).NotTo(HaveOccurred())

	Eventually(func() ClusterResourceList {
		list, err = client.List(clients.ListOpts{})
		Expect(err).NotTo(HaveOccurred())
		return list
	}, time.Second*10).Should(ContainElement(r1))
	Eventually(func() ClusterResourceList {
		list, err = client.List(clients.ListOpts{})
		Expect(err).NotTo(HaveOccurred())
		return list
	}, time.Second*10).ShouldNot(ContainElement(r2))
	w, errs, err := client.Watch(clients.WatchOpts{
		RefreshRate: time.Hour,
	})
	Expect(err).NotTo(HaveOccurred())

	var r3 resources.Resource
	wait := make(chan struct{})
	go func() {
		defer close(wait)
		defer GinkgoRecover()

		resources.UpdateMetadata(r2, func(meta *core.Metadata) {
			meta.ResourceVersion = ""
		})
		r2, err = client.Write(r2, clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())

		name = name3
		input = &ClusterResource{}
		Expect(err).NotTo(HaveOccurred())
		input.SetMetadata(core.Metadata{
			Name: name,
		})

		r3, err = client.Write(input, clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())
	}()
	<-wait

	select {
	case err := <-errs:
		Expect(err).NotTo(HaveOccurred())
	case list = <-w:
	case <-time.After(time.Millisecond * 5):
		Fail("expected a message in channel")
	}

	go func() {
		defer GinkgoRecover()
		for {
			select {
			case err := <-errs:
				Expect(err).NotTo(HaveOccurred())
			case <-time.After(time.Second / 4):
				return
			}
		}
	}()

	Eventually(w, time.Second*5, time.Second/10).Should(Receive(And(ContainElement(r1), ContainElement(r3), ContainElement(r3))))
}
func ClusterResourceMultiClusterClientCrudErrorsTest(client ClusterResourceMultiClusterClient) {
	_, err := client.Read("bar", clients.ReadOpts{Cluster: "read"})
	Expect(err).To(HaveOccurred())
	Expect(err.Error()).To(Equal(NoClusterResourceClientForClusterError("read").Error()))
	_, err = client.List(clients.ListOpts{Cluster: "list"})
	Expect(err).To(HaveOccurred())
	Expect(err.Error()).To(Equal(NoClusterResourceClientForClusterError("list").Error()))
	err = client.Delete("bar", clients.DeleteOpts{Cluster: "delete"})
	Expect(err).To(HaveOccurred())
	Expect(err.Error()).To(Equal(NoClusterResourceClientForClusterError("delete").Error()))

	input := &ClusterResource{}
	input.SetMetadata(core.Metadata{
		Cluster: "write",
		Name:    "bar",
	})
	_, err = client.Write(input, clients.WriteOpts{})
	Expect(err).To(HaveOccurred())
	Expect(err.Error()).To(Equal(NoClusterResourceClientForClusterError("write").Error()))
	_, _, err = client.Watch(clients.WatchOpts{Cluster: "watch"})
	Expect(err).To(HaveOccurred())
	Expect(err.Error()).To(Equal(NoClusterResourceClientForClusterError("watch").Error()))
}
func ClusterResourceMultiClusterClientWatchAggregationTest(client ClusterResourceMultiClusterClient, aggregator wrapper.WatchAggregator) {
	w, errs, err := aggregator.Watch(clients.WatchOpts{})
	Expect(err).NotTo(HaveOccurred())
	go func() {
		defer GinkgoRecover()
		for {
			select {
			case err := <-errs:
				Expect(err).NotTo(HaveOccurred())
			case <-time.After(time.Second / 4):
				return
			}
		}
	}()

	cfg, err := kubeutils.GetConfig("", "")
	Expect(err).NotTo(HaveOccurred())
	client.ClusterAdded("", cfg)
	input := &ClusterResource{}
	input.SetMetadata(core.Metadata{
		Cluster: "write",
		Name:    "bar",
	})
	_, err = client.Write(input, clients.WriteOpts{})
	written, err := client.Write(input, clients.WriteOpts{})
	Expect(err).NotTo(HaveOccurred())
	Eventually(w, time.Second*5, time.Second/10).Should(Receive(And(ContainElement(written))))
}
