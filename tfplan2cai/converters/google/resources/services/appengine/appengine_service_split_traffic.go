// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/appengine/ServiceSplitTraffic.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package appengine

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const AppEngineServiceSplitTrafficAssetType string = "appengine.googleapis.com/ServiceSplitTraffic"

func ResourceConverterAppEngineServiceSplitTraffic() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: AppEngineServiceSplitTrafficAssetType,
		Convert:   GetAppEngineServiceSplitTrafficCaiObject,
	}
}

func GetAppEngineServiceSplitTrafficCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//appengine.googleapis.com/apps/{{project}}/services/{{service}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetAppEngineServiceSplitTrafficApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: AppEngineServiceSplitTrafficAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/appengine/v1beta/rest",
				DiscoveryName:        "ServiceSplitTraffic",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetAppEngineServiceSplitTrafficApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	idProp, err := expandAppEngineServiceSplitTrafficService(d.Get("service"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service"); !tpgresource.IsEmptyValue(reflect.ValueOf(idProp)) && (ok || !reflect.DeepEqual(v, idProp)) {
		obj["id"] = idProp
	}
	splitProp, err := expandAppEngineServiceSplitTrafficSplit(d.Get("split"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("split"); !tpgresource.IsEmptyValue(reflect.ValueOf(splitProp)) && (ok || !reflect.DeepEqual(v, splitProp)) {
		obj["split"] = splitProp
	}

	return obj, nil
}

func expandAppEngineServiceSplitTrafficService(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineServiceSplitTrafficSplit(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedShardBy, err := expandAppEngineServiceSplitTrafficSplitShardBy(original["shard_by"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedShardBy); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["shardBy"] = transformedShardBy
	}

	transformedAllocations, err := expandAppEngineServiceSplitTrafficSplitAllocations(original["allocations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAllocations); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["allocations"] = transformedAllocations
	}

	return transformed, nil
}

func expandAppEngineServiceSplitTrafficSplitShardBy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAppEngineServiceSplitTrafficSplitAllocations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
