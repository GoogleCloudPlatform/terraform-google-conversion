// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/ResourcePolicy.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Suppresses a diff on cases like 1:00 when it should be 01:00.
// Because API will normalize this value
func HourlyFormatSuppressDiff(_, old, new string, _ *schema.ResourceData) bool {
	return old == "0"+new
}

const ComputeResourcePolicyAssetType string = "compute.googleapis.com/ResourcePolicy"

func ResourceConverterComputeResourcePolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ComputeResourcePolicyAssetType,
		Convert:   GetComputeResourcePolicyCaiObject,
	}
}

func GetComputeResourcePolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/regions/{{region}}/resourcePolicies/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetComputeResourcePolicyApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ComputeResourcePolicyAssetType,
			Resource: &cai.AssetResource{
				Version:              "beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/compute/beta/rest",
				DiscoveryName:        "ResourcePolicy",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetComputeResourcePolicyApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandComputeResourcePolicyName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandComputeResourcePolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	snapshotSchedulePolicyProp, err := expandComputeResourcePolicySnapshotSchedulePolicy(d.Get("snapshot_schedule_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("snapshot_schedule_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(snapshotSchedulePolicyProp)) && (ok || !reflect.DeepEqual(v, snapshotSchedulePolicyProp)) {
		obj["snapshotSchedulePolicy"] = snapshotSchedulePolicyProp
	}
	groupPlacementPolicyProp, err := expandComputeResourcePolicyGroupPlacementPolicy(d.Get("group_placement_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("group_placement_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(groupPlacementPolicyProp)) && (ok || !reflect.DeepEqual(v, groupPlacementPolicyProp)) {
		obj["groupPlacementPolicy"] = groupPlacementPolicyProp
	}
	instanceSchedulePolicyProp, err := expandComputeResourcePolicyInstanceSchedulePolicy(d.Get("instance_schedule_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("instance_schedule_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(instanceSchedulePolicyProp)) && (ok || !reflect.DeepEqual(v, instanceSchedulePolicyProp)) {
		obj["instanceSchedulePolicy"] = instanceSchedulePolicyProp
	}
	diskConsistencyGroupPolicyProp, err := expandComputeResourcePolicyDiskConsistencyGroupPolicy(d.Get("disk_consistency_group_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disk_consistency_group_policy"); ok || !reflect.DeepEqual(v, diskConsistencyGroupPolicyProp) {
		obj["diskConsistencyGroupPolicy"] = diskConsistencyGroupPolicyProp
	}
	workloadPolicyProp, err := expandComputeResourcePolicyWorkloadPolicy(d.Get("workload_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("workload_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(workloadPolicyProp)) && (ok || !reflect.DeepEqual(v, workloadPolicyProp)) {
		obj["workloadPolicy"] = workloadPolicyProp
	}
	regionProp, err := expandComputeResourcePolicyRegion(d.Get("region"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("region"); !tpgresource.IsEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	return obj, nil
}

func expandComputeResourcePolicyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSchedule, err := expandComputeResourcePolicySnapshotSchedulePolicySchedule(original["schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSchedule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["schedule"] = transformedSchedule
	}

	transformedRetentionPolicy, err := expandComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy(original["retention_policy"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRetentionPolicy); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["retentionPolicy"] = transformedRetentionPolicy
	}

	transformedSnapshotProperties, err := expandComputeResourcePolicySnapshotSchedulePolicySnapshotProperties(original["snapshot_properties"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSnapshotProperties); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["snapshotProperties"] = transformedSnapshotProperties
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicySchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHourlySchedule, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule(original["hourly_schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHourlySchedule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["hourlySchedule"] = transformedHourlySchedule
	}

	transformedDailySchedule, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule(original["daily_schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDailySchedule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dailySchedule"] = transformedDailySchedule
	}

	transformedWeeklySchedule, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule(original["weekly_schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWeeklySchedule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["weeklySchedule"] = transformedWeeklySchedule
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHoursInCycle, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleHoursInCycle(original["hours_in_cycle"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHoursInCycle); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["hoursInCycle"] = transformedHoursInCycle
	}

	transformedStartTime, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleHoursInCycle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleHourlyScheduleStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDaysInCycle, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleDaysInCycle(original["days_in_cycle"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDaysInCycle); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["daysInCycle"] = transformedDaysInCycle
	}

	transformedStartTime, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleDaysInCycle(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleDailyScheduleStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDayOfWeeks, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeks(original["day_of_weeks"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDayOfWeeks); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dayOfWeeks"] = transformedDayOfWeeks
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedStartTime, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeksStartTime(original["start_time"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["startTime"] = transformedStartTime
		}

		transformedDay, err := expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeksDay(original["day"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDay); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["day"] = transformedDay
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeksStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeksDay(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMaxRetentionDays, err := expandComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyMaxRetentionDays(original["max_retention_days"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxRetentionDays); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxRetentionDays"] = transformedMaxRetentionDays
	}

	transformedOnSourceDiskDelete, err := expandComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOnSourceDiskDelete(original["on_source_disk_delete"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOnSourceDiskDelete); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["onSourceDiskDelete"] = transformedOnSourceDiskDelete
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyMaxRetentionDays(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOnSourceDiskDelete(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicySnapshotProperties(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLabels, err := expandComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesLabels(original["labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["labels"] = transformedLabels
	}

	transformedStorageLocations, err := expandComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesStorageLocations(original["storage_locations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStorageLocations); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["storageLocations"] = transformedStorageLocations
	}

	transformedGuestFlush, err := expandComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesGuestFlush(original["guest_flush"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGuestFlush); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["guestFlush"] = transformedGuestFlush
	}

	transformedChainName, err := expandComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesChainName(original["chain_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedChainName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["chainName"] = transformedChainName
	}

	return transformed, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesStorageLocations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesGuestFlush(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesChainName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicyGroupPlacementPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedVmCount, err := expandComputeResourcePolicyGroupPlacementPolicyVmCount(original["vm_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVmCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vmCount"] = transformedVmCount
	}

	transformedAvailabilityDomainCount, err := expandComputeResourcePolicyGroupPlacementPolicyAvailabilityDomainCount(original["availability_domain_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAvailabilityDomainCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["availabilityDomainCount"] = transformedAvailabilityDomainCount
	}

	transformedCollocation, err := expandComputeResourcePolicyGroupPlacementPolicyCollocation(original["collocation"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCollocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["collocation"] = transformedCollocation
	}

	transformedMaxDistance, err := expandComputeResourcePolicyGroupPlacementPolicyMaxDistance(original["max_distance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxDistance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxDistance"] = transformedMaxDistance
	}

	transformedGpuTopology, err := expandComputeResourcePolicyGroupPlacementPolicyGpuTopology(original["gpu_topology"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGpuTopology); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gpuTopology"] = transformedGpuTopology
	}

	transformedTpuTopology, err := expandComputeResourcePolicyGroupPlacementPolicyTpuTopology(original["tpu_topology"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTpuTopology); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tpuTopology"] = transformedTpuTopology
	}

	return transformed, nil
}

func expandComputeResourcePolicyGroupPlacementPolicyVmCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicyGroupPlacementPolicyAvailabilityDomainCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicyGroupPlacementPolicyCollocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicyGroupPlacementPolicyMaxDistance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicyGroupPlacementPolicyGpuTopology(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicyGroupPlacementPolicyTpuTopology(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicyInstanceSchedulePolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedVmStartSchedule, err := expandComputeResourcePolicyInstanceSchedulePolicyVmStartSchedule(original["vm_start_schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVmStartSchedule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vmStartSchedule"] = transformedVmStartSchedule
	}

	transformedVmStopSchedule, err := expandComputeResourcePolicyInstanceSchedulePolicyVmStopSchedule(original["vm_stop_schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedVmStopSchedule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["vmStopSchedule"] = transformedVmStopSchedule
	}

	transformedTimeZone, err := expandComputeResourcePolicyInstanceSchedulePolicyTimeZone(original["time_zone"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTimeZone); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["timeZone"] = transformedTimeZone
	}

	transformedStartTime, err := expandComputeResourcePolicyInstanceSchedulePolicyStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	transformedExpirationTime, err := expandComputeResourcePolicyInstanceSchedulePolicyExpirationTime(original["expiration_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExpirationTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["expirationTime"] = transformedExpirationTime
	}

	return transformed, nil
}

func expandComputeResourcePolicyInstanceSchedulePolicyVmStartSchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSchedule, err := expandComputeResourcePolicyInstanceSchedulePolicyVmStartScheduleSchedule(original["schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSchedule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["schedule"] = transformedSchedule
	}

	return transformed, nil
}

func expandComputeResourcePolicyInstanceSchedulePolicyVmStartScheduleSchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicyInstanceSchedulePolicyVmStopSchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSchedule, err := expandComputeResourcePolicyInstanceSchedulePolicyVmStopScheduleSchedule(original["schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSchedule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["schedule"] = transformedSchedule
	}

	return transformed, nil
}

func expandComputeResourcePolicyInstanceSchedulePolicyVmStopScheduleSchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicyInstanceSchedulePolicyTimeZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicyInstanceSchedulePolicyStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicyInstanceSchedulePolicyExpirationTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicyDiskConsistencyGroupPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})

	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	if isEnabled, ok := original["enabled"]; ok {
		if !isEnabled.(bool) {
			return nil, nil
		}
	}
	transformed := make(map[string]interface{})
	return transformed, nil
}

func expandComputeResourcePolicyWorkloadPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedType, err := expandComputeResourcePolicyWorkloadPolicyType(original["type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["type"] = transformedType
	}

	transformedMaxTopologyDistance, err := expandComputeResourcePolicyWorkloadPolicyMaxTopologyDistance(original["max_topology_distance"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxTopologyDistance); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxTopologyDistance"] = transformedMaxTopologyDistance
	}

	transformedAcceleratorTopology, err := expandComputeResourcePolicyWorkloadPolicyAcceleratorTopology(original["accelerator_topology"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAcceleratorTopology); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["acceleratorTopology"] = transformedAcceleratorTopology
	}

	return transformed, nil
}

func expandComputeResourcePolicyWorkloadPolicyType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicyWorkloadPolicyMaxTopologyDistance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicyWorkloadPolicyAcceleratorTopology(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputeResourcePolicyRegion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	f, err := tpgresource.ParseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
