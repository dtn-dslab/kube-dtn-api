package v1

const (
	KubeDTNPrefix             = "kubedtn.dslab.sjtu"
	KubeDTNSystemK8sNamespace = "kube-dtn-system"

	KubeDTNDaemonDefaultPodCNIFuncMark = "kubedtn.dslab.sjtu/default-cni-func"

	// Please keep with the name in yaml config/daemon/daemon-host.yaml !
	KubeDTNHostDaemonSetName = "kubedtn-host"
	KubeDTNDPUDaemonSetName  = "kubedtn-dpu"

	K8sNodeRoleKubeDTNKey       = "role"
	K8sNodeRoleKubeDTNHostValue = "kubedtn-host"
	K8sNodeRoleKubeDTNDPUValue  = "kubedtn-dpu"

	KubeDTNK8sHostNodeModeHost = "host"
	KubeDTNK8sHostNodeModeDPU  = "dpu"

	// For kubedtn yaml apply, the only effective k8sName is "kubedtn", the only effective k8sNamespace is "kube-dtn-system"
	KubeDTNClusterInfoAssignedK8sName = "kubedtn-cluster-info"

	// For Host/DPU DaemonSet created Pod, they have a label "kubedtn.dslab.sjtu/instance" with value "daemon-host" or "daemon-dpu"
	KubeDTNPodInstanceKey             = KubeDTNPrefix + "/instance"
	KubeDTNPodInstanceDaemonHostValue = "daemon-host"
	KubeDTNPodInstanceDaemonDPUValue  = "daemon-dpu"
)

const (
	InternalLinkDirectionSender   = "sender"
	InternalLinkDirectionReceiver = "receiver"
)
