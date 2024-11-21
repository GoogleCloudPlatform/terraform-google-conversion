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

package spanner

import (
	"fmt"
	"log"
	"reflect"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func deleteSpannerBackups(d *schema.ResourceData, config *transport_tpg.Config, res map[string]interface{}, userAgent string, billingProject string) error {
	var v interface{}
	var ok bool

	v, ok = res["backups"]
	if !ok || v == nil {
		return nil
	}

	// Iterate over the list and delete each backup.
	for _, itemRaw := range v.([]interface{}) {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		backupName := item["name"].(string)

		log.Printf("[DEBUG] Found backups for resource %q: %#v)", d.Id(), item)

		path := "{{SpannerBasePath}}" + backupName

		url, err := tpgresource.ReplaceVars(d, config, path)
		if err != nil {
			return err
		}

		_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "DELETE",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func resourceSpannerInstanceVirtualUpdate(d *schema.ResourceData, resourceSchema map[string]*schema.Schema) bool {
	// force_destroy is the only virtual field
	if d.HasChange("force_destroy") {
		for field := range resourceSchema {
			if field == "force_destroy" {
				continue
			}
			if d.HasChange(field) {
				return false
			}
		}
		return true
	}
	return false
}

const SpannerInstanceAssetType string = "spanner.googleapis.com/Instance"

func ResourceConverterSpannerInstance() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: SpannerInstanceAssetType,
		Convert:   GetSpannerInstanceCaiObject,
	}
}

func GetSpannerInstanceCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//spanner.googleapis.com/projects/{{project}}/instances/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetSpannerInstanceApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: SpannerInstanceAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/spanner/v1/rest",
				DiscoveryName:        "Instance",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetSpannerInstanceApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandSpannerInstanceName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	configProp, err := expandSpannerInstanceConfig(d.Get("config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("config"); !tpgresource.IsEmptyValue(reflect.ValueOf(configProp)) && (ok || !reflect.DeepEqual(v, configProp)) {
		obj["config"] = configProp
	}
	displayNameProp, err := expandSpannerInstanceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	nodeCountProp, err := expandSpannerInstanceNumNodes(d.Get("num_nodes"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("num_nodes"); !tpgresource.IsEmptyValue(reflect.ValueOf(nodeCountProp)) && (ok || !reflect.DeepEqual(v, nodeCountProp)) {
		obj["nodeCount"] = nodeCountProp
	}
	processingUnitsProp, err := expandSpannerInstanceProcessingUnits(d.Get("processing_units"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("processing_units"); !tpgresource.IsEmptyValue(reflect.ValueOf(processingUnitsProp)) && (ok || !reflect.DeepEqual(v, processingUnitsProp)) {
		obj["processingUnits"] = processingUnitsProp
	}
	autoscalingConfigProp, err := expandSpannerInstanceAutoscalingConfig(d.Get("autoscaling_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("autoscaling_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(autoscalingConfigProp)) && (ok || !reflect.DeepEqual(v, autoscalingConfigProp)) {
		obj["autoscalingConfig"] = autoscalingConfigProp
	}
	editionProp, err := expandSpannerInstanceEdition(d.Get("edition"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("edition"); !tpgresource.IsEmptyValue(reflect.ValueOf(editionProp)) && (ok || !reflect.DeepEqual(v, editionProp)) {
		obj["edition"] = editionProp
	}
	defaultBackupScheduleTypeProp, err := expandSpannerInstanceDefaultBackupScheduleType(d.Get("default_backup_schedule_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("default_backup_schedule_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(defaultBackupScheduleTypeProp)) && (ok || !reflect.DeepEqual(v, defaultBackupScheduleTypeProp)) {
		obj["defaultBackupScheduleType"] = defaultBackupScheduleTypeProp
	}
	labelsProp, err := expandSpannerInstanceEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return resourceSpannerInstanceEncoder(d, config, obj)
}

func resourceSpannerInstanceEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Temp Logic to accommodate autoscaling_config, processing_units and num_nodes
	if obj["processingUnits"] == nil && obj["nodeCount"] == nil && obj["autoscalingConfig"] == nil {
		obj["nodeCount"] = 1
	}
	newObj := make(map[string]interface{})
	newObj["instance"] = obj
	if obj["name"] == nil {
		if err := d.Set("name", id.PrefixedUniqueId("tfgen-spanid-")[:30]); err != nil {
			return nil, fmt.Errorf("Error setting name: %s", err)
		}
		newObj["instanceId"] = d.Get("name").(string)
	} else {
		newObj["instanceId"] = obj["name"]
	}
	delete(obj, "name")
	return newObj, nil
}

func expandSpannerInstanceName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	r := regexp.MustCompile("projects/(.+)/instanceConfigs/(.+)")
	if r.MatchString(v.(string)) {
		return v.(string), nil
	}

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return nil, err
	}

	return fmt.Sprintf("projects/%s/instanceConfigs/%s", project, v.(string)), nil
}

func expandSpannerInstanceDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceNumNodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceProcessingUnits(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceAutoscalingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAutoscalingLimits, err := expandSpannerInstanceAutoscalingConfigAutoscalingLimits(original["autoscaling_limits"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAutoscalingLimits); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["autoscalingLimits"] = transformedAutoscalingLimits
	}

	transformedAutoscalingTargets, err := expandSpannerInstanceAutoscalingConfigAutoscalingTargets(original["autoscaling_targets"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAutoscalingTargets); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["autoscalingTargets"] = transformedAutoscalingTargets
	}

	transformedAsymmetricAutoscalingOptions, err := expandSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptions(original["asymmetric_autoscaling_options"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAsymmetricAutoscalingOptions); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["asymmetricAutoscalingOptions"] = transformedAsymmetricAutoscalingOptions
	}

	return transformed, nil
}

func expandSpannerInstanceAutoscalingConfigAutoscalingLimits(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinProcessingUnits, err := expandSpannerInstanceAutoscalingConfigAutoscalingLimitsMinProcessingUnits(original["min_processing_units"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinProcessingUnits); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minProcessingUnits"] = transformedMinProcessingUnits
	}

	transformedMaxProcessingUnits, err := expandSpannerInstanceAutoscalingConfigAutoscalingLimitsMaxProcessingUnits(original["max_processing_units"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxProcessingUnits); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxProcessingUnits"] = transformedMaxProcessingUnits
	}

	transformedMinNodes, err := expandSpannerInstanceAutoscalingConfigAutoscalingLimitsMinNodes(original["min_nodes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinNodes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minNodes"] = transformedMinNodes
	}

	transformedMaxNodes, err := expandSpannerInstanceAutoscalingConfigAutoscalingLimitsMaxNodes(original["max_nodes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxNodes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxNodes"] = transformedMaxNodes
	}

	return transformed, nil
}

func expandSpannerInstanceAutoscalingConfigAutoscalingLimitsMinProcessingUnits(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceAutoscalingConfigAutoscalingLimitsMaxProcessingUnits(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceAutoscalingConfigAutoscalingLimitsMinNodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceAutoscalingConfigAutoscalingLimitsMaxNodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceAutoscalingConfigAutoscalingTargets(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedHighPriorityCpuUtilizationPercent, err := expandSpannerInstanceAutoscalingConfigAutoscalingTargetsHighPriorityCpuUtilizationPercent(original["high_priority_cpu_utilization_percent"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedHighPriorityCpuUtilizationPercent); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["highPriorityCpuUtilizationPercent"] = transformedHighPriorityCpuUtilizationPercent
	}

	transformedStorageUtilizationPercent, err := expandSpannerInstanceAutoscalingConfigAutoscalingTargetsStorageUtilizationPercent(original["storage_utilization_percent"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStorageUtilizationPercent); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["storageUtilizationPercent"] = transformedStorageUtilizationPercent
	}

	return transformed, nil
}

func expandSpannerInstanceAutoscalingConfigAutoscalingTargetsHighPriorityCpuUtilizationPercent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceAutoscalingConfigAutoscalingTargetsStorageUtilizationPercent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedReplicaSelection, err := expandSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsReplicaSelection(original["replica_selection"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedReplicaSelection); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["replicaSelection"] = transformedReplicaSelection
		}

		transformedOverrides, err := expandSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsOverrides(original["overrides"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedOverrides); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["overrides"] = transformedOverrides
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsReplicaSelection(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLocation, err := expandSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsReplicaSelectionLocation(original["location"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["location"] = transformedLocation
	}

	return transformed, nil
}

func expandSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsReplicaSelectionLocation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsOverrides(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAutoscalingLimits, err := expandSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsOverridesAutoscalingLimits(original["autoscaling_limits"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAutoscalingLimits); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["autoscalingLimits"] = transformedAutoscalingLimits
	}

	return transformed, nil
}

func expandSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsOverridesAutoscalingLimits(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinNodes, err := expandSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsOverridesAutoscalingLimitsMinNodes(original["min_nodes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinNodes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minNodes"] = transformedMinNodes
	}

	transformedMaxNodes, err := expandSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsOverridesAutoscalingLimitsMaxNodes(original["max_nodes"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxNodes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxNodes"] = transformedMaxNodes
	}

	return transformed, nil
}

func expandSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsOverridesAutoscalingLimitsMinNodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsOverridesAutoscalingLimitsMaxNodes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceEdition(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceDefaultBackupScheduleType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSpannerInstanceEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
