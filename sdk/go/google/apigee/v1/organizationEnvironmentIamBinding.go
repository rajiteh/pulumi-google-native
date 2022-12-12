// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	iam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the IAM policy on an environment, if the policy already exists it will be replaced. For more information, see [Manage users, roles, and permissions using the API](https://cloud.google.com/apigee/docs/api-platform/system-administration/manage-users-roles). You must have the `apigee.environments.setIamPolicy` permission to call this API.
type OrganizationEnvironmentIamBinding struct {
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

// NewOrganizationEnvironmentIamBinding registers a new resource with the given unique name, arguments, and options.
func NewOrganizationEnvironmentIamBinding(ctx *pulumi.Context,
	name string, args *OrganizationEnvironmentIamBindingArgs, opts ...pulumi.ResourceOption) (*OrganizationEnvironmentIamBinding, error) {
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
	var resource OrganizationEnvironmentIamBinding
	err := ctx.RegisterResource("google-native:apigee/v1:OrganizationEnvironmentIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationEnvironmentIamBinding gets an existing OrganizationEnvironmentIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationEnvironmentIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationEnvironmentIamBindingState, opts ...pulumi.ResourceOption) (*OrganizationEnvironmentIamBinding, error) {
	var resource OrganizationEnvironmentIamBinding
	err := ctx.ReadResource("google-native:apigee/v1:OrganizationEnvironmentIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationEnvironmentIamBinding resources.
type organizationEnvironmentIamBindingState struct {
}

type OrganizationEnvironmentIamBindingState struct {
}

func (OrganizationEnvironmentIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationEnvironmentIamBindingState)(nil)).Elem()
}

type organizationEnvironmentIamBindingArgs struct {
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

// The set of arguments for constructing a OrganizationEnvironmentIamBinding resource.
type OrganizationEnvironmentIamBindingArgs struct {
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

func (OrganizationEnvironmentIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationEnvironmentIamBindingArgs)(nil)).Elem()
}

type OrganizationEnvironmentIamBindingInput interface {
	pulumi.Input

	ToOrganizationEnvironmentIamBindingOutput() OrganizationEnvironmentIamBindingOutput
	ToOrganizationEnvironmentIamBindingOutputWithContext(ctx context.Context) OrganizationEnvironmentIamBindingOutput
}

func (*OrganizationEnvironmentIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationEnvironmentIamBinding)(nil)).Elem()
}

func (i *OrganizationEnvironmentIamBinding) ToOrganizationEnvironmentIamBindingOutput() OrganizationEnvironmentIamBindingOutput {
	return i.ToOrganizationEnvironmentIamBindingOutputWithContext(context.Background())
}

func (i *OrganizationEnvironmentIamBinding) ToOrganizationEnvironmentIamBindingOutputWithContext(ctx context.Context) OrganizationEnvironmentIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationEnvironmentIamBindingOutput)
}

type OrganizationEnvironmentIamBindingOutput struct{ *pulumi.OutputState }

func (OrganizationEnvironmentIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationEnvironmentIamBinding)(nil)).Elem()
}

func (o OrganizationEnvironmentIamBindingOutput) ToOrganizationEnvironmentIamBindingOutput() OrganizationEnvironmentIamBindingOutput {
	return o
}

func (o OrganizationEnvironmentIamBindingOutput) ToOrganizationEnvironmentIamBindingOutputWithContext(ctx context.Context) OrganizationEnvironmentIamBindingOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o OrganizationEnvironmentIamBindingOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *OrganizationEnvironmentIamBinding) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o OrganizationEnvironmentIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationEnvironmentIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Identities that will be granted the privilege in role. Each entry can have one of the following values:
//
//   - user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
//   - domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o OrganizationEnvironmentIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrganizationEnvironmentIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The name of the resource to manage IAM policies for.
func (o OrganizationEnvironmentIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationEnvironmentIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o OrganizationEnvironmentIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationEnvironmentIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one `IamBinding` can be used per role.
func (o OrganizationEnvironmentIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationEnvironmentIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationEnvironmentIamBindingInput)(nil)).Elem(), &OrganizationEnvironmentIamBinding{})
	pulumi.RegisterOutputType(OrganizationEnvironmentIamBindingOutput{})
}