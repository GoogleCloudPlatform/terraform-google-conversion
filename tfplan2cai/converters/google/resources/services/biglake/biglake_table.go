// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/biglake/Table.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package biglake

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const BiglakeTableAssetType string = "biglake.googleapis.com/Table"

func ResourceConverterBiglakeTable() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: BiglakeTableAssetType,
		Convert:   GetBiglakeTableCaiObject,
	}
}

func GetBiglakeTableCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//biglake.googleapis.com/{{database}}/tables/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetBiglakeTableApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: BiglakeTableAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/biglake/v1/rest",
				DiscoveryName:        "Table",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetBiglakeTableApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	typeProp, err := expandBiglakeTableType(d.Get("type"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	hiveOptionsProp, err := expandBiglakeTableHiveOptions(d.Get("hive_options"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("hive_options"); !tpgresource.IsEmptyValue(reflect.ValueOf(hiveOptionsProp)) && (ok || !reflect.DeepEqual(v, hiveOptionsProp)) {
		obj["hiveOptions"] = hiveOptionsProp
	}

	return obj, nil
}

func expandBiglakeTableType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBiglakeTableHiveOptions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedParameters, err := expandBiglakeTableHiveOptionsParameters(original["parameters"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedParameters); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["parameters"] = transformedParameters
	}

	transformedTableType, err := expandBiglakeTableHiveOptionsTableType(original["table_type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTableType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["tableType"] = transformedTableType
	}

	transformedStorageDescriptor, err := expandBiglakeTableHiveOptionsStorageDescriptor(original["storage_descriptor"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStorageDescriptor); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["storageDescriptor"] = transformedStorageDescriptor
	}

	return transformed, nil
}

func expandBiglakeTableHiveOptionsParameters(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandBiglakeTableHiveOptionsTableType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBiglakeTableHiveOptionsStorageDescriptor(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedLocationUri, err := expandBiglakeTableHiveOptionsStorageDescriptorLocationUri(original["location_uri"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLocationUri); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["locationUri"] = transformedLocationUri
	}

	transformedInputFormat, err := expandBiglakeTableHiveOptionsStorageDescriptorInputFormat(original["input_format"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedInputFormat); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["inputFormat"] = transformedInputFormat
	}

	transformedOutputFormat, err := expandBiglakeTableHiveOptionsStorageDescriptorOutputFormat(original["output_format"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedOutputFormat); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["outputFormat"] = transformedOutputFormat
	}

	return transformed, nil
}

func expandBiglakeTableHiveOptionsStorageDescriptorLocationUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBiglakeTableHiveOptionsStorageDescriptorInputFormat(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandBiglakeTableHiveOptionsStorageDescriptorOutputFormat(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
