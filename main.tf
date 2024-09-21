# core.tf
terraform {
  backend "gcs" {
    bucket  = "brucearctor-demo-ingestion-tf-state"
    prefix  = "tofu/state"
  }

  required_version = ">= 1.8"

  required_providers {
  google = {
    source  = "hashicorp/google"
    version = "~> 6.3"
    }
  }
}

provider "google" {
  project = var.gcp_project_id
  region  = var.gcp_region
}



#=========================

# variables.tf
variable "gcp_project_id" {
  description = "The ID of the project in which to provision resources."
  type        = string
  # using default for now ...
  # will want tfvars
  default     = "brucearctor-demo-ingestion"
}

variable "gcp_region" {
  description = "The region in which to provision resources."
  type        = string
  # using default for now ...
  # will want tfvars
  default     = "us-central1"
}

#=========================



# etc

resource "google_pubsub_topic" "demo" {
  name = "demo-topic"
  # TODO:  SCHEMA
  # TODO:  Encrypt
}




variable "gcp_service_list" {
  description ="The list of apis necessary for the project"
  type = list(string)
  default = [
    "firestore.googleapis.com",
    "artifactregistry.googleapis.com"
  ]
}

resource "google_project_service" "gcp_services" {
  for_each = toset(var.gcp_service_list)
  disable_on_destroy = false
  service = each.key
}


# use following if tofu destroy and need to recreate
# tofu import google_firestore_database.database demo-ingestion
resource "google_firestore_database" "database" {
  name        = "demo-ingestion"
  location_id = "nam5"
  type        = "FIRESTORE_NATIVE"
}



resource "google_container_registry" "registry" {
  location = "US"
}


# resource "google_cloud_run_service" "default" {
#   name     = "cloudrun-srv"
#   location = "us-central1"

#   template {
#     spec {
#       containers {
#         image = "us-docker.pkg.dev/cloudrun/container/hello"
#       }
#     }
#   }
# }

# data "google_iam_policy" "noauth" {
#   binding {
#     role = "roles/run.invoker"
#     members = [
#       "allUsers",
#     ]
#   }
# }

# resource "google_cloud_run_service_iam_policy" "noauth" {
#   location    = google_cloud_run_service.default.location
#   project     = google_cloud_run_service.default.project
#   service     = google_cloud_run_service.default.name

#   policy_data = data.google_iam_policy.noauth.policy_data
# }
