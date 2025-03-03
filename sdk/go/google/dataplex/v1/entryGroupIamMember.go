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
type EntryGroupIamMember struct {
	pulumi.CustomResourceState

	// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
	Condition iam.ConditionPtrOutput `pulumi:"condition"`
	// The etag of the resource's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Identity that will be granted the privilege in role. The entry can have one of the following values:
	//
	//  * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	//  * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	//  * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
	//  * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Member pulumi.StringOutput `pulumi:"member"`
	// The name of the resource to manage IAM policies for.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project in which the resource belongs. If it is not provided, a default will be supplied.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewEntryGroupIamMember registers a new resource with the given unique name, arguments, and options.
func NewEntryGroupIamMember(ctx *pulumi.Context,
	name string, args *EntryGroupIamMemberArgs, opts ...pulumi.ResourceOption) (*EntryGroupIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource EntryGroupIamMember
	err := ctx.RegisterResource("google-native:dataplex/v1:EntryGroupIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEntryGroupIamMember gets an existing EntryGroupIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEntryGroupIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EntryGroupIamMemberState, opts ...pulumi.ResourceOption) (*EntryGroupIamMember, error) {
	var resource EntryGroupIamMember
	err := ctx.ReadResource("google-native:dataplex/v1:EntryGroupIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EntryGroupIamMember resources.
type entryGroupIamMemberState struct {
}

type EntryGroupIamMemberState struct {
}

func (EntryGroupIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*entryGroupIamMemberState)(nil)).Elem()
}

type entryGroupIamMemberArgs struct {
	// An IAM Condition for a given binding.
	Condition *iam.Condition `pulumi:"condition"`
	// Identity that will be granted the privilege in role. The entry can have one of the following values:
	//
	//  * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	//  * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	//  * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
	//  * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Member string `pulumi:"member"`
	// The name of the resource to manage IAM policies for.
	Name string `pulumi:"name"`
	// The role that should be applied.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a EntryGroupIamMember resource.
type EntryGroupIamMemberArgs struct {
	// An IAM Condition for a given binding.
	Condition iam.ConditionPtrInput
	// Identity that will be granted the privilege in role. The entry can have one of the following values:
	//
	//  * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	//  * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	//  * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
	//  * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Member pulumi.StringInput
	// The name of the resource to manage IAM policies for.
	Name pulumi.StringInput
	// The role that should be applied.
	Role pulumi.StringInput
}

func (EntryGroupIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*entryGroupIamMemberArgs)(nil)).Elem()
}

type EntryGroupIamMemberInput interface {
	pulumi.Input

	ToEntryGroupIamMemberOutput() EntryGroupIamMemberOutput
	ToEntryGroupIamMemberOutputWithContext(ctx context.Context) EntryGroupIamMemberOutput
}

func (*EntryGroupIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**EntryGroupIamMember)(nil)).Elem()
}

func (i *EntryGroupIamMember) ToEntryGroupIamMemberOutput() EntryGroupIamMemberOutput {
	return i.ToEntryGroupIamMemberOutputWithContext(context.Background())
}

func (i *EntryGroupIamMember) ToEntryGroupIamMemberOutputWithContext(ctx context.Context) EntryGroupIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntryGroupIamMemberOutput)
}

type EntryGroupIamMemberOutput struct{ *pulumi.OutputState }

func (EntryGroupIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntryGroupIamMember)(nil)).Elem()
}

func (o EntryGroupIamMemberOutput) ToEntryGroupIamMemberOutput() EntryGroupIamMemberOutput {
	return o
}

func (o EntryGroupIamMemberOutput) ToEntryGroupIamMemberOutputWithContext(ctx context.Context) EntryGroupIamMemberOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o EntryGroupIamMemberOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *EntryGroupIamMember) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o EntryGroupIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryGroupIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Identity that will be granted the privilege in role. The entry can have one of the following values:
//
//   - user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
//   - domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o EntryGroupIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryGroupIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The name of the resource to manage IAM policies for.
func (o EntryGroupIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryGroupIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o EntryGroupIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryGroupIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied.
func (o EntryGroupIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *EntryGroupIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EntryGroupIamMemberInput)(nil)).Elem(), &EntryGroupIamMember{})
	pulumi.RegisterOutputType(EntryGroupIamMemberOutput{})
}
