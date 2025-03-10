// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/vertexai/Featurestore.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter_iam.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package vertexai

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const VertexAIFeaturestoreIAMAssetType string = "aiplatform.googleapis.com/Featurestore"

func ResourceConverterVertexAIFeaturestoreIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         VertexAIFeaturestoreIAMAssetType,
		Convert:           GetVertexAIFeaturestoreIamPolicyCaiObject,
		MergeCreateUpdate: MergeVertexAIFeaturestoreIamPolicy,
	}
}

func ResourceConverterVertexAIFeaturestoreIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         VertexAIFeaturestoreIAMAssetType,
		Convert:           GetVertexAIFeaturestoreIamBindingCaiObject,
		FetchFullResource: FetchVertexAIFeaturestoreIamPolicy,
		MergeCreateUpdate: MergeVertexAIFeaturestoreIamBinding,
		MergeDelete:       MergeVertexAIFeaturestoreIamBindingDelete,
	}
}

func ResourceConverterVertexAIFeaturestoreIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         VertexAIFeaturestoreIAMAssetType,
		Convert:           GetVertexAIFeaturestoreIamMemberCaiObject,
		FetchFullResource: FetchVertexAIFeaturestoreIamPolicy,
		MergeCreateUpdate: MergeVertexAIFeaturestoreIamMember,
		MergeDelete:       MergeVertexAIFeaturestoreIamMemberDelete,
	}
}

func GetVertexAIFeaturestoreIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newVertexAIFeaturestoreIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetVertexAIFeaturestoreIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newVertexAIFeaturestoreIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetVertexAIFeaturestoreIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newVertexAIFeaturestoreIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeVertexAIFeaturestoreIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeVertexAIFeaturestoreIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeVertexAIFeaturestoreIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeVertexAIFeaturestoreIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeVertexAIFeaturestoreIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newVertexAIFeaturestoreIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//aiplatform.googleapis.com/projects/{{project}}/locations/{{region}}/featurestores/{{featurestore}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: VertexAIFeaturestoreIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchVertexAIFeaturestoreIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("region"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("featurestore"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		VertexAIFeaturestoreIamUpdaterProducer,
		d,
		config,
		"//aiplatform.googleapis.com/projects/{{project}}/locations/{{region}}/featurestores/{{featurestore}}",
		VertexAIFeaturestoreIAMAssetType,
	)
}
