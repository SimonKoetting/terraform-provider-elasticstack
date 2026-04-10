## 1. Resource schema and validation

- [x] 1.1 Add an optional `preset` attribute to `elasticstack_fleet_output` schema and model.
- [x] 1.2 Add conditional validation so `preset` is accepted only for `type = "elasticsearch"` and `type = "remote_elasticsearch"`.
- [x] 1.3 Ensure schema diagnostics clearly fail when `preset` is configured for unsupported types.

## 2. API and state mapping

- [x] 2.1 Update create/update request builders to include `preset` for supported output types when configured.
- [x] 2.2 Update read/type-dispatch mapping to persist API-returned `preset` into Terraform state for supported output types.
- [x] 2.3 Ensure behavior remains backward compatible when `preset` is omitted.

## 3. Test coverage and docs

- [x] 3.1 Add or extend unit tests for type-conditional validation and request/response mapping of `preset`.
- [x] 3.2 Add or extend acceptance tests for Fleet output lifecycle scenarios that include `preset` on supported types.
- [x] 3.3 Update `docs/resources/fleet_output.md` to document `preset`, supported types, and invalid-type behavior.

## 4. Verification

- [x] 4.1 Run targeted Go tests for touched Fleet output packages and fix regressions.
- [x] 4.2 Run `make build` to verify provider build integrity (requires `make setup` or at least `make tools` and `make golangci-lint-custom` so the `acctestconfigdirlint` plugin is available to `golangci-lint-custom`).
- [x] 4.3 Confirm OpenSpec artifact readiness with `openspec status --change fleet-output-elasticsearch-add-preset-parameter`.
