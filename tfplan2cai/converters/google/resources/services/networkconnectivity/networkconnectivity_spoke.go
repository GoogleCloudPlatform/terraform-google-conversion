// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networkconnectivity/Spoke.yaml
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

const NetworkConnectivitySpokeAssetType string = "networkconnectivity.googleapis.com/Spoke"

func ResourceConverterNetworkConnectivitySpoke() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetworkConnectivitySpokeAssetType,
		Convert:   GetNetworkConnectivitySpokeCaiObject,
	}
}

func GetNetworkConnectivitySpokeCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//networkconnectivity.googleapis.com/projects/{{project}}/locations/{{location}}/spokes/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetworkConnectivitySpokeApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetworkConnectivitySpokeAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networkconnectivity/v1/rest",
				DiscoveryName:        "Spoke",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetworkConnectivitySpokeApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandNetworkConnectivitySpokeName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandNetworkConnectivitySpokeDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	hubProp, err := expandNetworkConnectivitySpokeHub(d.Get("hub"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("hub"); !tpgresource.IsEmptyValue(reflect.ValueOf(hubProp)) && (ok || !reflect.DeepEqual(v, hubProp)) {
		obj["hub"] = hubProp
	}
	groupProp, err := expandNetworkConnectivitySpokeGroup(d.Get("group"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("group"); !tpgresource.IsEmptyValue(reflect.ValueOf(groupProp)) && (ok || !reflect.DeepEqual(v, groupProp)) {
		obj["group"] = groupProp
	}
	linkedVpnTunnelsProp, err := expandNetworkConnectivitySpokeLinkedVpnTunnels(d.Get("linked_vpn_tunnels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("linked_vpn_tunnels"); !tpgresource.IsEmptyValue(reflect.ValueOf(linkedVpnTunnelsProp)) && (ok || !reflect.DeepEqual(v, linkedVpnTunnelsProp)) {
		obj["linkedVpnTunnels"] = linkedVpnTunnelsProp
	}
	linkedInterconnectAttachmentsProp, err := expandNetworkConnectivitySpokeLinkedInterconnectAttachments(d.Get("linked_interconnect_attachments"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("linked_interconnect_attachments"); !tpgresource.IsEmptyValue(reflect.ValueOf(linkedInterconnectAttachmentsProp)) && (ok || !reflect.DeepEqual(v, linkedInterconnectAttachmentsProp)) {
		obj["linkedInterconnectAttachments"] = linkedInterconnectAttachmentsProp
	}
	linkedRouterApplianceInstancesProp, err := expandNetworkConnectivitySpokeLinkedRouterApplianceInstances(d.Get("linked_router_appliance_instances"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("linked_router_appliance_instances"); !tpgresource.IsEmptyValue(reflect.ValueOf(linkedRouterApplianceInstancesProp)) && (ok || !reflect.DeepEqual(v, linkedRouterApplianceInstancesProp)) {
		obj["linkedRouterApplianceInstances"] = linkedRouterApplianceInstancesProp
	}
	linkedVpcNetworkProp, err := expandNetworkConnectivitySpokeLinkedVpcNetwork(d.Get("linked_vpc_network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("linked_vpc_network"); !tpgresource.IsEmptyValue(reflect.ValueOf(linkedVpcNetworkProp)) && (ok || !reflect.DeepEqual(v, linkedVpcNetworkProp)) {
		obj["linkedVpcNetwork"] = linkedVpcNetworkProp
	}
	linkedProducerVpcNetworkProp, err := expandNetworkConnectivitySpokeLinkedProducerVpcNetwork(d.Get("linked_producer_vpc_network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("linked_producer_vpc_network"); !tpgresource.IsEmptyValue(reflect.ValueOf(linkedProducerVpcNetworkProp)) && (ok || !reflect.DeepEqual(v, linkedProducerVpcNetworkProp)) {
		obj["linkedProducerVpcNetwork"] = linkedProducerVpcNetworkProp
	}
	labelsProp, err := expandNetworkConnectivitySpokeEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetworkConnectivitySpokeName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeHub(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeGroup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedVpnTunnels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUris, err := expandNetworkConnectivitySpokeLinkedVpnTunnelsUris(original["uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uris"] = transformedUris
	}

	transformedSiteToSiteDataTransfer, err := expandNetworkConnectivitySpokeLinkedVpnTunnelsSiteToSiteDataTransfer(original["site_to_site_data_transfer"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSiteToSiteDataTransfer); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["siteToSiteDataTransfer"] = transformedSiteToSiteDataTransfer
	}

	transformedIncludeImportRanges, err := expandNetworkConnectivitySpokeLinkedVpnTunnelsIncludeImportRanges(original["include_import_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludeImportRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["includeImportRanges"] = transformedIncludeImportRanges
	}

	return transformed, nil
}

func expandNetworkConnectivitySpokeLinkedVpnTunnelsUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedVpnTunnelsSiteToSiteDataTransfer(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedVpnTunnelsIncludeImportRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedInterconnectAttachments(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUris, err := expandNetworkConnectivitySpokeLinkedInterconnectAttachmentsUris(original["uris"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUris); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uris"] = transformedUris
	}

	transformedSiteToSiteDataTransfer, err := expandNetworkConnectivitySpokeLinkedInterconnectAttachmentsSiteToSiteDataTransfer(original["site_to_site_data_transfer"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSiteToSiteDataTransfer); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["siteToSiteDataTransfer"] = transformedSiteToSiteDataTransfer
	}

	transformedIncludeImportRanges, err := expandNetworkConnectivitySpokeLinkedInterconnectAttachmentsIncludeImportRanges(original["include_import_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludeImportRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["includeImportRanges"] = transformedIncludeImportRanges
	}

	return transformed, nil
}

func expandNetworkConnectivitySpokeLinkedInterconnectAttachmentsUris(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedInterconnectAttachmentsSiteToSiteDataTransfer(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedInterconnectAttachmentsIncludeImportRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedRouterApplianceInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedInstances, err := expandNetworkConnectivitySpokeLinkedRouterApplianceInstancesInstances(original["instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstances); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["instances"] = transformedInstances
	}

	transformedSiteToSiteDataTransfer, err := expandNetworkConnectivitySpokeLinkedRouterApplianceInstancesSiteToSiteDataTransfer(original["site_to_site_data_transfer"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSiteToSiteDataTransfer); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["siteToSiteDataTransfer"] = transformedSiteToSiteDataTransfer
	}

	transformedIncludeImportRanges, err := expandNetworkConnectivitySpokeLinkedRouterApplianceInstancesIncludeImportRanges(original["include_import_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludeImportRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["includeImportRanges"] = transformedIncludeImportRanges
	}

	return transformed, nil
}

func expandNetworkConnectivitySpokeLinkedRouterApplianceInstancesInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedVirtualMachine, err := expandNetworkConnectivitySpokeLinkedRouterApplianceInstancesInstancesVirtualMachine(original["virtual_machine"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedVirtualMachine); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["virtualMachine"] = transformedVirtualMachine
		}

		transformedIpAddress, err := expandNetworkConnectivitySpokeLinkedRouterApplianceInstancesInstancesIpAddress(original["ip_address"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpAddress); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipAddress"] = transformedIpAddress
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandNetworkConnectivitySpokeLinkedRouterApplianceInstancesInstancesVirtualMachine(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedRouterApplianceInstancesInstancesIpAddress(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedRouterApplianceInstancesSiteToSiteDataTransfer(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedRouterApplianceInstancesIncludeImportRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedVpcNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUri, err := expandNetworkConnectivitySpokeLinkedVpcNetworkUri(original["uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uri"] = transformedUri
	}

	transformedExcludeExportRanges, err := expandNetworkConnectivitySpokeLinkedVpcNetworkExcludeExportRanges(original["exclude_export_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExcludeExportRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["excludeExportRanges"] = transformedExcludeExportRanges
	}

	transformedIncludeExportRanges, err := expandNetworkConnectivitySpokeLinkedVpcNetworkIncludeExportRanges(original["include_export_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludeExportRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["includeExportRanges"] = transformedIncludeExportRanges
	}

	return transformed, nil
}

func expandNetworkConnectivitySpokeLinkedVpcNetworkUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedVpcNetworkExcludeExportRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedVpcNetworkIncludeExportRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedProducerVpcNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNetwork, err := expandNetworkConnectivitySpokeLinkedProducerVpcNetworkNetwork(original["network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["network"] = transformedNetwork
	}

	transformedPeering, err := expandNetworkConnectivitySpokeLinkedProducerVpcNetworkPeering(original["peering"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPeering); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["peering"] = transformedPeering
	}

	transformedProducerNetwork, err := expandNetworkConnectivitySpokeLinkedProducerVpcNetworkProducerNetwork(original["producer_network"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProducerNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["producerNetwork"] = transformedProducerNetwork
	}

	transformedIncludeExportRanges, err := expandNetworkConnectivitySpokeLinkedProducerVpcNetworkIncludeExportRanges(original["include_export_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludeExportRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["includeExportRanges"] = transformedIncludeExportRanges
	}

	transformedExcludeExportRanges, err := expandNetworkConnectivitySpokeLinkedProducerVpcNetworkExcludeExportRanges(original["exclude_export_ranges"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExcludeExportRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["excludeExportRanges"] = transformedExcludeExportRanges
	}

	return transformed, nil
}

func expandNetworkConnectivitySpokeLinkedProducerVpcNetworkNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedProducerVpcNetworkPeering(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedProducerVpcNetworkProducerNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedProducerVpcNetworkIncludeExportRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeLinkedProducerVpcNetworkExcludeExportRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivitySpokeEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
