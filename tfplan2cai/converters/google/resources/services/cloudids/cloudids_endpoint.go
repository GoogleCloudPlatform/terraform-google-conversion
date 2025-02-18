// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/cloudids/Endpoint.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package cloudids

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const CloudIdsEndpointAssetType string = "ids.googleapis.com/Endpoint"

func ResourceConverterCloudIdsEndpoint() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: CloudIdsEndpointAssetType,
		Convert:   GetCloudIdsEndpointCaiObject,
	}
}

func GetCloudIdsEndpointCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//ids.googleapis.com/projects/{{project}}/locations/{{location}}/endpoints/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetCloudIdsEndpointApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: CloudIdsEndpointAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/ids/v1/rest",
				DiscoveryName:        "Endpoint",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetCloudIdsEndpointApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandCloudIdsEndpointName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandCloudIdsEndpointNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	descriptionProp, err := expandCloudIdsEndpointDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	severityProp, err := expandCloudIdsEndpointSeverity(d.Get("severity"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("severity"); !tpgresource.IsEmptyValue(reflect.ValueOf(severityProp)) && (ok || !reflect.DeepEqual(v, severityProp)) {
		obj["severity"] = severityProp
	}
	threatExceptionsProp, err := expandCloudIdsEndpointThreatExceptions(d.Get("threat_exceptions"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("threat_exceptions"); !tpgresource.IsEmptyValue(reflect.ValueOf(threatExceptionsProp)) && (ok || !reflect.DeepEqual(v, threatExceptionsProp)) {
		obj["threatExceptions"] = threatExceptionsProp
	}

	return obj, nil
}

func expandCloudIdsEndpointName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/endpoints/{{name}}")
}

func expandCloudIdsEndpointNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdsEndpointDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdsEndpointSeverity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandCloudIdsEndpointThreatExceptions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
