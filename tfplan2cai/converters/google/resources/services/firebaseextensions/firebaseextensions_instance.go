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

package firebaseextensions

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const FirebaseExtensionsInstanceAssetType string = "firebaseextensions.googleapis.com/Instance"

func ResourceConverterFirebaseExtensionsInstance() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: FirebaseExtensionsInstanceAssetType,
		Convert:   GetFirebaseExtensionsInstanceCaiObject,
	}
}

func GetFirebaseExtensionsInstanceCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//firebaseextensions.googleapis.com/projects/{{project}}/instances/{{instance_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetFirebaseExtensionsInstanceApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: FirebaseExtensionsInstanceAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/firebaseextensions/v1beta/rest",
				DiscoveryName:        "Instance",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetFirebaseExtensionsInstanceApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	configProp, err := expandFirebaseExtensionsInstanceConfig(d.Get("config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("config"); !tpgresource.IsEmptyValue(reflect.ValueOf(configProp)) && (ok || !reflect.DeepEqual(v, configProp)) {
		obj["config"] = configProp
	}
	etagProp, err := expandFirebaseExtensionsInstanceEtag(d.Get("etag"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}

	return obj, nil
}

func expandFirebaseExtensionsInstanceConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandFirebaseExtensionsInstanceConfigName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["name"] = transformedName
	}

	transformedCreateTime, err := expandFirebaseExtensionsInstanceConfigCreateTime(original["create_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCreateTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["createTime"] = transformedCreateTime
	}

	transformedParams, err := expandFirebaseExtensionsInstanceConfigParams(original["params"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedParams); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["params"] = transformedParams
	}

	transformedSystemParams, err := expandFirebaseExtensionsInstanceConfigSystemParams(original["system_params"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSystemParams); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["systemParams"] = transformedSystemParams
	}

	transformedExtensionRef, err := expandFirebaseExtensionsInstanceConfigExtensionRef(original["extension_ref"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExtensionRef); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["extensionRef"] = transformedExtensionRef
	}

	transformedExtensionVersion, err := expandFirebaseExtensionsInstanceConfigExtensionVersion(original["extension_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExtensionVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["extensionVersion"] = transformedExtensionVersion
	}

	transformedAllowedEventTypes, err := expandFirebaseExtensionsInstanceConfigAllowedEventTypes(original["allowed_event_types"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllowedEventTypes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allowedEventTypes"] = transformedAllowedEventTypes
	}

	transformedEventarcChannel, err := expandFirebaseExtensionsInstanceConfigEventarcChannel(original["eventarc_channel"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEventarcChannel); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["eventarcChannel"] = transformedEventarcChannel
	}

	transformedPopulatedPostinstallContent, err := expandFirebaseExtensionsInstanceConfigPopulatedPostinstallContent(original["populated_postinstall_content"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPopulatedPostinstallContent); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["populatedPostinstallContent"] = transformedPopulatedPostinstallContent
	}

	return transformed, nil
}

func expandFirebaseExtensionsInstanceConfigName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseExtensionsInstanceConfigCreateTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseExtensionsInstanceConfigParams(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandFirebaseExtensionsInstanceConfigSystemParams(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandFirebaseExtensionsInstanceConfigExtensionRef(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseExtensionsInstanceConfigExtensionVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseExtensionsInstanceConfigAllowedEventTypes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseExtensionsInstanceConfigEventarcChannel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseExtensionsInstanceConfigPopulatedPostinstallContent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseExtensionsInstanceEtag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
