package google

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/tfplan2cai/converters/google/resources/tpgresource"
)

func canonicalizeServiceScope(scope string) string {
	return tpgresource.CanonicalizeServiceScope(scope)
}

func canonicalizeServiceScopes(scopes []string) []string {
	return tpgresource.CanonicalizeServiceScopes(scopes)
}

func stringScopeHashcode(v interface{}) int {
	return tpgresource.StringScopeHashcode(v)
}
