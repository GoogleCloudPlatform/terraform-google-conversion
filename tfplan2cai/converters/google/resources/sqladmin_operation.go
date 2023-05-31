package google

import (
	"time"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/services/sql"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

// Deprecated: For backward compatibility SqlAdminOperationWaitTime is still working,
// but all new code should use SqlAdminOperationWaitTime in the sql package instead.
func SqlAdminOperationWaitTime(config *transport_tpg.Config, res interface{}, project, activity, userAgent string, timeout time.Duration) error {
	return sql.SqlAdminOperationWaitTime(config, res, project, activity, userAgent, timeout)
}

// Retry if Cloud SQL operation returns a 429 with a specific message for
// concurrent operations.
//
// Deprecated: For backward compatibility IsSqlInternalError is still working,
// but all new code should use IsSqlInternalError in the sql package instead.
func IsSqlInternalError(err error) (bool, string) {
	return sql.IsSqlInternalError(err)
}
