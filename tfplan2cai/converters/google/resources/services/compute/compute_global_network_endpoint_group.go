// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/GlobalNetworkEndpointGroup.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ComputeGlobalNetworkEndpointGroupAssetType string = "compute.googleapis.com/GlobalNetworkEndpointGroup"

func ResourceConverterComputeGlobalNetworkEndpointGroup() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeGlobalNetworkEndpointGroupAssetType,
		Convert:   GetComputeGlobalNetworkEndpointGroupCaiObject,
	}
}

func GetComputeGlobalNetworkEndpointGroupCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/networkEndpointGroups/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeGlobalNetworkEndpointGroupApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeGlobalNetworkEndpointGroupAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "GlobalNetworkEndpointGroup",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeGlobalNetworkEndpointGroupApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeGlobalNetworkEndpointGroupName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeGlobalNetworkEndpointGroupDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	networkEndpointTypeProp, err := expandComputeGlobalNetworkEndpointGroupNetworkEndpointType(d.Get("network_endpoint_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network_endpoint_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkEndpointTypeProp)) && (ok || !reflect.DeepEqual(v, networkEndpointTypeProp)) {
		obj["networkEndpointType"] = networkEndpointTypeProp
	}
	defaultPortProp, err := expandComputeGlobalNetworkEndpointGroupDefaultPort(d.Get("default_port"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_port"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultPortProp)) && (ok || !reflect.DeepEqual(v, defaultPortProp)) {
		obj["defaultPort"] = defaultPortProp
	}

	return obj, nil
}

func expandComputeGlobalNetworkEndpointGroupName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalNetworkEndpointGroupDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalNetworkEndpointGroupNetworkEndpointType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalNetworkEndpointGroupDefaultPort(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
