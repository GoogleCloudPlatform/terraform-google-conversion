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

package dataplex

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DataplexEntryTypeAssetType string = "dataplex.googleapis.com/EntryType"

func ResourceConverterDataplexEntryType() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DataplexEntryTypeAssetType,
		Convert:   GetDataplexEntryTypeCaiObject,
	}
}

func GetDataplexEntryTypeCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//dataplex.googleapis.com/projects/{{project}}/locations/{{location}}/entryTypes/{{entry_type_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDataplexEntryTypeApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DataplexEntryTypeAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dataplex/v1/rest",
				DiscoveryName:        "EntryType",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDataplexEntryTypeApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandDataplexEntryTypeDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	displayNameProp, err := expandDataplexEntryTypeDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	typeAliasesProp, err := expandDataplexEntryTypeTypeAliases(d.Get("type_aliases"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("type_aliases"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeAliasesProp)) && (ok || !reflect.DeepEqual(v, typeAliasesProp)) {
		obj["typeAliases"] = typeAliasesProp
	}
	platformProp, err := expandDataplexEntryTypePlatform(d.Get("platform"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("platform"); !tpgresource.IsEmptyValue(reflect.ValueOf(platformProp)) && (ok || !reflect.DeepEqual(v, platformProp)) {
		obj["platform"] = platformProp
	}
	systemProp, err := expandDataplexEntryTypeSystem(d.Get("system"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("system"); !tpgresource.IsEmptyValue(reflect.ValueOf(systemProp)) && (ok || !reflect.DeepEqual(v, systemProp)) {
		obj["system"] = systemProp
	}
	requiredAspectsProp, err := expandDataplexEntryTypeRequiredAspects(d.Get("required_aspects"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("required_aspects"); !tpgresource.IsEmptyValue(reflect.ValueOf(requiredAspectsProp)) && (ok || !reflect.DeepEqual(v, requiredAspectsProp)) {
		obj["requiredAspects"] = requiredAspectsProp
	}
	labelsProp, err := expandDataplexEntryTypeEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandDataplexEntryTypeDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexEntryTypeDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexEntryTypeTypeAliases(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexEntryTypePlatform(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexEntryTypeSystem(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexEntryTypeRequiredAspects(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedType, err := expandDataplexEntryTypeRequiredAspectsType(original["type"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedType); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["type"] = transformedType
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDataplexEntryTypeRequiredAspectsType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDataplexEntryTypeEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
