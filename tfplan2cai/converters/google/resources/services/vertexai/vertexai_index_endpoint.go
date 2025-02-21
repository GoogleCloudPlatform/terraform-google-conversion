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

package vertexai

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const VertexAIIndexEndpointAssetType string = "aiplatform.googleapis.com/IndexEndpoint"

func ResourceConverterVertexAIIndexEndpoint() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: VertexAIIndexEndpointAssetType,
		Convert:   GetVertexAIIndexEndpointCaiObject,
	}
}

func GetVertexAIIndexEndpointCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//aiplatform.googleapis.com/projects/{{project}}/locations/{{region}}/indexEndpoints/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetVertexAIIndexEndpointApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: VertexAIIndexEndpointAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/aiplatform/v1beta1/rest",
				DiscoveryName:        "IndexEndpoint",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetVertexAIIndexEndpointApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandVertexAIIndexEndpointDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandVertexAIIndexEndpointDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	networkProp, err := expandVertexAIIndexEndpointNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	privateServiceConnectConfigProp, err := expandVertexAIIndexEndpointPrivateServiceConnectConfig(d.Get("private_service_connect_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("private_service_connect_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(privateServiceConnectConfigProp)) && (ok || !reflect.DeepEqual(v, privateServiceConnectConfigProp)) {
		obj["privateServiceConnectConfig"] = privateServiceConnectConfigProp
	}
	publicEndpointEnabledProp, err := expandVertexAIIndexEndpointPublicEndpointEnabled(d.Get("public_endpoint_enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("public_endpoint_enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(publicEndpointEnabledProp)) && (ok || !reflect.DeepEqual(v, publicEndpointEnabledProp)) {
		obj["publicEndpointEnabled"] = publicEndpointEnabledProp
	}
	labelsProp, err := expandVertexAIIndexEndpointEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandVertexAIIndexEndpointDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointPrivateServiceConnectConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnablePrivateServiceConnect, err := expandVertexAIIndexEndpointPrivateServiceConnectConfigEnablePrivateServiceConnect(original["enable_private_service_connect"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnablePrivateServiceConnect); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enablePrivateServiceConnect"] = transformedEnablePrivateServiceConnect
	}

	transformedProjectAllowlist, err := expandVertexAIIndexEndpointPrivateServiceConnectConfigProjectAllowlist(original["project_allowlist"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectAllowlist); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectAllowlist"] = transformedProjectAllowlist
	}

	return transformed, nil
}

func expandVertexAIIndexEndpointPrivateServiceConnectConfigEnablePrivateServiceConnect(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointPrivateServiceConnectConfigProjectAllowlist(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointPublicEndpointEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVertexAIIndexEndpointEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
