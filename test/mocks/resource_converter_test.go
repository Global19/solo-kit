package mocks_test

// TODO joekelley pkg name

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	// TODO joekelley pkg name
	"github.com/solo-io/solo-kit/test/mocks"
	"github.com/solo-io/solo-kit/test/mocks/v1"
	v1 "github.com/solo-io/solo-kit/test/mocks/v1"
	"github.com/solo-io/solo-kit/test/mocks/v1alpha1"
	"github.com/solo-io/solo-kit/test/mocks/v2alpha1"
)

var converter crd.Converter

var _ = Describe("FakeResourceConverter", func() {
	BeforeEach(func() {
		converter = mocks.NewFakeResourceConverter(fakeResourceUpConverter{}, fakeResourceDownConverter{})
	})

	Describe("Convert", func() {
		It("works for noop conversions", func() {
			src := &v1alpha1.FakeResource{Metadata: core.Metadata{Name: "test"}}
			dst := &v1alpha1.FakeResource{}
			err := converter.Convert(src, dst)
			Expect(err).NotTo(HaveOccurred())
			Expect(dst.GetMetadata().Name).To(Equal("test"))
		})

		It("converts all the way up", func() {
			src := &v1alpha1.FakeResource{}
			dst := &v1.FakeResource{}
			err := converter.Convert(src, dst)
			Expect(err).NotTo(HaveOccurred())
			Expect(dst.GetMetadata().Name).To(Equal("v2alpha1"))
		})

		It("converts all the way down", func() {
			src := &v1.FakeResource{}
			dst := &v1alpha1.FakeResource{}
			err := converter.Convert(src, dst)
			Expect(err).NotTo(HaveOccurred())
			Expect(dst.GetMetadata().Name).To(Equal("v1alpha1"))
		})
	})
})

type fakeResourceUpConverter struct{}

func (fakeResourceUpConverter) FromV1Alpha1ToV1(src *v1alpha1.FakeResource) *v1.FakeResource {
	return &v1.FakeResource{Metadata: core.Metadata{Name: "v1"}}
}

type fakeResourceDownConverter struct{}

func (fakeResourceDownConverter) FromV1ToV1Alpha1(src *v1.FakeResource) *v1alpha1.FakeResource {
	return &v1alpha1.FakeResource{Metadata: core.Metadata{Name: "v1alpha1"}}
}

var _ = Describe("MockResourceConverter", func() {
	BeforeEach(func() {
		converter = mocks.NewMockResourceConverter(mockResourceUpConverter{}, mockResourceDownConverter{})
	})

	Describe("Convert", func() {
		It("works for noop conversions", func() {
			src := &v1alpha1.MockResource{Metadata: core.Metadata{Name: "test"}}
			dst := &v1alpha1.MockResource{}
			err := converter.Convert(src, dst)
			Expect(err).NotTo(HaveOccurred())
			Expect(dst.GetMetadata().Name).To(Equal("test"))
		})

		It("converts all the way up", func() {
			src := &v1alpha1.MockResource{}
			dst := &v2alpha1.MockResource{}
			err := converter.Convert(src, dst)
			Expect(err).NotTo(HaveOccurred())
			Expect(dst.GetMetadata().Name).To(Equal("v2alpha1"))
		})

		It("converts all the way down", func() {
			src := &v2alpha1.MockResource{}
			dst := &v1alpha1.MockResource{}
			err := converter.Convert(src, dst)
			Expect(err).NotTo(HaveOccurred())
			Expect(dst.GetMetadata().Name).To(Equal("v1alpha1"))
		})
	})
})

type mockResourceUpConverter struct{}

func (mockResourceUpConverter) FromV1Alpha1ToV1(src *v1alpha1.MockResource) *v1.MockResource {
	return &v1.MockResource{Metadata: core.Metadata{Name: "v1"}}
}
func (mockResourceUpConverter) FromV1ToV2Alpha1(src *v1.MockResource) *v2alpha1.MockResource {
	return &v2alpha1.MockResource{Metadata: core.Metadata{Name: "v2alpha1"}}
}

type mockResourceDownConverter struct{}

func (mockResourceDownConverter) FromV1ToV1Alpha1(src *v1.MockResource) *v1alpha1.MockResource {
	return &v1alpha1.MockResource{Metadata: core.Metadata{Name: "v1alpha1"}}
}
func (mockResourceDownConverter) FromV2Alpha1ToV1(src *v2alpha1.MockResource) *v1.MockResource {
	return &v1.MockResource{Metadata: core.Metadata{Name: "v1"}}
}
