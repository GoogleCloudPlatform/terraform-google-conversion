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

import "fmt"

func GetBigQueryTableIamPolicyCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newBigQueryTableIamAsset(d, config, expandIamPolicyBindings)
}

func GetBigQueryTableIamBindingCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newBigQueryTableIamAsset(d, config, expandIamRoleBindings)
}

func GetBigQueryTableIamMemberCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	return newBigQueryTableIamAsset(d, config, expandIamMemberBindings)
}

func MergeBigQueryTableIamPolicy(existing, incoming Asset) Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeBigQueryTableIamBinding(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAuthoritativeBindings)
}

func MergeBigQueryTableIamBindingDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAuthoritativeBindings)
}

func MergeBigQueryTableIamMember(existing, incoming Asset) Asset {
	return mergeIamAssets(existing, incoming, mergeAdditiveBindings)
}

func MergeBigQueryTableIamMemberDelete(existing, incoming Asset) Asset {
	return mergeDeleteIamAssets(existing, incoming, mergeDeleteAdditiveBindings)
}

func newBigQueryTableIamAsset(
	d TerraformResourceData,
	config *Config,
	expandBindings func(d TerraformResourceData) ([]IAMBinding, error),
) ([]Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := assetName(d, config, "//bigquery.googleapis.com/{{table}}")
	if err != nil {
		return []Asset{}, err
	}

	return []Asset{{
		Name: name,
		Type: "bigquery.googleapis.com/Table",
		IAMPolicy: &IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchBigQueryTableIamPolicy(d TerraformResourceData, config *Config) (Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("{{table}}"); !ok {
		return Asset{}, ErrEmptyIdentityField
	}

	return fetchIamPolicy(
		BigQueryTableIamUpdaterProducer,
		d,
		config,
		"//bigquery.googleapis.com/{{table}}",
		"bigquery.googleapis.com/Table",
	)
}
