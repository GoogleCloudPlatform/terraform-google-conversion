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

package dataplex

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DataplexDatascanAssetType string = "dataplex.googleapis.com/Datascan"

func ResourceConverterDataplexDatascan() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DataplexDatascanAssetType,
		Convert:   GetDataplexDatascanCaiObject,
	}
}

func GetDataplexDatascanCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//dataplex.googleapis.com/projects/{{project}}/locations/{{location}}/dataScans/{{data_scan_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDataplexDatascanApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DataplexDatascanAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dataplex/v1/rest",
				DiscoveryName:        "Datascan",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDataplexDatascanApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandDataplexDatascanDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandDataplexDatascanDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	dataProp, err := expandDataplexDatascanData(d.Get("data"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("data"); !tpgresource.IsEmptyValue(reflect.ValueOf(dataProp)) && (ok || !reflect.DeepEqual(v, dataProp)) {
		obj["data"] = dataProp
	}
	executionSpecProp, err := expandDataplexDatascanExecutionSpec(d.Get("execution_spec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("execution_spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(executionSpecProp)) && (ok || !reflect.DeepEqual(v, executionSpecProp)) {
		obj["executionSpec"] = executionSpecProp
	}
	dataQualitySpecProp, err := expandDataplexDatascanDataQualitySpec(d.Get("data_quality_spec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("data_quality_spec"); !tpgresource.IsEmptyValue(reflect.ValueOf(dataQualitySpecProp)) && (ok || !reflect.DeepEqual(v, dataQualitySpecProp)) {
		obj["dataQualitySpec"] = dataQualitySpecProp
	}
	dataProfileSpecProp, err := expandDataplexDatascanDataProfileSpec(d.Get("data_profile_spec"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("data_profile_spec"); ok || !reflect.DeepEqual(v, dataProfileSpecProp) {
		obj["dataProfileSpec"] = dataProfileSpecProp
	}
	labelsProp, err := expandDataplexDatascanEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandDataplexDatascanDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanData(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEntity, err := expandDataplexDatascanDataEntity(original["entity"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEntity); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["entity"] = transformedEntity
	}

	transformedResource, err := expandDataplexDatascanDataResource(original["resource"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResource); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["resource"] = transformedResource
	}

	return transformed, nil
}

func expandDataplexDatascanDataEntity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataResource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanExecutionSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTrigger, err := expandDataplexDatascanExecutionSpecTrigger(original["trigger"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTrigger); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["trigger"] = transformedTrigger
	}

	transformedField, err := expandDataplexDatascanExecutionSpecField(original["field"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedField); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["field"] = transformedField
	}

	return transformed, nil
}

func expandDataplexDatascanExecutionSpecTrigger(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedOnDemand, err := expandDataplexDatascanExecutionSpecTriggerOnDemand(original["on_demand"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["onDemand"] = transformedOnDemand
	}

	transformedSchedule, err := expandDataplexDatascanExecutionSpecTriggerSchedule(original["schedule"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSchedule); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["schedule"] = transformedSchedule
	}

	return transformed, nil
}

func expandDataplexDatascanExecutionSpecTriggerOnDemand(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	transformed := make(map[string]interface{})

	return transformed, nil
}

func expandDataplexDatascanExecutionSpecTriggerSchedule(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCron, err := expandDataplexDatascanExecutionSpecTriggerScheduleCron(original["cron"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCron); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["cron"] = transformedCron
	}

	return transformed, nil
}

func expandDataplexDatascanExecutionSpecTriggerScheduleCron(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanExecutionSpecField(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSamplingPercent, err := expandDataplexDatascanDataQualitySpecSamplingPercent(original["sampling_percent"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSamplingPercent); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["samplingPercent"] = transformedSamplingPercent
	}

	transformedRowFilter, err := expandDataplexDatascanDataQualitySpecRowFilter(original["row_filter"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRowFilter); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rowFilter"] = transformedRowFilter
	}

	transformedPostScanActions, err := expandDataplexDatascanDataQualitySpecPostScanActions(original["post_scan_actions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPostScanActions); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["postScanActions"] = transformedPostScanActions
	}

	transformedRules, err := expandDataplexDatascanDataQualitySpecRules(original["rules"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRules); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rules"] = transformedRules
	}

	return transformed, nil
}

func expandDataplexDatascanDataQualitySpecSamplingPercent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecRowFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecPostScanActions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBigqueryExport, err := expandDataplexDatascanDataQualitySpecPostScanActionsBigqueryExport(original["bigquery_export"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBigqueryExport); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["bigqueryExport"] = transformedBigqueryExport
	}

	return transformed, nil
}

func expandDataplexDatascanDataQualitySpecPostScanActionsBigqueryExport(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResultsTable, err := expandDataplexDatascanDataQualitySpecPostScanActionsBigqueryExportResultsTable(original["results_table"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResultsTable); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["resultsTable"] = transformedResultsTable
	}

	return transformed, nil
}

func expandDataplexDatascanDataQualitySpecPostScanActionsBigqueryExportResultsTable(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecRules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedColumn, err := expandDataplexDatascanDataQualitySpecRulesColumn(original["column"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedColumn); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["column"] = transformedColumn
		}

		transformedIgnoreNull, err := expandDataplexDatascanDataQualitySpecRulesIgnoreNull(original["ignore_null"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIgnoreNull); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ignoreNull"] = transformedIgnoreNull
		}

		transformedDimension, err := expandDataplexDatascanDataQualitySpecRulesDimension(original["dimension"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDimension); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["dimension"] = transformedDimension
		}

		transformedThreshold, err := expandDataplexDatascanDataQualitySpecRulesThreshold(original["threshold"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedThreshold); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["threshold"] = transformedThreshold
		}

		transformedName, err := expandDataplexDatascanDataQualitySpecRulesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedDescription, err := expandDataplexDatascanDataQualitySpecRulesDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedRangeExpectation, err := expandDataplexDatascanDataQualitySpecRulesRangeExpectation(original["range_expectation"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRangeExpectation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["rangeExpectation"] = transformedRangeExpectation
		}

		transformedNonNullExpectation, err := expandDataplexDatascanDataQualitySpecRulesNonNullExpectation(original["non_null_expectation"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["nonNullExpectation"] = transformedNonNullExpectation
		}

		transformedSetExpectation, err := expandDataplexDatascanDataQualitySpecRulesSetExpectation(original["set_expectation"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSetExpectation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["setExpectation"] = transformedSetExpectation
		}

		transformedRegexExpectation, err := expandDataplexDatascanDataQualitySpecRulesRegexExpectation(original["regex_expectation"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRegexExpectation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["regexExpectation"] = transformedRegexExpectation
		}

		transformedUniquenessExpectation, err := expandDataplexDatascanDataQualitySpecRulesUniquenessExpectation(original["uniqueness_expectation"], d, config)
		if err != nil {
			return nil, err
		} else {
			transformed["uniquenessExpectation"] = transformedUniquenessExpectation
		}

		transformedStatisticRangeExpectation, err := expandDataplexDatascanDataQualitySpecRulesStatisticRangeExpectation(original["statistic_range_expectation"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedStatisticRangeExpectation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["statisticRangeExpectation"] = transformedStatisticRangeExpectation
		}

		transformedRowConditionExpectation, err := expandDataplexDatascanDataQualitySpecRulesRowConditionExpectation(original["row_condition_expectation"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedRowConditionExpectation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["rowConditionExpectation"] = transformedRowConditionExpectation
		}

		transformedTableConditionExpectation, err := expandDataplexDatascanDataQualitySpecRulesTableConditionExpectation(original["table_condition_expectation"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedTableConditionExpectation); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["tableConditionExpectation"] = transformedTableConditionExpectation
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataplexDatascanDataQualitySpecRulesColumn(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecRulesIgnoreNull(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecRulesDimension(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecRulesThreshold(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecRulesName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecRulesDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecRulesRangeExpectation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedMinValue, err := expandDataplexDatascanDataQualitySpecRulesRangeExpectationMinValue(original["min_value"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minValue"] = transformedMinValue
	}

	transformedMaxValue, err := expandDataplexDatascanDataQualitySpecRulesRangeExpectationMaxValue(original["max_value"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxValue"] = transformedMaxValue
	}

	transformedStrictMinEnabled, err := expandDataplexDatascanDataQualitySpecRulesRangeExpectationStrictMinEnabled(original["strict_min_enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStrictMinEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["strictMinEnabled"] = transformedStrictMinEnabled
	}

	transformedStrictMaxEnabled, err := expandDataplexDatascanDataQualitySpecRulesRangeExpectationStrictMaxEnabled(original["strict_max_enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStrictMaxEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["strictMaxEnabled"] = transformedStrictMaxEnabled
	}

	return transformed, nil
}

func expandDataplexDatascanDataQualitySpecRulesRangeExpectationMinValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecRulesRangeExpectationMaxValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecRulesRangeExpectationStrictMinEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecRulesRangeExpectationStrictMaxEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecRulesNonNullExpectation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	transformed := make(map[string]interface{})

	return transformed, nil
}

func expandDataplexDatascanDataQualitySpecRulesSetExpectation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedValues, err := expandDataplexDatascanDataQualitySpecRulesSetExpectationValues(original["values"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedValues); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["values"] = transformedValues
	}

	return transformed, nil
}

func expandDataplexDatascanDataQualitySpecRulesSetExpectationValues(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecRulesRegexExpectation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRegex, err := expandDataplexDatascanDataQualitySpecRulesRegexExpectationRegex(original["regex"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRegex); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["regex"] = transformedRegex
	}

	return transformed, nil
}

func expandDataplexDatascanDataQualitySpecRulesRegexExpectationRegex(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecRulesUniquenessExpectation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	if l[0] == nil {
		transformed := make(map[string]interface{})
		return transformed, nil
	}
	transformed := make(map[string]interface{})

	return transformed, nil
}

func expandDataplexDatascanDataQualitySpecRulesStatisticRangeExpectation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedStatistic, err := expandDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationStatistic(original["statistic"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStatistic); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["statistic"] = transformedStatistic
	}

	transformedMinValue, err := expandDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationMinValue(original["min_value"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMinValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["minValue"] = transformedMinValue
	}

	transformedMaxValue, err := expandDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationMaxValue(original["max_value"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedMaxValue); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["maxValue"] = transformedMaxValue
	}

	transformedStrictMinEnabled, err := expandDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationStrictMinEnabled(original["strict_min_enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStrictMinEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["strictMinEnabled"] = transformedStrictMinEnabled
	}

	transformedStrictMaxEnabled, err := expandDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationStrictMaxEnabled(original["strict_max_enabled"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStrictMaxEnabled); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["strictMaxEnabled"] = transformedStrictMaxEnabled
	}

	return transformed, nil
}

func expandDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationStatistic(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationMinValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationMaxValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationStrictMinEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecRulesStatisticRangeExpectationStrictMaxEnabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecRulesRowConditionExpectation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSqlExpression, err := expandDataplexDatascanDataQualitySpecRulesRowConditionExpectationSqlExpression(original["sql_expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSqlExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sqlExpression"] = transformedSqlExpression
	}

	return transformed, nil
}

func expandDataplexDatascanDataQualitySpecRulesRowConditionExpectationSqlExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataQualitySpecRulesTableConditionExpectation(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSqlExpression, err := expandDataplexDatascanDataQualitySpecRulesTableConditionExpectationSqlExpression(original["sql_expression"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSqlExpression); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["sqlExpression"] = transformedSqlExpression
	}

	return transformed, nil
}

func expandDataplexDatascanDataQualitySpecRulesTableConditionExpectationSqlExpression(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataProfileSpec(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
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

	transformedSamplingPercent, err := expandDataplexDatascanDataProfileSpecSamplingPercent(original["sampling_percent"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSamplingPercent); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["samplingPercent"] = transformedSamplingPercent
	}

	transformedRowFilter, err := expandDataplexDatascanDataProfileSpecRowFilter(original["row_filter"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRowFilter); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rowFilter"] = transformedRowFilter
	}

	transformedPostScanActions, err := expandDataplexDatascanDataProfileSpecPostScanActions(original["post_scan_actions"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPostScanActions); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["postScanActions"] = transformedPostScanActions
	}

	transformedIncludeFields, err := expandDataplexDatascanDataProfileSpecIncludeFields(original["include_fields"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIncludeFields); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["includeFields"] = transformedIncludeFields
	}

	transformedExcludeFields, err := expandDataplexDatascanDataProfileSpecExcludeFields(original["exclude_fields"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedExcludeFields); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["excludeFields"] = transformedExcludeFields
	}

	return transformed, nil
}

func expandDataplexDatascanDataProfileSpecSamplingPercent(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataProfileSpecRowFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataProfileSpecPostScanActions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedBigqueryExport, err := expandDataplexDatascanDataProfileSpecPostScanActionsBigqueryExport(original["bigquery_export"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedBigqueryExport); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["bigqueryExport"] = transformedBigqueryExport
	}

	return transformed, nil
}

func expandDataplexDatascanDataProfileSpecPostScanActionsBigqueryExport(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedResultsTable, err := expandDataplexDatascanDataProfileSpecPostScanActionsBigqueryExportResultsTable(original["results_table"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedResultsTable); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["resultsTable"] = transformedResultsTable
	}

	return transformed, nil
}

func expandDataplexDatascanDataProfileSpecPostScanActionsBigqueryExportResultsTable(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataProfileSpecIncludeFields(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFieldNames, err := expandDataplexDatascanDataProfileSpecIncludeFieldsFieldNames(original["field_names"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFieldNames); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fieldNames"] = transformedFieldNames
	}

	return transformed, nil
}

func expandDataplexDatascanDataProfileSpecIncludeFieldsFieldNames(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanDataProfileSpecExcludeFields(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedFieldNames, err := expandDataplexDatascanDataProfileSpecExcludeFieldsFieldNames(original["field_names"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedFieldNames); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["fieldNames"] = transformedFieldNames
	}

	return transformed, nil
}

func expandDataplexDatascanDataProfileSpecExcludeFieldsFieldNames(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexDatascanEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
