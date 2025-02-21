// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/Network.yaml
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

const ComputeNetworkAssetType string = "compute.googleapis.com/Network"

func ResourceConverterComputeNetwork() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeNetworkAssetType,
		Convert:   GetComputeNetworkCaiObject,
	}
}

func GetComputeNetworkCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/global/networks/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeNetworkApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeNetworkAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "Network",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeNetworkApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeNetworkDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeNetworkName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	autoCreateSubnetworksProp, err := expandComputeNetworkAutoCreateSubnetworks(d.Get("auto_create_subnetworks"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("auto_create_subnetworks"); ok || !reflect.DeepEqual(v, autoCreateSubnetworksProp) {
		obj["autoCreateSubnetworks"] = autoCreateSubnetworksProp
	}
	routingConfigProp, err := expandComputeNetworkRoutingConfig(nil, d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("routing_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(routingConfigProp)) && (ok || !reflect.DeepEqual(v, routingConfigProp)) {
		obj["routingConfig"] = routingConfigProp
	}
	mtuProp, err := expandComputeNetworkMtu(d.Get("mtu"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("mtu"); !tpgresource.IsEmptyValue(reflect.ValueOf(mtuProp)) && (ok || !reflect.DeepEqual(v, mtuProp)) {
		obj["mtu"] = mtuProp
	}
	enableUlaInternalIpv6Prop, err := expandComputeNetworkEnableUlaInternalIpv6(d.Get("enable_ula_internal_ipv6"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("enable_ula_internal_ipv6"); !tpgresource.IsEmptyValue(reflect.ValueOf(enableUlaInternalIpv6Prop)) && (ok || !reflect.DeepEqual(v, enableUlaInternalIpv6Prop)) {
		obj["enableUlaInternalIpv6"] = enableUlaInternalIpv6Prop
	}
	internalIpv6RangeProp, err := expandComputeNetworkInternalIpv6Range(d.Get("internal_ipv6_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("internal_ipv6_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(internalIpv6RangeProp)) && (ok || !reflect.DeepEqual(v, internalIpv6RangeProp)) {
		obj["internalIpv6Range"] = internalIpv6RangeProp
	}
	networkFirewallPolicyEnforcementOrderProp, err := expandComputeNetworkNetworkFirewallPolicyEnforcementOrder(d.Get("network_firewall_policy_enforcement_order"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network_firewall_policy_enforcement_order"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkFirewallPolicyEnforcementOrderProp)) && (ok || !reflect.DeepEqual(v, networkFirewallPolicyEnforcementOrderProp)) {
		obj["networkFirewallPolicyEnforcementOrder"] = networkFirewallPolicyEnforcementOrderProp
	}
	networkProfileProp, err := expandComputeNetworkNetworkProfile(d.Get("network_profile"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network_profile"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProfileProp)) && (ok || !reflect.DeepEqual(v, networkProfileProp)) {
		obj["networkProfile"] = networkProfileProp
	}

	return resourceComputeNetworkEncoder(d, config, obj)
}

func resourceComputeNetworkEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	delete(obj, "numeric_id") // Field doesn't exist in the API
	return obj, nil
}

func expandComputeNetworkDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkAutoCreateSubnetworks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkRoutingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedRoutingMode, err := expandComputeNetworkRoutingConfigRoutingMode(d.Get("routing_mode"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRoutingMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["routingMode"] = transformedRoutingMode
	}

	transformedBgpBestPathSelectionMode, err := expandComputeNetworkRoutingConfigBgpBestPathSelectionMode(d.Get("bgp_best_path_selection_mode"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBgpBestPathSelectionMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["bgpBestPathSelectionMode"] = transformedBgpBestPathSelectionMode
	}

	transformedBgpAlwaysCompareMed, err := expandComputeNetworkRoutingConfigBgpAlwaysCompareMed(d.Get("bgp_always_compare_med"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBgpAlwaysCompareMed); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["bgpAlwaysCompareMed"] = transformedBgpAlwaysCompareMed
	}

	transformedBgpInterRegionCost, err := expandComputeNetworkRoutingConfigBgpInterRegionCost(d.Get("bgp_inter_region_cost"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBgpInterRegionCost); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["bgpInterRegionCost"] = transformedBgpInterRegionCost
	}

	return transformed, nil
}

func expandComputeNetworkRoutingConfigRoutingMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkRoutingConfigBgpBestPathSelectionMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkRoutingConfigBgpAlwaysCompareMed(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkRoutingConfigBgpInterRegionCost(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkMtu(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkEnableUlaInternalIpv6(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkInternalIpv6Range(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkNetworkFirewallPolicyEnforcementOrder(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeNetworkNetworkProfile(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
