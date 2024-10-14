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

package oracledatabase

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const OracleDatabaseCloudExadataInfrastructureAssetType string = "oracledatabase.googleapis.com/CloudExadataInfrastructure"

func ResourceConverterOracleDatabaseCloudExadataInfrastructure() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: OracleDatabaseCloudExadataInfrastructureAssetType,
		Convert:   GetOracleDatabaseCloudExadataInfrastructureCaiObject,
	}
}

func GetOracleDatabaseCloudExadataInfrastructureCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//oracledatabase.googleapis.com/projects/{{project}}/locations/{{location}}/cloudExadataInfrastructures/{{cloud_exadata_infrastructure_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetOracleDatabaseCloudExadataInfrastructureApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: OracleDatabaseCloudExadataInfrastructureAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/oracledatabase/v1/rest",
				DiscoveryName:        "CloudExadataInfrastructure",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetOracleDatabaseCloudExadataInfrastructureApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandOracleDatabaseCloudExadataInfrastructureDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	gcpOracleZoneProp, err := expandOracleDatabaseCloudExadataInfrastructureGcpOracleZone(d.Get("gcp_oracle_zone"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("gcp_oracle_zone"); !tpgresource.IsEmptyValue(reflect.ValueOf(gcpOracleZoneProp)) && (ok || !reflect.DeepEqual(v, gcpOracleZoneProp)) {
		obj["gcpOracleZone"] = gcpOracleZoneProp
	}
	propertiesProp, err := expandOracleDatabaseCloudExadataInfrastructureProperties(d.Get("properties"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("properties"); !tpgresource.IsEmptyValue(reflect.ValueOf(propertiesProp)) && (ok || !reflect.DeepEqual(v, propertiesProp)) {
		obj["properties"] = propertiesProp
	}
	labelsProp, err := expandOracleDatabaseCloudExadataInfrastructureEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandOracleDatabaseCloudExadataInfrastructureDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructureGcpOracleZone(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructureProperties(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedOcid, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesOcid(original["ocid"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOcid); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ocid"] = transformedOcid
	}

	transformedComputeCount, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesComputeCount(original["compute_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedComputeCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["computeCount"] = transformedComputeCount
	}

	transformedStorageCount, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesStorageCount(original["storage_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStorageCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["storageCount"] = transformedStorageCount
	}

	transformedTotalStorageSizeGb, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesTotalStorageSizeGb(original["total_storage_size_gb"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTotalStorageSizeGb); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["totalStorageSizeGb"] = transformedTotalStorageSizeGb
	}

	transformedAvailableStorageSizeGb, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesAvailableStorageSizeGb(original["available_storage_size_gb"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAvailableStorageSizeGb); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["availableStorageSizeGb"] = transformedAvailableStorageSizeGb
	}

	transformedMaintenanceWindow, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow(original["maintenance_window"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaintenanceWindow); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maintenanceWindow"] = transformedMaintenanceWindow
	}

	transformedState, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesState(original["state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedState); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["state"] = transformedState
	}

	transformedShape, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesShape(original["shape"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedShape); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["shape"] = transformedShape
	}

	transformedOciUrl, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesOciUrl(original["oci_url"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOciUrl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ociUrl"] = transformedOciUrl
	}

	transformedCpuCount, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesCpuCount(original["cpu_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCpuCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cpuCount"] = transformedCpuCount
	}

	transformedMaxCpuCount, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesMaxCpuCount(original["max_cpu_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxCpuCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxCpuCount"] = transformedMaxCpuCount
	}

	transformedMemorySizeGb, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesMemorySizeGb(original["memory_size_gb"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMemorySizeGb); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["memorySizeGb"] = transformedMemorySizeGb
	}

	transformedMaxMemoryGb, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesMaxMemoryGb(original["max_memory_gb"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxMemoryGb); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxMemoryGb"] = transformedMaxMemoryGb
	}

	transformedDbNodeStorageSizeGb, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesDbNodeStorageSizeGb(original["db_node_storage_size_gb"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDbNodeStorageSizeGb); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dbNodeStorageSizeGb"] = transformedDbNodeStorageSizeGb
	}

	transformedMaxDbNodeStorageSizeGb, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesMaxDbNodeStorageSizeGb(original["max_db_node_storage_size_gb"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxDbNodeStorageSizeGb); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxDbNodeStorageSizeGb"] = transformedMaxDbNodeStorageSizeGb
	}

	transformedDataStorageSizeTb, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesDataStorageSizeTb(original["data_storage_size_tb"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDataStorageSizeTb); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dataStorageSizeTb"] = transformedDataStorageSizeTb
	}

	transformedMaxDataStorageTb, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesMaxDataStorageTb(original["max_data_storage_tb"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxDataStorageTb); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxDataStorageTb"] = transformedMaxDataStorageTb
	}

	transformedActivatedStorageCount, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesActivatedStorageCount(original["activated_storage_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedActivatedStorageCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["activatedStorageCount"] = transformedActivatedStorageCount
	}

	transformedAdditionalStorageCount, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesAdditionalStorageCount(original["additional_storage_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAdditionalStorageCount); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["additionalStorageCount"] = transformedAdditionalStorageCount
	}

	transformedDbServerVersion, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesDbServerVersion(original["db_server_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDbServerVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["dbServerVersion"] = transformedDbServerVersion
	}

	transformedStorageServerVersion, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesStorageServerVersion(original["storage_server_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStorageServerVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["storageServerVersion"] = transformedStorageServerVersion
	}

	transformedNextMaintenanceRunId, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesNextMaintenanceRunId(original["next_maintenance_run_id"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNextMaintenanceRunId); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nextMaintenanceRunId"] = transformedNextMaintenanceRunId
	}

	transformedNextMaintenanceRunTime, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesNextMaintenanceRunTime(original["next_maintenance_run_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNextMaintenanceRunTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nextMaintenanceRunTime"] = transformedNextMaintenanceRunTime
	}

	transformedNextSecurityMaintenanceRunTime, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesNextSecurityMaintenanceRunTime(original["next_security_maintenance_run_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNextSecurityMaintenanceRunTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["nextSecurityMaintenanceRunTime"] = transformedNextSecurityMaintenanceRunTime
	}

	transformedCustomerContacts, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesCustomerContacts(original["customer_contacts"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCustomerContacts); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["customerContacts"] = transformedCustomerContacts
	}

	transformedMonthlyStorageServerVersion, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesMonthlyStorageServerVersion(original["monthly_storage_server_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMonthlyStorageServerVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["monthlyStorageServerVersion"] = transformedMonthlyStorageServerVersion
	}

	transformedMonthlyDbServerVersion, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesMonthlyDbServerVersion(original["monthly_db_server_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMonthlyDbServerVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["monthlyDbServerVersion"] = transformedMonthlyDbServerVersion
	}

	return transformed, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesOcid(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesComputeCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesStorageCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesTotalStorageSizeGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesAvailableStorageSizeGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedPreference, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowPreference(original["preference"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPreference); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["preference"] = transformedPreference
	}

	transformedMonths, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowMonths(original["months"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMonths); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["months"] = transformedMonths
	}

	transformedWeeksOfMonth, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowWeeksOfMonth(original["weeks_of_month"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWeeksOfMonth); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["weeksOfMonth"] = transformedWeeksOfMonth
	}

	transformedDaysOfWeek, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowDaysOfWeek(original["days_of_week"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDaysOfWeek); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["daysOfWeek"] = transformedDaysOfWeek
	}

	transformedHoursOfDay, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowHoursOfDay(original["hours_of_day"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHoursOfDay); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["hoursOfDay"] = transformedHoursOfDay
	}

	transformedLeadTimeWeek, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowLeadTimeWeek(original["lead_time_week"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLeadTimeWeek); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["leadTimeWeek"] = transformedLeadTimeWeek
	}

	transformedPatchingMode, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowPatchingMode(original["patching_mode"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPatchingMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["patchingMode"] = transformedPatchingMode
	}

	transformedCustomActionTimeoutMins, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowCustomActionTimeoutMins(original["custom_action_timeout_mins"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCustomActionTimeoutMins); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["customActionTimeoutMins"] = transformedCustomActionTimeoutMins
	}

	transformedIsCustomActionTimeoutEnabled, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowIsCustomActionTimeoutEnabled(original["is_custom_action_timeout_enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIsCustomActionTimeoutEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["isCustomActionTimeoutEnabled"] = transformedIsCustomActionTimeoutEnabled
	}

	return transformed, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowPreference(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowMonths(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowWeeksOfMonth(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowDaysOfWeek(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowHoursOfDay(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowLeadTimeWeek(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowPatchingMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowCustomActionTimeoutMins(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowIsCustomActionTimeoutEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesState(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesShape(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesOciUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesCpuCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesMaxCpuCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesMemorySizeGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesMaxMemoryGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesDbNodeStorageSizeGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesMaxDbNodeStorageSizeGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesDataStorageSizeTb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesMaxDataStorageTb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesActivatedStorageCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesAdditionalStorageCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesDbServerVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesStorageServerVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesNextMaintenanceRunId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesNextMaintenanceRunTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesNextSecurityMaintenanceRunTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesCustomerContacts(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedEmail, err := expandOracleDatabaseCloudExadataInfrastructurePropertiesCustomerContactsEmail(original["email"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["email"] = transformedEmail
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesCustomerContactsEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesMonthlyStorageServerVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructurePropertiesMonthlyDbServerVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandOracleDatabaseCloudExadataInfrastructureEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}