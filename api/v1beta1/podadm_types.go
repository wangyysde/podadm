/*
Copyright 2024 Wayne Wang<net_use@bzhy.com>.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PodAdmSpec defines the desired state of PodAdm
type PodAdmSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// NameSpace is the namespace of the configMap or Secret
	NameSpace string `json:"nameSpace,omitempty"`

	// ObjectName name of the configMap or Secret
	ObjectName string `json:"objectName,omitempty"`

	// Kind of configMap or Secret
	Kind string `json:"kind,omitempty"`
}

// PodAdmStatus defines the observed state of PodAdm
type PodAdmStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// RefObjects list the objects which using the configMap or Secret
	RefObjects []RefObject `json:"refObjects"`
}

type RefObject struct {
	// TypeMeta data of the object which using the configMap or Secret
	metav1.TypeMeta `json:",inline"`

	// ObjectName name of the object which using the configMap or Secret
	ObjectName string `json:"objectName,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PodAdm is the Schema for the podadms API
type PodAdm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodAdmSpec   `json:"spec,omitempty"`
	Status PodAdmStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PodAdmList contains a list of PodAdm
type PodAdmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodAdm `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PodAdm{}, &PodAdmList{})
}
