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

package networksecurity

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetworkSecurityMirroringDeploymentAssetType string = "networksecurity.googleapis.com/MirroringDeployment"

func ResourceConverterNetworkSecurityMirroringDeployment() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetworkSecurityMirroringDeploymentAssetType,
		Convert:   GetNetworkSecurityMirroringDeploymentCaiObject,
	}
}

func GetNetworkSecurityMirroringDeploymentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//networksecurity.googleapis.com/projects/{{project}}/locations/{{location}}/mirroringDeployments/{{mirroring_deployment_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetworkSecurityMirroringDeploymentApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetworkSecurityMirroringDeploymentAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networksecurity/v1beta1/rest",
				DiscoveryName:        "MirroringDeployment",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetworkSecurityMirroringDeploymentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	forwardingRuleProp, err := expandNetworkSecurityMirroringDeploymentForwardingRule(d.Get("forwarding_rule"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("forwarding_rule"); !tpgresource.IsEmptyValue(reflect.ValueOf(forwardingRuleProp)) && (ok || !reflect.DeepEqual(v, forwardingRuleProp)) {
		obj["forwardingRule"] = forwardingRuleProp
	}
	mirroringDeploymentGroupProp, err := expandNetworkSecurityMirroringDeploymentMirroringDeploymentGroup(d.Get("mirroring_deployment_group"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("mirroring_deployment_group"); !tpgresource.IsEmptyValue(reflect.ValueOf(mirroringDeploymentGroupProp)) && (ok || !reflect.DeepEqual(v, mirroringDeploymentGroupProp)) {
		obj["mirroringDeploymentGroup"] = mirroringDeploymentGroupProp
	}
	labelsProp, err := expandNetworkSecurityMirroringDeploymentEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetworkSecurityMirroringDeploymentForwardingRule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityMirroringDeploymentMirroringDeploymentGroup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityMirroringDeploymentEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
