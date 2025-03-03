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

// Sets the access control policy on the specified resource. Replaces any existing policy.Can return NOT_FOUND, INVALID_ARGUMENT, and PERMISSION_DENIED errors.
type EntryTypeIamBinding struct {
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

// NewEntryTypeIamBinding registers a new resource with the given unique name, arguments, and options.
func NewEntryTypeIamBinding(ctx *pulumi.Context,
	name string, args *EntryTypeIamBindingArgs, opts ...pulumi.ResourceOption) (*EntryTypeIamBinding, error) {
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
	var resource EntryTypeIamBinding
	err := ctx.RegisterResource("google-native:dataplex/v1:EntryTypeIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntryTypeIamBinding gets an existing EntryTypeIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntryTypeIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntryTypeIamBindingState, opts ...pulumi.ResourceOption) (*EntryTypeIamBinding, error) {
	var resource EntryTypeIamBinding
	err := ctx.ReadResource("google-native:dataplex/v1:EntryTypeIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EntryTypeIamBinding resources.
type entryTypeIamBindingState struct {
}

type EntryTypeIamBindingState struct {
}

func (EntryTypeIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*entryTypeIamBindingState)(nil)).Elem()
}

type entryTypeIamBindingArgs struct {
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

// The set of arguments for constructing a EntryTypeIamBinding resource.
type EntryTypeIamBindingArgs struct {
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

func (EntryTypeIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entryTypeIamBindingArgs)(nil)).Elem()
}

type EntryTypeIamBindingInput interface {
	pulumi.Input

	ToEntryTypeIamBindingOutput() EntryTypeIamBindingOutput
	ToEntryTypeIamBindingOutputWithContext(ctx context.Context) EntryTypeIamBindingOutput
}

func (*EntryTypeIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**EntryTypeIamBinding)(nil)).Elem()
}

func (i *EntryTypeIamBinding) ToEntryTypeIamBindingOutput() EntryTypeIamBindingOutput {
	return i.ToEntryTypeIamBindingOutputWithContext(context.Background())
}

func (i *EntryTypeIamBinding) ToEntryTypeIamBindingOutputWithContext(ctx context.Context) EntryTypeIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntryTypeIamBindingOutput)
}

type EntryTypeIamBindingOutput struct{ *pulumi.OutputState }

func (EntryTypeIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntryTypeIamBinding)(nil)).Elem()
}

func (o EntryTypeIamBindingOutput) ToEntryTypeIamBindingOutput() EntryTypeIamBindingOutput {
	return o
}

func (o EntryTypeIamBindingOutput) ToEntryTypeIamBindingOutputWithContext(ctx context.Context) EntryTypeIamBindingOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o EntryTypeIamBindingOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *EntryTypeIamBinding) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o EntryTypeIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryTypeIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Identities that will be granted the privilege in role. Each entry can have one of the following values:
//
//   - user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
//   - domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o EntryTypeIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EntryTypeIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

// The name of the resource to manage IAM policies for.
func (o EntryTypeIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryTypeIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o EntryTypeIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryTypeIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied. Only one `IamBinding` can be used per role.
func (o EntryTypeIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryTypeIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EntryTypeIamBindingInput)(nil)).Elem(), &EntryTypeIamBinding{})
	pulumi.RegisterOutputType(EntryTypeIamBindingOutput{})
}
