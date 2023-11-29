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

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const MonitoringGroupAssetType string = "monitoring.googleapis.com/Group"

func ResourceConverterMonitoringGroup() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: MonitoringGroupAssetType,
		Convert:   GetMonitoringGroupCaiObject,
	}
}

func GetMonitoringGroupCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//monitoring.googleapis.com/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetMonitoringGroupApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: MonitoringGroupAssetType,
			Resource: &cai.AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/monitoring/v3/rest",
				DiscoveryName:        "Group",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetMonitoringGroupApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	parentNameProp, err := expandMonitoringGroupParentName(d.Get("parent_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("parent_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(parentNameProp)) && (ok || !reflect.DeepEqual(v, parentNameProp)) {
		obj["parentName"] = parentNameProp
	}
	isClusterProp, err := expandMonitoringGroupIsCluster(d.Get("is_cluster"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("is_cluster"); !tpgresource.IsEmptyValue(reflect.ValueOf(isClusterProp)) && (ok || !reflect.DeepEqual(v, isClusterProp)) {
		obj["isCluster"] = isClusterProp
	}
	displayNameProp, err := expandMonitoringGroupDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	filterProp, err := expandMonitoringGroupFilter(d.Get("filter"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("filter"); !tpgresource.IsEmptyValue(reflect.ValueOf(filterProp)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}

	return obj, nil
}

func expandMonitoringGroupParentName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringGroupIsCluster(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringGroupDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandMonitoringGroupFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
