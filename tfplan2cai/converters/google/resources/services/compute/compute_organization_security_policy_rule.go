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
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeOrganizationSecurityPolicyRuleAssetType string = "compute.googleapis.com/OrganizationSecurityPolicyRule"

func ResourceConverterComputeOrganizationSecurityPolicyRule() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeOrganizationSecurityPolicyRuleAssetType,
		Convert:   GetComputeOrganizationSecurityPolicyRuleCaiObject,
	}
}

func GetComputeOrganizationSecurityPolicyRuleCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/{{policy_id}}/getRule?priority={{priority}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeOrganizationSecurityPolicyRuleApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeOrganizationSecurityPolicyRuleAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "OrganizationSecurityPolicyRule",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeOrganizationSecurityPolicyRuleApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeOrganizationSecurityPolicyRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	priorityProp, err := expandComputeOrganizationSecurityPolicyRulePriority(d.Get("priority"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("priority"); !tpgresource.IsEmptyValue(reflect.ValueOf(priorityProp)) && (ok || !reflect.DeepEqual(v, priorityProp)) {
		obj["priority"] = priorityProp
	}
	matchProp, err := expandComputeOrganizationSecurityPolicyRuleMatch(d.Get("match"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("match"); !tpgresource.IsEmptyValue(reflect.ValueOf(matchProp)) && (ok || !reflect.DeepEqual(v, matchProp)) {
		obj["match"] = matchProp
	}
	actionProp, err := expandComputeOrganizationSecurityPolicyRuleAction(d.Get("action"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("action"); !tpgresource.IsEmptyValue(reflect.ValueOf(actionProp)) && (ok || !reflect.DeepEqual(v, actionProp)) {
		obj["action"] = actionProp
	}
	previewProp, err := expandComputeOrganizationSecurityPolicyRulePreview(d.Get("preview"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("preview"); !tpgresource.IsEmptyValue(reflect.ValueOf(previewProp)) && (ok || !reflect.DeepEqual(v, previewProp)) {
		obj["preview"] = previewProp
	}
	directionProp, err := expandComputeOrganizationSecurityPolicyRuleDirection(d.Get("direction"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("direction"); !tpgresource.IsEmptyValue(reflect.ValueOf(directionProp)) && (ok || !reflect.DeepEqual(v, directionProp)) {
		obj["direction"] = directionProp
	}
	targetResourcesProp, err := expandComputeOrganizationSecurityPolicyRuleTargetResources(d.Get("target_resources"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target_resources"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetResourcesProp)) && (ok || !reflect.DeepEqual(v, targetResourcesProp)) {
		obj["targetResources"] = targetResourcesProp
	}
	enableLoggingProp, err := expandComputeOrganizationSecurityPolicyRuleEnableLogging(d.Get("enable_logging"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_logging"); ok || !reflect.DeepEqual(v, enableLoggingProp) {
		obj["enableLogging"] = enableLoggingProp
	}
	targetServiceAccountsProp, err := expandComputeOrganizationSecurityPolicyRuleTargetServiceAccounts(d.Get("target_service_accounts"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target_service_accounts"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetServiceAccountsProp)) && (ok || !reflect.DeepEqual(v, targetServiceAccountsProp)) {
		obj["targetServiceAccounts"] = targetServiceAccountsProp
	}

	return obj, nil
}

func expandComputeOrganizationSecurityPolicyRuleDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRulePriority(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleMatch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDescription, err := expandComputeOrganizationSecurityPolicyRuleMatchDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedVersionedExpr, err := expandComputeOrganizationSecurityPolicyRuleMatchVersionedExpr(original["versioned_expr"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVersionedExpr); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["versionedExpr"] = transformedVersionedExpr
	}

	transformedConfig, err := expandComputeOrganizationSecurityPolicyRuleMatchConfig(original["config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["config"] = transformedConfig
	}

	return transformed, nil
}

func expandComputeOrganizationSecurityPolicyRuleMatchDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleMatchVersionedExpr(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleMatchConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSrcIpRanges, err := expandComputeOrganizationSecurityPolicyRuleMatchConfigSrcIpRanges(original["src_ip_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSrcIpRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["srcIpRanges"] = transformedSrcIpRanges
	}

	transformedDestIpRanges, err := expandComputeOrganizationSecurityPolicyRuleMatchConfigDestIpRanges(original["dest_ip_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDestIpRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["destIpRanges"] = transformedDestIpRanges
	}

	transformedLayer4Config, err := expandComputeOrganizationSecurityPolicyRuleMatchConfigLayer4Config(original["layer4_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLayer4Config); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["layer4Configs"] = transformedLayer4Config
	}

	return transformed, nil
}

func expandComputeOrganizationSecurityPolicyRuleMatchConfigSrcIpRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleMatchConfigDestIpRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleMatchConfigLayer4Config(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedIpProtocol, err := expandComputeOrganizationSecurityPolicyRuleMatchConfigLayer4ConfigIpProtocol(original["ip_protocol"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpProtocol); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipProtocol"] = transformedIpProtocol
		}

		transformedPorts, err := expandComputeOrganizationSecurityPolicyRuleMatchConfigLayer4ConfigPorts(original["ports"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedPorts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ports"] = transformedPorts
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeOrganizationSecurityPolicyRuleMatchConfigLayer4ConfigIpProtocol(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleMatchConfigLayer4ConfigPorts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRulePreview(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleDirection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleTargetResources(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleEnableLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeOrganizationSecurityPolicyRuleTargetServiceAccounts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
