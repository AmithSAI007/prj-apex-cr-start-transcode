resource "google_cloud_run_service" "apex_cr_start_transcode" {
  name     = var.service_name
  location = var.project_region
  template {

    metadata {
      annotations = {
        "autoscaling.knative.dev/minScale" = tostring(var.min_instance_count)
        "autoscaling.knative.dev/maxScale" = tostring(var.max_instance_count)
      }
    }
    spec {
      service_account_name = var.service_account_name
      containers {
        image = var.container_image

        resources {
          limits = {
            memory = var.memory_limit
            cpu    = var.cpu_limit
          }
        }

        env {
          name  = "APP_ENV"
          value = var.app_env
        }
        env {
          name  = "HTTP_PORT"
          value = var.http_port
        }
        env {
          name  = "PROJECT_ID"
          value = var.project_id
        }
        env {
          name = "TEMPLATE_ID"
          value_from {
            secret_key_ref {
              name = var.template_id_secret_key
              key  = "latest"
            }
          }
        }
        env {
          name = "OUTPUT_URI"
          value_from {
            secret_key_ref {
              name = var.output_uri_secret_key
              key  = "latest"
            }
          }
        }
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}
