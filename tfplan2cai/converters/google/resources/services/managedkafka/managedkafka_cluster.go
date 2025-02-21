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

package managedkafka

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ManagedKafkaClusterAssetType string = "managedkafka.googleapis.com/Cluster"

func ResourceConverterManagedKafkaCluster() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ManagedKafkaClusterAssetType,
		Convert:   GetManagedKafkaClusterCaiObject,
	}
}

func GetManagedKafkaClusterCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//managedkafka.googleapis.com/projects/{{project}}/locations/{{location}}/clusters/{{cluster_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetManagedKafkaClusterApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ManagedKafkaClusterAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/managedkafka/v1/rest",
				DiscoveryName:        "Cluster",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetManagedKafkaClusterApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	gcpConfigProp, err := expandManagedKafkaClusterGcpConfig(d.Get("gcp_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("gcp_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(gcpConfigProp)) && (ok || !reflect.DeepEqual(v, gcpConfigProp)) {
		obj["gcpConfig"] = gcpConfigProp
	}
	capacityConfigProp, err := expandManagedKafkaClusterCapacityConfig(d.Get("capacity_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("capacity_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(capacityConfigProp)) && (ok || !reflect.DeepEqual(v, capacityConfigProp)) {
		obj["capacityConfig"] = capacityConfigProp
	}
	rebalanceConfigProp, err := expandManagedKafkaClusterRebalanceConfig(d.Get("rebalance_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("rebalance_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(rebalanceConfigProp)) && (ok || !reflect.DeepEqual(v, rebalanceConfigProp)) {
		obj["rebalanceConfig"] = rebalanceConfigProp
	}
	labelsProp, err := expandManagedKafkaClusterEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandManagedKafkaClusterGcpConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAccessConfig, err := expandManagedKafkaClusterGcpConfigAccessConfig(original["access_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAccessConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["accessConfig"] = transformedAccessConfig
	}

	transformedKmsKey, err := expandManagedKafkaClusterGcpConfigKmsKey(original["kms_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKey"] = transformedKmsKey
	}

	return transformed, nil
}

func expandManagedKafkaClusterGcpConfigAccessConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNetworkConfigs, err := expandManagedKafkaClusterGcpConfigAccessConfigNetworkConfigs(original["network_configs"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNetworkConfigs); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["networkConfigs"] = transformedNetworkConfigs
	}

	return transformed, nil
}

func expandManagedKafkaClusterGcpConfigAccessConfigNetworkConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedSubnet, err := expandManagedKafkaClusterGcpConfigAccessConfigNetworkConfigsSubnet(original["subnet"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSubnet); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["subnet"] = transformedSubnet
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandManagedKafkaClusterGcpConfigAccessConfigNetworkConfigsSubnet(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandManagedKafkaClusterGcpConfigKmsKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandManagedKafkaClusterCapacityConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedVcpuCount, err := expandManagedKafkaClusterCapacityConfigVcpuCount(original["vcpu_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVcpuCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vcpuCount"] = transformedVcpuCount
	}

	transformedMemoryBytes, err := expandManagedKafkaClusterCapacityConfigMemoryBytes(original["memory_bytes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMemoryBytes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["memoryBytes"] = transformedMemoryBytes
	}

	return transformed, nil
}

func expandManagedKafkaClusterCapacityConfigVcpuCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandManagedKafkaClusterCapacityConfigMemoryBytes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandManagedKafkaClusterRebalanceConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMode, err := expandManagedKafkaClusterRebalanceConfigMode(original["mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mode"] = transformedMode
	}

	return transformed, nil
}

func expandManagedKafkaClusterRebalanceConfigMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandManagedKafkaClusterEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
