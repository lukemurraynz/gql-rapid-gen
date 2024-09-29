resource "azurerm_dns_zone" "default" {
  name                = var.cert_domain
  resource_group_name = azurerm_resource_group.main.name
  tags = {
    Environment = terraform.workspace
  }
}