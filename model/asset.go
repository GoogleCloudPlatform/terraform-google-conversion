package model

import (
	"github.com/GoogleCloudPlatform/terraform-validator/converters/google"
)

// Asset represents a CAI Asset.
// In the long run, this will be defined in this package; for now,
// re-export the Asset type from terraform-validator.
type Asset = google.Asset
