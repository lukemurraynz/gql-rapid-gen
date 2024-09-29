#!/bin/bash

SUBSCRIPTION_ID=${1}
RESOURCE_GROUP=${2}
STORAGE_ACCOUNT=${3}
CONTAINER=${4}
STATE_KEY=${5}
LOCATION=${6}

# Login to Azure (assumes you have already authenticated)
az account set --subscription "${SUBSCRIPTION_ID}"

# Create Resource Group
az group create --name "${RESOURCE_GROUP}" --location "${LOCATION}"

# Create Storage Account
az storage account create --name "${STORAGE_ACCOUNT}" --resource-group "${RESOURCE_GROUP}" --location "${LOCATION}" --sku Standard_LRS

# Get Storage Account Key
ACCOUNT_KEY=$(az storage account keys list --resource-group "${RESOURCE_GROUP}" --account-name "${STORAGE_ACCOUNT}" --query '[0].value' --output tsv)

# Create Blob Container
az storage container create --name "${CONTAINER}" --account-name "${STORAGE_ACCOUNT}" --account-key "${ACCOUNT_KEY}"

# Generate providers.tf
tee ./providers.tf > /dev/null <<EOF
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 4.3.0"
    }
  }
  backend "azurerm" {
    resource_group_name   = "${RESOURCE_GROUP}"
    storage_account_name  = "${STORAGE_ACCOUNT}"
    container_name        = "${CONTAINER}"
    key                   = "${STATE_KEY}"
  }
}

provider "azurerm" {
  features {}
}
EOF