package tpgresource

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
	cloudresourcemanager "google.golang.org/api/cloudresourcemanager/v1"
)

// Wraps Op.Error in an implementation of built-in Error
type CommonOpError struct {
	*cloudresourcemanager.Status
}

func (e *CommonOpError) Error() string {
	return fmt.Sprintf("Error code %v, message: %s", e.Code, e.Message)
}

type Waiter interface {
	// State returns the current status of the operation.
	State() string

	// Error returns an error embedded in the operation we're waiting on, or nil
	// if the operation has no current error.
	Error() error

	// IsRetryable returns whether a given error should be retried.
	IsRetryable(error) bool

	// SetOp sets the operation we're waiting on in a Waiter struct so that it
	// can be used in other methods.
	SetOp(interface{}) error

	// QueryOp sends a request to the server to get the current status of the
	// operation. It's expected that QueryOp will return exactly one of an
	// operation or an error as non-nil, and that requests will be retried by
	// specific implementations of the method.
	QueryOp() (interface{}, error)

	// OpName is the name of the operation and is used to log its status.
	OpName() string

	// PendingStates contains the values of State() that cause us to continue
	// refreshing the operation.
	PendingStates() []string

	// TargetStates contain the values of State() that cause us to finish
	// refreshing the operation.
	TargetStates() []string
}

type CommonOperationWaiter struct {
	Op CommonOperation
}

func (w *CommonOperationWaiter) State() string {
	if w == nil {
		return fmt.Sprintf("Operation is nil!")
	}

	return fmt.Sprintf("done: %v", w.Op.Done)
}

func (w *CommonOperationWaiter) Error() error {
	if w != nil && w.Op.Error != nil {
		return &CommonOpError{w.Op.Error}
	}
	return nil
}

func (w *CommonOperationWaiter) IsRetryable(error) bool {
	return false
}

func (w *CommonOperationWaiter) SetOp(op interface{}) error {
	if err := Convert(op, &w.Op); err != nil {
		return err
	}
	return nil
}

func (w *CommonOperationWaiter) OpName() string {
	if w == nil {
		return "<nil>"
	}

	return w.Op.Name
}

func (w *CommonOperationWaiter) PendingStates() []string {
	return []string{"done: false"}
}

func (w *CommonOperationWaiter) TargetStates() []string {
	return []string{"done: true"}
}

func OperationDone(w Waiter) bool {
	for _, s := range w.TargetStates() {
		if s == w.State() {
			return true
		}
	}
	return false
}

func CommonRefreshFunc(w Waiter) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		op, err := w.QueryOp()
		if err != nil {
			// Retry 404 when getting operation (not resource state)
			if transport_tpg.IsRetryableError(err, transport_tpg.IsNotFoundRetryableError("GET operation")) {
				log.Printf("[DEBUG] Dismissed retryable error on GET operation %q: %s", w.OpName(), err)
				return nil, "done: false", nil
			}
			return nil, "", fmt.Errorf("error while retrieving operation: %s", err)
		}

		if err = w.SetOp(op); err != nil {
			return nil, "", fmt.Errorf("Cannot continue, unable to use operation: %s", err)
		}

		if err = w.Error(); err != nil {
			if w.IsRetryable(err) {
				log.Printf("[DEBUG] Retrying operation GET based on retryable err: %s", err)
				return nil, w.State(), nil
			}
			return nil, "", err
		}

		log.Printf("[DEBUG] Got %v while polling for operation %s's status", w.State(), w.OpName())
		return op, w.State(), nil
	}
}

func OperationWait(w Waiter, activity string, timeout time.Duration, pollInterval time.Duration) error {
	if OperationDone(w) {
		if w.Error() != nil {
			return w.Error()
		}
		return nil
	}

	c := &resource.StateChangeConf{
		Pending:      w.PendingStates(),
		Target:       w.TargetStates(),
		Refresh:      CommonRefreshFunc(w),
		Timeout:      timeout,
		MinTimeout:   2 * time.Second,
		PollInterval: pollInterval,
	}
	opRaw, err := c.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for %s: %s", activity, err)
	}

	err = w.SetOp(opRaw)
	if err != nil {
		return err
	}
	if w.Error() != nil {
		return w.Error()
	}

	return nil
}

// The cloud resource manager API operation is an example of one of many
// interchangeable API operations. Choose it somewhat arbitrarily to represent
// the "common" operation.
type CommonOperation cloudresourcemanager.Operation
