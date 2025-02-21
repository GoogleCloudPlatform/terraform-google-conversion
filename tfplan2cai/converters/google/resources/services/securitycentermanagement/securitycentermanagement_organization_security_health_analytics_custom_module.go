// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/securitycentermanagement/OrganizationSecurityHealthAnalyticsCustomModule.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package securitycentermanagement

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const SecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleAssetType string = "securitycentermanagement.googleapis.com/OrganizationSecurityHealthAnalyticsCustomModule"

func ResourceConverterSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModule() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: SecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleAssetType,
		Convert:   GetSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCaiObject,
	}
}

func GetSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//securitycentermanagement.googleapis.com/organizations/{{organization}}/locations/{{location}}/securityHealthAnalyticsCustomModules/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: SecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/securitycentermanagement/v1/rest",
				DiscoveryName:        "OrganizationSecurityHealthAnalyticsCustomModule",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	enablementStateProp, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleEnablementState(d.Get("enablement_state"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enablement_state"); !tpgresource.IsEmptyValue(reflect.ValueOf(enablementStateProp)) && (ok || !reflect.DeepEqual(v, enablementStateProp)) {
		obj["enablementState"] = enablementStateProp
	}
	customConfigProp, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfig(d.Get("custom_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("custom_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(customConfigProp)) && (ok || !reflect.DeepEqual(v, customConfigProp)) {
		obj["customConfig"] = customConfigProp
	}

	return obj, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleEnablementState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPredicate, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigPredicate(original["predicate"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPredicate); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["predicate"] = transformedPredicate
	}

	transformedCustomOutput, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput(original["custom_output"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCustomOutput); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["customOutput"] = transformedCustomOutput
	}

	transformedResourceSelector, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector(original["resource_selector"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceSelector); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["resourceSelector"] = transformedResourceSelector
	}

	transformedSeverity, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigSeverity(original["severity"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSeverity); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["severity"] = transformedSeverity
	}

	transformedDescription, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedRecommendation, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigRecommendation(original["recommendation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRecommendation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["recommendation"] = transformedRecommendation
	}

	return transformed, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigPredicate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigPredicateExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigPredicateTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigPredicateDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigPredicateLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigPredicateExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigPredicateTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigPredicateDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigPredicateLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProperties, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputProperties(original["properties"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProperties); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["properties"] = transformedProperties
	}

	return transformed, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputProperties(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputPropertiesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedValueExpression, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputPropertiesValueExpression(original["value_expression"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValueExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["valueExpression"] = transformedValueExpression
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputPropertiesName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputPropertiesValueExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedExpression, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputPropertiesValueExpressionExpression(original["expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expression"] = transformedExpression
	}

	transformedTitle, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputPropertiesValueExpressionTitle(original["title"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTitle); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["title"] = transformedTitle
	}

	transformedDescription, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputPropertiesValueExpressionDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedLocation, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputPropertiesValueExpressionLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputPropertiesValueExpressionExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputPropertiesValueExpressionTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputPropertiesValueExpressionDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputPropertiesValueExpressionLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResourceTypes, err := expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelectorResourceTypes(original["resource_types"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResourceTypes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["resourceTypes"] = transformedResourceTypes
	}

	return transformed, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelectorResourceTypes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigSeverity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigRecommendation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
