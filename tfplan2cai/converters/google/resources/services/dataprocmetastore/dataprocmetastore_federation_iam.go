// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/metastore/Federation.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter_iam.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package dataprocmetastore

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const DataprocMetastoreFederationIAMAssetType string = "metastore.googleapis.com/Federation"

func ResourceConverterDataprocMetastoreFederationIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataprocMetastoreFederationIAMAssetType,
		Convert:           GetDataprocMetastoreFederationIamPolicyCaiObject,
		MergeCreateUpdate: MergeDataprocMetastoreFederationIamPolicy,
	}
}

func ResourceConverterDataprocMetastoreFederationIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataprocMetastoreFederationIAMAssetType,
		Convert:           GetDataprocMetastoreFederationIamBindingCaiObject,
		FetchFullResource: FetchDataprocMetastoreFederationIamPolicy,
		MergeCreateUpdate: MergeDataprocMetastoreFederationIamBinding,
		MergeDelete:       MergeDataprocMetastoreFederationIamBindingDelete,
	}
}

func ResourceConverterDataprocMetastoreFederationIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         DataprocMetastoreFederationIAMAssetType,
		Convert:           GetDataprocMetastoreFederationIamMemberCaiObject,
		FetchFullResource: FetchDataprocMetastoreFederationIamPolicy,
		MergeCreateUpdate: MergeDataprocMetastoreFederationIamMember,
		MergeDelete:       MergeDataprocMetastoreFederationIamMemberDelete,
	}
}

func GetDataprocMetastoreFederationIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataprocMetastoreFederationIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetDataprocMetastoreFederationIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataprocMetastoreFederationIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetDataprocMetastoreFederationIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newDataprocMetastoreFederationIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeDataprocMetastoreFederationIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeDataprocMetastoreFederationIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeDataprocMetastoreFederationIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeDataprocMetastoreFederationIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeDataprocMetastoreFederationIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newDataprocMetastoreFederationIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//metastore.googleapis.com/projects/{{project}}/locations/{{location}}/federations/{{federation_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: DataprocMetastoreFederationIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchDataprocMetastoreFederationIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("federation_id"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		DataprocMetastoreFederationIamUpdaterProducer,
		d,
		config,
		"//metastore.googleapis.com/projects/{{project}}/locations/{{location}}/federations/{{federation_id}}",
		DataprocMetastoreFederationIAMAssetType,
	)
}
