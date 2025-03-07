resource "google_project" "678722381013" {
  project_id          = "tf-test-tgc"
  name                = "tf-test-tgc"
  org_id              = "529579013760"
  billing_account     = "01129A-55361F-811C70" # requires billing to enable compute API
  deletion_policy = "DELETE"
  auto_create_network = false
}
