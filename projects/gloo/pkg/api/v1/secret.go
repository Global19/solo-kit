package v1

import (
	"sort"

	"github.com/gogo/protobuf/proto"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TODO: modify as needed to populate additional fields
func NewSecret(namespace, name string) *Secret {
	return &Secret{
		Metadata: core.Metadata{
			Name:      name,
			Namespace: namespace,
		},
	}
}

func (r *Secret) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

type SecretList []*Secret
type SecretsByNamespace map[string]SecretList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list SecretList) Find(namespace, name string) (*Secret, error) {
	for _, secret := range list {
		if secret.Metadata.Name == name {
			if namespace == "" || secret.Metadata.Namespace == namespace {
				return secret, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find secret %v.%v", namespace, name)
}

func (list SecretList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, secret := range list {
		ress = append(ress, secret)
	}
	return ress
}

func (list SecretList) Names() []string {
	var names []string
	for _, secret := range list {
		names = append(names, secret.Metadata.Name)
	}
	return names
}

func (list SecretList) NamespacesDotNames() []string {
	var names []string
	for _, secret := range list {
		names = append(names, secret.Metadata.Namespace+"."+secret.Metadata.Name)
	}
	return names
}

func (list SecretList) Sort() SecretList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Metadata.Less(list[j].Metadata)
	})
	return list
}

func (list SecretList) Clone() SecretList {
	var secretList SecretList
	for _, secret := range list {
		secretList = append(secretList, proto.Clone(secret).(*Secret))
	}
	return secretList
}

func (list SecretList) ByNamespace() SecretsByNamespace {
	byNamespace := make(SecretsByNamespace)
	for _, secret := range list {
		byNamespace.Add(secret)
	}
	return byNamespace
}

func (byNamespace SecretsByNamespace) Add(secret ...*Secret) {
	for _, item := range secret {
		byNamespace[item.Metadata.Namespace] = append(byNamespace[item.Metadata.Namespace], item)
	}
}

func (byNamespace SecretsByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace SecretsByNamespace) List() SecretList {
	var list SecretList
	for _, secretList := range byNamespace {
		list = append(list, secretList...)
	}
	return list.Sort()
}

func (byNamespace SecretsByNamespace) Clone() SecretsByNamespace {
	return byNamespace.List().Clone().ByNamespace()
}

var _ resources.Resource = &Secret{}

// Kubernetes Adapter for Secret

func (o *Secret) GetObjectKind() schema.ObjectKind {
	t := SecretCrd.TypeMeta()
	return &t
}

func (o *Secret) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*Secret)
}

var SecretCrd = crd.NewCrd("gloo.solo.io",
	"secrets",
	"gloo.solo.io",
	"v1",
	"Secret",
	"sec",
	&Secret{})
