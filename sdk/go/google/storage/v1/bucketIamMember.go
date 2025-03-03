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

// Updates an IAM policy for the specified bucket.
type BucketIamMember struct {
	pulumi.CustomResourceState

	// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
	Condition iam.ConditionPtrOutput `pulumi:"condition"`
	// The etag of the resource's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A collection of identifiers for members who may assume the provided role. Recognized identifiers are as follows:
	// - allUsers — A special identifier that represents anyone on the internet; with or without a Google account.
	// - allAuthenticatedUsers — A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// - user:emailid — An email address that represents a specific account. For example, user:alice@gmail.com or user:joe@example.com.
	// - serviceAccount:emailid — An email address that represents a service account. For example,  serviceAccount:my-other-app@appspot.gserviceaccount.com .
	// - group:emailid — An email address that represents a Google group. For example, group:admins@example.com.
	// - domain:domain — A Google Apps domain name that represents all the users of that domain. For example, domain:google.com or domain:example.com.
	// - projectOwner:projectid — Owners of the given project. For example, projectOwner:my-example-project
	// - projectEditor:projectid — Editors of the given project. For example, projectEditor:my-example-project
	// - projectViewer:projectid — Viewers of the given project. For example, projectViewer:my-example-project
	Member pulumi.StringOutput `pulumi:"member"`
	// The name of the resource to manage IAM policies for.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project in which the resource belongs. If it is not provided, a default will be supplied.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role to which members belong. Two types of roles are supported: new IAM roles, which grant permissions that do not map directly to those provided by ACLs, and legacy IAM roles, which do map directly to ACL permissions. All roles are of the format roles/storage.specificRole.
	// The new IAM roles are:
	// - roles/storage.admin — Full control of Google Cloud Storage resources.
	// - roles/storage.objectViewer — Read-Only access to Google Cloud Storage objects.
	// - roles/storage.objectCreator — Access to create objects in Google Cloud Storage.
	// - roles/storage.objectAdmin — Full control of Google Cloud Storage objects.   The legacy IAM roles are:
	// - roles/storage.legacyObjectReader — Read-only access to objects without listing. Equivalent to an ACL entry on an object with the READER role.
	// - roles/storage.legacyObjectOwner — Read/write access to existing objects without listing. Equivalent to an ACL entry on an object with the OWNER role.
	// - roles/storage.legacyBucketReader — Read access to buckets with object listing. Equivalent to an ACL entry on a bucket with the READER role.
	// - roles/storage.legacyBucketWriter — Read access to buckets with object listing/creation/deletion. Equivalent to an ACL entry on a bucket with the WRITER role.
	// - roles/storage.legacyBucketOwner — Read and write access to existing buckets with object listing/creation/deletion. Equivalent to an ACL entry on a bucket with the OWNER role.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewBucketIamMember registers a new resource with the given unique name, arguments, and options.
func NewBucketIamMember(ctx *pulumi.Context,
	name string, args *BucketIamMemberArgs, opts ...pulumi.ResourceOption) (*BucketIamMember, error) {
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
	var resource BucketIamMember
	err := ctx.RegisterResource("google-native:storage/v1:BucketIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketIamMember gets an existing BucketIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketIamMemberState, opts ...pulumi.ResourceOption) (*BucketIamMember, error) {
	var resource BucketIamMember
	err := ctx.ReadResource("google-native:storage/v1:BucketIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketIamMember resources.
type bucketIamMemberState struct {
}

type BucketIamMemberState struct {
}

func (BucketIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketIamMemberState)(nil)).Elem()
}

type bucketIamMemberArgs struct {
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

// The set of arguments for constructing a BucketIamMember resource.
type BucketIamMemberArgs struct {
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

func (BucketIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketIamMemberArgs)(nil)).Elem()
}

type BucketIamMemberInput interface {
	pulumi.Input

	ToBucketIamMemberOutput() BucketIamMemberOutput
	ToBucketIamMemberOutputWithContext(ctx context.Context) BucketIamMemberOutput
}

func (*BucketIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketIamMember)(nil)).Elem()
}

func (i *BucketIamMember) ToBucketIamMemberOutput() BucketIamMemberOutput {
	return i.ToBucketIamMemberOutputWithContext(context.Background())
}

func (i *BucketIamMember) ToBucketIamMemberOutputWithContext(ctx context.Context) BucketIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketIamMemberOutput)
}

type BucketIamMemberOutput struct{ *pulumi.OutputState }

func (BucketIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketIamMember)(nil)).Elem()
}

func (o BucketIamMemberOutput) ToBucketIamMemberOutput() BucketIamMemberOutput {
	return o
}

func (o BucketIamMemberOutput) ToBucketIamMemberOutputWithContext(ctx context.Context) BucketIamMemberOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o BucketIamMemberOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *BucketIamMember) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o BucketIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// A collection of identifiers for members who may assume the provided role. Recognized identifiers are as follows:
// - allUsers — A special identifier that represents anyone on the internet; with or without a Google account.
// - allAuthenticatedUsers — A special identifier that represents anyone who is authenticated with a Google account or a service account.
// - user:emailid — An email address that represents a specific account. For example, user:alice@gmail.com or user:joe@example.com.
// - serviceAccount:emailid — An email address that represents a service account. For example,  serviceAccount:my-other-app@appspot.gserviceaccount.com .
// - group:emailid — An email address that represents a Google group. For example, group:admins@example.com.
// - domain:domain — A Google Apps domain name that represents all the users of that domain. For example, domain:google.com or domain:example.com.
// - projectOwner:projectid — Owners of the given project. For example, projectOwner:my-example-project
// - projectEditor:projectid — Editors of the given project. For example, projectEditor:my-example-project
// - projectViewer:projectid — Viewers of the given project. For example, projectViewer:my-example-project
func (o BucketIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The name of the resource to manage IAM policies for.
func (o BucketIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o BucketIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The role to which members belong. Two types of roles are supported: new IAM roles, which grant permissions that do not map directly to those provided by ACLs, and legacy IAM roles, which do map directly to ACL permissions. All roles are of the format roles/storage.specificRole.
// The new IAM roles are:
// - roles/storage.admin — Full control of Google Cloud Storage resources.
// - roles/storage.objectViewer — Read-Only access to Google Cloud Storage objects.
// - roles/storage.objectCreator — Access to create objects in Google Cloud Storage.
// - roles/storage.objectAdmin — Full control of Google Cloud Storage objects.   The legacy IAM roles are:
// - roles/storage.legacyObjectReader — Read-only access to objects without listing. Equivalent to an ACL entry on an object with the READER role.
// - roles/storage.legacyObjectOwner — Read/write access to existing objects without listing. Equivalent to an ACL entry on an object with the OWNER role.
// - roles/storage.legacyBucketReader — Read access to buckets with object listing. Equivalent to an ACL entry on a bucket with the READER role.
// - roles/storage.legacyBucketWriter — Read access to buckets with object listing/creation/deletion. Equivalent to an ACL entry on a bucket with the WRITER role.
// - roles/storage.legacyBucketOwner — Read and write access to existing buckets with object listing/creation/deletion. Equivalent to an ACL entry on a bucket with the OWNER role.
func (o BucketIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketIamMemberInput)(nil)).Elem(), &BucketIamMember{})
	pulumi.RegisterOutputType(BucketIamMemberOutput{})
}
