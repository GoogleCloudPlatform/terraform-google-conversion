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

package integrationconnectors

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const IntegrationConnectorsEndpointAttachmentAssetType string = "connectors.googleapis.com/EndpointAttachment"

func ResourceConverterIntegrationConnectorsEndpointAttachment() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: IntegrationConnectorsEndpointAttachmentAssetType,
		Convert:   GetIntegrationConnectorsEndpointAttachmentCaiObject,
	}
}

func GetIntegrationConnectorsEndpointAttachmentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//connectors.googleapis.com/projects/{{project}}/locations/{{location}}/endpointAttachments/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetIntegrationConnectorsEndpointAttachmentApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: IntegrationConnectorsEndpointAttachmentAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/connectors/v1/rest",
				DiscoveryName:        "EndpointAttachment",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetIntegrationConnectorsEndpointAttachmentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandIntegrationConnectorsEndpointAttachmentDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	serviceAttachmentProp, err := expandIntegrationConnectorsEndpointAttachmentServiceAttachment(d.Get("service_attachment"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_attachment"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceAttachmentProp)) && (ok || !reflect.DeepEqual(v, serviceAttachmentProp)) {
		obj["serviceAttachment"] = serviceAttachmentProp
	}
	endpointGlobalAccessProp, err := expandIntegrationConnectorsEndpointAttachmentEndpointGlobalAccess(d.Get("endpoint_global_access"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("endpoint_global_access"); !tpgresource.IsEmptyValue(reflect.ValueOf(endpointGlobalAccessProp)) && (ok || !reflect.DeepEqual(v, endpointGlobalAccessProp)) {
		obj["endpointGlobalAccess"] = endpointGlobalAccessProp
	}
	labelsProp, err := expandIntegrationConnectorsEndpointAttachmentEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	return obj, nil
}

func expandIntegrationConnectorsEndpointAttachmentDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIntegrationConnectorsEndpointAttachmentServiceAttachment(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIntegrationConnectorsEndpointAttachmentEndpointGlobalAccess(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIntegrationConnectorsEndpointAttachmentEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
