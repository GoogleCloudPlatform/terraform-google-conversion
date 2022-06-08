package conversion

import (
	"go.uber.org/zap"
	"net/http"
)

// Single struct for options so that adding new options does not require
// updating function signatures all along the pipe
type TFPlanToCAIOptions struct {
	ConvertUnchanged bool
	ErrorLogger      *zap.Logger
	Offline          bool
	DefaultProject   string
	DefaultRegion    string
	DefaultZone      string
	// UserAgent for all requests (if online)
	UserAgent string
	// HTTPClient for all requests (if online)
	HTTPClient *http.Client
	// Map hierarchy resource (like projects/<number> or folders/<number>)
	// to a cached ancestry
	AncestryCache map[string]string
}

func TFPlanToCAI(ctx context.Context, jsonPlan []byte, options *TFPlanToCAIOptions) ([]Asset, error) {
	// Creates ancestry manager and converter internally; they are
	// implementation details private to this package.
	var assets []Asset
	return assets, nil
}
