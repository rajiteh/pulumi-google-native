// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Updates an IAM policy for the specified object.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type ObjectIamPolicy struct {
	pulumi.CustomResourceState

	// An association between a role, which comes with a set of permissions, and members who may assume that role.
	Bindings ObjectIamPolicyBindingsItemResponseArrayOutput `pulumi:"bindings"`
	Bucket   pulumi.StringOutput                            `pulumi:"bucket"`
	// HTTP 1.1  Entity tag for the policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// If present, selects a specific revision of this object (as opposed to the latest version, the default).
	Generation pulumi.StringPtrOutput `pulumi:"generation"`
	// The kind of item this is. For policies, this is always storage#policy. This field is ignored on input.
	Kind   pulumi.StringOutput `pulumi:"kind"`
	Object pulumi.StringOutput `pulumi:"object"`
	// The ID of the resource to which this policy belongs. Will be of the form projects/_/buckets/bucket for buckets, and projects/_/buckets/bucket/objects/object for objects. A specific generation may be specified by appending #generationNumber to the end of the object name, e.g. projects/_/buckets/my-bucket/objects/data.txt#17. The current generation can be denoted with #0. This field is ignored on input.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The project to be billed for this request. Required for Requester Pays buckets.
	UserProject pulumi.StringPtrOutput `pulumi:"userProject"`
	// The IAM policy format version.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewObjectIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewObjectIamPolicy(ctx *pulumi.Context,
	name string, args *ObjectIamPolicyArgs, opts ...pulumi.ResourceOption) (*ObjectIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Object == nil {
		return nil, errors.New("invalid value for required argument 'Object'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"bucket",
		"object",
	})
	opts = append(opts, replaceOnChanges)
	var resource ObjectIamPolicy
	err := ctx.RegisterResource("google-native:storage/v1:ObjectIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectIamPolicy gets an existing ObjectIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectIamPolicyState, opts ...pulumi.ResourceOption) (*ObjectIamPolicy, error) {
	var resource ObjectIamPolicy
	err := ctx.ReadResource("google-native:storage/v1:ObjectIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectIamPolicy resources.
type objectIamPolicyState struct {
}

type ObjectIamPolicyState struct {
}

func (ObjectIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectIamPolicyState)(nil)).Elem()
}

type objectIamPolicyArgs struct {
	// An association between a role, which comes with a set of permissions, and members who may assume that role.
	Bindings []ObjectIamPolicyBindingsItem `pulumi:"bindings"`
	Bucket   string                        `pulumi:"bucket"`
	// HTTP 1.1  Entity tag for the policy.
	Etag *string `pulumi:"etag"`
	// If present, selects a specific revision of this object (as opposed to the latest version, the default).
	Generation *string `pulumi:"generation"`
	// The kind of item this is. For policies, this is always storage#policy. This field is ignored on input.
	Kind   *string `pulumi:"kind"`
	Object string  `pulumi:"object"`
	// The ID of the resource to which this policy belongs. Will be of the form projects/_/buckets/bucket for buckets, and projects/_/buckets/bucket/objects/object for objects. A specific generation may be specified by appending #generationNumber to the end of the object name, e.g. projects/_/buckets/my-bucket/objects/data.txt#17. The current generation can be denoted with #0. This field is ignored on input.
	ResourceId *string `pulumi:"resourceId"`
	// The project to be billed for this request. Required for Requester Pays buckets.
	UserProject *string `pulumi:"userProject"`
	// The IAM policy format version.
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a ObjectIamPolicy resource.
type ObjectIamPolicyArgs struct {
	// An association between a role, which comes with a set of permissions, and members who may assume that role.
	Bindings ObjectIamPolicyBindingsItemArrayInput
	Bucket   pulumi.StringInput
	// HTTP 1.1  Entity tag for the policy.
	Etag pulumi.StringPtrInput
	// If present, selects a specific revision of this object (as opposed to the latest version, the default).
	Generation pulumi.StringPtrInput
	// The kind of item this is. For policies, this is always storage#policy. This field is ignored on input.
	Kind   pulumi.StringPtrInput
	Object pulumi.StringInput
	// The ID of the resource to which this policy belongs. Will be of the form projects/_/buckets/bucket for buckets, and projects/_/buckets/bucket/objects/object for objects. A specific generation may be specified by appending #generationNumber to the end of the object name, e.g. projects/_/buckets/my-bucket/objects/data.txt#17. The current generation can be denoted with #0. This field is ignored on input.
	ResourceId pulumi.StringPtrInput
	// The project to be billed for this request. Required for Requester Pays buckets.
	UserProject pulumi.StringPtrInput
	// The IAM policy format version.
	Version pulumi.IntPtrInput
}

func (ObjectIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectIamPolicyArgs)(nil)).Elem()
}

type ObjectIamPolicyInput interface {
	pulumi.Input

	ToObjectIamPolicyOutput() ObjectIamPolicyOutput
	ToObjectIamPolicyOutputWithContext(ctx context.Context) ObjectIamPolicyOutput
}

func (*ObjectIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectIamPolicy)(nil)).Elem()
}

func (i *ObjectIamPolicy) ToObjectIamPolicyOutput() ObjectIamPolicyOutput {
	return i.ToObjectIamPolicyOutputWithContext(context.Background())
}

func (i *ObjectIamPolicy) ToObjectIamPolicyOutputWithContext(ctx context.Context) ObjectIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectIamPolicyOutput)
}

type ObjectIamPolicyOutput struct{ *pulumi.OutputState }

func (ObjectIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectIamPolicy)(nil)).Elem()
}

func (o ObjectIamPolicyOutput) ToObjectIamPolicyOutput() ObjectIamPolicyOutput {
	return o
}

func (o ObjectIamPolicyOutput) ToObjectIamPolicyOutputWithContext(ctx context.Context) ObjectIamPolicyOutput {
	return o
}

// An association between a role, which comes with a set of permissions, and members who may assume that role.
func (o ObjectIamPolicyOutput) Bindings() ObjectIamPolicyBindingsItemResponseArrayOutput {
	return o.ApplyT(func(v *ObjectIamPolicy) ObjectIamPolicyBindingsItemResponseArrayOutput { return v.Bindings }).(ObjectIamPolicyBindingsItemResponseArrayOutput)
}

func (o ObjectIamPolicyOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectIamPolicy) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// HTTP 1.1  Entity tag for the policy.
func (o ObjectIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// If present, selects a specific revision of this object (as opposed to the latest version, the default).
func (o ObjectIamPolicyOutput) Generation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectIamPolicy) pulumi.StringPtrOutput { return v.Generation }).(pulumi.StringPtrOutput)
}

// The kind of item this is. For policies, this is always storage#policy. This field is ignored on input.
func (o ObjectIamPolicyOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectIamPolicy) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o ObjectIamPolicyOutput) Object() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectIamPolicy) pulumi.StringOutput { return v.Object }).(pulumi.StringOutput)
}

// The ID of the resource to which this policy belongs. Will be of the form projects/_/buckets/bucket for buckets, and projects/_/buckets/bucket/objects/object for objects. A specific generation may be specified by appending #generationNumber to the end of the object name, e.g. projects/_/buckets/my-bucket/objects/data.txt#17. The current generation can be denoted with #0. This field is ignored on input.
func (o ObjectIamPolicyOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectIamPolicy) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// The project to be billed for this request. Required for Requester Pays buckets.
func (o ObjectIamPolicyOutput) UserProject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ObjectIamPolicy) pulumi.StringPtrOutput { return v.UserProject }).(pulumi.StringPtrOutput)
}

// The IAM policy format version.
func (o ObjectIamPolicyOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *ObjectIamPolicy) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectIamPolicyInput)(nil)).Elem(), &ObjectIamPolicy{})
	pulumi.RegisterOutputType(ObjectIamPolicyOutput{})
}
