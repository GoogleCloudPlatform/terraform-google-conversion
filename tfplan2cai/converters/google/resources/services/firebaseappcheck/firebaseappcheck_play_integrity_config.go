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

package firebaseappcheck

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const FirebaseAppCheckPlayIntegrityConfigAssetType string = "firebaseappcheck.googleapis.com/PlayIntegrityConfig"

func ResourceConverterFirebaseAppCheckPlayIntegrityConfig() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: FirebaseAppCheckPlayIntegrityConfigAssetType,
		Convert:   GetFirebaseAppCheckPlayIntegrityConfigCaiObject,
	}
}

func GetFirebaseAppCheckPlayIntegrityConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//firebaseappcheck.googleapis.com/projects/{{project}}/apps/{{app_id}}/playIntegrityConfig")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetFirebaseAppCheckPlayIntegrityConfigApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: FirebaseAppCheckPlayIntegrityConfigAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/firebaseappcheck/v1beta/rest",
				DiscoveryName:        "PlayIntegrityConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetFirebaseAppCheckPlayIntegrityConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	tokenTtlProp, err := expandFirebaseAppCheckPlayIntegrityConfigTokenTtl(d.Get("token_ttl"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("token_ttl"); !tpgresource.IsEmptyValue(reflect.ValueOf(tokenTtlProp)) && (ok || !reflect.DeepEqual(v, tokenTtlProp)) {
		obj["tokenTtl"] = tokenTtlProp
	}

	return obj, nil
}

func expandFirebaseAppCheckPlayIntegrityConfigTokenTtl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
