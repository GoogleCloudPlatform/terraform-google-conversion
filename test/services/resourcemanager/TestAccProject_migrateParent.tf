resource "google_project" "178044935624" {
  project_id = "tf-test-tgc"
  name       = "tf-test-tgc"
  deletion_policy = "DELETE"

  folder_id = google_folder.folder1.id
}

resource "google_folder" "folder1" {
  display_name = "tf-testtgc"
  parent       = "organizations/529579013760"
  deletion_protection = false
}
