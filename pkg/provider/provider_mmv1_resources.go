package provider

import (
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/resourcemanager"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/alloydb"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/apigee"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/apphub"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/artifactregistry"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/backupdr"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/beyondcorp"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/bigquery"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/binaryauthorization"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/blockchainnodeengine"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/certificatemanager"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/cloudasset"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/cloudbuild"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/cloudbuildv2"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/clouddeploy"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/cloudfunctions2"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/compute"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/dataproc"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/datastream"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/filestore"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/firebasedataconnect"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/iambeta"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/kms"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/networksecurity"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/pubsub"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/redis"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/secretmanager"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/secretmanagerregional"
)

var handwrittenTfplan2caiResources = map[string]*schema.Resource{
	// ####### START handwritten resources ###########
	"google_compute_instance": compute.ResourceComputeInstance(),
	"google_project":          resourcemanager.ResourceGoogleProject(),
	// ####### END handwritten resources ###########
}

// Generated resources: 72
var generatedResources = map[string]*schema.Resource{
	"google_alloydb_backup":                                alloydb.ResourceAlloydbBackup(),
	"google_alloydb_cluster":                               alloydb.ResourceAlloydbCluster(),
	"google_alloydb_instance":                              alloydb.ResourceAlloydbInstance(),
	"google_apigee_instance":                               apigee.ResourceApigeeInstance(),
	"google_apphub_application":                            apphub.ResourceApphubApplication(),
	"google_apphub_service":                                apphub.ResourceApphubService(),
	"google_apphub_service_project_attachment":             apphub.ResourceApphubServiceProjectAttachment(),
	"google_apphub_workload":                               apphub.ResourceApphubWorkload(),
	"google_artifact_registry_repository":                  artifactregistry.ResourceArtifactRegistryRepository(),
	"google_backup_dr_backup_plan":                         backupdr.ResourceBackupDRBackupPlan(),
	"google_backup_dr_backup_vault":                        backupdr.ResourceBackupDRBackupVault(),
	"google_beyondcorp_app_connection":                     beyondcorp.ResourceBeyondcorpAppConnection(),
	"google_beyondcorp_app_connector":                      beyondcorp.ResourceBeyondcorpAppConnector(),
	"google_beyondcorp_app_gateway":                        beyondcorp.ResourceBeyondcorpAppGateway(),
	"google_bigquery_dataset":                              bigquery.ResourceBigQueryDataset(),
	"google_binary_authorization_attestor":                 binaryauthorization.ResourceBinaryAuthorizationAttestor(),
	"google_blockchain_node_engine_blockchain_nodes":       blockchainnodeengine.ResourceBlockchainNodeEngineBlockchainNodes(),
	"google_certificate_manager_certificate":               certificatemanager.ResourceCertificateManagerCertificate(),
	"google_cloud_asset_folder_feed":                       cloudasset.ResourceCloudAssetFolderFeed(),
	"google_cloud_asset_organization_feed":                 cloudasset.ResourceCloudAssetOrganizationFeed(),
	"google_cloud_asset_project_feed":                      cloudasset.ResourceCloudAssetProjectFeed(),
	"google_cloudbuild_bitbucket_server_config":            cloudbuild.ResourceCloudBuildBitbucketServerConfig(),
	"google_cloudbuildv2_connection":                       cloudbuildv2.ResourceCloudbuildv2Connection(),
	"google_cloudbuildv2_repository":                       cloudbuildv2.ResourceCloudbuildv2Repository(),
	"google_clouddeploy_automation":                        clouddeploy.ResourceClouddeployAutomation(),
	"google_clouddeploy_custom_target_type":                clouddeploy.ResourceClouddeployCustomTargetType(),
	"google_clouddeploy_deploy_policy":                     clouddeploy.ResourceClouddeployDeployPolicy(),
	"google_cloudfunctions2_function":                      cloudfunctions2.ResourceCloudfunctions2function(),
	"google_compute_address":                               compute.ResourceComputeAddress(),
	"google_compute_autoscaler":                            compute.ResourceComputeAutoscaler(),
	"google_compute_backend_bucket":                        compute.ResourceComputeBackendBucket(),
	"google_compute_backend_service":                       compute.ResourceComputeBackendService(),
	"google_compute_disk":                                  compute.ResourceComputeDisk(),
	"google_compute_external_vpn_gateway":                  compute.ResourceComputeExternalVpnGateway(),
	"google_compute_firewall":                              compute.ResourceComputeFirewall(),
	"google_compute_firewall_policy":                       compute.ResourceComputeFirewallPolicy(),
	"google_compute_global_address":                        compute.ResourceComputeGlobalAddress(),
	"google_compute_health_check":                          compute.ResourceComputeHealthCheck(),
	"google_compute_image":                                 compute.ResourceComputeImage(),
	"google_compute_network":                               compute.ResourceComputeNetwork(),
	"google_compute_region_autoscaler":                     compute.ResourceComputeRegionAutoscaler(),
	"google_compute_region_health_check":                   compute.ResourceComputeRegionHealthCheck(),
	"google_compute_subnetwork":                            compute.ResourceComputeSubnetwork(),
	"google_compute_url_map":                               compute.ResourceComputeUrlMap(),
	"google_dataproc_batch":                                dataproc.ResourceDataprocBatch(),
	"google_datastream_connection_profile":                 datastream.ResourceDatastreamConnectionProfile(),
	"google_datastream_private_connection":                 datastream.ResourceDatastreamPrivateConnection(),
	"google_datastream_stream":                             datastream.ResourceDatastreamStream(),
	"google_filestore_backup":                              filestore.ResourceFilestoreBackup(),
	"google_filestore_instance":                            filestore.ResourceFilestoreInstance(),
	"google_filestore_snapshot":                            filestore.ResourceFilestoreSnapshot(),
	"google_firebase_data_connect_service":                 firebasedataconnect.ResourceFirebaseDataConnectService(),
	"google_iam_workload_identity_pool":                    iambeta.ResourceIAMBetaWorkloadIdentityPool(),
	"google_iam_workload_identity_pool_provider":           iambeta.ResourceIAMBetaWorkloadIdentityPoolProvider(),
	"google_kms_autokey_config":                            kms.ResourceKMSAutokeyConfig(),
	"google_kms_key_handle":                                kms.ResourceKMSKeyHandle(),
	"google_network_security_address_group":                networksecurity.ResourceNetworkSecurityAddressGroup(),
	"google_network_security_client_tls_policy":            networksecurity.ResourceNetworkSecurityClientTlsPolicy(),
	"google_network_security_gateway_security_policy":      networksecurity.ResourceNetworkSecurityGatewaySecurityPolicy(),
	"google_network_security_gateway_security_policy_rule": networksecurity.ResourceNetworkSecurityGatewaySecurityPolicyRule(),
	"google_network_security_security_profile":             networksecurity.ResourceNetworkSecuritySecurityProfile(),
	"google_network_security_security_profile_group":       networksecurity.ResourceNetworkSecuritySecurityProfileGroup(),
	"google_network_security_server_tls_policy":            networksecurity.ResourceNetworkSecurityServerTlsPolicy(),
	"google_network_security_url_lists":                    networksecurity.ResourceNetworkSecurityUrlLists(),
	"google_pubsub_subscription":                           pubsub.ResourcePubsubSubscription(),
	"google_pubsub_topic":                                  pubsub.ResourcePubsubTopic(),
	"google_redis_cluster":                                 redis.ResourceRedisCluster(),
	"google_redis_instance":                                redis.ResourceRedisInstance(),
	"google_secret_manager_secret":                         secretmanager.ResourceSecretManagerSecret(),
	"google_secret_manager_secret_version":                 secretmanager.ResourceSecretManagerSecretVersion(),
	"google_secret_manager_regional_secret":                secretmanagerregional.ResourceSecretManagerRegionalRegionalSecret(),
	"google_secret_manager_regional_secret_version":        secretmanagerregional.ResourceSecretManagerRegionalRegionalSecretVersion(),
}
