// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/bigquerydatatransfer/Config.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package bigquerydatatransfer

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

var sensitiveParams = []string{"secret_access_key"}

func sensitiveParamCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	for _, sp := range sensitiveParams {
		mapLabel := diff.Get("params." + sp).(string)
		authLabel := diff.Get("sensitive_params.0." + sp).(string)
		if mapLabel != "" && authLabel != "" {
			return fmt.Errorf("Sensitive param [%s] cannot be set in both `params` and the `sensitive_params` block.", sp)
		}
	}
	return nil
}

// This customizeDiff is to use ForceNew for params fields data_path_template or data_path and
// destination_table_name_template only if the value of "data_source_id" is "google_cloud_storage" or "amazon_s3".
func ParamsCustomizeDiffFunc(diff tpgresource.TerraformResourceDiff) error {
	old, new := diff.GetChange("params")
	dsId := diff.Get("data_source_id").(string)
	oldParams := old.(map[string]interface{})
	newParams := new.(map[string]interface{})
	var err error

	switch dsId {
	case "google_cloud_storage":
		if oldParams["data_path_template"] != nil && newParams["data_path_template"] != nil && oldParams["data_path_template"].(string) != newParams["data_path_template"].(string) {
			err = diff.ForceNew("params")
			if err != nil {
				return fmt.Errorf("ForceNew failed for params, old - %v and new - %v", oldParams, newParams)
			}
			return nil
		}

		if oldParams["destination_table_name_template"] != nil && newParams["destination_table_name_template"] != nil && oldParams["destination_table_name_template"].(string) != newParams["destination_table_name_template"].(string) {
			err = diff.ForceNew("params")
			if err != nil {
				return fmt.Errorf("ForceNew failed for params, old - %v and new - %v", oldParams, newParams)
			}
			return nil
		}
	case "amazon_s3":
		if oldParams["data_path"] != nil && newParams["data_path"] != nil && oldParams["data_path"].(string) != newParams["data_path"].(string) {
			err = diff.ForceNew("params")
			if err != nil {
				return fmt.Errorf("ForceNew failed for params, old - %v and new - %v", oldParams, newParams)
			}
			return nil
		}

		if oldParams["destination_table_name_template"] != nil && newParams["destination_table_name_template"] != nil && oldParams["destination_table_name_template"].(string) != newParams["destination_table_name_template"].(string) {
			err = diff.ForceNew("params")
			if err != nil {
				return fmt.Errorf("ForceNew failed for params, old - %v and new - %v", oldParams, newParams)
			}
			return nil
		}
	}
	return nil
}
func paramsCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	return ParamsCustomizeDiffFunc(diff)
}

const BigqueryDataTransferConfigAssetType string = "bigquerydatatransfer.googleapis.com/Config"

func ResourceConverterBigqueryDataTransferConfig() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: BigqueryDataTransferConfigAssetType,
		Convert:   GetBigqueryDataTransferConfigCaiObject,
	}
}

func GetBigqueryDataTransferConfigCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//bigquerydatatransfer.googleapis.com/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetBigqueryDataTransferConfigApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: BigqueryDataTransferConfigAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/bigquerydatatransfer/v1/rest",
				DiscoveryName:        "Config",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetBigqueryDataTransferConfigApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandBigqueryDataTransferConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	destinationDatasetIdProp, err := expandBigqueryDataTransferConfigDestinationDatasetId(d.Get("destination_dataset_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("destination_dataset_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(destinationDatasetIdProp)) && (ok || !reflect.DeepEqual(v, destinationDatasetIdProp)) {
		obj["destinationDatasetId"] = destinationDatasetIdProp
	}
	dataSourceIdProp, err := expandBigqueryDataTransferConfigDataSourceId(d.Get("data_source_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("data_source_id"); !tpgresource.IsEmptyValue(reflect.ValueOf(dataSourceIdProp)) && (ok || !reflect.DeepEqual(v, dataSourceIdProp)) {
		obj["dataSourceId"] = dataSourceIdProp
	}
	scheduleProp, err := expandBigqueryDataTransferConfigSchedule(d.Get("schedule"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("schedule"); !tpgresource.IsEmptyValue(reflect.ValueOf(scheduleProp)) && (ok || !reflect.DeepEqual(v, scheduleProp)) {
		obj["schedule"] = scheduleProp
	}
	scheduleOptionsProp, err := expandBigqueryDataTransferConfigScheduleOptions(d.Get("schedule_options"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("schedule_options"); !tpgresource.IsEmptyValue(reflect.ValueOf(scheduleOptionsProp)) && (ok || !reflect.DeepEqual(v, scheduleOptionsProp)) {
		obj["scheduleOptions"] = scheduleOptionsProp
	}
	emailPreferencesProp, err := expandBigqueryDataTransferConfigEmailPreferences(d.Get("email_preferences"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("email_preferences"); !tpgresource.IsEmptyValue(reflect.ValueOf(emailPreferencesProp)) && (ok || !reflect.DeepEqual(v, emailPreferencesProp)) {
		obj["emailPreferences"] = emailPreferencesProp
	}
	notificationPubsubTopicProp, err := expandBigqueryDataTransferConfigNotificationPubsubTopic(d.Get("notification_pubsub_topic"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("notification_pubsub_topic"); !tpgresource.IsEmptyValue(reflect.ValueOf(notificationPubsubTopicProp)) && (ok || !reflect.DeepEqual(v, notificationPubsubTopicProp)) {
		obj["notificationPubsubTopic"] = notificationPubsubTopicProp
	}
	dataRefreshWindowDaysProp, err := expandBigqueryDataTransferConfigDataRefreshWindowDays(d.Get("data_refresh_window_days"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("data_refresh_window_days"); !tpgresource.IsEmptyValue(reflect.ValueOf(dataRefreshWindowDaysProp)) && (ok || !reflect.DeepEqual(v, dataRefreshWindowDaysProp)) {
		obj["dataRefreshWindowDays"] = dataRefreshWindowDaysProp
	}
	encryptionConfigurationProp, err := expandBigqueryDataTransferConfigEncryptionConfiguration(d.Get("encryption_configuration"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("encryption_configuration"); !tpgresource.IsEmptyValue(reflect.ValueOf(encryptionConfigurationProp)) && (ok || !reflect.DeepEqual(v, encryptionConfigurationProp)) {
		obj["encryptionConfiguration"] = encryptionConfigurationProp
	}
	disabledProp, err := expandBigqueryDataTransferConfigDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	paramsProp, err := expandBigqueryDataTransferConfigParams(d.Get("params"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("params"); !tpgresource.IsEmptyValue(reflect.ValueOf(paramsProp)) && (ok || !reflect.DeepEqual(v, paramsProp)) {
		obj["params"] = paramsProp
	}

	return resourceBigqueryDataTransferConfigEncoder(d, config, obj)
}

func resourceBigqueryDataTransferConfigEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	paramMap, ok := obj["params"]
	if !ok {
		paramMap = make(map[string]string)
	}

	params := map[string]interface{}{}

	for k, v := range paramMap.(map[string]string) {
		var value interface{}
		if err := json.Unmarshal([]byte(v), &value); err != nil {
			// If the value is a string, don't convert it to anything.
			params[k] = v
		} else {
			switch value.(type) {
			case float64:
				// If the value is a number, keep the string representation.
				params[k] = v
			default:
				// If the value is another JSON type, keep the unmarshalled type as is.
				params[k] = value
			}
		}
	}

	for _, sp := range sensitiveParams {
		if auth, _ := d.GetOkExists("sensitive_params.0." + sp); auth != "" {
			params[sp] = auth.(string)
		}
	}

	obj["params"] = params

	return obj, nil
}

func expandBigqueryDataTransferConfigDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigDestinationDatasetId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigDataSourceId(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigSchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigScheduleOptions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDisableAutoScheduling, err := expandBigqueryDataTransferConfigScheduleOptionsDisableAutoScheduling(original["disable_auto_scheduling"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisableAutoScheduling); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["disableAutoScheduling"] = transformedDisableAutoScheduling
	}

	transformedStartTime, err := expandBigqueryDataTransferConfigScheduleOptionsStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	transformedEndTime, err := expandBigqueryDataTransferConfigScheduleOptionsEndTime(original["end_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEndTime); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["endTime"] = transformedEndTime
	}

	return transformed, nil
}

func expandBigqueryDataTransferConfigScheduleOptionsDisableAutoScheduling(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigScheduleOptionsStartTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigScheduleOptionsEndTime(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigEmailPreferences(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnableFailureEmail, err := expandBigqueryDataTransferConfigEmailPreferencesEnableFailureEmail(original["enable_failure_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableFailureEmail); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableFailureEmail"] = transformedEnableFailureEmail
	}

	return transformed, nil
}

func expandBigqueryDataTransferConfigEmailPreferencesEnableFailureEmail(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigNotificationPubsubTopic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigDataRefreshWindowDays(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigEncryptionConfiguration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandBigqueryDataTransferConfigEncryptionConfigurationKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandBigqueryDataTransferConfigEncryptionConfigurationKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigParams(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
