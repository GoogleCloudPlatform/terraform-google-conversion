// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networksecurity/InterceptDeployment.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package networksecurity

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetworkSecurityInterceptDeploymentAssetType string = "networksecurity.googleapis.com/InterceptDeployment"

func ResourceConverterNetworkSecurityInterceptDeployment() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetworkSecurityInterceptDeploymentAssetType,
		Convert:   GetNetworkSecurityInterceptDeploymentCaiObject,
	}
}

func GetNetworkSecurityInterceptDeploymentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//networksecurity.googleapis.com/projects/{{project}}/locations/{{location}}/interceptDeployments/{{intercept_deployment_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetworkSecurityInterceptDeploymentApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetworkSecurityInterceptDeploymentAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networksecurity/v1beta1/rest",
				DiscoveryName:        "InterceptDeployment",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetworkSecurityInterceptDeploymentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	forwardingRuleProp, err := expandNetworkSecurityInterceptDeploymentForwardingRule(d.Get("forwarding_rule"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("forwarding_rule"); !tpgresource.IsEmptyValue(reflect.ValueOf(forwardingRuleProp)) && (ok || !reflect.DeepEqual(v, forwardingRuleProp)) {
		obj["forwardingRule"] = forwardingRuleProp
	}
	interceptDeploymentGroupProp, err := expandNetworkSecurityInterceptDeploymentInterceptDeploymentGroup(d.Get("intercept_deployment_group"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("intercept_deployment_group"); !tpgresource.IsEmptyValue(reflect.ValueOf(interceptDeploymentGroupProp)) && (ok || !reflect.DeepEqual(v, interceptDeploymentGroupProp)) {
		obj["interceptDeploymentGroup"] = interceptDeploymentGroupProp
	}
	labelsProp, err := expandNetworkSecurityInterceptDeploymentEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetworkSecurityInterceptDeploymentForwardingRule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityInterceptDeploymentInterceptDeploymentGroup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityInterceptDeploymentEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
