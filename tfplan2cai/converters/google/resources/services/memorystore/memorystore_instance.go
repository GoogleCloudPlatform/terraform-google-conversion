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

package memorystore

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const MemorystoreInstanceAssetType string = "memorystore.googleapis.com/Instance"

func ResourceConverterMemorystoreInstance() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: MemorystoreInstanceAssetType,
		Convert:   GetMemorystoreInstanceCaiObject,
	}
}

func GetMemorystoreInstanceCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//memorystore.googleapis.com/projects/{{project}}/locations/{{location}}/instances/{{instance_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetMemorystoreInstanceApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: MemorystoreInstanceAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/memorystore/v1beta/rest",
				DiscoveryName:        "Instance",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetMemorystoreInstanceApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	replicaCountProp, err := expandMemorystoreInstanceReplicaCount(d.Get("replica_count"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("replica_count"); !tpgresource.IsEmptyValue(reflect.ValueOf(replicaCountProp)) && (ok || !reflect.DeepEqual(v, replicaCountProp)) {
		obj["replicaCount"] = replicaCountProp
	}
	authorizationModeProp, err := expandMemorystoreInstanceAuthorizationMode(d.Get("authorization_mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("authorization_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(authorizationModeProp)) && (ok || !reflect.DeepEqual(v, authorizationModeProp)) {
		obj["authorizationMode"] = authorizationModeProp
	}
	transitEncryptionModeProp, err := expandMemorystoreInstanceTransitEncryptionMode(d.Get("transit_encryption_mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("transit_encryption_mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(transitEncryptionModeProp)) && (ok || !reflect.DeepEqual(v, transitEncryptionModeProp)) {
		obj["transitEncryptionMode"] = transitEncryptionModeProp
	}
	shardCountProp, err := expandMemorystoreInstanceShardCount(d.Get("shard_count"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("shard_count"); !tpgresource.IsEmptyValue(reflect.ValueOf(shardCountProp)) && (ok || !reflect.DeepEqual(v, shardCountProp)) {
		obj["shardCount"] = shardCountProp
	}
	nodeTypeProp, err := expandMemorystoreInstanceNodeType(d.Get("node_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("node_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeTypeProp)) && (ok || !reflect.DeepEqual(v, nodeTypeProp)) {
		obj["nodeType"] = nodeTypeProp
	}
	persistenceConfigProp, err := expandMemorystoreInstancePersistenceConfig(d.Get("persistence_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("persistence_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(persistenceConfigProp)) && (ok || !reflect.DeepEqual(v, persistenceConfigProp)) {
		obj["persistenceConfig"] = persistenceConfigProp
	}
	engineVersionProp, err := expandMemorystoreInstanceEngineVersion(d.Get("engine_version"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("engine_version"); !tpgresource.IsEmptyValue(reflect.ValueOf(engineVersionProp)) && (ok || !reflect.DeepEqual(v, engineVersionProp)) {
		obj["engineVersion"] = engineVersionProp
	}
	engineConfigsProp, err := expandMemorystoreInstanceEngineConfigs(d.Get("engine_configs"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("engine_configs"); !tpgresource.IsEmptyValue(reflect.ValueOf(engineConfigsProp)) && (ok || !reflect.DeepEqual(v, engineConfigsProp)) {
		obj["engineConfigs"] = engineConfigsProp
	}
	zoneDistributionConfigProp, err := expandMemorystoreInstanceZoneDistributionConfig(d.Get("zone_distribution_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("zone_distribution_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(zoneDistributionConfigProp)) && (ok || !reflect.DeepEqual(v, zoneDistributionConfigProp)) {
		obj["zoneDistributionConfig"] = zoneDistributionConfigProp
	}
	deletionProtectionEnabledProp, err := expandMemorystoreInstanceDeletionProtectionEnabled(d.Get("deletion_protection_enabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("deletion_protection_enabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(deletionProtectionEnabledProp)) && (ok || !reflect.DeepEqual(v, deletionProtectionEnabledProp)) {
		obj["deletionProtectionEnabled"] = deletionProtectionEnabledProp
	}
	modeProp, err := expandMemorystoreInstanceMode(d.Get("mode"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(modeProp)) && (ok || !reflect.DeepEqual(v, modeProp)) {
		obj["mode"] = modeProp
	}
	labelsProp, err := expandMemorystoreInstanceEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return resourceMemorystoreInstanceEncoder(d, config, obj)
}

func resourceMemorystoreInstanceEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	v, ok := d.GetOk("desired_psc_auto_connections")
	if !ok {
		return obj, nil // No desired connections, nothing to update
	}
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		desiredConnection := raw.(map[string]interface{})
		connectionReq := make(map[string]interface{})

		projectId := desiredConnection["project_id"]
		if val := reflect.ValueOf(projectId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			connectionReq["projectId"] = projectId
		}

		network := desiredConnection["network"]
		if val := reflect.ValueOf(network); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			connectionReq["network"] = network
		}

		req = append(req, connectionReq)
	}
	obj["pscAutoConnections"] = req
	return obj, nil
}

func expandMemorystoreInstanceReplicaCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceAuthorizationMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceTransitEncryptionMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceShardCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceNodeType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstancePersistenceConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMode, err := expandMemorystoreInstancePersistenceConfigMode(original["mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mode"] = transformedMode
	}

	transformedRdbConfig, err := expandMemorystoreInstancePersistenceConfigRdbConfig(original["rdb_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRdbConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rdbConfig"] = transformedRdbConfig
	}

	transformedAofConfig, err := expandMemorystoreInstancePersistenceConfigAofConfig(original["aof_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAofConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["aofConfig"] = transformedAofConfig
	}

	return transformed, nil
}

func expandMemorystoreInstancePersistenceConfigMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstancePersistenceConfigRdbConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRdbSnapshotPeriod, err := expandMemorystoreInstancePersistenceConfigRdbConfigRdbSnapshotPeriod(original["rdb_snapshot_period"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRdbSnapshotPeriod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rdbSnapshotPeriod"] = transformedRdbSnapshotPeriod
	}

	transformedRdbSnapshotStartTime, err := expandMemorystoreInstancePersistenceConfigRdbConfigRdbSnapshotStartTime(original["rdb_snapshot_start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRdbSnapshotStartTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rdbSnapshotStartTime"] = transformedRdbSnapshotStartTime
	}

	return transformed, nil
}

func expandMemorystoreInstancePersistenceConfigRdbConfigRdbSnapshotPeriod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstancePersistenceConfigRdbConfigRdbSnapshotStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstancePersistenceConfigAofConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAppendFsync, err := expandMemorystoreInstancePersistenceConfigAofConfigAppendFsync(original["append_fsync"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAppendFsync); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["appendFsync"] = transformedAppendFsync
	}

	return transformed, nil
}

func expandMemorystoreInstancePersistenceConfigAofConfigAppendFsync(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceEngineVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceEngineConfigs(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandMemorystoreInstanceZoneDistributionConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedZone, err := expandMemorystoreInstanceZoneDistributionConfigZone(original["zone"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedZone); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["zone"] = transformedZone
	}

	transformedMode, err := expandMemorystoreInstanceZoneDistributionConfigMode(original["mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mode"] = transformedMode
	}

	return transformed, nil
}

func expandMemorystoreInstanceZoneDistributionConfigZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceZoneDistributionConfigMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceDeletionProtectionEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
