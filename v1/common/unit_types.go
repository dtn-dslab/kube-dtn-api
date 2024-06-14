package common

// Float percentage between 0 and 100
// +kubebuilder:validation:Pattern=`^(100(\.0+)?|\d{1,2}(\.\d+)?)$`
type Percentage string

// Duration string, e.g. "300ms", "1.5s".
// +kubebuilder:validation:Pattern=`^(\d+(\.\d+)?(ns|us|µs|μs|ms|s|m|h))+$`
type Duration string
