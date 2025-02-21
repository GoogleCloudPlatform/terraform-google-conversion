// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/gkebackup/BackupPlan.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter_iam.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package gkebackup

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

// Provide a separate asset type constant so we don't have to worry about name conflicts between IAM and non-IAM converter files
const GKEBackupBackupPlanIAMAssetType string = "gkebackup.googleapis.com/BackupPlan"

func ResourceConverterGKEBackupBackupPlanIamPolicy() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         GKEBackupBackupPlanIAMAssetType,
		Convert:           GetGKEBackupBackupPlanIamPolicyCaiObject,
		MergeCreateUpdate: MergeGKEBackupBackupPlanIamPolicy,
	}
}

func ResourceConverterGKEBackupBackupPlanIamBinding() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         GKEBackupBackupPlanIAMAssetType,
		Convert:           GetGKEBackupBackupPlanIamBindingCaiObject,
		FetchFullResource: FetchGKEBackupBackupPlanIamPolicy,
		MergeCreateUpdate: MergeGKEBackupBackupPlanIamBinding,
		MergeDelete:       MergeGKEBackupBackupPlanIamBindingDelete,
	}
}

func ResourceConverterGKEBackupBackupPlanIamMember() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType:         GKEBackupBackupPlanIAMAssetType,
		Convert:           GetGKEBackupBackupPlanIamMemberCaiObject,
		FetchFullResource: FetchGKEBackupBackupPlanIamPolicy,
		MergeCreateUpdate: MergeGKEBackupBackupPlanIamMember,
		MergeDelete:       MergeGKEBackupBackupPlanIamMemberDelete,
	}
}

func GetGKEBackupBackupPlanIamPolicyCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newGKEBackupBackupPlanIamAsset(d, config, cai.ExpandIamPolicyBindings)
}

func GetGKEBackupBackupPlanIamBindingCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newGKEBackupBackupPlanIamAsset(d, config, cai.ExpandIamRoleBindings)
}

func GetGKEBackupBackupPlanIamMemberCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	return newGKEBackupBackupPlanIamAsset(d, config, cai.ExpandIamMemberBindings)
}

func MergeGKEBackupBackupPlanIamPolicy(existing, incoming cai.Asset) cai.Asset {
	existing.IAMPolicy = incoming.IAMPolicy
	return existing
}

func MergeGKEBackupBackupPlanIamBinding(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAuthoritativeBindings)
}

func MergeGKEBackupBackupPlanIamBindingDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAuthoritativeBindings)
}

func MergeGKEBackupBackupPlanIamMember(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeIamAssets(existing, incoming, cai.MergeAdditiveBindings)
}

func MergeGKEBackupBackupPlanIamMemberDelete(existing, incoming cai.Asset) cai.Asset {
	return cai.MergeDeleteIamAssets(existing, incoming, cai.MergeDeleteAdditiveBindings)
}

func newGKEBackupBackupPlanIamAsset(
	d tpgresource.TerraformResourceData,
	config *transport_tpg.Config,
	expandBindings func(d tpgresource.TerraformResourceData) ([]cai.IAMBinding, error),
) ([]cai.Asset, error) {
	bindings, err := expandBindings(d)
	if err != nil {
		return []cai.Asset{}, fmt.Errorf("expanding bindings: %v", err)
	}

	name, err := cai.AssetName(d, config, "//gkebackup.googleapis.com/projects/{{project}}/locations/{{location}}/backupPlans/{{name}}")
	if err != nil {
		return []cai.Asset{}, err
	}

	return []cai.Asset{{
		Name: name,
		Type: GKEBackupBackupPlanIAMAssetType,
		IAMPolicy: &cai.IAMPolicy{
			Bindings: bindings,
		},
	}}, nil
}

func FetchGKEBackupBackupPlanIamPolicy(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (cai.Asset, error) {
	// Check if the identity field returns a value
	if _, ok := d.GetOk("location"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}
	if _, ok := d.GetOk("name"); !ok {
		return cai.Asset{}, cai.ErrEmptyIdentityField
	}

	return cai.FetchIamPolicy(
		GKEBackupBackupPlanIamUpdaterProducer,
		d,
		config,
		"//gkebackup.googleapis.com/projects/{{project}}/locations/{{location}}/backupPlans/{{name}}",
		GKEBackupBackupPlanIAMAssetType,
	)
}
