// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/securitycenter/FolderCustomModule.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package securitycenter

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const SecurityCenterFolderCustomModuleAssetType string = "securitycenter.googleapis.com/FolderCustomModule"

func ResourceConverterSecurityCenterFolderCustomModule() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: SecurityCenterFolderCustomModuleAssetType,
		Convert:   GetSecurityCenterFolderCustomModuleCaiObject,
	}
}

func GetSecurityCenterFolderCustomModuleCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//securitycenter.googleapis.com/folders/{{folder}}/securityHealthAnalyticsSettings/customModules/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetSecurityCenterFolderCustomModuleApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: SecurityCenterFolderCustomModuleAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/securitycenter/v1/rest",
				DiscoveryName:        "FolderCustomModule",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetSecurityCenterFolderCustomModuleApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandSecurityCenterFolderCustomModuleDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	enablementStateProp, err := expandSecurityCenterFolderCustomModuleEnablementState(d.Get("enablement_state"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enablement_state"); !tpgresource.IsEmptyValue(reflect.ValueOf(enablementStateProp)) && (ok || !reflect.DeepEqual(v, enablementStateProp)) {
		obj["enablementState"] = enablementStateProp
	}
	customConfigProp, err := expandSecurityCenterFolderCustomModuleCustomConfig(d.Get("custom_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("custom_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(customConfigProp)) && (ok || !reflect.DeepEqual(v, customConfigProp)) {
		obj["customConfig"] = customConfigProp
	}

	return obj, nil
}

func expandSecurityCenterFolderCustomModuleDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleEnablementState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPredicate, err := expandSecurityCenterFolderCustomModuleCustomConfigPredicate(original["predicate"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPredicate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["predicate"] = transformedPredicate
	}

	transformedCustomOutput, err := expandSecurityCenterFolderCustomModuleCustomConfigCustomOutput(original["custom_output"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCustomOutput); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["customOutput"] = transformedCustomOutput
	}

	transformedResourceSelector, err := expandSecurityCenterFolderCustomModuleCustomConfigResourceSelector(original["resource_selector"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceSelector); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["resourceSelector"] = transformedResourceSelector
	}

	transformedSeverity, err := expandSecurityCenterFolderCustomModuleCustomConfigSeverity(original["severity"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSeverity); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["severity"] = transformedSeverity
	}

	transformedDescription, err := expandSecurityCenterFolderCustomModuleCustomConfigDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedRecommendation, err := expandSecurityCenterFolderCustomModuleCustomConfigRecommendation(original["recommendation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRecommendation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["recommendation"] = transformedRecommendation
	}

	return transformed, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigPredicate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandSecurityCenterFolderCustomModuleCustomConfigPredicateExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandSecurityCenterFolderCustomModuleCustomConfigPredicateTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandSecurityCenterFolderCustomModuleCustomConfigPredicateDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandSecurityCenterFolderCustomModuleCustomConfigPredicateLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigPredicateExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigPredicateTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigPredicateDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigPredicateLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigCustomOutput(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProperties, err := expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputProperties(original["properties"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProperties); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["properties"] = transformedProperties
	}

	return transformed, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputProperties(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedValueExpression, err := expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpression(original["value_expression"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValueExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["valueExpression"] = transformedValueExpression
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigCustomOutputPropertiesValueExpressionLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigResourceSelector(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResourceTypes, err := expandSecurityCenterFolderCustomModuleCustomConfigResourceSelectorResourceTypes(original["resource_types"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceTypes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["resourceTypes"] = transformedResourceTypes
	}

	return transformed, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigResourceSelectorResourceTypes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigSeverity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterFolderCustomModuleCustomConfigRecommendation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
