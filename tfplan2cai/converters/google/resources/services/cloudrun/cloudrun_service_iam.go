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

package cloudrun

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const CloudRunServiceIAMAssetType string = "run.googleapis.com/Service"

func ResourceConverterCloudRunServiceIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         CloudRunServiceIAMAssetType,
		Convert:           GetCloudRunServiceIamPolicyCaiObject,
		MergeCreateUpdate: MergeCloudRunServiceIamPolicy,
	}
}

func ResourceConverterCloudRunServiceIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         CloudRunServiceIAMAssetType,
		Convert:           GetCloudRunServiceIamBindingCaiObject,
		FetchFullResource: FetchCloudRunServiceIamPolicy,
		MergeCreateUpdate: MergeCloudRunServiceIamBinding,
		MergeDelete:       MergeCloudRunServiceIamBindingDelete,
	}
}

func ResourceConverterCloudRunServiceIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         CloudRunServiceIAMAssetType,
		Convert:           GetCloudRunServiceIamMemberCaiObject,
		FetchFullResource: FetchCloudRunServiceIamPolicy,
		MergeCreateUpdate: MergeCloudRunServiceIamMember,
		MergeDelete:       MergeCloudRunServiceIamMemberDelete,
	}
}

func GetCloudRunServiceIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newCloudRunServiceIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetCloudRunServiceIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newCloudRunServiceIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetCloudRunServiceIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newCloudRunServiceIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeCloudRunServiceIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeCloudRunServiceIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeCloudRunServiceIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeCloudRunServiceIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeCloudRunServiceIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newCloudRunServiceIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//run.googleapis.com/projects/{{project}}/locations/{{location}}/services/{{service}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: CloudRunServiceIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchCloudRunServiceIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("service"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		CloudRunServiceIamUpdaterProducer,
		d,
		config,
		"//run.googleapis.com/projects/{{project}}/locations/{{location}}/services/{{service}}",
		CloudRunServiceIAMAssetType,
	)
}
