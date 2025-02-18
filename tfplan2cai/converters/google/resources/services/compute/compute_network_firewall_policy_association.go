// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/NetworkFirewallPolicyAssociation.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeNetworkFirewallPolicyAssociationAssetType string = "compute.googleapis.com/NetworkFirewallPolicyAssociation"

func ResourceConverterComputeNetworkFirewallPolicyAssociation() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeNetworkFirewallPolicyAssociationAssetType,
		Convert:   GetComputeNetworkFirewallPolicyAssociationCaiObject,
	}
}

func GetComputeNetworkFirewallPolicyAssociationCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/firewallPolicies/{{firewall_policy}}/getAssociation?name={{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeNetworkFirewallPolicyAssociationApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeNetworkFirewallPolicyAssociationAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "NetworkFirewallPolicyAssociation",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeNetworkFirewallPolicyAssociationApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeNetworkFirewallPolicyAssociationName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	attachmentTargetProp, err := expandComputeNetworkFirewallPolicyAssociationAttachmentTarget(d.Get("attachment_target"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("attachment_target"); !tpgresource.IsEmptyValue(reflect.ValueOf(attachmentTargetProp)) && (ok || !reflect.DeepEqual(v, attachmentTargetProp)) {
		obj["attachmentTarget"] = attachmentTargetProp
	}

	return obj, nil
}

func expandComputeNetworkFirewallPolicyAssociationName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkFirewallPolicyAssociationAttachmentTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
