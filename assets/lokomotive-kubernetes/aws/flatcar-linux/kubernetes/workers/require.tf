# Terraform version and plugin versions

terraform {
  required_version = ">= 0.12.0"

  required_providers {
    aws      = "2.48.0"
    ct       = "0.5.0"
    template = "~> 2.1"
  }
}
