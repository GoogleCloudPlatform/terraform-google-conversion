// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/apphub/ServiceProjectAttachment.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package apphub

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Suppress all diff for the field Service Project
func ServiceProjectDiffSuppress(_, _, _ string, _ *schema.ResourceData) bool {
	return true
}

const ApphubServiceProjectAttachmentAssetType string = "apphub.googleapis.com/ServiceProjectAttachment"

func ResourceConverterApphubServiceProjectAttachment() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: ApphubServiceProjectAttachmentAssetType,
		Convert:   GetApphubServiceProjectAttachmentCaiObject,
	}
}

func GetApphubServiceProjectAttachmentCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//apphub.googleapis.com/projects/{{project}}/locations/global/serviceProjectAttachments/{{service_project_attachment_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetApphubServiceProjectAttachmentApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: ApphubServiceProjectAttachmentAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/apphub/v1/rest",
				DiscoveryName:        "ServiceProjectAttachment",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetApphubServiceProjectAttachmentApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	serviceProjectProp, err := expandApphubServiceProjectAttachmentServiceProject(d.Get("service_project"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("service_project"); !tpgresource.IsEmptyValue(reflect.ValueOf(serviceProjectProp)) && (ok || !reflect.DeepEqual(v, serviceProjectProp)) {
		obj["serviceProject"] = serviceProjectProp
	}

	return obj, nil
}

func expandApphubServiceProjectAttachmentServiceProject(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {

	service_project := "projects/" + d.Get("service_project_attachment_id").(string)

	return service_project, nil
}
