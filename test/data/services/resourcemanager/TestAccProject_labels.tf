provider "google" {
  add_terraform_attribution_label = false
}

resource "google_project" "810776189846" {
  project_id = "tf-test-tgc"
  name       = "tf-test-tgc"
  org_id     = "529579013760"
  deletion_policy = "DELETE"
  labels = {"test" = "that"
}
}
