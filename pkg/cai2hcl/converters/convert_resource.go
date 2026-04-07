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
			if true && strings.Contains(asset.Name, "folders") {
				converter = ConverterMap[asset.Type]["CloudAssetFolderFeed"]
			} else if true && strings.Contains(asset.Name, "organizations") {
				converter = ConverterMap[asset.Type]["CloudAssetOrganizationFeed"]
			} else if true && strings.Contains(asset.Name, "projects") {
				converter = ConverterMap[asset.Type]["CloudAssetProjectFeed"]
			}
		case "compute.googleapis.com/Autoscaler":
			if true && strings.Contains(asset.Name, "zones") {
				converter = ConverterMap[asset.Type]["ComputeAutoscaler"]
			} else if true && strings.Contains(asset.Name, "regions") {
				converter = ConverterMap[asset.Type]["ComputeRegionAutoscaler"]
			}
		case "compute.googleapis.com/FirewallPolicy":
			if true && strings.Contains(asset.Name, "locations") && strings.Contains(asset.Name, "global") {
				converter = ConverterMap[asset.Type]["ComputeFirewallPolicy"]
			} else if true && strings.Contains(asset.Name, "projects") && strings.Contains(asset.Name, "global") {
				converter = ConverterMap[asset.Type]["ComputeNetworkFirewallPolicy"]
			} else if true && strings.Contains(asset.Name, "projects") && strings.Contains(asset.Name, "regions") {
				converter = ConverterMap[asset.Type]["ComputeRegionNetworkFirewallPolicy"]
			}
		case "compute.googleapis.com/HealthCheck":
			if true && strings.Contains(asset.Name, "global") {
				converter = ConverterMap[asset.Type]["ComputeHealthCheck"]
			} else if true && strings.Contains(asset.Name, "regions") {
				converter = ConverterMap[asset.Type]["ComputeRegionHealthCheck"]
			}
		case "compute.googleapis.com/NetworkEndpointGroup":
			if true && strings.Contains(asset.Name, "global") {
				converter = ConverterMap[asset.Type]["ComputeGlobalNetworkEndpointGroup"]
			} else if true && strings.Contains(asset.Name, "zones") {
				converter = ConverterMap[asset.Type]["ComputeNetworkEndpointGroup"]
			} else if true && strings.Contains(asset.Name, "regions") {
				converter = ConverterMap[asset.Type]["ComputeRegionNetworkEndpointGroup"]
			}
		case "compute.googleapis.com/SslCertificate":
			if strings.Contains(asset.Name, "regions") {
				converter = ConverterMap[asset.Type]["ComputeRegionSslCertificate"]
			} else if typeVal, ok := asset.Resource.Data["type"]; ok && typeVal == "MANAGED" {
				converter = ConverterMap[asset.Type]["ComputeManagedSslCertificate"]
			} else {
				converter = ConverterMap[asset.Type]["ComputeSslCertificate"]
			}
		case "compute.googleapis.com/SslPolicy":
			if true && strings.Contains(asset.Name, "regions") {
				converter = ConverterMap[asset.Type]["ComputeRegionSslPolicy"]
			} else if true && strings.Contains(asset.Name, "global") {
				converter = ConverterMap[asset.Type]["ComputeSslPolicy"]
			}
		case "compute.googleapis.com/TargetHttpProxy":
			if true && strings.Contains(asset.Name, "regions") {
				converter = ConverterMap[asset.Type]["ComputeRegionTargetHttpProxy"]
			} else if true && strings.Contains(asset.Name, "global") {
				converter = ConverterMap[asset.Type]["ComputeTargetHttpProxy"]
			}
		case "compute.googleapis.com/TargetHttpsProxy":
			if true && strings.Contains(asset.Name, "regions") {
				converter = ConverterMap[asset.Type]["ComputeRegionTargetHttpsProxy"]
			} else if true && strings.Contains(asset.Name, "global") {
				converter = ConverterMap[asset.Type]["ComputeTargetHttpsProxy"]
			}
		case "compute.googleapis.com/TargetTcpProxy":
			if true && strings.Contains(asset.Name, "regions") {
				converter = ConverterMap[asset.Type]["ComputeRegionTargetTcpProxy"]
			} else if true && strings.Contains(asset.Name, "global") {
				converter = ConverterMap[asset.Type]["ComputeTargetTcpProxy"]
			}
		case "compute.googleapis.com/UrlMap":
			if true && strings.Contains(asset.Name, "regions") {
				converter = ConverterMap[asset.Type]["ComputeRegionUrlMap"]
			} else if true && strings.Contains(asset.Name, "global") {
				converter = ConverterMap[asset.Type]["ComputeUrlMap"]
			}
		case "secretmanager.googleapis.com/Secret":
			if true && strings.Contains(asset.Name, "locations") {
				converter = ConverterMap[asset.Type]["SecretManagerRegionalRegionalSecret"]
			} else if true {
				converter = ConverterMap[asset.Type]["SecretManagerSecret"]
			}
		case "secretmanager.googleapis.com/SecretVersion":
			if true && strings.Contains(asset.Name, "locations") {
				converter = ConverterMap[asset.Type]["SecretManagerRegionalRegionalSecretVersion"]
			} else if true {
				converter = ConverterMap[asset.Type]["SecretManagerSecretVersion"]
			}
		case "securitycenter.googleapis.com/MuteConfig":
			if true && strings.Contains(asset.Name, "folders") {
				converter = ConverterMap[asset.Type]["SecurityCenterV2FolderMuteConfig"]
			} else if true && strings.Contains(asset.Name, "organizations") {
				converter = ConverterMap[asset.Type]["SecurityCenterV2OrganizationMuteConfig"]
			}
		}
	}

	if converter == nil {
		return nil, nil
	}
	return converter.Convert(asset)
}
