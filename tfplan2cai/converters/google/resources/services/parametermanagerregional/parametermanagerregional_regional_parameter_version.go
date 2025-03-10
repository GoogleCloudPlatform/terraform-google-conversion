// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/parametermanagerregional/RegionalParameterVersion.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package parametermanagerregional

import (
	"encoding/base64"
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ParameterManagerRegionalRegionalParameterVersionAssetType string = "parametermanager.googleapis.com/RegionalParameterVersion"

func ResourceConverterParameterManagerRegionalRegionalParameterVersion() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ParameterManagerRegionalRegionalParameterVersionAssetType,
		Convert:   GetParameterManagerRegionalRegionalParameterVersionCaiObject,
	}
}

func GetParameterManagerRegionalRegionalParameterVersionCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//parametermanager.googleapis.com/{{parameter}}/versions/{{parameter_version_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetParameterManagerRegionalRegionalParameterVersionApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ParameterManagerRegionalRegionalParameterVersionAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/parametermanager/v1/rest",
				DiscoveryName:        "RegionalParameterVersion",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetParameterManagerRegionalRegionalParameterVersionApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	disabledProp, err := expandParameterManagerRegionalRegionalParameterVersionDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	payloadProp, err := expandParameterManagerRegionalRegionalParameterVersionPayload(nil, d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("payload"); !tpgresource.IsEmptyValue(reflect.ValueOf(payloadProp)) && (ok || !reflect.DeepEqual(v, payloadProp)) {
		obj["payload"] = payloadProp
	}

	return obj, nil
}

func expandParameterManagerRegionalRegionalParameterVersionDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandParameterManagerRegionalRegionalParameterVersionPayload(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedParameterData, err := expandParameterManagerRegionalRegionalParameterVersionPayloadParameterData(d.Get("parameter_data"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedParameterData); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["data"] = transformedParameterData
	}

	return transformed, nil
}

func expandParameterManagerRegionalRegionalParameterVersionPayloadParameterData(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	return base64.StdEncoding.EncodeToString([]byte(v.(string))), nil
}
