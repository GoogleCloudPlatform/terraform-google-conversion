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

package kms

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func folderPrefixSuppress(_, old, new string, d *schema.ResourceData) bool {
	prefix := "folders/"
	return prefix+old == new || prefix+new == old
}

const KMSAutokeyConfigAssetType string = "cloudkms.googleapis.com/AutokeyConfig"

func ResourceConverterKMSAutokeyConfig() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: KMSAutokeyConfigAssetType,
		Convert:   GetKMSAutokeyConfigCaiObject,
	}
}

func GetKMSAutokeyConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//cloudkms.googleapis.com/folders/{{folder}}/autokeyConfig")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetKMSAutokeyConfigApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: KMSAutokeyConfigAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/cloudkms/v1/rest",
				DiscoveryName:        "AutokeyConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetKMSAutokeyConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	keyProjectProp, err := expandKMSAutokeyConfigKeyProject(d.Get("key_project"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("key_project"); !tpgresource.IsEmptyValue(reflect.ValueOf(keyProjectProp)) && (ok || !reflect.DeepEqual(v, keyProjectProp)) {
		obj["keyProject"] = keyProjectProp
	}

	return obj, nil
}

func expandKMSAutokeyConfigKeyProject(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
