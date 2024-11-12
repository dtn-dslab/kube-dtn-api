package common

import (
	"dslab.sjtu/kube-dtn/api/v1/pb"
)

type NetworkIntf struct {
	// UID of the network interface, unique within all network interfaces, not within a pod
	UID int32 `json:"uid"`

	// Name of the network interface
	Name string `json:"name"`

	// MAC address of the network interface
	Mac Mac `json:"mac"`

	// IPv4 address of the network interface, Optional
	// +optional
	IPv4 MaskedIPv4 `json:"ipv4,omitempty"`

	// IPv6 address of the network interface, Optional
	// +optional
	IPv6 MaskedIPv6 `json:"ipv6,omitempty"`

	// Owner pod's k8s namespace, uid only unique within the namespace
	Ns string `json:"Ns,omitempty"`
}

func (n *NetworkIntf) ToProto() *pb.NetworkIntf {
	return &pb.NetworkIntf{
		Uid:  n.UID,
		Name: n.Name,
		Mac:  string(n.Mac),
		Ipv4: string(n.IPv4),
		Ipv6: string(n.IPv6),
		Ns:   string(n.Ns),
	}
}

func (n *NetworkIntf) Equal(other *NetworkIntf) bool {
	if n.UID != other.UID {
		return false
	}
	if n.Name != other.Name {
		return false
	}
	if n.Mac != other.Mac {
		return false
	}
	if n.IPv4 != other.IPv4 {
		return false
	}
	if n.IPv6 != other.IPv6 {
		return false
	}
	if n.Ns != other.Ns {
		return false
	}
	return true
}

func GetIntfIDKey(ns string, uid int32) string {
	return ns + "-" + string(uid)
}
