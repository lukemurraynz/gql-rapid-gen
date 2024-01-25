
resource "aws_vpc" "default" {
  cidr_block                       = "10.0.0.0/16"
  assign_generated_ipv6_cidr_block = true

  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name = "${var.client-code}-${terraform.workspace}"
  }
}

resource "aws_subnet" "public_subnet" {
  count                           = length(data.aws_availability_zones.available.names)
  vpc_id                          = aws_vpc.default.id
  cidr_block                      = cidrsubnet(aws_vpc.default.cidr_block, 8, 0 + count.index)
  ipv6_cidr_block                 = cidrsubnet(aws_vpc.default.ipv6_cidr_block, 8, 0 + count.index)
  availability_zone               = data.aws_availability_zones.available.names[count.index]
  map_public_ip_on_launch         = true
  assign_ipv6_address_on_creation = true

  tags = {
    Name = "Public-${data.aws_availability_zones.available.names[count.index]}-${terraform.workspace}"
  }
}

resource "aws_route_table_association" "public_subnet" {
  count          = length(data.aws_availability_zones.available.names)
  subnet_id      = aws_subnet.public_subnet[count.index].id
  route_table_id = aws_route_table.public.id
}
/*
resource "aws_subnet" "nat_subnet" {
  count                           = length(data.aws_availability_zones.available.names)
  vpc_id                          = aws_vpc.default.id
  cidr_block                      = cidrsubnet(aws_vpc.default.cidr_block, 8, 8 + count.index)
  ipv6_cidr_block                 = cidrsubnet(aws_vpc.default.ipv6_cidr_block, 8, 8 + count.index)
  availability_zone               = data.aws_availability_zones.available.names[count.index]
  map_public_ip_on_launch         = true
  assign_ipv6_address_on_creation = true

  tags = {
    Name = "NAT-${data.aws_availability_zones.available.names[count.index]}-${terraform.workspace}"
  }
}

resource "aws_route_table_association" "nat_subnet" {
  count          = length(data.aws_availability_zones.available.names)
  subnet_id      = aws_subnet.nat_subnet[count.index].id
  route_table_id = aws_route_table.nat.id
}*/

resource "aws_subnet" "private_subnet" {
  count                           = length(data.aws_availability_zones.available.names)
  vpc_id                          = aws_vpc.default.id
  cidr_block                      = cidrsubnet(aws_vpc.default.cidr_block, 8, 16 + count.index)
  ipv6_cidr_block                 = cidrsubnet(aws_vpc.default.ipv6_cidr_block, 8, 16 + count.index)
  availability_zone               = data.aws_availability_zones.available.names[count.index]
  map_public_ip_on_launch         = true
  assign_ipv6_address_on_creation = true

  tags = {
    Name = "Private-${data.aws_availability_zones.available.names[count.index]}-${terraform.workspace}"
  }
}

resource "aws_route_table_association" "private_subnet" {
  count          = length(data.aws_availability_zones.available.names)
  subnet_id      = aws_subnet.private_subnet[count.index].id
  route_table_id = aws_route_table.private.id
}

#resource "aws_eip" "nat" {
#  tags = {
#    Name = "vpc-nat-${terraform.workspace}"
#  }
#}
/*
resource "aws_nat_gateway" "nat" {
  tags = {
    Name = "vpc-nat-${terraform.workspace}"
  }
  allocation_id = aws_eip.nat.id
  subnet_id     = aws_subnet.public_subnet[0].id

  depends_on = [aws_internet_gateway.gw]
}*/

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.default.id
  tags = {
    Name = "Public-${terraform.workspace}"
  }

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gw.id
  }
  route {
    ipv6_cidr_block        = "::/0"
    egress_only_gateway_id = aws_egress_only_internet_gateway.gw.id
  }
}
/*
resource "aws_route_table" "nat" {
  vpc_id = aws_vpc.default.id
  tags = {
    Name = "NAT-${terraform.workspace}"
  }

  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.nat.id
  }
  route {
    ipv6_cidr_block = "::/0"
    nat_gateway_id  = aws_nat_gateway.nat.id
  }
}*/

resource "aws_route_table" "private" {
  vpc_id = aws_vpc.default.id
  tags = {
    Name = "Private-${terraform.workspace}"
  }

}

resource "aws_internet_gateway" "gw" {
  vpc_id = aws_vpc.default.id

  tags = {
    Name = "internet-${terraform.workspace}"
  }
}

resource "aws_egress_only_internet_gateway" "gw" {
  vpc_id = aws_vpc.default.id

  tags = {
    Name = "internet-${terraform.workspace}"
  }
}