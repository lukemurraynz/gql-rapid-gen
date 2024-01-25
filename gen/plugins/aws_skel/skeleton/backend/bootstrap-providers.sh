#!/bin/bash

REGION=${1}
BUCKET=${2}
STATE_KEY=${3}

export AWS_DEFAULT_REGION=${REGION}

aws s3api create-bucket –bucket "${BUCKET}" –region "${REGION}" –create-bucket-configuration "LocationConstraint=${REGION}"
aws s3api put-bucket-encryption –bucket "${BUCKET}" –server-side-encryption-configuration "{\"Rules\": [{\"ApplyServerSideEncryptionByDefault\":{\"SSEAlgorithm\": \"AES256\"}}]}"

aws dynamodb create-table –table-name "terraform-state" –attribute-definitions "AttributeName=LockID,AttributeType=S" –key-schema "AttributeName=LockID,KeyType=HASH" –provisioned-throughput "ReadCapacityUnits=5,WriteCapacityUnits=5"

tee ./providers.tf > /dev/null <<EOF
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.14"
    }
  }
  backend "s3" {
    bucket         = "${BUCKET}"
    key            = "${STATE_KEY}"
    region         = "${REGION}"
    dynamodb_table = "terraform-state"
  }
}

provider "aws" {
  region = "${REGION}"
}

provider "aws" {
  alias  = "us-east-1"
  region = "us-east-1"
}

provider "aws" {
  alias  = "us-west-2"
  region = "us-west-2"
}

data "aws_availability_zones" "available" {

}
EOF
