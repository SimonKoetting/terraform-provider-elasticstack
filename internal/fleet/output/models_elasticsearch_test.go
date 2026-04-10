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
	"testing"

	"github.com/elastic/terraform-provider-elasticstack/generated/kbapi"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOutputModelToAPICreateElasticsearchModelIncludesPreset(t *testing.T) {
	t.Parallel()

	model := outputModel{
		OutputID: types.StringValue("elasticsearch-output-id"),
		Name:     types.StringValue("elasticsearch-output"),
		Type:     types.StringValue("elasticsearch"),
		Hosts:    types.ListValueMust(types.StringType, []attr.Value{types.StringValue("https://elasticsearch:9200")}),
		Preset:   types.StringValue("throughput"),
	}

	union, diags := model.toAPICreateElasticsearchModel(context.Background())
	require.False(t, diags.HasError())

	body, err := union.AsNewOutputElasticsearch()
	require.NoError(t, err)
	require.NotNil(t, body.Preset)
	assert.Equal(t, kbapi.KibanaHTTPAPIsNewOutputElasticsearchPreset("throughput"), *body.Preset)
}

func TestOutputModelToAPIUpdateElasticsearchModelIncludesPreset(t *testing.T) {
	t.Parallel()

	model := outputModel{
		Name:   types.StringValue("updated-elasticsearch-output"),
		Type:   types.StringValue("elasticsearch"),
		Hosts:  types.ListValueMust(types.StringType, []attr.Value{types.StringValue("https://elasticsearch:9200")}),
		Preset: types.StringValue("balanced"),
	}

	union, diags := model.toAPIUpdateElasticsearchModel(context.Background())
	require.False(t, diags.HasError())

	body, err := union.AsUpdateOutputElasticsearch()
	require.NoError(t, err)
	require.NotNil(t, body.Preset)
	assert.Equal(t, kbapi.UpdateOutputElasticsearchPreset("balanced"), *body.Preset)
}

func TestOutputModelFromAPIElasticsearchModelMapsPreset(t *testing.T) {
	t.Parallel()

	model := outputModel{
		SpaceIDs: types.SetNull(types.StringType),
	}

	diags := model.fromAPIElasticsearchModel(context.Background(), &kbapi.OutputElasticsearch{
		Id:    new("output-id"),
		Name:  "elasticsearch-output",
		Type:  kbapi.KibanaHTTPAPIsOutputElasticsearchTypeElasticsearch,
		Hosts: []string{"https://elasticsearch:9200"},
		Preset: func() *kbapi.KibanaHTTPAPIsOutputElasticsearchPreset {
			value := kbapi.KibanaHTTPAPIsOutputElasticsearchPreset("latency")
			return &value
		}(),
	})
	require.False(t, diags.HasError())

	assert.Equal(t, "latency", model.Preset.ValueString())
}
