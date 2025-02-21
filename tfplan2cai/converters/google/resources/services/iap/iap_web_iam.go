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

package iap

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const IapWebIAMAssetType string = "iap.googleapis.com/Web"

func ResourceConverterIapWebIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         IapWebIAMAssetType,
		Convert:           GetIapWebIamPolicyCaiObject,
		MergeCreateUpdate: MergeIapWebIamPolicy,
	}
}

func ResourceConverterIapWebIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         IapWebIAMAssetType,
		Convert:           GetIapWebIamBindingCaiObject,
		FetchFullResource: FetchIapWebIamPolicy,
		MergeCreateUpdate: MergeIapWebIamBinding,
		MergeDelete:       MergeIapWebIamBindingDelete,
	}
}

func ResourceConverterIapWebIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         IapWebIAMAssetType,
		Convert:           GetIapWebIamMemberCaiObject,
		FetchFullResource: FetchIapWebIamPolicy,
		MergeCreateUpdate: MergeIapWebIamMember,
		MergeDelete:       MergeIapWebIamMemberDelete,
	}
}

func GetIapWebIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newIapWebIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetIapWebIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newIapWebIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetIapWebIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newIapWebIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeIapWebIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeIapWebIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeIapWebIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeIapWebIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeIapWebIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newIapWebIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//iap.googleapis.com/projects/{{project}}/iap_web")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: IapWebIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchIapWebIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value

	return cai.FetchIamPolicy(
		IapWebIamUpdaterProducer,
		d,
		config,
		"//iap.googleapis.com/projects/{{project}}/iap_web",
		IapWebIAMAssetType,
	)
}
