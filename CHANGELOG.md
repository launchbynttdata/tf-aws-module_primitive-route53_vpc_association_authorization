# Changelog

## [Unreleased]

### Breaking Changes

- Widened AWS provider constraint from `~> 5.100` to `>= 5.0` (supports both v5 and v6+)
- Updated Terraform version constraint from `~> 1.9` to `~> 1.10`

### Added

- `examples/min_provider` example with `~> 5.0` constraint for minimum provider testing
- Terratest functions for minimum provider verification

### Fixed

- Replaced deprecated `data.aws_region.current.name` with `.id` in example

### Validation

- Reviewed AWS provider v6 upgrade guide: no breaking changes to `aws_route53_vpc_association_authorization`
- Reviewed resource documentation for provider v6
- Module validates with both AWS provider v5 and v6
- No argument or attribute changes required; module is already feature complete
