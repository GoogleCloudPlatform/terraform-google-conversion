package google

import (
	"time"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/compute"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

// Deprecated: For backward compatibility ComputeOperationWaitTime is still working,
// but all new code should use ComputeOperationWaitTime in the compute package instead.
func ComputeOperationWaitTime(config *transport_tpg.Config, res interface{}, project, activity, userAgent string, timeout time.Duration) error {
	return compute.ComputeOperationWaitTime(config, res, project, activity, userAgent, timeout)
}

// Deprecated: For backward compatibility ComputeOrgOperationWaitTimeWithResponse is still working,
// but all new code should use ComputeOrgOperationWaitTimeWithResponse in the compute package instead.
func ComputeOrgOperationWaitTimeWithResponse(config *transport_tpg.Config, res interface{}, response *map[string]interface{}, parent, activity, userAgent string, timeout time.Duration) error {
	return compute.ComputeOrgOperationWaitTimeWithResponse(config, res, response, parent, activity, userAgent, timeout)
}
