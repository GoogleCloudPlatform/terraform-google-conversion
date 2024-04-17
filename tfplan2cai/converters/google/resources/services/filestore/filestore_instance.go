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

package filestore

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const FilestoreInstanceAssetType string = "file.googleapis.com/Instance"

func ResourceConverterFilestoreInstance() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: FilestoreInstanceAssetType,
		Convert:   GetFilestoreInstanceCaiObject,
	}
}

func GetFilestoreInstanceCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//file.googleapis.com/projects/{{project}}/locations/{{location}}/instances/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetFilestoreInstanceApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: FilestoreInstanceAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1beta1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/file/v1beta1/rest",
				DiscoveryName:        "Instance",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetFilestoreInstanceApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandFilestoreInstanceDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	tierProp, err := expandFilestoreInstanceTier(d.Get("tier"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("tier"); !tpgresource.IsEmptyValue(reflect.ValueOf(tierProp)) && (ok || !reflect.DeepEqual(v, tierProp)) {
		obj["tier"] = tierProp
	}
	protocolProp, err := expandFilestoreInstanceProtocol(d.Get("protocol"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("protocol"); !tpgresource.IsEmptyValue(reflect.ValueOf(protocolProp)) && (ok || !reflect.DeepEqual(v, protocolProp)) {
		obj["protocol"] = protocolProp
	}
	fileSharesProp, err := expandFilestoreInstanceFileShares(d.Get("file_shares"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("file_shares"); !tpgresource.IsEmptyValue(reflect.ValueOf(fileSharesProp)) && (ok || !reflect.DeepEqual(v, fileSharesProp)) {
		obj["fileShares"] = fileSharesProp
	}
	networksProp, err := expandFilestoreInstanceNetworks(d.Get("networks"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("networks"); !tpgresource.IsEmptyValue(reflect.ValueOf(networksProp)) && (ok || !reflect.DeepEqual(v, networksProp)) {
		obj["networks"] = networksProp
	}
	kmsKeyNameProp, err := expandFilestoreInstanceKmsKeyName(d.Get("kms_key_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("kms_key_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(kmsKeyNameProp)) && (ok || !reflect.DeepEqual(v, kmsKeyNameProp)) {
		obj["kmsKeyName"] = kmsKeyNameProp
	}
	labelsProp, err := expandFilestoreInstanceEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandFilestoreInstanceDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceTier(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceProtocol(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceFileShares(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandFilestoreInstanceFileSharesName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedCapacityGb, err := expandFilestoreInstanceFileSharesCapacityGb(original["capacity_gb"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCapacityGb); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["capacityGb"] = transformedCapacityGb
		}

		transformedSourceBackup, err := expandFilestoreInstanceFileSharesSourceBackup(original["source_backup"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSourceBackup); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["sourceBackup"] = transformedSourceBackup
		}

		transformedNfsExportOptions, err := expandFilestoreInstanceFileSharesNfsExportOptions(original["nfs_export_options"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNfsExportOptions); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["nfsExportOptions"] = transformedNfsExportOptions
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandFilestoreInstanceFileSharesName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceFileSharesCapacityGb(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceFileSharesSourceBackup(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceFileSharesNfsExportOptions(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedIpRanges, err := expandFilestoreInstanceFileSharesNfsExportOptionsIpRanges(original["ip_ranges"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpRanges); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipRanges"] = transformedIpRanges
		}

		transformedAccessMode, err := expandFilestoreInstanceFileSharesNfsExportOptionsAccessMode(original["access_mode"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAccessMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["accessMode"] = transformedAccessMode
		}

		transformedSquashMode, err := expandFilestoreInstanceFileSharesNfsExportOptionsSquashMode(original["squash_mode"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSquashMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["squashMode"] = transformedSquashMode
		}

		transformedAnonUid, err := expandFilestoreInstanceFileSharesNfsExportOptionsAnonUid(original["anon_uid"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAnonUid); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["anonUid"] = transformedAnonUid
		}

		transformedAnonGid, err := expandFilestoreInstanceFileSharesNfsExportOptionsAnonGid(original["anon_gid"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedAnonGid); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["anonGid"] = transformedAnonGid
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandFilestoreInstanceFileSharesNfsExportOptionsIpRanges(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceFileSharesNfsExportOptionsAccessMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceFileSharesNfsExportOptionsSquashMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceFileSharesNfsExportOptionsAnonUid(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceFileSharesNfsExportOptionsAnonGid(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceNetworks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNetwork, err := expandFilestoreInstanceNetworksNetwork(original["network"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNetwork); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["network"] = transformedNetwork
		}

		transformedModes, err := expandFilestoreInstanceNetworksModes(original["modes"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedModes); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["modes"] = transformedModes
		}

		transformedReservedIpRange, err := expandFilestoreInstanceNetworksReservedIpRange(original["reserved_ip_range"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedReservedIpRange); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["reservedIpRange"] = transformedReservedIpRange
		}

		transformedIpAddresses, err := expandFilestoreInstanceNetworksIpAddresses(original["ip_addresses"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpAddresses); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipAddresses"] = transformedIpAddresses
		}

		transformedConnectMode, err := expandFilestoreInstanceNetworksConnectMode(original["connect_mode"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedConnectMode); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["connectMode"] = transformedConnectMode
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandFilestoreInstanceNetworksNetwork(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceNetworksModes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceNetworksReservedIpRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceNetworksIpAddresses(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceNetworksConnectMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandFilestoreInstanceEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
