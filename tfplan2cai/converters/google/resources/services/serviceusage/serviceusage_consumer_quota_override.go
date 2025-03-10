// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/serviceusage/ConsumerQuotaOverride.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package serviceusage

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const ServiceUsageConsumerQuotaOverrideAssetType string = "serviceusage.googleapis.com/ConsumerQuotaOverride"

func ResourceConverterServiceUsageConsumerQuotaOverride() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ServiceUsageConsumerQuotaOverrideAssetType,
		Convert:   GetServiceUsageConsumerQuotaOverrideCaiObject,
	}
}

func GetServiceUsageConsumerQuotaOverrideCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//serviceusage.googleapis.com/projects/{{project}}/services/{{service}}/consumerQuotaMetrics/{{metric}}/limits/{{limit}}/consumerOverrides/")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetServiceUsageConsumerQuotaOverrideApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ServiceUsageConsumerQuotaOverrideAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/serviceusage/v1beta1/rest",
				DiscoveryName:        "ConsumerQuotaOverride",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetServiceUsageConsumerQuotaOverrideApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	overrideValueProp, err := expandServiceUsageConsumerQuotaOverrideOverrideValue(d.Get("override_value"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("override_value"); !tpgresource.IsEmptyValue(reflect.ValueOf(overrideValueProp)) && (ok || !reflect.DeepEqual(v, overrideValueProp)) {
		obj["overrideValue"] = overrideValueProp
	}
	dimensionsProp, err := expandServiceUsageConsumerQuotaOverrideDimensions(d.Get("dimensions"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dimensions"); !tpgresource.IsEmptyValue(reflect.ValueOf(dimensionsProp)) && (ok || !reflect.DeepEqual(v, dimensionsProp)) {
		obj["dimensions"] = dimensionsProp
	}

	return obj, nil
}

func expandServiceUsageConsumerQuotaOverrideOverrideValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandServiceUsageConsumerQuotaOverrideDimensions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
