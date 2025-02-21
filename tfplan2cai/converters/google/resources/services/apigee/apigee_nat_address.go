// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/apigee/NatAddress.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package apigee

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ApigeeNatAddressAssetType string = "apigee.googleapis.com/NatAddress"

func ResourceConverterApigeeNatAddress() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ApigeeNatAddressAssetType,
		Convert:   GetApigeeNatAddressCaiObject,
	}
}

func GetApigeeNatAddressCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//apigee.googleapis.com/{{instance_id}}/natAddresses/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetApigeeNatAddressApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ApigeeNatAddressAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/apigee/v1/rest",
				DiscoveryName:        "NatAddress",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetApigeeNatAddressApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandApigeeNatAddressName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	activateProp, err := expandApigeeNatAddressActivate(d.Get("activate"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("activate"); !tpgresource.IsEmptyValue(reflect.ValueOf(activateProp)) && (ok || !reflect.DeepEqual(v, activateProp)) {
		obj["activate"] = activateProp
	}

	return resourceApigeeNatAddressEncoder(d, config, obj)
}

func resourceApigeeNatAddressEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// cannot include activate prop in the body
	delete(obj, "activate")
	return obj, nil
}

func expandApigeeNatAddressName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeNatAddressActivate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
