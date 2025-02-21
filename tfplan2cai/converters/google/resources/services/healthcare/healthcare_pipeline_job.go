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

package healthcare

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const HealthcarePipelineJobAssetType string = "healthcare.googleapis.com/PipelineJob"

func ResourceConverterHealthcarePipelineJob() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: HealthcarePipelineJobAssetType,
		Convert:   GetHealthcarePipelineJobCaiObject,
	}
}

func GetHealthcarePipelineJobCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//healthcare.googleapis.com/{{dataset}}/pipelineJobs/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetHealthcarePipelineJobApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: HealthcarePipelineJobAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/healthcare/v1beta1/rest",
				DiscoveryName:        "PipelineJob",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetHealthcarePipelineJobApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandHealthcarePipelineJobName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	disableLineageProp, err := expandHealthcarePipelineJobDisableLineage(d.Get("disable_lineage"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disable_lineage"); !tpgresource.IsEmptyValue(reflect.ValueOf(disableLineageProp)) && (ok || !reflect.DeepEqual(v, disableLineageProp)) {
		obj["disableLineage"] = disableLineageProp
	}
	mappingPipelineJobProp, err := expandHealthcarePipelineJobMappingPipelineJob(d.Get("mapping_pipeline_job"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("mapping_pipeline_job"); !tpgresource.IsEmptyValue(reflect.ValueOf(mappingPipelineJobProp)) && (ok || !reflect.DeepEqual(v, mappingPipelineJobProp)) {
		obj["mappingPipelineJob"] = mappingPipelineJobProp
	}
	reconciliationPipelineJobProp, err := expandHealthcarePipelineJobReconciliationPipelineJob(d.Get("reconciliation_pipeline_job"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("reconciliation_pipeline_job"); !tpgresource.IsEmptyValue(reflect.ValueOf(reconciliationPipelineJobProp)) && (ok || !reflect.DeepEqual(v, reconciliationPipelineJobProp)) {
		obj["reconciliationPipelineJob"] = reconciliationPipelineJobProp
	}
	backfillPipelineJobProp, err := expandHealthcarePipelineJobBackfillPipelineJob(d.Get("backfill_pipeline_job"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("backfill_pipeline_job"); !tpgresource.IsEmptyValue(reflect.ValueOf(backfillPipelineJobProp)) && (ok || !reflect.DeepEqual(v, backfillPipelineJobProp)) {
		obj["backfillPipelineJob"] = backfillPipelineJobProp
	}
	labelsProp, err := expandHealthcarePipelineJobEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandHealthcarePipelineJobName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcarePipelineJobDisableLineage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcarePipelineJobMappingPipelineJob(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMappingConfig, err := expandHealthcarePipelineJobMappingPipelineJobMappingConfig(original["mapping_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMappingConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mappingConfig"] = transformedMappingConfig
	}

	transformedFhirStreamingSource, err := expandHealthcarePipelineJobMappingPipelineJobFhirStreamingSource(original["fhir_streaming_source"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFhirStreamingSource); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fhirStreamingSource"] = transformedFhirStreamingSource
	}

	transformedFhirStoreDestination, err := expandHealthcarePipelineJobMappingPipelineJobFhirStoreDestination(original["fhir_store_destination"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFhirStoreDestination); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fhirStoreDestination"] = transformedFhirStoreDestination
	}

	transformedReconciliationDestination, err := expandHealthcarePipelineJobMappingPipelineJobReconciliationDestination(original["reconciliation_destination"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedReconciliationDestination); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["reconciliationDestination"] = transformedReconciliationDestination
	}

	return transformed, nil
}

func expandHealthcarePipelineJobMappingPipelineJobMappingConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDescription, err := expandHealthcarePipelineJobMappingPipelineJobMappingConfigDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedWhistleConfigSource, err := expandHealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSource(original["whistle_config_source"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWhistleConfigSource); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["whistleConfigSource"] = transformedWhistleConfigSource
	}

	return transformed, nil
}

func expandHealthcarePipelineJobMappingPipelineJobMappingConfigDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUri, err := expandHealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceUri(original["uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uri"] = transformedUri
	}

	transformedImportUriPrefix, err := expandHealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceImportUriPrefix(original["import_uri_prefix"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImportUriPrefix); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["importUriPrefix"] = transformedImportUriPrefix
	}

	return transformed, nil
}

func expandHealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcarePipelineJobMappingPipelineJobMappingConfigWhistleConfigSourceImportUriPrefix(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcarePipelineJobMappingPipelineJobFhirStreamingSource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFhirStore, err := expandHealthcarePipelineJobMappingPipelineJobFhirStreamingSourceFhirStore(original["fhir_store"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFhirStore); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fhirStore"] = transformedFhirStore
	}

	transformedDescription, err := expandHealthcarePipelineJobMappingPipelineJobFhirStreamingSourceDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	return transformed, nil
}

func expandHealthcarePipelineJobMappingPipelineJobFhirStreamingSourceFhirStore(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcarePipelineJobMappingPipelineJobFhirStreamingSourceDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcarePipelineJobMappingPipelineJobFhirStoreDestination(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcarePipelineJobMappingPipelineJobReconciliationDestination(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcarePipelineJobReconciliationPipelineJob(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMergeConfig, err := expandHealthcarePipelineJobReconciliationPipelineJobMergeConfig(original["merge_config"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMergeConfig); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mergeConfig"] = transformedMergeConfig
	}

	transformedMatchingUriPrefix, err := expandHealthcarePipelineJobReconciliationPipelineJobMatchingUriPrefix(original["matching_uri_prefix"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMatchingUriPrefix); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["matchingUriPrefix"] = transformedMatchingUriPrefix
	}

	transformedFhirStoreDestination, err := expandHealthcarePipelineJobReconciliationPipelineJobFhirStoreDestination(original["fhir_store_destination"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFhirStoreDestination); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fhirStoreDestination"] = transformedFhirStoreDestination
	}

	return transformed, nil
}

func expandHealthcarePipelineJobReconciliationPipelineJobMergeConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDescription, err := expandHealthcarePipelineJobReconciliationPipelineJobMergeConfigDescription(original["description"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["description"] = transformedDescription
	}

	transformedWhistleConfigSource, err := expandHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSource(original["whistle_config_source"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedWhistleConfigSource); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["whistleConfigSource"] = transformedWhistleConfigSource
	}

	return transformed, nil
}

func expandHealthcarePipelineJobReconciliationPipelineJobMergeConfigDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedUri, err := expandHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceUri(original["uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["uri"] = transformedUri
	}

	transformedImportUriPrefix, err := expandHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceImportUriPrefix(original["import_uri_prefix"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImportUriPrefix); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["importUriPrefix"] = transformedImportUriPrefix
	}

	return transformed, nil
}

func expandHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceImportUriPrefix(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcarePipelineJobReconciliationPipelineJobMatchingUriPrefix(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcarePipelineJobReconciliationPipelineJobFhirStoreDestination(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcarePipelineJobBackfillPipelineJob(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMappingPipelineJob, err := expandHealthcarePipelineJobBackfillPipelineJobMappingPipelineJob(original["mapping_pipeline_job"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMappingPipelineJob); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["mappingPipelineJob"] = transformedMappingPipelineJob
	}

	return transformed, nil
}

func expandHealthcarePipelineJobBackfillPipelineJobMappingPipelineJob(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandHealthcarePipelineJobEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
