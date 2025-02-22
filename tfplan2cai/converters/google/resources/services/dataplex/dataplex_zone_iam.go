// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/dataplex/Zone.yaml
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
const DataplexZoneIAMAssetType string = "dataplex.googleapis.com/Zone"

func ResourceConverterDataplexZoneIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataplexZoneIAMAssetType,
		Convert:           GetDataplexZoneIamPolicyCaiObject,
		MergeCreateUpdate: MergeDataplexZoneIamPolicy,
	}
}

func ResourceConverterDataplexZoneIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataplexZoneIAMAssetType,
		Convert:           GetDataplexZoneIamBindingCaiObject,
		FetchFullResource: FetchDataplexZoneIamPolicy,
		MergeCreateUpdate: MergeDataplexZoneIamBinding,
		MergeDelete:       MergeDataplexZoneIamBindingDelete,
	}
}

func ResourceConverterDataplexZoneIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataplexZoneIAMAssetType,
		Convert:           GetDataplexZoneIamMemberCaiObject,
		FetchFullResource: FetchDataplexZoneIamPolicy,
		MergeCreateUpdate: MergeDataplexZoneIamMember,
		MergeDelete:       MergeDataplexZoneIamMemberDelete,
	}
}

func GetDataplexZoneIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataplexZoneIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetDataplexZoneIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataplexZoneIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetDataplexZoneIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataplexZoneIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeDataplexZoneIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeDataplexZoneIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeDataplexZoneIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeDataplexZoneIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeDataplexZoneIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newDataplexZoneIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//dataplex.googleapis.com/projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{dataplex_zone}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: DataplexZoneIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchDataplexZoneIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("lake"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("dataplex_zone"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		DataplexZoneIamUpdaterProducer,
		d,
		config,
		"//dataplex.googleapis.com/projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{dataplex_zone}}",
		DataplexZoneIAMAssetType,
	)
}
