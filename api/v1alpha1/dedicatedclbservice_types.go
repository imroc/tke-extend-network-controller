/*
Copyright 2024.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DedicatedCLBServiceSpec defines the desired state of DedicatedCLBService
type DedicatedCLBServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +default=500
	MinPort int32 `json:"minPort"`
	// +default=550
	MaxPort     int32 `json:"maxPort"`
	ServiceName int32 `json:"serviceName"`
	// +optional
	ExtensiveParameters string `json:"extensiveParameters"`
	// +optional
	ExistedLbIds []string `json:"existedLbIds"`
}

// DedicatedCLBServiceStatus defines the observed state of DedicatedCLBService
type DedicatedCLBServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DedicatedCLBService is the Schema for the dedicatedclbservices API
type DedicatedCLBService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DedicatedCLBServiceSpec   `json:"spec,omitempty"`
	Status DedicatedCLBServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedCLBServiceList contains a list of DedicatedCLBService
type DedicatedCLBServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DedicatedCLBService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DedicatedCLBService{}, &DedicatedCLBServiceList{})
}
