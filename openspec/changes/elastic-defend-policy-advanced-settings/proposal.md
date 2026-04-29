## Why

The Elastic Defend resource currently does not expose the full set of advanced policy settings available in the Fleet API, which prevents users from managing all supported Defend behavior declaratively through Terraform. Expanding coverage now closes parity gaps and avoids forced out-of-band configuration drift.

## What Changes

- Expand `elasticstack_fleet_elastic_defend_integration_policy` so `policy` supports advanced settings available in the Elastic Defend API payload.
- Add explicit requirements for end-to-end mapping of supported advanced policy settings across create, read, update, and import flows.
- Define behavior for unknown or server-managed advanced fields so provider behavior remains stable while preserving full user-controlled functionality.
- Add acceptance coverage expectations for advanced settings across supported operating systems and policy branches.

## Capabilities

### New Capabilities
<!-- Capabilities being introduced. Replace <name> with kebab-case identifier (e.g., user-auth, data-export, api-rate-limiting). Each creates specs/<name>/spec.md -->

### Modified Capabilities
<!-- Existing capabilities whose REQUIREMENTS are changing (not just implementation).
     Only list here if spec-level behavior changes. Each needs a delta spec file.
     Use existing spec names from openspec/specs/. Leave empty if no requirement changes. -->
- `fleet-elastic-defend-integration-policy`: extend typed `policy` requirements so advanced settings can be configured with API-aligned behavior and full CRUD/import parity.

## Impact

- Affected code: `internal/fleet/elastic_defend_integration_policy` schema, mapping, and request/response translation logic.
- Affected tests: acceptance tests for Elastic Defend policy resource and related fixture/config helpers.
- External behavior: expanded Terraform surface within `policy` to align with API capabilities while maintaining typed, validated configuration.
