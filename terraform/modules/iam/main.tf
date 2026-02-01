data "google_service_account" "eventarc_sa" {
  account_id = var.service_account_name
}

data "google_projects" "project" {
  filter = "projectId:${var.project_id}"
}

locals {
  project_number = data.google_project.project.number
  gcs_sa_email   = "service-${local.project_number}@gcp-sa-cloud-storage.iam.gserviceaccount.com"
}

resource "google_project_iam_member" "gcs_pubsub_publisher" {
  project = var.project_id
  role    = "roles/pubsub.publisher"
  member  = "serviceAccount:${local.gcs_sa_email}"
}
