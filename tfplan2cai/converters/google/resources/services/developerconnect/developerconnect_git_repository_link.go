// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/developerconnect/GitRepositoryLink.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package developerconnect

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const DeveloperConnectGitRepositoryLinkAssetType string = "developerconnect.googleapis.com/GitRepositoryLink"

func ResourceConverterDeveloperConnectGitRepositoryLink() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: DeveloperConnectGitRepositoryLinkAssetType,
		Convert:   GetDeveloperConnectGitRepositoryLinkCaiObject,
	}
}

func GetDeveloperConnectGitRepositoryLinkCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//developerconnect.googleapis.com/projects/{{project}}/locations/{{location}}/connections/{{parent_connection}}/gitRepositoryLinks/{{git_repository_link_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetDeveloperConnectGitRepositoryLinkApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: DeveloperConnectGitRepositoryLinkAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/developerconnect/v1/rest",
				DiscoveryName:        "GitRepositoryLink",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetDeveloperConnectGitRepositoryLinkApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	cloneUriProp, err := expandDeveloperConnectGitRepositoryLinkCloneUri(d.Get("clone_uri"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("clone_uri"); !tpgresource.IsEmptyValue(reflect.ValueOf(cloneUriProp)) && (ok || !reflect.DeepEqual(v, cloneUriProp)) {
		obj["cloneUri"] = cloneUriProp
	}
	etagProp, err := expandDeveloperConnectGitRepositoryLinkEtag(d.Get("etag"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("etag"); !tpgresource.IsEmptyValue(reflect.ValueOf(etagProp)) && (ok || !reflect.DeepEqual(v, etagProp)) {
		obj["etag"] = etagProp
	}
	labelsProp, err := expandDeveloperConnectGitRepositoryLinkEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandDeveloperConnectGitRepositoryLinkEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	return obj, nil
}

func expandDeveloperConnectGitRepositoryLinkCloneUri(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDeveloperConnectGitRepositoryLinkEtag(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDeveloperConnectGitRepositoryLinkEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandDeveloperConnectGitRepositoryLinkEffectiveAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
