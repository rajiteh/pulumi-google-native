// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"errors"
	iam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the access control policy on the specified resource. Replaces any existing policy.
type InterconnectAttachmentIamMember struct {
	pulumi.CustomResourceState

	// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
	Condition iam.ConditionPtrOutput `pulumi:"condition"`
	// The etag of the resource's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding.
	Member pulumi.StringOutput `pulumi:"member"`
	// The name of the resource to manage IAM policies for.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project in which the resource belongs. If it is not provided, a default will be supplied.
	Project pulumi.StringOutput `pulumi:"project"`
	// Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewInterconnectAttachmentIamMember registers a new resource with the given unique name, arguments, and options.
func NewInterconnectAttachmentIamMember(ctx *pulumi.Context,
	name string, args *InterconnectAttachmentIamMemberArgs, opts ...pulumi.ResourceOption) (*InterconnectAttachmentIamMember, error) {
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
	var resource InterconnectAttachmentIamMember
	err := ctx.RegisterResource("google-native:compute/alpha:InterconnectAttachmentIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInterconnectAttachmentIamMember gets an existing InterconnectAttachmentIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInterconnectAttachmentIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InterconnectAttachmentIamMemberState, opts ...pulumi.ResourceOption) (*InterconnectAttachmentIamMember, error) {
	var resource InterconnectAttachmentIamMember
	err := ctx.ReadResource("google-native:compute/alpha:InterconnectAttachmentIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InterconnectAttachmentIamMember resources.
type interconnectAttachmentIamMemberState struct {
}

type InterconnectAttachmentIamMemberState struct {
}

func (InterconnectAttachmentIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*interconnectAttachmentIamMemberState)(nil)).Elem()
}

type interconnectAttachmentIamMemberArgs struct {
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

// The set of arguments for constructing a InterconnectAttachmentIamMember resource.
type InterconnectAttachmentIamMemberArgs struct {
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

func (InterconnectAttachmentIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*interconnectAttachmentIamMemberArgs)(nil)).Elem()
}

type InterconnectAttachmentIamMemberInput interface {
	pulumi.Input

	ToInterconnectAttachmentIamMemberOutput() InterconnectAttachmentIamMemberOutput
	ToInterconnectAttachmentIamMemberOutputWithContext(ctx context.Context) InterconnectAttachmentIamMemberOutput
}

func (*InterconnectAttachmentIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**InterconnectAttachmentIamMember)(nil)).Elem()
}

func (i *InterconnectAttachmentIamMember) ToInterconnectAttachmentIamMemberOutput() InterconnectAttachmentIamMemberOutput {
	return i.ToInterconnectAttachmentIamMemberOutputWithContext(context.Background())
}

func (i *InterconnectAttachmentIamMember) ToInterconnectAttachmentIamMemberOutputWithContext(ctx context.Context) InterconnectAttachmentIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InterconnectAttachmentIamMemberOutput)
}

type InterconnectAttachmentIamMemberOutput struct{ *pulumi.OutputState }

func (InterconnectAttachmentIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InterconnectAttachmentIamMember)(nil)).Elem()
}

func (o InterconnectAttachmentIamMemberOutput) ToInterconnectAttachmentIamMemberOutput() InterconnectAttachmentIamMemberOutput {
	return o
}

func (o InterconnectAttachmentIamMemberOutput) ToInterconnectAttachmentIamMemberOutputWithContext(ctx context.Context) InterconnectAttachmentIamMemberOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o InterconnectAttachmentIamMemberOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *InterconnectAttachmentIamMember) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o InterconnectAttachmentIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *InterconnectAttachmentIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding.
func (o InterconnectAttachmentIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *InterconnectAttachmentIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The name of the resource to manage IAM policies for.
func (o InterconnectAttachmentIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InterconnectAttachmentIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o InterconnectAttachmentIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *InterconnectAttachmentIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
func (o InterconnectAttachmentIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *InterconnectAttachmentIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InterconnectAttachmentIamMemberInput)(nil)).Elem(), &InterconnectAttachmentIamMember{})
	pulumi.RegisterOutputType(InterconnectAttachmentIamMemberOutput{})
}
