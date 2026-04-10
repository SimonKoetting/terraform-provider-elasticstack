// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package output

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSchemaIncludesRemoteElasticsearchTypeAndServiceToken(t *testing.T) {
	t.Parallel()

	s := getSchema()

	typeAttr, ok := s.Attributes["type"].(schema.StringAttribute)
	require.True(t, ok)
	require.NotEmpty(t, typeAttr.Validators)

	allowedType := false
	for _, validator := range typeAttr.Validators {
		if strings.Contains(validator.Description(context.Background()), "remote_elasticsearch") {
			allowedType = true
			break
		}
	}
	assert.True(t, allowedType, "expected remote_elasticsearch to be an allowed type")

	serviceTokenAttr, ok := s.Attributes["service_token"].(schema.StringAttribute)
	require.True(t, ok)
	assert.True(t, serviceTokenAttr.Sensitive)
	assert.True(t, serviceTokenAttr.Optional)
	assert.NotEmpty(t, serviceTokenAttr.Validators)

	presetAttr, ok := s.Attributes["preset"].(schema.StringAttribute)
	require.True(t, ok)
	assert.True(t, presetAttr.Optional)
	assert.NotEmpty(t, presetAttr.Validators)

	allowedPresetType := false
	for _, validator := range presetAttr.Validators {
		description := validator.Description(context.Background())
		if strings.Contains(description, "elasticsearch") && strings.Contains(description, "remote_elasticsearch") {
			allowedPresetType = true
			break
		}
	}
	assert.True(t, allowedPresetType, "expected preset to be restricted to Elasticsearch-family output types")

	hasNonEmptyWhenSet := false
	for _, validator := range presetAttr.Validators {
		desc := validator.Description(context.Background())
		if strings.Contains(strings.ToLower(desc), "at least") || strings.Contains(strings.ToLower(desc), "length") {
			hasNonEmptyWhenSet = true
			break
		}
	}
	assert.True(t, hasNonEmptyWhenSet, "expected preset to enforce non-empty when set")
}
