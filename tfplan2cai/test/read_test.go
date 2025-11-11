package test

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/tfplan2cai"
	"github.com/google/go-cmp/cmp"
	"go.uber.org/zap/zaptest"
)

func TestReadPlannedAssetsCoverage(t *testing.T) {
	cases := []struct {
		name string
	}{
		// read-only, the following tests are not in cli_test or
		// have unique parameters that separate them
		{name: "example_folder_iam_binding"},
		{name: "example_folder_iam_member"},
		{name: "example_project_create"},
		{name: "example_project_update"},
		{name: "example_project_iam_binding"},
		{name: "example_project_iam_member"},
		{name: "example_storage_bucket"},
		{name: "example_storage_bucket_empty_project_id"},
		{name: "example_storage_bucket_iam_binding"},
		{name: "example_storage_bucket_iam_member"},
		{name: "example_project_create_empty_project_id"},
		{name: "example_project_iam_member_empty_project"},
		// auto inserted tests that are not in list above or manually inserted in cli_test.go
		{name: "bucket"},
		{name: "compute_image"},
		{name: "disk"},
		{name: "example_access_context_manager_access_policy"},
		{name: "example_access_context_manager_service_perimeter"},
		{name: "example_alloydb_instance"},
		{name: "example_apikeys_key"},
		{name: "example_app_engine_application"},
		{name: "example_artifact_registry_repository"},
		{name: "example_bigquery_dataset"},
		{name: "example_bigquery_dataset_iam_binding"},
		{name: "example_bigquery_dataset_iam_member"},
		{name: "example_bigquery_dataset_iam_policy"},
		{name: "example_bigquery_dataset_iam_policy_empty_policy_data"},
		{name: "example_bigquery_table"},
		{name: "example_bigtable_instance"},
		{name: "example_cloud_run_mapping"},
		{name: "example_cloud_run_service"},
		{name: "example_cloud_run_service_iam_binding"},
		{name: "example_cloud_run_service_iam_member"},
		{name: "example_cloud_run_service_iam_policy"},
		{name: "example_cloud_run_v2_job"},
		{name: "example_cloud_tasks_queue"},
		{name: "example_compute_address"},
		{name: "example_compute_disk"},
		{name: "example_compute_disk_empty_image"},
		{name: "example_compute_firewall"},
		{name: "example_compute_forwarding_rule"},
		{name: "example_compute_global_address"},
		{name: "example_compute_global_forwarding_rule"},
		{name: "example_compute_health_check"},
		{name: "example_compute_instance_iam_binding"},
		{name: "example_compute_instance_iam_member"},
		{name: "example_compute_instance_iam_policy"},
		{name: "example_compute_network"},
		{name: "example_compute_route"},
		{name: "example_compute_security_policy"},
		{name: "example_compute_snapshot"},
		{name: "example_compute_ssl_policy"},
		{name: "example_compute_subnetwork"},
		{name: "example_compute_target_https_proxy"},
		{name: "example_compute_target_ssl_proxy"},
		{name: "example_compute_url_map"},
		{name: "example_compute_vpn_tunnel"},
		{name: "example_container_cluster"},
		{name: "example_dns_managed_zone"},
		{name: "example_dns_policy"},
		{name: "example_filestore_instance"},
		{name: "example_folder_iam_member_empty_folder"},
		{name: "example_folder_iam_policy"},
		{name: "example_folder_organization_policy"},
		{name: "example_gke_hub_feature"},
		{name: "example_google_app_engine_standard_app_version"},
		{name: "example_google_cloudfunctions_function"},
		{name: "example_google_composer_environment"},
		{name: "example_google_compute_autoscaler"},
		{name: "example_google_compute_instance_group"},
		{name: "example_google_compute_network_endpoint_group"},
		{name: "example_google_compute_node_group"},
		{name: "example_google_compute_node_template"},
		{name: "example_google_compute_region_commitment"},
		{name: "example_google_compute_resource_policy"},
		{name: "example_google_compute_router"},
		{name: "example_google_compute_ssl_certificate"},
		{name: "example_google_compute_target_http_proxy"},
		{name: "example_google_compute_target_pool"},
		{name: "example_google_dataflow_job"},
		{name: "example_google_dataproc_autoscaling_policy"},
		{name: "example_google_dataproc_cluster"},
		{name: "example_google_datastream_connection_profile"},
		{name: "example_google_datastream_private_connection"},
		{name: "example_google_datastream_stream"},
		{name: "example_google_datastream_stream_append_only"},
		{name: "example_google_firebase_project"},
		{name: "example_google_gke_hub_membership"},
		{name: "example_google_logging_billing_account_bucket_config"},
		{name: "example_google_logging_folder_bucket_config"},
		{name: "example_google_logging_folder_sink"},
		{name: "example_google_logging_organization_bucket_config"},
		{name: "example_google_logging_project_bucket_config"},
		{name: "example_google_logging_project_sink"},
		{name: "example_google_sql_database"},
		{name: "example_kms_crypto_key"},
		{name: "example_kms_crypto_key_iam_binding"},
		{name: "example_kms_crypto_key_iam_member"},
		{name: "example_kms_crypto_key_iam_policy"},
		{name: "example_kms_import_job"},
		{name: "example_kms_key_ring"},
		{name: "example_kms_key_ring_iam_binding"},
		{name: "example_kms_key_ring_iam_member"},
		{name: "example_kms_key_ring_iam_policy"},
		{name: "example_logging_metric"},
		{name: "example_monitoring_alert_policy"},
		{name: "example_monitoring_notification_channel"},
		{name: "example_org_policy_custom_constraint"},
		{name: "example_org_policy_policy"},
		{name: "example_organization_iam_binding"},
		{name: "example_organization_iam_custom_role"},
		{name: "example_organization_iam_member"},
		{name: "example_organization_iam_policy"},
		{name: "example_organization_policy"},
		{name: "example_project_iam"},
		{name: "example_project_iam_custom_role"},
		{name: "example_project_iam_policy"},
		{name: "example_project_in_folder"},
		{name: "example_project_in_org"},
		{name: "example_project_organization_policy"},
		{name: "example_project_service"},
		{name: "example_pubsub_lite_reservation"},
		{name: "example_pubsub_lite_subscription"},
		{name: "example_pubsub_lite_topic"},
		{name: "example_pubsub_schema"},
		{name: "example_pubsub_subscription"},
		{name: "example_pubsub_subscription_iam_binding"},
		{name: "example_pubsub_subscription_iam_member"},
		{name: "example_pubsub_subscription_iam_policy"},
		{name: "example_pubsub_topic"},
		{name: "example_redis_instance"},
		{name: "example_secret_manager_secret_iam_binding"},
		{name: "example_secret_manager_secret_iam_member"},
		{name: "example_secret_manager_secret_iam_policy"},
		{name: "example_service_account"},
		{name: "example_service_account_key"},
		{name: "example_service_account_update"},
		{name: "example_spanner_database"},
		{name: "example_spanner_database_iam_binding"},
		{name: "example_spanner_database_iam_member"},
		{name: "example_spanner_database_iam_policy"},
		{name: "example_spanner_instance_iam_binding"},
		{name: "example_spanner_instance_iam_member"},
		{name: "example_spanner_instance_iam_policy"},
		{name: "example_sql_database_instance"},
		{name: "example_storage_bucket_iam_member_random_suffix"},
		{name: "example_storage_bucket_iam_policy"},
		{name: "example_vertex_ai_dataset"},
		{name: "example_vpc_access_connector"},
		{name: "firewall"},
		{name: "full_compute_firewall"},
		{name: "full_compute_instance"},
		{name: "full_container_cluster"},
		{name: "full_container_node_pool"},
		{name: "full_spanner_instance"},
		{name: "full_sql_database_instance"},
		{name: "full_storage_bucket"},
		{name: "google_iam_workload_identity_pool_provider"},
		{name: "google_vmwareengine_cluster"},
		{name: "google_vmwareengine_external_address"},
		{name: "google_vmwareengine_network_peering"},
		{name: "google_vmwareengine_private_cloud"},
		{name: "instance"},
		{name: "kms_crypto_key_version"},
		{name: "logging_organization_sink"},
		{name: "sql"},
		{name: "vmwareengine_network_policy"},
		{name: "workbench_instance"},
	}
	for i := range cases {
		// Allocate a variable to make sure test can run in parallel.
		c := cases[i]
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			// Create a temporary directory for running terraform.
			dir, err := os.MkdirTemp(tmpDir, "terraform")
			if err != nil {
				log.Fatal(err)
			}
			defer os.RemoveAll(dir)

			generateTestFiles(t, "../testdata/templates", dir, c.name+".json")
			generateTestFiles(t, "../testdata/templates", dir, c.name+".tf")

			// tfstate files are for cases testing updates, eg. project update.
			// Uses glob matching to match generateTestFiles internals.
			tfstateMatches, err := filepath.Glob(filepath.Join("../testdata/templates", c.name+".tfstate"))
			if err != nil {
				t.Fatalf("malformed glob: %v", err)
			}
			if tfstateMatches != nil {
				generateTestFiles(t, "../testdata/templates", dir, c.name+".tfstate")
				err = os.Rename(
					filepath.Join(dir, c.name+".tfstate"),
					filepath.Join(dir, "terraform.tfstate"),
				)
				if err != nil {
					t.Fatalf("renaming tfstate: %v", err)
				}
			}

			// Run terraform init and terraform apply to generate tfplan.json files
			terraformWorkflow(t, dir, c.name)

			// Unmarshal payload from testfile into `want` variable.
			f := filepath.Join(dir, c.name+".json")
			want, err := readExpectedTestFile(f)
			if err != nil {
				t.Fatal(err)
			}

			planfile := filepath.Join(dir, c.name+".tfplan.json")
			ctx := context.Background()
			ancestryCache := map[string]string{
				data.Provider["project"]: data.Ancestry,
			}

			jsonPlan, err := os.ReadFile(planfile)
			if err != nil {
				t.Fatalf("Error parsing %s: %s", f, err)
			}
			got, err := tfplan2cai.Convert(ctx, jsonPlan, &tfplan2cai.Options{
				ConvertUnchanged: false,
				ErrorLogger:      zaptest.NewLogger(t),
				Offline:          true,
				DefaultProject:   data.Provider["project"],
				DefaultRegion:    "",
				DefaultZone:      "",
				UserAgent:        "",
				AncestryCache:    ancestryCache,
			})
			if err != nil {
				t.Fatalf("Convert(%s, %s, \"\", \"\", %s, offline): %v", planfile, data.Provider["project"], ancestryCache, err)
			}
			expectedAssets := normalizeAssets(t, want, true)
			actualAssets := normalizeAssets(t, got, true)
			if diff := cmp.Diff(expectedAssets, actualAssets); diff != "" {
				t.Errorf("%v diff(-want, +got):\n%s", t.Name(), diff)
			}
		})
	}
}

func TestReadPlannedAssetsCoverage_WithoutDefaultProject(t *testing.T) {
	cases := []struct {
		name string
	}{
		{name: "example_project_create_empty_project_id"},
		{name: "example_storage_bucket_empty_project_id"},
		{name: "example_project_iam_member_empty_project"},
	}
	for i := range cases {
		// Allocate a variable to make sure test can run in parallel.
		c := cases[i]
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			// Create a temporary directory for running terraform.
			dir, err := os.MkdirTemp(tmpDir, "terraform")
			if err != nil {
				log.Fatal(err)
			}
			defer os.RemoveAll(dir)

			generateTestFiles(t, "../testdata/templates", dir, c.name+"_without_default_project.json")
			generateTestFiles(t, "../testdata/templates", dir, c.name+".tf")

			// Run terraform init and terraform plan to generate tfplan.json files
			terraformWorkflow(t, dir, c.name)

			// Unmarshal payload from testfile into `want` variable.
			f := filepath.Join(dir, c.name+"_without_default_project.json")
			want, err := readExpectedTestFile(f)
			if err != nil {
				t.Fatal(err)
			}

			planfile := filepath.Join(dir, c.name+".tfplan.json")
			ctx := context.Background()

			jsonPlan, err := os.ReadFile(planfile)
			if err != nil {
				t.Fatalf("Error parsing %s: %s", f, err)
			}
			got, err := tfplan2cai.Convert(ctx, jsonPlan, &tfplan2cai.Options{
				ConvertUnchanged: false,
				ErrorLogger:      zaptest.NewLogger(t),
				Offline:          true,
				DefaultProject:   "",
				DefaultRegion:    "",
				DefaultZone:      "",
				UserAgent:        "",
				AncestryCache:    map[string]string{},
			})
			if err != nil {
				t.Fatalf("WithoutProject: Convert(%s, offline): %v", planfile, err)
			}
			expectedAssets := normalizeAssets(t, want, true)
			actualAssets := normalizeAssets(t, got, true)
			if diff := cmp.Diff(expectedAssets, actualAssets); diff != "" {
				t.Errorf("%v diff(-want, +got):\n%s", t.Name(), diff)
			}
		})
	}
}
