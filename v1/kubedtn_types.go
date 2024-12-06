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

type KubeDTNK8sNodeInfo struct {
	NodeName string `json:"node_name,omitempty"`
	NodeRole string `json:"node_role,omitempty"` // can be "kubedtn-host" or "kubedtn-dpu". Set by the Reconciler
	NodeIP   string `json:"node_ip,omitempty"`
}

func (n *KubeDTNK8sNodeInfo) DeepEqual(other *KubeDTNK8sNodeInfo) bool {
	return n.NodeName == other.NodeName && n.NodeRole == other.NodeRole && n.NodeIP == other.NodeIP
}

type KubeDTNK8sDPUYamlSpec struct {
	DPUNodeName string `json:"dpu_node_name,omitempty"`
}

type KubeDTNK8sHostNodeYamlSpec struct {
	Name string `json:"name,omitempty"`
	Mode string `json:"mype,omitempty"` // Can be "host" or "dpu"

	// Currently only have dpu_node_name field
	KubeDTNK8sDPUYamlSpec `json:",inline,omitempty"`
}

// KubeDTNSpec defines the desired state of KubeDTN
type KubeDTNSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	HostNodes []KubeDTNK8sHostNodeYamlSpec `json:"host_nodes,omitempty"`
}

// KubeDTNStatus defines the observed state of KubeDTN
type KubeDTNStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KubeDTN is the Schema for the kubedtns API
type KubeDTN struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubeDTNSpec   `json:"spec,omitempty"`
	Status KubeDTNStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KubeDTNList contains a list of KubeDTN
type KubeDTNList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubeDTN `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubeDTN{}, &KubeDTNList{})
}
