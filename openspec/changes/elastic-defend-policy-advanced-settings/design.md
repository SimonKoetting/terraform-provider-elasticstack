## Context

The `elasticstack_fleet_elastic_defend_integration_policy` resource already models a broad typed `policy` structure, but it does not currently guarantee full coverage of advanced Elastic Defend policy settings available in the Fleet API. Users who need these settings must apply out-of-band changes, which introduces drift and breaks declarative ownership expectations. The change must preserve the existing typed-resource model, two-phase Defend create flow, private state handling, and API-validated update semantics.

## Goals / Non-Goals

**Goals:**
- Add explicit API-aligned support for advanced Defend settings inside `policy`.
- Ensure supported advanced settings are consistently handled in create, read, update, and import.
- Keep the schema typed and Terraform-validatable, without introducing raw JSON policy input.
- Preserve stability for unknown or server-managed fields that remain outside the schema boundary.

**Non-Goals:**
- Replacing the typed Defend resource with a generic JSON passthrough model.
- Changing package identity, bootstrap/update sequencing, or private-state concurrency handling.
- Expanding scope to unrelated Fleet package policy resources.

## Decisions

- Continue with a typed schema-first approach and extend nested attributes only for advanced settings that are stable and user-manageable in the Defend API.
  - Alternative considered: expose raw advanced policy JSON as a free-form field. Rejected because it weakens validation, degrades plan readability, and conflicts with existing typed capability boundaries.
- Define requirement-level CRUD/import parity for advanced settings so behavior is testable and consistent across lifecycle operations.
  - Alternative considered: support advanced settings only on create/update. Rejected because refresh/import drift would still occur.
- Treat server-managed or unsupported advanced fields as provider-internal/ignored unless explicitly modeled.
  - Alternative considered: preserve all unknown advanced fields in public Terraform state. Rejected because opaque state expansion creates unstable diffs and breaks typed-state expectations.
- Require acceptance test expansion to cover representative advanced branches per operating system and verify round-trip behavior.

## Task 1.1 inventory artifact

- Field-by-field inclusion/exclusion rationale for the initial advanced subset is tracked in `openspec/changes/elastic-defend-policy-advanced-settings/advanced-fields-inventory.md`.

## Risks / Trade-offs

- [Risk] API advanced settings evolve faster than schema updates -> Mitigation: constrain requirements to API-supported, stable user-managed fields and expand tests as new fields are added.
- [Risk] Larger nested schema raises maintenance cost -> Mitigation: group advanced fields by existing policy branch structure and keep mapping code modular.
- [Risk] Partial branch coverage could leave hidden drift paths -> Mitigation: require CRUD/import parity and targeted acceptance assertions for advanced fields.
- [Risk] Backward compatibility concerns for existing configurations -> Mitigation: additive schema changes only; preserve existing defaults and behavior for currently modeled fields.
