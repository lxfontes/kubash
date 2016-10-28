package install

import (
	"github.com/lxfontes/kubash/types"
	"k8s.io/client-go/1.5/pkg/api/unversioned"
	"k8s.io/client-go/1.5/pkg/apimachinery/announced"
	"k8s.io/client-go/1.5/pkg/runtime"
)

func init() {
	if err := announced.NewGroupMetaFactory(
		&announced.GroupMetaFactoryArgs{
			GroupName:              types.GroupName,
			VersionPreferenceOrder: []string{types.GroupVersion.Version},
			ImportPrefix:           "github.com/lxfontes/kubash/types",
		},
		announced.VersionToSchemeFunc{
			types.GroupVersion.Version: AddToScheme,
		},
	).Announce().RegisterAndEnable(); err != nil {
		panic(err)
	}
}

// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) unversioned.GroupKind {
	return types.GroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) unversioned.GroupResource {
	return types.GroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	// TODO this gets cleaned up when the types are fixed
	scheme.AddKnownTypes(types.GroupVersion,
		&types.RunItem{},
		&types.RunItemList{},
	)
	return nil
}
