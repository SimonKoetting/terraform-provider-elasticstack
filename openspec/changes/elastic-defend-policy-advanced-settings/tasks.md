## 1. Schema and mapping extension

- [ ] 1.1 Inventory advanced Elastic Defend policy fields in the Fleet API that are stable and user-manageable, then map them to typed Terraform schema paths under `policy`.
- [ ] 1.2 Extend `elasticstack_fleet_elastic_defend_integration_policy` nested schema to include the selected advanced settings while preserving existing defaults and validation boundaries.
- [ ] 1.3 Update API translation logic so advanced settings round-trip consistently in create, read, update, and import flows.

## 2. Lifecycle and drift safety

- [ ] 2.1 Ensure create/update payload building includes modeled advanced settings in the Defend typed input shape without introducing unsupported server-managed fields.
- [ ] 2.2 Confirm read/import behavior maps modeled advanced settings into state and ignores unmodeled server-managed fields to avoid unstable diffs.
- [ ] 2.3 Add or refine diagnostics for unsupported or invalid advanced-setting combinations surfaced by API responses.

## 3. Validation and acceptance coverage

- [ ] 3.1 Add acceptance test scenarios that configure representative advanced settings across Windows, macOS, and Linux policy branches.
- [ ] 3.2 Add import/refresh parity checks to verify advanced settings do not drift after import and subsequent plan.
- [ ] 3.3 Run targeted tests for the Elastic Defend resource and fix any regressions in schema defaults or API mapping behavior.
