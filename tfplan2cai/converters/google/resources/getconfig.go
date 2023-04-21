package google

import (
	"context"
	"net/http"

	"github.com/pkg/errors"

	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"
)

func NewConfig(ctx context.Context, project, zone, region string, offline bool, userAgent string, client *http.Client) (*transport_tpg.Config, error) {
	cfg := &transport_tpg.Config{
		Project:   project,
		Zone:      zone,
		Region:    region,
		UserAgent: userAgent,
	}

	// Search for default credentials
	cfg.Credentials = MultiEnvSearch([]string{
		"GOOGLE_CREDENTIALS",
		"GOOGLE_CLOUD_KEYFILE_JSON",
		"GCLOUD_KEYFILE_JSON",
	})

	cfg.AccessToken = MultiEnvSearch([]string{
		"GOOGLE_OAUTH_ACCESS_TOKEN",
	})

	cfg.ImpersonateServiceAccount = MultiEnvSearch([]string{
		"GOOGLE_IMPERSONATE_SERVICE_ACCOUNT",
	})

	if !offline {
		transport_tpg.ConfigureBasePaths(cfg)
		if err := cfg.LoadAndValidate(ctx); err != nil {
			return nil, errors.Wrap(err, "load and validate config")
		}
		if client != nil {
			cfg.Client = client
		}
	}

	return cfg, nil
}
