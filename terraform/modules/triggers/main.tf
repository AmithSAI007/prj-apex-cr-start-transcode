resource "google_eventarc_trigger" "storage_trigger" {
  name     = var.trigger_name
  location = var.project_region
  matching_criteria {
    attribute = "type"
    value     = "google.cloud.storage.object.v1.finalized"
  }
  matching_criteria {
    attribute = "bucket"
    value     = google_storage_bucket.default.name
  }

  destination {
    cloud_run_service {
      service = var.service_name
      region  = var.project_region
    }
  }


}
