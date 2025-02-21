// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networksecurity/SecurityProfile.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package networksecurity

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetworkSecuritySecurityProfileAssetType string = "networksecurity.googleapis.com/SecurityProfile"

func ResourceConverterNetworkSecuritySecurityProfile() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetworkSecuritySecurityProfileAssetType,
		Convert:   GetNetworkSecuritySecurityProfileCaiObject,
	}
}

func GetNetworkSecuritySecurityProfileCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//networksecurity.googleapis.com/{{parent}}/locations/{{location}}/securityProfiles/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetworkSecuritySecurityProfileApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetworkSecuritySecurityProfileAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networksecurity/v1beta1/rest",
				DiscoveryName:        "SecurityProfile",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetworkSecuritySecurityProfileApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecuritySecurityProfileDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	threatPreventionProfileProp, err := expandNetworkSecuritySecurityProfileThreatPreventionProfile(d.Get("threat_prevention_profile"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("threat_prevention_profile"); !tpgresource.IsEmptyValue(reflect.ValueOf(threatPreventionProfileProp)) && (ok || !reflect.DeepEqual(v, threatPreventionProfileProp)) {
		obj["threatPreventionProfile"] = threatPreventionProfileProp
	}
	customMirroringProfileProp, err := expandNetworkSecuritySecurityProfileCustomMirroringProfile(d.Get("custom_mirroring_profile"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("custom_mirroring_profile"); !tpgresource.IsEmptyValue(reflect.ValueOf(customMirroringProfileProp)) && (ok || !reflect.DeepEqual(v, customMirroringProfileProp)) {
		obj["customMirroringProfile"] = customMirroringProfileProp
	}
	customInterceptProfileProp, err := expandNetworkSecuritySecurityProfileCustomInterceptProfile(d.Get("custom_intercept_profile"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("custom_intercept_profile"); !tpgresource.IsEmptyValue(reflect.ValueOf(customInterceptProfileProp)) && (ok || !reflect.DeepEqual(v, customInterceptProfileProp)) {
		obj["customInterceptProfile"] = customInterceptProfileProp
	}
	typeProp, err := expandNetworkSecuritySecurityProfileType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	labelsProp, err := expandNetworkSecuritySecurityProfileEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetworkSecuritySecurityProfileDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileThreatPreventionProfile(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSeverityOverrides, err := expandNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverrides(original["severity_overrides"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSeverityOverrides); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["severityOverrides"] = transformedSeverityOverrides
	}

	transformedThreatOverrides, err := expandNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverrides(original["threat_overrides"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedThreatOverrides); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["threatOverrides"] = transformedThreatOverrides
	}

	return transformed, nil
}

func expandNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverrides(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedAction, err := expandNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverridesAction(original["action"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["action"] = transformedAction
		}

		transformedSeverity, err := expandNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverridesSeverity(original["severity"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSeverity); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["severity"] = transformedSeverity
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverridesAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileThreatPreventionProfileSeverityOverridesSeverity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverrides(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedAction, err := expandNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesAction(original["action"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["action"] = transformedAction
		}

		transformedThreatId, err := expandNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesThreatId(original["threat_id"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedThreatId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["threatId"] = transformedThreatId
		}

		transformedType, err := expandNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["type"] = transformedType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesAction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesThreatId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileThreatPreventionProfileThreatOverridesType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileCustomMirroringProfile(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMirroringEndpointGroup, err := expandNetworkSecuritySecurityProfileCustomMirroringProfileMirroringEndpointGroup(original["mirroring_endpoint_group"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMirroringEndpointGroup); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mirroringEndpointGroup"] = transformedMirroringEndpointGroup
	}

	return transformed, nil
}

func expandNetworkSecuritySecurityProfileCustomMirroringProfileMirroringEndpointGroup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileCustomInterceptProfile(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedInterceptEndpointGroup, err := expandNetworkSecuritySecurityProfileCustomInterceptProfileInterceptEndpointGroup(original["intercept_endpoint_group"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInterceptEndpointGroup); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["interceptEndpointGroup"] = transformedInterceptEndpointGroup
	}

	return transformed, nil
}

func expandNetworkSecuritySecurityProfileCustomInterceptProfileInterceptEndpointGroup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecuritySecurityProfileEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
