package cai2hcl

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"

	"github.com/zclconf/go-cty/cty"
)

// Converter interface for resources.
type Converter interface {
	// Convert converts an Asset into a cty.Value with its ID.
	Convert(asset *caiasset.Asset) (string, cty.Value, error)
	// ConvertIAM converts an Asset's iam policy into a cty.Value with its ID.
	ConvertIAM(asset *caiasset.Asset) (string, cty.Value, error)
	// TFResourceType returns terraform resource type.
	TFResourceType() string
}

// Converters returns a map of resource converters.
// The map key is the CAI Asset type.
func Converters() map[string]Converter {
	return map[string]Converter{
		ComputeInstanceAssetType: NewComputeInstanceConverter(),
	}
}
