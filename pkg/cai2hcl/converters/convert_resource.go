package converters

import (
	"strings"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/cai2hcl/models"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/caiasset"
)

func ConvertResource(asset caiasset.Asset) ([]*models.TerraformResourceBlock, error) {
	converters, ok := ConverterMap[asset.Type]
	if !ok || len(converters) == 0 {
		return nil, nil
	}

	var converter models.Cai2hclConverter
	// Normally, one asset type has only one converter.
	if len(converters) == 1 {
		converter = converters["Default"]
	} else {
		// Edge cases
		// Handle the edge case that multiple Terraform resources share the same CAI asset type
		switch asset.Type {
		case "cloudasset.googleapis.com/Feed":
			if strings.Contains(asset.Name, "folders") {
				converter = ConverterMap[asset.Type]["CloudAssetFolderFeed"]
			} else if strings.Contains(asset.Name, "organizations") {
				converter = ConverterMap[asset.Type]["CloudAssetOrganizationFeed"]
			} else if strings.Contains(asset.Name, "projects") {
				converter = ConverterMap[asset.Type]["CloudAssetProjectFeed"]
			}
		case "compute.googleapis.com/Autoscaler":
			if strings.Contains(asset.Name, "zones") {
				converter = ConverterMap[asset.Type]["ComputeAutoscaler"]
			} else if strings.Contains(asset.Name, "regions") {
				converter = ConverterMap[asset.Type]["ComputeRegionAutoscaler"]
			}
		case "compute.googleapis.com/HealthCheck":
			if strings.Contains(asset.Name, "global") {
				converter = ConverterMap[asset.Type]["ComputeHealthCheck"]
			} else if strings.Contains(asset.Name, "regions") {
				converter = ConverterMap[asset.Type]["ComputeRegionHealthCheck"]
			}
		case "secretmanager.googleapis.com/Secret":
			if strings.Contains(asset.Name, "locations") {
				converter = ConverterMap[asset.Type]["SecretManagerRegionalRegionalSecret"]
			} else {
				converter = ConverterMap[asset.Type]["SecretManagerSecret"]
			}
		case "secretmanager.googleapis.com/SecretVersion":
			if strings.Contains(asset.Name, "") {
				converter = ConverterMap[asset.Type]["SecretManagerRegionalRegionalSecretVersion"]
			} else {
				converter = ConverterMap[asset.Type]["SecretManagerSecretVersion"]
			}
		}
	}

	if converter == nil {
		return nil, nil
	}
	return converter.Convert(asset)
}
