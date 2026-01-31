variable "project_id" {
  type        = string
  description = "The unique identifier for the GCP project for resource organization and billing."
  validation {
    condition     = length(var.project_id) > 0
    error_message = "The project_id must not be empty."
  }
}

variable "project_region" {
  type        = string
  description = "The GCP region where the resources will be deployed, impacting latency and compliance."
  validation {
    condition     = length(var.project_region) > 0
    error_message = "The project_region must be specified."
  }
}

variable "service_name" {
  description = "The name of the Cloud Run service."
  type        = string
  default     = "apex-cr-start-transcode"
}

variable "container_image" {
  description = "The container image to be used for the Cloud Run service."
  type        = string
}

variable "service_account_name" {
  description = "The service account name to be used by the Cloud Run service."
  type        = string
}

variable "min_instance_count" {
  description = "The minimum number of instances for the Cloud Run service."
  type        = number
  default     = 0
}

variable "max_instance_count" {
  description = "The maximum number of instances for the Cloud Run service."
  type        = number
  default     = 2
}

variable "memory_limit" {
  description = "The memory limit for the Cloud Run service."
  type        = string
  default     = "512Mi"
}

variable "cpu_limit" {
  description = "The CPU limit for the Cloud Run service."
  type        = string
  default     = "1"
}

variable "app_env" {
  description = "The application environment variable."
  type        = string
  default     = "production"
}

variable "http_port" {
  description = "The HTTP port for the Cloud Run service."
  type        = string
  default     = "8080"
}

variable "template_id_secret_name" {
  description = "The name of the secret containing the template ID."
  type        = string
}

variable "template_id_secret_key" {
  description = "The key within the secret that contains the template ID."
  type        = string
}

variable "output_uri_secret_name" {
  description = "The name of the secret containing the output URI."
  type        = string
}

variable "output_uri_secret_key" {
  description = "The key within the secret that contains the output URI."
  type        = string
}
