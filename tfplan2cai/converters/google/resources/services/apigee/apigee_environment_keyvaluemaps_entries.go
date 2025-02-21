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

package apigee

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ApigeeEnvironmentKeyvaluemapsEntriesAssetType string = "apigee.googleapis.com/EnvironmentKeyvaluemapsEntries"

func ResourceConverterApigeeEnvironmentKeyvaluemapsEntries() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ApigeeEnvironmentKeyvaluemapsEntriesAssetType,
		Convert:   GetApigeeEnvironmentKeyvaluemapsEntriesCaiObject,
	}
}

func GetApigeeEnvironmentKeyvaluemapsEntriesCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//apigee.googleapis.com/{{env_keyvaluemap_id}}/entries/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetApigeeEnvironmentKeyvaluemapsEntriesApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ApigeeEnvironmentKeyvaluemapsEntriesAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/apigee/v1/rest",
				DiscoveryName:        "EnvironmentKeyvaluemapsEntries",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetApigeeEnvironmentKeyvaluemapsEntriesApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandApigeeEnvironmentKeyvaluemapsEntriesName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	valueProp, err := expandApigeeEnvironmentKeyvaluemapsEntriesValue(d.Get("value"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("value"); !tpgresource.IsEmptyValue(reflect.ValueOf(valueProp)) && (ok || !reflect.DeepEqual(v, valueProp)) {
		obj["value"] = valueProp
	}

	return obj, nil
}

func expandApigeeEnvironmentKeyvaluemapsEntriesName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentKeyvaluemapsEntriesValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
