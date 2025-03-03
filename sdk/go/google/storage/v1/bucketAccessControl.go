// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new ACL entry on the specified bucket.
// Auto-naming is currently not supported for this resource.
type BucketAccessControl struct {
	pulumi.CustomResourceState

	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The domain associated with the entity, if any.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The email address associated with the entity, if any.
	Email pulumi.StringOutput `pulumi:"email"`
	// The entity holding the permission, in one of the following forms:
	// - user-userId
	// - user-email
	// - group-groupId
	// - group-email
	// - domain-domain
	// - project-team-projectId
	// - allUsers
	// - allAuthenticatedUsers Examples:
	// - The user liz@example.com would be user-liz@example.com.
	// - The group example@googlegroups.com would be group-example@googlegroups.com.
	// - To refer to all members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
	Entity pulumi.StringOutput `pulumi:"entity"`
	// The ID for the entity, if any.
	EntityId pulumi.StringOutput `pulumi:"entityId"`
	// HTTP 1.1 Entity tag for the access-control entry.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The kind of item this is. For bucket access control entries, this is always storage#bucketAccessControl.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The project team associated with the entity, if any.
	ProjectTeam BucketAccessControlProjectTeamResponseOutput `pulumi:"projectTeam"`
	// The access permission for the entity.
	Role pulumi.StringOutput `pulumi:"role"`
	// The link to this access-control entry.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The project to be billed for this request. Required for Requester Pays buckets.
	UserProject pulumi.StringPtrOutput `pulumi:"userProject"`
}

// NewBucketAccessControl registers a new resource with the given unique name, arguments, and options.
func NewBucketAccessControl(ctx *pulumi.Context,
	name string, args *BucketAccessControlArgs, opts ...pulumi.ResourceOption) (*BucketAccessControl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"bucket",
	})
	opts = append(opts, replaceOnChanges)
	var resource BucketAccessControl
	err := ctx.RegisterResource("google-native:storage/v1:BucketAccessControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketAccessControl gets an existing BucketAccessControl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketAccessControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketAccessControlState, opts ...pulumi.ResourceOption) (*BucketAccessControl, error) {
	var resource BucketAccessControl
	err := ctx.ReadResource("google-native:storage/v1:BucketAccessControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketAccessControl resources.
type bucketAccessControlState struct {
}

type BucketAccessControlState struct {
}

func (BucketAccessControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketAccessControlState)(nil)).Elem()
}

type bucketAccessControlArgs struct {
	// The name of the bucket.
	Bucket string `pulumi:"bucket"`
	// The domain associated with the entity, if any.
	Domain *string `pulumi:"domain"`
	// The email address associated with the entity, if any.
	Email *string `pulumi:"email"`
	// The entity holding the permission, in one of the following forms:
	// - user-userId
	// - user-email
	// - group-groupId
	// - group-email
	// - domain-domain
	// - project-team-projectId
	// - allUsers
	// - allAuthenticatedUsers Examples:
	// - The user liz@example.com would be user-liz@example.com.
	// - The group example@googlegroups.com would be group-example@googlegroups.com.
	// - To refer to all members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
	Entity *string `pulumi:"entity"`
	// The ID for the entity, if any.
	EntityId *string `pulumi:"entityId"`
	// HTTP 1.1 Entity tag for the access-control entry.
	Etag *string `pulumi:"etag"`
	// The ID of the access-control entry.
	Id *string `pulumi:"id"`
	// The kind of item this is. For bucket access control entries, this is always storage#bucketAccessControl.
	Kind *string `pulumi:"kind"`
	// The project team associated with the entity, if any.
	ProjectTeam *BucketAccessControlProjectTeam `pulumi:"projectTeam"`
	// The access permission for the entity.
	Role *string `pulumi:"role"`
	// The link to this access-control entry.
	SelfLink *string `pulumi:"selfLink"`
	// The project to be billed for this request. Required for Requester Pays buckets.
	UserProject *string `pulumi:"userProject"`
}

// The set of arguments for constructing a BucketAccessControl resource.
type BucketAccessControlArgs struct {
	// The name of the bucket.
	Bucket pulumi.StringInput
	// The domain associated with the entity, if any.
	Domain pulumi.StringPtrInput
	// The email address associated with the entity, if any.
	Email pulumi.StringPtrInput
	// The entity holding the permission, in one of the following forms:
	// - user-userId
	// - user-email
	// - group-groupId
	// - group-email
	// - domain-domain
	// - project-team-projectId
	// - allUsers
	// - allAuthenticatedUsers Examples:
	// - The user liz@example.com would be user-liz@example.com.
	// - The group example@googlegroups.com would be group-example@googlegroups.com.
	// - To refer to all members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
	Entity pulumi.StringPtrInput
	// The ID for the entity, if any.
	EntityId pulumi.StringPtrInput
	// HTTP 1.1 Entity tag for the access-control entry.
	Etag pulumi.StringPtrInput
	// The ID of the access-control entry.
	Id pulumi.StringPtrInput
	// The kind of item this is. For bucket access control entries, this is always storage#bucketAccessControl.
	Kind pulumi.StringPtrInput
	// The project team associated with the entity, if any.
	ProjectTeam BucketAccessControlProjectTeamPtrInput
	// The access permission for the entity.
	Role pulumi.StringPtrInput
	// The link to this access-control entry.
	SelfLink pulumi.StringPtrInput
	// The project to be billed for this request. Required for Requester Pays buckets.
	UserProject pulumi.StringPtrInput
}

func (BucketAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketAccessControlArgs)(nil)).Elem()
}

type BucketAccessControlInput interface {
	pulumi.Input

	ToBucketAccessControlOutput() BucketAccessControlOutput
	ToBucketAccessControlOutputWithContext(ctx context.Context) BucketAccessControlOutput
}

func (*BucketAccessControl) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketAccessControl)(nil)).Elem()
}

func (i *BucketAccessControl) ToBucketAccessControlOutput() BucketAccessControlOutput {
	return i.ToBucketAccessControlOutputWithContext(context.Background())
}

func (i *BucketAccessControl) ToBucketAccessControlOutputWithContext(ctx context.Context) BucketAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketAccessControlOutput)
}

type BucketAccessControlOutput struct{ *pulumi.OutputState }

func (BucketAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketAccessControl)(nil)).Elem()
}

func (o BucketAccessControlOutput) ToBucketAccessControlOutput() BucketAccessControlOutput {
	return o
}

func (o BucketAccessControlOutput) ToBucketAccessControlOutputWithContext(ctx context.Context) BucketAccessControlOutput {
	return o
}

func (o BucketAccessControlOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketAccessControl) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The domain associated with the entity, if any.
func (o BucketAccessControlOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketAccessControl) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// The email address associated with the entity, if any.
func (o BucketAccessControlOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketAccessControl) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// The entity holding the permission, in one of the following forms:
// - user-userId
// - user-email
// - group-groupId
// - group-email
// - domain-domain
// - project-team-projectId
// - allUsers
// - allAuthenticatedUsers Examples:
// - The user liz@example.com would be user-liz@example.com.
// - The group example@googlegroups.com would be group-example@googlegroups.com.
// - To refer to all members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
func (o BucketAccessControlOutput) Entity() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketAccessControl) pulumi.StringOutput { return v.Entity }).(pulumi.StringOutput)
}

// The ID for the entity, if any.
func (o BucketAccessControlOutput) EntityId() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketAccessControl) pulumi.StringOutput { return v.EntityId }).(pulumi.StringOutput)
}

// HTTP 1.1 Entity tag for the access-control entry.
func (o BucketAccessControlOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketAccessControl) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The kind of item this is. For bucket access control entries, this is always storage#bucketAccessControl.
func (o BucketAccessControlOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketAccessControl) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The project team associated with the entity, if any.
func (o BucketAccessControlOutput) ProjectTeam() BucketAccessControlProjectTeamResponseOutput {
	return o.ApplyT(func(v *BucketAccessControl) BucketAccessControlProjectTeamResponseOutput { return v.ProjectTeam }).(BucketAccessControlProjectTeamResponseOutput)
}

// The access permission for the entity.
func (o BucketAccessControlOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketAccessControl) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// The link to this access-control entry.
func (o BucketAccessControlOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *BucketAccessControl) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The project to be billed for this request. Required for Requester Pays buckets.
func (o BucketAccessControlOutput) UserProject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BucketAccessControl) pulumi.StringPtrOutput { return v.UserProject }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketAccessControlInput)(nil)).Elem(), &BucketAccessControl{})
	pulumi.RegisterOutputType(BucketAccessControlOutput{})
}
