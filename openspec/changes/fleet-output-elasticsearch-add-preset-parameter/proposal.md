## Why

The Fleet Outputs API supports a `preset` option for Elasticsearch-family output types, but `elasticstack_fleet_output` does not expose that setting today. This prevents users from expressing the full server-supported configuration for `elasticsearch` and `remote_elasticsearch` outputs in Terraform.

## What Changes

- Add a `preset` attribute to `elasticstack_fleet_output`.
- Accept `preset` when `type` is `elasticsearch` or `remote_elasticsearch`.
- Reject `preset` for non-Elasticsearch output types (`logstash`, `kafka`) with schema validation.
- Include `preset` in create/update API payload mapping and read-state mapping for supported types.
- Preserve existing behavior for resources that do not set `preset`.

## Capabilities

### New Capabilities
- None.

### Modified Capabilities
- `fleet-output`: extend Fleet output resource requirements and schema mapping to support the `preset` option for `elasticsearch` and `remote_elasticsearch` types.

## Impact

- Affected specs: `openspec/specs/fleet-output/spec.md` (delta in change folder).
- Affected provider code: `internal/fleet/output` resource schema, validators, request/response mapping, and tests.
- API surface alignment: match Fleet Outputs API behavior documented for output details.
