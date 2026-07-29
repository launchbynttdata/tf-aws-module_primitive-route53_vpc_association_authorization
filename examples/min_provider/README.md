# min_provider

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.10 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.0 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_resource_names"></a> [resource\_names](#module\_resource\_names) | terraform.registry.launch.nttdata.com/module_library/resource_name/launch | ~> 2.0 |
| <a name="module_vpc_association_authorization"></a> [vpc\_association\_authorization](#module\_vpc\_association\_authorization) | ../.. | n/a |

## Resources

| Name | Type |
|------|------|
| [aws_default_security_group.authorized_vpc](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/default_security_group) | resource |
| [aws_default_security_group.zone_vpc](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/default_security_group) | resource |
| [aws_route53_zone.zone](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_zone) | resource |
| [aws_vpc.authorized_vpc](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc) | resource |
| [aws_vpc.zone_vpc](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/vpc) | resource |
| [aws_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_authorized_vpc_cidr_block"></a> [authorized\_vpc\_cidr\_block](#input\_authorized\_vpc\_cidr\_block) | CIDR block for the VPC to be authorized for association. | `string` | `"10.1.0.0/16"` | no |
| <a name="input_class_env"></a> [class\_env](#input\_class\_env) | Environment class for the resource naming module (e.g. dev, staging, prod). | `string` | `"dev"` | no |
| <a name="input_hosted_zone_domain"></a> [hosted\_zone\_domain](#input\_hosted\_zone\_domain) | The base domain suffix for the private hosted zone name. | `string` | `"internal"` | no |
| <a name="input_instance_env"></a> [instance\_env](#input\_instance\_env) | Environment instance number for the resource naming module. | `number` | `0` | no |
| <a name="input_instance_resource"></a> [instance\_resource](#input\_instance\_resource) | Resource instance number for the resource naming module. | `number` | `0` | no |
| <a name="input_logical_product_family"></a> [logical\_product\_family](#input\_logical\_product\_family) | Logical product family name used by the resource naming module. | `string` | `"launch"` | no |
| <a name="input_logical_product_service"></a> [logical\_product\_service](#input\_logical\_product\_service) | Logical product service name used by the resource naming module. | `string` | `"r53auth"` | no |
| <a name="input_resource_names_map"></a> [resource\_names\_map](#input\_resource\_names\_map) | Map of resource names with cloud\_resource\_type and max\_length for the resource naming module. | <pre>map(object({<br/>    name       = string<br/>    max_length = optional(number, 60)<br/>  }))</pre> | <pre>{<br/>  "authorized_vpc": {<br/>    "max_length": 60,<br/>    "name": "authvpc1"<br/>  },<br/>  "hosted_zone": {<br/>    "max_length": 60,<br/>    "name": "hz1"<br/>  },<br/>  "zone_vpc": {<br/>    "max_length": 60,<br/>    "name": "zonevpc1"<br/>  }<br/>}</pre> | no |
| <a name="input_tags"></a> [tags](#input\_tags) | Map of tags to assign to the resources. | `map(string)` | `{}` | no |
| <a name="input_vpc_region"></a> [vpc\_region](#input\_vpc\_region) | The VPC's region. Defaults to the region of the AWS provider if not specified. | `string` | `null` | no |
| <a name="input_zone_vpc_cidr_block"></a> [zone\_vpc\_cidr\_block](#input\_zone\_vpc\_cidr\_block) | CIDR block for the VPC that owns the private hosted zone. | `string` | `"10.0.0.0/16"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_hosted_zone_name"></a> [hosted\_zone\_name](#output\_hosted\_zone\_name) | The name of the private hosted zone. |
| <a name="output_id"></a> [id](#output\_id) | The calculated unique identifier for the association authorization. |
| <a name="output_vpc_id"></a> [vpc\_id](#output\_vpc\_id) | The ID of the authorized VPC. |
| <a name="output_vpc_region"></a> [vpc\_region](#output\_vpc\_region) | The region of the authorized VPC. |
| <a name="output_zone_id"></a> [zone\_id](#output\_zone\_id) | The ID of the private hosted zone. |
| <a name="output_zone_vpc_id"></a> [zone\_vpc\_id](#output\_zone\_vpc\_id) | The ID of the VPC that owns the hosted zone. |
<!-- END_TF_DOCS -->
