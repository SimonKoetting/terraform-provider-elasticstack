## ADDED Requirements

### Requirement: Fleet output preset attribute for Elasticsearch-family types
The `elasticstack_fleet_output` resource SHALL expose an optional `preset` attribute. The provider SHALL accept `preset` only when `type` is `elasticsearch` or `remote_elasticsearch`.

#### Scenario: Configure preset for elasticsearch output
- **WHEN** a resource configuration sets `type = "elasticsearch"` and sets `preset`
- **THEN** schema validation SHALL accept the configuration

#### Scenario: Configure preset for remote_elasticsearch output
- **WHEN** a resource configuration sets `type = "remote_elasticsearch"` and sets `preset`
- **THEN** schema validation SHALL accept the configuration

### Requirement: Preset is invalid for unsupported output types
The provider SHALL reject configurations that set `preset` when `type` is not `elasticsearch` or `remote_elasticsearch`.

#### Scenario: Configure preset for logstash output
- **WHEN** a resource configuration sets `type = "logstash"` and sets `preset`
- **THEN** schema validation SHALL return an error

#### Scenario: Configure preset for kafka output
- **WHEN** a resource configuration sets `type = "kafka"` and sets `preset`
- **THEN** schema validation SHALL return an error

### Requirement: Preset request and state mapping
For supported types, the provider SHALL include `preset` in Fleet create and update requests when configured, and SHALL map `preset` from Fleet get responses into Terraform state during read.

#### Scenario: Create or update with preset
- **WHEN** a supported output type is created or updated with `preset` configured
- **THEN** the provider SHALL send `preset` in the Fleet API request body

#### Scenario: Read maps preset into state
- **WHEN** Fleet returns `preset` for an `elasticsearch` or `remote_elasticsearch` output
- **THEN** the provider SHALL persist that `preset` value in Terraform state
