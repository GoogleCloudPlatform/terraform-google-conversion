// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/firebasehosting/CustomDomain.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package firebasehosting

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const FirebaseHostingCustomDomainAssetType string = "firebasehosting.googleapis.com/CustomDomain"

func ResourceConverterFirebaseHostingCustomDomain() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: FirebaseHostingCustomDomainAssetType,
		Convert:   GetFirebaseHostingCustomDomainCaiObject,
	}
}

func GetFirebaseHostingCustomDomainCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//firebasehosting.googleapis.com/projects/{{project}}/sites/{{site_id}}/customDomains/{{custom_domain}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetFirebaseHostingCustomDomainApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: FirebaseHostingCustomDomainAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/firebasehosting/v1beta1/rest",
				DiscoveryName:        "CustomDomain",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetFirebaseHostingCustomDomainApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	etagProp, err := expandFirebaseHostingCustomDomainEtag(d.Get("etag"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}
	certPreferenceProp, err := expandFirebaseHostingCustomDomainCertPreference(d.Get("cert_preference"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cert_preference"); !tpgresource.IsEmptyValue(reflect.ValueOf(certPreferenceProp)) && (ok || !reflect.DeepEqual(v, certPreferenceProp)) {
		obj["certPreference"] = certPreferenceProp
	}
	redirectTargetProp, err := expandFirebaseHostingCustomDomainRedirectTarget(d.Get("redirect_target"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("redirect_target"); !tpgresource.IsEmptyValue(reflect.ValueOf(redirectTargetProp)) && (ok || !reflect.DeepEqual(v, redirectTargetProp)) {
		obj["redirectTarget"] = redirectTargetProp
	}

	return obj, nil
}

func expandFirebaseHostingCustomDomainEtag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingCustomDomainCertPreference(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFirebaseHostingCustomDomainRedirectTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
