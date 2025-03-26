terraform {
  required_providers {
    google = {
      source = "hashicorp/google-beta"
      version = "~> {{.Provider.version}}"
    }
  }
}

provider "google" {
  {{if .Provider.credentials }}credentials = "{{.Provider.credentials}}"{{end}}
}

resource "random_string" "suffix" {
  length  = 10
  upper   = false
  special = false
}

resource "google_folder" "test-folder" {
  display_name = "test-folder-${random_string.suffix.result}"
  parent       = "organizations/{{.OrgID}}"
}

resource "google_project" "test-project" {
  folder_id       = google_folder.test-folder.name
  name       = "My Project ${random_string.suffix.result}"
  project_id = "test-project-${random_string.suffix.result}"
}

