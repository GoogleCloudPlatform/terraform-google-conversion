package hcl

import (
	"fmt"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v2/helper"

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

	accel, err := convertInstanceGuestAccelerators(instance.GuestAccelerators)
	if err != nil {
		return "", nil, err
	}

	networkInterfaces, err := convertNetworkInterfaces(instance.NetworkInterfaces)
	if err != nil {
		return "", nil, err
	}
	bootDisks, scratchDisks, attachedDisks, err := convertDisks(instance.Disks)
	if err != nil {
		return "", nil, err
	}
	hclData := make(map[string]interface{})
	hclData["can_ip_forward"] = instance.CanIpForward
	hclData["description"] = instance.Description
	hclData["boot_disk"] = bootDisks
	hclData["scratch_disk"] = scratchDisks
	hclData["attached_disk"] = attachedDisks
	hclData["machine_type"] = helper.ParseFieldValue(instance.MachineType, "machineTypes")
	hclData["name"] = instance.Name
	hclData["network_interface"] = networkInterfaces
	hclData["tags"] = instance.Tags.Items
	hclData["tags_fingerprint"] = instance.Tags.Fingerprint
	hclData["labels"] = instance.Labels
	hclData["service_account"] = convertServiceAccounts(instance.ServiceAccounts)
	hclData["guest_accelerator"] = accel
	hclData["min_cpu_platform"] = instance.MinCpuPlatform
	hclData["scheduling"] = convertScheduling(instance.Scheduling)
	hclData["deletion_protection"] = instance.DeletionProtection
	hclData["hostname"] = instance.Hostname
	hclData["shielded_instance_config"] = convertShieldedInstanceConfig(instance.ShieldedInstanceConfig)
	hclData["enable_display"] = convertDisplayDevice(instance.DisplayDevice)

	if instance.Zone == "" {
		instance.Zone = helper.ParseFieldValue(asset.Name, "zones")
	}
	hclData["zone"] = instance.Zone
	metadata := convertMetadata(instance.Metadata)
	for k, v := range metadata {
		hclData[k] = v
	}
	return instance.Name, hclData, nil
}

func convertDisks(disks []*compute.AttachedDisk) (bootDisks []map[string]interface{}, scratchDisks []map[string]interface{}, attachedDisks []map[string]interface{}, err error) {
	for _, disk := range disks {
		if disk.Boot {
			bootDisk, err := convertBootDisk(disk)
			if err != nil {
				return nil, nil, nil, err
			}
			bootDisks = append(bootDisks, bootDisk)
			continue
		}
		if disk.Type == "SCRATCH" {
			scratchDisk, err := convertScratchDisk(disk)
			if err != nil {
				return nil, nil, nil, err
			}
			scratchDisks = append(scratchDisks, scratchDisk)
			continue
		}
		attachedDisk, err := convertAttachedDisk(disk)
		if err != nil {
			return nil, nil, nil, err
		}
		attachedDisks = append(attachedDisks, attachedDisk)
	}
	return
}

func convertBootDisk(disk *compute.AttachedDisk) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"auto_delete": disk.AutoDelete,
		"device_name": disk.DeviceName,
		"source":      helper.ParseFieldValue(disk.Source, "disks"),
		// "initialize_params": map[string]interface{}{
		// 	"size":  disk.InitializeParams.DiskSizeGb,
		// 	"type":  helper.ParseFieldValue(disk.InitializeParams.DiskType, "diskTypes"),
		// 	"image": disk.InitializeParams.SourceImage,
		// },
		"mode": disk.Mode,
	}
	if disk.DiskEncryptionKey != nil {
		data["disk_encryption_key_raw"] = disk.DiskEncryptionKey.RawKey
		data["kms_key_self_link"] = disk.DiskEncryptionKey.KmsKeyName
	}
	return data, nil
}

func convertScratchDisk(disk *compute.AttachedDisk) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"interface": disk.Interface,
	}
	return data, nil
}

func convertAttachedDisk(disk *compute.AttachedDisk) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"source":      disk.Source,
		"mode":        disk.Mode,
		"device_name": disk.DeviceName,
	}
	if disk.DiskEncryptionKey != nil {
		data["disk_encryption_key_raw"] = disk.DiskEncryptionKey.RawKey
		data["kms_key_self_link"] = disk.DiskEncryptionKey.KmsKeyName
	}
	return data, nil
}

func convertInstanceGuestAccelerators(configs []*compute.AcceleratorConfig) ([]map[string]interface{}, error) {
	arr := make([]map[string]interface{}, len(configs))
	for ix, config := range configs {
		arr[ix] = make(map[string]interface{})
		arr[ix]["count"] = config.AcceleratorCount
		arr[ix]["type"] = helper.ParseFieldValue(config.AcceleratorType, "acceleratorTypes")
	}
	return arr, nil
}

func convertNetworkInterfaces(interfaces []*compute.NetworkInterface) ([]map[string]interface{}, error) {
	arr := make([]map[string]interface{}, len(interfaces))
	for ix, item := range interfaces {
		arr[ix] = make(map[string]interface{})
		arr[ix]["network"] = helper.ParseFieldValue(item.Network, "networks")
		arr[ix]["subnetwork"] = item.Subnetwork
		arr[ix]["network_ip"] = item.NetworkIP
		arr[ix]["nic_type"] = item.NicType
		arr[ix]["queue_count"] = item.QueueCount
		arr[ix]["ipv6_access_config"] = convertIpv6AccessConfigs(item.Ipv6AccessConfigs)
		arr[ix]["access_config"] = convertAccessConfigs(item.AccessConfigs)
		arr[ix]["alias_ip_range"] = convertAliasIpRanges(item.AliasIpRanges)
		arr[ix]["stack_type"] = item.StackType
	}
	return arr, nil
}

func convertAliasIpRanges(ranges []*compute.AliasIpRange) []map[string]interface{} {
	arr := make([]map[string]interface{}, len(ranges))
	for ix, item := range ranges {
		arr[ix] = make(map[string]interface{})
		arr[ix]["ip_cidr_range"] = item.IpCidrRange
		arr[ix]["subnetwork_range_name"] = item.SubnetworkRangeName
	}
	return arr
}

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

func convertAccessConfigs(configs []*compute.AccessConfig) []map[string]interface{} {
	arr := make([]map[string]interface{}, len(configs))
	for ix, item := range configs {
		arr[ix] = make(map[string]interface{})
		arr[ix]["nat_ip"] = item.NatIP
		arr[ix]["network_tier"] = item.NetworkTier
		if len(item.PublicPtrDomainName) > 0 {
			arr[ix]["public_ptr_domain_name"] = item.PublicPtrDomainName
		}
	}
	return arr
}

func convertIpv6AccessConfigs(configs []*compute.AccessConfig) []map[string]interface{} {
	arr := make([]map[string]interface{}, len(configs))
	for ix, item := range configs {
		arr[ix] = make(map[string]interface{})
		arr[ix]["network_tier"] = item.NetworkTier
		if len(item.PublicPtrDomainName) > 0 {
			arr[ix]["public_ptr_domain_name"] = item.PublicPtrDomainName
		}
	}
	return arr
}

func convertServiceAccounts(accounts []*compute.ServiceAccount) []map[string]interface{} {
	arr := make([]map[string]interface{}, len(accounts))
	for ix, item := range accounts {
		arr[ix] = make(map[string]interface{})
		arr[ix]["email"] = item.Email
		var scopes []string
		for _, url := range item.Scopes {
			scope := helper.ConvertServiceScope(url)
			scopes = append(scopes, scope)
		}
		arr[ix]["scopes"] = scopes
	}
	return arr
}

func convertShieldedInstanceConfig(item *compute.ShieldedInstanceConfig) []map[string]interface{} {
	if item == nil {
		return nil
	}
	return []map[string]interface{}{
		{
			"enable_secure_boot":          item.EnableSecureBoot,
			"enable_vtpm":                 item.EnableVtpm,
			"enable_integrity_monitoring": item.EnableIntegrityMonitoring,
		},
	}
}

func convertConfidentialInstanceConfig(item *compute.ConfidentialInstanceConfig) (map[string]interface{}, error) {
	if item == nil {
		return nil, nil
	}
	return map[string]interface{}{
		"enable_confidential_compute": item.EnableConfidentialCompute,
	}, nil
}

func convertAdvancedMachineFeatures(item *compute.AdvancedMachineFeatures) (map[string]interface{}, error) {
	if item == nil {
		return nil, nil
	}
	return map[string]interface{}{
		"enable_nested_virtualization": item.EnableNestedVirtualization,
		"threads_per_core":             item.ThreadsPerCore,
	}, nil
}

func convertDisplayDevice(item *compute.DisplayDevice) map[string]interface{} {
	if item == nil {
		return nil
	}
	return map[string]interface{}{
		"enable_display": item.EnableDisplay,
	}
}

func convertMetadata(metadata *compute.Metadata) map[string]interface{} {
	data := map[string]interface{}{
		"metadata_fingerprint": metadata.Fingerprint,
	}
	kv := map[string]interface{}{}
	for _, item := range metadata.Items {
		if item.Value == nil {
			continue
		}
		if item.Key == "startup-script" {
			data[item.Key] = *item.Value
		} else {
			kv[item.Key] = *item.Value
		}
	}
	data["metadata"] = kv
	return data
}

// func resourceInstanceMetadata(d TerraformResourceData) (*compute.Metadata, error) {
// 	m := &compute.Metadata{}
// 	mdMap := d.Get("metadata").(map[string]interface{})
// 	if v, ok := d.GetOk("metadata_startup_script"); ok && v.(string) != "" {
// 		if _, ok := mdMap["startup-script"]; ok {
// 			return nil, errors.New("Cannot provide both metadata_startup_script and metadata.startup-script.")
// 		}
// 		mdMap["startup-script"] = v
// 	}
// 	if len(mdMap) > 0 {
// 		m.Items = make([]*compute.MetadataItems, 0, len(mdMap))
// 		var keys []string
// 		for k := range mdMap {
// 			keys = append(keys, k)
// 		}
// 		sort.Strings(keys)
// 		for _, k := range keys {
// 			v := mdMap[k].(string)
// 			m.Items = append(m.Items, &compute.MetadataItems{
// 				Key:   k,
// 				Value: &v,
// 			})
// 		}

// 		// Set the fingerprint. If the metadata has never been set before
// 		// then this will just be blank.
// 		m.Fingerprint = d.Get("metadata_fingerprint").(string)
// 	}

// 	return m, nil
// }
