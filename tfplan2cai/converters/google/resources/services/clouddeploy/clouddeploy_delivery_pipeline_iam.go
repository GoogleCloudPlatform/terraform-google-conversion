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

package clouddeploy

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v5/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const ClouddeployDeliveryPipelineIAMAssetType string = "clouddeploy.googleapis.com/DeliveryPipeline"

func ResourceConverterClouddeployDeliveryPipelineIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ClouddeployDeliveryPipelineIAMAssetType,
		Convert:           GetClouddeployDeliveryPipelineIamPolicyCaiObject,
		MergeCreateUpdate: MergeClouddeployDeliveryPipelineIamPolicy,
	}
}

func ResourceConverterClouddeployDeliveryPipelineIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ClouddeployDeliveryPipelineIAMAssetType,
		Convert:           GetClouddeployDeliveryPipelineIamBindingCaiObject,
		FetchFullResource: FetchClouddeployDeliveryPipelineIamPolicy,
		MergeCreateUpdate: MergeClouddeployDeliveryPipelineIamBinding,
		MergeDelete:       MergeClouddeployDeliveryPipelineIamBindingDelete,
	}
}

func ResourceConverterClouddeployDeliveryPipelineIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ClouddeployDeliveryPipelineIAMAssetType,
		Convert:           GetClouddeployDeliveryPipelineIamMemberCaiObject,
		FetchFullResource: FetchClouddeployDeliveryPipelineIamPolicy,
		MergeCreateUpdate: MergeClouddeployDeliveryPipelineIamMember,
		MergeDelete:       MergeClouddeployDeliveryPipelineIamMemberDelete,
	}
}

func GetClouddeployDeliveryPipelineIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newClouddeployDeliveryPipelineIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetClouddeployDeliveryPipelineIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newClouddeployDeliveryPipelineIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetClouddeployDeliveryPipelineIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newClouddeployDeliveryPipelineIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeClouddeployDeliveryPipelineIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeClouddeployDeliveryPipelineIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeClouddeployDeliveryPipelineIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeClouddeployDeliveryPipelineIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeClouddeployDeliveryPipelineIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newClouddeployDeliveryPipelineIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//clouddeploy.googleapis.com/projects/{{project}}/locations/{{location}}/deliveryPipelines/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: ClouddeployDeliveryPipelineIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchClouddeployDeliveryPipelineIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("name"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		ClouddeployDeliveryPipelineIamUpdaterProducer,
		d,
		config,
		"//clouddeploy.googleapis.com/projects/{{project}}/locations/{{location}}/deliveryPipelines/{{name}}",
		ClouddeployDeliveryPipelineIAMAssetType,
	)
}
