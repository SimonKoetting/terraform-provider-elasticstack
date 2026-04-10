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

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func presetDefaultBalancedForElasticsearchFamily() planmodifier.String {
	return presetDefaultBalancedModifier{}
}

type presetDefaultBalancedModifier struct{}

func (presetDefaultBalancedModifier) Description(_ context.Context) string {
	return "Defaults preset to balanced when type is elasticsearch or remote_elasticsearch and preset is omitted in configuration."
}

func (m presetDefaultBalancedModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (presetDefaultBalancedModifier) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if !req.ConfigValue.IsNull() {
		return
	}

	var outputType types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("type"), &outputType)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if outputType.IsNull() || outputType.IsUnknown() {
		return
	}

	switch outputType.ValueString() {
	case "elasticsearch", "remote_elasticsearch":
		resp.PlanValue = types.StringValue(defaultFleetOutputPreset)
	}
}
