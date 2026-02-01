variable "project_id" {
  type        = string
  description = "The unique identifier for the GCP project for resource organization and billing."
  validation {
    condition     = length(var.project_id) > 0
    error_message = "The project_id must not be empty."
  }
}

variable "service_account_name" {
  description = "The name of the Cloud Storage bucket to monitor for new video uploads."
  type        = string
  default     = "prj-apex-cr-start-transcode-sa"
}
