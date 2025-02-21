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

package healthcare

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const HealthcareConsentStoreIAMAssetType string = "healthcare.googleapis.com/ConsentStore"

func ResourceConverterHealthcareConsentStoreIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         HealthcareConsentStoreIAMAssetType,
		Convert:           GetHealthcareConsentStoreIamPolicyCaiObject,
		MergeCreateUpdate: MergeHealthcareConsentStoreIamPolicy,
	}
}

func ResourceConverterHealthcareConsentStoreIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         HealthcareConsentStoreIAMAssetType,
		Convert:           GetHealthcareConsentStoreIamBindingCaiObject,
		FetchFullResource: FetchHealthcareConsentStoreIamPolicy,
		MergeCreateUpdate: MergeHealthcareConsentStoreIamBinding,
		MergeDelete:       MergeHealthcareConsentStoreIamBindingDelete,
	}
}

func ResourceConverterHealthcareConsentStoreIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         HealthcareConsentStoreIAMAssetType,
		Convert:           GetHealthcareConsentStoreIamMemberCaiObject,
		FetchFullResource: FetchHealthcareConsentStoreIamPolicy,
		MergeCreateUpdate: MergeHealthcareConsentStoreIamMember,
		MergeDelete:       MergeHealthcareConsentStoreIamMemberDelete,
	}
}

func GetHealthcareConsentStoreIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newHealthcareConsentStoreIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetHealthcareConsentStoreIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newHealthcareConsentStoreIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetHealthcareConsentStoreIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newHealthcareConsentStoreIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeHealthcareConsentStoreIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeHealthcareConsentStoreIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeHealthcareConsentStoreIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeHealthcareConsentStoreIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeHealthcareConsentStoreIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newHealthcareConsentStoreIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//healthcare.googleapis.com/{{dataset}}/consentStores/{{consent_store_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: HealthcareConsentStoreIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchHealthcareConsentStoreIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("dataset"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("consent_store_id"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		HealthcareConsentStoreIamUpdaterProducer,
		d,
		config,
		"//healthcare.googleapis.com/{{dataset}}/consentStores/{{consent_store_id}}",
		HealthcareConsentStoreIAMAssetType,
	)
}
