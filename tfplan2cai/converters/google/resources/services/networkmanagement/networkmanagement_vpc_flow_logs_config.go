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

package networkmanagement

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetworkManagementVpcFlowLogsConfigAssetType string = "networkmanagement.googleapis.com/VpcFlowLogsConfig"

func ResourceConverterNetworkManagementVpcFlowLogsConfig() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetworkManagementVpcFlowLogsConfigAssetType,
		Convert:   GetNetworkManagementVpcFlowLogsConfigCaiObject,
	}
}

func GetNetworkManagementVpcFlowLogsConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//networkmanagement.googleapis.com/projects/{{project}}/locations/{{location}}/vpcFlowLogsConfigs/{{vpc_flow_logs_config_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetworkManagementVpcFlowLogsConfigApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetworkManagementVpcFlowLogsConfigAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networkmanagement/v1beta1/rest",
				DiscoveryName:        "VpcFlowLogsConfig",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetworkManagementVpcFlowLogsConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkManagementVpcFlowLogsConfigDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	stateProp, err := expandNetworkManagementVpcFlowLogsConfigState(d.Get("state"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("state"); !tpgresource.IsEmptyValue(reflect.ValueOf(stateProp)) && (ok || !reflect.DeepEqual(v, stateProp)) {
		obj["state"] = stateProp
	}
	aggregationIntervalProp, err := expandNetworkManagementVpcFlowLogsConfigAggregationInterval(d.Get("aggregation_interval"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("aggregation_interval"); !tpgresource.IsEmptyValue(reflect.ValueOf(aggregationIntervalProp)) && (ok || !reflect.DeepEqual(v, aggregationIntervalProp)) {
		obj["aggregationInterval"] = aggregationIntervalProp
	}
	flowSamplingProp, err := expandNetworkManagementVpcFlowLogsConfigFlowSampling(d.Get("flow_sampling"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("flow_sampling"); !tpgresource.IsEmptyValue(reflect.ValueOf(flowSamplingProp)) && (ok || !reflect.DeepEqual(v, flowSamplingProp)) {
		obj["flowSampling"] = flowSamplingProp
	}
	metadataProp, err := expandNetworkManagementVpcFlowLogsConfigMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("metadata"); !tpgresource.IsEmptyValue(reflect.ValueOf(metadataProp)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}
	metadataFieldsProp, err := expandNetworkManagementVpcFlowLogsConfigMetadataFields(d.Get("metadata_fields"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("metadata_fields"); !tpgresource.IsEmptyValue(reflect.ValueOf(metadataFieldsProp)) && (ok || !reflect.DeepEqual(v, metadataFieldsProp)) {
		obj["metadataFields"] = metadataFieldsProp
	}
	filterExprProp, err := expandNetworkManagementVpcFlowLogsConfigFilterExpr(d.Get("filter_expr"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("filter_expr"); !tpgresource.IsEmptyValue(reflect.ValueOf(filterExprProp)) && (ok || !reflect.DeepEqual(v, filterExprProp)) {
		obj["filterExpr"] = filterExprProp
	}
	interconnectAttachmentProp, err := expandNetworkManagementVpcFlowLogsConfigInterconnectAttachment(d.Get("interconnect_attachment"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("interconnect_attachment"); !tpgresource.IsEmptyValue(reflect.ValueOf(interconnectAttachmentProp)) && (ok || !reflect.DeepEqual(v, interconnectAttachmentProp)) {
		obj["interconnectAttachment"] = interconnectAttachmentProp
	}
	vpnTunnelProp, err := expandNetworkManagementVpcFlowLogsConfigVpnTunnel(d.Get("vpn_tunnel"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("vpn_tunnel"); !tpgresource.IsEmptyValue(reflect.ValueOf(vpnTunnelProp)) && (ok || !reflect.DeepEqual(v, vpnTunnelProp)) {
		obj["vpnTunnel"] = vpnTunnelProp
	}
	labelsProp, err := expandNetworkManagementVpcFlowLogsConfigEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetworkManagementVpcFlowLogsConfigDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementVpcFlowLogsConfigState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementVpcFlowLogsConfigAggregationInterval(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementVpcFlowLogsConfigFlowSampling(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementVpcFlowLogsConfigMetadata(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementVpcFlowLogsConfigMetadataFields(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementVpcFlowLogsConfigFilterExpr(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementVpcFlowLogsConfigInterconnectAttachment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementVpcFlowLogsConfigVpnTunnel(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkManagementVpcFlowLogsConfigEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}