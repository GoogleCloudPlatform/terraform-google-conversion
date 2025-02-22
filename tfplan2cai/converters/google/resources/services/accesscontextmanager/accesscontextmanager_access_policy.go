// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/accesscontextmanager/AccessPolicy.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package accesscontextmanager

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const AccessContextManagerAccessPolicyAssetType string = "accesscontextmanager.googleapis.com/AccessPolicy"

func ResourceConverterAccessContextManagerAccessPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: AccessContextManagerAccessPolicyAssetType,
		Convert:   GetAccessContextManagerAccessPolicyCaiObject,
	}
}

func GetAccessContextManagerAccessPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//accesscontextmanager.googleapis.com/accessPolicies/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetAccessContextManagerAccessPolicyApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: AccessContextManagerAccessPolicyAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/accesscontextmanager/v1/rest",
				DiscoveryName:        "AccessPolicy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetAccessContextManagerAccessPolicyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	parentProp, err := expandAccessContextManagerAccessPolicyParent(d.Get("parent"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("parent"); !tpgresource.IsEmptyValue(reflect.ValueOf(parentProp)) && (ok || !reflect.DeepEqual(v, parentProp)) {
		obj["parent"] = parentProp
	}
	titleProp, err := expandAccessContextManagerAccessPolicyTitle(d.Get("title"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("title"); !tpgresource.IsEmptyValue(reflect.ValueOf(titleProp)) && (ok || !reflect.DeepEqual(v, titleProp)) {
		obj["title"] = titleProp
	}
	scopesProp, err := expandAccessContextManagerAccessPolicyScopes(d.Get("scopes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("scopes"); !tpgresource.IsEmptyValue(reflect.ValueOf(scopesProp)) && (ok || !reflect.DeepEqual(v, scopesProp)) {
		obj["scopes"] = scopesProp
	}

	return obj, nil
}

func expandAccessContextManagerAccessPolicyParent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessPolicyTitle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAccessContextManagerAccessPolicyScopes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
