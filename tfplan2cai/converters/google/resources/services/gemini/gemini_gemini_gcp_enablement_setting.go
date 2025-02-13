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

package gemini

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const GeminiGeminiGcpEnablementSettingAssetType string = "cloudaicompanion.googleapis.com/GeminiGcpEnablementSetting"

func ResourceConverterGeminiGeminiGcpEnablementSetting() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: GeminiGeminiGcpEnablementSettingAssetType,
		Convert:   GetGeminiGeminiGcpEnablementSettingCaiObject,
	}
}

func GetGeminiGeminiGcpEnablementSettingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//cloudaicompanion.googleapis.com/projects/{{project}}/locations/{{location}}/geminiGcpEnablementSettings/{{gemini_gcp_enablement_setting_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetGeminiGeminiGcpEnablementSettingApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: GeminiGeminiGcpEnablementSettingAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudaicompanion/v1/rest",
				DiscoveryName:        "GeminiGcpEnablementSetting",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetGeminiGeminiGcpEnablementSettingApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	enableCustomerDataSharingProp, err := expandGeminiGeminiGcpEnablementSettingEnableCustomerDataSharing(d.Get("enable_customer_data_sharing"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_customer_data_sharing"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableCustomerDataSharingProp)) && (ok || !reflect.DeepEqual(v, enableCustomerDataSharingProp)) {
		obj["enableCustomerDataSharing"] = enableCustomerDataSharingProp
	}
	labelsProp, err := expandGeminiGeminiGcpEnablementSettingEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandGeminiGeminiGcpEnablementSettingEnableCustomerDataSharing(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGeminiGeminiGcpEnablementSettingEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
