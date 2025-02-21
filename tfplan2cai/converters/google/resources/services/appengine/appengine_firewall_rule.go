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

const AppEngineFirewallRuleAssetType string = "appengine.googleapis.com/FirewallRule"

func ResourceConverterAppEngineFirewallRule() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: AppEngineFirewallRuleAssetType,
		Convert:   GetAppEngineFirewallRuleCaiObject,
	}
}

func GetAppEngineFirewallRuleCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//appengine.googleapis.com/apps/{{project}}/firewall/ingressRules/{{priority}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetAppEngineFirewallRuleApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: AppEngineFirewallRuleAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/appengine/v1beta/rest",
				DiscoveryName:        "FirewallRule",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetAppEngineFirewallRuleApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandAppEngineFirewallRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	sourceRangeProp, err := expandAppEngineFirewallRuleSourceRange(d.Get("source_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceRangeProp)) && (ok || !reflect.DeepEqual(v, sourceRangeProp)) {
		obj["sourceRange"] = sourceRangeProp
	}
	actionProp, err := expandAppEngineFirewallRuleAction(d.Get("action"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("action"); !tpgresource.IsEmptyValue(reflect.ValueOf(actionProp)) && (ok || !reflect.DeepEqual(v, actionProp)) {
		obj["action"] = actionProp
	}
	priorityProp, err := expandAppEngineFirewallRulePriority(d.Get("priority"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("priority"); !tpgresource.IsEmptyValue(reflect.ValueOf(priorityProp)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}

	return obj, nil
}

func expandAppEngineFirewallRuleDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineFirewallRuleSourceRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineFirewallRuleAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineFirewallRulePriority(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
