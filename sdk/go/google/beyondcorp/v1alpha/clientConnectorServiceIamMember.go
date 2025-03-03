// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"errors"
	iam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy. Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
type ClientConnectorServiceIamMember struct {
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

// NewClientConnectorServiceIamMember registers a new resource with the given unique name, arguments, and options.
func NewClientConnectorServiceIamMember(ctx *pulumi.Context,
	name string, args *ClientConnectorServiceIamMemberArgs, opts ...pulumi.ResourceOption) (*ClientConnectorServiceIamMember, error) {
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
	var resource ClientConnectorServiceIamMember
	err := ctx.RegisterResource("google-native:beyondcorp/v1alpha:ClientConnectorServiceIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientConnectorServiceIamMember gets an existing ClientConnectorServiceIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientConnectorServiceIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientConnectorServiceIamMemberState, opts ...pulumi.ResourceOption) (*ClientConnectorServiceIamMember, error) {
	var resource ClientConnectorServiceIamMember
	err := ctx.ReadResource("google-native:beyondcorp/v1alpha:ClientConnectorServiceIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientConnectorServiceIamMember resources.
type clientConnectorServiceIamMemberState struct {
}

type ClientConnectorServiceIamMemberState struct {
}

func (ClientConnectorServiceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientConnectorServiceIamMemberState)(nil)).Elem()
}

type clientConnectorServiceIamMemberArgs struct {
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

// The set of arguments for constructing a ClientConnectorServiceIamMember resource.
type ClientConnectorServiceIamMemberArgs struct {
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

func (ClientConnectorServiceIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientConnectorServiceIamMemberArgs)(nil)).Elem()
}

type ClientConnectorServiceIamMemberInput interface {
	pulumi.Input

	ToClientConnectorServiceIamMemberOutput() ClientConnectorServiceIamMemberOutput
	ToClientConnectorServiceIamMemberOutputWithContext(ctx context.Context) ClientConnectorServiceIamMemberOutput
}

func (*ClientConnectorServiceIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientConnectorServiceIamMember)(nil)).Elem()
}

func (i *ClientConnectorServiceIamMember) ToClientConnectorServiceIamMemberOutput() ClientConnectorServiceIamMemberOutput {
	return i.ToClientConnectorServiceIamMemberOutputWithContext(context.Background())
}

func (i *ClientConnectorServiceIamMember) ToClientConnectorServiceIamMemberOutputWithContext(ctx context.Context) ClientConnectorServiceIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientConnectorServiceIamMemberOutput)
}

type ClientConnectorServiceIamMemberOutput struct{ *pulumi.OutputState }

func (ClientConnectorServiceIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientConnectorServiceIamMember)(nil)).Elem()
}

func (o ClientConnectorServiceIamMemberOutput) ToClientConnectorServiceIamMemberOutput() ClientConnectorServiceIamMemberOutput {
	return o
}

func (o ClientConnectorServiceIamMemberOutput) ToClientConnectorServiceIamMemberOutputWithContext(ctx context.Context) ClientConnectorServiceIamMemberOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o ClientConnectorServiceIamMemberOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *ClientConnectorServiceIamMember) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o ClientConnectorServiceIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientConnectorServiceIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Identity that will be granted the privilege in role. The entry can have one of the following values:
//
//   - user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
//   - serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
//   - group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
//   - domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
func (o ClientConnectorServiceIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientConnectorServiceIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The name of the resource to manage IAM policies for.
func (o ClientConnectorServiceIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientConnectorServiceIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o ClientConnectorServiceIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientConnectorServiceIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role that should be applied.
func (o ClientConnectorServiceIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ClientConnectorServiceIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientConnectorServiceIamMemberInput)(nil)).Elem(), &ClientConnectorServiceIamMember{})
	pulumi.RegisterOutputType(ClientConnectorServiceIamMemberOutput{})
}
