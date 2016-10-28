package types

import (
	"k8s.io/client-go/1.5/pkg/api"
	"k8s.io/client-go/1.5/pkg/api/unversioned"
)

type RunItemList struct {
	unversioned.TypeMeta `json:",inline"`
	unversioned.ListMeta `json:"metadata,omitempty"`
	Items                []RunItem `json:"items"`
}

type RunItem struct {
	unversioned.TypeMeta `json:",inline"`
	api.ObjectMeta       `json:"metadata,omitempty"`
	Script               string   `json:"script"`
	Arguments            []string `json:"arguments,omitempty"`
}

// GroupName is the group name use in this package
const GroupName = "kubash.com"

// GroupVersion is group version used to register these objects
var GroupVersion = unversioned.GroupVersion{Group: GroupName, Version: "v1"}

func (obj *RunItem) GetObjectKind() unversioned.ObjectKind { return &obj.TypeMeta }
