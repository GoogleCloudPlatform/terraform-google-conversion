package cai2hcl_test

import (
	"testing"

	cai2hclTesting "github.com/GoogleCloudPlatform/terraform-google-conversion/v7/cai2hcl/testing"
)

func TestConvertCompute(t *testing.T) {
	cai2hclTesting.AssertTestFiles(
		t,
		"./services/compute/testdata",
		[]string{
			"full_compute_instance",
		})
}

func TestConvertResourcemanager(t *testing.T) {
	cai2hclTesting.AssertTestFiles(
		t,
		"./services/resourcemanager/testdata",
		[]string{
			"project_create",
		})
}

func TestConvertNetworksecurity(t *testing.T) {
	cai2hclTesting.AssertTestFiles(
		t,
		"./services/networksecurity/testdata",
		[]string{
			"server_tls_policy",
			"backend_authentication_config",
		})
}
