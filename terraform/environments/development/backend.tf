terraform {
  backend "gcs" {
    bucket = "prj-apex-infra-terraform-state"
    prefix = "terraform/apex_cr_start_transcode/state"
  }
}
