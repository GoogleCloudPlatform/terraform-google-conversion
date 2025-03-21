// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/vmwareengine/NetworkPeering.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package vmwareengine

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const VmwareengineNetworkPeeringAssetType string = "vmwareengine.googleapis.com/NetworkPeering"

func ResourceConverterVmwareengineNetworkPeering() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: VmwareengineNetworkPeeringAssetType,
		Convert:   GetVmwareengineNetworkPeeringCaiObject,
	}
}

func GetVmwareengineNetworkPeeringCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//vmwareengine.googleapis.com/projects/{{project}}/locations/global/networkPeerings/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetVmwareengineNetworkPeeringApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: VmwareengineNetworkPeeringAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/vmwareengine/v1/rest",
				DiscoveryName:        "NetworkPeering",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetVmwareengineNetworkPeeringApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	peerNetworkProp, err := expandVmwareengineNetworkPeeringPeerNetwork(d.Get("peer_network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("peer_network"); !tpgresource.IsEmptyValue(reflect.ValueOf(peerNetworkProp)) && (ok || !reflect.DeepEqual(v, peerNetworkProp)) {
		obj["peerNetwork"] = peerNetworkProp
	}
	exportCustomRoutesProp, err := expandVmwareengineNetworkPeeringExportCustomRoutes(d.Get("export_custom_routes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("export_custom_routes"); ok || !reflect.DeepEqual(v, exportCustomRoutesProp) {
		obj["exportCustomRoutes"] = exportCustomRoutesProp
	}
	importCustomRoutesProp, err := expandVmwareengineNetworkPeeringImportCustomRoutes(d.Get("import_custom_routes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("import_custom_routes"); ok || !reflect.DeepEqual(v, importCustomRoutesProp) {
		obj["importCustomRoutes"] = importCustomRoutesProp
	}
	exportCustomRoutesWithPublicIpProp, err := expandVmwareengineNetworkPeeringExportCustomRoutesWithPublicIp(d.Get("export_custom_routes_with_public_ip"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("export_custom_routes_with_public_ip"); ok || !reflect.DeepEqual(v, exportCustomRoutesWithPublicIpProp) {
		obj["exportCustomRoutesWithPublicIp"] = exportCustomRoutesWithPublicIpProp
	}
	importCustomRoutesWithPublicIpProp, err := expandVmwareengineNetworkPeeringImportCustomRoutesWithPublicIp(d.Get("import_custom_routes_with_public_ip"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("import_custom_routes_with_public_ip"); ok || !reflect.DeepEqual(v, importCustomRoutesWithPublicIpProp) {
		obj["importCustomRoutesWithPublicIp"] = importCustomRoutesWithPublicIpProp
	}
	peerNetworkTypeProp, err := expandVmwareengineNetworkPeeringPeerNetworkType(d.Get("peer_network_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("peer_network_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(peerNetworkTypeProp)) && (ok || !reflect.DeepEqual(v, peerNetworkTypeProp)) {
		obj["peerNetworkType"] = peerNetworkTypeProp
	}
	vmwareEngineNetworkProp, err := expandVmwareengineNetworkPeeringVmwareEngineNetwork(d.Get("vmware_engine_network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("vmware_engine_network"); !tpgresource.IsEmptyValue(reflect.ValueOf(vmwareEngineNetworkProp)) && (ok || !reflect.DeepEqual(v, vmwareEngineNetworkProp)) {
		obj["vmwareEngineNetwork"] = vmwareEngineNetworkProp
	}
	descriptionProp, err := expandVmwareengineNetworkPeeringDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	return obj, nil
}

func expandVmwareengineNetworkPeeringPeerNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineNetworkPeeringExportCustomRoutes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineNetworkPeeringImportCustomRoutes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineNetworkPeeringExportCustomRoutesWithPublicIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineNetworkPeeringImportCustomRoutesWithPublicIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineNetworkPeeringPeerNetworkType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineNetworkPeeringVmwareEngineNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandVmwareengineNetworkPeeringDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
