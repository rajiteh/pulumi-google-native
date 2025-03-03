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

// Sets the IAM policy.
type DataExchangeIamMember struct {
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

// NewDataExchangeIamMember registers a new resource with the given unique name, arguments, and options.
func NewDataExchangeIamMember(ctx *pulumi.Context,
	name string, args *DataExchangeIamMemberArgs, opts ...pulumi.ResourceOption) (*DataExchangeIamMember, error) {
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
	var resource DataExchangeIamMember
	err := ctx.RegisterResource("google-native:analyticshub/v1:DataExchangeIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataExchangeIamMember gets an existing DataExchangeIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataExchangeIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataExchangeIamMemberState, opts ...pulumi.ResourceOption) (*DataExchangeIamMember, error) {
	var resource DataExchangeIamMember
	err := ctx.ReadResource("google-native:analyticshub/v1:DataExchangeIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataExchangeIamMember resources.
type dataExchangeIamMemberState struct {
}

type DataExchangeIamMemberState struct {
}

func (DataExchangeIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataExchangeIamMemberState)(nil)).Elem()
}

type dataExchangeIamMemberArgs struct {
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

// The set of arguments for constructing a DataExchangeIamMember resource.
type DataExchangeIamMemberArgs struct {
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

func (DataExchangeIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataExchangeIamMemberArgs)(nil)).Elem()
}

type DataExchangeIamMemberInput interface {
	pulumi.Input

	ToDataExchangeIamMemberOutput() DataExchangeIamMemberOutput
	ToDataExchangeIamMemberOutputWithContext(ctx context.Context) DataExchangeIamMemberOutput
}

func (*DataExchangeIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**DataExchangeIamMember)(nil)).Elem()
}

func (i *DataExchangeIamMember) ToDataExchangeIamMemberOutput() DataExchangeIamMemberOutput {
	return i.ToDataExchangeIamMemberOutputWithContext(context.Background())
}

func (i *DataExchangeIamMember) ToDataExchangeIamMemberOutputWithContext(ctx context.Context) DataExchangeIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataExchangeIamMemberOutput)
}

type DataExchangeIamMemberOutput struct{ *pulumi.OutputState }

func (DataExchangeIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataExchangeIamMember)(nil)).Elem()
}

func (o DataExchangeIamMemberOutput) ToDataExchangeIamMemberOutput() DataExchangeIamMemberOutput {
	return o
}

func (o DataExchangeIamMemberOutput) ToDataExchangeIamMemberOutputWithContext(ctx context.Context) DataExchangeIamMemberOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o DataExchangeIamMemberOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *DataExchangeIamMember) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o DataExchangeIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchangeIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding.
func (o DataExchangeIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchangeIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The name of the resource to manage IAM policies for.
func (o DataExchangeIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchangeIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o DataExchangeIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchangeIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
func (o DataExchangeIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchangeIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataExchangeIamMemberInput)(nil)).Elem(), &DataExchangeIamMember{})
	pulumi.RegisterOutputType(DataExchangeIamMemberOutput{})
}
