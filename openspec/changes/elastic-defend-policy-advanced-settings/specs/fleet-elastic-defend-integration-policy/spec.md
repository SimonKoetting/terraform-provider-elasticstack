## MODIFIED Requirements

### Requirement: Typed Defend configuration schema (REQ-006)

The resource SHALL model Defend-owned configuration through typed Terraform attributes and nested attributes. The `preset` attribute SHALL map to `config.integration_config.value.endpointConfig.preset` in read/update payloads. The `policy` attribute SHALL contain optional `windows`, `mac`, and `linux` nested attributes, each with a distinct schema containing only the fields applicable to that operating system. Structurally invalid combinations (such as `policy.linux.ransomware`) SHALL be impossible at plan time without requiring custom validation. The typed schema SHALL also model the provider-known Defend defaults for popup entries, protection mode objects, behavior protection, Windows antivirus registration, and Windows attack surface reduction credential hardening so omitted or empty nested blocks plan with the same effective values the Defend policy uses. The typed schema SHALL include supported advanced policy settings exposed by the Elastic Defend API so users can configure the full supported Defend functionality from Terraform without out-of-band edits.

#### Scenario: Policy settings are modeled as typed attributes
- **WHEN** a configuration enables or disables Defend protections, event collection, or supported advanced settings
- **THEN** those settings SHALL be represented by typed resource attributes and nested attributes
- **AND** the configuration SHALL NOT require users to provide raw `policy` JSON

#### Scenario: Linux event settings include documented Linux-specific leaves
- **WHEN** Terraform maps the `policy.linux.events` schema to and from the API
- **THEN** the typed schema SHALL include the documented Linux-specific event flags
- **AND** those flags SHALL include `session_data` and `tty_io`

#### Scenario: Omitted nested policy settings use modeled defaults
- **WHEN** a configuration omits or leaves empty a popup item, protection-mode object, behavior-protection object, Windows antivirus registration object, or Windows attack-surface-reduction credential hardening object
- **THEN** the omitted or empty nested object SHALL use the modeled defaults
- **AND** `policy.windows.popup` SHALL default to an object whose `malware`, `ransomware`, `memory_protection`, and `behavior_protection` entries each use the popup-item defaults
- **AND** popup items SHALL default to `{ message = "", enabled = false }`
- **AND** protection-mode objects SHALL default to `{ mode = "off", supported = true }`
- **AND** behavior-protection objects SHALL default to `{ mode = "off", supported = true, reputation_service = false }`
- **AND** `policy.windows.antivirus_registration` SHALL default to `{ mode = "disabled", enabled = false }`
- **AND** `policy.windows.attack_surface_reduction.credential_hardening` SHALL default to `{ enabled = false }`

#### Scenario: Advanced policy settings maintain API-aligned shape
- **WHEN** Terraform plans and applies advanced Defend settings under supported `policy` branches
- **THEN** the provider SHALL map those settings to the API using the corresponding Defend typed input shape
- **AND** read and refresh SHALL map the returned advanced values back into the same typed schema paths

### Requirement: Read and import map only modeled fields to state (REQ-011)

On read and import, the resource SHALL parse the Defend-specific package policy response and populate only the modeled Terraform schema fields. The provider SHALL map `preset` from the Defend `integration_config` payload and SHALL map the typed `policy` payload into the corresponding operating-system nested attributes, including supported advanced settings modeled by the schema. The provider SHALL ignore unmodeled server-managed Defend payloads in Terraform state, except for preserving any opaque values required for future updates in internal provider-managed state.

#### Scenario: Read ignores unmodeled server-managed Defend fields
- **WHEN** a Defend package policy response includes `artifact_manifest` and other server-managed Defend data
- **THEN** Terraform state SHALL include only the modeled schema fields
- **AND** the provider SHALL preserve any required opaque update data internally

#### Scenario: Import round-trips modeled advanced policy settings
- **WHEN** an existing Defend package policy with supported advanced settings is imported and read
- **THEN** Terraform state SHALL include those modeled advanced settings in their typed schema paths
- **AND** a subsequent plan with matching configuration SHALL produce no advanced-setting drift
