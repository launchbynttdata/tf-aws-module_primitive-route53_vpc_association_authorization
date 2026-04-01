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

output "id" {
  description = "The calculated unique identifier for the association authorization."
  value       = module.vpc_association_authorization.id
}

output "zone_id" {
  description = "The ID of the private hosted zone."
  value       = module.vpc_association_authorization.zone_id
}

output "vpc_id" {
  description = "The ID of the authorized VPC."
  value       = module.vpc_association_authorization.vpc_id
}

output "vpc_region" {
  description = "The region of the authorized VPC."
  value       = module.vpc_association_authorization.vpc_region
}

output "zone_vpc_id" {
  description = "The ID of the VPC that owns the hosted zone."
  value       = aws_vpc.zone_vpc.id
}

output "hosted_zone_name" {
  description = "The name of the private hosted zone."
  value       = aws_route53_zone.zone.name
}
