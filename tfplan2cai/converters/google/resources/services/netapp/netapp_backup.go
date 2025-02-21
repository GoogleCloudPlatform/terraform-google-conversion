// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/netapp/Backup.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package netapp

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const NetappBackupAssetType string = "netapp.googleapis.com/Backup"

func ResourceConverterNetappBackup() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: NetappBackupAssetType,
		Convert:   GetNetappBackupCaiObject,
	}
}

func GetNetappBackupCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//netapp.googleapis.com/projects/{{project}}/locations/{{location}}/backupVaults/{{vault_name}}/backups/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetNetappBackupApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: NetappBackupAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/netapp/v1beta1/rest",
				DiscoveryName:        "Backup",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetNetappBackupApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandNetappBackupDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	sourceVolumeProp, err := expandNetappBackupSourceVolume(d.Get("source_volume"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_volume"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceVolumeProp)) && (ok || !reflect.DeepEqual(v, sourceVolumeProp)) {
		obj["sourceVolume"] = sourceVolumeProp
	}
	sourceSnapshotProp, err := expandNetappBackupSourceSnapshot(d.Get("source_snapshot"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("source_snapshot"); !tpgresource.IsEmptyValue(reflect.ValueOf(sourceSnapshotProp)) && (ok || !reflect.DeepEqual(v, sourceSnapshotProp)) {
		obj["sourceSnapshot"] = sourceSnapshotProp
	}
	labelsProp, err := expandNetappBackupEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandNetappBackupDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappBackupSourceVolume(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappBackupSourceSnapshot(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandNetappBackupEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
