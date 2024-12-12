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

// CNIExclusionSpec defines the desired state of CNIExclusion
type CNIExclusionSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// ExclusiveNamespace is a list of namespaces that are exclusive for CNI.
	// User can specify more than one namespaces that are exclusive for CNI in one yaml.
	// Effect: All "CNIExclusion" with all namespaces in "ExclusiveNamespaces" will be excluded from CNI, after the rule is created.
	// +optional
	ExclusiveNamespaces []string `json:"exclusive_namespaces,omitempty"`
}

// CNIExclusionStatus defines the observed state of CNIExclusion
type CNIExclusionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CNIExclusion is the Schema for the cniexclusions API
type CNIExclusion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CNIExclusionSpec   `json:"spec,omitempty"`
	Status CNIExclusionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CNIExclusionList contains a list of CNIExclusion
type CNIExclusionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CNIExclusion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CNIExclusion{}, &CNIExclusionList{})
}
