// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy. Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type OrganizationTenantProxyConfigIamPolicy struct {
	pulumi.CustomResourceState

	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs GoogleIamV1AuditConfigResponseArrayOutput `pulumi:"auditConfigs"`
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings GoogleIamV1BindingResponseArrayOutput `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag           pulumi.StringOutput `pulumi:"etag"`
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	ProxyConfigId  pulumi.StringOutput `pulumi:"proxyConfigId"`
	TenantId       pulumi.StringOutput `pulumi:"tenantId"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewOrganizationTenantProxyConfigIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewOrganizationTenantProxyConfigIamPolicy(ctx *pulumi.Context,
	name string, args *OrganizationTenantProxyConfigIamPolicyArgs, opts ...pulumi.ResourceOption) (*OrganizationTenantProxyConfigIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	if args.ProxyConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ProxyConfigId'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"organizationId",
		"proxyConfigId",
		"tenantId",
	})
	opts = append(opts, replaceOnChanges)
	var resource OrganizationTenantProxyConfigIamPolicy
	err := ctx.RegisterResource("google-native:beyondcorp/v1:OrganizationTenantProxyConfigIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationTenantProxyConfigIamPolicy gets an existing OrganizationTenantProxyConfigIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationTenantProxyConfigIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationTenantProxyConfigIamPolicyState, opts ...pulumi.ResourceOption) (*OrganizationTenantProxyConfigIamPolicy, error) {
	var resource OrganizationTenantProxyConfigIamPolicy
	err := ctx.ReadResource("google-native:beyondcorp/v1:OrganizationTenantProxyConfigIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationTenantProxyConfigIamPolicy resources.
type organizationTenantProxyConfigIamPolicyState struct {
}

type OrganizationTenantProxyConfigIamPolicyState struct {
}

func (OrganizationTenantProxyConfigIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationTenantProxyConfigIamPolicyState)(nil)).Elem()
}

type organizationTenantProxyConfigIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []GoogleIamV1AuditConfig `pulumi:"auditConfigs"`
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings []GoogleIamV1Binding `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag           *string `pulumi:"etag"`
	OrganizationId string  `pulumi:"organizationId"`
	ProxyConfigId  string  `pulumi:"proxyConfigId"`
	TenantId       string  `pulumi:"tenantId"`
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask *string `pulumi:"updateMask"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a OrganizationTenantProxyConfigIamPolicy resource.
type OrganizationTenantProxyConfigIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs GoogleIamV1AuditConfigArrayInput
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings GoogleIamV1BindingArrayInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag           pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
	ProxyConfigId  pulumi.StringInput
	TenantId       pulumi.StringInput
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask pulumi.StringPtrInput
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (OrganizationTenantProxyConfigIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationTenantProxyConfigIamPolicyArgs)(nil)).Elem()
}

type OrganizationTenantProxyConfigIamPolicyInput interface {
	pulumi.Input

	ToOrganizationTenantProxyConfigIamPolicyOutput() OrganizationTenantProxyConfigIamPolicyOutput
	ToOrganizationTenantProxyConfigIamPolicyOutputWithContext(ctx context.Context) OrganizationTenantProxyConfigIamPolicyOutput
}

func (*OrganizationTenantProxyConfigIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationTenantProxyConfigIamPolicy)(nil)).Elem()
}

func (i *OrganizationTenantProxyConfigIamPolicy) ToOrganizationTenantProxyConfigIamPolicyOutput() OrganizationTenantProxyConfigIamPolicyOutput {
	return i.ToOrganizationTenantProxyConfigIamPolicyOutputWithContext(context.Background())
}

func (i *OrganizationTenantProxyConfigIamPolicy) ToOrganizationTenantProxyConfigIamPolicyOutputWithContext(ctx context.Context) OrganizationTenantProxyConfigIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationTenantProxyConfigIamPolicyOutput)
}

type OrganizationTenantProxyConfigIamPolicyOutput struct{ *pulumi.OutputState }

func (OrganizationTenantProxyConfigIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationTenantProxyConfigIamPolicy)(nil)).Elem()
}

func (o OrganizationTenantProxyConfigIamPolicyOutput) ToOrganizationTenantProxyConfigIamPolicyOutput() OrganizationTenantProxyConfigIamPolicyOutput {
	return o
}

func (o OrganizationTenantProxyConfigIamPolicyOutput) ToOrganizationTenantProxyConfigIamPolicyOutputWithContext(ctx context.Context) OrganizationTenantProxyConfigIamPolicyOutput {
	return o
}

// Specifies cloud audit logging configuration for this policy.
func (o OrganizationTenantProxyConfigIamPolicyOutput) AuditConfigs() GoogleIamV1AuditConfigResponseArrayOutput {
	return o.ApplyT(func(v *OrganizationTenantProxyConfigIamPolicy) GoogleIamV1AuditConfigResponseArrayOutput {
		return v.AuditConfigs
	}).(GoogleIamV1AuditConfigResponseArrayOutput)
}

// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
func (o OrganizationTenantProxyConfigIamPolicyOutput) Bindings() GoogleIamV1BindingResponseArrayOutput {
	return o.ApplyT(func(v *OrganizationTenantProxyConfigIamPolicy) GoogleIamV1BindingResponseArrayOutput {
		return v.Bindings
	}).(GoogleIamV1BindingResponseArrayOutput)
}

// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
func (o OrganizationTenantProxyConfigIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationTenantProxyConfigIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o OrganizationTenantProxyConfigIamPolicyOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationTenantProxyConfigIamPolicy) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o OrganizationTenantProxyConfigIamPolicyOutput) ProxyConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationTenantProxyConfigIamPolicy) pulumi.StringOutput { return v.ProxyConfigId }).(pulumi.StringOutput)
}

func (o OrganizationTenantProxyConfigIamPolicyOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationTenantProxyConfigIamPolicy) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
func (o OrganizationTenantProxyConfigIamPolicyOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *OrganizationTenantProxyConfigIamPolicy) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationTenantProxyConfigIamPolicyInput)(nil)).Elem(), &OrganizationTenantProxyConfigIamPolicy{})
	pulumi.RegisterOutputType(OrganizationTenantProxyConfigIamPolicyOutput{})
}
