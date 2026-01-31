output "trancoder_service_name" {
  description = "The name of the Cloud Run service used for transcoding."
  value       = google_cloud_run_service.apex_cr_start_transcode.name
}
