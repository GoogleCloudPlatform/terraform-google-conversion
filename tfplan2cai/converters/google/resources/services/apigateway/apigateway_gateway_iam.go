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

package apigateway

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgiamresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const ApiGatewayGatewayIAMAssetType string = "apigateway.googleapis.com/Gateway"

func ResourceConverterApiGatewayGatewayIamPolicy() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ApiGatewayGatewayIAMAssetType,
		Convert:           GetApiGatewayGatewayIamPolicyCaiObject,
		MergeCreateUpdate: MergeApiGatewayGatewayIamPolicy,
	}
}

func ResourceConverterApiGatewayGatewayIamBinding() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ApiGatewayGatewayIAMAssetType,
		Convert:           GetApiGatewayGatewayIamBindingCaiObject,
		FetchFullResource: FetchApiGatewayGatewayIamPolicy,
		MergeCreateUpdate: MergeApiGatewayGatewayIamBinding,
		MergeDelete:       MergeApiGatewayGatewayIamBindingDelete,
	}
}

func ResourceConverterApiGatewayGatewayIamMember() tpgresource.ResourceConverter {
	return tpgresource.ResourceConverter{
		AssetType:         ApiGatewayGatewayIAMAssetType,
		Convert:           GetApiGatewayGatewayIamMemberCaiObject,
		FetchFullResource: FetchApiGatewayGatewayIamPolicy,
		MergeCreateUpdate: MergeApiGatewayGatewayIamMember,
		MergeDelete:       MergeApiGatewayGatewayIamMemberDelete,
	}
}

func GetApiGatewayGatewayIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newApiGatewayGatewayIamAsset(d, config, tpgiamresource.ExpandIamPolicyBindings)
}

func GetApiGatewayGatewayIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newApiGatewayGatewayIamAsset(d, config, tpgiamresource.ExpandIamRoleBindings)
}

func GetApiGatewayGatewayIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]tpgresource.Asset, error) {
	return newApiGatewayGatewayIamAsset(d, config, tpgiamresource.ExpandIamMemberBindings)
}

func MergeApiGatewayGatewayIamPolicy(existing, incoming tpgresource.Asset) tpgresource.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeApiGatewayGatewayIamBinding(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAuthoritativeBindings)
}

func MergeApiGatewayGatewayIamBindingDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAuthoritativeBindings)
}

func MergeApiGatewayGatewayIamMember(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeIamAssets(existing, incoming, tpgiamresource.MergeAdditiveBindings)
}

func MergeApiGatewayGatewayIamMemberDelete(existing, incoming tpgresource.Asset) tpgresource.Asset {
	return tpgiamresource.MergeDeleteIamAssets(existing, incoming, tpgiamresource.MergeDeleteAdditiveBindings)
}

func newApiGatewayGatewayIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]tpgresource.IAMBinding, error),
) ([]tpgresource.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []tpgresource.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := tpgresource.AssetName(d, config, "//apigateway.googleapis.com/projects/{{project}}/locations/{{region}}/gateways/{{gateway}}")
	if err != nil {
		return []tpgresource.Asset{}, err
	}

	return []tpgresource.Asset{{
		Name: name,
		Type: ApiGatewayGatewayIAMAssetType,
		IAMPolicy: &tpgresource.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchApiGatewayGatewayIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (tpgresource.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("region"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("gateway"); !ok {
		return tpgresource.Asset{}, tpgresource.ErrEmptyIdentityField
	}

	return tpgiamresource.FetchIamPolicy(
		ApiGatewayGatewayIamUpdaterProducer,
		d,
		config,
		"//apigateway.googleapis.com/projects/{{project}}/locations/{{region}}/gateways/{{gateway}}",
		ApiGatewayGatewayIAMAssetType,
	)
}
