// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

data "aws_region" "current" {}

module "resource_names" {
  source  = "terraform.registry.launch.nttdata.com/module_library/resource_name/launch"
  version = "~> 2.0"

  for_each = var.resource_names_map

  logical_product_family  = var.logical_product_family
  logical_product_service = var.logical_product_service
  class_env               = var.class_env
  instance_env            = var.instance_env
  instance_resource       = var.instance_resource
  cloud_resource_type     = each.value.name
  maximum_length          = each.value.max_length
  region                  = join("", split("-", data.aws_region.current.name))
}

// VPC that owns the private hosted zone
resource "aws_vpc" "zone_vpc" {
  cidr_block           = var.zone_vpc_cidr_block
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = merge(var.tags, {
    Name = module.resource_names["zone_vpc"].standard
  })
}

// Private hosted zone associated with the zone VPC
resource "aws_route53_zone" "zone" {
  name = "${module.resource_names["hosted_zone"].standard}.${var.hosted_zone_domain}"

  vpc {
    vpc_id = aws_vpc.zone_vpc.id
  }

  lifecycle {
    ignore_changes = [vpc]
  }

  tags = var.tags
}

// Restrict all traffic on the zone VPC default security group (FG_R00089)
resource "aws_default_security_group" "zone_vpc" {
  vpc_id = aws_vpc.zone_vpc.id

  tags = var.tags
}

// VPC to be authorized for association
resource "aws_vpc" "authorized_vpc" {
  cidr_block           = var.authorized_vpc_cidr_block
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = merge(var.tags, {
    Name = module.resource_names["authorized_vpc"].standard
  })
}

// Restrict all traffic on the authorized VPC default security group (FG_R00089)
resource "aws_default_security_group" "authorized_vpc" {
  vpc_id = aws_vpc.authorized_vpc.id

  tags = var.tags
}

// The primitive module under test
module "vpc_association_authorization" {
  source = "../.."

  zone_id    = aws_route53_zone.zone.zone_id
  vpc_id     = aws_vpc.authorized_vpc.id
  vpc_region = var.vpc_region
}
