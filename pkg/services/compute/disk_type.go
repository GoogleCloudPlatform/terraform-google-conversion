package compute

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/tpgresource"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/transport"
)

// readDiskType finds the disk type with the given name.
func readDiskType(c *transport_tpg.Config, d tpgresource.TerraformResourceData, name string) (*tpgresource.ZonalFieldValue, error) {
	return tpgresource.ParseZonalFieldValue("diskTypes", name, "project", "zone", d, c, false)
}

// readRegionDiskType finds the disk type with the given name.
func readRegionDiskType(c *transport_tpg.Config, d tpgresource.TerraformResourceData, name string) (*tpgresource.RegionalFieldValue, error) {
	return tpgresource.ParseRegionalFieldValue("diskTypes", name, "project", "region", "zone", d, c, false)
}
