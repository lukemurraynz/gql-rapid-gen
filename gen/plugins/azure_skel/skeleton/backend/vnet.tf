provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "main" {
  name     = "${var.client-code}-${terraform.workspace}-rg"
  location = var.location
}

resource "azurerm_virtual_network" "main" {
  name                = "${var.client-code}-${terraform.workspace}-vnet"
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name

  tags = {
    Name = "${var.client-code}-${terraform.workspace}"
  }
}

resource "azurerm_subnet" "public_subnet" {
  count                = length(data.azurerm_availability_zones.available.names)
  name                 = "Public-${count.index}-${terraform.workspace}"
  resource_group_name  = azurerm_resource_group.main.name
  virtual_network_name = azurerm_virtual_network.main.name
  address_prefixes     = [cidrsubnet(azurerm_virtual_network.main.address_space[0], 8, 0 + count.index)]

  tags = {
    Name = "Public-${data.azurerm_availability_zones.available.names[count.index]}-${terraform.workspace}"
  }
}

resource "azurerm_subnet" "private_subnet" {
  count                = length(data.azurerm_availability_zones.available.names)
  name                 = "Private-${count.index}-${terraform.workspace}"
  resource_group_name  = azurerm_resource_group.main.name
  virtual_network_name = azurerm_virtual_network.main.name
  address_prefixes     = [cidrsubnet(azurerm_virtual_network.main.address_space[0], 8, 16 + count.index)]

  tags = {
    Name = "Private-${data.azurerm_availability_zones.available.names[count.index]}-${terraform.workspace}"
  }
}

resource "azurerm_route_table" "public" {
  name                = "Public-${terraform.workspace}"
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name

  route {
    name                   = "internet-route"
    address_prefix         = "0.0.0.0/0"
    next_hop_type          = "Internet"
  }

  tags = {
    Name = "Public-${terraform.workspace}"
  }
}

resource "azurerm_route_table" "private" {
  name                = "Private-${terraform.workspace}"
  location            = azurerm_resource_group.main.location
  resource_group_name = azurerm_resource_group.main.name

  tags = {
    Name = "Private-${terraform.workspace}"
  }
}

resource "azurerm_subnet_route_table_association" "public" {
  count          = length(data.azurerm_availability_zones.available.names)
  subnet_id      = azurerm_subnet.public_subnet[count.index].id
  route_table_id = azurerm_route_table.public.id
}

resource "azurerm_subnet_route_table_association" "private" {
  count          = length(data.azurerm_availability_zones.available.names)
  subnet_id      = azurerm_subnet.private_subnet[count.index].id
  route_table_id = azurerm_route_table.private.id
}