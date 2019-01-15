// Code generated by solo-kit. DO NOT EDIT.

// +build solokit

package v1

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-kit/test/tests/typed"
)

var _ = FDescribe("ClusterResourceClient", func() {
	for _, test := range []typed.ResourceClientTester{
		&typed.KubeRcTester{Crd: ClusterResourceCrd},
	} {
		Context("resource client backed by "+test.Description(), func() {
			var (
				client              ClusterResourceClient
				err                 error
				name1, name2, name3 = "foo" + helpers.RandString(3), "boo" + helpers.RandString(3), "goo" + helpers.RandString(3)
			)

			BeforeEach(func() {
				factory := test.Setup("")
				client, err = NewClusterResourceClient(factory)
				Expect(err).NotTo(HaveOccurred())
			})
			AfterEach(func() {
				client.Delete(name1, clients.DeleteOpts{})
				client.Delete(name2, clients.DeleteOpts{})
				client.Delete(name3, clients.DeleteOpts{})
			})
			It("CRUDs ClusterResources", func() {
				ClusterResourceClientTest(client, name1, name2, name3)
			})
		})
	}
})

func ClusterResourceClientTest(client ClusterResourceClient, name1, name2, name3 string) {
	err := client.Register()
	Expect(err).NotTo(HaveOccurred())

	name := name1
	input := NewClusterResource("", name)
	r1, err := client.Write(input, clients.WriteOpts{})
	Expect(err).NotTo(HaveOccurred())

	_, err = client.Write(input, clients.WriteOpts{})
	Expect(err).To(HaveOccurred())
	Expect(errors.IsExist(err)).To(BeTrue())

	Expect(r1).To(BeAssignableToTypeOf(&ClusterResource{}))
	Expect(r1.GetMetadata().Name).To(Equal(name))
	Expect(r1.Metadata.ResourceVersion).NotTo(Equal(input.Metadata.ResourceVersion))
	Expect(r1.Metadata.Ref()).To(Equal(input.Metadata.Ref()))
	Expect(r1.Status).To(Equal(input.Status))
	Expect(r1.BasicField).To(Equal(input.BasicField))

	_, err = client.Write(input, clients.WriteOpts{
		OverwriteExisting: true,
	})
	Expect(err).To(HaveOccurred())

	input.Metadata.ResourceVersion = r1.GetMetadata().ResourceVersion
	r1, err = client.Write(input, clients.WriteOpts{
		OverwriteExisting: true,
	})
	Expect(err).NotTo(HaveOccurred())
	read, err := client.Read(name, clients.ReadOpts{})
	Expect(err).NotTo(HaveOccurred())
	Expect(read).To(Equal(r1))
	_, err = client.Read(name, clients.ReadOpts{})
	Expect(err).To(HaveOccurred())
	Expect(errors.IsNotExist(err)).To(BeTrue())

	name = name2
	input = &ClusterResource{}

	input.Metadata = core.Metadata{
		Name: name,
	}

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
	list, err = client.List(clients.ListOpts{})
	Expect(err).NotTo(HaveOccurred())
	Expect(list).To(ContainElement(r1))
	Expect(list).NotTo(ContainElement(r2))
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
		input.Metadata = core.Metadata{
			Name: name,
		}

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

drain:
	for {
		select {
		case list = <-w:
		case err := <-errs:
			Expect(err).NotTo(HaveOccurred())
		case <-time.After(time.Millisecond * 500):
			break drain
		}
	}

	Expect(list).To(ContainElement(r1))
	Expect(list).To(ContainElement(r2))
	Expect(list).To(ContainElement(r3))
}
