// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/monitoring/MetricDescriptor.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package monitoring

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const MonitoringMetricDescriptorAssetType string = "monitoring.googleapis.com/MetricDescriptor"

func ResourceConverterMonitoringMetricDescriptor() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: MonitoringMetricDescriptorAssetType,
		Convert:   GetMonitoringMetricDescriptorCaiObject,
	}
}

func GetMonitoringMetricDescriptorCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//monitoring.googleapis.com/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetMonitoringMetricDescriptorApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: MonitoringMetricDescriptorAssetType,
			Resource: &cai.AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/monitoring/v3/rest",
				DiscoveryName:        "MetricDescriptor",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetMonitoringMetricDescriptorApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	typeProp, err := expandMonitoringMetricDescriptorType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	labelsProp, err := expandMonitoringMetricDescriptorLabels(d.Get("labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	metricKindProp, err := expandMonitoringMetricDescriptorMetricKind(d.Get("metric_kind"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("metric_kind"); !tpgresource.IsEmptyValue(reflect.ValueOf(metricKindProp)) && (ok || !reflect.DeepEqual(v, metricKindProp)) {
		obj["metricKind"] = metricKindProp
	}
	valueTypeProp, err := expandMonitoringMetricDescriptorValueType(d.Get("value_type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("value_type"); !tpgresource.IsEmptyValue(reflect.ValueOf(valueTypeProp)) && (ok || !reflect.DeepEqual(v, valueTypeProp)) {
		obj["valueType"] = valueTypeProp
	}
	unitProp, err := expandMonitoringMetricDescriptorUnit(d.Get("unit"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("unit"); !tpgresource.IsEmptyValue(reflect.ValueOf(unitProp)) && (ok || !reflect.DeepEqual(v, unitProp)) {
		obj["unit"] = unitProp
	}
	descriptionProp, err := expandMonitoringMetricDescriptorDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandMonitoringMetricDescriptorDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	metadataProp, err := expandMonitoringMetricDescriptorMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("metadata"); !tpgresource.IsEmptyValue(reflect.ValueOf(metadataProp)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}
	launchStageProp, err := expandMonitoringMetricDescriptorLaunchStage(d.Get("launch_stage"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("launch_stage"); !tpgresource.IsEmptyValue(reflect.ValueOf(launchStageProp)) && (ok || !reflect.DeepEqual(v, launchStageProp)) {
		obj["launchStage"] = launchStageProp
	}

	return obj, nil
}

func expandMonitoringMetricDescriptorType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedKey, err := expandMonitoringMetricDescriptorLabelsKey(original["key"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedKey); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["key"] = transformedKey
		}

		transformedValueType, err := expandMonitoringMetricDescriptorLabelsValueType(original["value_type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValueType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["valueType"] = transformedValueType
		}

		transformedDescription, err := expandMonitoringMetricDescriptorLabelsDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandMonitoringMetricDescriptorLabelsKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorLabelsValueType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorLabelsDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorMetricKind(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorValueType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorUnit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorMetadata(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSamplePeriod, err := expandMonitoringMetricDescriptorMetadataSamplePeriod(original["sample_period"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSamplePeriod); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["samplePeriod"] = transformedSamplePeriod
	}

	transformedIngestDelay, err := expandMonitoringMetricDescriptorMetadataIngestDelay(original["ingest_delay"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedIngestDelay); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["ingestDelay"] = transformedIngestDelay
	}

	return transformed, nil
}

func expandMonitoringMetricDescriptorMetadataSamplePeriod(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorMetadataIngestDelay(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringMetricDescriptorLaunchStage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
