package generic

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	. "github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/test/mocks"
)

// Call within "It"
func TestCrudClient(namespace string, client ResourceClient) {
	client.Register()
	foo := "foo"
	input := mocks.NewMockData(namespace, foo)
	data := map[string]string{"hello": "goodbye"}
	input.Data = data
	labels := map[string]string{"pick": "me"}
	input.Metadata.Labels = labels

	err := client.Register()
	Expect(err).NotTo(HaveOccurred())

	r1, err := client.Write(input, clients.WriteOpts{})
	Expect(err).NotTo(HaveOccurred())

	_, err = client.Write(input, clients.WriteOpts{})
	Expect(err).To(HaveOccurred())
	Expect(errors.IsExist(err)).To(BeTrue())

	Expect(r1).To(BeAssignableToTypeOf(&mocks.MockData{}))
	Expect(r1.GetMetadata().Name).To(Equal(foo))
	if namespace == "" {
		namespace = DefaultNamespace
	}
	Expect(r1.GetMetadata().Namespace).To(Equal(namespace))
	Expect(r1.GetMetadata().ResourceVersion).NotTo(Equal(""))
	Expect(r1.(*mocks.MockData).Data).To(Equal(data))

	// if exists and resource ver was not updated, error
	_, err = client.Write(input, clients.WriteOpts{
		OverwriteExisting: true,
	})
	Expect(err).To(HaveOccurred())

	resources.UpdateMetadata(input, func(meta *core.Metadata) {
		meta.ResourceVersion = r1.GetMetadata().ResourceVersion
	})
	data = map[string]string{"asdf": "qwer"}
	input.Data = data

	oldRv := r1.GetMetadata().ResourceVersion

	r1, err = client.Write(input, clients.WriteOpts{
		OverwriteExisting: true,
	})
	Expect(err).NotTo(HaveOccurred())

	read, err := client.Read(namespace, foo, clients.ReadOpts{})
	Expect(err).NotTo(HaveOccurred())
	// it should update the resource version on the new write
	Expect(read.GetMetadata().ResourceVersion).NotTo(Equal(oldRv))
	Expect(read).To(Equal(r1))

	_, err = client.Read("doesntexist", foo, clients.ReadOpts{})
	Expect(err).To(HaveOccurred())
	Expect(errors.IsNotExist(err)).To(BeTrue())

	boo := "boo"
	input = &mocks.MockData{
		Data: data,
		Metadata: core.Metadata{
			Name:      boo,
			Namespace: namespace,
		},
	}
	r2, err := client.Write(input, clients.WriteOpts{})
	Expect(err).NotTo(HaveOccurred())

	// with labels
	list, err := client.List(namespace, clients.ListOpts{
		Selector: labels,
	})
	Expect(err).NotTo(HaveOccurred())
	Expect(list).To(ContainElement(r1))
	Expect(list).NotTo(ContainElement(r2))

	// without
	list, err = client.List(namespace, clients.ListOpts{})
	Expect(err).NotTo(HaveOccurred())
	Expect(list).To(ContainElement(r1))
	Expect(list).To(ContainElement(r2))

	err = client.Delete(namespace, "adsfw", clients.DeleteOpts{})
	Expect(err).To(HaveOccurred())
	Expect(errors.IsNotExist(err)).To(BeTrue())

	err = client.Delete(namespace, "adsfw", clients.DeleteOpts{
		IgnoreNotExist: true,
	})
	Expect(err).NotTo(HaveOccurred())

	err = client.Delete(namespace, r2.GetMetadata().Name, clients.DeleteOpts{})
	Expect(err).NotTo(HaveOccurred())
	list, err = client.List(namespace, clients.ListOpts{})
	Expect(err).NotTo(HaveOccurred())
	Expect(list).To(ContainElement(r1))
	Expect(list).NotTo(ContainElement(r2))

	w, errs, err := client.Watch(namespace, clients.WatchOpts{RefreshRate: time.Millisecond})
	Expect(err).NotTo(HaveOccurred())

	var r3 resources.Resource
	wait := make(chan struct{})
	go func() {
		defer GinkgoRecover()
		defer close(wait)
		resources.UpdateMetadata(r2, func(meta *core.Metadata) {
			meta.ResourceVersion = ""
		})
		r2, err = client.Write(r2, clients.WriteOpts{})
		Expect(err).NotTo(HaveOccurred())

		input = &mocks.MockData{
			Data: data,
			Metadata: core.Metadata{
				Name:      "goo",
				Namespace: namespace,
			},
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
		case <-time.After(time.Second / 4):
			break drain
		}
	}

	Expect(list).To(ContainElement(r1))
	Expect(list).To(ContainElement(r2))
	Expect(list).To(ContainElement(r3))
}
