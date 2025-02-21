// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/oslogin/SSHPublicKey.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package oslogin

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const OSLoginSSHPublicKeyAssetType string = "oslogin.googleapis.com/SSHPublicKey"

func ResourceConverterOSLoginSSHPublicKey() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: OSLoginSSHPublicKeyAssetType,
		Convert:   GetOSLoginSSHPublicKeyCaiObject,
	}
}

func GetOSLoginSSHPublicKeyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//oslogin.googleapis.com/users/{{user}}/sshPublicKeys/{{fingerprint}}/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetOSLoginSSHPublicKeyApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: OSLoginSSHPublicKeyAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/oslogin/v1/rest",
				DiscoveryName:        "SSHPublicKey",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetOSLoginSSHPublicKeyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	keyProp, err := expandOSLoginSSHPublicKeyKey(d.Get("key"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("key"); !tpgresource.IsEmptyValue(reflect.ValueOf(keyProp)) && (ok || !reflect.DeepEqual(v, keyProp)) {
		obj["key"] = keyProp
	}
	expirationTimeUsecProp, err := expandOSLoginSSHPublicKeyExpirationTimeUsec(d.Get("expiration_time_usec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("expiration_time_usec"); !tpgresource.IsEmptyValue(reflect.ValueOf(expirationTimeUsecProp)) && (ok || !reflect.DeepEqual(v, expirationTimeUsecProp)) {
		obj["expirationTimeUsec"] = expirationTimeUsecProp
	}

	return obj, nil
}

func expandOSLoginSSHPublicKeyKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOSLoginSSHPublicKeyExpirationTimeUsec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
