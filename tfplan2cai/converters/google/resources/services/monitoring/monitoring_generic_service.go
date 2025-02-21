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

package monitoring

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const MonitoringGenericServiceAssetType string = "monitoring.googleapis.com/GenericService"

func ResourceConverterMonitoringGenericService() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: MonitoringGenericServiceAssetType,
		Convert:   GetMonitoringGenericServiceCaiObject,
	}
}

func GetMonitoringGenericServiceCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//monitoring.googleapis.com/projects/{{project}}/services/{{service_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetMonitoringGenericServiceApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: MonitoringGenericServiceAssetType,
			Resource: &cai.AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/monitoring/v3/rest",
				DiscoveryName:        "GenericService",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetMonitoringGenericServiceApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandMonitoringGenericServiceDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	userLabelsProp, err := expandMonitoringGenericServiceUserLabels(d.Get("user_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("user_labels"); ok || !reflect.DeepEqual(v, userLabelsProp) {
		obj["userLabels"] = userLabelsProp
	}
	basicServiceProp, err := expandMonitoringGenericServiceBasicService(d.Get("basic_service"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("basic_service"); !tpgresource.IsEmptyValue(reflect.ValueOf(basicServiceProp)) && (ok || !reflect.DeepEqual(v, basicServiceProp)) {
		obj["basicService"] = basicServiceProp
	}

	return obj, nil
}

func expandMonitoringGenericServiceDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringGenericServiceUserLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandMonitoringGenericServiceBasicService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedServiceType, err := expandMonitoringGenericServiceBasicServiceServiceType(original["service_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceType"] = transformedServiceType
	}

	transformedServiceLabels, err := expandMonitoringGenericServiceBasicServiceServiceLabels(original["service_labels"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedServiceLabels); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["serviceLabels"] = transformedServiceLabels
	}

	return transformed, nil
}

func expandMonitoringGenericServiceBasicServiceServiceType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringGenericServiceBasicServiceServiceLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
