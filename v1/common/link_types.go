package common

import (
	pb "dslab.sjtu/kube-dtn/internal/api/v1/pb"
)

type Link struct {
	// Unique identifier of a p2p link
	UID int `json:"uid"`

	// Source node of the link
	Src int `json:"src"`

	// Destination node of the link
	Dst int `json:"dst"`

	// Uni-directional link, default is false to make it bi-directional
	// +optional
	// +kubebuilder:default=false
	Unidirectional bool `json:"unidirectional,omitempty"`

	// Link properties, latency, bandwidth, etc
	// +optional
	Properties LinkProperties `json:"properties,omitempty"`
}

type LinkProperties struct {
	// Latency in duration string format, e.g. "300ms", "1.5s".
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// +optional
	Latency Duration `json:"latency,omitempty"`

	// Latency correlation in float percentage
	// +optional
	LatencyCorr Percentage `json:"latency_corr,omitempty"`

	// Jitter in duration string format, e.g. "300ms", "1.5s".
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// +optional
	Jitter Duration `json:"jitter,omitempty"`

	// Loss rate in float percentage
	// +optional
	Loss Percentage `json:"loss,omitempty"`

	// Loss correlation in float percentage
	// +optional
	LossCorr Percentage `json:"loss_corr,omitempty"`

	// Bandwidth rate limit, e.g. 1000(bit/s), 100kbit, 100Mbps, 1Gibps.
	// For more information, refer to https://man7.org/linux/man-pages/man8/tc.8.html.
	// +optional
	// +kubebuilder:validation:Pattern=`^\d+(\.\d+)?([KkMmGg]i?)?(bit|bps)?$`
	Rate string `json:"rate,omitempty"`

	// Gap every N packets
	// +optional
	// +kubebuilder:validation:Minimum=0
	Gap uint32 `json:"gap,omitempty"`

	// Duplicate rate in float percentage
	// +optional
	Duplicate Percentage `json:"duplicate,omitempty"`

	// Duplicate correlation in float percentage
	// +optional
	DuplicateCorr Percentage `json:"duplicate_corr,omitempty"`

	// Reorder probability in float percentage
	// +optional
	ReorderProb Percentage `json:"reorder_prob,omitempty"`

	// Reorder correlation in float percentage
	// +optional
	ReorderCorr Percentage `json:"reorder_corr,omitempty"`

	// Corrupt probability in float percentage
	// +optional
	CorruptProb Percentage `json:"corrupt_prob,omitempty"`

	// Corrupt correlation in float percentage
	// +optional
	CorruptCorr Percentage `json:"corrupt_corr,omitempty"`
}

func (l *Link) ToProto() *pb.Link {
	return &pb.Link{
		Uid:            int64(l.UID),
		Src:            int64(l.Src),
		Dst:            int64(l.Dst),
		Properties:     l.Properties.ToProto(),
		UniDirectional: l.Unidirectional,
	}
}

func (p *LinkProperties) ToProto() *pb.LinkProperties {
	return &pb.LinkProperties{
		Latency:       string(p.Latency),
		LatencyCorr:   string(p.LatencyCorr),
		Jitter:        string(p.Jitter),
		Loss:          string(p.Loss),
		LossCorr:      string(p.LossCorr),
		Rate:          p.Rate,
		Gap:           p.Gap,
		Duplicate:     string(p.Duplicate),
		DuplicateCorr: string(p.DuplicateCorr),
		ReorderProb:   string(p.ReorderProb),
		ReorderCorr:   string(p.ReorderCorr),
		CorruptProb:   string(p.CorruptProb),
		CorruptCorr:   string(p.CorruptCorr),
	}
}
