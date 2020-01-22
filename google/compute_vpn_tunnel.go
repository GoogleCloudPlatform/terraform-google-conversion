// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

package google

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func GetComputeVpnTunnelCaiObject(d TerraformResourceData, config *Config) (Asset, error) {
	name, err := assetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}}")
	if err != nil {
		return Asset{}, err
	}
	if obj, err := GetComputeVpnTunnelApiObject(d, config); err == nil {
		return Asset{
			Name: name,
			Type: "compute.googleapis.com/VpnTunnel",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/v1/rest",
				DiscoveryName:        "VpnTunnel",
				Data:                 obj,
			},
		}, nil
	} else {
		return Asset{}, err
	}
}

func GetComputeVpnTunnelApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeVpnTunnelName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeVpnTunnelDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	targetVpnGatewayProp, err := expandComputeVpnTunnelTargetVpnGateway(d.Get("target_vpn_gateway"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target_vpn_gateway"); !isEmptyValue(reflect.ValueOf(targetVpnGatewayProp)) && (ok || !reflect.DeepEqual(v, targetVpnGatewayProp)) {
		obj["targetVpnGateway"] = targetVpnGatewayProp
	}
	routerProp, err := expandComputeVpnTunnelRouter(d.Get("router"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("router"); !isEmptyValue(reflect.ValueOf(routerProp)) && (ok || !reflect.DeepEqual(v, routerProp)) {
		obj["router"] = routerProp
	}
	peerIpProp, err := expandComputeVpnTunnelPeerIp(d.Get("peer_ip"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("peer_ip"); !isEmptyValue(reflect.ValueOf(peerIpProp)) && (ok || !reflect.DeepEqual(v, peerIpProp)) {
		obj["peerIp"] = peerIpProp
	}
	sharedSecretProp, err := expandComputeVpnTunnelSharedSecret(d.Get("shared_secret"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("shared_secret"); !isEmptyValue(reflect.ValueOf(sharedSecretProp)) && (ok || !reflect.DeepEqual(v, sharedSecretProp)) {
		obj["sharedSecret"] = sharedSecretProp
	}
	ikeVersionProp, err := expandComputeVpnTunnelIkeVersion(d.Get("ike_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ike_version"); !isEmptyValue(reflect.ValueOf(ikeVersionProp)) && (ok || !reflect.DeepEqual(v, ikeVersionProp)) {
		obj["ikeVersion"] = ikeVersionProp
	}
	localTrafficSelectorProp, err := expandComputeVpnTunnelLocalTrafficSelector(d.Get("local_traffic_selector"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("local_traffic_selector"); !isEmptyValue(reflect.ValueOf(localTrafficSelectorProp)) && (ok || !reflect.DeepEqual(v, localTrafficSelectorProp)) {
		obj["localTrafficSelector"] = localTrafficSelectorProp
	}
	remoteTrafficSelectorProp, err := expandComputeVpnTunnelRemoteTrafficSelector(d.Get("remote_traffic_selector"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("remote_traffic_selector"); !isEmptyValue(reflect.ValueOf(remoteTrafficSelectorProp)) && (ok || !reflect.DeepEqual(v, remoteTrafficSelectorProp)) {
		obj["remoteTrafficSelector"] = remoteTrafficSelectorProp
	}
	regionProp, err := expandComputeVpnTunnelRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return resourceComputeVpnTunnelEncoder(d, config, obj)
}

func resourceComputeVpnTunnelEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	config := meta.(*Config)
	f, err := parseRegionalFieldValue("targetVpnGateways", d.Get("target_vpn_gateway").(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, err
	}
	if _, ok := d.GetOk("project"); !ok {
		d.Set("project", f.Project)
	}
	if _, ok := d.GetOk("region"); !ok {
		d.Set("region", f.Region)
	}
	return obj, nil
}

func expandComputeVpnTunnelName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelTargetVpnGateway(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseRegionalFieldValue("targetVpnGateways", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for target_vpn_gateway: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeVpnTunnelRouter(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if v == nil || v.(string) == "" {
		return "", nil
	}
	f, err := parseRegionalFieldValue("routers", v.(string), "project", "region", "zone", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for router: %s", err)
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}"+f.RelativeLink())
	if err != nil {
		return nil, err
	}

	return url, nil
}

func expandComputeVpnTunnelPeerIp(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelSharedSecret(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelIkeVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeVpnTunnelLocalTrafficSelector(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeVpnTunnelRemoteTrafficSelector(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeVpnTunnelRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
