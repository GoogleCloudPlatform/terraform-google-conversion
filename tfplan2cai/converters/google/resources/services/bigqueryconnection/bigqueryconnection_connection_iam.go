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

package bigqueryconnection

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const BigqueryConnectionConnectionIAMAssetType string = "bigqueryconnection.googleapis.com/Connection"

func ResourceConverterBigqueryConnectionConnectionIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         BigqueryConnectionConnectionIAMAssetType,
		Convert:           GetBigqueryConnectionConnectionIamPolicyCaiObject,
		MergeCreateUpdate: MergeBigqueryConnectionConnectionIamPolicy,
	}
}

func ResourceConverterBigqueryConnectionConnectionIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         BigqueryConnectionConnectionIAMAssetType,
		Convert:           GetBigqueryConnectionConnectionIamBindingCaiObject,
		FetchFullResource: FetchBigqueryConnectionConnectionIamPolicy,
		MergeCreateUpdate: MergeBigqueryConnectionConnectionIamBinding,
		MergeDelete:       MergeBigqueryConnectionConnectionIamBindingDelete,
	}
}

func ResourceConverterBigqueryConnectionConnectionIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         BigqueryConnectionConnectionIAMAssetType,
		Convert:           GetBigqueryConnectionConnectionIamMemberCaiObject,
		FetchFullResource: FetchBigqueryConnectionConnectionIamPolicy,
		MergeCreateUpdate: MergeBigqueryConnectionConnectionIamMember,
		MergeDelete:       MergeBigqueryConnectionConnectionIamMemberDelete,
	}
}

func GetBigqueryConnectionConnectionIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newBigqueryConnectionConnectionIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetBigqueryConnectionConnectionIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newBigqueryConnectionConnectionIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetBigqueryConnectionConnectionIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newBigqueryConnectionConnectionIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeBigqueryConnectionConnectionIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeBigqueryConnectionConnectionIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeBigqueryConnectionConnectionIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeBigqueryConnectionConnectionIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeBigqueryConnectionConnectionIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newBigqueryConnectionConnectionIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//bigqueryconnection.googleapis.com/projects/{{project}}/locations/{{location}}/connections/{{connection_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: BigqueryConnectionConnectionIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchBigqueryConnectionConnectionIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("connection_id"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		BigqueryConnectionConnectionIamUpdaterProducer,
		d,
		config,
		"//bigqueryconnection.googleapis.com/projects/{{project}}/locations/{{location}}/connections/{{connection_id}}",
		BigqueryConnectionConnectionIAMAssetType,
	)
}
