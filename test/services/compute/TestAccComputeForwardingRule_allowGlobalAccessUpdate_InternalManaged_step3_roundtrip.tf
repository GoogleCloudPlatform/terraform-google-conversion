resource "google_compute_forwarding_rule" "tf-test-mxc84acit3" {
  allow_global_access     = true
  allow_psc_global_access = false
  ip_address              = "10.0.0.3"
  ip_protocol             = "TCP"
  load_balancing_scheme   = "INTERNAL_MANAGED"
  name                    = "tf-test-mxc84acit3"
  network                 = "projects/ci-test-project-nightly-ga/global/networks/tf-test-mxc84acit3-network"
  network_tier            = "PREMIUM"
  no_automate_dns_zone    = false
  port_range              = "80-80"
  project                 = "ci-test-project-nightly-ga"
  region                  = "us-central1"
  subnetwork              = "projects/ci-test-project-nightly-ga/regions/us-central1/subnetworks/tf-test-mxc84acit3-subnet"
  target                  = "https://www.googleapis.com/compute/v1/projects/ci-test-project-nightly-ga/regions/us-central1/targetHttpProxies/tf-test-mxc84acit3-http-proxy"
}
