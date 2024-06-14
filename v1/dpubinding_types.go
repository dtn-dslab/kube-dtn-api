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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DPUBindingSpec defines the desired state of DPUBinding
type DPUBindingSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// DPU is the node name of the DPU node
	// +kubebuilder:validation:Required
	DPU string `json:"dpu"`

	// Networks is the list of network ids to bind to the DPU for offloading
	// +kubebuilder:validation:Required
	Networks []int `json:"networks"`
}

// DPUBindingStatus defines the observed state of DPUBinding
type DPUBindingStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Active indicates whether the DPUBinding is active
	Active bool `json:"active"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DPUBinding is the Schema for the dpubindings API
type DPUBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DPUBindingSpec   `json:"spec,omitempty"`
	Status DPUBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DPUBindingList contains a list of DPUBinding
type DPUBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DPUBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DPUBinding{}, &DPUBindingList{})
}
