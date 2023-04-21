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

package google

import (
	"reflect"

	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const DataLossPreventionInspectTemplateAssetType string = "dlp.googleapis.com/InspectTemplate"

func resourceConverterDataLossPreventionInspectTemplate() ResourceConverter {
	return ResourceConverter{
		AssetType: DataLossPreventionInspectTemplateAssetType,
		Convert:   GetDataLossPreventionInspectTemplateCaiObject,
	}
}

func GetDataLossPreventionInspectTemplateCaiObject(d TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//dlp.googleapis.com/{{parent}}/inspectTemplates/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetDataLossPreventionInspectTemplateApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: DataLossPreventionInspectTemplateAssetType,
			Resource: &AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dlp/v2/rest",
				DiscoveryName:        "InspectTemplate",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetDataLossPreventionInspectTemplateApiObject(d TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandDataLossPreventionInspectTemplateDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandDataLossPreventionInspectTemplateDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	inspectConfigProp, err := expandDataLossPreventionInspectTemplateInspectConfig(d.Get("inspect_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("inspect_config"); !isEmptyValue(reflect.ValueOf(inspectConfigProp)) && (ok || !reflect.DeepEqual(v, inspectConfigProp)) {
		obj["inspectConfig"] = inspectConfigProp
	}

	return resourceDataLossPreventionInspectTemplateEncoder(d, config, obj)
}

func resourceDataLossPreventionInspectTemplateEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	newObj := make(map[string]interface{})
	newObj["inspectTemplate"] = obj
	return newObj, nil
}

func expandDataLossPreventionInspectTemplateDescription(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateDisplayName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfig(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExcludeInfoTypes, err := expandDataLossPreventionInspectTemplateInspectConfigExcludeInfoTypes(original["exclude_info_types"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExcludeInfoTypes); val.IsValid() && !isEmptyValue(val) {
		transformed["excludeInfoTypes"] = transformedExcludeInfoTypes
	}

	transformedIncludeQuote, err := expandDataLossPreventionInspectTemplateInspectConfigIncludeQuote(original["include_quote"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludeQuote); val.IsValid() && !isEmptyValue(val) {
		transformed["includeQuote"] = transformedIncludeQuote
	}

	transformedMinLikelihood, err := expandDataLossPreventionInspectTemplateInspectConfigMinLikelihood(original["min_likelihood"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinLikelihood); val.IsValid() && !isEmptyValue(val) {
		transformed["minLikelihood"] = transformedMinLikelihood
	}

	transformedLimits, err := expandDataLossPreventionInspectTemplateInspectConfigLimits(original["limits"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLimits); val.IsValid() && !isEmptyValue(val) {
		transformed["limits"] = transformedLimits
	}

	transformedInfoTypes, err := expandDataLossPreventionInspectTemplateInspectConfigInfoTypes(original["info_types"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInfoTypes); val.IsValid() && !isEmptyValue(val) {
		transformed["infoTypes"] = transformedInfoTypes
	}

	transformedContentOptions, err := expandDataLossPreventionInspectTemplateInspectConfigContentOptions(original["content_options"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedContentOptions); val.IsValid() && !isEmptyValue(val) {
		transformed["contentOptions"] = transformedContentOptions
	}

	transformedRuleSet, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSet(original["rule_set"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRuleSet); val.IsValid() && !isEmptyValue(val) {
		transformed["ruleSet"] = transformedRuleSet
	}

	transformedCustomInfoTypes, err := expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypes(original["custom_info_types"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCustomInfoTypes); val.IsValid() && !isEmptyValue(val) {
		transformed["customInfoTypes"] = transformedCustomInfoTypes
	}

	return transformed, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigExcludeInfoTypes(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigIncludeQuote(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigMinLikelihood(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigLimits(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMaxFindingsPerItem, err := expandDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerItem(original["max_findings_per_item"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxFindingsPerItem); val.IsValid() && !isEmptyValue(val) {
		transformed["maxFindingsPerItem"] = transformedMaxFindingsPerItem
	}

	transformedMaxFindingsPerRequest, err := expandDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerRequest(original["max_findings_per_request"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxFindingsPerRequest); val.IsValid() && !isEmptyValue(val) {
		transformed["maxFindingsPerRequest"] = transformedMaxFindingsPerRequest
	}

	transformedMaxFindingsPerInfoType, err := expandDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(original["max_findings_per_info_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxFindingsPerInfoType); val.IsValid() && !isEmptyValue(val) {
		transformed["maxFindingsPerInfoType"] = transformedMaxFindingsPerInfoType
	}

	return transformed, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerItem(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerRequest(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoType(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedInfoType, err := expandDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(original["info_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedInfoType); val.IsValid() && !isEmptyValue(val) {
			transformed["infoType"] = transformedInfoType
		}

		transformedMaxFindings, err := expandDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeMaxFindings(original["max_findings"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedMaxFindings); val.IsValid() && !isEmptyValue(val) {
			transformed["maxFindings"] = transformedMaxFindings
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoType(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
		transformed["name"] = transformedName
	}

	return transformed, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeInfoTypeName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigLimitsMaxFindingsPerInfoTypeMaxFindings(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigInfoTypes(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandDataLossPreventionInspectTemplateInspectConfigInfoTypesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedVersion, err := expandDataLossPreventionInspectTemplateInspectConfigInfoTypesVersion(original["version"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedVersion); val.IsValid() && !isEmptyValue(val) {
			transformed["version"] = transformedVersion
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigInfoTypesName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigInfoTypesVersion(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigContentOptions(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSet(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedInfoTypes, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetInfoTypes(original["info_types"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedInfoTypes); val.IsValid() && !isEmptyValue(val) {
			transformed["infoTypes"] = transformedInfoTypes
		}

		transformedRules, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRules(original["rules"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRules); val.IsValid() && !isEmptyValue(val) {
			transformed["rules"] = transformedRules
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetInfoTypes(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetInfoTypesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetInfoTypesName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRules(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedHotwordRule, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRule(original["hotword_rule"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedHotwordRule); val.IsValid() && !isEmptyValue(val) {
			transformed["hotwordRule"] = transformedHotwordRule
		}

		transformedExclusionRule, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRule(original["exclusion_rule"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedExclusionRule); val.IsValid() && !isEmptyValue(val) {
			transformed["exclusionRule"] = transformedExclusionRule
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRule(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHotwordRegex, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(original["hotword_regex"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHotwordRegex); val.IsValid() && !isEmptyValue(val) {
		transformed["hotwordRegex"] = transformedHotwordRegex
	}

	transformedProximity, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(original["proximity"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProximity); val.IsValid() && !isEmptyValue(val) {
		transformed["proximity"] = transformedProximity
	}

	transformedLikelihoodAdjustment, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(original["likelihood_adjustment"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLikelihoodAdjustment); val.IsValid() && !isEmptyValue(val) {
		transformed["likelihoodAdjustment"] = transformedLikelihoodAdjustment
	}

	return transformed, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegex(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPattern, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexPattern(original["pattern"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPattern); val.IsValid() && !isEmptyValue(val) {
		transformed["pattern"] = transformedPattern
	}

	transformedGroupIndexes, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexGroupIndexes(original["group_indexes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGroupIndexes); val.IsValid() && !isEmptyValue(val) {
		transformed["groupIndexes"] = transformedGroupIndexes
	}

	return transformed, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexPattern(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleHotwordRegexGroupIndexes(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximity(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedWindowBefore, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityWindowBefore(original["window_before"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWindowBefore); val.IsValid() && !isEmptyValue(val) {
		transformed["windowBefore"] = transformedWindowBefore
	}

	transformedWindowAfter, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityWindowAfter(original["window_after"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWindowAfter); val.IsValid() && !isEmptyValue(val) {
		transformed["windowAfter"] = transformedWindowAfter
	}

	return transformed, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityWindowBefore(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleProximityWindowAfter(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustment(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFixedLikelihood, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihood(original["fixed_likelihood"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFixedLikelihood); val.IsValid() && !isEmptyValue(val) {
		transformed["fixedLikelihood"] = transformedFixedLikelihood
	}

	transformedRelativeLikelihood, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentRelativeLikelihood(original["relative_likelihood"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRelativeLikelihood); val.IsValid() && !isEmptyValue(val) {
		transformed["relativeLikelihood"] = transformedRelativeLikelihood
	}

	return transformed, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentFixedLikelihood(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesHotwordRuleLikelihoodAdjustmentRelativeLikelihood(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRule(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMatchingType, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingType(original["matching_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMatchingType); val.IsValid() && !isEmptyValue(val) {
		transformed["matchingType"] = transformedMatchingType
	}

	transformedDictionary, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(original["dictionary"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDictionary); val.IsValid() && !isEmptyValue(val) {
		transformed["dictionary"] = transformedDictionary
	}

	transformedRegex, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(original["regex"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRegex); val.IsValid() && !isEmptyValue(val) {
		transformed["regex"] = transformedRegex
	}

	transformedExcludeInfoTypes, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(original["exclude_info_types"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExcludeInfoTypes); val.IsValid() && !isEmptyValue(val) {
		transformed["excludeInfoTypes"] = transformedExcludeInfoTypes
	}

	return transformed, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleMatchingType(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionary(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedWordList, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(original["word_list"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWordList); val.IsValid() && !isEmptyValue(val) {
		transformed["wordList"] = transformedWordList
	}

	transformedCloudStoragePath, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(original["cloud_storage_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCloudStoragePath); val.IsValid() && !isEmptyValue(val) {
		transformed["cloudStoragePath"] = transformedCloudStoragePath
	}

	return transformed, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordList(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedWords, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListWords(original["words"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWords); val.IsValid() && !isEmptyValue(val) {
		transformed["words"] = transformedWords
	}

	return transformed, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryWordListWords(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePath(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPath, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathPath(original["path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !isEmptyValue(val) {
		transformed["path"] = transformedPath
	}

	return transformed, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleDictionaryCloudStoragePathPath(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegex(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPattern, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexPattern(original["pattern"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPattern); val.IsValid() && !isEmptyValue(val) {
		transformed["pattern"] = transformedPattern
	}

	transformedGroupIndexes, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexGroupIndexes(original["group_indexes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGroupIndexes); val.IsValid() && !isEmptyValue(val) {
		transformed["groupIndexes"] = transformedGroupIndexes
	}

	return transformed, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexPattern(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleRegexGroupIndexes(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypes(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedInfoTypes, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(original["info_types"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInfoTypes); val.IsValid() && !isEmptyValue(val) {
		transformed["infoTypes"] = transformedInfoTypes
	}

	return transformed, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypes(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigRuleSetRulesExclusionRuleExcludeInfoTypesInfoTypesName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypes(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedInfoType, err := expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoType(original["info_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedInfoType); val.IsValid() && !isEmptyValue(val) {
			transformed["infoType"] = transformedInfoType
		}

		transformedLikelihood, err := expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesLikelihood(original["likelihood"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedLikelihood); val.IsValid() && !isEmptyValue(val) {
			transformed["likelihood"] = transformedLikelihood
		}

		transformedExclusionType, err := expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesExclusionType(original["exclusion_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedExclusionType); val.IsValid() && !isEmptyValue(val) {
			transformed["exclusionType"] = transformedExclusionType
		}

		transformedRegex, err := expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegex(original["regex"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRegex); val.IsValid() && !isEmptyValue(val) {
			transformed["regex"] = transformedRegex
		}

		transformedDictionary, err := expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionary(original["dictionary"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDictionary); val.IsValid() && !isEmptyValue(val) {
			transformed["dictionary"] = transformedDictionary
		}

		transformedStoredType, err := expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredType(original["stored_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedStoredType); val.IsValid() && !isEmptyValue(val) {
			transformed["storedType"] = transformedStoredType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoType(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoTypeName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
		transformed["name"] = transformedName
	}

	return transformed, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoTypeName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesLikelihood(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesExclusionType(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegex(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPattern, err := expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegexPattern(original["pattern"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPattern); val.IsValid() && !isEmptyValue(val) {
		transformed["pattern"] = transformedPattern
	}

	transformedGroupIndexes, err := expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegexGroupIndexes(original["group_indexes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGroupIndexes); val.IsValid() && !isEmptyValue(val) {
		transformed["groupIndexes"] = transformedGroupIndexes
	}

	return transformed, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegexPattern(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegexGroupIndexes(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionary(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedWordList, err := expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(original["word_list"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWordList); val.IsValid() && !isEmptyValue(val) {
		transformed["wordList"] = transformedWordList
	}

	transformedCloudStoragePath, err := expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(original["cloud_storage_path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCloudStoragePath); val.IsValid() && !isEmptyValue(val) {
		transformed["cloudStoragePath"] = transformedCloudStoragePath
	}

	return transformed, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionaryWordList(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedWords, err := expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListWords(original["words"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWords); val.IsValid() && !isEmptyValue(val) {
		transformed["words"] = transformedWords
	}

	return transformed, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionaryWordListWords(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePath(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPath, err := expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathPath(original["path"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPath); val.IsValid() && !isEmptyValue(val) {
		transformed["path"] = transformedPath
	}

	return transformed, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionaryCloudStoragePathPath(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredType(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedName, err := expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredTypeName(original["name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
		transformed["name"] = transformedName
	}

	return transformed, nil
}

func expandDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredTypeName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
