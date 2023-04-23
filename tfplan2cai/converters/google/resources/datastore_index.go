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

package google

import (
	"reflect"

	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

const DatastoreIndexAssetType string = "datastore.googleapis.com/Index"

func resourceConverterDatastoreIndex() ResourceConverter {
	return ResourceConverter{
		AssetType: DatastoreIndexAssetType,
		Convert:   GetDatastoreIndexCaiObject,
	}
}

func GetDatastoreIndexCaiObject(d TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	name, err := assetName(d, config, "//datastore.googleapis.com/projects/{{project}}/indexes/{{index_id}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetDatastoreIndexApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: DatastoreIndexAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/datastore/v1/rest",
				DiscoveryName:        "Index",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetDatastoreIndexApiObject(d TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	kindProp, err := expandDatastoreIndexKind(d.Get("kind"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("kind"); !isEmptyValue(reflect.ValueOf(kindProp)) && (ok || !reflect.DeepEqual(v, kindProp)) {
		obj["kind"] = kindProp
	}
	ancestorProp, err := expandDatastoreIndexAncestor(d.Get("ancestor"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ancestor"); !isEmptyValue(reflect.ValueOf(ancestorProp)) && (ok || !reflect.DeepEqual(v, ancestorProp)) {
		obj["ancestor"] = ancestorProp
	}
	propertiesProp, err := expandDatastoreIndexProperties(d.Get("properties"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("properties"); !isEmptyValue(reflect.ValueOf(propertiesProp)) && (ok || !reflect.DeepEqual(v, propertiesProp)) {
		obj["properties"] = propertiesProp
	}

	return obj, nil
}

func expandDatastoreIndexKind(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatastoreIndexAncestor(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatastoreIndexProperties(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandDatastoreIndexPropertiesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedDirection, err := expandDatastoreIndexPropertiesDirection(original["direction"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDirection); val.IsValid() && !isEmptyValue(val) {
			transformed["direction"] = transformedDirection
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDatastoreIndexPropertiesName(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDatastoreIndexPropertiesDirection(v interface{}, d TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
