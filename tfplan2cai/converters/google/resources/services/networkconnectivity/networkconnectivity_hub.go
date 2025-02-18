// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networkconnectivity/Hub.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package networkconnectivity

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetworkConnectivityHubAssetType string = "networkconnectivity.googleapis.com/Hub"

func ResourceConverterNetworkConnectivityHub() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetworkConnectivityHubAssetType,
		Convert:   GetNetworkConnectivityHubCaiObject,
	}
}

func GetNetworkConnectivityHubCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//networkconnectivity.googleapis.com/projects/{{project}}/locations/global/hubs/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetworkConnectivityHubApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetworkConnectivityHubAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networkconnectivity/v1/rest",
				DiscoveryName:        "Hub",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetworkConnectivityHubApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandNetworkConnectivityHubName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandNetworkConnectivityHubDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	presetTopologyProp, err := expandNetworkConnectivityHubPresetTopology(d.Get("preset_topology"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("preset_topology"); !tpgresource.IsEmptyValue(reflect.ValueOf(presetTopologyProp)) && (ok || !reflect.DeepEqual(v, presetTopologyProp)) {
		obj["presetTopology"] = presetTopologyProp
	}
	exportPscProp, err := expandNetworkConnectivityHubExportPsc(d.Get("export_psc"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("export_psc"); !tpgresource.IsEmptyValue(reflect.ValueOf(exportPscProp)) && (ok || !reflect.DeepEqual(v, exportPscProp)) {
		obj["exportPsc"] = exportPscProp
	}
	labelsProp, err := expandNetworkConnectivityHubEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetworkConnectivityHubName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityHubDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityHubPresetTopology(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityHubExportPsc(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityHubEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
