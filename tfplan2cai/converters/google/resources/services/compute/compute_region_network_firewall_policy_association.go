// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package compute

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeRegionNetworkFirewallPolicyAssociationAssetType string = "compute.googleapis.com/RegionNetworkFirewallPolicyAssociation"

func ResourceConverterComputeRegionNetworkFirewallPolicyAssociation() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeRegionNetworkFirewallPolicyAssociationAssetType,
		Convert:   GetComputeRegionNetworkFirewallPolicyAssociationCaiObject,
	}
}

func GetComputeRegionNetworkFirewallPolicyAssociationCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/firewallPolicies/{{firewall_policy}}/getAssociation?name={{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeRegionNetworkFirewallPolicyAssociationApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeRegionNetworkFirewallPolicyAssociationAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "RegionNetworkFirewallPolicyAssociation",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeRegionNetworkFirewallPolicyAssociationApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeRegionNetworkFirewallPolicyAssociationName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	attachmentTargetProp, err := expandComputeRegionNetworkFirewallPolicyAssociationAttachmentTarget(d.Get("attachment_target"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("attachment_target"); !tpgresource.IsEmptyValue(reflect.ValueOf(attachmentTargetProp)) && (ok || !reflect.DeepEqual(v, attachmentTargetProp)) {
		obj["attachmentTarget"] = attachmentTargetProp
	}

	return obj, nil
}

func expandComputeRegionNetworkFirewallPolicyAssociationName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionNetworkFirewallPolicyAssociationAttachmentTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
