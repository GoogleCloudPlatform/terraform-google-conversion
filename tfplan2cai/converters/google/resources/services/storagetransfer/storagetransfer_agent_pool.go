// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/storagetransfer/AgentPool.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package storagetransfer

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const StorageTransferAgentPoolAssetType string = "storagetransfer.googleapis.com/AgentPool"

func ResourceConverterStorageTransferAgentPool() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: StorageTransferAgentPoolAssetType,
		Convert:   GetStorageTransferAgentPoolCaiObject,
	}
}

func GetStorageTransferAgentPoolCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//storagetransfer.googleapis.com/projects/{{project}}/agentPools/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetStorageTransferAgentPoolApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: StorageTransferAgentPoolAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/storagetransfer/v1/rest",
				DiscoveryName:        "AgentPool",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetStorageTransferAgentPoolApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandStorageTransferAgentPoolDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	bandwidthLimitProp, err := expandStorageTransferAgentPoolBandwidthLimit(d.Get("bandwidth_limit"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("bandwidth_limit"); !tpgresource.IsEmptyValue(reflect.ValueOf(bandwidthLimitProp)) && (ok || !reflect.DeepEqual(v, bandwidthLimitProp)) {
		obj["bandwidthLimit"] = bandwidthLimitProp
	}

	return obj, nil
}

func expandStorageTransferAgentPoolDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageTransferAgentPoolBandwidthLimit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLimitMbps, err := expandStorageTransferAgentPoolBandwidthLimitLimitMbps(original["limit_mbps"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLimitMbps); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["limitMbps"] = transformedLimitMbps
	}

	return transformed, nil
}

func expandStorageTransferAgentPoolBandwidthLimitLimitMbps(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
