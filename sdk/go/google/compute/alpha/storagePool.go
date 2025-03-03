// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a storage pool in the specified project using the data in the request.
type StoragePool struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Type of the resource. Always compute#storagePool for storage pools.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A fingerprint for the labels being applied to this storage pool, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a storage pool.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this storage pool. These can be later modified by the setLabels method.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Provsioned IOPS of the storage pool.
	ProvisionedIops pulumi.StringOutput `pulumi:"provisionedIops"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// Status information for the storage pool resource.
	ResourceStatus StoragePoolResourceStatusResponseOutput `pulumi:"resourceStatus"`
	// Server-defined fully-qualified URL for this resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource's resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// Size, in GiB, of the storage pool.
	SizeGb pulumi.StringOutput `pulumi:"sizeGb"`
	// The status of storage pool creation. - CREATING: Storage pool is provisioning. storagePool. - FAILED: Storage pool creation failed. - READY: Storage pool is ready for use. - DELETING: Storage pool is deleting.
	State pulumi.StringOutput `pulumi:"state"`
	// Type of the storage pool
	Type pulumi.StringOutput `pulumi:"type"`
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewStoragePool registers a new resource with the given unique name, arguments, and options.
func NewStoragePool(ctx *pulumi.Context,
	name string, args *StoragePoolArgs, opts ...pulumi.ResourceOption) (*StoragePool, error) {
	if args == nil {
		args = &StoragePoolArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
		"zone",
	})
	opts = append(opts, replaceOnChanges)
	var resource StoragePool
	err := ctx.RegisterResource("google-native:compute/alpha:StoragePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStoragePool gets an existing StoragePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStoragePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StoragePoolState, opts ...pulumi.ResourceOption) (*StoragePool, error) {
	var resource StoragePool
	err := ctx.ReadResource("google-native:compute/alpha:StoragePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StoragePool resources.
type storagePoolState struct {
}

type StoragePoolState struct {
}

func (StoragePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*storagePoolState)(nil)).Elem()
}

type storagePoolArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Labels to apply to this storage pool. These can be later modified by the setLabels method.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Provsioned IOPS of the storage pool.
	ProvisionedIops *string `pulumi:"provisionedIops"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// Size, in GiB, of the storage pool.
	SizeGb *string `pulumi:"sizeGb"`
	// Type of the storage pool
	Type *StoragePoolType `pulumi:"type"`
	Zone *string          `pulumi:"zone"`
}

// The set of arguments for constructing a StoragePool resource.
type StoragePoolArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Labels to apply to this storage pool. These can be later modified by the setLabels method.
	Labels pulumi.StringMapInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Provsioned IOPS of the storage pool.
	ProvisionedIops pulumi.StringPtrInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// Size, in GiB, of the storage pool.
	SizeGb pulumi.StringPtrInput
	// Type of the storage pool
	Type StoragePoolTypePtrInput
	Zone pulumi.StringPtrInput
}

func (StoragePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storagePoolArgs)(nil)).Elem()
}

type StoragePoolInput interface {
	pulumi.Input

	ToStoragePoolOutput() StoragePoolOutput
	ToStoragePoolOutputWithContext(ctx context.Context) StoragePoolOutput
}

func (*StoragePool) ElementType() reflect.Type {
	return reflect.TypeOf((**StoragePool)(nil)).Elem()
}

func (i *StoragePool) ToStoragePoolOutput() StoragePoolOutput {
	return i.ToStoragePoolOutputWithContext(context.Background())
}

func (i *StoragePool) ToStoragePoolOutputWithContext(ctx context.Context) StoragePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoragePoolOutput)
}

type StoragePoolOutput struct{ *pulumi.OutputState }

func (StoragePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StoragePool)(nil)).Elem()
}

func (o StoragePoolOutput) ToStoragePoolOutput() StoragePoolOutput {
	return o
}

func (o StoragePoolOutput) ToStoragePoolOutputWithContext(ctx context.Context) StoragePoolOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o StoragePoolOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *StoragePool) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o StoragePoolOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *StoragePool) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Type of the resource. Always compute#storagePool for storage pools.
func (o StoragePoolOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *StoragePool) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// A fingerprint for the labels being applied to this storage pool, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a storage pool.
func (o StoragePoolOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *StoragePool) pulumi.StringOutput { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels to apply to this storage pool. These can be later modified by the setLabels method.
func (o StoragePoolOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StoragePool) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o StoragePoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StoragePool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StoragePoolOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *StoragePool) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Provsioned IOPS of the storage pool.
func (o StoragePoolOutput) ProvisionedIops() pulumi.StringOutput {
	return o.ApplyT(func(v *StoragePool) pulumi.StringOutput { return v.ProvisionedIops }).(pulumi.StringOutput)
}

// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
func (o StoragePoolOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StoragePool) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// Status information for the storage pool resource.
func (o StoragePoolOutput) ResourceStatus() StoragePoolResourceStatusResponseOutput {
	return o.ApplyT(func(v *StoragePool) StoragePoolResourceStatusResponseOutput { return v.ResourceStatus }).(StoragePoolResourceStatusResponseOutput)
}

// Server-defined fully-qualified URL for this resource.
func (o StoragePoolOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *StoragePool) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource's resource id.
func (o StoragePoolOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v *StoragePool) pulumi.StringOutput { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// Size, in GiB, of the storage pool.
func (o StoragePoolOutput) SizeGb() pulumi.StringOutput {
	return o.ApplyT(func(v *StoragePool) pulumi.StringOutput { return v.SizeGb }).(pulumi.StringOutput)
}

// The status of storage pool creation. - CREATING: Storage pool is provisioning. storagePool. - FAILED: Storage pool creation failed. - READY: Storage pool is ready for use. - DELETING: Storage pool is deleting.
func (o StoragePoolOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *StoragePool) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Type of the storage pool
func (o StoragePoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StoragePool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o StoragePoolOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *StoragePool) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StoragePoolInput)(nil)).Elem(), &StoragePool{})
	pulumi.RegisterOutputType(StoragePoolOutput{})
}
