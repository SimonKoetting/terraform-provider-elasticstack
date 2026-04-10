## 1. Resource schema and validation

- [ ] 1.1 Add an optional `preset` attribute to `elasticstack_fleet_output` schema and model.
- [ ] 1.2 Add conditional validation so `preset` is accepted only for `type = "elasticsearch"` and `type = "remote_elasticsearch"`.
- [ ] 1.3 Ensure schema diagnostics clearly fail when `preset` is configured for unsupported types.

## 2. API and state mapping

- [ ] 2.1 Update create/update request builders to include `preset` for supported output types when configured.
- [ ] 2.2 Update read/type-dispatch mapping to persist API-returned `preset` into Terraform state for supported output types.
- [ ] 2.3 Ensure behavior remains backward compatible when `preset` is omitted.

## 3. Test coverage and docs

- [ ] 3.1 Add or extend unit tests for type-conditional validation and request/response mapping of `preset`.
- [ ] 3.2 Add or extend acceptance tests for Fleet output lifecycle scenarios that include `preset` on supported types.
- [ ] 3.3 Update `docs/resources/fleet_output.md` to document `preset`, supported types, and invalid-type behavior.

## 4. Verification

- [ ] 4.1 Run targeted Go tests for touched Fleet output packages and fix regressions.
- [ ] 4.2 Run `make build` to verify provider build integrity.
- [ ] 4.3 Confirm OpenSpec artifact readiness with `openspec status --change fleet-output-elasticsearch-add-preset-parameter`.
