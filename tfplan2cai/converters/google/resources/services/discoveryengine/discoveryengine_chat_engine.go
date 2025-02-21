// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/discoveryengine/ChatEngine.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package discoveryengine

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DiscoveryEngineChatEngineAssetType string = "{{location}}-discoveryengine.googleapis.com/ChatEngine"

func ResourceConverterDiscoveryEngineChatEngine() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DiscoveryEngineChatEngineAssetType,
		Convert:   GetDiscoveryEngineChatEngineCaiObject,
	}
}

func GetDiscoveryEngineChatEngineCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//{{location}}-discoveryengine.googleapis.com/projects/{{project}}/locations/{{location}}/collections/{{collection_id}}/engines/{{engine_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDiscoveryEngineChatEngineApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DiscoveryEngineChatEngineAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/{{location}}-discoveryengine/v1/rest",
				DiscoveryName:        "ChatEngine",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDiscoveryEngineChatEngineApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	industryVerticalProp, err := expandDiscoveryEngineChatEngineIndustryVertical(d.Get("industry_vertical"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("industry_vertical"); !tpgresource.IsEmptyValue(reflect.ValueOf(industryVerticalProp)) && (ok || !reflect.DeepEqual(v, industryVerticalProp)) {
		obj["industryVertical"] = industryVerticalProp
	}
	displayNameProp, err := expandDiscoveryEngineChatEngineDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	dataStoreIdsProp, err := expandDiscoveryEngineChatEngineDataStoreIds(d.Get("data_store_ids"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("data_store_ids"); !tpgresource.IsEmptyValue(reflect.ValueOf(dataStoreIdsProp)) && (ok || !reflect.DeepEqual(v, dataStoreIdsProp)) {
		obj["dataStoreIds"] = dataStoreIdsProp
	}
	chatEngineConfigProp, err := expandDiscoveryEngineChatEngineChatEngineConfig(d.Get("chat_engine_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("chat_engine_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(chatEngineConfigProp)) && (ok || !reflect.DeepEqual(v, chatEngineConfigProp)) {
		obj["chatEngineConfig"] = chatEngineConfigProp
	}
	commonConfigProp, err := expandDiscoveryEngineChatEngineCommonConfig(d.Get("common_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("common_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(commonConfigProp)) && (ok || !reflect.DeepEqual(v, commonConfigProp)) {
		obj["commonConfig"] = commonConfigProp
	}

	return resourceDiscoveryEngineChatEngineEncoder(d, config, obj)
}

func resourceDiscoveryEngineChatEngineEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// hard code solutionType to "SOLUTION_TYPE_CHAT" for chat engine resource
	obj["solutionType"] = "SOLUTION_TYPE_CHAT"
	return obj, nil
}

func expandDiscoveryEngineChatEngineIndustryVertical(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineChatEngineDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineChatEngineDataStoreIds(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineChatEngineChatEngineConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAgentCreationConfig, err := expandDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfig(original["agent_creation_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAgentCreationConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["agentCreationConfig"] = transformedAgentCreationConfig
	}

	transformedDialogflowAgentToLink, err := expandDiscoveryEngineChatEngineChatEngineConfigDialogflowAgentToLink(original["dialogflow_agent_to_link"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDialogflowAgentToLink); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dialogflowAgentToLink"] = transformedDialogflowAgentToLink
	}

	return transformed, nil
}

func expandDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBusiness, err := expandDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfigBusiness(original["business"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBusiness); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["business"] = transformedBusiness
	}

	transformedDefaultLanguageCode, err := expandDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfigDefaultLanguageCode(original["default_language_code"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDefaultLanguageCode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["defaultLanguageCode"] = transformedDefaultLanguageCode
	}

	transformedTimeZone, err := expandDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfigTimeZone(original["time_zone"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimeZone); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["timeZone"] = transformedTimeZone
	}

	transformedLocation, err := expandDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfigLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfigBusiness(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfigDefaultLanguageCode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfigTimeZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfigLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineChatEngineChatEngineConfigDialogflowAgentToLink(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDiscoveryEngineChatEngineCommonConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCompanyName, err := expandDiscoveryEngineChatEngineCommonConfigCompanyName(original["company_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCompanyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["companyName"] = transformedCompanyName
	}

	return transformed, nil
}

func expandDiscoveryEngineChatEngineCommonConfigCompanyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
