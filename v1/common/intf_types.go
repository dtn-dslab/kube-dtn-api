package common

import (
	"dslab.sjtu/kube-dtn/internal/api/v1/pb"
)

type NetworkIntf struct {
	// UID of the network interface, unique within all network interfaces, not within a pod
	UID int64 `json:"uid"`

	// Name of the network interface
	Name string `json:"name"`

	// MAC address of the network interface
	Mac Mac `json:"mac"`

	// IPv4 address of the network interface, Optional
	// +optional
	IPv4 IPv4 `json:"ipv4,omitempty"`

	// IPv6 address of the network interface, Optional
	// +optional
	IPv6 IPv6 `json:"ipv6,omitempty"`
}

func (n *NetworkIntf) ToProto() *pb.NetworkIntf {
	return &pb.NetworkIntf{
		Uid:  n.UID,
		Name: n.Name,
		Mac:  string(n.Mac),
		Ipv4: string(n.IPv4),
		Ipv6: string(n.IPv6),
	}
}
