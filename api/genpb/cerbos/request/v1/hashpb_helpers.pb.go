// Code generated by protoc-gen-go-hashpb. Do not edit.
// protoc-gen-go-hashpb v0.1.0

package requestv1

import (
	v1 "github.com/cerbos/cerbos/api/genpb/cerbos/engine/v1"
	v11 "github.com/cerbos/cerbos/api/genpb/cerbos/policy/v1"
	v12 "github.com/cerbos/cerbos/api/genpb/cerbos/schema/v1"
	protowire "google.golang.org/protobuf/encoding/protowire"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	hash "hash"
	math "math"
	sort "sort"
)

func cerbos_engine_v1_Principal_hashpb_sum(m *v1.Principal, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.Principal.id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Id))

	}
	if _, ok := ignore["cerbos.engine.v1.Principal.policy_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.PolicyVersion))

	}
	if _, ok := ignore["cerbos.engine.v1.Principal.roles"]; !ok {
		if len(m.Roles) > 0 {
			for _, v := range m.Roles {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.Principal.attr"]; !ok {
		if len(m.Attr) > 0 {
			keys := make([]string, len(m.Attr))
			i := 0
			for k := range m.Attr {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Attr[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Attr[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.Principal.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Scope))

	}
}

func cerbos_engine_v1_Resource_hashpb_sum(m *v1.Resource, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.Resource.kind"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Kind))

	}
	if _, ok := ignore["cerbos.engine.v1.Resource.policy_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.PolicyVersion))

	}
	if _, ok := ignore["cerbos.engine.v1.Resource.id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Id))

	}
	if _, ok := ignore["cerbos.engine.v1.Resource.attr"]; !ok {
		if len(m.Attr) > 0 {
			keys := make([]string, len(m.Attr))
			i := 0
			for k := range m.Attr {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Attr[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Attr[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.Resource.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Scope))

	}
}

func cerbos_engine_v1_ResourcesQueryPlanRequest_Resource_hashpb_sum(m *v1.ResourcesQueryPlanRequest_Resource, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanRequest.Resource.kind"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Kind))

	}
	if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanRequest.Resource.attr"]; !ok {
		if len(m.Attr) > 0 {
			keys := make([]string, len(m.Attr))
			i := 0
			for k := range m.Attr {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Attr[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Attr[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanRequest.Resource.policy_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.PolicyVersion))

	}
	if _, ok := ignore["cerbos.engine.v1.ResourcesQueryPlanRequest.Resource.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Scope))

	}
}

func cerbos_policy_v1_Condition_hashpb_sum(m *v11.Condition, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Condition != nil {
		if _, ok := ignore["cerbos.policy.v1.Condition.condition"]; !ok {
			switch t := m.Condition.(type) {
			case *v11.Condition_Match:
				if t.Match != nil {
					cerbos_policy_v1_Match_hashpb_sum(t.Match, hasher, ignore)
				}

			case *v11.Condition_Script:
				_, _ = hasher.Write(protowire.AppendString(nil, t.Script))

			}
		}
	}
}

func cerbos_policy_v1_DerivedRoles_hashpb_sum(m *v11.DerivedRoles, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.DerivedRoles.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
	if _, ok := ignore["cerbos.policy.v1.DerivedRoles.definitions"]; !ok {
		if len(m.Definitions) > 0 {
			for _, v := range m.Definitions {
				if v != nil {
					cerbos_policy_v1_RoleDef_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_policy_v1_Match_ExprList_hashpb_sum(m *v11.Match_ExprList, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Match.ExprList.of"]; !ok {
		if len(m.Of) > 0 {
			for _, v := range m.Of {
				if v != nil {
					cerbos_policy_v1_Match_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_policy_v1_Match_hashpb_sum(m *v11.Match, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Op != nil {
		if _, ok := ignore["cerbos.policy.v1.Match.op"]; !ok {
			switch t := m.Op.(type) {
			case *v11.Match_All:
				if t.All != nil {
					cerbos_policy_v1_Match_ExprList_hashpb_sum(t.All, hasher, ignore)
				}

			case *v11.Match_Any:
				if t.Any != nil {
					cerbos_policy_v1_Match_ExprList_hashpb_sum(t.Any, hasher, ignore)
				}

			case *v11.Match_None:
				if t.None != nil {
					cerbos_policy_v1_Match_ExprList_hashpb_sum(t.None, hasher, ignore)
				}

			case *v11.Match_Expr:
				_, _ = hasher.Write(protowire.AppendString(nil, t.Expr))

			}
		}
	}
}

func cerbos_policy_v1_Metadata_hashpb_sum(m *v11.Metadata, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Metadata.source_file"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.SourceFile))

	}
	if _, ok := ignore["cerbos.policy.v1.Metadata.annotations"]; !ok {
		if len(m.Annotations) > 0 {
			keys := make([]string, len(m.Annotations))
			i := 0
			for k := range m.Annotations {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendString(nil, m.Annotations[k]))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.Metadata.hash"]; !ok {
		if m.Hash != nil {
			google_protobuf_UInt64Value_hashpb_sum(m.Hash, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.Metadata.store_identifer"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.StoreIdentifer))

	}
}

func cerbos_policy_v1_Policy_hashpb_sum(m *v11.Policy, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Policy.api_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.ApiVersion))

	}
	if _, ok := ignore["cerbos.policy.v1.Policy.disabled"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(m.Disabled)))

	}
	if _, ok := ignore["cerbos.policy.v1.Policy.description"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Description))

	}
	if _, ok := ignore["cerbos.policy.v1.Policy.metadata"]; !ok {
		if m.Metadata != nil {
			cerbos_policy_v1_Metadata_hashpb_sum(m.Metadata, hasher, ignore)
		}

	}
	if m.PolicyType != nil {
		if _, ok := ignore["cerbos.policy.v1.Policy.policy_type"]; !ok {
			switch t := m.PolicyType.(type) {
			case *v11.Policy_ResourcePolicy:
				if t.ResourcePolicy != nil {
					cerbos_policy_v1_ResourcePolicy_hashpb_sum(t.ResourcePolicy, hasher, ignore)
				}

			case *v11.Policy_PrincipalPolicy:
				if t.PrincipalPolicy != nil {
					cerbos_policy_v1_PrincipalPolicy_hashpb_sum(t.PrincipalPolicy, hasher, ignore)
				}

			case *v11.Policy_DerivedRoles:
				if t.DerivedRoles != nil {
					cerbos_policy_v1_DerivedRoles_hashpb_sum(t.DerivedRoles, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.Policy.variables"]; !ok {
		if len(m.Variables) > 0 {
			keys := make([]string, len(m.Variables))
			i := 0
			for k := range m.Variables {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				_, _ = hasher.Write(protowire.AppendString(nil, m.Variables[k]))

			}
		}
	}
}

func cerbos_policy_v1_PrincipalPolicy_hashpb_sum(m *v11.PrincipalPolicy, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.PrincipalPolicy.principal"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Principal))

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalPolicy.version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Version))

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalPolicy.rules"]; !ok {
		if len(m.Rules) > 0 {
			for _, v := range m.Rules {
				if v != nil {
					cerbos_policy_v1_PrincipalRule_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalPolicy.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Scope))

	}
}

func cerbos_policy_v1_PrincipalRule_Action_hashpb_sum(m *v11.PrincipalRule_Action, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.PrincipalRule.Action.action"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Action))

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalRule.Action.condition"]; !ok {
		if m.Condition != nil {
			cerbos_policy_v1_Condition_hashpb_sum(m.Condition, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalRule.Action.effect"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Effect)))

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalRule.Action.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
}

func cerbos_policy_v1_PrincipalRule_hashpb_sum(m *v11.PrincipalRule, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.PrincipalRule.resource"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Resource))

	}
	if _, ok := ignore["cerbos.policy.v1.PrincipalRule.actions"]; !ok {
		if len(m.Actions) > 0 {
			for _, v := range m.Actions {
				if v != nil {
					cerbos_policy_v1_PrincipalRule_Action_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_policy_v1_ResourcePolicy_hashpb_sum(m *v11.ResourcePolicy, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.resource"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Resource))

	}
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Version))

	}
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.import_derived_roles"]; !ok {
		if len(m.ImportDerivedRoles) > 0 {
			for _, v := range m.ImportDerivedRoles {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.rules"]; !ok {
		if len(m.Rules) > 0 {
			for _, v := range m.Rules {
				if v != nil {
					cerbos_policy_v1_ResourceRule_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Scope))

	}
	if _, ok := ignore["cerbos.policy.v1.ResourcePolicy.schemas"]; !ok {
		if m.Schemas != nil {
			cerbos_policy_v1_Schemas_hashpb_sum(m.Schemas, hasher, ignore)
		}

	}
}

func cerbos_policy_v1_ResourceRule_hashpb_sum(m *v11.ResourceRule, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.ResourceRule.actions"]; !ok {
		if len(m.Actions) > 0 {
			for _, v := range m.Actions {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.ResourceRule.derived_roles"]; !ok {
		if len(m.DerivedRoles) > 0 {
			for _, v := range m.DerivedRoles {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.ResourceRule.roles"]; !ok {
		if len(m.Roles) > 0 {
			for _, v := range m.Roles {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.ResourceRule.condition"]; !ok {
		if m.Condition != nil {
			cerbos_policy_v1_Condition_hashpb_sum(m.Condition, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.ResourceRule.effect"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Effect)))

	}
	if _, ok := ignore["cerbos.policy.v1.ResourceRule.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
}

func cerbos_policy_v1_RoleDef_hashpb_sum(m *v11.RoleDef, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.RoleDef.name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Name))

	}
	if _, ok := ignore["cerbos.policy.v1.RoleDef.parent_roles"]; !ok {
		if len(m.ParentRoles) > 0 {
			for _, v := range m.ParentRoles {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.policy.v1.RoleDef.condition"]; !ok {
		if m.Condition != nil {
			cerbos_policy_v1_Condition_hashpb_sum(m.Condition, hasher, ignore)
		}

	}
}

func cerbos_policy_v1_Schemas_IgnoreWhen_hashpb_sum(m *v11.Schemas_IgnoreWhen, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Schemas.IgnoreWhen.actions"]; !ok {
		if len(m.Actions) > 0 {
			for _, v := range m.Actions {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
}

func cerbos_policy_v1_Schemas_Schema_hashpb_sum(m *v11.Schemas_Schema, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Schemas.Schema.ref"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Ref))

	}
	if _, ok := ignore["cerbos.policy.v1.Schemas.Schema.ignore_when"]; !ok {
		if m.IgnoreWhen != nil {
			cerbos_policy_v1_Schemas_IgnoreWhen_hashpb_sum(m.IgnoreWhen, hasher, ignore)
		}

	}
}

func cerbos_policy_v1_Schemas_hashpb_sum(m *v11.Schemas, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.policy.v1.Schemas.principal_schema"]; !ok {
		if m.PrincipalSchema != nil {
			cerbos_policy_v1_Schemas_Schema_hashpb_sum(m.PrincipalSchema, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.policy.v1.Schemas.resource_schema"]; !ok {
		if m.ResourceSchema != nil {
			cerbos_policy_v1_Schemas_Schema_hashpb_sum(m.ResourceSchema, hasher, ignore)
		}

	}
}

func cerbos_request_v1_AddOrUpdatePolicyRequest_hashpb_sum(m *AddOrUpdatePolicyRequest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.AddOrUpdatePolicyRequest.policies"]; !ok {
		if len(m.Policies) > 0 {
			for _, v := range m.Policies {
				if v != nil {
					cerbos_policy_v1_Policy_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_request_v1_AddOrUpdateSchemaRequest_hashpb_sum(m *AddOrUpdateSchemaRequest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.AddOrUpdateSchemaRequest.schemas"]; !ok {
		if len(m.Schemas) > 0 {
			for _, v := range m.Schemas {
				if v != nil {
					cerbos_schema_v1_Schema_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_request_v1_AttributesMap_hashpb_sum(m *AttributesMap, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.AttributesMap.attr"]; !ok {
		if len(m.Attr) > 0 {
			keys := make([]string, len(m.Attr))
			i := 0
			for k := range m.Attr {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Attr[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Attr[k], hasher, ignore)
				}

			}
		}
	}
}

func cerbos_request_v1_AuxData_JWT_hashpb_sum(m *AuxData_JWT, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.AuxData.JWT.token"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Token))

	}
	if _, ok := ignore["cerbos.request.v1.AuxData.JWT.key_set_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.KeySetId))

	}
}

func cerbos_request_v1_AuxData_hashpb_sum(m *AuxData, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.AuxData.jwt"]; !ok {
		if m.Jwt != nil {
			cerbos_request_v1_AuxData_JWT_hashpb_sum(m.Jwt, hasher, ignore)
		}

	}
}

func cerbos_request_v1_CheckResourceBatchRequest_BatchEntry_hashpb_sum(m *CheckResourceBatchRequest_BatchEntry, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.CheckResourceBatchRequest.BatchEntry.actions"]; !ok {
		if len(m.Actions) > 0 {
			for _, v := range m.Actions {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.request.v1.CheckResourceBatchRequest.BatchEntry.resource"]; !ok {
		if m.Resource != nil {
			cerbos_engine_v1_Resource_hashpb_sum(m.Resource, hasher, ignore)
		}

	}
}

func cerbos_request_v1_CheckResourceBatchRequest_hashpb_sum(m *CheckResourceBatchRequest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.CheckResourceBatchRequest.request_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.RequestId))

	}
	if _, ok := ignore["cerbos.request.v1.CheckResourceBatchRequest.principal"]; !ok {
		if m.Principal != nil {
			cerbos_engine_v1_Principal_hashpb_sum(m.Principal, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.request.v1.CheckResourceBatchRequest.resources"]; !ok {
		if len(m.Resources) > 0 {
			for _, v := range m.Resources {
				if v != nil {
					cerbos_request_v1_CheckResourceBatchRequest_BatchEntry_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.request.v1.CheckResourceBatchRequest.aux_data"]; !ok {
		if m.AuxData != nil {
			cerbos_request_v1_AuxData_hashpb_sum(m.AuxData, hasher, ignore)
		}

	}
}

func cerbos_request_v1_CheckResourceSetRequest_hashpb_sum(m *CheckResourceSetRequest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.CheckResourceSetRequest.request_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.RequestId))

	}
	if _, ok := ignore["cerbos.request.v1.CheckResourceSetRequest.actions"]; !ok {
		if len(m.Actions) > 0 {
			for _, v := range m.Actions {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.request.v1.CheckResourceSetRequest.principal"]; !ok {
		if m.Principal != nil {
			cerbos_engine_v1_Principal_hashpb_sum(m.Principal, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.request.v1.CheckResourceSetRequest.resource"]; !ok {
		if m.Resource != nil {
			cerbos_request_v1_ResourceSet_hashpb_sum(m.Resource, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.request.v1.CheckResourceSetRequest.include_meta"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(m.IncludeMeta)))

	}
	if _, ok := ignore["cerbos.request.v1.CheckResourceSetRequest.aux_data"]; !ok {
		if m.AuxData != nil {
			cerbos_request_v1_AuxData_hashpb_sum(m.AuxData, hasher, ignore)
		}

	}
}

func cerbos_request_v1_CheckResourcesRequest_ResourceEntry_hashpb_sum(m *CheckResourcesRequest_ResourceEntry, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.CheckResourcesRequest.ResourceEntry.actions"]; !ok {
		if len(m.Actions) > 0 {
			for _, v := range m.Actions {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.request.v1.CheckResourcesRequest.ResourceEntry.resource"]; !ok {
		if m.Resource != nil {
			cerbos_engine_v1_Resource_hashpb_sum(m.Resource, hasher, ignore)
		}

	}
}

func cerbos_request_v1_CheckResourcesRequest_hashpb_sum(m *CheckResourcesRequest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.CheckResourcesRequest.request_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.RequestId))

	}
	if _, ok := ignore["cerbos.request.v1.CheckResourcesRequest.include_meta"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(m.IncludeMeta)))

	}
	if _, ok := ignore["cerbos.request.v1.CheckResourcesRequest.principal"]; !ok {
		if m.Principal != nil {
			cerbos_engine_v1_Principal_hashpb_sum(m.Principal, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.request.v1.CheckResourcesRequest.resources"]; !ok {
		if len(m.Resources) > 0 {
			for _, v := range m.Resources {
				if v != nil {
					cerbos_request_v1_CheckResourcesRequest_ResourceEntry_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.request.v1.CheckResourcesRequest.aux_data"]; !ok {
		if m.AuxData != nil {
			cerbos_request_v1_AuxData_hashpb_sum(m.AuxData, hasher, ignore)
		}

	}
}

func cerbos_request_v1_DeleteSchemaRequest_hashpb_sum(m *DeleteSchemaRequest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.DeleteSchemaRequest.id"]; !ok {
		if len(m.Id) > 0 {
			for _, v := range m.Id {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
}

func cerbos_request_v1_File_hashpb_sum(m *File, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.File.file_name"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.FileName))

	}
	if _, ok := ignore["cerbos.request.v1.File.contents"]; !ok {
		_, _ = hasher.Write(protowire.AppendBytes(nil, m.Contents))

	}
}

func cerbos_request_v1_GetPolicyRequest_hashpb_sum(m *GetPolicyRequest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.GetPolicyRequest.id"]; !ok {
		if len(m.Id) > 0 {
			for _, v := range m.Id {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
}

func cerbos_request_v1_GetSchemaRequest_hashpb_sum(m *GetSchemaRequest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.GetSchemaRequest.id"]; !ok {
		if len(m.Id) > 0 {
			for _, v := range m.Id {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
}

func cerbos_request_v1_ListAuditLogEntriesRequest_TimeRange_hashpb_sum(m *ListAuditLogEntriesRequest_TimeRange, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.ListAuditLogEntriesRequest.TimeRange.start"]; !ok {
		if m.Start != nil {
			google_protobuf_Timestamp_hashpb_sum(m.Start, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.request.v1.ListAuditLogEntriesRequest.TimeRange.end"]; !ok {
		if m.End != nil {
			google_protobuf_Timestamp_hashpb_sum(m.End, hasher, ignore)
		}

	}
}

func cerbos_request_v1_ListAuditLogEntriesRequest_hashpb_sum(m *ListAuditLogEntriesRequest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.ListAuditLogEntriesRequest.kind"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Kind)))

	}
	if m.Filter != nil {
		if _, ok := ignore["cerbos.request.v1.ListAuditLogEntriesRequest.filter"]; !ok {
			switch t := m.Filter.(type) {
			case *ListAuditLogEntriesRequest_Tail:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.Tail)))

			case *ListAuditLogEntriesRequest_Between:
				if t.Between != nil {
					cerbos_request_v1_ListAuditLogEntriesRequest_TimeRange_hashpb_sum(t.Between, hasher, ignore)
				}

			case *ListAuditLogEntriesRequest_Since:
				if t.Since != nil {
					google_protobuf_Duration_hashpb_sum(t.Since, hasher, ignore)
				}

			case *ListAuditLogEntriesRequest_Lookup:
				_, _ = hasher.Write(protowire.AppendString(nil, t.Lookup))

			}
		}
	}
}

func cerbos_request_v1_ListPoliciesRequest_hashpb_sum(m *ListPoliciesRequest, hasher hash.Hash, ignore map[string]struct{}) {
}

func cerbos_request_v1_ListSchemasRequest_hashpb_sum(m *ListSchemasRequest, hasher hash.Hash, ignore map[string]struct{}) {
}

func cerbos_request_v1_PlaygroundEvaluateRequest_hashpb_sum(m *PlaygroundEvaluateRequest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.PlaygroundEvaluateRequest.playground_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.PlaygroundId))

	}
	if _, ok := ignore["cerbos.request.v1.PlaygroundEvaluateRequest.files"]; !ok {
		if len(m.Files) > 0 {
			for _, v := range m.Files {
				if v != nil {
					cerbos_request_v1_File_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.request.v1.PlaygroundEvaluateRequest.principal"]; !ok {
		if m.Principal != nil {
			cerbos_engine_v1_Principal_hashpb_sum(m.Principal, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.request.v1.PlaygroundEvaluateRequest.resource"]; !ok {
		if m.Resource != nil {
			cerbos_engine_v1_Resource_hashpb_sum(m.Resource, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.request.v1.PlaygroundEvaluateRequest.actions"]; !ok {
		if len(m.Actions) > 0 {
			for _, v := range m.Actions {
				_, _ = hasher.Write(protowire.AppendString(nil, v))

			}
		}
	}
	if _, ok := ignore["cerbos.request.v1.PlaygroundEvaluateRequest.aux_data"]; !ok {
		if m.AuxData != nil {
			cerbos_request_v1_AuxData_hashpb_sum(m.AuxData, hasher, ignore)
		}

	}
}

func cerbos_request_v1_PlaygroundProxyRequest_hashpb_sum(m *PlaygroundProxyRequest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.PlaygroundProxyRequest.playground_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.PlaygroundId))

	}
	if _, ok := ignore["cerbos.request.v1.PlaygroundProxyRequest.files"]; !ok {
		if len(m.Files) > 0 {
			for _, v := range m.Files {
				if v != nil {
					cerbos_request_v1_File_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
	if m.ProxyRequest != nil {
		if _, ok := ignore["cerbos.request.v1.PlaygroundProxyRequest.proxy_request"]; !ok {
			switch t := m.ProxyRequest.(type) {
			case *PlaygroundProxyRequest_CheckResourceSet:
				if t.CheckResourceSet != nil {
					cerbos_request_v1_CheckResourceSetRequest_hashpb_sum(t.CheckResourceSet, hasher, ignore)
				}

			case *PlaygroundProxyRequest_CheckResourceBatch:
				if t.CheckResourceBatch != nil {
					cerbos_request_v1_CheckResourceBatchRequest_hashpb_sum(t.CheckResourceBatch, hasher, ignore)
				}

			case *PlaygroundProxyRequest_ResourcesQueryPlan:
				if t.ResourcesQueryPlan != nil {
					cerbos_request_v1_ResourcesQueryPlanRequest_hashpb_sum(t.ResourcesQueryPlan, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_request_v1_PlaygroundTestRequest_hashpb_sum(m *PlaygroundTestRequest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.PlaygroundTestRequest.playground_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.PlaygroundId))

	}
	if _, ok := ignore["cerbos.request.v1.PlaygroundTestRequest.files"]; !ok {
		if len(m.Files) > 0 {
			for _, v := range m.Files {
				if v != nil {
					cerbos_request_v1_File_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_request_v1_PlaygroundValidateRequest_hashpb_sum(m *PlaygroundValidateRequest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.PlaygroundValidateRequest.playground_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.PlaygroundId))

	}
	if _, ok := ignore["cerbos.request.v1.PlaygroundValidateRequest.files"]; !ok {
		if len(m.Files) > 0 {
			for _, v := range m.Files {
				if v != nil {
					cerbos_request_v1_File_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func cerbos_request_v1_ReloadStoreRequest_hashpb_sum(m *ReloadStoreRequest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.ReloadStoreRequest.wait"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(m.Wait)))

	}
}

func cerbos_request_v1_ResourceSet_hashpb_sum(m *ResourceSet, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.ResourceSet.kind"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Kind))

	}
	if _, ok := ignore["cerbos.request.v1.ResourceSet.policy_version"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.PolicyVersion))

	}
	if _, ok := ignore["cerbos.request.v1.ResourceSet.instances"]; !ok {
		if len(m.Instances) > 0 {
			keys := make([]string, len(m.Instances))
			i := 0
			for k := range m.Instances {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Instances[k] != nil {
					cerbos_request_v1_AttributesMap_hashpb_sum(m.Instances[k], hasher, ignore)
				}

			}
		}
	}
	if _, ok := ignore["cerbos.request.v1.ResourceSet.scope"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Scope))

	}
}

func cerbos_request_v1_ResourcesQueryPlanRequest_hashpb_sum(m *ResourcesQueryPlanRequest, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.request.v1.ResourcesQueryPlanRequest.request_id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.RequestId))

	}
	if _, ok := ignore["cerbos.request.v1.ResourcesQueryPlanRequest.action"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Action))

	}
	if _, ok := ignore["cerbos.request.v1.ResourcesQueryPlanRequest.principal"]; !ok {
		if m.Principal != nil {
			cerbos_engine_v1_Principal_hashpb_sum(m.Principal, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.request.v1.ResourcesQueryPlanRequest.resource"]; !ok {
		if m.Resource != nil {
			cerbos_engine_v1_ResourcesQueryPlanRequest_Resource_hashpb_sum(m.Resource, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.request.v1.ResourcesQueryPlanRequest.aux_data"]; !ok {
		if m.AuxData != nil {
			cerbos_request_v1_AuxData_hashpb_sum(m.AuxData, hasher, ignore)
		}

	}
	if _, ok := ignore["cerbos.request.v1.ResourcesQueryPlanRequest.include_meta"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(m.IncludeMeta)))

	}
}

func cerbos_request_v1_ServerInfoRequest_hashpb_sum(m *ServerInfoRequest, hasher hash.Hash, ignore map[string]struct{}) {
}

func cerbos_schema_v1_Schema_hashpb_sum(m *v12.Schema, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["cerbos.schema.v1.Schema.id"]; !ok {
		_, _ = hasher.Write(protowire.AppendString(nil, m.Id))

	}
	if _, ok := ignore["cerbos.schema.v1.Schema.definition"]; !ok {
		_, _ = hasher.Write(protowire.AppendBytes(nil, m.Definition))

	}
}

func google_protobuf_Duration_hashpb_sum(m *durationpb.Duration, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.Duration.seconds"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Seconds)))

	}
	if _, ok := ignore["google.protobuf.Duration.nanos"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Nanos)))

	}
}

func google_protobuf_ListValue_hashpb_sum(m *structpb.ListValue, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.ListValue.values"]; !ok {
		if len(m.Values) > 0 {
			for _, v := range m.Values {
				if v != nil {
					google_protobuf_Value_hashpb_sum(v, hasher, ignore)
				}

			}
		}
	}
}

func google_protobuf_Struct_hashpb_sum(m *structpb.Struct, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.Struct.fields"]; !ok {
		if len(m.Fields) > 0 {
			keys := make([]string, len(m.Fields))
			i := 0
			for k := range m.Fields {
				keys[i] = k
				i++
			}

			sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

			for _, k := range keys {
				if m.Fields[k] != nil {
					google_protobuf_Value_hashpb_sum(m.Fields[k], hasher, ignore)
				}

			}
		}
	}
}

func google_protobuf_Timestamp_hashpb_sum(m *timestamppb.Timestamp, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.Timestamp.seconds"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Seconds)))

	}
	if _, ok := ignore["google.protobuf.Timestamp.nanos"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(m.Nanos)))

	}
}

func google_protobuf_UInt64Value_hashpb_sum(m *wrapperspb.UInt64Value, hasher hash.Hash, ignore map[string]struct{}) {
	if _, ok := ignore["google.protobuf.UInt64Value.value"]; !ok {
		_, _ = hasher.Write(protowire.AppendVarint(nil, m.Value))

	}
}

func google_protobuf_Value_hashpb_sum(m *structpb.Value, hasher hash.Hash, ignore map[string]struct{}) {
	if m.Kind != nil {
		if _, ok := ignore["google.protobuf.Value.kind"]; !ok {
			switch t := m.Kind.(type) {
			case *structpb.Value_NullValue:
				_, _ = hasher.Write(protowire.AppendVarint(nil, uint64(t.NullValue)))

			case *structpb.Value_NumberValue:
				_, _ = hasher.Write(protowire.AppendFixed64(nil, math.Float64bits(t.NumberValue)))

			case *structpb.Value_StringValue:
				_, _ = hasher.Write(protowire.AppendString(nil, t.StringValue))

			case *structpb.Value_BoolValue:
				_, _ = hasher.Write(protowire.AppendVarint(nil, protowire.EncodeBool(t.BoolValue)))

			case *structpb.Value_StructValue:
				if t.StructValue != nil {
					google_protobuf_Struct_hashpb_sum(t.StructValue, hasher, ignore)
				}

			case *structpb.Value_ListValue:
				if t.ListValue != nil {
					google_protobuf_ListValue_hashpb_sum(t.ListValue, hasher, ignore)
				}

			}
		}
	}
}
