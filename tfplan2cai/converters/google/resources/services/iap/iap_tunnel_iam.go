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

package iap

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const IapTunnelIAMAssetType string = "iap.googleapis.com/Tunnel"

func ResourceConverterIapTunnelIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         IapTunnelIAMAssetType,
		Convert:           GetIapTunnelIamPolicyCaiObject,
		MergeCreateUpdate: MergeIapTunnelIamPolicy,
	}
}

func ResourceConverterIapTunnelIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         IapTunnelIAMAssetType,
		Convert:           GetIapTunnelIamBindingCaiObject,
		FetchFullResource: FetchIapTunnelIamPolicy,
		MergeCreateUpdate: MergeIapTunnelIamBinding,
		MergeDelete:       MergeIapTunnelIamBindingDelete,
	}
}

func ResourceConverterIapTunnelIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         IapTunnelIAMAssetType,
		Convert:           GetIapTunnelIamMemberCaiObject,
		FetchFullResource: FetchIapTunnelIamPolicy,
		MergeCreateUpdate: MergeIapTunnelIamMember,
		MergeDelete:       MergeIapTunnelIamMemberDelete,
	}
}

func GetIapTunnelIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newIapTunnelIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetIapTunnelIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newIapTunnelIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetIapTunnelIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newIapTunnelIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeIapTunnelIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeIapTunnelIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeIapTunnelIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeIapTunnelIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeIapTunnelIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newIapTunnelIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//iap.googleapis.com/projects/{{project}}/iap_tunnel")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: IapTunnelIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchIapTunnelIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value

	return cai.FetchIamPolicy(
		IapTunnelIamUpdaterProducer,
		d,
		config,
		"//iap.googleapis.com/projects/{{project}}/iap_tunnel",
		IapTunnelIAMAssetType,
	)
}
