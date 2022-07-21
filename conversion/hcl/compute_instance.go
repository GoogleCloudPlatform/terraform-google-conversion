package hcl

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"google.golang.org/api/compute/v1"
)

const ComputeInstanceAssetType string = "compute.googleapis.com/Instance"

func NewComputeInstanceConverter() *Converter {
	return &Converter{
		TFResourceName: "google_compute_instance",
		Convert:        convertComputeInstance,
	}
}

func convertComputeInstance(asset *Asset) (string, map[string]interface{}, error) {
	if asset == nil || asset.Resource == nil || asset.Resource.Data == nil {
		return "", nil, fmt.Errorf("asset does not provide enough data for conversion")
	}
	var instance *compute.Instance
	if err := helper.DecodeJSON(asset.Resource.Data, &instance); err != nil {
		return "", nil, err
	}

	bootDisks, scratchDisks, attachedDisks := convertDisks(instance.Disks)

	hclData := make(map[string]interface{})
	hclData["can_ip_forward"] = instance.CanIpForward
	hclData["description"] = instance.Description
	hclData["boot_disk"] = bootDisks
	hclData["scratch_disk"] = scratchDisks
	hclData["attached_disk"] = attachedDisks
	hclData["machine_type"] = helper.ParseFieldValue(instance.MachineType, "machineTypes")
	hclData["name"] = instance.Name
	hclData["network_interface"] = flattenNetworkInterfaces(instance.NetworkInterfaces)
	hclData["tags"] = instance.Tags.Items
	hclData["tags_fingerprint"] = instance.Tags.Fingerprint
	hclData["labels"] = instance.Labels
	hclData["service_account"] = flattenServiceAccounts(instance.ServiceAccounts)
	hclData["guest_accelerator"] = flattenGuestAccelerators(instance.GuestAccelerators)
	hclData["min_cpu_platform"] = instance.MinCpuPlatform
	hclData["scheduling"] = convertScheduling(instance.Scheduling)
	hclData["deletion_protection"] = instance.DeletionProtection
	hclData["hostname"] = instance.Hostname
	hclData["shielded_instance_config"] = flattenShieldedVmConfig(instance.ShieldedInstanceConfig)
	hclData["enable_display"] = flattenEnableDisplay(instance.DisplayDevice)
	hclData["metadata_fingerprint"] = instance.Metadata.Fingerprint
	hclData["metadata"] = convertMetadata(instance.Metadata)

	if instance.Zone == "" {
		instance.Zone = helper.ParseFieldValue(asset.Name, "zones")
	}
	hclData["zone"] = instance.Zone

	return instance.Name, hclData, nil
}

func convertDisks(disks []*compute.AttachedDisk) (bootDisks []map[string]interface{}, scratchDisks []map[string]interface{}, attachedDisks []map[string]interface{}) {
	for _, disk := range disks {
		if disk.Boot {
			bootDisks = append(bootDisks, convertBootDisk(disk))
			continue
		}
		if disk.Type == "SCRATCH" {
			scratchDisks = append(scratchDisks, flattenScratchDisk(disk))
			continue
		}
		attachedDisks = append(attachedDisks, convertAttachedDisk(disk))
	}
	return
}

func convertBootDisk(disk *compute.AttachedDisk) map[string]interface{} {
	data := map[string]interface{}{
		"auto_delete": disk.AutoDelete,
		"device_name": disk.DeviceName,
		"source":      disk.Source,
		"mode":        disk.Mode,
	}
	if disk.InitializeParams != nil {
		data["initialize_params"] = []map[string]interface{}{
			{
				"size":   disk.InitializeParams.DiskSizeGb,
				"type":   helper.ParseFieldValue(disk.InitializeParams.DiskType, "diskTypes"),
				"image":  disk.InitializeParams.SourceImage,
				"labels": disk.InitializeParams.Labels,
			},
		}
	}
	if disk.DiskEncryptionKey != nil {
		if disk.DiskEncryptionKey.RawKey != "" {
			data["disk_encryption_key_raw"] = disk.DiskEncryptionKey.RawKey
		}
		if disk.DiskEncryptionKey.Sha256 != "" {
			data["disk_encryption_key_sha256"] = disk.DiskEncryptionKey.Sha256
		}
		if disk.DiskEncryptionKey.KmsKeyName != "" {
			data["kms_key_self_link"] = disk.DiskEncryptionKey.KmsKeyName
		}
	}

	return data
}

func convertAttachedDisk(disk *compute.AttachedDisk) map[string]interface{} {
	data := map[string]interface{}{
		"source":      disk.Source,
		"mode":        disk.Mode,
		"device_name": disk.DeviceName,
	}
	if disk.DiskEncryptionKey != nil {
		if disk.DiskEncryptionKey.RawKey != "" {
			data["disk_encryption_key_raw"] = disk.DiskEncryptionKey.RawKey
		}
		if disk.DiskEncryptionKey.Sha256 != "" {
			data["disk_encryption_key_sha256"] = disk.DiskEncryptionKey.Sha256
		}
		if disk.DiskEncryptionKey.KmsKeyName != "" {
			data["kms_key_self_link"] = disk.DiskEncryptionKey.KmsKeyName
		}
	}
	return data
}

// func convertInstanceGuestAccelerators(configs []*compute.AcceleratorConfig) []map[string]interface{} {
// 	arr := make([]map[string]interface{}, len(configs))
// 	for ix, config := range configs {
// 		arr[ix] = make(map[string]interface{})
// 		arr[ix]["count"] = config.AcceleratorCount
// 		arr[ix]["type"] = helper.ParseFieldValue(config.AcceleratorType, "acceleratorTypes")
// 	}
// 	return arr
// }

// func convertNetworkInterfaces(interfaces []*compute.NetworkInterface) ([]map[string]interface{}, error) {
// 	arr := make([]map[string]interface{}, len(interfaces))
// 	for ix, item := range interfaces {
// 		arr[ix] = make(map[string]interface{})
// 		arr[ix]["network"] = item.Network
// 		arr[ix]["subnetwork"] = item.Subnetwork
// 		arr[ix]["network_ip"] = item.NetworkIP
// 		arr[ix]["nic_type"] = item.NicType
// 		arr[ix]["queue_count"] = item.QueueCount
// 		arr[ix]["ipv6_access_config"] = convertIpv6AccessConfigs(item.Ipv6AccessConfigs)
// 		arr[ix]["access_config"] = convertAccessConfigs(item.AccessConfigs)
// 		arr[ix]["alias_ip_range"] = convertAliasIpRanges(item.AliasIpRanges)
// 		arr[ix]["stack_type"] = item.StackType
// 	}
// 	return arr, nil
// }

// func convertAliasIpRanges(ranges []*compute.AliasIpRange) []map[string]interface{} {
// 	arr := make([]map[string]interface{}, len(ranges))
// 	for ix, item := range ranges {
// 		arr[ix] = make(map[string]interface{})
// 		arr[ix]["ip_cidr_range"] = item.IpCidrRange
// 		arr[ix]["subnetwork_range_name"] = item.SubnetworkRangeName
// 	}
// 	return arr
// }

func convertScheduling(sched *compute.Scheduling) []map[string]interface{} {
	data := map[string]interface{}{
		"automatic_restart":   sched.AutomaticRestart,
		"preemptible":         sched.Preemptible,
		"on_host_maintenance": sched.OnHostMaintenance,
		// node_affinities are not converted into cai
		"node_affinities": convertSchedulingNodeAffinity(sched.NodeAffinities),
	}
	if sched.MinNodeCpus > 0 {
		data["min_node_cpus"] = sched.MinNodeCpus
	}
	if len(sched.ProvisioningModel) > 0 {
		data["provisioning_model"] = sched.ProvisioningModel
	}
	return []map[string]interface{}{data}
}

func convertSchedulingNodeAffinity(items []*compute.SchedulingNodeAffinity) []map[string]interface{} {
	arr := make([]map[string]interface{}, len(items))
	for ix, item := range items {
		arr[ix] = make(map[string]interface{})
		arr[ix]["key"] = item.Key
		arr[ix]["operator"] = item.Operator
		arr[ix]["values"] = item.Values
	}
	return arr
}

// func convertAccessConfigs(configs []*compute.AccessConfig) []map[string]interface{} {
// 	arr := make([]map[string]interface{}, len(configs))
// 	for ix, item := range configs {
// 		arr[ix] = make(map[string]interface{})
// 		arr[ix]["nat_ip"] = item.NatIP
// 		arr[ix]["network_tier"] = item.NetworkTier
// 		if len(item.PublicPtrDomainName) > 0 {
// 			arr[ix]["public_ptr_domain_name"] = item.PublicPtrDomainName
// 		}
// 	}
// 	return arr
// }

// func convertIpv6AccessConfigs(configs []*compute.AccessConfig) []map[string]interface{} {
// 	arr := make([]map[string]interface{}, len(configs))
// 	for ix, item := range configs {
// 		arr[ix] = make(map[string]interface{})
// 		arr[ix]["network_tier"] = item.NetworkTier
// 		if len(item.PublicPtrDomainName) > 0 {
// 			arr[ix]["public_ptr_domain_name"] = item.PublicPtrDomainName
// 		}
// 	}
// 	return arr
// }

// func convertServiceAccounts(accounts []*compute.ServiceAccount) []map[string]interface{} {
// 	arr := make([]map[string]interface{}, len(accounts))
// 	for ix, item := range accounts {
// 		arr[ix] = make(map[string]interface{})
// 		arr[ix]["email"] = item.Email
// 		var scopes []string
// 		for _, url := range item.Scopes {
// 			scope := helper.ConvertServiceScope(url)
// 			scopes = append(scopes, scope)
// 		}
// 		arr[ix]["scopes"] = scopes
// 	}
// 	return arr
// }

// func convertShieldedInstanceConfig(item *compute.ShieldedInstanceConfig) []map[string]interface{} {
// 	if item == nil {
// 		return nil
// 	}
// 	return []map[string]interface{}{
// 		{
// 			"enable_secure_boot":          item.EnableSecureBoot,
// 			"enable_vtpm":                 item.EnableVtpm,
// 			"enable_integrity_monitoring": item.EnableIntegrityMonitoring,
// 		},
// 	}
// }

func convertMetadata(metadata *compute.Metadata) map[string]interface{} {
	md := flattenMetadata(metadata)

	// If the existing state contains "metadata_startup_script" instead of "metadata.startup-script",
	// we should move the remote metadata.startup-script to metadata_startup_script to avoid
	// specifying it in two places.
	if _, ok := md["metadata_startup_script"]; ok {
		md["metadata_startup_script"] = md["startup-script"]
		delete(md, "startup-script")
	}

	return md
}

func flattenMetadata(metadata *compute.Metadata) map[string]interface{} {
	metadataMap := make(map[string]interface{})
	for _, item := range metadata.Items {
		metadataMap[item.Key] = *item.Value
	}
	return metadataMap
}

func flattenGuestAccelerators(accelerators []*compute.AcceleratorConfig) []map[string]interface{} {
	acceleratorsSchema := make([]map[string]interface{}, len(accelerators))
	for i, accelerator := range accelerators {
		acceleratorsSchema[i] = map[string]interface{}{
			"count": accelerator.AcceleratorCount,
			"type":  accelerator.AcceleratorType,
		}
	}
	return acceleratorsSchema
}

func flattenShieldedVmConfig(shieldedVmConfig *compute.ShieldedInstanceConfig) []map[string]bool {
	if shieldedVmConfig == nil {
		return nil
	}

	return []map[string]bool{{
		"enable_secure_boot":          shieldedVmConfig.EnableSecureBoot,
		"enable_vtpm":                 shieldedVmConfig.EnableVtpm,
		"enable_integrity_monitoring": shieldedVmConfig.EnableIntegrityMonitoring,
	}}
}

func flattenEnableDisplay(displayDevice *compute.DisplayDevice) interface{} {
	if displayDevice == nil {
		return nil
	}

	return displayDevice.EnableDisplay
}

func flattenServiceAccounts(serviceAccounts []*compute.ServiceAccount) []map[string]interface{} {
	result := make([]map[string]interface{}, len(serviceAccounts))
	for i, serviceAccount := range serviceAccounts {
		result[i] = map[string]interface{}{
			"email":  serviceAccount.Email,
			"scopes": serviceAccount.Scopes,
		}
	}
	return result
}

func flattenNetworkInterfaces(networkInterfaces []*compute.NetworkInterface) []map[string]interface{} {
	flattened := make([]map[string]interface{}, len(networkInterfaces))
	for i, iface := range networkInterfaces {
		var ac []map[string]interface{}
		ac, _ = flattenAccessConfigs(iface.AccessConfigs)

		flattened[i] = map[string]interface{}{
			"network_ip":         iface.NetworkIP,
			"network":            iface.Network,
			"subnetwork":         iface.Subnetwork,
			"access_config":      ac,
			"alias_ip_range":     flattenAliasIpRange(iface.AliasIpRanges),
			"nic_type":           iface.NicType,
			"stack_type":         iface.StackType,
			"ipv6_access_config": flattenIpv6AccessConfigs(iface.Ipv6AccessConfigs),
			"queue_count":        iface.QueueCount,
		}
		// Instance template interfaces never have names, so they're absent
		// in the instance template network_interface schema. We want to use the
		// same flattening code for both resource types, so we avoid trying to
		// set the name field when it's not set at the GCE end.
		if iface.Name != "" {
			flattened[i]["name"] = iface.Name
		}
	}
	return flattened
}

func flattenIpv6AccessConfigs(ipv6AccessConfigs []*compute.AccessConfig) []map[string]interface{} {
	flattened := make([]map[string]interface{}, len(ipv6AccessConfigs))
	for i, ac := range ipv6AccessConfigs {
		flattened[i] = map[string]interface{}{
			"network_tier": ac.NetworkTier,
		}
		flattened[i]["public_ptr_domain_name"] = ac.PublicPtrDomainName
		flattened[i]["external_ipv6"] = ac.ExternalIpv6
	}
	return flattened
}

func flattenAccessConfigs(accessConfigs []*compute.AccessConfig) ([]map[string]interface{}, string) {
	flattened := make([]map[string]interface{}, len(accessConfigs))
	natIP := ""
	for i, ac := range accessConfigs {
		flattened[i] = map[string]interface{}{
			"nat_ip":       ac.NatIP,
			"network_tier": ac.NetworkTier,
		}
		if ac.SetPublicPtr {
			flattened[i]["public_ptr_domain_name"] = ac.PublicPtrDomainName
		}
		if natIP == "" {
			natIP = ac.NatIP
		}
	}
	return flattened, natIP
}

func flattenAliasIpRange(ranges []*compute.AliasIpRange) []map[string]interface{} {
	rangesSchema := make([]map[string]interface{}, 0, len(ranges))
	for _, ipRange := range ranges {
		rangesSchema = append(rangesSchema, map[string]interface{}{
			"ip_cidr_range":         ipRange.IpCidrRange,
			"subnetwork_range_name": ipRange.SubnetworkRangeName,
		})
	}
	return rangesSchema
}

func instanceSchedulingNodeAffinitiesElemSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"operator": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"IN", "NOT_IN"}, false),
			},
			"values": {
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func convertStringArrToInterface(strs []string) []interface{} {
	arr := make([]interface{}, len(strs))
	for i, str := range strs {
		arr[i] = str
	}
	return arr
}

func flattenScratchDisk(disk *compute.AttachedDisk) map[string]interface{} {
	result := map[string]interface{}{
		"interface": disk.Interface,
	}
	return result
}
