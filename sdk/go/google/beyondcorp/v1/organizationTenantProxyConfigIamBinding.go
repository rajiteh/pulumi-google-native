// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	iam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy. Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
type OrganizationTenantProxyConfigIamBinding struct {
	pulumi.CustomResourceState

	// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
	Condition iam.ConditionPtrOutput `pulumi:"condition"`
	// The etag of the resource's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Identities that will be granted the privilege in role. Each entry can have one of the following values:
	//
	//  * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	//  * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	//  * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
	//  * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The name of the resource to manage IAM policies for.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project in which the resource belongs. If it is not provided, a default will be supplied.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one `IamBinding` can be used per role.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewOrganizationTenantProxyConfigIamBinding registers a new resource with the given unique name, arguments, and options.
func NewOrganizationTenantProxyConfigIamBinding(ctx *pulumi.Context,
	name string, args *OrganizationTenantProxyConfigIamBindingArgs, opts ...pulumi.ResourceOption) (*OrganizationTenantProxyConfigIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource OrganizationTenantProxyConfigIamBinding
	err := ctx.RegisterResource("google-native:beyondcorp/v1:OrganizationTenantProxyConfigIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationTenantProxyConfigIamBinding gets an existing OrganizationTenantProxyConfigIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationTenantProxyConfigIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationTenantProxyConfigIamBindingState, opts ...pulumi.ResourceOption) (*OrganizationTenantProxyConfigIamBinding, error) {
	var resource OrganizationTenantProxyConfigIamBinding
	err := ctx.ReadResource("google-native:beyondcorp/v1:OrganizationTenantProxyConfigIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationTenantProxyConfigIamBinding resources.
type organizationTenantProxyConfigIamBindingState struct {
}

type OrganizationTenantProxyConfigIamBindingState struct {
}

func (OrganizationTenantProxyConfigIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationTenantProxyConfigIamBindingState)(nil)).Elem()
}

type organizationTenantProxyConfigIamBindingArgs struct {
	// An IAM Condition for a given binding.
	Condition *iam.Condition `pulumi:"condition"`
	// Identities that will be granted the privilege in role. Each entry can have one of the following values:
	//
	//  * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	//  * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	//  * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
	//  * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Members []string `pulumi:"members"`
	// The name of the resource to manage IAM policies for.
	Name string `pulumi:"name"`
	// The role that should be applied. Only one `IamBinding` can be used per role.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a OrganizationTenantProxyConfigIamBinding resource.
type OrganizationTenantProxyConfigIamBindingArgs struct {
	// An IAM Condition for a given binding.
	Condition iam.ConditionPtrInput
	// Identities that will be granted the privilege in role. Each entry can have one of the following values:
	//
	//  * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	//  * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	//  * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
	//  * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Members pulumi.StringArrayInput
	// The name of the resource to manage IAM policies for.
	Name pulumi.StringInput
	// The role that should be applied. Only one `IamBinding` can be used per role.
	Role pulumi.StringInput
}

func (OrganizationTenantProxyConfigIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationTenantProxyConfigIamBindingArgs)(nil)).Elem()
}

type OrganizationTenantProxyConfigIamBindingInput interface {
	pulumi.Input

	ToOrganizationTenantProxyConfigIamBindingOutput() OrganizationTenantProxyConfigIamBindingOutput
	ToOrganizationTenantProxyConfigIamBindingOutputWithContext(ctx context.Context) OrganizationTenantProxyConfigIamBindingOutput
}

func (*OrganizationTenantProxyConfigIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationTenantProxyConfigIamBinding)(nil)).Elem()
}

func (i *OrganizationTenantProxyConfigIamBinding) ToOrganizationTenantProxyConfigIamBindingOutput() OrganizationTenantProxyConfigIamBindingOutput {
	return i.ToOrganizationTenantProxyConfigIamBindingOutputWithContext(context.Background())
}

func (i *OrganizationTenantProxyConfigIamBinding) ToOrganizationTenantProxyConfigIamBindingOutputWithContext(ctx context.Context) OrganizationTenantProxyConfigIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationTenantProxyConfigIamBindingOutput)
}

type OrganizationTenantProxyConfigIamBindingOutput struct{ *pulumi.OutputState }

func (OrganizationTenantProxyConfigIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationTenantProxyConfigIamBinding)(nil)).Elem()
}

func (o OrganizationTenantProxyConfigIamBindingOutput) ToOrganizationTenantProxyConfigIamBindingOutput() OrganizationTenantProxyConfigIamBindingOutput {
	return o
}

func (o OrganizationTenantProxyConfigIamBindingOutput) ToOrganizationTenantProxyConfigIamBindingOutputWithContext(ctx context.Context) OrganizationTenantProxyConfigIamBindingOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o OrganizationTenantProxyConfigIamBindingOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *OrganizationTenantProxyConfigIamBinding) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o OrganizationTenantProxyConfigIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationTenantProxyConfigIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Identities that will be granted the privilege in role. Each entry can have one of the following values:
//
//   - user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
//   - domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o OrganizationTenantProxyConfigIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrganizationTenantProxyConfigIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The name of the resource to manage IAM policies for.
func (o OrganizationTenantProxyConfigIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationTenantProxyConfigIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o OrganizationTenantProxyConfigIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationTenantProxyConfigIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one `IamBinding` can be used per role.
func (o OrganizationTenantProxyConfigIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationTenantProxyConfigIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationTenantProxyConfigIamBindingInput)(nil)).Elem(), &OrganizationTenantProxyConfigIamBinding{})
	pulumi.RegisterOutputType(OrganizationTenantProxyConfigIamBindingOutput{})
}
