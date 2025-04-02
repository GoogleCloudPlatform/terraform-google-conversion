// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/memorystore/Instance.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package memorystore

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
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
	automatedBackupConfigProp, err := expandMemorystoreInstanceAutomatedBackupConfig(d.Get("automated_backup_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("automated_backup_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(automatedBackupConfigProp)) && (ok || !reflect.DeepEqual(v, automatedBackupConfigProp)) {
		obj["automatedBackupConfig"] = automatedBackupConfigProp
	}
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
	maintenancePolicyProp, err := expandMemorystoreInstanceMaintenancePolicy(d.Get("maintenance_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("maintenance_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(maintenancePolicyProp)) && (ok || !reflect.DeepEqual(v, maintenancePolicyProp)) {
		obj["maintenancePolicy"] = maintenancePolicyProp
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
	crossInstanceReplicationConfigProp, err := expandMemorystoreInstanceCrossInstanceReplicationConfig(d.Get("cross_instance_replication_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cross_instance_replication_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(crossInstanceReplicationConfigProp)) && (ok || !reflect.DeepEqual(v, crossInstanceReplicationConfigProp)) {
		obj["crossInstanceReplicationConfig"] = crossInstanceReplicationConfigProp
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
	// if the automated_backup_config is not defined, automatedBackupMode needs to be passed and set to DISABLED in the expand
	if obj["automatedBackupConfig"] == nil {
		config := meta.(*transport_tpg.Config)
		automatedBackupConfigProp, _ := expandMemorystoreInstanceAutomatedBackupConfig(d.Get("automated_backup_config"), d, config)
		obj["automatedBackupConfig"] = automatedBackupConfigProp
	}
	return obj, nil
}

func expandMemorystoreInstanceAutomatedBackupConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})

	// The automated_backup_config block is not specified, so automatedBackupMode should be DISABLED
	transformed := make(map[string]interface{})
	if len(d.Get("automated_backup_config").([]interface{})) < 1 {
		transformed["automatedBackupMode"] = "DISABLED"
		return transformed, nil
	}
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})

	// The automated_backup_config block is specified, so automatedBackupMode should be ENABLED
	transformed["automatedBackupMode"] = "ENABLED"
	transformedFixedFrequencySchedule, err := expandMemorystoreInstanceAutomatedBackupConfigFixedFrequencySchedule(original["fixed_frequency_schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFixedFrequencySchedule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fixedFrequencySchedule"] = transformedFixedFrequencySchedule
	}

	transformedRetention, err := expandMemorystoreInstanceAutomatedBackupConfigRetention(original["retention"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRetention); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["retention"] = transformedRetention
	}

	return transformed, nil
}

func expandMemorystoreInstanceAutomatedBackupConfigFixedFrequencySchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStartTime, err := expandMemorystoreInstanceAutomatedBackupConfigFixedFrequencyScheduleStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	return transformed, nil
}

func expandMemorystoreInstanceAutomatedBackupConfigFixedFrequencyScheduleStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHours, err := expandMemorystoreInstanceAutomatedBackupConfigFixedFrequencyScheduleStartTimeHours(original["hours"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHours); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["hours"] = transformedHours
	}

	return transformed, nil
}

func expandMemorystoreInstanceAutomatedBackupConfigFixedFrequencyScheduleStartTimeHours(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceAutomatedBackupConfigRetention(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
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

func expandMemorystoreInstanceMaintenancePolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCreateTime, err := expandMemorystoreInstanceMaintenancePolicyCreateTime(original["create_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCreateTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["createTime"] = transformedCreateTime
	}

	transformedUpdateTime, err := expandMemorystoreInstanceMaintenancePolicyUpdateTime(original["update_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUpdateTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["updateTime"] = transformedUpdateTime
	}

	transformedWeeklyMaintenanceWindow, err := expandMemorystoreInstanceMaintenancePolicyWeeklyMaintenanceWindow(original["weekly_maintenance_window"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWeeklyMaintenanceWindow); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["weeklyMaintenanceWindow"] = transformedWeeklyMaintenanceWindow
	}

	return transformed, nil
}

func expandMemorystoreInstanceMaintenancePolicyCreateTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceMaintenancePolicyUpdateTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceMaintenancePolicyWeeklyMaintenanceWindow(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDay, err := expandMemorystoreInstanceMaintenancePolicyWeeklyMaintenanceWindowDay(original["day"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDay); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["day"] = transformedDay
		}

		transformedDuration, err := expandMemorystoreInstanceMaintenancePolicyWeeklyMaintenanceWindowDuration(original["duration"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDuration); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["duration"] = transformedDuration
		}

		transformedStartTime, err := expandMemorystoreInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(original["start_time"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["startTime"] = transformedStartTime
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandMemorystoreInstanceMaintenancePolicyWeeklyMaintenanceWindowDay(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceMaintenancePolicyWeeklyMaintenanceWindowDuration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHours, err := expandMemorystoreInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeHours(original["hours"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHours); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["hours"] = transformedHours
	}

	transformedMinutes, err := expandMemorystoreInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeMinutes(original["minutes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinutes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minutes"] = transformedMinutes
	}

	transformedSeconds, err := expandMemorystoreInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeSeconds(original["seconds"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSeconds); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["seconds"] = transformedSeconds
	}

	transformedNanos, err := expandMemorystoreInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeNanos(original["nanos"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNanos); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nanos"] = transformedNanos
	}

	return transformed, nil
}

func expandMemorystoreInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeHours(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeMinutes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeSeconds(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTimeNanos(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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

func expandMemorystoreInstanceCrossInstanceReplicationConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedInstanceRole, err := expandMemorystoreInstanceCrossInstanceReplicationConfigInstanceRole(original["instance_role"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstanceRole); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["instanceRole"] = transformedInstanceRole
	}

	transformedPrimaryInstance, err := expandMemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstance(original["primary_instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrimaryInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["primaryInstance"] = transformedPrimaryInstance
	}

	transformedSecondaryInstances, err := expandMemorystoreInstanceCrossInstanceReplicationConfigSecondaryInstances(original["secondary_instances"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSecondaryInstances); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["secondaryInstances"] = transformedSecondaryInstances
	}

	transformedMembership, err := expandMemorystoreInstanceCrossInstanceReplicationConfigMembership(original["membership"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMembership); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["membership"] = transformedMembership
	}

	transformedUpdateTime, err := expandMemorystoreInstanceCrossInstanceReplicationConfigUpdateTime(original["update_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUpdateTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["updateTime"] = transformedUpdateTime
	}

	return transformed, nil
}

func expandMemorystoreInstanceCrossInstanceReplicationConfigInstanceRole(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedInstance, err := expandMemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstanceInstance(original["instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["instance"] = transformedInstance
	}

	transformedUid, err := expandMemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstanceUid(original["uid"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUid); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uid"] = transformedUid
	}

	return transformed, nil
}

func expandMemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstanceInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstanceUid(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceCrossInstanceReplicationConfigSecondaryInstances(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedInstance, err := expandMemorystoreInstanceCrossInstanceReplicationConfigSecondaryInstancesInstance(original["instance"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["instance"] = transformedInstance
		}

		transformedUid, err := expandMemorystoreInstanceCrossInstanceReplicationConfigSecondaryInstancesUid(original["uid"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedUid); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["uid"] = transformedUid
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandMemorystoreInstanceCrossInstanceReplicationConfigSecondaryInstancesInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceCrossInstanceReplicationConfigSecondaryInstancesUid(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceCrossInstanceReplicationConfigMembership(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPrimaryInstance, err := expandMemorystoreInstanceCrossInstanceReplicationConfigMembershipPrimaryInstance(original["primary_instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrimaryInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["primaryInstance"] = transformedPrimaryInstance
	}

	transformedSecondaryInstance, err := expandMemorystoreInstanceCrossInstanceReplicationConfigMembershipSecondaryInstance(original["secondary_instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSecondaryInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["secondaryInstance"] = transformedSecondaryInstance
	}

	return transformed, nil
}

func expandMemorystoreInstanceCrossInstanceReplicationConfigMembershipPrimaryInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedInstance, err := expandMemorystoreInstanceCrossInstanceReplicationConfigMembershipPrimaryInstanceInstance(original["instance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["instance"] = transformedInstance
	}

	transformedUid, err := expandMemorystoreInstanceCrossInstanceReplicationConfigMembershipPrimaryInstanceUid(original["uid"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUid); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uid"] = transformedUid
	}

	return transformed, nil
}

func expandMemorystoreInstanceCrossInstanceReplicationConfigMembershipPrimaryInstanceInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceCrossInstanceReplicationConfigMembershipPrimaryInstanceUid(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceCrossInstanceReplicationConfigMembershipSecondaryInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedInstance, err := expandMemorystoreInstanceCrossInstanceReplicationConfigMembershipSecondaryInstanceInstance(original["instance"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedInstance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["instance"] = transformedInstance
		}

		transformedUid, err := expandMemorystoreInstanceCrossInstanceReplicationConfigMembershipSecondaryInstanceUid(original["uid"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedUid); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["uid"] = transformedUid
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandMemorystoreInstanceCrossInstanceReplicationConfigMembershipSecondaryInstanceInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceCrossInstanceReplicationConfigMembershipSecondaryInstanceUid(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMemorystoreInstanceCrossInstanceReplicationConfigUpdateTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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
