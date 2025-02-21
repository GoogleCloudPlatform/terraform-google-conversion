// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/securitycenterv2/OrganizationSccBigQueryExport.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package securitycenterv2

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const SecurityCenterV2OrganizationSccBigQueryExportAssetType string = "securitycenter.googleapis.com/OrganizationSccBigQueryExport"

func ResourceConverterSecurityCenterV2OrganizationSccBigQueryExport() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: SecurityCenterV2OrganizationSccBigQueryExportAssetType,
		Convert:   GetSecurityCenterV2OrganizationSccBigQueryExportCaiObject,
	}
}

func GetSecurityCenterV2OrganizationSccBigQueryExportCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//securitycenter.googleapis.com/organizations/{{organization}}/locations/{{location}}/bigQueryExports/{{big_query_export_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetSecurityCenterV2OrganizationSccBigQueryExportApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: SecurityCenterV2OrganizationSccBigQueryExportAssetType,
			Resource: &cai.AssetResource{
				Version:              "v2",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/securitycenter/v2/rest",
				DiscoveryName:        "OrganizationSccBigQueryExport",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetSecurityCenterV2OrganizationSccBigQueryExportApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	nameProp, err := expandSecurityCenterV2OrganizationSccBigQueryExportName(d.Get("name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	descriptionProp, err := expandSecurityCenterV2OrganizationSccBigQueryExportDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	datasetProp, err := expandSecurityCenterV2OrganizationSccBigQueryExportDataset(d.Get("dataset"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dataset"); !tpgresource.IsEmptyValue(reflect.ValueOf(datasetProp)) && (ok || !reflect.DeepEqual(v, datasetProp)) {
		obj["dataset"] = datasetProp
	}
	filterProp, err := expandSecurityCenterV2OrganizationSccBigQueryExportFilter(d.Get("filter"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("filter"); !tpgresource.IsEmptyValue(reflect.ValueOf(filterProp)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}

	return obj, nil
}

func expandSecurityCenterV2OrganizationSccBigQueryExportName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterV2OrganizationSccBigQueryExportDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterV2OrganizationSccBigQueryExportDataset(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterV2OrganizationSccBigQueryExportFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
