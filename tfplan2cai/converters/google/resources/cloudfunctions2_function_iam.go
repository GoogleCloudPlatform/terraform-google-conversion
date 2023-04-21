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

package google

import (
	"fmt"

	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const Cloudfunctions2functionIAMAssetType string = "cloudfunctions.googleapis.com/function"

func resourceConverterCloudfunctions2functionIamPolicy() ResourceConverter {
	return ResourceConverter{
		AssetType:         Cloudfunctions2functionIAMAssetType,
		Convert:           GetCloudfunctions2functionIamPolicyCaiObject,
		MergeCreateUpdate: MergeCloudfunctions2functionIamPolicy,
	}
}

func resourceConverterCloudfunctions2functionIamBinding() ResourceConverter {
	return ResourceConverter{
		AssetType:         Cloudfunctions2functionIAMAssetType,
		Convert:           GetCloudfunctions2functionIamBindingCaiObject,
		FetchFullResource: FetchCloudfunctions2functionIamPolicy,
		MergeCreateUpdate: MergeCloudfunctions2functionIamBinding,
		MergeDelete:       MergeCloudfunctions2functionIamBindingDelete,
	}
}

func resourceConverterCloudfunctions2functionIamMember() ResourceConverter {
	return ResourceConverter{
		AssetType:         Cloudfunctions2functionIAMAssetType,
		Convert:           GetCloudfunctions2functionIamMemberCaiObject,
		FetchFullResource: FetchCloudfunctions2functionIamPolicy,
		MergeCreateUpdate: MergeCloudfunctions2functionIamMember,
		MergeDelete:       MergeCloudfunctions2functionIamMemberDelete,
	}
}

func GetCloudfunctions2functionIamPolicyCaiObject(d TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newCloudfunctions2functionIamAsset(d, config, expandIamPolicyBindings)
}

func GetCloudfunctions2functionIamBindingCaiObject(d TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newCloudfunctions2functionIamAsset(d, config, expandIamRoleBindings)
}

func GetCloudfunctions2functionIamMemberCaiObject(d TerraformResourceData, config *transport_tpg.Config) ([]Asset, error) {
	return newCloudfunctions2functionIamAsset(d, config, expandIamMemberBindings)
}

func MergeCloudfunctions2functionIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeCloudfunctions2functionIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeCloudfunctions2functionIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeCloudfunctions2functionIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeCloudfunctions2functionIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newCloudfunctions2functionIamAsset(
	d TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//cloudfunctions.googleapis.com/projects/{{project}}/locations/{{location}}/functions/{{cloud_function}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: Cloudfunctions2functionIAMAssetType,
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchCloudfunctions2functionIamPolicy(d TerraformResourceData, config *transport_tpg.Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("cloud_function"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		Cloudfunctions2functionIamUpdaterProducer,
		d,
		config,
		"//cloudfunctions.googleapis.com/projects/{{project}}/locations/{{location}}/functions/{{cloud_function}}",
		Cloudfunctions2functionIAMAssetType,
	)
}
