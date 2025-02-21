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

const ManagedKafkaTopicAssetType string = "managedkafka.googleapis.com/Topic"

func ResourceConverterManagedKafkaTopic() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ManagedKafkaTopicAssetType,
		Convert:   GetManagedKafkaTopicCaiObject,
	}
}

func GetManagedKafkaTopicCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//managedkafka.googleapis.com/projects/{{project}}/locations/{{location}}/clusters/{{cluster}}/topics/{{topic_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetManagedKafkaTopicApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ManagedKafkaTopicAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/managedkafka/v1/rest",
				DiscoveryName:        "Topic",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetManagedKafkaTopicApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	partitionCountProp, err := expandManagedKafkaTopicPartitionCount(d.Get("partition_count"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("partition_count"); !tpgresource.IsEmptyValue(reflect.ValueOf(partitionCountProp)) && (ok || !reflect.DeepEqual(v, partitionCountProp)) {
		obj["partitionCount"] = partitionCountProp
	}
	replicationFactorProp, err := expandManagedKafkaTopicReplicationFactor(d.Get("replication_factor"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("replication_factor"); !tpgresource.IsEmptyValue(reflect.ValueOf(replicationFactorProp)) && (ok || !reflect.DeepEqual(v, replicationFactorProp)) {
		obj["replicationFactor"] = replicationFactorProp
	}
	configsProp, err := expandManagedKafkaTopicConfigs(d.Get("configs"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("configs"); !tpgresource.IsEmptyValue(reflect.ValueOf(configsProp)) && (ok || !reflect.DeepEqual(v, configsProp)) {
		obj["configs"] = configsProp
	}

	return obj, nil
}

func expandManagedKafkaTopicPartitionCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandManagedKafkaTopicReplicationFactor(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandManagedKafkaTopicConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
