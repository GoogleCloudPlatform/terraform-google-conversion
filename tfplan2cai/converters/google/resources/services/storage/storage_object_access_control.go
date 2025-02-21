// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/storage/ObjectAccessControl.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package storage

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const StorageObjectAccessControlAssetType string = "storage.googleapis.com/ObjectAccessControl"

func ResourceConverterStorageObjectAccessControl() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: StorageObjectAccessControlAssetType,
		Convert:   GetStorageObjectAccessControlCaiObject,
	}
}

func GetStorageObjectAccessControlCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//storage.googleapis.com/b/{{bucket}}/o/{{%object}}/acl/{{entity}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetStorageObjectAccessControlApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: StorageObjectAccessControlAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/storage/v1/rest",
				DiscoveryName:        "ObjectAccessControl",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetStorageObjectAccessControlApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	bucketProp, err := expandStorageObjectAccessControlBucket(d.Get("bucket"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("bucket"); !tpgresource.IsEmptyValue(reflect.ValueOf(bucketProp)) && (ok || !reflect.DeepEqual(v, bucketProp)) {
		obj["bucket"] = bucketProp
	}
	entityProp, err := expandStorageObjectAccessControlEntity(d.Get("entity"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("entity"); !tpgresource.IsEmptyValue(reflect.ValueOf(entityProp)) && (ok || !reflect.DeepEqual(v, entityProp)) {
		obj["entity"] = entityProp
	}
	objectProp, err := expandStorageObjectAccessControlObject(d.Get("object"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("object"); !tpgresource.IsEmptyValue(reflect.ValueOf(objectProp)) && (ok || !reflect.DeepEqual(v, objectProp)) {
		obj["object"] = objectProp
	}
	roleProp, err := expandStorageObjectAccessControlRole(d.Get("role"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("role"); !tpgresource.IsEmptyValue(reflect.ValueOf(roleProp)) && (ok || !reflect.DeepEqual(v, roleProp)) {
		obj["role"] = roleProp
	}

	return obj, nil
}

func expandStorageObjectAccessControlBucket(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageObjectAccessControlEntity(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageObjectAccessControlObject(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandStorageObjectAccessControlRole(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
