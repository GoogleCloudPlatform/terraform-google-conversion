// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networksecurity/AddressGroup.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package networksecurity

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetworkSecurityAddressGroupAssetType string = "networksecurity.googleapis.com/AddressGroup"

func ResourceConverterNetworkSecurityAddressGroup() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetworkSecurityAddressGroupAssetType,
		Convert:   GetNetworkSecurityAddressGroupCaiObject,
	}
}

func GetNetworkSecurityAddressGroupCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//networksecurity.googleapis.com/{{parent}}/locations/{{location}}/addressGroups/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetworkSecurityAddressGroupApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetworkSecurityAddressGroupAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networksecurity/v1beta1/rest",
				DiscoveryName:        "AddressGroup",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetworkSecurityAddressGroupApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkSecurityAddressGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	typeProp, err := expandNetworkSecurityAddressGroupType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	itemsProp, err := expandNetworkSecurityAddressGroupItems(d.Get("items"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("items"); !tpgresource.IsEmptyValue(reflect.ValueOf(itemsProp)) && (ok || !reflect.DeepEqual(v, itemsProp)) {
		obj["items"] = itemsProp
	}
	capacityProp, err := expandNetworkSecurityAddressGroupCapacity(d.Get("capacity"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("capacity"); !tpgresource.IsEmptyValue(reflect.ValueOf(capacityProp)) && (ok || !reflect.DeepEqual(v, capacityProp)) {
		obj["capacity"] = capacityProp
	}
	purposeProp, err := expandNetworkSecurityAddressGroupPurpose(d.Get("purpose"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("purpose"); !tpgresource.IsEmptyValue(reflect.ValueOf(purposeProp)) && (ok || !reflect.DeepEqual(v, purposeProp)) {
		obj["purpose"] = purposeProp
	}
	labelsProp, err := expandNetworkSecurityAddressGroupEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetworkSecurityAddressGroupDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityAddressGroupType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityAddressGroupItems(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityAddressGroupCapacity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityAddressGroupPurpose(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkSecurityAddressGroupEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
