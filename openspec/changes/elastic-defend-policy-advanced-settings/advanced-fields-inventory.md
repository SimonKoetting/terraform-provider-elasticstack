## Advanced Field Inventory (Task 1.1)

Source references considered:
- Elastic Defend API create/customize policy examples
- Elastic Defend advanced settings reference

Included in task group 1 implementation (stable, user-manageable, API-aligned typed paths):
- `policy.windows.advanced.agent.connection_delay` -> `windows.advanced.agent.connection_delay`
- `policy.mac.advanced.agent.connection_delay` -> `mac.advanced.agent.connection_delay`
- `policy.linux.advanced.agent.connection_delay` -> `linux.advanced.agent.connection_delay`
- `policy.windows.advanced.alerts.hash.md5` -> `windows.advanced.alerts.hash.md5`
- `policy.windows.advanced.alerts.hash.sha1` -> `windows.advanced.alerts.hash.sha1`
- `policy.mac.advanced.alerts.hash.md5` -> `mac.advanced.alerts.hash.md5`
- `policy.mac.advanced.alerts.hash.sha1` -> `mac.advanced.alerts.hash.sha1`
- `policy.linux.advanced.alerts.hash.md5` -> `linux.advanced.alerts.hash.md5`
- `policy.linux.advanced.alerts.hash.sha1` -> `linux.advanced.alerts.hash.sha1`
- `policy.windows.advanced.alerts.cloud_lookup` -> `windows.advanced.alerts.cloud_lookup`
- `policy.mac.advanced.alerts.cloud_lookup` -> `mac.advanced.alerts.cloud_lookup`

Explicitly excluded from task group 1:
- Artifact, transport, TLS, proxy, and other operational internals under `advanced.*` that are either server/environment managed or high-churn and not yet modeled by provider typed schema.
- Broad event tuning branches (`advanced.events.*`, callstacks, kernel/fanotify details) that significantly expand schema surface and validation scope.
- Rollback/self-healing and similar specialized branches requiring deeper lifecycle/acceptance coverage.

Rationale:
- Keep task group 1 focused on a minimal stable subset that round-trips safely in typed schema.
- Avoid introducing unbounded advanced-setting surface before broader test coverage work in later task groups.
