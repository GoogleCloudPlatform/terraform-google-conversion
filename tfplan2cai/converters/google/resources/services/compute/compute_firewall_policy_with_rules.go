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

package compute

import (
	"fmt"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func firewallPolicyWithRulesConvertPriorityToInt(v interface{}) (int64, error) {
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal, nil
		}
	}

	if intVal, ok := v.(int64); ok {
		return intVal, nil
	}

	if floatVal, ok := v.(float64); ok {
		intVal := int64(floatVal)
		return intVal, nil
	}

	return 0, fmt.Errorf("Incorrect rule priority: %s. Priority must be a number", v)
}

func firewallPolicyWithRulesIsPredefinedRule(rule map[string]interface{}) (bool, error) {
	// Priorities from 2147483548 to 2147483647 are reserved and cannot be modified by the user.
	const ReservedPriorityStart = 2147483548

	priority := rule["priority"]
	priorityInt, err := firewallPolicyWithRulesConvertPriorityToInt(priority)

	if err != nil {
		return false, err
	}

	return priorityInt >= ReservedPriorityStart, nil
}

func firewallPolicyWithRulesSplitPredefinedRules(allRules []interface{}) ([]interface{}, []interface{}, error) {
	predefinedRules := make([]interface{}, 0)
	rules := make([]interface{}, 0)
	for _, rule := range allRules {
		isPredefined, err := firewallPolicyWithRulesIsPredefinedRule(rule.(map[string]interface{}))
		if err != nil {
			return nil, nil, err
		}

		if isPredefined {
			predefinedRules = append(predefinedRules, rule)
		} else {
			rules = append(rules, rule)
		}
	}
	return rules, predefinedRules, nil
}

const ComputeFirewallPolicyWithRulesAssetType string = "compute.googleapis.com/FirewallPolicyWithRules"

func ResourceConverterComputeFirewallPolicyWithRules() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeFirewallPolicyWithRulesAssetType,
		Convert:   GetComputeFirewallPolicyWithRulesCaiObject,
	}
}

func GetComputeFirewallPolicyWithRulesCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/locations/global/firewallPolicies/{{policy_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeFirewallPolicyWithRulesApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeFirewallPolicyWithRulesAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "FirewallPolicyWithRules",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeFirewallPolicyWithRulesApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	shortNameProp, err := expandComputeFirewallPolicyWithRulesShortName(d.Get("short_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("short_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(shortNameProp)) && (ok || !reflect.DeepEqual(v, shortNameProp)) {
		obj["shortName"] = shortNameProp
	}
	descriptionProp, err := expandComputeFirewallPolicyWithRulesDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	rulesProp, err := expandComputeFirewallPolicyWithRulesRule(d.Get("rule"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("rule"); !tpgresource.IsEmptyValue(reflect.ValueOf(rulesProp)) && (ok || !reflect.DeepEqual(v, rulesProp)) {
		obj["rules"] = rulesProp
	}
	fingerprintProp, err := expandComputeFirewallPolicyWithRulesFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("fingerprint"); !tpgresource.IsEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	parentProp, err := expandComputeFirewallPolicyWithRulesParent(d.Get("parent"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("parent"); !tpgresource.IsEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}

	return resourceComputeFirewallPolicyWithRulesEncoder(d, config, obj)
}

func resourceComputeFirewallPolicyWithRulesEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "rules") // Rules are not supported in the create API
	return obj, nil
}

func expandComputeFirewallPolicyWithRulesShortName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDescription, err := expandComputeFirewallPolicyWithRulesRuleDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedRuleName, err := expandComputeFirewallPolicyWithRulesRuleRuleName(original["rule_name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRuleName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ruleName"] = transformedRuleName
		}

		transformedPriority, err := expandComputeFirewallPolicyWithRulesRulePriority(original["priority"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPriority); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["priority"] = transformedPriority
		}

		transformedMatch, err := expandComputeFirewallPolicyWithRulesRuleMatch(original["match"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMatch); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["match"] = transformedMatch
		}

		transformedAction, err := expandComputeFirewallPolicyWithRulesRuleAction(original["action"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["action"] = transformedAction
		}

		transformedDirection, err := expandComputeFirewallPolicyWithRulesRuleDirection(original["direction"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDirection); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["direction"] = transformedDirection
		}

		transformedEnableLogging, err := expandComputeFirewallPolicyWithRulesRuleEnableLogging(original["enable_logging"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["enableLogging"] = transformedEnableLogging
		}

		transformedTargetServiceAccounts, err := expandComputeFirewallPolicyWithRulesRuleTargetServiceAccounts(original["target_service_accounts"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTargetServiceAccounts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["targetServiceAccounts"] = transformedTargetServiceAccounts
		}

		transformedSecurityProfileGroup, err := expandComputeFirewallPolicyWithRulesRuleSecurityProfileGroup(original["security_profile_group"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSecurityProfileGroup); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["securityProfileGroup"] = transformedSecurityProfileGroup
		}

		transformedTlsInspect, err := expandComputeFirewallPolicyWithRulesRuleTlsInspect(original["tls_inspect"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTlsInspect); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["tlsInspect"] = transformedTlsInspect
		}

		transformedTargetResources, err := expandComputeFirewallPolicyWithRulesRuleTargetResources(original["target_resources"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTargetResources); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["targetResources"] = transformedTargetResources
		}

		transformedDisabled, err := expandComputeFirewallPolicyWithRulesRuleDisabled(original["disabled"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDisabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["disabled"] = transformedDisabled
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeFirewallPolicyWithRulesRuleDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleRuleName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRulePriority(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleMatch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSrcIpRanges, err := expandComputeFirewallPolicyWithRulesRuleMatchSrcIpRanges(original["src_ip_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcIpRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcIpRanges"] = transformedSrcIpRanges
	}

	transformedDestIpRanges, err := expandComputeFirewallPolicyWithRulesRuleMatchDestIpRanges(original["dest_ip_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestIpRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destIpRanges"] = transformedDestIpRanges
	}

	transformedSrcAddressGroups, err := expandComputeFirewallPolicyWithRulesRuleMatchSrcAddressGroups(original["src_address_groups"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcAddressGroups); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcAddressGroups"] = transformedSrcAddressGroups
	}

	transformedDestAddressGroups, err := expandComputeFirewallPolicyWithRulesRuleMatchDestAddressGroups(original["dest_address_groups"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestAddressGroups); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destAddressGroups"] = transformedDestAddressGroups
	}

	transformedSrcFqdns, err := expandComputeFirewallPolicyWithRulesRuleMatchSrcFqdns(original["src_fqdns"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcFqdns); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcFqdns"] = transformedSrcFqdns
	}

	transformedDestFqdns, err := expandComputeFirewallPolicyWithRulesRuleMatchDestFqdns(original["dest_fqdns"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestFqdns); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destFqdns"] = transformedDestFqdns
	}

	transformedSrcNetworkScope, err := expandComputeFirewallPolicyWithRulesRuleMatchSrcNetworkScope(original["src_network_scope"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcNetworkScope); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcNetworkScope"] = transformedSrcNetworkScope
	}

	transformedSrcNetworks, err := expandComputeFirewallPolicyWithRulesRuleMatchSrcNetworks(original["src_networks"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcNetworks); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcNetworks"] = transformedSrcNetworks
	}

	transformedDestNetworkScope, err := expandComputeFirewallPolicyWithRulesRuleMatchDestNetworkScope(original["dest_network_scope"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestNetworkScope); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destNetworkScope"] = transformedDestNetworkScope
	}

	transformedSrcRegionCodes, err := expandComputeFirewallPolicyWithRulesRuleMatchSrcRegionCodes(original["src_region_codes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcRegionCodes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcRegionCodes"] = transformedSrcRegionCodes
	}

	transformedDestRegionCodes, err := expandComputeFirewallPolicyWithRulesRuleMatchDestRegionCodes(original["dest_region_codes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestRegionCodes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destRegionCodes"] = transformedDestRegionCodes
	}

	transformedSrcThreatIntelligences, err := expandComputeFirewallPolicyWithRulesRuleMatchSrcThreatIntelligences(original["src_threat_intelligences"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcThreatIntelligences); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcThreatIntelligences"] = transformedSrcThreatIntelligences
	}

	transformedDestThreatIntelligences, err := expandComputeFirewallPolicyWithRulesRuleMatchDestThreatIntelligences(original["dest_threat_intelligences"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestThreatIntelligences); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destThreatIntelligences"] = transformedDestThreatIntelligences
	}

	transformedLayer4Config, err := expandComputeFirewallPolicyWithRulesRuleMatchLayer4Config(original["layer4_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLayer4Config); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["layer4Configs"] = transformedLayer4Config
	}

	return transformed, nil
}

func expandComputeFirewallPolicyWithRulesRuleMatchSrcIpRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleMatchDestIpRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleMatchSrcAddressGroups(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleMatchDestAddressGroups(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleMatchSrcFqdns(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleMatchDestFqdns(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleMatchSrcNetworkScope(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleMatchSrcNetworks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleMatchDestNetworkScope(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleMatchSrcRegionCodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleMatchDestRegionCodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleMatchSrcThreatIntelligences(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleMatchDestThreatIntelligences(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleMatchLayer4Config(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedIpProtocol, err := expandComputeFirewallPolicyWithRulesRuleMatchLayer4ConfigIpProtocol(original["ip_protocol"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpProtocol); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipProtocol"] = transformedIpProtocol
		}

		transformedPorts, err := expandComputeFirewallPolicyWithRulesRuleMatchLayer4ConfigPorts(original["ports"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPorts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ports"] = transformedPorts
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeFirewallPolicyWithRulesRuleMatchLayer4ConfigIpProtocol(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleMatchLayer4ConfigPorts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleDirection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleEnableLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleTargetServiceAccounts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleSecurityProfileGroup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleTlsInspect(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleTargetResources(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesRuleDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesFingerprint(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeFirewallPolicyWithRulesParent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
