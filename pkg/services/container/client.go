package container

import (
	"google.golang.org/api/container/v1"
	"google.golang.org/api/option"
	"log"

	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/transport"
)

func NewClient(c *transport_tpg.Config, userAgent string) *container.Service {
	containerClientBasePath := transport_tpg.RemoveBasePathVersion(c.ContainerBasePath)
	log.Printf("[INFO] Instantiating GKE client for path %s", containerClientBasePath)
	clientContainer, err := container.NewService(c.Context, option.WithHTTPClient(c.Client))
	if err != nil {
		log.Printf("[WARN] Error creating client container: %s", err)
		return nil
	}
	clientContainer.UserAgent = userAgent
	clientContainer.BasePath = containerClientBasePath

	return clientContainer
}
