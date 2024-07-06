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
	common "dslab.sjtu/kube-dtn/api/v1/common"
	"dslab.sjtu/kube-dtn/internal/api/v1/pb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NetworkNodeSpec defines the desired state of NetworkNode
type NetworkNodeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Name equates to the name of the network Pod
	Name string `json:"name"`

	// Namespace equates to the namespace of the network Pod
	Namespace string `json:"namespace"`

	// NetworkIntfs is a list of network interfaces
	// +optional
	NetworkIntfs []common.NetworkIntf `json:"network_intfs,omitempty"`
}

func (n *NetworkNodeSpec) ToProto() *pb.NetworkNode {
	return &pb.NetworkNode{
		Name:      n.Name,
		Namespace: n.Namespace,
		NetworkIntfs: func() []*pb.NetworkIntf {
			var networkIntfs []*pb.NetworkIntf
			for _, networkIntf := range n.NetworkIntfs {
				networkIntfs = append(networkIntfs, networkIntf.ToProto())
			}
			return networkIntfs
		}(),
	}
}

// NetworkNodeStatus defines the observed state of NetworkNode
type NetworkNodeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Ready equates to the readiness of the network node
	Ready bool `json:"ready"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// NetworkNode is the Schema for the networknodes API
type NetworkNode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkNodeSpec   `json:"spec,omitempty"`
	Status NetworkNodeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkNodeList contains a list of NetworkNode
type NetworkNodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkNode `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkNode{}, &NetworkNodeList{})
}