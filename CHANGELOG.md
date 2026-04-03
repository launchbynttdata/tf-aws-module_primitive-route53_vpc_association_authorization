# Changelog

## [Unreleased]

### Breaking Changes

- Upgraded AWS provider from `~> 5.100` to `~> 6.0`
- Updated Terraform version constraint from `~> 1.9` to `~> 1.10`

### Fixed

- Replaced deprecated `data.aws_region.current.name` with `.id` in example

### Validation

- Reviewed AWS provider v6 upgrade guide: no breaking changes to `aws_route53_vpc_association_authorization`
- Reviewed resource documentation for provider v6
- No argument or attribute changes required; module is already feature complete
