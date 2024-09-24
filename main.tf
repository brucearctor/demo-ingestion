# core.tf
terraform {
  backend "gcs" {
    bucket = "brucearctor-demo-ingestion-tf-state"
    prefix = "tofu/state"
  }

  required_version = ">= 1.8"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 6.3"
    }
    google-beta = {
      source  = "hashicorp/google-beta"
      version = "~> 5.0"
    }
  }
}

provider "google" {
  project = var.gcp_project_id
  region  = var.gcp_region
}

provider "google-beta" {
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
  default = "brucearctor-demo-ingestion"
}

variable "gcp_region" {
  description = "The region in which to provision resources."
  type        = string
  # using default for now ...
  # will want tfvars
  default = "us-central1"
}

variable "gcp_project_number" {
  description = "The project number"
  type        = string
  # using default for now ...
  # will want tfvars
  default = "909394038571"
}


#=========================



# etc

resource "google_pubsub_schema" "demo" {
  name       = "demo"
  type       = "PROTOCOL_BUFFER"
  definition = file("pubsub_proto_stripped/demo.proto_4_pubsub")
}

resource "google_pubsub_topic" "demo" {
  name = "demo-topic"

  depends_on = [google_pubsub_schema.demo]
  schema_settings {
    schema   = google_pubsub_schema.demo.id
    encoding = "BINARY"
  }
}





variable "gcp_service_list" {
  description = "The list of apis necessary for the project"
  type        = list(string)
  default = [
    "firestore.googleapis.com",
    "artifactregistry.googleapis.com",
    "cloudfunctions.googleapis.com",
    "cloudbuild.googleapis.com",
    "run.googleapis.com",
    "eventarc.googleapis.com"
  ]
}

resource "google_project_service" "gcp_services" {
  for_each           = toset(var.gcp_service_list)
  disable_on_destroy = false
  service            = each.key
}


resource "google_firebase_project" "demo" {
  provider = google-beta
}


# use following if tofu destroy and need to recreate
# tofu import google_firestore_database.database demo-ingestion
resource "google_firestore_database" "database" {
  name        = "demo-ingestion"
  location_id = "nam5"
  type        = "FIRESTORE_NATIVE"
  # project_id  = google_firebase_project.demo.project_id
}



resource "google_storage_bucket" "default" {
  name                        = "${var.gcp_project_id}-gcf-source"
  location                    = "US"
  uniform_bucket_level_access = true
}

data "archive_file" "collector" {
  type        = "zip"
  output_path = "/tmp/collector-function-source.zip"
  source_dir  = "collector"
}


resource "google_storage_bucket_object" "collector_object" {
  name   = "collector-function-source.zip"
  bucket = google_storage_bucket.default.name
  source = data.archive_file.collector.output_path
}

resource "google_cloudfunctions2_function" "collector" {
  name        = "collector"
  location    = "us-central1"
  description = "a collector function"

  build_config {
    runtime     = "go122"
    entry_point = "ReceiveAndPublish"
    # environment_variables = {
    #     GCP_PROJECT = var.gcp_project_id
    #     TOPIC = google_pubsub_topic.demo.name
    # }
    source {
      storage_source {
        bucket = google_storage_bucket.default.name
        object = google_storage_bucket_object.collector_object.name
      }
    }
  }

  # TODO: right size.  Took example/default values.
  service_config {
    max_instance_count = 1
    available_memory   = "256M"
    timeout_seconds    = 60
  }
}

resource "google_cloud_run_service_iam_member" "member" {
  location = google_cloudfunctions2_function.collector.location
  service  = google_cloudfunctions2_function.collector.name
  role     = "roles/run.invoker"
  member   = "allUsers"
}

output "function_uri" {
  value = google_cloudfunctions2_function.collector.service_config[0].uri
}

# // tofu import google_project.project brucearctor-demo-ingestion
# resource "google_project" "project" {
#   name = var.gcp_project_id
#   project_id = var.gcp_project_id
#   // deletion_policy = PREVENT
#   // deletion policy of PREVENT is default!
# }


resource "google_pubsub_subscription" "demo" {
  name  = "demo-subscription"
  topic = google_pubsub_topic.demo.id



  bigquery_config {
    use_topic_schema = true
    table            = "${google_bigquery_table.demo.project}.${google_bigquery_table.demo.dataset_id}.${google_bigquery_table.demo.table_id}"
  }

  depends_on = [google_project_iam_member.viewer, google_project_iam_member.editor]
}


resource "google_project_iam_member" "viewer" {
  project = var.gcp_project_id
  role    = "roles/bigquery.metadataViewer"
  member  = "serviceAccount:service-${var.gcp_project_number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}

resource "google_project_iam_member" "editor" {
  project = var.gcp_project_id
  role    = "roles/bigquery.dataEditor"
  member  = "serviceAccount:service-${var.gcp_project_number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}

resource "google_bigquery_dataset" "demo" {
  dataset_id = "demo"
}

resource "google_bigquery_table" "demo" {
  deletion_protection = false
  table_id            = "flight_status"
  dataset_id          = google_bigquery_dataset.demo.dataset_id

  schema = file("gen/bq/demo_ingest/flight_status.schema")
}



data "archive_file" "corefirestore" {
  type        = "zip"
  output_path = "/tmp/corefirestore-function-source.zip"
  source_dir  = "corefirestore"
}


resource "google_storage_bucket_object" "corefirestore_object" {
  name   = "corefirestore-function-source.zip"
  bucket = google_storage_bucket.default.name
  source = data.archive_file.corefirestore.output_path
}



resource "google_cloudfunctions2_function" "corefirestore" {
  depends_on  = [google_storage_bucket_object.collector_object]
  name        = "corefirestore"
  location    = "us-central1"
  description = "a core firestore function"

  build_config {
    runtime     = "go122"
    entry_point = "ReceivePushAndInsert"
    # environment_variables = {
    #     GCP_PROJECT = var.gcp_project_id
    #     TOPIC = google_pubsub_topic.demo.name
    # }
    source {
      storage_source {
        bucket = google_storage_bucket.default.name
        object = google_storage_bucket_object.corefirestore_object.name
      }
    }
  }

  # TODO: right size.  Took example/default values.
  service_config {
    max_instance_count = 1
    available_memory   = "256M"
    timeout_seconds    = 60
  }
}


resource "google_cloud_run_service_iam_member" "corefirestore_push_member" {
  depends_on = [google_service_account.demo_pubsub_invoker, google_cloudfunctions2_function.corefirestore]
  location   = google_cloudfunctions2_function.corefirestore.location
  service    = google_cloudfunctions2_function.corefirestore.name
  role       = "roles/run.invoker"
  member     = "serviceAccount:${google_service_account.demo_pubsub_invoker.email}"
}


resource "google_service_account" "demo_pubsub_invoker" {
  account_id  = "myinvoker"
  description = "A service account for pubsub to trigger cloud functions"
}


resource "google_pubsub_subscription" "corefirestore" {

  name  = "corefirestore-push-subscription"
  topic = google_pubsub_topic.demo.id

  ack_deadline_seconds = 20

  push_config {
    push_endpoint = google_cloudfunctions2_function.corefirestore.url
    oidc_token {
      service_account_email = google_service_account.demo_pubsub_invoker.email
    }
    no_wrapper {
      write_metadata = false
    }
  }

}

###### Skipping CurrentFlights ... by way of triggering on cloud firestore insert.
###### This seems overkill [ at least, I do not want to do the parsing, currently ]

# # https://cloud.google.com/eventarc/docs/creating-triggers-terraform
# data "archive_file" "currentflights" {
#   type        = "zip"
#   output_path = "/tmp/currentflights-function-source.zip"
#   source_dir  = "currentflights"
# }


# resource "google_storage_bucket_object" "currentflights_object" {
#   name   = "currentflights-function-source.zip"
#   bucket = google_storage_bucket.default.name
#   source = data.archive_file.currentflights.output_path
# }


# resource "google_cloudfunctions2_function" "currentflights" {
#   depends_on  = [google_storage_bucket_object.collector_object]
#   name        = "currentflights"
#   location    = "us-central1"
#   description = "currentflights function"

#   build_config {
#     runtime     = "go122"
#     entry_point = "SetMostRecentFlights"
#     # environment_variables = {
#     #     GCP_PROJECT = var.gcp_project_id
#     #     TOPIC = google_pubsub_topic.demo.name
#     # }
#     source {
#       storage_source {
#         bucket = google_storage_bucket.default.name
#         object = google_storage_bucket_object.currentflights_object.name
#       }
#     }
#   }
#   # # TODO: Look at these values
#   # service_config {
#   #   max_instance_count  = 3
#   #   min_instance_count = 1
#   #   available_memory    = "256M"
#   #   timeout_seconds     = 60
#   #   environment_variables = {
#   #       SERVICE_CONFIG_TEST = "config_test"
#   #   }
#   #   ingress_settings = "ALLOW_INTERNAL_ONLY"
#   #   all_traffic_on_latest_revision = true
#   #   service_account_email = google_service_account.account.email
#   # }

#   # TODO: WILL NEED MORE FILTERING
#   event_trigger {
#     trigger_region = "nam5"
#     event_type     = "google.cloud.firestore.document.v1.created"
#     retry_policy   = "RETRY_POLICY_RETRY"
#     event_filters {
#       attribute = "database"
#       // TODO: get this from tfconfig
#       value = "demo-ingestion"
#     }
#     # event_filters {
#     #   attribute = "document"
#     #   value     = "flights/*"
#     #   operator  = "match-path-pattern" # This allows path patterns to be used in the value field
#     # }
#   }
# }

# # resource "google_service_account" "account" {
# #   account_id   = "gcf-sa"
# #   display_name = "Test Service Account - used for both the cloud function and eventarc trigger in the test"
# # }

# # # Permissions on the service account used by the function and Eventarc trigger
# # resource "google_project_iam_member" "invoking" {
# #   project = var.gcp_project_id
# #   role    = "roles/run.invoker"
# #   member  = "serviceAccount:${google_service_account.account.email}"
# #   # TODO: WHY WAS THE FOLLOWING IN THE EXAMPLE?
# #   #depends_on = [google_project_iam_member.gcs-pubsub-publishing]
# # }

# # resource "google_project_iam_member" "event-receiving" {
# #   project = var.gcp_project_id
# #   role    = "roles/eventarc.eventReceiver"
# #   member  = "serviceAccount:${google_service_account.account.email}"
# #   depends_on = [google_project_iam_member.invoking]
# # }

# # resource "google_project_iam_member" "artifactregistry-reader" {
# #   project = var.gcp_project_id
# #   role     = "roles/artifactregistry.reader"
# #   member   = "serviceAccount:${google_service_account.account.email}"
# #   depends_on = [google_project_iam_member.event-receiving]
# # }


## TODO:  Add relevant bits for inairflights
## UNLESS just adding that code to the corefirestore function
## Pros/Cons to discuss to understand organizational culture/preferences
