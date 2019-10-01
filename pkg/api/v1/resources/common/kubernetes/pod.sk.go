// Code generated by solo-kit. DO NOT EDIT.

package kubernetes

import (
	"sort"

	github_com_solo_io_solo_kit_api_external_kubernetes_pod "github.com/solo-io/solo-kit/api/external/kubernetes/pod"

	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewPod(namespace, name string) *Pod {
	pod := &Pod{}
	pod.Pod.SetMetadata(core.Metadata{
		Name:      name,
		Namespace: namespace,
	})
	return pod
}

// require custom resource to implement Clone() as well as resources.Resource interface

type CloneablePod interface {
	resources.Resource
	Clone() *github_com_solo_io_solo_kit_api_external_kubernetes_pod.Pod
}

var _ CloneablePod = &github_com_solo_io_solo_kit_api_external_kubernetes_pod.Pod{}

type Pod struct {
	github_com_solo_io_solo_kit_api_external_kubernetes_pod.Pod
}

func (r *Pod) Clone() resources.Resource {
	return &Pod{Pod: *r.Pod.Clone()}
}

func (r *Pod) Hash() uint64 {
	clone := r.Pod.Clone()

	resources.UpdateMetadata(clone, func(meta *core.Metadata) {
		meta.ResourceVersion = ""
	})

	return hashutils.HashAll(clone)
}

func (r *Pod) GroupVersionKind() schema.GroupVersionKind {
	return PodGVK
}

type PodList []*Pod

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list PodList) Find(namespace, name string) (*Pod, error) {
	for _, pod := range list {
		if pod.GetMetadata().Name == name {
			if namespace == "" || pod.GetMetadata().Namespace == namespace {
				return pod, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find pod %v.%v", namespace, name)
}

func (list PodList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, pod := range list {
		ress = append(ress, pod)
	}
	return ress
}

func (list PodList) Names() []string {
	var names []string
	for _, pod := range list {
		names = append(names, pod.GetMetadata().Name)
	}
	return names
}

func (list PodList) NamespacesDotNames() []string {
	var names []string
	for _, pod := range list {
		names = append(names, pod.GetMetadata().Namespace+"."+pod.GetMetadata().Name)
	}
	return names
}

func (list PodList) Sort() PodList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].GetMetadata().Less(list[j].GetMetadata())
	})
	return list
}

func (list PodList) Clone() PodList {
	var podList PodList
	for _, pod := range list {
		podList = append(podList, resources.Clone(pod).(*Pod))
	}
	return podList
}

func (list PodList) Each(f func(element *Pod)) {
	for _, pod := range list {
		f(pod)
	}
}

func (list PodList) EachResource(f func(element resources.Resource)) {
	for _, pod := range list {
		f(pod)
	}
}

func (list PodList) AsInterfaces() []interface{} {
	var asInterfaces []interface{}
	list.Each(func(element *Pod) {
		asInterfaces = append(asInterfaces, element)
	})
	return asInterfaces
}

var (
	PodGVK = schema.GroupVersionKind{
		Version: "kubernetes",
		Group:   "kubernetes.solo.io",
		Kind:    "Pod",
	}
)
