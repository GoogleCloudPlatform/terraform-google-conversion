package cai2hcl

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/caiasset"

	"google.golang.org/api/compute/v1"
)

// Convert converts asset to HCL resource blocks.
func ConvertComputeForwardingRules(assets []*caiasset.Asset, context *ConverterContext) ([]*HCLResourceBlock, error) {
	var blocks []*HCLResourceBlock
	for _, asset := range assets {
		if asset == nil {
			continue
		}
		if asset.Resource != nil && asset.Resource.Data != nil {
			block, err := convertComputeForwardingRule(asset, context)
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
func convertComputeForwardingRule(asset *caiasset.Asset, context *ConverterContext) (*HCLResourceBlock, error) {
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

	ctyVal, err := mapToCtyValWithSchema(hcl, context.schema)
	if err != nil {
		return nil, err
	}
	return &HCLResourceBlock{
		Labels: []string{context.name, resource.Name},
		Value:  ctyVal,
	}, nil
}
