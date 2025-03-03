// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a snapshot in the specified project using the data included in the request. For regular snapshot creation, consider using this method instead of disks.createSnapshot, as this method supports more features, such as creating snapshots in a project different from the source disk project.
type Snapshot struct {
	pulumi.CustomResourceState

	// The architecture of the snapshot. Valid values are ARM64 or X86_64.
	Architecture pulumi.StringOutput `pulumi:"architecture"`
	// Set to true if snapshots are automatically created by applying resource policy on the target disk.
	AutoCreated pulumi.BoolOutput `pulumi:"autoCreated"`
	// Creates the new snapshot in the snapshot chain labeled with the specified name. The chain name must be 1-63 characters long and comply with RFC1035. This is an uncommon option only for advanced service owners who needs to create separate snapshot chains, for example, for chargeback tracking. When you describe your snapshot resource, this field is visible only if it has a non-empty value.
	ChainName pulumi.StringOutput `pulumi:"chainName"`
	// Size in bytes of the snapshot at creation time.
	CreationSizeBytes pulumi.StringOutput `pulumi:"creationSizeBytes"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Size of the source disk, specified in GB.
	DiskSizeGb pulumi.StringOutput `pulumi:"diskSizeGb"`
	// Number of bytes downloaded to restore a snapshot to a disk.
	DownloadBytes pulumi.StringOutput `pulumi:"downloadBytes"`
	// Whether this snapshot is created from a confidential compute mode disk. see go/confidential-mode-in-arcus for details. [Output Only]: This field is not set by user, but from source disk.
	EnableConfidentialCompute pulumi.BoolOutput `pulumi:"enableConfidentialCompute"`
	// [Input Only] Whether to attempt an application consistent snapshot by informing the OS to prepare for the snapshot process.
	GuestFlush pulumi.BoolOutput `pulumi:"guestFlush"`
	// A list of features to enable on the guest operating system. Applicable only for bootable images. Read Enabling guest operating system features to see a list of available options.
	GuestOsFeatures GuestOsFeatureResponseArrayOutput `pulumi:"guestOsFeatures"`
	// Type of the resource. Always compute#snapshot for Snapshot resources.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A fingerprint for the labels being applied to this snapshot, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a snapshot.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this snapshot. These can be later modified by the setLabels method. Label values may be empty.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Integer license codes indicating which licenses are attached to this snapshot.
	LicenseCodes pulumi.StringArrayOutput `pulumi:"licenseCodes"`
	// A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses attached (such as a Windows image).
	Licenses pulumi.StringArrayOutput `pulumi:"licenses"`
	// An opaque location hint used to place the snapshot close to other resources. This field is for use by internal tools that use the public API.
	LocationHint pulumi.StringOutput `pulumi:"locationHint"`
	// Number of days the snapshot should be retained before being deleted automatically.
	MaxRetentionDays pulumi.IntOutput `pulumi:"maxRetentionDays"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// Reserved for future use.
	SatisfiesPzs pulumi.BoolOutput `pulumi:"satisfiesPzs"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource's resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// Encrypts the snapshot using a customer-supplied encryption key. After you encrypt a snapshot using a customer-supplied key, you must provide the same key if you use the snapshot later. For example, you must provide the encryption key when you create a disk from the encrypted snapshot in a future request. Customer-supplied encryption keys do not protect access to metadata of the snapshot. If you do not provide an encryption key when creating the snapshot, then the snapshot will be encrypted using an automatically generated key and you do not need to provide a key to use the snapshot later.
	SnapshotEncryptionKey CustomerEncryptionKeyResponseOutput `pulumi:"snapshotEncryptionKey"`
	// Indicates the type of the snapshot.
	SnapshotType pulumi.StringOutput `pulumi:"snapshotType"`
	// The source disk used to create this snapshot.
	SourceDisk pulumi.StringOutput `pulumi:"sourceDisk"`
	// The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
	SourceDiskEncryptionKey CustomerEncryptionKeyResponseOutput `pulumi:"sourceDiskEncryptionKey"`
	// The source disk whose recovery checkpoint will be used to create this snapshot.
	SourceDiskForRecoveryCheckpoint pulumi.StringOutput `pulumi:"sourceDiskForRecoveryCheckpoint"`
	// The ID value of the disk used to create this snapshot. This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given disk name.
	SourceDiskId pulumi.StringOutput `pulumi:"sourceDiskId"`
	// The source instant snapshot used to create this snapshot. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instantSnapshots/instantSnapshot - projects/project/zones/zone/instantSnapshots/instantSnapshot - zones/zone/instantSnapshots/instantSnapshot
	SourceInstantSnapshot pulumi.StringOutput `pulumi:"sourceInstantSnapshot"`
	// The unique ID of the instant snapshot used to create this snapshot. This value identifies the exact instant snapshot that was used to create this persistent disk. For example, if you created the persistent disk from an instant snapshot that was later deleted and recreated under the same name, the source instant snapshot ID would identify the exact instant snapshot that was used.
	SourceInstantSnapshotId pulumi.StringOutput `pulumi:"sourceInstantSnapshotId"`
	// URL of the resource policy which created this scheduled snapshot.
	SourceSnapshotSchedulePolicy pulumi.StringOutput `pulumi:"sourceSnapshotSchedulePolicy"`
	// ID of the resource policy which created this scheduled snapshot.
	SourceSnapshotSchedulePolicyId pulumi.StringOutput `pulumi:"sourceSnapshotSchedulePolicyId"`
	// The status of the snapshot. This can be CREATING, DELETING, FAILED, READY, or UPLOADING.
	Status pulumi.StringOutput `pulumi:"status"`
	// A size of the storage used by the snapshot. As snapshots share storage, this number is expected to change with snapshot creation/deletion.
	StorageBytes pulumi.StringOutput `pulumi:"storageBytes"`
	// An indicator whether storageBytes is in a stable state or it is being adjusted as a result of shared storage reallocation. This status can either be UPDATING, meaning the size of the snapshot is being updated, or UP_TO_DATE, meaning the size of the snapshot is up-to-date.
	StorageBytesStatus pulumi.StringOutput `pulumi:"storageBytesStatus"`
	// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
	StorageLocations pulumi.StringArrayOutput `pulumi:"storageLocations"`
	// A list of user provided licenses represented by a list of URLs to the license resource.
	UserLicenses pulumi.StringArrayOutput `pulumi:"userLicenses"`
}

// NewSnapshot registers a new resource with the given unique name, arguments, and options.
func NewSnapshot(ctx *pulumi.Context,
	name string, args *SnapshotArgs, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	if args == nil {
		args = &SnapshotArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource Snapshot
	err := ctx.RegisterResource("google-native:compute/alpha:Snapshot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnapshot gets an existing Snapshot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnapshot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnapshotState, opts ...pulumi.ResourceOption) (*Snapshot, error) {
	var resource Snapshot
	err := ctx.ReadResource("google-native:compute/alpha:Snapshot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snapshot resources.
type snapshotState struct {
}

type SnapshotState struct {
}

func (SnapshotState) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotState)(nil)).Elem()
}

type snapshotArgs struct {
	// Creates the new snapshot in the snapshot chain labeled with the specified name. The chain name must be 1-63 characters long and comply with RFC1035. This is an uncommon option only for advanced service owners who needs to create separate snapshot chains, for example, for chargeback tracking. When you describe your snapshot resource, this field is visible only if it has a non-empty value.
	ChainName *string `pulumi:"chainName"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Whether this snapshot is created from a confidential compute mode disk. see go/confidential-mode-in-arcus for details. [Output Only]: This field is not set by user, but from source disk.
	EnableConfidentialCompute *bool `pulumi:"enableConfidentialCompute"`
	// [Input Only] Whether to attempt an application consistent snapshot by informing the OS to prepare for the snapshot process.
	GuestFlush *bool `pulumi:"guestFlush"`
	// Labels to apply to this snapshot. These can be later modified by the setLabels method. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// An opaque location hint used to place the snapshot close to other resources. This field is for use by internal tools that use the public API.
	LocationHint *string `pulumi:"locationHint"`
	// Number of days the snapshot should be retained before being deleted automatically.
	MaxRetentionDays *int `pulumi:"maxRetentionDays"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// Encrypts the snapshot using a customer-supplied encryption key. After you encrypt a snapshot using a customer-supplied key, you must provide the same key if you use the snapshot later. For example, you must provide the encryption key when you create a disk from the encrypted snapshot in a future request. Customer-supplied encryption keys do not protect access to metadata of the snapshot. If you do not provide an encryption key when creating the snapshot, then the snapshot will be encrypted using an automatically generated key and you do not need to provide a key to use the snapshot later.
	SnapshotEncryptionKey *CustomerEncryptionKey `pulumi:"snapshotEncryptionKey"`
	// Indicates the type of the snapshot.
	SnapshotType *SnapshotSnapshotType `pulumi:"snapshotType"`
	// The source disk used to create this snapshot.
	SourceDisk *string `pulumi:"sourceDisk"`
	// The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
	SourceDiskEncryptionKey *CustomerEncryptionKey `pulumi:"sourceDiskEncryptionKey"`
	// The source disk whose recovery checkpoint will be used to create this snapshot.
	SourceDiskForRecoveryCheckpoint *string `pulumi:"sourceDiskForRecoveryCheckpoint"`
	// The source instant snapshot used to create this snapshot. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instantSnapshots/instantSnapshot - projects/project/zones/zone/instantSnapshots/instantSnapshot - zones/zone/instantSnapshots/instantSnapshot
	SourceInstantSnapshot *string `pulumi:"sourceInstantSnapshot"`
	// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
	StorageLocations []string `pulumi:"storageLocations"`
}

// The set of arguments for constructing a Snapshot resource.
type SnapshotArgs struct {
	// Creates the new snapshot in the snapshot chain labeled with the specified name. The chain name must be 1-63 characters long and comply with RFC1035. This is an uncommon option only for advanced service owners who needs to create separate snapshot chains, for example, for chargeback tracking. When you describe your snapshot resource, this field is visible only if it has a non-empty value.
	ChainName pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Whether this snapshot is created from a confidential compute mode disk. see go/confidential-mode-in-arcus for details. [Output Only]: This field is not set by user, but from source disk.
	EnableConfidentialCompute pulumi.BoolPtrInput
	// [Input Only] Whether to attempt an application consistent snapshot by informing the OS to prepare for the snapshot process.
	GuestFlush pulumi.BoolPtrInput
	// Labels to apply to this snapshot. These can be later modified by the setLabels method. Label values may be empty.
	Labels pulumi.StringMapInput
	// An opaque location hint used to place the snapshot close to other resources. This field is for use by internal tools that use the public API.
	LocationHint pulumi.StringPtrInput
	// Number of days the snapshot should be retained before being deleted automatically.
	MaxRetentionDays pulumi.IntPtrInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// Encrypts the snapshot using a customer-supplied encryption key. After you encrypt a snapshot using a customer-supplied key, you must provide the same key if you use the snapshot later. For example, you must provide the encryption key when you create a disk from the encrypted snapshot in a future request. Customer-supplied encryption keys do not protect access to metadata of the snapshot. If you do not provide an encryption key when creating the snapshot, then the snapshot will be encrypted using an automatically generated key and you do not need to provide a key to use the snapshot later.
	SnapshotEncryptionKey CustomerEncryptionKeyPtrInput
	// Indicates the type of the snapshot.
	SnapshotType SnapshotSnapshotTypePtrInput
	// The source disk used to create this snapshot.
	SourceDisk pulumi.StringPtrInput
	// The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
	SourceDiskEncryptionKey CustomerEncryptionKeyPtrInput
	// The source disk whose recovery checkpoint will be used to create this snapshot.
	SourceDiskForRecoveryCheckpoint pulumi.StringPtrInput
	// The source instant snapshot used to create this snapshot. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instantSnapshots/instantSnapshot - projects/project/zones/zone/instantSnapshots/instantSnapshot - zones/zone/instantSnapshots/instantSnapshot
	SourceInstantSnapshot pulumi.StringPtrInput
	// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
	StorageLocations pulumi.StringArrayInput
}

func (SnapshotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snapshotArgs)(nil)).Elem()
}

type SnapshotInput interface {
	pulumi.Input

	ToSnapshotOutput() SnapshotOutput
	ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput
}

func (*Snapshot) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (i *Snapshot) ToSnapshotOutput() SnapshotOutput {
	return i.ToSnapshotOutputWithContext(context.Background())
}

func (i *Snapshot) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotOutput)
}

type SnapshotOutput struct{ *pulumi.OutputState }

func (SnapshotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snapshot)(nil)).Elem()
}

func (o SnapshotOutput) ToSnapshotOutput() SnapshotOutput {
	return o
}

func (o SnapshotOutput) ToSnapshotOutputWithContext(ctx context.Context) SnapshotOutput {
	return o
}

// The architecture of the snapshot. Valid values are ARM64 or X86_64.
func (o SnapshotOutput) Architecture() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Architecture }).(pulumi.StringOutput)
}

// Set to true if snapshots are automatically created by applying resource policy on the target disk.
func (o SnapshotOutput) AutoCreated() pulumi.BoolOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.BoolOutput { return v.AutoCreated }).(pulumi.BoolOutput)
}

// Creates the new snapshot in the snapshot chain labeled with the specified name. The chain name must be 1-63 characters long and comply with RFC1035. This is an uncommon option only for advanced service owners who needs to create separate snapshot chains, for example, for chargeback tracking. When you describe your snapshot resource, this field is visible only if it has a non-empty value.
func (o SnapshotOutput) ChainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.ChainName }).(pulumi.StringOutput)
}

// Size in bytes of the snapshot at creation time.
func (o SnapshotOutput) CreationSizeBytes() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.CreationSizeBytes }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o SnapshotOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o SnapshotOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Size of the source disk, specified in GB.
func (o SnapshotOutput) DiskSizeGb() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.DiskSizeGb }).(pulumi.StringOutput)
}

// Number of bytes downloaded to restore a snapshot to a disk.
func (o SnapshotOutput) DownloadBytes() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.DownloadBytes }).(pulumi.StringOutput)
}

// Whether this snapshot is created from a confidential compute mode disk. see go/confidential-mode-in-arcus for details. [Output Only]: This field is not set by user, but from source disk.
func (o SnapshotOutput) EnableConfidentialCompute() pulumi.BoolOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.BoolOutput { return v.EnableConfidentialCompute }).(pulumi.BoolOutput)
}

// [Input Only] Whether to attempt an application consistent snapshot by informing the OS to prepare for the snapshot process.
func (o SnapshotOutput) GuestFlush() pulumi.BoolOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.BoolOutput { return v.GuestFlush }).(pulumi.BoolOutput)
}

// A list of features to enable on the guest operating system. Applicable only for bootable images. Read Enabling guest operating system features to see a list of available options.
func (o SnapshotOutput) GuestOsFeatures() GuestOsFeatureResponseArrayOutput {
	return o.ApplyT(func(v *Snapshot) GuestOsFeatureResponseArrayOutput { return v.GuestOsFeatures }).(GuestOsFeatureResponseArrayOutput)
}

// Type of the resource. Always compute#snapshot for Snapshot resources.
func (o SnapshotOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// A fingerprint for the labels being applied to this snapshot, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a snapshot.
func (o SnapshotOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels to apply to this snapshot. These can be later modified by the setLabels method. Label values may be empty.
func (o SnapshotOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Integer license codes indicating which licenses are attached to this snapshot.
func (o SnapshotOutput) LicenseCodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringArrayOutput { return v.LicenseCodes }).(pulumi.StringArrayOutput)
}

// A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses attached (such as a Windows image).
func (o SnapshotOutput) Licenses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringArrayOutput { return v.Licenses }).(pulumi.StringArrayOutput)
}

// An opaque location hint used to place the snapshot close to other resources. This field is for use by internal tools that use the public API.
func (o SnapshotOutput) LocationHint() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.LocationHint }).(pulumi.StringOutput)
}

// Number of days the snapshot should be retained before being deleted automatically.
func (o SnapshotOutput) MaxRetentionDays() pulumi.IntOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.IntOutput { return v.MaxRetentionDays }).(pulumi.IntOutput)
}

// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o SnapshotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SnapshotOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
func (o SnapshotOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// Reserved for future use.
func (o SnapshotOutput) SatisfiesPzs() pulumi.BoolOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.BoolOutput { return v.SatisfiesPzs }).(pulumi.BoolOutput)
}

// Server-defined URL for the resource.
func (o SnapshotOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource's resource id.
func (o SnapshotOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// Encrypts the snapshot using a customer-supplied encryption key. After you encrypt a snapshot using a customer-supplied key, you must provide the same key if you use the snapshot later. For example, you must provide the encryption key when you create a disk from the encrypted snapshot in a future request. Customer-supplied encryption keys do not protect access to metadata of the snapshot. If you do not provide an encryption key when creating the snapshot, then the snapshot will be encrypted using an automatically generated key and you do not need to provide a key to use the snapshot later.
func (o SnapshotOutput) SnapshotEncryptionKey() CustomerEncryptionKeyResponseOutput {
	return o.ApplyT(func(v *Snapshot) CustomerEncryptionKeyResponseOutput { return v.SnapshotEncryptionKey }).(CustomerEncryptionKeyResponseOutput)
}

// Indicates the type of the snapshot.
func (o SnapshotOutput) SnapshotType() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SnapshotType }).(pulumi.StringOutput)
}

// The source disk used to create this snapshot.
func (o SnapshotOutput) SourceDisk() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SourceDisk }).(pulumi.StringOutput)
}

// The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
func (o SnapshotOutput) SourceDiskEncryptionKey() CustomerEncryptionKeyResponseOutput {
	return o.ApplyT(func(v *Snapshot) CustomerEncryptionKeyResponseOutput { return v.SourceDiskEncryptionKey }).(CustomerEncryptionKeyResponseOutput)
}

// The source disk whose recovery checkpoint will be used to create this snapshot.
func (o SnapshotOutput) SourceDiskForRecoveryCheckpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SourceDiskForRecoveryCheckpoint }).(pulumi.StringOutput)
}

// The ID value of the disk used to create this snapshot. This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given disk name.
func (o SnapshotOutput) SourceDiskId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SourceDiskId }).(pulumi.StringOutput)
}

// The source instant snapshot used to create this snapshot. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instantSnapshots/instantSnapshot - projects/project/zones/zone/instantSnapshots/instantSnapshot - zones/zone/instantSnapshots/instantSnapshot
func (o SnapshotOutput) SourceInstantSnapshot() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SourceInstantSnapshot }).(pulumi.StringOutput)
}

// The unique ID of the instant snapshot used to create this snapshot. This value identifies the exact instant snapshot that was used to create this persistent disk. For example, if you created the persistent disk from an instant snapshot that was later deleted and recreated under the same name, the source instant snapshot ID would identify the exact instant snapshot that was used.
func (o SnapshotOutput) SourceInstantSnapshotId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SourceInstantSnapshotId }).(pulumi.StringOutput)
}

// URL of the resource policy which created this scheduled snapshot.
func (o SnapshotOutput) SourceSnapshotSchedulePolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SourceSnapshotSchedulePolicy }).(pulumi.StringOutput)
}

// ID of the resource policy which created this scheduled snapshot.
func (o SnapshotOutput) SourceSnapshotSchedulePolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.SourceSnapshotSchedulePolicyId }).(pulumi.StringOutput)
}

// The status of the snapshot. This can be CREATING, DELETING, FAILED, READY, or UPLOADING.
func (o SnapshotOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A size of the storage used by the snapshot. As snapshots share storage, this number is expected to change with snapshot creation/deletion.
func (o SnapshotOutput) StorageBytes() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.StorageBytes }).(pulumi.StringOutput)
}

// An indicator whether storageBytes is in a stable state or it is being adjusted as a result of shared storage reallocation. This status can either be UPDATING, meaning the size of the snapshot is being updated, or UP_TO_DATE, meaning the size of the snapshot is up-to-date.
func (o SnapshotOutput) StorageBytesStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringOutput { return v.StorageBytesStatus }).(pulumi.StringOutput)
}

// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
func (o SnapshotOutput) StorageLocations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringArrayOutput { return v.StorageLocations }).(pulumi.StringArrayOutput)
}

// A list of user provided licenses represented by a list of URLs to the license resource.
func (o SnapshotOutput) UserLicenses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Snapshot) pulumi.StringArrayOutput { return v.UserLicenses }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotInput)(nil)).Elem(), &Snapshot{})
	pulumi.RegisterOutputType(SnapshotOutput{})
}
