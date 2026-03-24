
resource "google_compute_network" "default" {
  name                    = "tf-test-mxc84acit3-network"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  name          = "tf-test-mxc84acit3-subnet"
  ip_cidr_range = "10.0.0.0/16"
  region        = "us-central1"
  network       = google_compute_network.default.id
}

resource "google_compute_subnetwork" "proxy" {
  name          = "tf-test-mxc84acit3-proxy-subnet"
  ip_cidr_range = "10.1.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.id
  purpose       = "REGIONAL_MANAGED_PROXY"
  role          = "ACTIVE"
}

resource "google_compute_health_check" "default" {
  name               = "tf-test-mxc84acit3-hc"
  check_interval_sec = 1
  timeout_sec        = 1
  http_health_check {
    port = "80"
  }
}

resource "google_compute_region_backend_service" "default" {
  name                  = "tf-test-mxc84acit3-backend"
  region                = "us-central1"
  health_checks         = [google_compute_health_check.default.id]
  load_balancing_scheme = "INTERNAL_MANAGED"
  protocol              = "HTTP"
}

resource "google_compute_region_url_map" "default" {
  name            = "tf-test-mxc84acit3-url-map"
  region          = "us-central1"
  default_service = google_compute_region_backend_service.default.id
}

resource "google_compute_region_target_http_proxy" "default" {
  name    = "tf-test-mxc84acit3-http-proxy"
  region  = "us-central1"
  url_map = google_compute_region_url_map.default.id
}

resource "google_compute_forwarding_rule" "foobar" {
  name                  = "tf-test-mxc84acit3"
  region                = "us-central1"
  network               = google_compute_network.default.id
  subnetwork            = google_compute_subnetwork.default.id
  load_balancing_scheme = "INTERNAL_MANAGED"
  target                = google_compute_region_target_http_proxy.default.id
  port_range            = "80"
  allow_global_access   = true

  depends_on = [google_compute_subnetwork.proxy]
}
