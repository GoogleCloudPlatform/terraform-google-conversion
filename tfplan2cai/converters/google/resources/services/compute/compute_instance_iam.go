// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/Instance.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter_iam.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const ComputeInstanceIAMAssetType string = "compute.googleapis.com/Instance"

func ResourceConverterComputeInstanceIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ComputeInstanceIAMAssetType,
		Convert:           GetComputeInstanceIamPolicyCaiObject,
		MergeCreateUpdate: MergeComputeInstanceIamPolicy,
	}
}

func ResourceConverterComputeInstanceIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ComputeInstanceIAMAssetType,
		Convert:           GetComputeInstanceIamBindingCaiObject,
		FetchFullResource: FetchComputeInstanceIamPolicy,
		MergeCreateUpdate: MergeComputeInstanceIamBinding,
		MergeDelete:       MergeComputeInstanceIamBindingDelete,
	}
}

func ResourceConverterComputeInstanceIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         ComputeInstanceIAMAssetType,
		Convert:           GetComputeInstanceIamMemberCaiObject,
		FetchFullResource: FetchComputeInstanceIamPolicy,
		MergeCreateUpdate: MergeComputeInstanceIamMember,
		MergeDelete:       MergeComputeInstanceIamMemberDelete,
	}
}

func GetComputeInstanceIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newComputeInstanceIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetComputeInstanceIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newComputeInstanceIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetComputeInstanceIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newComputeInstanceIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeComputeInstanceIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeComputeInstanceIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeComputeInstanceIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeComputeInstanceIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeComputeInstanceIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newComputeInstanceIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/instances/{{instance_name}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: ComputeInstanceIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchComputeInstanceIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("zone"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("instance_name"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		ComputeInstanceIamUpdaterProducer,
		d,
		config,
		"//compute.googleapis.com/projects/{{project}}/zones/{{zone}}/instances/{{instance_name}}",
		ComputeInstanceIAMAssetType,
	)
}
