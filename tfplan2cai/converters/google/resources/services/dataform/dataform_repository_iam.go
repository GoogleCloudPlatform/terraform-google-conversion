// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/dataform/Repository.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter_iam.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package dataform

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const DataformRepositoryIAMAssetType string = "dataform.googleapis.com/Repository"

func ResourceConverterDataformRepositoryIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataformRepositoryIAMAssetType,
		Convert:           GetDataformRepositoryIamPolicyCaiObject,
		MergeCreateUpdate: MergeDataformRepositoryIamPolicy,
	}
}

func ResourceConverterDataformRepositoryIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataformRepositoryIAMAssetType,
		Convert:           GetDataformRepositoryIamBindingCaiObject,
		FetchFullResource: FetchDataformRepositoryIamPolicy,
		MergeCreateUpdate: MergeDataformRepositoryIamBinding,
		MergeDelete:       MergeDataformRepositoryIamBindingDelete,
	}
}

func ResourceConverterDataformRepositoryIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataformRepositoryIAMAssetType,
		Convert:           GetDataformRepositoryIamMemberCaiObject,
		FetchFullResource: FetchDataformRepositoryIamPolicy,
		MergeCreateUpdate: MergeDataformRepositoryIamMember,
		MergeDelete:       MergeDataformRepositoryIamMemberDelete,
	}
}

func GetDataformRepositoryIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataformRepositoryIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetDataformRepositoryIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataformRepositoryIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetDataformRepositoryIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataformRepositoryIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeDataformRepositoryIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeDataformRepositoryIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeDataformRepositoryIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeDataformRepositoryIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeDataformRepositoryIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newDataformRepositoryIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//dataform.googleapis.com/projects/{{project}}/locations/{{region}}/repositories/{{repository}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: DataformRepositoryIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchDataformRepositoryIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("region"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("repository"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		DataformRepositoryIamUpdaterProducer,
		d,
		config,
		"//dataform.googleapis.com/projects/{{project}}/locations/{{region}}/repositories/{{repository}}",
		DataformRepositoryIAMAssetType,
	)
}
