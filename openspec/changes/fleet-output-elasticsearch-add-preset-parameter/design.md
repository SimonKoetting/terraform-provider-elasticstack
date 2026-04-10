## Context

The Fleet output resource already supports multiple output types and uses type-dispatched request/response mapping for CRUD operations. The Fleet Outputs API includes a `preset` field that is relevant to Elasticsearch-family outputs, but the Terraform resource schema currently does not expose it, which creates a configuration gap.

This change should be narrow and low-risk: it extends an existing resource with one optional attribute and constrains its use to compatible output types (`elasticsearch` and `remote_elasticsearch`).

## Goals / Non-Goals

**Goals:**
- Add `preset` as an optional attribute on `elasticstack_fleet_output`.
- Ensure `preset` is only valid for `elasticsearch` and `remote_elasticsearch`.
- Map `preset` in create/update requests and read responses for supported types.
- Add tests for validation and state/API mapping behavior.

**Non-Goals:**
- Introducing new Fleet output types.
- Changing default behavior for existing fields.
- Defining or validating server-specific allowed `preset` values beyond Fleet/API diagnostics unless already represented in provider validation patterns.

## Decisions

- Add `preset` as a top-level optional string attribute on the resource model and schema.
  - Rationale: `preset` is output-level configuration, not specific to nested SSL/Kafka blocks.
  - Alternative considered: embedding in `config_yaml`; rejected because it weakens explicit schema support and drift visibility.

- Enforce type-conditional validation so `preset` is accepted only when `type` is `elasticsearch` or `remote_elasticsearch`.
  - Rationale: this follows API intent and prevents invalid plans for unsupported output types.
  - Alternative considered: allowing `preset` for all types and relying on API errors; rejected to provide earlier and clearer Terraform validation feedback.

- Extend existing type-dispatch mapping paths (create/update/read) rather than introducing a new mapping abstraction.
  - Rationale: keeps the implementation consistent with current resource patterns and reduces regression risk.
  - Alternative considered: larger refactor of output mappers; rejected as out of scope for a single-field enhancement.

## Risks / Trade-offs

- [Fleet response variability for `preset`] -> Preserve behavior when field is absent in read responses and cover with unit tests against nil/empty mappings.
- [Type-conditional validator mismatch with current schema patterns] -> Reuse existing conditional validation approach already used for output-type-specific fields.
- [Potential drift for users who previously set equivalent values via `config_yaml`] -> Document precedence/expectation in resource docs and keep behavior backward compatible by not changing `config_yaml` semantics.

## Migration Plan

- No state upgrade is required because `preset` is a new optional attribute.
- Existing resources continue to work unchanged when `preset` is unset.
- Rollback is straightforward: remove provider version with `preset` support; resources using `preset` may require config cleanup before planning with older versions.

## Open Questions

- Whether Fleet enforces a strict enum of `preset` values across all supported versions, and if that enum should be provider-validated in a follow-up.
