package google

import transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/transport"

// readDiskType finds the disk type with the given name.
func readDiskType(c *transport_tpg.Config, d TerraformResourceData, name string) (*ZonalFieldValue, error) {
	return parseZonalFieldValue("diskTypes", name, "project", "zone", d, c, false)
}

// readRegionDiskType finds the disk type with the given name.
func readRegionDiskType(c *transport_tpg.Config, d TerraformResourceData, name string) (*RegionalFieldValue, error) {
	return parseRegionalFieldValue("diskTypes", name, "project", "region", "zone", d, c, false)
}
