# k8s-custom-resources
A project containing all custom resource definitions used by HPE

Requirements for generating new CRDs.

#### 1. Add new types:

Open pkg/apis/hpestorage/v1/types.go and add your new type:

```
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type NewType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              NewTypeSpec `json:"spec"`
}
```

By default, it must include the TypeMeta and ObjectMeta.  The Spec is where you define properties specific to the new type:

```
type NewTypeSpec struct {
	Foo     string `json:"foo,omitempty"`
	Bar     string `json:"bar,omitempty"`
}
```

#### 2. Generate client, informers, and listers:

* Download https://github.com/kubernetes/code-generator
* Run ./code-generator/generate-groups.sh all "github.com/hpe-storage/k8s-custom-resources/pkg/client" "github.com/hpe-storage/k8s-custom-resources/pkg/apis" "hpestorage:v1"

#### 3. Add the CRD definition to our yaml deployment file:

```
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  creationTimestamp: null
  name: newtypes.storage.hpe.com
spec:
  group: storage.hpe.com
  names:
    kind: NewType
    plural: newtypes
  scope: Cluster
  validation:
    openAPIV3Schema:
      properties:
        hpeNodes:
          description: List of HPE new types.
          items:
            properties:
              foo:
                description: The foo description.
                type: string
              bar:
                description: The bar description.
                type: string
          type: array
  version: v1
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
```
