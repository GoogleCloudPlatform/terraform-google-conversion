resource "google_compute_instance" "test1" {
  attached_disk {
    device_name       = "test-device_name"
    kms_key_self_link = "test-kms_key_self_link"
    mode              = "READ_ONLY"
    source            = "projects/test-project/zones/us-central1-a/disks/test-source"
  }

  attached_disk {
    mode   = "READ_WRITE"
    source = "projects/test-project/zones/us-central1-a/disks/test-source2"
  }

  boot_disk {
    auto_delete             = true
    device_name             = "test-device_name"
    disk_encryption_key_raw = "test-disk_encryption_key_raw"
    mode                    = "READ_WRITE"
    source                  = "test-source"
  }

  can_ip_forward      = true
  deletion_protection = true
  description         = "test-description"

  guest_accelerator {
    count = 42
    type  = "test-guest_accelerator-type1"
  }

  guest_accelerator {
    count = 42
    type  = "test-guest_accelerator-type2"
  }

  hostname = "test-hostname"

  labels = {
    label_foo1 = "label-bar1"
  }

  machine_type = "n1-standard-1"

  metadata = {
    metadata_foo1 = "metadata-bar1"
  }

  min_cpu_platform = "test-min_cpu_platform"
  name             = "test1"

  network_interface {
    access_config {
      nat_ip = "192.168.0.42"
    }

    access_config {
      network_tier = "STANDARD"
    }

    access_config {
      public_ptr_domain_name = "test-public_ptr_domain_name"
    }

    alias_ip_range {
      ip_cidr_range         = "test-ip_cidr_range"
      subnetwork_range_name = "test-subnetwork_range_name"
    }

    network     = "default"
    network_ip  = "test-network_ip"
    queue_count = 0
  }

  network_interface {
    queue_count = 0
    subnetwork  = "projects/test-subnetwork_project/regions/us-central1/subnetworks/test-subnetwork"
  }

  scheduling {
    automatic_restart   = true
    on_host_maintenance = "test-on_host_maintenance"
    preemptible         = true
  }

  scratch_disk {
    interface = "SCSI"
  }

  scratch_disk {
    interface = "SCSI"
  }

  service_account {
    email  = "test-email"
    scopes = ["cloud-platform"]
  }

  shielded_instance_config {
    enable_integrity_monitoring = true
    enable_secure_boot          = true
    enable_vtpm                 = true
  }

  tags = ["bar", "foo"]
  zone = "us-central1-a"
}
resource "google_compute_instance" "test2" {
  boot_disk {
    auto_delete       = true
    kms_key_self_link = "test-kms_key_self_link"
    mode              = "READ_WRITE"
  }

  can_ip_forward      = false
  deletion_protection = false
  machine_type        = "n1-standard-1"
  name                = "test2"

  network_interface {
    network     = "default"
    queue_count = 0
  }

  scheduling {
    automatic_restart = true
    preemptible       = false
  }

  zone = "us-central1-a"
}
