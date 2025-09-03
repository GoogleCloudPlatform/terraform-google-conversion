package provider

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/services/alloydb"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/services/apigee"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/services/apphub"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/services/backupdr"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/services/beyondcorp"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/services/bigquery"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/services/binaryauthorization"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/services/blockchainnodeengine"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/services/certificatemanager"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/services/cloudasset"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/services/cloudbuild"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/services/cloudbuildv2"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/services/cloudfunctions2"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/services/compute"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/services/networksecurity"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/services/pubsub"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/services/resourcemanager"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var handwrittenTfplan2caiResources = map[string]*schema.Resource{
	// ####### START handwritten resources ###########
	"google_compute_instance": compute.ResourceComputeInstance(),
	"google_project":          resourcemanager.ResourceGoogleProject(),
	// ####### END handwritten resources ###########
}

// Generated resources: 41
var generatedResources = map[string]*schema.Resource{
	"google_alloydb_backup":                          alloydb.ResourceAlloydbBackup(),
	"google_alloydb_cluster":                         alloydb.ResourceAlloydbCluster(),
	"google_alloydb_instance":                        alloydb.ResourceAlloydbInstance(),
	"google_apigee_instance":                         apigee.ResourceApigeeInstance(),
	"google_apphub_application":                      apphub.ResourceApphubApplication(),
	"google_apphub_service":                          apphub.ResourceApphubService(),
	"google_apphub_workload":                         apphub.ResourceApphubWorkload(),
	"google_backup_dr_backup_plan":                   backupdr.ResourceBackupDRBackupPlan(),
	"google_backup_dr_backup_vault":                  backupdr.ResourceBackupDRBackupVault(),
	"google_beyondcorp_app_connection":               beyondcorp.ResourceBeyondcorpAppConnection(),
	"google_beyondcorp_app_connector":                beyondcorp.ResourceBeyondcorpAppConnector(),
	"google_beyondcorp_app_gateway":                  beyondcorp.ResourceBeyondcorpAppGateway(),
	"google_bigquery_dataset":                        bigquery.ResourceBigQueryDataset(),
	"google_binary_authorization_attestor":           binaryauthorization.ResourceBinaryAuthorizationAttestor(),
	"google_blockchain_node_engine_blockchain_nodes": blockchainnodeengine.ResourceBlockchainNodeEngineBlockchainNodes(),
	"google_certificate_manager_certificate":         certificatemanager.ResourceCertificateManagerCertificate(),
	"google_cloud_asset_folder_feed":                 cloudasset.ResourceCloudAssetFolderFeed(),
	"google_cloud_asset_organization_feed":           cloudasset.ResourceCloudAssetOrganizationFeed(),
	"google_cloud_asset_project_feed":                cloudasset.ResourceCloudAssetProjectFeed(),
	"google_cloudbuild_bitbucket_server_config":      cloudbuild.ResourceCloudBuildBitbucketServerConfig(),
	"google_cloudbuildv2_connection":                 cloudbuildv2.ResourceCloudbuildv2Connection(),
	"google_cloudbuildv2_repository":                 cloudbuildv2.ResourceCloudbuildv2Repository(),
	"google_cloudfunctions2_function":                cloudfunctions2.ResourceCloudfunctions2function(),
	"google_compute_address":                         compute.ResourceComputeAddress(),
	"google_compute_autoscaler":                      compute.ResourceComputeAutoscaler(),
	"google_compute_backend_bucket":                  compute.ResourceComputeBackendBucket(),
	"google_compute_backend_service":                 compute.ResourceComputeBackendService(),
	"google_compute_disk":                            compute.ResourceComputeDisk(),
	"google_compute_external_vpn_gateway":            compute.ResourceComputeExternalVpnGateway(),
	"google_compute_firewall":                        compute.ResourceComputeFirewall(),
	"google_compute_firewall_policy":                 compute.ResourceComputeFirewallPolicy(),
	"google_compute_global_address":                  compute.ResourceComputeGlobalAddress(),
	"google_compute_health_check":                    compute.ResourceComputeHealthCheck(),
	"google_compute_image":                           compute.ResourceComputeImage(),
	"google_compute_network":                         compute.ResourceComputeNetwork(),
	"google_compute_region_autoscaler":               compute.ResourceComputeRegionAutoscaler(),
	"google_compute_subnetwork":                      compute.ResourceComputeSubnetwork(),
	"google_compute_url_map":                         compute.ResourceComputeUrlMap(),
	"google_network_security_server_tls_policy":      networksecurity.ResourceNetworkSecurityServerTlsPolicy(),
	"google_pubsub_subscription":                     pubsub.ResourcePubsubSubscription(),
	"google_pubsub_topic":                            pubsub.ResourcePubsubTopic(),
}
