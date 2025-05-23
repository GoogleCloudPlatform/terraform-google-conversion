// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networkservices/LbTrafficExtension.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package networkservices

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetworkServicesLbTrafficExtensionAssetType string = "networkservices.googleapis.com/LbTrafficExtension"

func ResourceConverterNetworkServicesLbTrafficExtension() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetworkServicesLbTrafficExtensionAssetType,
		Convert:   GetNetworkServicesLbTrafficExtensionCaiObject,
	}
}

func GetNetworkServicesLbTrafficExtensionCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//networkservices.googleapis.com/projects/{{project}}/locations/{{location}}/lbTrafficExtensions/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetworkServicesLbTrafficExtensionApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetworkServicesLbTrafficExtensionAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networkservices/v1/rest",
				DiscoveryName:        "LbTrafficExtension",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetworkServicesLbTrafficExtensionApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkServicesLbTrafficExtensionDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	forwardingRulesProp, err := expandNetworkServicesLbTrafficExtensionForwardingRules(d.Get("forwarding_rules"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("forwarding_rules"); !tpgresource.IsEmptyValue(reflect.ValueOf(forwardingRulesProp)) && (ok || !reflect.DeepEqual(v, forwardingRulesProp)) {
		obj["forwardingRules"] = forwardingRulesProp
	}
	extensionChainsProp, err := expandNetworkServicesLbTrafficExtensionExtensionChains(d.Get("extension_chains"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("extension_chains"); !tpgresource.IsEmptyValue(reflect.ValueOf(extensionChainsProp)) && (ok || !reflect.DeepEqual(v, extensionChainsProp)) {
		obj["extensionChains"] = extensionChainsProp
	}
	loadBalancingSchemeProp, err := expandNetworkServicesLbTrafficExtensionLoadBalancingScheme(d.Get("load_balancing_scheme"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("load_balancing_scheme"); !tpgresource.IsEmptyValue(reflect.ValueOf(loadBalancingSchemeProp)) && (ok || !reflect.DeepEqual(v, loadBalancingSchemeProp)) {
		obj["loadBalancingScheme"] = loadBalancingSchemeProp
	}
	labelsProp, err := expandNetworkServicesLbTrafficExtensionEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetworkServicesLbTrafficExtensionDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesLbTrafficExtensionForwardingRules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesLbTrafficExtensionExtensionChains(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandNetworkServicesLbTrafficExtensionExtensionChainsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedMatchCondition, err := expandNetworkServicesLbTrafficExtensionExtensionChainsMatchCondition(original["match_condition"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMatchCondition); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["matchCondition"] = transformedMatchCondition
		}

		transformedExtensions, err := expandNetworkServicesLbTrafficExtensionExtensionChainsExtensions(original["extensions"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedExtensions); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["extensions"] = transformedExtensions
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkServicesLbTrafficExtensionExtensionChainsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesLbTrafficExtensionExtensionChainsMatchCondition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCelExpression, err := expandNetworkServicesLbTrafficExtensionExtensionChainsMatchConditionCelExpression(original["cel_expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCelExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["celExpression"] = transformedCelExpression
	}

	return transformed, nil
}

func expandNetworkServicesLbTrafficExtensionExtensionChainsMatchConditionCelExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesLbTrafficExtensionExtensionChainsExtensions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandNetworkServicesLbTrafficExtensionExtensionChainsExtensionsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedAuthority, err := expandNetworkServicesLbTrafficExtensionExtensionChainsExtensionsAuthority(original["authority"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAuthority); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["authority"] = transformedAuthority
		}

		transformedService, err := expandNetworkServicesLbTrafficExtensionExtensionChainsExtensionsService(original["service"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedService); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["service"] = transformedService
		}

		transformedTimeout, err := expandNetworkServicesLbTrafficExtensionExtensionChainsExtensionsTimeout(original["timeout"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTimeout); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["timeout"] = transformedTimeout
		}

		transformedFailOpen, err := expandNetworkServicesLbTrafficExtensionExtensionChainsExtensionsFailOpen(original["fail_open"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFailOpen); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["failOpen"] = transformedFailOpen
		}

		transformedForwardHeaders, err := expandNetworkServicesLbTrafficExtensionExtensionChainsExtensionsForwardHeaders(original["forward_headers"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedForwardHeaders); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["forwardHeaders"] = transformedForwardHeaders
		}

		transformedSupportedEvents, err := expandNetworkServicesLbTrafficExtensionExtensionChainsExtensionsSupportedEvents(original["supported_events"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSupportedEvents); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["supportedEvents"] = transformedSupportedEvents
		}

		transformedMetadata, err := expandNetworkServicesLbTrafficExtensionExtensionChainsExtensionsMetadata(original["metadata"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMetadata); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["metadata"] = transformedMetadata
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkServicesLbTrafficExtensionExtensionChainsExtensionsName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesLbTrafficExtensionExtensionChainsExtensionsAuthority(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesLbTrafficExtensionExtensionChainsExtensionsService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesLbTrafficExtensionExtensionChainsExtensionsTimeout(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesLbTrafficExtensionExtensionChainsExtensionsFailOpen(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesLbTrafficExtensionExtensionChainsExtensionsForwardHeaders(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesLbTrafficExtensionExtensionChainsExtensionsSupportedEvents(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesLbTrafficExtensionExtensionChainsExtensionsMetadata(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandNetworkServicesLbTrafficExtensionLoadBalancingScheme(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkServicesLbTrafficExtensionEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
