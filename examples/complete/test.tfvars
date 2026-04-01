logical_product_family  = "launch"
logical_product_service = "r53auth"
class_env               = "dev"
instance_env            = 0
instance_resource       = 0

hosted_zone_domain        = "internal"
zone_vpc_cidr_block       = "10.0.0.0/16"
authorized_vpc_cidr_block = "10.1.0.0/16"

tags = {
  Environment = "test"
  Module      = "tf-aws-module_primitive-route53_vpc_association_authorization"
  ManagedBy   = "Terraform"
  Purpose     = "terratest"
}
