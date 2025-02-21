// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networkconnectivity/Group.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package networkconnectivity

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetworkConnectivityGroupAssetType string = "networkconnectivity.googleapis.com/Group"

func ResourceConverterNetworkConnectivityGroup() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetworkConnectivityGroupAssetType,
		Convert:   GetNetworkConnectivityGroupCaiObject,
	}
}

func GetNetworkConnectivityGroupCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//networkconnectivity.googleapis.com/projects/{{project}}/locations/global/hubs/{{hub}}/groups/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetworkConnectivityGroupApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetworkConnectivityGroupAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networkconnectivity/v1/rest",
				DiscoveryName:        "Group",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetworkConnectivityGroupApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandNetworkConnectivityGroupName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandNetworkConnectivityGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	autoAcceptProp, err := expandNetworkConnectivityGroupAutoAccept(d.Get("auto_accept"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("auto_accept"); !tpgresource.IsEmptyValue(reflect.ValueOf(autoAcceptProp)) && (ok || !reflect.DeepEqual(v, autoAcceptProp)) {
		obj["autoAccept"] = autoAcceptProp
	}
	labelsProp, err := expandNetworkConnectivityGroupEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetworkConnectivityGroupName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityGroupDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityGroupAutoAccept(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAutoAcceptProjects, err := expandNetworkConnectivityGroupAutoAcceptAutoAcceptProjects(original["auto_accept_projects"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAutoAcceptProjects); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["autoAcceptProjects"] = transformedAutoAcceptProjects
	}

	return transformed, nil
}

func expandNetworkConnectivityGroupAutoAcceptAutoAcceptProjects(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityGroupEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
