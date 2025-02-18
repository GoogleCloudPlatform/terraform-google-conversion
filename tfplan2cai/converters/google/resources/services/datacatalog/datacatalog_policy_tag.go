// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/datacatalog/PolicyTag.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package datacatalog

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DataCatalogPolicyTagAssetType string = "datacatalog.googleapis.com/PolicyTag"

func ResourceConverterDataCatalogPolicyTag() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DataCatalogPolicyTagAssetType,
		Convert:   GetDataCatalogPolicyTagCaiObject,
	}
}

func GetDataCatalogPolicyTagCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//datacatalog.googleapis.com/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDataCatalogPolicyTagApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DataCatalogPolicyTagAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/datacatalog/v1beta1/rest",
				DiscoveryName:        "PolicyTag",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDataCatalogPolicyTagApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDataCatalogPolicyTagDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandDataCatalogPolicyTagDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	parentPolicyTagProp, err := expandDataCatalogPolicyTagParentPolicyTag(d.Get("parent_policy_tag"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("parent_policy_tag"); !tpgresource.IsEmptyValue(reflect.ValueOf(parentPolicyTagProp)) && (ok || !reflect.DeepEqual(v, parentPolicyTagProp)) {
		obj["parentPolicyTag"] = parentPolicyTagProp
	}

	return obj, nil
}

func expandDataCatalogPolicyTagDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogPolicyTagDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataCatalogPolicyTagParentPolicyTag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
