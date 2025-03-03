// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy. Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type InternalRangeIamPolicy struct {
	pulumi.CustomResourceState

	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigResponseArrayOutput `pulumi:"auditConfigs"`
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings BindingResponseArrayOutput `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag            pulumi.StringOutput `pulumi:"etag"`
	InternalRangeId pulumi.StringOutput `pulumi:"internalRangeId"`
	Location        pulumi.StringOutput `pulumi:"location"`
	Project         pulumi.StringOutput `pulumi:"project"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewInternalRangeIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewInternalRangeIamPolicy(ctx *pulumi.Context,
	name string, args *InternalRangeIamPolicyArgs, opts ...pulumi.ResourceOption) (*InternalRangeIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InternalRangeId == nil {
		return nil, errors.New("invalid value for required argument 'InternalRangeId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"internalRangeId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource InternalRangeIamPolicy
	err := ctx.RegisterResource("google-native:networkconnectivity/v1alpha1:InternalRangeIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInternalRangeIamPolicy gets an existing InternalRangeIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInternalRangeIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InternalRangeIamPolicyState, opts ...pulumi.ResourceOption) (*InternalRangeIamPolicy, error) {
	var resource InternalRangeIamPolicy
	err := ctx.ReadResource("google-native:networkconnectivity/v1alpha1:InternalRangeIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InternalRangeIamPolicy resources.
type internalRangeIamPolicyState struct {
}

type InternalRangeIamPolicyState struct {
}

func (InternalRangeIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*internalRangeIamPolicyState)(nil)).Elem()
}

type internalRangeIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []AuditConfig `pulumi:"auditConfigs"`
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings []Binding `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag            *string `pulumi:"etag"`
	InternalRangeId string  `pulumi:"internalRangeId"`
	Location        *string `pulumi:"location"`
	Project         *string `pulumi:"project"`
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask *string `pulumi:"updateMask"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a InternalRangeIamPolicy resource.
type InternalRangeIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigArrayInput
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings BindingArrayInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag            pulumi.StringPtrInput
	InternalRangeId pulumi.StringInput
	Location        pulumi.StringPtrInput
	Project         pulumi.StringPtrInput
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask pulumi.StringPtrInput
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (InternalRangeIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*internalRangeIamPolicyArgs)(nil)).Elem()
}

type InternalRangeIamPolicyInput interface {
	pulumi.Input

	ToInternalRangeIamPolicyOutput() InternalRangeIamPolicyOutput
	ToInternalRangeIamPolicyOutputWithContext(ctx context.Context) InternalRangeIamPolicyOutput
}

func (*InternalRangeIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**InternalRangeIamPolicy)(nil)).Elem()
}

func (i *InternalRangeIamPolicy) ToInternalRangeIamPolicyOutput() InternalRangeIamPolicyOutput {
	return i.ToInternalRangeIamPolicyOutputWithContext(context.Background())
}

func (i *InternalRangeIamPolicy) ToInternalRangeIamPolicyOutputWithContext(ctx context.Context) InternalRangeIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InternalRangeIamPolicyOutput)
}

type InternalRangeIamPolicyOutput struct{ *pulumi.OutputState }

func (InternalRangeIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InternalRangeIamPolicy)(nil)).Elem()
}

func (o InternalRangeIamPolicyOutput) ToInternalRangeIamPolicyOutput() InternalRangeIamPolicyOutput {
	return o
}

func (o InternalRangeIamPolicyOutput) ToInternalRangeIamPolicyOutputWithContext(ctx context.Context) InternalRangeIamPolicyOutput {
	return o
}

// Specifies cloud audit logging configuration for this policy.
func (o InternalRangeIamPolicyOutput) AuditConfigs() AuditConfigResponseArrayOutput {
	return o.ApplyT(func(v *InternalRangeIamPolicy) AuditConfigResponseArrayOutput { return v.AuditConfigs }).(AuditConfigResponseArrayOutput)
}

// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
func (o InternalRangeIamPolicyOutput) Bindings() BindingResponseArrayOutput {
	return o.ApplyT(func(v *InternalRangeIamPolicy) BindingResponseArrayOutput { return v.Bindings }).(BindingResponseArrayOutput)
}

// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
func (o InternalRangeIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *InternalRangeIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o InternalRangeIamPolicyOutput) InternalRangeId() pulumi.StringOutput {
	return o.ApplyT(func(v *InternalRangeIamPolicy) pulumi.StringOutput { return v.InternalRangeId }).(pulumi.StringOutput)
}

func (o InternalRangeIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *InternalRangeIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o InternalRangeIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *InternalRangeIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
func (o InternalRangeIamPolicyOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *InternalRangeIamPolicy) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InternalRangeIamPolicyInput)(nil)).Elem(), &InternalRangeIamPolicy{})
	pulumi.RegisterOutputType(InternalRangeIamPolicyOutput{})
}
