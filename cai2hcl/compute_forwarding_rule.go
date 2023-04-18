package cai2hcl

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/compute/v1"
)

// ComputeForwardingRuleAssetType is the CAI asset type name for compute instance.
const ComputeForwardingRuleAssetType string = "compute.googleapis.com/ForwardingRule"

// ComputeForwardingRuleConverter for regional forwarding rule.
type ComputeForwardingRuleConverter struct {
	name   string
	schema map[string]*tfschema.Schema
}

// NewComputeForwardingRuleConverter returns an HCL converter for compute instance.
func NewComputeForwardingRuleConverter() *ComputeForwardingRuleConverter {
	return &ComputeForwardingRuleConverter{
		name:   "google_compute_forwarding_rule",
		schema: schemaProvider.ResourcesMap["google_compute_forwarding_rule"].Schema,
	}
}

// Convert converts asset to HCL resource blocks.
func (c *ComputeForwardingRuleConverter) Convert(assets []*caiasset.Asset) ([]*HCLResourceBlock, error) {
	var blocks []*HCLResourceBlock
	for _, asset := range assets {
		if asset == nil {
			continue
		}
		if asset.Resource != nil && asset.Resource.Data != nil {
			block, err := c.convertResourceData(asset)
			if err != nil {
				return nil, err
			}
			blocks = append(blocks, block)
		}
	}
	return blocks, nil
}

// Convert REST payload to JSON/
// Ported from https://github.com/hashicorp/terraform-provider-google/blob/main/google/resource_compute_forwarding_rule.go#L351
func (c *ComputeForwardingRuleConverter) convertResourceData(asset *caiasset.Asset) (*HCLResourceBlock, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	var resource *compute.ForwardingRule
	if err := decodeJSON(asset.Resource.Data, &resource); err != nil {
		return nil, err
	}

	hcl := map[string]interface{}{
		"name":                            resource.Name,
		"all_ports":                       resource.AllPorts,
		"allow_global_access":             resource.AllowGlobalAccess,
		"backend_service":                 resource.BackendService,
		"description":                     resource.Description,
		"ip_address":                      resource.IPAddress,
		"ip_protocol":                     resource.IPProtocol,
		"is_mirroring_collector":          resource.IsMirroringCollector,
		"labels":                          resource.Labels,
		"load_balancing_scheme":           resource.LoadBalancingScheme,
		"network":                         resource.Network,
		"network_tier":                    resource.NetworkTier,
		"port_range":                      resource.PortRange,
		"ports":                           resource.Ports,
		"subnetwork":                      resource.Subnetwork,
		"target":                          resource.Target,
		"service_label":                   resource.ServiceLabel,
		"service_directory_registrations": resource.ServiceDirectoryRegistrations,
	}

	if resource.Region != "" {
		hcl["region"] = parseFieldValue(resource.Region, "regions")
	}

	ctyVal, err := mapToCtyValWithSchema(hcl, c.schema)
	if err != nil {
		return nil, err
	}
	return &HCLResourceBlock{
		Labels: []string{c.name, resource.Name},
		Value:  ctyVal,
	}, nil
}
