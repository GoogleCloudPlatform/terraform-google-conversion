// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/chronicle/RuleDeployment.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package chronicle

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func chronicleRuleDeploymentNilRunFrequencyDiffSuppressFunc(_, old, new string, _ *schema.ResourceData) bool {
	if new == "" {
		return true
	}
	return false
}

const ChronicleRuleDeploymentAssetType string = "{{location}}-chronicle.googleapis.com/RuleDeployment"

func ResourceConverterChronicleRuleDeployment() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ChronicleRuleDeploymentAssetType,
		Convert:   GetChronicleRuleDeploymentCaiObject,
	}
}

func GetChronicleRuleDeploymentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//{{location}}-chronicle.googleapis.com/projects/{{project}}/locations/{{location}}/instances/{{instance}}/rules/{{rule}}/deployment")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetChronicleRuleDeploymentApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ChronicleRuleDeploymentAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/{{location}}-chronicle/v1beta/rest",
				DiscoveryName:        "RuleDeployment",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetChronicleRuleDeploymentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	enabledProp, err := expandChronicleRuleDeploymentEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(enabledProp)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}
	alertingProp, err := expandChronicleRuleDeploymentAlerting(d.Get("alerting"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("alerting"); !tpgresource.IsEmptyValue(reflect.ValueOf(alertingProp)) && (ok || !reflect.DeepEqual(v, alertingProp)) {
		obj["alerting"] = alertingProp
	}
	archivedProp, err := expandChronicleRuleDeploymentArchived(d.Get("archived"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("archived"); !tpgresource.IsEmptyValue(reflect.ValueOf(archivedProp)) && (ok || !reflect.DeepEqual(v, archivedProp)) {
		obj["archived"] = archivedProp
	}
	runFrequencyProp, err := expandChronicleRuleDeploymentRunFrequency(d.Get("run_frequency"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("run_frequency"); !tpgresource.IsEmptyValue(reflect.ValueOf(runFrequencyProp)) && (ok || !reflect.DeepEqual(v, runFrequencyProp)) {
		obj["runFrequency"] = runFrequencyProp
	}

	return obj, nil
}

func expandChronicleRuleDeploymentEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleRuleDeploymentAlerting(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleRuleDeploymentArchived(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandChronicleRuleDeploymentRunFrequency(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
