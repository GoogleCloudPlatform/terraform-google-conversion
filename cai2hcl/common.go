package cai2hcl

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ConvertFunc converts an Asset type into a map[string]interface{} struct.
type ConvertFunc func(asset *caiasset.Asset) (string, map[string]interface{}, error)

// Converter for individual resource.
type Converter struct {
	TFResourceName string
	Convert        ConvertFunc
	Resource       *tfschema.Resource
}

// Converters returns a map of resource converters.
// The map key is the CAI Asset type.
func Converters() map[string]*Converter {
	return map[string]*Converter{
		ComputeInstanceAssetType: NewComputeInstanceConverter(),
	}
}
