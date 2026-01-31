variable "project_region" {
  description = "The region where the Cloud Run service will be deployed."
  type        = string
  validation {
    condition     = length(var.project_region) > 0
    error_message = "The project_region must be specified."
  }
}

variable "service_name" {
  description = "The name of the Cloud Run service."
  type        = string
}

variable "trigger_name" {
  description = "The name of the Eventarc trigger."
  type        = string
  default     = "apex-transcoder-storage-trigger"
}

variable "raw_videos_bucket_name" {
  description = "The name of the Cloud Storage bucket to monitor for new video uploads."
  type        = string
}
