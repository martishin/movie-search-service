# Configure AWS provider
provider "aws" {
  alias  = "acm"
  region = var.aws_region
}

# VPC and networking components
resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"

  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = {
    Name = "main-vpc"
  }
}
