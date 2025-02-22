// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/iam3/PrincipalAccessBoundaryPolicy.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package iam3

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const IAM3PrincipalAccessBoundaryPolicyAssetType string = "iam.googleapis.com/PrincipalAccessBoundaryPolicy"

func ResourceConverterIAM3PrincipalAccessBoundaryPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: IAM3PrincipalAccessBoundaryPolicyAssetType,
		Convert:   GetIAM3PrincipalAccessBoundaryPolicyCaiObject,
	}
}

func GetIAM3PrincipalAccessBoundaryPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//iam.googleapis.com/organizations/{{organization}}/locations/{{location}}/principalAccessBoundaryPolicies/{{principal_access_boundary_policy_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetIAM3PrincipalAccessBoundaryPolicyApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: IAM3PrincipalAccessBoundaryPolicyAssetType,
			Resource: &cai.AssetResource{
				Version:              "v3beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/iam/v3beta/rest",
				DiscoveryName:        "PrincipalAccessBoundaryPolicy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetIAM3PrincipalAccessBoundaryPolicyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandIAM3PrincipalAccessBoundaryPolicyDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	detailsProp, err := expandIAM3PrincipalAccessBoundaryPolicyDetails(d.Get("details"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("details"); !tpgresource.IsEmptyValue(reflect.ValueOf(detailsProp)) && (ok || !reflect.DeepEqual(v, detailsProp)) {
		obj["details"] = detailsProp
	}
	annotationsProp, err := expandIAM3PrincipalAccessBoundaryPolicyEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	return obj, nil
}

func expandIAM3PrincipalAccessBoundaryPolicyDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3PrincipalAccessBoundaryPolicyDetails(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRules, err := expandIAM3PrincipalAccessBoundaryPolicyDetailsRules(original["rules"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRules); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rules"] = transformedRules
	}

	transformedEnforcementVersion, err := expandIAM3PrincipalAccessBoundaryPolicyDetailsEnforcementVersion(original["enforcement_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnforcementVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enforcementVersion"] = transformedEnforcementVersion
	}

	return transformed, nil
}

func expandIAM3PrincipalAccessBoundaryPolicyDetailsRules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDescription, err := expandIAM3PrincipalAccessBoundaryPolicyDetailsRulesDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedResources, err := expandIAM3PrincipalAccessBoundaryPolicyDetailsRulesResources(original["resources"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedResources); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["resources"] = transformedResources
		}

		transformedEffect, err := expandIAM3PrincipalAccessBoundaryPolicyDetailsRulesEffect(original["effect"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEffect); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["effect"] = transformedEffect
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandIAM3PrincipalAccessBoundaryPolicyDetailsRulesDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3PrincipalAccessBoundaryPolicyDetailsRulesResources(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3PrincipalAccessBoundaryPolicyDetailsRulesEffect(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3PrincipalAccessBoundaryPolicyDetailsEnforcementVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3PrincipalAccessBoundaryPolicyEffectiveAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
