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

package gkebackup

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const GKEBackupBackupPlanAssetType string = "gkebackup.googleapis.com/BackupPlan"

func ResourceConverterGKEBackupBackupPlan() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: GKEBackupBackupPlanAssetType,
		Convert:   GetGKEBackupBackupPlanCaiObject,
	}
}

func GetGKEBackupBackupPlanCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//gkebackup.googleapis.com/projects/{{project}}/locations/{{location}}/backupPlans/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetGKEBackupBackupPlanApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: GKEBackupBackupPlanAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/gkebackup/v1/rest",
				DiscoveryName:        "BackupPlan",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetGKEBackupBackupPlanApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandGKEBackupBackupPlanName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandGKEBackupBackupPlanDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	clusterProp, err := expandGKEBackupBackupPlanCluster(d.Get("cluster"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("cluster"); !tpgresource.IsEmptyValue(reflect.ValueOf(clusterProp)) && (ok || !reflect.DeepEqual(v, clusterProp)) {
		obj["cluster"] = clusterProp
	}
	retentionPolicyProp, err := expandGKEBackupBackupPlanRetentionPolicy(d.Get("retention_policy"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("retention_policy"); !tpgresource.IsEmptyValue(reflect.ValueOf(retentionPolicyProp)) && (ok || !reflect.DeepEqual(v, retentionPolicyProp)) {
		obj["retentionPolicy"] = retentionPolicyProp
	}
	backupScheduleProp, err := expandGKEBackupBackupPlanBackupSchedule(d.Get("backup_schedule"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("backup_schedule"); !tpgresource.IsEmptyValue(reflect.ValueOf(backupScheduleProp)) && (ok || !reflect.DeepEqual(v, backupScheduleProp)) {
		obj["backupSchedule"] = backupScheduleProp
	}
	deactivatedProp, err := expandGKEBackupBackupPlanDeactivated(d.Get("deactivated"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("deactivated"); !tpgresource.IsEmptyValue(reflect.ValueOf(deactivatedProp)) && (ok || !reflect.DeepEqual(v, deactivatedProp)) {
		obj["deactivated"] = deactivatedProp
	}
	backupConfigProp, err := expandGKEBackupBackupPlanBackupConfig(d.Get("backup_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("backup_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(backupConfigProp)) && (ok || !reflect.DeepEqual(v, backupConfigProp)) {
		obj["backupConfig"] = backupConfigProp
	}
	labelsProp, err := expandGKEBackupBackupPlanEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandGKEBackupBackupPlanName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/backupPlans/{{name}}")
}

func expandGKEBackupBackupPlanDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanCluster(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanRetentionPolicy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBackupDeleteLockDays, err := expandGKEBackupBackupPlanRetentionPolicyBackupDeleteLockDays(original["backup_delete_lock_days"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBackupDeleteLockDays); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["backupDeleteLockDays"] = transformedBackupDeleteLockDays
	}

	transformedBackupRetainDays, err := expandGKEBackupBackupPlanRetentionPolicyBackupRetainDays(original["backup_retain_days"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBackupRetainDays); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["backupRetainDays"] = transformedBackupRetainDays
	}

	transformedLocked, err := expandGKEBackupBackupPlanRetentionPolicyLocked(original["locked"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocked); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["locked"] = transformedLocked
	}

	return transformed, nil
}

func expandGKEBackupBackupPlanRetentionPolicyBackupDeleteLockDays(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanRetentionPolicyBackupRetainDays(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanRetentionPolicyLocked(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanBackupSchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCronSchedule, err := expandGKEBackupBackupPlanBackupScheduleCronSchedule(original["cron_schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCronSchedule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cronSchedule"] = transformedCronSchedule
	}

	transformedPaused, err := expandGKEBackupBackupPlanBackupSchedulePaused(original["paused"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPaused); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["paused"] = transformedPaused
	}

	return transformed, nil
}

func expandGKEBackupBackupPlanBackupScheduleCronSchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanBackupSchedulePaused(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanDeactivated(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanBackupConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedIncludeVolumeData, err := expandGKEBackupBackupPlanBackupConfigIncludeVolumeData(original["include_volume_data"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludeVolumeData); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["includeVolumeData"] = transformedIncludeVolumeData
	}

	transformedIncludeSecrets, err := expandGKEBackupBackupPlanBackupConfigIncludeSecrets(original["include_secrets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludeSecrets); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["includeSecrets"] = transformedIncludeSecrets
	}

	transformedEncryptionKey, err := expandGKEBackupBackupPlanBackupConfigEncryptionKey(original["encryption_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEncryptionKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["encryptionKey"] = transformedEncryptionKey
	}

	transformedAllNamespaces, err := expandGKEBackupBackupPlanBackupConfigAllNamespaces(original["all_namespaces"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllNamespaces); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allNamespaces"] = transformedAllNamespaces
	}

	transformedSelectedNamespaces, err := expandGKEBackupBackupPlanBackupConfigSelectedNamespaces(original["selected_namespaces"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSelectedNamespaces); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["selectedNamespaces"] = transformedSelectedNamespaces
	}

	transformedSelectedApplications, err := expandGKEBackupBackupPlanBackupConfigSelectedApplications(original["selected_applications"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSelectedApplications); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["selectedApplications"] = transformedSelectedApplications
	}

	return transformed, nil
}

func expandGKEBackupBackupPlanBackupConfigIncludeVolumeData(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanBackupConfigIncludeSecrets(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanBackupConfigEncryptionKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGcpKmsEncryptionKey, err := expandGKEBackupBackupPlanBackupConfigEncryptionKeyGcpKmsEncryptionKey(original["gcp_kms_encryption_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGcpKmsEncryptionKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gcpKmsEncryptionKey"] = transformedGcpKmsEncryptionKey
	}

	return transformed, nil
}

func expandGKEBackupBackupPlanBackupConfigEncryptionKeyGcpKmsEncryptionKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanBackupConfigAllNamespaces(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanBackupConfigSelectedNamespaces(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNamespaces, err := expandGKEBackupBackupPlanBackupConfigSelectedNamespacesNamespaces(original["namespaces"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespaces); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["namespaces"] = transformedNamespaces
	}

	return transformed, nil
}

func expandGKEBackupBackupPlanBackupConfigSelectedNamespacesNamespaces(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanBackupConfigSelectedApplications(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedNamespacedNames, err := expandGKEBackupBackupPlanBackupConfigSelectedApplicationsNamespacedNames(original["namespaced_names"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedNamespacedNames); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["namespacedNames"] = transformedNamespacedNames
	}

	return transformed, nil
}

func expandGKEBackupBackupPlanBackupConfigSelectedApplicationsNamespacedNames(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNamespace, err := expandGKEBackupBackupPlanBackupConfigSelectedApplicationsNamespacedNamesNamespace(original["namespace"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNamespace); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["namespace"] = transformedNamespace
		}

		transformedName, err := expandGKEBackupBackupPlanBackupConfigSelectedApplicationsNamespacedNamesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandGKEBackupBackupPlanBackupConfigSelectedApplicationsNamespacedNamesNamespace(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanBackupConfigSelectedApplicationsNamespacedNamesName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupPlanEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
