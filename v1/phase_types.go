package v1

const (
	// Some transient status. Expect the Reconciler to do the next step, or Reconciler is doing now.

	// Expect Reconciler to truly apply the current view.
	PhasePending = "Pending"
	// Expect Reconciler to just do grpc to remove the data plane entity, but not the Persistence Storage.
	PhaseCleaning = "Cleaning"
	// If get user's "kubectl delete" command, the phase will be Terminating, till all related resources are deleted(Ignore logic error, just retry to delete).
	// Equal to K8s's DeletionTimestamp != nil, but also necessary for HealthCheckServer.
	PhaseTerminating = "Terminating"

	// Some Steady Status, can be checked by HealthCheckServer

	PhaseRunning     = "Running"
	PhaseUnscheduled = "Unscheduled"

	// If business logic error, the phase will be Error. It can be retried periodically. Maybe replaced by ErrorSuffix
	PhaseError = "Error"
	// Help know the error is from which part. Equal to PhaseError. Maybe only transient phase can be appended with ErrorSuffix.
	ErrorSuffix = "Error"

	// If HealthCheckServer check the K8s Node/Daemon is unhealthy, the phase will be Unhealthy.
	PhaseUnhealthy  = "Unhealthy"
	UnhealthySuffix = "Unhealthy"

	FinalizerName = "kubedtn.dslab.sjtu/finalizer"

	// When kClient do write, it can set this annotation to skip once Reconcile process, Reconciler need to remove it (Currently only for networkNode)
	AnnotationReconcileSkip = "kubedtn.dslab.sjtu/reconcile-skip"
)

const (
	// Used by daemon to wait for expected phase
	PhasePollingSleepMs = 50
	PhaseWaitTimeoutMs  = 3000
)

func IsInTransientPhase(s string) bool {
	return s == PhasePending || s == PhaseCleaning || s == PhaseTerminating
}

func IsInErrorPhase(s string) bool {
	return len(s) >= len(ErrorSuffix) && s[len(s)-len(ErrorSuffix):] == ErrorSuffix
}

// UnhealthySuffix can be appended to one with ErrorSuffix, but the inverse is not.
func IsInUnhealthyPhase(s string) bool {
	return len(s) >= len(UnhealthySuffix) && s[len(s)-len(UnhealthySuffix):] == UnhealthySuffix
}

// Since ErrorSuffix == "Error" == ErrorPhase, len(s) == len(ErrorSuffix) is ok.
func AppendErrorSuffixIfNotExist(s string) string {
	if IsInErrorPhase(s) {
		return s
	}
	return s + ErrorSuffix
}

func RemoveErrorSuffix(s string) string {
	if IsInErrorPhase(s) {
		return s[:len(s)-len(ErrorSuffix)]
	}
	return s
}

func AppendUnhealthySuffixIfNotExist(s string) string {
	if IsInUnhealthyPhase(s) {
		return s
	}
	return s + UnhealthySuffix
}

func RemoveUnhealthySuffix(s string) string {
	if IsInUnhealthyPhase(s) {
		return s[:len(s)-len(UnhealthySuffix)]
	}
	return s
}
