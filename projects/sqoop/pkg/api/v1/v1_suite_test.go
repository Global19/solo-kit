package v1

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestResolverMap(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ResolverMap Suite")
}