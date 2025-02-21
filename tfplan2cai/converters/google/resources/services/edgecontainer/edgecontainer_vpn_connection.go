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

package edgecontainer

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const EdgecontainerVpnConnectionAssetType string = "edgecontainer.googleapis.com/VpnConnection"

func ResourceConverterEdgecontainerVpnConnection() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: EdgecontainerVpnConnectionAssetType,
		Convert:   GetEdgecontainerVpnConnectionCaiObject,
	}
}

func GetEdgecontainerVpnConnectionCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//edgecontainer.googleapis.com/projects/{{project}}/locations/{{location}}/vpnConnections/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetEdgecontainerVpnConnectionApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: EdgecontainerVpnConnectionAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/edgecontainer/v1/rest",
				DiscoveryName:        "VpnConnection",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetEdgecontainerVpnConnectionApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	natGatewayIpProp, err := expandEdgecontainerVpnConnectionNatGatewayIp(d.Get("nat_gateway_ip"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("nat_gateway_ip"); !tpgresource.IsEmptyValue(reflect.ValueOf(natGatewayIpProp)) && (ok || !reflect.DeepEqual(v, natGatewayIpProp)) {
		obj["natGatewayIp"] = natGatewayIpProp
	}
	clusterProp, err := expandEdgecontainerVpnConnectionCluster(d.Get("cluster"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cluster"); !tpgresource.IsEmptyValue(reflect.ValueOf(clusterProp)) && (ok || !reflect.DeepEqual(v, clusterProp)) {
		obj["cluster"] = clusterProp
	}
	vpcProp, err := expandEdgecontainerVpnConnectionVpc(d.Get("vpc"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("vpc"); !tpgresource.IsEmptyValue(reflect.ValueOf(vpcProp)) && (ok || !reflect.DeepEqual(v, vpcProp)) {
		obj["vpc"] = vpcProp
	}
	vpcProjectProp, err := expandEdgecontainerVpnConnectionVpcProject(d.Get("vpc_project"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("vpc_project"); !tpgresource.IsEmptyValue(reflect.ValueOf(vpcProjectProp)) && (ok || !reflect.DeepEqual(v, vpcProjectProp)) {
		obj["vpcProject"] = vpcProjectProp
	}
	enableHighAvailabilityProp, err := expandEdgecontainerVpnConnectionEnableHighAvailability(d.Get("enable_high_availability"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_high_availability"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableHighAvailabilityProp)) && (ok || !reflect.DeepEqual(v, enableHighAvailabilityProp)) {
		obj["enableHighAvailability"] = enableHighAvailabilityProp
	}
	routerProp, err := expandEdgecontainerVpnConnectionRouter(d.Get("router"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("router"); !tpgresource.IsEmptyValue(reflect.ValueOf(routerProp)) && (ok || !reflect.DeepEqual(v, routerProp)) {
		obj["router"] = routerProp
	}
	labelsProp, err := expandEdgecontainerVpnConnectionEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandEdgecontainerVpnConnectionNatGatewayIp(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerVpnConnectionCluster(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerVpnConnectionVpc(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerVpnConnectionVpcProject(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProjectId, err := expandEdgecontainerVpnConnectionVpcProjectProjectId(original["project_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProjectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["projectId"] = transformedProjectId
	}

	return transformed, nil
}

func expandEdgecontainerVpnConnectionVpcProjectProjectId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerVpnConnectionEnableHighAvailability(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerVpnConnectionRouter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandEdgecontainerVpnConnectionEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
