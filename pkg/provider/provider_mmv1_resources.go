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
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/cloudtasks"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/compute"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/databasemigrationservice"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/datafusion"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/dataproc"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/datastream"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/developerconnect"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/dialogflow"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/filestore"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/firebasedataconnect"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/gemini"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/gkebackup"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/gkehub"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/gkeonprem"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/iambeta"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/iap"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/identityplatform"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/kms"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/logging"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/looker"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/managedkafka"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/memcache"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/monitoring"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/netapp"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/networkconnectivity"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/networksecurity"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/privateca"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/pubsub"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/redis"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/secretmanager"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/secretmanagerregional"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/securesourcemanager"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/securitycenterv2"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/spanner"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/storageinsights"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/vmwareengine"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/vpcaccess"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v7/pkg/services/workbench"
)

var handwrittenTfplan2caiResources = map[string]*schema.Resource{
	// ####### START handwritten resources ###########
	"google_compute_instance": compute.ResourceComputeInstance(),
	"google_project":          resourcemanager.ResourceGoogleProject(),
	// ####### END handwritten resources ###########
}

// Generated resources: 136
var generatedResources = map[string]*schema.Resource{
	"google_alloydb_backup":                                  alloydb.ResourceAlloydbBackup(),
	"google_alloydb_cluster":                                 alloydb.ResourceAlloydbCluster(),
	"google_alloydb_instance":                                alloydb.ResourceAlloydbInstance(),
	"google_apigee_instance":                                 apigee.ResourceApigeeInstance(),
	"google_apphub_application":                              apphub.ResourceApphubApplication(),
	"google_apphub_service":                                  apphub.ResourceApphubService(),
	"google_apphub_service_project_attachment":               apphub.ResourceApphubServiceProjectAttachment(),
	"google_apphub_workload":                                 apphub.ResourceApphubWorkload(),
	"google_artifact_registry_repository":                    artifactregistry.ResourceArtifactRegistryRepository(),
	"google_backup_dr_backup_plan":                           backupdr.ResourceBackupDRBackupPlan(),
	"google_backup_dr_backup_vault":                          backupdr.ResourceBackupDRBackupVault(),
	"google_beyondcorp_app_connection":                       beyondcorp.ResourceBeyondcorpAppConnection(),
	"google_beyondcorp_app_connector":                        beyondcorp.ResourceBeyondcorpAppConnector(),
	"google_beyondcorp_app_gateway":                          beyondcorp.ResourceBeyondcorpAppGateway(),
	"google_bigquery_dataset":                                bigquery.ResourceBigQueryDataset(),
	"google_binary_authorization_attestor":                   binaryauthorization.ResourceBinaryAuthorizationAttestor(),
	"google_blockchain_node_engine_blockchain_nodes":         blockchainnodeengine.ResourceBlockchainNodeEngineBlockchainNodes(),
	"google_certificate_manager_certificate":                 certificatemanager.ResourceCertificateManagerCertificate(),
	"google_certificate_manager_certificate_issuance_config": certificatemanager.ResourceCertificateManagerCertificateIssuanceConfig(),
	"google_certificate_manager_certificate_map":             certificatemanager.ResourceCertificateManagerCertificateMap(),
	"google_cloud_asset_folder_feed":                         cloudasset.ResourceCloudAssetFolderFeed(),
	"google_cloud_asset_organization_feed":                   cloudasset.ResourceCloudAssetOrganizationFeed(),
	"google_cloud_asset_project_feed":                        cloudasset.ResourceCloudAssetProjectFeed(),
	"google_cloudbuild_bitbucket_server_config":              cloudbuild.ResourceCloudBuildBitbucketServerConfig(),
	"google_cloudbuildv2_connection":                         cloudbuildv2.ResourceCloudbuildv2Connection(),
	"google_cloudbuildv2_repository":                         cloudbuildv2.ResourceCloudbuildv2Repository(),
	"google_clouddeploy_automation":                          clouddeploy.ResourceClouddeployAutomation(),
	"google_clouddeploy_custom_target_type":                  clouddeploy.ResourceClouddeployCustomTargetType(),
	"google_clouddeploy_deploy_policy":                       clouddeploy.ResourceClouddeployDeployPolicy(),
	"google_cloudfunctions2_function":                        cloudfunctions2.ResourceCloudfunctions2function(),
	"google_cloud_tasks_queue":                               cloudtasks.ResourceCloudTasksQueue(),
	"google_compute_address":                                 compute.ResourceComputeAddress(),
	"google_compute_autoscaler":                              compute.ResourceComputeAutoscaler(),
	"google_compute_backend_bucket":                          compute.ResourceComputeBackendBucket(),
	"google_compute_backend_service":                         compute.ResourceComputeBackendService(),
	"google_compute_disk":                                    compute.ResourceComputeDisk(),
	"google_compute_external_vpn_gateway":                    compute.ResourceComputeExternalVpnGateway(),
	"google_compute_firewall":                                compute.ResourceComputeFirewall(),
	"google_compute_firewall_policy":                         compute.ResourceComputeFirewallPolicy(),
	"google_compute_global_address":                          compute.ResourceComputeGlobalAddress(),
	"google_compute_ha_vpn_gateway":                          compute.ResourceComputeHaVpnGateway(),
	"google_compute_health_check":                            compute.ResourceComputeHealthCheck(),
	"google_compute_image":                                   compute.ResourceComputeImage(),
	"google_compute_network":                                 compute.ResourceComputeNetwork(),
	"google_compute_node_group":                              compute.ResourceComputeNodeGroup(),
	"google_compute_node_template":                           compute.ResourceComputeNodeTemplate(),
	"google_compute_region_autoscaler":                       compute.ResourceComputeRegionAutoscaler(),
	"google_compute_region_health_check":                     compute.ResourceComputeRegionHealthCheck(),
	"google_compute_region_ssl_certificate":                  compute.ResourceComputeRegionSslCertificate(),
	"google_compute_resource_policy":                         compute.ResourceComputeResourcePolicy(),
	"google_compute_route":                                   compute.ResourceComputeRoute(),
	"google_compute_router":                                  compute.ResourceComputeRouter(),
	"google_compute_storage_pool":                            compute.ResourceComputeStoragePool(),
	"google_compute_subnetwork":                              compute.ResourceComputeSubnetwork(),
	"google_compute_url_map":                                 compute.ResourceComputeUrlMap(),
	"google_database_migration_service_migration_job":        databasemigrationservice.ResourceDatabaseMigrationServiceMigrationJob(),
	"google_data_fusion_instance":                            datafusion.ResourceDataFusionInstance(),
	"google_dataproc_autoscaling_policy":                     dataproc.ResourceDataprocAutoscalingPolicy(),
	"google_dataproc_batch":                                  dataproc.ResourceDataprocBatch(),
	"google_datastream_connection_profile":                   datastream.ResourceDatastreamConnectionProfile(),
	"google_datastream_private_connection":                   datastream.ResourceDatastreamPrivateConnection(),
	"google_datastream_stream":                               datastream.ResourceDatastreamStream(),
	"google_developer_connect_connection":                    developerconnect.ResourceDeveloperConnectConnection(),
	"google_developer_connect_git_repository_link":           developerconnect.ResourceDeveloperConnectGitRepositoryLink(),
	"google_dialogflow_agent":                                dialogflow.ResourceDialogflowAgent(),
	"google_filestore_backup":                                filestore.ResourceFilestoreBackup(),
	"google_filestore_instance":                              filestore.ResourceFilestoreInstance(),
	"google_filestore_snapshot":                              filestore.ResourceFilestoreSnapshot(),
	"google_firebase_data_connect_service":                   firebasedataconnect.ResourceFirebaseDataConnectService(),
	"google_gemini_code_repository_index":                    gemini.ResourceGeminiCodeRepositoryIndex(),
	"google_gemini_repository_group":                         gemini.ResourceGeminiRepositoryGroup(),
	"google_gke_backup_backup_plan":                          gkebackup.ResourceGKEBackupBackupPlan(),
	"google_gke_backup_restore_plan":                         gkebackup.ResourceGKEBackupRestorePlan(),
	"google_gke_hub_membership":                              gkehub.ResourceGKEHubMembership(),
	"google_gkeonprem_bare_metal_admin_cluster":              gkeonprem.ResourceGkeonpremBareMetalAdminCluster(),
	"google_gkeonprem_bare_metal_cluster":                    gkeonprem.ResourceGkeonpremBareMetalCluster(),
	"google_gkeonprem_bare_metal_node_pool":                  gkeonprem.ResourceGkeonpremBareMetalNodePool(),
	"google_gkeonprem_vmware_cluster":                        gkeonprem.ResourceGkeonpremVmwareCluster(),
	"google_gkeonprem_vmware_node_pool":                      gkeonprem.ResourceGkeonpremVmwareNodePool(),
	"google_iam_workload_identity_pool":                      iambeta.ResourceIAMBetaWorkloadIdentityPool(),
	"google_iam_workload_identity_pool_provider":             iambeta.ResourceIAMBetaWorkloadIdentityPoolProvider(),
	"google_iap_tunnel_dest_group":                           iap.ResourceIapTunnelDestGroup(),
	"google_identity_platform_default_supported_idp_config":  identityplatform.ResourceIdentityPlatformDefaultSupportedIdpConfig(),
	"google_identity_platform_inbound_saml_config":           identityplatform.ResourceIdentityPlatformInboundSamlConfig(),
	"google_identity_platform_tenant":                        identityplatform.ResourceIdentityPlatformTenant(),
	"google_kms_autokey_config":                              kms.ResourceKMSAutokeyConfig(),
	"google_kms_crypto_key":                                  kms.ResourceKMSCryptoKey(),
	"google_kms_crypto_key_version":                          kms.ResourceKMSCryptoKeyVersion(),
	"google_kms_ekm_connection":                              kms.ResourceKMSEkmConnection(),
	"google_kms_key_handle":                                  kms.ResourceKMSKeyHandle(),
	"google_kms_key_ring":                                    kms.ResourceKMSKeyRing(),
	"google_kms_key_ring_import_job":                         kms.ResourceKMSKeyRingImportJob(),
	"google_logging_metric":                                  logging.ResourceLoggingMetric(),
	"google_looker_instance":                                 looker.ResourceLookerInstance(),
	"google_managed_kafka_cluster":                           managedkafka.ResourceManagedKafkaCluster(),
	"google_memcache_instance":                               memcache.ResourceMemcacheInstance(),
	"google_monitoring_alert_policy":                         monitoring.ResourceMonitoringAlertPolicy(),
	"google_monitoring_notification_channel":                 monitoring.ResourceMonitoringNotificationChannel(),
	"google_monitoring_uptime_check_config":                  monitoring.ResourceMonitoringUptimeCheckConfig(),
	"google_netapp_active_directory":                         netapp.ResourceNetappActiveDirectory(),
	"google_netapp_backup":                                   netapp.ResourceNetappBackup(),
	"google_network_connectivity_group":                      networkconnectivity.ResourceNetworkConnectivityGroup(),
	"google_network_connectivity_policy_based_route":         networkconnectivity.ResourceNetworkConnectivityPolicyBasedRoute(),
	"google_network_security_address_group":                  networksecurity.ResourceNetworkSecurityAddressGroup(),
	"google_network_security_authz_policy":                   networksecurity.ResourceNetworkSecurityAuthzPolicy(),
	"google_network_security_client_tls_policy":              networksecurity.ResourceNetworkSecurityClientTlsPolicy(),
	"google_network_security_firewall_endpoint":              networksecurity.ResourceNetworkSecurityFirewallEndpoint(),
	"google_network_security_firewall_endpoint_association":  networksecurity.ResourceNetworkSecurityFirewallEndpointAssociation(),
	"google_network_security_gateway_security_policy":        networksecurity.ResourceNetworkSecurityGatewaySecurityPolicy(),
	"google_network_security_gateway_security_policy_rule":   networksecurity.ResourceNetworkSecurityGatewaySecurityPolicyRule(),
	"google_network_security_security_profile":               networksecurity.ResourceNetworkSecuritySecurityProfile(),
	"google_network_security_security_profile_group":         networksecurity.ResourceNetworkSecuritySecurityProfileGroup(),
	"google_network_security_server_tls_policy":              networksecurity.ResourceNetworkSecurityServerTlsPolicy(),
	"google_network_security_tls_inspection_policy":          networksecurity.ResourceNetworkSecurityTlsInspectionPolicy(),
	"google_network_security_url_lists":                      networksecurity.ResourceNetworkSecurityUrlLists(),
	"google_privateca_ca_pool":                               privateca.ResourcePrivatecaCaPool(),
	"google_privateca_certificate":                           privateca.ResourcePrivatecaCertificate(),
	"google_pubsub_subscription":                             pubsub.ResourcePubsubSubscription(),
	"google_pubsub_topic":                                    pubsub.ResourcePubsubTopic(),
	"google_redis_cluster":                                   redis.ResourceRedisCluster(),
	"google_redis_instance":                                  redis.ResourceRedisInstance(),
	"google_secret_manager_secret":                           secretmanager.ResourceSecretManagerSecret(),
	"google_secret_manager_secret_version":                   secretmanager.ResourceSecretManagerSecretVersion(),
	"google_secret_manager_regional_secret":                  secretmanagerregional.ResourceSecretManagerRegionalRegionalSecret(),
	"google_secret_manager_regional_secret_version":          secretmanagerregional.ResourceSecretManagerRegionalRegionalSecretVersion(),
	"google_secure_source_manager_instance":                  securesourcemanager.ResourceSecureSourceManagerInstance(),
	"google_scc_v2_folder_mute_config":                       securitycenterv2.ResourceSecurityCenterV2FolderMuteConfig(),
	"google_scc_v2_organization_mute_config":                 securitycenterv2.ResourceSecurityCenterV2OrganizationMuteConfig(),
	"google_spanner_database":                                spanner.ResourceSpannerDatabase(),
	"google_spanner_instance":                                spanner.ResourceSpannerInstance(),
	"google_storage_insights_report_config":                  storageinsights.ResourceStorageInsightsReportConfig(),
	"google_vmwareengine_network":                            vmwareengine.ResourceVmwareengineNetwork(),
	"google_vmwareengine_network_peering":                    vmwareengine.ResourceVmwareengineNetworkPeering(),
	"google_vmwareengine_network_policy":                     vmwareengine.ResourceVmwareengineNetworkPolicy(),
	"google_vpc_access_connector":                            vpcaccess.ResourceVPCAccessConnector(),
	"google_workbench_instance":                              workbench.ResourceWorkbenchInstance(),
}
