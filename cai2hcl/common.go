package cai2hcl

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zclconf/go-cty/cty"
)

// Shared converters context.
type ConverterContext struct {
	name   string
	schema map[string]*tfschema.Schema
}

// Convert assets into hcl blocks.
type Converter func(asset []*caiasset.Asset, context *ConverterContext) ([]*HCLResourceBlock, error)

// HCLResourceBlock identifies the HCL block's labels and content.
type HCLResourceBlock struct {
	Labels []string
	Value  cty.Value
}

// Map key is the CAI Asset type, value is the TF resource name.
var converterNames = map[string]string{
	"compute.googleapis.com/Instance":                "google_compute_instance",
	"compute.googleapis.com/ForwardingRule":          "google_compute_forwarding_rule",
	"cloudresourcemanager.googleapis.com/Project":    "google_project",
	"cloudbilling.googleapis.com/ProjectBillingInfo": "google_project",
}

// Map key is the TF resource name, value - conversion function.
var converterMap = map[string]Converter{
	"google_compute_instance":        ConvertComputeInstances,
	"google_compute_forwarding_rule": ConvertComputeForwardingRules,
	"google_project":                 ConvertProjects,
}
