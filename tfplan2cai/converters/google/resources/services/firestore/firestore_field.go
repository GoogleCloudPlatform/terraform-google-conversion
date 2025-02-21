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

package firestore

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const FirestoreFieldAssetType string = "firestore.googleapis.com/Field"

func ResourceConverterFirestoreField() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: FirestoreFieldAssetType,
		Convert:   GetFirestoreFieldCaiObject,
	}
}

func GetFirestoreFieldCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//firestore.googleapis.com/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetFirestoreFieldApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: FirestoreFieldAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/firestore/v1/rest",
				DiscoveryName:        "Field",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetFirestoreFieldApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	indexConfigProp, err := expandFirestoreFieldIndexConfig(d.Get("index_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("index_config"); ok || !reflect.DeepEqual(v, indexConfigProp) {
		obj["indexConfig"] = indexConfigProp
	}
	ttlConfigProp, err := expandFirestoreFieldTtlConfig(d.Get("ttl_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("ttl_config"); ok || !reflect.DeepEqual(v, ttlConfigProp) {
		obj["ttlConfig"] = ttlConfigProp
	}

	return resourceFirestoreFieldEncoder(d, config, obj)
}

func resourceFirestoreFieldEncoder(d tpgresource.TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {

	// We've added project / database / collection / field as split fields of the name, but
	// the API doesn't expect them.  Make sure we remove them from any requests.

	delete(obj, "project")
	delete(obj, "database")
	delete(obj, "collection")
	delete(obj, "field")
	return obj, nil
}

func expandFirestoreFieldIndexConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	// We drop all output only fields as they are unnecessary.
	if v == nil {
		return nil, nil
	}

	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}

	transformedIndexConfig := make(map[string]interface{})

	// A configured, but empty, index_config block should be sent. This is how a user would remove all indexes.
	if l[0] == nil {
		return transformedIndexConfig, nil
	}

	indexConfig := l[0].(map[string]interface{})

	// For Single field indexes, we put the field configuration on the index to avoid forced nesting.
	// Push all order/arrayConfig down into a single element fields list.
	l = indexConfig["indexes"].(*schema.Set).List()
	transformedIndexes := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})
		transformedField := make(map[string]interface{})

		if val := reflect.ValueOf(original["query_scope"]); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["queryScope"] = original["query_scope"]
		}

		if val := reflect.ValueOf(original["order"]); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformedField["order"] = original["order"]
		}

		if val := reflect.ValueOf(original["array_config"]); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformedField["arrayConfig"] = original["array_config"]
		}
		transformed["fields"] = [1]interface{}{
			transformedField,
		}

		transformedIndexes = append(transformedIndexes, transformed)
	}
	transformedIndexConfig["indexes"] = transformedIndexes
	return transformedIndexConfig, nil
}

/*
 * Expands an empty terraform config into an empty object.
 *
 * Used to differentate a user specifying an empty block versus a null/unset block.
 *
 * This is unique from send_empty_value, which will send an explicit null value
 * for empty configuration blocks.
 */
func expandFirestoreFieldTtlConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	l := v.([]interface{})
	if len(l) == 0 {
		return nil, nil
	}
	// A set, but empty object.
	return struct{}{}, nil
}
