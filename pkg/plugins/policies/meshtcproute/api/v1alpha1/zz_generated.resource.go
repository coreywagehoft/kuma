// Generated by tools/resource-gen.
// Run "make generate" to update this file.

// nolint:whitespace
package v1alpha1

import (
	_ "embed"
	"fmt"

	"k8s.io/kube-openapi/pkg/validation/spec"
	"sigs.k8s.io/yaml"

	"github.com/kumahq/kuma/pkg/core/resources/model"
)

//go:embed schema.yaml
var rawSchema []byte

func init() {
	var schema spec.Schema
	if err := yaml.Unmarshal(rawSchema, &schema); err != nil {
		panic(err)
	}
	rawSchema = nil
	MeshTCPRouteResourceTypeDescriptor.Schema = &schema
}

const (
	MeshTCPRouteType model.ResourceType = "MeshTCPRoute"
)

var _ model.Resource = &MeshTCPRouteResource{}

type MeshTCPRouteResource struct {
	Meta model.ResourceMeta
	Spec *MeshTCPRoute
}

func NewMeshTCPRouteResource() *MeshTCPRouteResource {
	return &MeshTCPRouteResource{
		Spec: &MeshTCPRoute{},
	}
}

func (t *MeshTCPRouteResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *MeshTCPRouteResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *MeshTCPRouteResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *MeshTCPRouteResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*MeshTCPRoute)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &MeshTCPRoute{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *MeshTCPRouteResource) Descriptor() model.ResourceTypeDescriptor {
	return MeshTCPRouteResourceTypeDescriptor
}

func (t *MeshTCPRouteResource) Validate() error {
	if v, ok := interface{}(t).(interface{ validate() error }); !ok {
		return nil
	} else {
		return v.validate()
	}
}

var _ model.ResourceList = &MeshTCPRouteResourceList{}

type MeshTCPRouteResourceList struct {
	Items      []*MeshTCPRouteResource
	Pagination model.Pagination
}

func (l *MeshTCPRouteResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *MeshTCPRouteResourceList) GetItemType() model.ResourceType {
	return MeshTCPRouteType
}

func (l *MeshTCPRouteResourceList) NewItem() model.Resource {
	return NewMeshTCPRouteResource()
}

func (l *MeshTCPRouteResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*MeshTCPRouteResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*MeshTCPRouteResource)(nil), r)
	}
}

func (l *MeshTCPRouteResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

var MeshTCPRouteResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:                MeshTCPRouteType,
	Resource:            NewMeshTCPRouteResource(),
	ResourceList:        &MeshTCPRouteResourceList{},
	Scope:               model.ScopeMesh,
	KDSFlags:            model.FromGlobalToZone,
	WsPath:              "meshtcproutes",
	KumactlArg:          "meshtcproute",
	KumactlListArg:      "meshtcproutes",
	AllowToInspect:      true,
	IsPolicy:            true,
	IsExperimental:      false,
	SingularDisplayName: "Mesh TCP Route",
	PluralDisplayName:   "Mesh TCP Routes",
	IsPluginOriginated:  true,
}
