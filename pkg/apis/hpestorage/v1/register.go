// Copyright 2019 Hewlett Packard Enterprise Development LP

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GroupName is the group name used in this package
const GroupName = "storage.hpe.com"

var (
	// SchemeBuilder is the new scheme builder
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	// AddToScheme adds to the scheme
	AddToScheme = SchemeBuilder.AddToScheme
	// SchemeGroupVersion is the roup version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1"}
)

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func init() {
	SchemeBuilder.Register(addKnownTypes)
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&HPENodeInfo{},
		&HPENodeInfoList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
