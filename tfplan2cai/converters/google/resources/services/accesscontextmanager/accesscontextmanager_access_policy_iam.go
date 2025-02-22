// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/accesscontextmanager/AccessPolicy.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter_iam.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package accesscontextmanager

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const AccessContextManagerAccessPolicyIAMAssetType string = "accesscontextmanager.googleapis.com/AccessPolicy"

func ResourceConverterAccessContextManagerAccessPolicyIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         AccessContextManagerAccessPolicyIAMAssetType,
		Convert:           GetAccessContextManagerAccessPolicyIamPolicyCaiObject,
		MergeCreateUpdate: MergeAccessContextManagerAccessPolicyIamPolicy,
	}
}

func ResourceConverterAccessContextManagerAccessPolicyIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         AccessContextManagerAccessPolicyIAMAssetType,
		Convert:           GetAccessContextManagerAccessPolicyIamBindingCaiObject,
		FetchFullResource: FetchAccessContextManagerAccessPolicyIamPolicy,
		MergeCreateUpdate: MergeAccessContextManagerAccessPolicyIamBinding,
		MergeDelete:       MergeAccessContextManagerAccessPolicyIamBindingDelete,
	}
}

func ResourceConverterAccessContextManagerAccessPolicyIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         AccessContextManagerAccessPolicyIAMAssetType,
		Convert:           GetAccessContextManagerAccessPolicyIamMemberCaiObject,
		FetchFullResource: FetchAccessContextManagerAccessPolicyIamPolicy,
		MergeCreateUpdate: MergeAccessContextManagerAccessPolicyIamMember,
		MergeDelete:       MergeAccessContextManagerAccessPolicyIamMemberDelete,
	}
}

func GetAccessContextManagerAccessPolicyIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newAccessContextManagerAccessPolicyIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetAccessContextManagerAccessPolicyIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newAccessContextManagerAccessPolicyIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetAccessContextManagerAccessPolicyIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newAccessContextManagerAccessPolicyIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeAccessContextManagerAccessPolicyIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeAccessContextManagerAccessPolicyIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeAccessContextManagerAccessPolicyIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeAccessContextManagerAccessPolicyIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeAccessContextManagerAccessPolicyIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newAccessContextManagerAccessPolicyIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//accesscontextmanager.googleapis.com/accessPolicies/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: AccessContextManagerAccessPolicyIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchAccessContextManagerAccessPolicyIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("name"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		AccessContextManagerAccessPolicyIamUpdaterProducer,
		d,
		config,
		"//accesscontextmanager.googleapis.com/accessPolicies/{{name}}",
		AccessContextManagerAccessPolicyIAMAssetType,
	)
}
