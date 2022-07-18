package conversion

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/conversion/hcl"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

type Converter struct {
	provider *schema.Provider
	opts     *CAIToHCLOptions
}

// Single struct for options so that adding new options does not require
// updating function signatures all along the pipe
type CAIToHCLOptions struct {
	ErrorLogger *zap.Logger
}

// Returns HCL as a byte slice.
func (c *Converter) CAIToHCL(assets []*hcl.Asset, options *CAIToHCLOptions) ([]byte, error) {
	// For now, no assets are supported.
	var hclBytes []byte
	for _, asset := range assets {
		converter, ok := hcl.Converters()[asset.Type]
		if !ok {
			c.opts.ErrorLogger.Warn(fmt.Sprintf("type %s is not supported", asset.Type))
			continue
		}

		tfResource, ok := c.provider.ResourcesMap[converter.TFResourceName]
		if !ok {
			c.opts.ErrorLogger.Warn(fmt.Sprintf("provider resource map does not have type %s", asset.Type))
			continue
		}

		id, data, err := converter.Convert(asset)
		if err != nil {
			c.opts.ErrorLogger.Warn(fmt.Sprintf("failed to convert asset: %s", err))
			continue
		}

		val := helper.MapToCtyValWithSchema(data, tfResource.Schema)
		instanceState := terraform.NewInstanceStateShimmedFromValue(val, tfResource.SchemaVersion)

		info := &terraform.InstanceInfo{
			Id:   id,
			Type: converter.TFResourceName,
		}
		ret, err := helper.InstanceStateToHCL(instanceState, info, c.provider)
		if err != nil {
			return nil, err
		}
		hclBytes = append(hclBytes, []byte(ret)...)
	}
	return hclBytes, nil
}
