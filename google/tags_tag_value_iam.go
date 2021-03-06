// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
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

import "fmt"

func GetTagsTagValueIamPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newTagsTagValueIamAsset(d, config, expandIamPolicyBindings)
}

func GetTagsTagValueIamBindingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newTagsTagValueIamAsset(d, config, expandIamRoleBindings)
}

func GetTagsTagValueIamMemberCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newTagsTagValueIamAsset(d, config, expandIamMemberBindings)
}

func MergeTagsTagValueIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeTagsTagValueIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeTagsTagValueIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeTagsTagValueIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeTagsTagValueIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newTagsTagValueIamAsset(
	d TerraformResourceData,
	config *Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//tags.googleapis.com/{{tagvalue}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: "tags.googleapis.com/TagValue",
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchTagsTagValueIamPolicy(d TerraformResourceData, config *Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("{{tagvalue}}"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		TagsTagValueIamUpdaterProducer,
		d,
		config,
		"//tags.googleapis.com/{{tagvalue}}",
		"tags.googleapis.com/TagValue",
	)
}
