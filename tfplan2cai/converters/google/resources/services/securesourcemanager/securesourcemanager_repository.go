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

package securesourcemanager

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const SecureSourceManagerRepositoryAssetType string = "securesourcemanager.googleapis.com/Repository"

func ResourceConverterSecureSourceManagerRepository() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: SecureSourceManagerRepositoryAssetType,
		Convert:   GetSecureSourceManagerRepositoryCaiObject,
	}
}

func GetSecureSourceManagerRepositoryCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//securesourcemanager.googleapis.com/projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetSecureSourceManagerRepositoryApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: SecureSourceManagerRepositoryAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/securesourcemanager/v1/rest",
				DiscoveryName:        "Repository",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetSecureSourceManagerRepositoryApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	descriptionProp, err := expandSecureSourceManagerRepositoryDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	instanceProp, err := expandSecureSourceManagerRepositoryInstance(d.Get("instance"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("instance"); !tpgresource.IsEmptyValue(reflect.ValueOf(instanceProp)) && (ok || !reflect.DeepEqual(v, instanceProp)) {
		obj["instance"] = instanceProp
	}
	initialConfigProp, err := expandSecureSourceManagerRepositoryInitialConfig(d.Get("initial_config"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("initial_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(initialConfigProp)) && (ok || !reflect.DeepEqual(v, initialConfigProp)) {
		obj["initialConfig"] = initialConfigProp
	}

	return obj, nil
}

func expandSecureSourceManagerRepositoryDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerRepositoryInstance(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerRepositoryInitialConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDefaultBranch, err := expandSecureSourceManagerRepositoryInitialConfigDefaultBranch(original["default_branch"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDefaultBranch); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["defaultBranch"] = transformedDefaultBranch
	}

	transformedGitignores, err := expandSecureSourceManagerRepositoryInitialConfigGitignores(original["gitignores"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGitignores); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gitignores"] = transformedGitignores
	}

	transformedLicense, err := expandSecureSourceManagerRepositoryInitialConfigLicense(original["license"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedLicense); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["license"] = transformedLicense
	}

	transformedReadme, err := expandSecureSourceManagerRepositoryInitialConfigReadme(original["readme"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedReadme); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["readme"] = transformedReadme
	}

	return transformed, nil
}

func expandSecureSourceManagerRepositoryInitialConfigDefaultBranch(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerRepositoryInitialConfigGitignores(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerRepositoryInitialConfigLicense(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerRepositoryInitialConfigReadme(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
