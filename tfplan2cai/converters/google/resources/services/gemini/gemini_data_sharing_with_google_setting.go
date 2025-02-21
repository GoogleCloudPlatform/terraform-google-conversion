// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/gemini/DataSharingWithGoogleSetting.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package gemini

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const GeminiDataSharingWithGoogleSettingAssetType string = "cloudaicompanion.googleapis.com/DataSharingWithGoogleSetting"

func ResourceConverterGeminiDataSharingWithGoogleSetting() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: GeminiDataSharingWithGoogleSettingAssetType,
		Convert:   GetGeminiDataSharingWithGoogleSettingCaiObject,
	}
}

func GetGeminiDataSharingWithGoogleSettingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//cloudaicompanion.googleapis.com/projects/{{project}}/locations/{{location}}/dataSharingWithGoogleSettings/{{data_sharing_with_google_setting_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetGeminiDataSharingWithGoogleSettingApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: GeminiDataSharingWithGoogleSettingAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudaicompanion/v1/rest",
				DiscoveryName:        "DataSharingWithGoogleSetting",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetGeminiDataSharingWithGoogleSettingApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	enablePreviewDataSharingProp, err := expandGeminiDataSharingWithGoogleSettingEnablePreviewDataSharing(d.Get("enable_preview_data_sharing"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_preview_data_sharing"); !tpgresource.IsEmptyValue(reflect.ValueOf(enablePreviewDataSharingProp)) && (ok || !reflect.DeepEqual(v, enablePreviewDataSharingProp)) {
		obj["enablePreviewDataSharing"] = enablePreviewDataSharingProp
	}
	labelsProp, err := expandGeminiDataSharingWithGoogleSettingEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandGeminiDataSharingWithGoogleSettingEnablePreviewDataSharing(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGeminiDataSharingWithGoogleSettingEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
