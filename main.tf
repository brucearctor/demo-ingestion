# TBD
terraform {
  backend "gcs" {
    bucket  = "brucearctor-demo-ingestion-tf-state"
    prefix  = "tofu/state"
  }
}
