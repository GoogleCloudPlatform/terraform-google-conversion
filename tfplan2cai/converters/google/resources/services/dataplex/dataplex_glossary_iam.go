// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/dataplex/Glossary.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter_iam.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package dataplex

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const DataplexGlossaryIAMAssetType string = "dataplex.googleapis.com/Glossary"

func ResourceConverterDataplexGlossaryIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataplexGlossaryIAMAssetType,
		Convert:           GetDataplexGlossaryIamPolicyCaiObject,
		MergeCreateUpdate: MergeDataplexGlossaryIamPolicy,
	}
}

func ResourceConverterDataplexGlossaryIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataplexGlossaryIAMAssetType,
		Convert:           GetDataplexGlossaryIamBindingCaiObject,
		FetchFullResource: FetchDataplexGlossaryIamPolicy,
		MergeCreateUpdate: MergeDataplexGlossaryIamBinding,
		MergeDelete:       MergeDataplexGlossaryIamBindingDelete,
	}
}

func ResourceConverterDataplexGlossaryIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataplexGlossaryIAMAssetType,
		Convert:           GetDataplexGlossaryIamMemberCaiObject,
		FetchFullResource: FetchDataplexGlossaryIamPolicy,
		MergeCreateUpdate: MergeDataplexGlossaryIamMember,
		MergeDelete:       MergeDataplexGlossaryIamMemberDelete,
	}
}

func GetDataplexGlossaryIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataplexGlossaryIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetDataplexGlossaryIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataplexGlossaryIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetDataplexGlossaryIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataplexGlossaryIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeDataplexGlossaryIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeDataplexGlossaryIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeDataplexGlossaryIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeDataplexGlossaryIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeDataplexGlossaryIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newDataplexGlossaryIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//dataplex.googleapis.com/projects/{{project}}/locations/{{location}}/glossaries/{{glossary_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: DataplexGlossaryIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchDataplexGlossaryIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("glossary_id"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		DataplexGlossaryIamUpdaterProducer,
		d,
		config,
		"//dataplex.googleapis.com/projects/{{project}}/locations/{{location}}/glossaries/{{glossary_id}}",
		DataplexGlossaryIAMAssetType,
	)
}
