output "service_account_name" {
  description = "The name of the service account used by the Eventarc trigger."
  value       = data.google_service_account.eventarc_sa.email
}
