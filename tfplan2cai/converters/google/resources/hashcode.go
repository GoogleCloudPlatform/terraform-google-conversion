package google

import "github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"

// hashcode hashes a string to a unique hashcode.
//
// crc32 returns a uint32, but for our use we need
// and non negative integer. Here we cast to an integer
// and invert it if the result is negative.
//
// Deprecated: For backward compatibility hashcode is still working,
// but all new code should use Hashcode in the tpgresource package instead.
func hashcode(s string) int {
	return tpgresource.Hashcode(s)
}
