package common

// MAC address
// +kubebuilder:validation:Pattern=`^(([0-9A-Fa-f]{2}[:-]){5}[0-9A-Fa-f]{2})?$`
type Mac string

// IPv4 address
type IPv4 string

// IPv6 address
type IPv6 string
