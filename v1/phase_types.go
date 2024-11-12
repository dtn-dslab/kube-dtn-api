package v1

import (
	"crypto/rand"
	"fmt"
	"time"
)

const (
	// Some transient status. Expect the Reconciler to do the next step, or Reconciler is doing now.

	// Expect Reconciler to truly apply the current view.
	PhasePending = "Pending"
	// Expect Reconciler to just do grpc to remove the data plane entity, but not the Persistence Storage.
	PhaseCleaning = "Cleaning"
	// If get user's "kubectl delete" command, the phase will be Terminating, till all related resources are deleted(Ignore logic error, just retry to delete).
	// Equal to K8s's DeletionTimestamp != nil, but also necessary for HealthCheckServer.
	PhaseTerminating = "Terminating"

	// Transient phase only for newly add link. Update Link is Pending.
	PhaseInitializing = "Initializing"
	// Node intf change (like IntfName, MAC Addr, etc.) and trigger the link to rebuild.
	PhaseRebuilding = "Rebuilding"

	// Some Steady Status, can be checked by HealthCheckServer

	PhaseRunning     = "Running"
	PhaseUnscheduled = "Unscheduled"

	// If business logic error, the phase will be Error. It can be retried periodically. Maybe replaced by ErrorSuffix
	PhaseError = "Error"
	// Help know the error is from which part. Equal to PhaseError. Maybe only transient phase can be appended with ErrorSuffix.
	ErrorSuffix = "Error"
	// For link apply error, try to clean the data plane entity.
	PhaseErrorAndCleaning = "ErrorAndCleaning"
	PhaseErrorAndCleaned  = "ErrorAndCleaned"

	// If HealthCheckServer check the K8s Node/Daemon is unhealthy, the phase will be Unhealthy.
	PhaseUnhealthy  = "Unhealthy"
	UnhealthySuffix = "Unhealthy"

	PhaseUnknown = "Unknown"

	FinalizerName    = "kubedtn.dslab.sjtu/finalizer"
	AnnotationIsFake = "kubedtn.dslab.sjtu/annotation/is-fake"
)

const (
	// Used by daemon to wait for expected phase
	PhasePollingSleepMs = 100
	PhaseWaitTimeoutMs  = 5000

	ErrorRetryDeleteCountDefault = 3
)

type VersionStatus struct {
	VersionId string `json:"version_id,omitempty"`

	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"` // may not used
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

func NewVersionStatus(now time.Time) *VersionStatus {
	tsInt := now.UnixNano()
	randomBytes := make([]byte, 8)
	rand.Read(randomBytes)

	return &VersionStatus{
		VersionId: fmt.Sprintf("%x-%x", tsInt, randomBytes),
		CreatedAt: now,
		UpdatedAt: now,
	}
}

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
