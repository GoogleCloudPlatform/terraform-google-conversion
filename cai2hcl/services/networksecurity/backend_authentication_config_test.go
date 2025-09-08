package networksecurity_test

import (
	cai2hcl_testing "github.com/GoogleCloudPlatform/terraform-google-conversion/v5/cai2hcl/testing"
	"testing"
)

func TestBackendAuthenticationConfig(t *testing.T) {
	cai2hcl_testing.AssertTestFiles(
		t,
		"./testdata",
		[]string{"backend_authentication_config"})
}
