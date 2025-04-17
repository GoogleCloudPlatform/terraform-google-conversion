// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/gkebackup/BackupChannel.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package gkebackup

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const GKEBackupBackupChannelAssetType string = "gkebackup.googleapis.com/BackupChannel"

func ResourceConverterGKEBackupBackupChannel() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: GKEBackupBackupChannelAssetType,
		Convert:   GetGKEBackupBackupChannelCaiObject,
	}
}

func GetGKEBackupBackupChannelCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//gkebackup.googleapis.com/projects/{{project}}/locations/{{location}}/backupChannels/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetGKEBackupBackupChannelApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: GKEBackupBackupChannelAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/gkebackup/v1/rest",
				DiscoveryName:        "BackupChannel",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetGKEBackupBackupChannelApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandGKEBackupBackupChannelName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	destinationProjectProp, err := expandGKEBackupBackupChannelDestinationProject(d.Get("destination_project"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("destination_project"); !tpgresource.IsEmptyValue(reflect.ValueOf(destinationProjectProp)) && (ok || !reflect.DeepEqual(v, destinationProjectProp)) {
		obj["destinationProject"] = destinationProjectProp
	}
	descriptionProp, err := expandGKEBackupBackupChannelDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandGKEBackupBackupChannelEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandGKEBackupBackupChannelName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/backupChannels/{{name}}")
}

func expandGKEBackupBackupChannelDestinationProject(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupChannelDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGKEBackupBackupChannelEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
