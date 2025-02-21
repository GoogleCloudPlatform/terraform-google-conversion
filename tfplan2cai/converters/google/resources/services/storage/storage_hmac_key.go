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

package storage

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const StorageHmacKeyAssetType string = "storage.googleapis.com/HmacKey"

func ResourceConverterStorageHmacKey() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: StorageHmacKeyAssetType,
		Convert:   GetStorageHmacKeyCaiObject,
	}
}

func GetStorageHmacKeyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//storage.googleapis.com/projects/{{project}}/hmacKeys/{{access_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetStorageHmacKeyApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: StorageHmacKeyAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/storage/v1/rest",
				DiscoveryName:        "HmacKey",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetStorageHmacKeyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	serviceAccountEmailProp, err := expandStorageHmacKeyServiceAccountEmail(d.Get("service_account_email"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_account_email"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceAccountEmailProp)) && (ok || !reflect.DeepEqual(v, serviceAccountEmailProp)) {
		obj["serviceAccountEmail"] = serviceAccountEmailProp
	}
	stateProp, err := expandStorageHmacKeyState(d.Get("state"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("state"); !tpgresource.IsEmptyValue(reflect.ValueOf(stateProp)) && (ok || !reflect.DeepEqual(v, stateProp)) {
		obj["state"] = stateProp
	}

	return obj, nil
}

func expandStorageHmacKeyServiceAccountEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageHmacKeyState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
