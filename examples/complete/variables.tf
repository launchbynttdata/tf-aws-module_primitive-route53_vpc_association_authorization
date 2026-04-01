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

// Resource naming variables
variable "resource_names_map" {
  description = "Map of resource names with cloud_resource_type and max_length for the resource naming module."
  type = map(object({
    name       = string
    max_length = optional(number, 60)
  }))
  default = {
    zone_vpc = {
      name       = "zonevpc1"
      max_length = 60
    }
    authorized_vpc = {
      name       = "authvpc1"
      max_length = 60
    }
    hosted_zone = {
      name       = "hz1"
      max_length = 60
    }
  }
}

variable "logical_product_family" {
  description = "Logical product family name used by the resource naming module."
  type        = string
  default     = "launch"
}

variable "logical_product_service" {
  description = "Logical product service name used by the resource naming module."
  type        = string
  default     = "r53auth"
}

variable "class_env" {
  description = "Environment class for the resource naming module (e.g. dev, staging, prod)."
  type        = string
  default     = "dev"
}

variable "instance_env" {
  description = "Environment instance number for the resource naming module."
  type        = number
  default     = 0
}

variable "instance_resource" {
  description = "Resource instance number for the resource naming module."
  type        = number
  default     = 0
}

// Example-specific variables
variable "hosted_zone_domain" {
  description = "The base domain suffix for the private hosted zone name."
  type        = string
  default     = "internal"
}

variable "zone_vpc_cidr_block" {
  description = "CIDR block for the VPC that owns the private hosted zone."
  type        = string
  default     = "10.0.0.0/16"
}

variable "authorized_vpc_cidr_block" {
  description = "CIDR block for the VPC to be authorized for association."
  type        = string
  default     = "10.1.0.0/16"
}

// Root module pass-through variables
variable "vpc_region" {
  description = "The VPC's region. Defaults to the region of the AWS provider if not specified."
  type        = string
  default     = null
}

variable "tags" {
  description = "Map of tags to assign to the resources."
  type        = map(string)
  default     = {}
}
