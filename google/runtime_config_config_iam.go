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

func GetRuntimeConfigConfigIamPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newRuntimeConfigConfigIamAsset(d, config, expandIamPolicyBindings)
}

func GetRuntimeConfigConfigIamBindingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newRuntimeConfigConfigIamAsset(d, config, expandIamRoleBindings)
}

func GetRuntimeConfigConfigIamMemberCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newRuntimeConfigConfigIamAsset(d, config, expandIamMemberBindings)
}

func MergeRuntimeConfigConfigIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeRuntimeConfigConfigIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeRuntimeConfigConfigIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeRuntimeConfigConfigIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeRuntimeConfigConfigIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newRuntimeConfigConfigIamAsset(
	d TerraformResourceData,
	config *Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//runtimeconfig.googleapis.com/{{config}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: "runtimeconfig.googleapis.com/Config",
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchRuntimeConfigConfigIamPolicy(d TerraformResourceData, config *Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("{{config}}"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		RuntimeConfigConfigIamUpdaterProducer,
		d,
		config,
		"//runtimeconfig.googleapis.com/{{config}}",
		"runtimeconfig.googleapis.com/Config",
	)
}
