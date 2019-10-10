package multicluster_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/multicluster"
)

var _ = Describe("Shared Cache Manager", func() {
	var manager multicluster.KubeSharedCacheManager

	BeforeEach(func() {
		manager = multicluster.NewKubeSharedCacheManager(context.Background())
	})

	It("works", func() {
		cluster1, cluster2 := "one", "two"

		cache1 := manager.GetCache(cluster1)
		Expect(cache1).NotTo(BeNil())
		cache2 := manager.GetCache(cluster2)
		Expect(cache2).NotTo(BeNil())
		Expect(cache1 == cache2).To(BeFalse(), "each has a unique cache")
	})
})
