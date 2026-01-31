module "cloud_run_service" {
  source                  = "./modules/cloud_run_service"
  service_name            = var.service_name
  project_region          = var.project_region
  min_instance_count      = var.min_instance_count
  max_instance_count      = var.max_instance_count
  service_account_name    = var.service_account_name
  container_image         = var.container_image
  memory_limit            = var.memory_limit
  cpu_limit               = var.cpu_limit
  app_env                 = var.app_env
  http_port               = var.http_port
  project_id              = var.project_id
  template_id_secret_name = var.template_id_secret_name
  template_id_secret_key  = var.template_id_secret_key
  output_uri_secret_name  = var.output_uri_secret_name
  output_uri_secret_key   = var.output_uri_secret_key
}
