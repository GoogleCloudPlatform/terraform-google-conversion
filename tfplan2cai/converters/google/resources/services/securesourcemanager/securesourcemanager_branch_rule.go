// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/securesourcemanager/BranchRule.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package securesourcemanager

import (
	"reflect"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai/converters/google/resources/cai"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

const SecureSourceManagerBranchRuleAssetType string = "securesourcemanager.googleapis.com/BranchRule"

func ResourceConverterSecureSourceManagerBranchRule() cai.ResourceConverter {
	return cai.ResourceConverter{
		AssetType: SecureSourceManagerBranchRuleAssetType,
		Convert:   GetSecureSourceManagerBranchRuleCaiObject,
	}
}

func GetSecureSourceManagerBranchRuleCaiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) ([]cai.Asset, error) {
	name, err := cai.AssetName(d, config, "//securesourcemanager.googleapis.com/projects/{{project}}/locations/{{location}}/repositories/{{repository_id}}/branchRules/{{branch_rule_id}}")
	if err != nil {
		return []cai.Asset{}, err
	}
	if obj, err := GetSecureSourceManagerBranchRuleApiObject(d, config); err == nil {
		return []cai.Asset{{
			Name: name,
			Type: SecureSourceManagerBranchRuleAssetType,
			Resource: &cai.AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/securesourcemanager/v1/rest",
				DiscoveryName:        "BranchRule",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []cai.Asset{}, err
	}
}

func GetSecureSourceManagerBranchRuleApiObject(d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	includePatternProp, err := expandSecureSourceManagerBranchRuleIncludePattern(d.Get("include_pattern"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("include_pattern"); !tpgresource.IsEmptyValue(reflect.ValueOf(includePatternProp)) && (ok || !reflect.DeepEqual(v, includePatternProp)) {
		obj["includePattern"] = includePatternProp
	}
	disabledProp, err := expandSecureSourceManagerBranchRuleDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	requirePullRequestProp, err := expandSecureSourceManagerBranchRuleRequirePullRequest(d.Get("require_pull_request"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("require_pull_request"); !tpgresource.IsEmptyValue(reflect.ValueOf(requirePullRequestProp)) && (ok || !reflect.DeepEqual(v, requirePullRequestProp)) {
		obj["requirePullRequest"] = requirePullRequestProp
	}
	minimumReviewsCountProp, err := expandSecureSourceManagerBranchRuleMinimumReviewsCount(d.Get("minimum_reviews_count"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("minimum_reviews_count"); !tpgresource.IsEmptyValue(reflect.ValueOf(minimumReviewsCountProp)) && (ok || !reflect.DeepEqual(v, minimumReviewsCountProp)) {
		obj["minimumReviewsCount"] = minimumReviewsCountProp
	}
	minimumApprovalsCountProp, err := expandSecureSourceManagerBranchRuleMinimumApprovalsCount(d.Get("minimum_approvals_count"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("minimum_approvals_count"); !tpgresource.IsEmptyValue(reflect.ValueOf(minimumApprovalsCountProp)) && (ok || !reflect.DeepEqual(v, minimumApprovalsCountProp)) {
		obj["minimumApprovalsCount"] = minimumApprovalsCountProp
	}
	requireCommentsResolvedProp, err := expandSecureSourceManagerBranchRuleRequireCommentsResolved(d.Get("require_comments_resolved"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("require_comments_resolved"); !tpgresource.IsEmptyValue(reflect.ValueOf(requireCommentsResolvedProp)) && (ok || !reflect.DeepEqual(v, requireCommentsResolvedProp)) {
		obj["requireCommentsResolved"] = requireCommentsResolvedProp
	}
	allowStaleReviewsProp, err := expandSecureSourceManagerBranchRuleAllowStaleReviews(d.Get("allow_stale_reviews"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("allow_stale_reviews"); !tpgresource.IsEmptyValue(reflect.ValueOf(allowStaleReviewsProp)) && (ok || !reflect.DeepEqual(v, allowStaleReviewsProp)) {
		obj["allowStaleReviews"] = allowStaleReviewsProp
	}
	requireLinearHistoryProp, err := expandSecureSourceManagerBranchRuleRequireLinearHistory(d.Get("require_linear_history"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("require_linear_history"); !tpgresource.IsEmptyValue(reflect.ValueOf(requireLinearHistoryProp)) && (ok || !reflect.DeepEqual(v, requireLinearHistoryProp)) {
		obj["requireLinearHistory"] = requireLinearHistoryProp
	}

	return obj, nil
}

func expandSecureSourceManagerBranchRuleIncludePattern(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerBranchRuleDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerBranchRuleRequirePullRequest(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerBranchRuleMinimumReviewsCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerBranchRuleMinimumApprovalsCount(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerBranchRuleRequireCommentsResolved(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerBranchRuleAllowStaleReviews(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandSecureSourceManagerBranchRuleRequireLinearHistory(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}
