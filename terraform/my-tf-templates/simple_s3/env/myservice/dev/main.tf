provider "aws" {
  region  = "us-east-2"
  profile = "default"

  default_tags {
    tags = {
      AppEnv  = local.env
      Project = "${local.env}-${local.project_name}"
    }
  }
}

terraform {
  required_version = ">= 1.0.0, < 2.0.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}
