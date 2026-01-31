module "cloud_run_service" {
  source                 = "../../modules/services"
  project_region         = var.project_region
  project_id             = var.project_id
  container_image        = var.container_image
  service_account_name   = var.service_account_name
  min_instance_count     = var.min_instance_count
  max_instance_count     = var.max_instance_count
  memory_limit           = var.memory_limit
  cpu_limit              = var.cpu_limit
  app_env                = var.app_env
  http_port              = var.http_port
  template_id_secret_key = var.template_id_secret_key
  output_uri_secret_key  = var.output_uri_secret_key
}

module "iam" {
  source               = "../../modules/iam"
  service_account_name = var.service_account_name
}

module "storage" {
  source = "../../modules/storage"
}

module "eventarc_trigger" {
  source                 = "../../modules/triggers"
  project_region         = var.project_region
  service_name           = module.cloud_run_service.trancoder_service_name
  raw_videos_bucket_name = module.storage.raw_videos_bucket_name
  service_account_name   = module.iam.service_account_name
}
