package compute

import (
	"log"

	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	compute "google.golang.org/api/compute/v0.beta"
	"google.golang.org/api/option"
)

func NewClient(c *transport_tpg.Config, userAgent string) *compute.Service {
	log.Printf("[INFO] Instantiating GCE client for path %s", c.ComputeBasePath)
	clientCompute, err := compute.NewService(c.Context, option.WithHTTPClient(c.Client))
	if err != nil {
		log.Printf("[WARN] Error creating client compute: %s", err)
		return nil
	}
	clientCompute.UserAgent = userAgent
	clientCompute.BasePath = c.ComputeBasePath

	return clientCompute
}
