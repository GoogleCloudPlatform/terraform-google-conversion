package cai2hcl

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"

	"github.com/zclconf/go-cty/cty"
	"google.golang.org/api/cloudresourcemanager/v1"
)

// Convert converts asset resource data.
func ConvertProjects(assets []*caiasset.Asset, context *ConverterContext) ([]*HCLResourceBlock, error) {
	var billings = make(map[string]string)

	// process billing info
	for _, asset := range assets {
		if asset == nil {
			continue
		}
		if asset.Type == "cloudbilling.googleapis.com/ProjectBillingInfo" {
			project := parseFieldValue(asset.Name, "projects")
			projectAssetName := fmt.Sprintf("//cloudresourcemanager.googleapis.com/projects/%s", project)
			billings[projectAssetName] = convertBilling(asset)
		}
	}

	var blocks []*HCLResourceBlock
	for _, asset := range assets {
		if asset == nil {
			continue
		}
		if asset.Type == "cloudbilling.googleapis.com/ProjectBillingInfo" {
			continue
		}
		if asset.IAMPolicy != nil {
			iamBlock, err := convertIam(asset, context)
			if err != nil {
				return nil, err
			}
			blocks = append(blocks, iamBlock)
		}
		if asset.Resource != nil && asset.Resource.Data != nil {
			block, err := convertProject(asset, context, billings)
			if err != nil {
				return nil, err
			}
			blocks = append(blocks, block)
		}
	}
	return blocks, nil
}

func convertIam(asset *caiasset.Asset, context *ConverterContext) (*HCLResourceBlock, error) {
	if asset.IAMPolicy == nil {
		return nil, fmt.Errorf("asset IAM policy is nil")
	}

	project := parseFieldValue(asset.Name, "projects")
	policyData, err := json.Marshal(asset.IAMPolicy)
	if err != nil {
		return nil, err
	}

	return &HCLResourceBlock{
		Labels: []string{
			context.name + "_iam_policy",
			project + "_iam_policy",
		},
		Value: cty.ObjectVal(map[string]cty.Value{
			"project":     cty.StringVal(project),
			"policy_data": cty.StringVal(string(policyData)),
		}),
	}, nil
}

func convertBilling(asset *caiasset.Asset) string {
	if asset != nil && asset.Resource != nil && asset.Resource.Data != nil {
		return strings.TrimPrefix(asset.Resource.Data["billingAccountName"].(string), "billingAccounts/")
	}
	return ""
}

func convertProject(asset *caiasset.Asset, context *ConverterContext, billings map[string]string) (*HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}
	var project *cloudresourcemanager.Project
	if err := decodeJSON(asset.Resource.Data, &project); err != nil {
		return nil, err
	}

	hclData := make(map[string]interface{})
	hclData["name"] = project.Name
	hclData["project_id"] = project.ProjectId
	hclData["labels"] = project.Labels
	if strings.Contains(asset.Resource.Parent, "folders/") {
		hclData["folder_id"] = parseFieldValue(asset.Resource.Parent, "folders")
	} else if strings.Contains(asset.Resource.Parent, "organizations/") {
		hclData["org_id"] = parseFieldValue(asset.Resource.Parent, "organizations")
	}

	if billingAccount, ok := billings[asset.Name]; ok {
		hclData["billing_account"] = billingAccount
	}

	ctyVal, err := mapToCtyValWithSchema(hclData, context.schema)
	if err != nil {
		return nil, err
	}
	return &HCLResourceBlock{
		Labels: []string{context.name, project.ProjectId},
		Value:  ctyVal,
	}, nil
}
