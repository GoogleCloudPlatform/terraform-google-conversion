// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/networkconnectivity/InternalRange.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package networkconnectivity

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetworkConnectivityInternalRangeAssetType string = "networkconnectivity.googleapis.com/InternalRange"

func ResourceConverterNetworkConnectivityInternalRange() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetworkConnectivityInternalRangeAssetType,
		Convert:   GetNetworkConnectivityInternalRangeCaiObject,
	}
}

func GetNetworkConnectivityInternalRangeCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//networkconnectivity.googleapis.com/projects/{{project}}/locations/global/internalRanges/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetworkConnectivityInternalRangeApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetworkConnectivityInternalRangeAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/networkconnectivity/v1/rest",
				DiscoveryName:        "InternalRange",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetworkConnectivityInternalRangeApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandNetworkConnectivityInternalRangeDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	ipCidrRangeProp, err := expandNetworkConnectivityInternalRangeIpCidrRange(d.Get("ip_cidr_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ip_cidr_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(ipCidrRangeProp)) && (ok || !reflect.DeepEqual(v, ipCidrRangeProp)) {
		obj["ipCidrRange"] = ipCidrRangeProp
	}
	networkProp, err := expandNetworkConnectivityInternalRangeNetwork(d.Get("network"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("network"); !tpgresource.IsEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	usageProp, err := expandNetworkConnectivityInternalRangeUsage(d.Get("usage"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("usage"); !tpgresource.IsEmptyValue(reflect.ValueOf(usageProp)) && (ok || !reflect.DeepEqual(v, usageProp)) {
		obj["usage"] = usageProp
	}
	peeringProp, err := expandNetworkConnectivityInternalRangePeering(d.Get("peering"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("peering"); !tpgresource.IsEmptyValue(reflect.ValueOf(peeringProp)) && (ok || !reflect.DeepEqual(v, peeringProp)) {
		obj["peering"] = peeringProp
	}
	prefixLengthProp, err := expandNetworkConnectivityInternalRangePrefixLength(d.Get("prefix_length"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("prefix_length"); !tpgresource.IsEmptyValue(reflect.ValueOf(prefixLengthProp)) && (ok || !reflect.DeepEqual(v, prefixLengthProp)) {
		obj["prefixLength"] = prefixLengthProp
	}
	targetCidrRangeProp, err := expandNetworkConnectivityInternalRangeTargetCidrRange(d.Get("target_cidr_range"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("target_cidr_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(targetCidrRangeProp)) && (ok || !reflect.DeepEqual(v, targetCidrRangeProp)) {
		obj["targetCidrRange"] = targetCidrRangeProp
	}
	overlapsProp, err := expandNetworkConnectivityInternalRangeOverlaps(d.Get("overlaps"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("overlaps"); !tpgresource.IsEmptyValue(reflect.ValueOf(overlapsProp)) && (ok || !reflect.DeepEqual(v, overlapsProp)) {
		obj["overlaps"] = overlapsProp
	}
	migrationProp, err := expandNetworkConnectivityInternalRangeMigration(d.Get("migration"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("migration"); !tpgresource.IsEmptyValue(reflect.ValueOf(migrationProp)) && (ok || !reflect.DeepEqual(v, migrationProp)) {
		obj["migration"] = migrationProp
	}
	labelsProp, err := expandNetworkConnectivityInternalRangeEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetworkConnectivityInternalRangeDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityInternalRangeIpCidrRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityInternalRangeNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityInternalRangeUsage(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityInternalRangePeering(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityInternalRangePrefixLength(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityInternalRangeTargetCidrRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityInternalRangeOverlaps(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityInternalRangeMigration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedSource, err := expandNetworkConnectivityInternalRangeMigrationSource(original["source"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSource); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["source"] = transformedSource
	}

	transformedTarget, err := expandNetworkConnectivityInternalRangeMigrationTarget(original["target"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTarget); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["target"] = transformedTarget
	}

	return transformed, nil
}

func expandNetworkConnectivityInternalRangeMigrationSource(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityInternalRangeMigrationTarget(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetworkConnectivityInternalRangeEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
