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

package securitycenter

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const SecurityCenterOrganizationSccBigQueryExportAssetType string = "securitycenter.googleapis.com/OrganizationSccBigQueryExport"

func ResourceConverterSecurityCenterOrganizationSccBigQueryExport() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: SecurityCenterOrganizationSccBigQueryExportAssetType,
		Convert:   GetSecurityCenterOrganizationSccBigQueryExportCaiObject,
	}
}

func GetSecurityCenterOrganizationSccBigQueryExportCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//securitycenter.googleapis.com/organizations/{{organization}}/bigQueryExports/{{big_query_export_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetSecurityCenterOrganizationSccBigQueryExportApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: SecurityCenterOrganizationSccBigQueryExportAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/securitycenter/v1/rest",
				DiscoveryName:        "OrganizationSccBigQueryExport",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetSecurityCenterOrganizationSccBigQueryExportApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandSecurityCenterOrganizationSccBigQueryExportDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	datasetProp, err := expandSecurityCenterOrganizationSccBigQueryExportDataset(d.Get("dataset"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("dataset"); !tpgresource.IsEmptyValue(reflect.ValueOf(datasetProp)) && (ok || !reflect.DeepEqual(v, datasetProp)) {
		obj["dataset"] = datasetProp
	}
	filterProp, err := expandSecurityCenterOrganizationSccBigQueryExportFilter(d.Get("filter"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("filter"); !tpgresource.IsEmptyValue(reflect.ValueOf(filterProp)) && (ok || !reflect.DeepEqual(v, filterProp)) {
		obj["filter"] = filterProp
	}

	return obj, nil
}

func expandSecurityCenterOrganizationSccBigQueryExportDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterOrganizationSccBigQueryExportDataset(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecurityCenterOrganizationSccBigQueryExportFilter(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
