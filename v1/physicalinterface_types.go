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
	pb "dslab.sjtu/kube-dtn/api/v1/pb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	VxlanBackend     = "vxlan"
	RawDeviceBackend = "raw_device"
)

// PhysicalInterfaceSpec defines the desired state of PhysicalInterface
type PhysicalInterfaceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// UID of the physical interface, unique within all interfaces, including network interfaces
	UID int32 `json:"uid"`

	// Backend is the backend of the physical interface, e.g. raw_device, vxlan
	Backend string `json:"backend"`

	// Name of the physical interface
	Name string `json:"name"`

	// Mac address of remote physical interface
	Mac common.Mac `json:"mac"`

	// NodeName is the name of the node where the virtual interface is configured
	NodeName string `json:"node_name"`

	// IPv4 address of the physical interface, Optional
	RawDevice *RawDeviceSpec `json:"raw_device,omitempty"`

	// Vxlan of the physical interface, Optional
	Vxlan *VxlanSpec `json:"vxlan,omitempty"`
}

func (p *PhysicalInterfaceSpec) ToProto() *pb.PhysicalIntf {
	intf := &pb.PhysicalIntf{
		Uid:      p.UID,
		Name:     p.Name,
		Mac:      string(p.Mac),
		NodeName: p.NodeName,
		Backend:  p.Backend,
	}

	switch p.Backend {
	case "RawDevice":
		intf.DeviceName = p.RawDevice.DeviceName
	case "Vxlan":
		intf.Vni = p.Vxlan.VNI
		intf.VtepIp = string(p.Vxlan.VtepIP)
		intf.DstPort = p.Vxlan.DstPort
	}

	return intf
}

type RawDeviceSpec struct {
	// Name of the raw device
	DeviceName string `json:"device_name"`
}

type VxlanSpec struct {
	// VNI is the Vxlan Network Identifier
	VNI int32 `json:"vni"`

	// VtepIP is the Vxlan Tunnel End Point IP address
	VtepIP common.IPv4 `json:"vtep_ip"`

	// DstPort is the destination port of the Vxlan, Optional
	DstPort int32 `json:"dst_port,omitempty"`
}

// PhysicalInterfaceStatus defines the observed state of PhysicalInterface
type PhysicalInterfaceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Phase equates to the phase of the physical interface, e.g. Pending, Running, Failed
	Phase string `json:"phase,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// PhysicalInterface is the Schema for the physicalinterfaces API
type PhysicalInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PhysicalInterfaceSpec   `json:"spec,omitempty"`
	Status PhysicalInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PhysicalInterfaceList contains a list of PhysicalInterface
type PhysicalInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PhysicalInterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PhysicalInterface{}, &PhysicalInterfaceList{})
}
