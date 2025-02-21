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

package appengine

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const AppEngineApplicationUrlDispatchRulesAssetType string = "appengine.googleapis.com/ApplicationUrlDispatchRules"

func ResourceConverterAppEngineApplicationUrlDispatchRules() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: AppEngineApplicationUrlDispatchRulesAssetType,
		Convert:   GetAppEngineApplicationUrlDispatchRulesCaiObject,
	}
}

func GetAppEngineApplicationUrlDispatchRulesCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//appengine.googleapis.com/apps/{{project}}/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetAppEngineApplicationUrlDispatchRulesApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: AppEngineApplicationUrlDispatchRulesAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/appengine/v1beta/rest",
				DiscoveryName:        "ApplicationUrlDispatchRules",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetAppEngineApplicationUrlDispatchRulesApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	dispatchRulesProp, err := expandAppEngineApplicationUrlDispatchRulesDispatchRules(d.Get("dispatch_rules"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dispatch_rules"); !tpgresource.IsEmptyValue(reflect.ValueOf(dispatchRulesProp)) && (ok || !reflect.DeepEqual(v, dispatchRulesProp)) {
		obj["dispatchRules"] = dispatchRulesProp
	}

	return obj, nil
}

func expandAppEngineApplicationUrlDispatchRulesDispatchRules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDomain, err := expandAppEngineApplicationUrlDispatchRulesDispatchRulesDomain(original["domain"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDomain); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["domain"] = transformedDomain
		}

		transformedPath, err := expandAppEngineApplicationUrlDispatchRulesDispatchRulesPath(original["path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["path"] = transformedPath
		}

		transformedService, err := expandAppEngineApplicationUrlDispatchRulesDispatchRulesService(original["service"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedService); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["service"] = transformedService
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandAppEngineApplicationUrlDispatchRulesDispatchRulesDomain(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineApplicationUrlDispatchRulesDispatchRulesPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineApplicationUrlDispatchRulesDispatchRulesService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
