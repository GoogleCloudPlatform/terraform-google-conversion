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

package identityplatform

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const IdentityPlatformDefaultSupportedIdpConfigAssetType string = "identitytoolkit.googleapis.com/DefaultSupportedIdpConfig"

func ResourceConverterIdentityPlatformDefaultSupportedIdpConfig() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: IdentityPlatformDefaultSupportedIdpConfigAssetType,
		Convert:   GetIdentityPlatformDefaultSupportedIdpConfigCaiObject,
	}
}

func GetIdentityPlatformDefaultSupportedIdpConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//identitytoolkit.googleapis.com/projects/{{project}}/defaultSupportedIdpConfigs/{{idp_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetIdentityPlatformDefaultSupportedIdpConfigApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: IdentityPlatformDefaultSupportedIdpConfigAssetType,
			Resource: &cai.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/identitytoolkit/v2/rest",
				DiscoveryName:        "DefaultSupportedIdpConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetIdentityPlatformDefaultSupportedIdpConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	clientIdProp, err := expandIdentityPlatformDefaultSupportedIdpConfigClientId(d.Get("client_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("client_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(clientIdProp)) && (ok || !reflect.DeepEqual(v, clientIdProp)) {
		obj["clientId"] = clientIdProp
	}
	clientSecretProp, err := expandIdentityPlatformDefaultSupportedIdpConfigClientSecret(d.Get("client_secret"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("client_secret"); !tpgresource.IsEmptyValue(reflect.ValueOf(clientSecretProp)) && (ok || !reflect.DeepEqual(v, clientSecretProp)) {
		obj["clientSecret"] = clientSecretProp
	}
	enabledProp, err := expandIdentityPlatformDefaultSupportedIdpConfigEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(enabledProp)) && (ok || !reflect.DeepEqual(v, enabledProp)) {
		obj["enabled"] = enabledProp
	}

	return obj, nil
}

func expandIdentityPlatformDefaultSupportedIdpConfigClientId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformDefaultSupportedIdpConfigClientSecret(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIdentityPlatformDefaultSupportedIdpConfigEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
