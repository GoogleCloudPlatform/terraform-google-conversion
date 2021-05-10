// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

import "reflect"

func GetPubsubSchemaCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//pubsub.googleapis.com/projects/{{project}}/schemas/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetPubsubSchemaApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: "pubsub.googleapis.com/Schema",
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/pubsub/v1/rest",
				DiscoveryName:        "Schema",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetPubsubSchemaApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	typeProp, err := expandPubsubSchemaType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !isEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	definitionProp, err := expandPubsubSchemaDefinition(d.Get("definition"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("definition"); !isEmptyValue(reflect.ValueOf(definitionProp)) && (ok || !reflect.DeepEqual(v, definitionProp)) {
		obj["definition"] = definitionProp
	}
	nameProp, err := expandPubsubSchemaName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}

	return obj, nil
}

func expandPubsubSchemaType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSchemaDefinition(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandPubsubSchemaName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return GetResourceNameFromSelfLink(v.(string)), nil
}
