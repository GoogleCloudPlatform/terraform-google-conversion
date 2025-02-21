// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/securitycenter/Source.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter_iam.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package securitycenter

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const SecurityCenterSourceIAMAssetType string = "securitycenter.googleapis.com/Source"

func ResourceConverterSecurityCenterSourceIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         SecurityCenterSourceIAMAssetType,
		Convert:           GetSecurityCenterSourceIamPolicyCaiObject,
		MergeCreateUpdate: MergeSecurityCenterSourceIamPolicy,
	}
}

func ResourceConverterSecurityCenterSourceIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         SecurityCenterSourceIAMAssetType,
		Convert:           GetSecurityCenterSourceIamBindingCaiObject,
		FetchFullResource: FetchSecurityCenterSourceIamPolicy,
		MergeCreateUpdate: MergeSecurityCenterSourceIamBinding,
		MergeDelete:       MergeSecurityCenterSourceIamBindingDelete,
	}
}

func ResourceConverterSecurityCenterSourceIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         SecurityCenterSourceIAMAssetType,
		Convert:           GetSecurityCenterSourceIamMemberCaiObject,
		FetchFullResource: FetchSecurityCenterSourceIamPolicy,
		MergeCreateUpdate: MergeSecurityCenterSourceIamMember,
		MergeDelete:       MergeSecurityCenterSourceIamMemberDelete,
	}
}

func GetSecurityCenterSourceIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newSecurityCenterSourceIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetSecurityCenterSourceIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newSecurityCenterSourceIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetSecurityCenterSourceIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newSecurityCenterSourceIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeSecurityCenterSourceIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeSecurityCenterSourceIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeSecurityCenterSourceIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeSecurityCenterSourceIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeSecurityCenterSourceIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newSecurityCenterSourceIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//securitycenter.googleapis.com/organizations/{{organization}}/sources/{{source}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: SecurityCenterSourceIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchSecurityCenterSourceIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("organization"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("source"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		SecurityCenterSourceIamUpdaterProducer,
		d,
		config,
		"//securitycenter.googleapis.com/organizations/{{organization}}/sources/{{source}}",
		SecurityCenterSourceIAMAssetType,
	)
}
