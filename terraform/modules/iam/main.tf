data "google_service_account" "eventarc_sa" {
  account_id = var.service_account_name
}
