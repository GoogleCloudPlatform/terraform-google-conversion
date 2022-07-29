package conversion

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/model"
	"github.com/GoogleCloudPlatform/terraform-validator/ancestrymanager"
	"github.com/GoogleCloudPlatform/terraform-validator/converters/google"
	resources "github.com/GoogleCloudPlatform/terraform-validator/converters/google/resources"
	"github.com/GoogleCloudPlatform/terraform-validator/tfplan"
	"go.uber.org/zap"
)

// Single struct for options so that adding new options does not require
// updating function signatures all along the pipe
type TFPlanToCAIOptions struct {
	ConvertUnchanged bool
	ErrorLogger      *zap.Logger
	Offline          bool
	DefaultProject   string
	// TODO:
	// DefaultRegion    string
	// DefaultZone      string
	// UserAgent for all requests (if online)
	UserAgent string
	// HTTPClient for all requests (if online)
	HTTPClient *http.Client
	// Map hierarchy resource (like projects/<number> or folders/<number>)
	// to an ancestry path (like organizations/123/folders/456/projects/789)
	AncestryCache map[string]string
}

type TFPlanToCAIFunc func(ctx context.Context, jsonPlan []byte, o *TFPlanToCAIOptions) ([]model.Asset, error)

func TFPlanToCAI(ctx context.Context, jsonPlan []byte, o *TFPlanToCAIOptions) ([]model.Asset, error) {
	// Creates ancestry manager and converter internally; they are
	// implementation details private to this package.

	errorLogger := o.ErrorLogger
	if errorLogger == nil {
		logger, err := zap.NewDevelopment()
		if err != nil {
			return nil, fmt.Errorf("building logger: %w", err)
		}
		errorLogger = logger
	}

	// contents of newConverter
	// Set up config and ancestry manager using the same http client and user agent
	cfg, err := resources.NewConfig(ctx, o.DefaultProject, o.Offline, o.UserAgent, o.HTTPClient)
	if err != nil {
		return nil, fmt.Errorf("building config: %w", err)
	}

	ancestryManager, err := ancestrymanager.New(cfg, o.Offline, o.AncestryCache, errorLogger)
	if err != nil {
		return nil, fmt.Errorf("building ancestry manager: %w", err)
	}

	converter := google.NewConverter(cfg, ancestryManager, o.Offline, o.ConvertUnchanged, errorLogger)

	// ReadResourceChanges
	changes, err := tfplan.ReadResourceChanges(jsonPlan)
	if err != nil {
		return nil, err
	}

	err = converter.AddResourceChanges(changes)
	if err != nil {
		return nil, err
	}

	return converter.Assets(), nil
}
