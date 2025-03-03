// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new backup run on demand.
// Auto-naming is currently not supported for this resource.
type BackupRun struct {
	pulumi.CustomResourceState

	// Specifies the kind of backup, PHYSICAL or DEFAULT_SNAPSHOT.
	BackupKind pulumi.StringOutput `pulumi:"backupKind"`
	// The description of this run, only applicable to on-demand backups.
	Description pulumi.StringOutput `pulumi:"description"`
	// Encryption configuration specific to a backup.
	DiskEncryptionConfiguration DiskEncryptionConfigurationResponseOutput `pulumi:"diskEncryptionConfiguration"`
	// Encryption status specific to a backup.
	DiskEncryptionStatus DiskEncryptionStatusResponseOutput `pulumi:"diskEncryptionStatus"`
	// The time the backup operation completed in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// The time the run was enqueued in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
	EnqueuedTime pulumi.StringOutput `pulumi:"enqueuedTime"`
	// Information about why the backup operation failed. This is only present if the run has the FAILED status.
	Error    OperationErrorResponseOutput `pulumi:"error"`
	Instance pulumi.StringOutput          `pulumi:"instance"`
	// This is always `sql#backupRun`.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Location of the backups.
	Location pulumi.StringOutput `pulumi:"location"`
	Project  pulumi.StringOutput `pulumi:"project"`
	// The URI of this resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The time the backup operation actually started in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// The status of this run.
	Status pulumi.StringOutput `pulumi:"status"`
	// Backup time zone to prevent restores to an instance with a different time zone. Now relevant only for SQL Server.
	TimeZone pulumi.StringOutput `pulumi:"timeZone"`
	// The type of this run; can be either "AUTOMATED" or "ON_DEMAND" or "FINAL". This field defaults to "ON_DEMAND" and is ignored, when specified for insert requests.
	Type pulumi.StringOutput `pulumi:"type"`
	// The start time of the backup window during which this the backup was attempted in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
	WindowStartTime pulumi.StringOutput `pulumi:"windowStartTime"`
}

// NewBackupRun registers a new resource with the given unique name, arguments, and options.
func NewBackupRun(ctx *pulumi.Context,
	name string, args *BackupRunArgs, opts ...pulumi.ResourceOption) (*BackupRun, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"instance",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource BackupRun
	err := ctx.RegisterResource("google-native:sqladmin/v1:BackupRun", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupRun gets an existing BackupRun resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupRun(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupRunState, opts ...pulumi.ResourceOption) (*BackupRun, error) {
	var resource BackupRun
	err := ctx.ReadResource("google-native:sqladmin/v1:BackupRun", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupRun resources.
type backupRunState struct {
}

type BackupRunState struct {
}

func (BackupRunState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupRunState)(nil)).Elem()
}

type backupRunArgs struct {
	// Specifies the kind of backup, PHYSICAL or DEFAULT_SNAPSHOT.
	BackupKind *BackupRunBackupKind `pulumi:"backupKind"`
	// The description of this run, only applicable to on-demand backups.
	Description *string `pulumi:"description"`
	// Encryption configuration specific to a backup.
	DiskEncryptionConfiguration *DiskEncryptionConfiguration `pulumi:"diskEncryptionConfiguration"`
	// Encryption status specific to a backup.
	DiskEncryptionStatus *DiskEncryptionStatus `pulumi:"diskEncryptionStatus"`
	// The time the backup operation completed in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
	EndTime *string `pulumi:"endTime"`
	// The time the run was enqueued in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
	EnqueuedTime *string `pulumi:"enqueuedTime"`
	// Information about why the backup operation failed. This is only present if the run has the FAILED status.
	Error *OperationError `pulumi:"error"`
	// The identifier for this backup run. Unique only for a specific Cloud SQL instance.
	Id *string `pulumi:"id"`
	// Name of the database instance.
	Instance string `pulumi:"instance"`
	// This is always `sql#backupRun`.
	Kind *string `pulumi:"kind"`
	// Location of the backups.
	Location *string `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// The URI of this resource.
	SelfLink *string `pulumi:"selfLink"`
	// The time the backup operation actually started in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
	StartTime *string `pulumi:"startTime"`
	// Backup time zone to prevent restores to an instance with a different time zone. Now relevant only for SQL Server.
	TimeZone *string `pulumi:"timeZone"`
	// The type of this run; can be either "AUTOMATED" or "ON_DEMAND" or "FINAL". This field defaults to "ON_DEMAND" and is ignored, when specified for insert requests.
	Type *BackupRunType `pulumi:"type"`
	// The start time of the backup window during which this the backup was attempted in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
	WindowStartTime *string `pulumi:"windowStartTime"`
}

// The set of arguments for constructing a BackupRun resource.
type BackupRunArgs struct {
	// Specifies the kind of backup, PHYSICAL or DEFAULT_SNAPSHOT.
	BackupKind BackupRunBackupKindPtrInput
	// The description of this run, only applicable to on-demand backups.
	Description pulumi.StringPtrInput
	// Encryption configuration specific to a backup.
	DiskEncryptionConfiguration DiskEncryptionConfigurationPtrInput
	// Encryption status specific to a backup.
	DiskEncryptionStatus DiskEncryptionStatusPtrInput
	// The time the backup operation completed in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
	EndTime pulumi.StringPtrInput
	// The time the run was enqueued in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
	EnqueuedTime pulumi.StringPtrInput
	// Information about why the backup operation failed. This is only present if the run has the FAILED status.
	Error OperationErrorPtrInput
	// The identifier for this backup run. Unique only for a specific Cloud SQL instance.
	Id pulumi.StringPtrInput
	// Name of the database instance.
	Instance pulumi.StringInput
	// This is always `sql#backupRun`.
	Kind pulumi.StringPtrInput
	// Location of the backups.
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// The URI of this resource.
	SelfLink pulumi.StringPtrInput
	// The time the backup operation actually started in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
	StartTime pulumi.StringPtrInput
	// Backup time zone to prevent restores to an instance with a different time zone. Now relevant only for SQL Server.
	TimeZone pulumi.StringPtrInput
	// The type of this run; can be either "AUTOMATED" or "ON_DEMAND" or "FINAL". This field defaults to "ON_DEMAND" and is ignored, when specified for insert requests.
	Type BackupRunTypePtrInput
	// The start time of the backup window during which this the backup was attempted in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
	WindowStartTime pulumi.StringPtrInput
}

func (BackupRunArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupRunArgs)(nil)).Elem()
}

type BackupRunInput interface {
	pulumi.Input

	ToBackupRunOutput() BackupRunOutput
	ToBackupRunOutputWithContext(ctx context.Context) BackupRunOutput
}

func (*BackupRun) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupRun)(nil)).Elem()
}

func (i *BackupRun) ToBackupRunOutput() BackupRunOutput {
	return i.ToBackupRunOutputWithContext(context.Background())
}

func (i *BackupRun) ToBackupRunOutputWithContext(ctx context.Context) BackupRunOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupRunOutput)
}

type BackupRunOutput struct{ *pulumi.OutputState }

func (BackupRunOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupRun)(nil)).Elem()
}

func (o BackupRunOutput) ToBackupRunOutput() BackupRunOutput {
	return o
}

func (o BackupRunOutput) ToBackupRunOutputWithContext(ctx context.Context) BackupRunOutput {
	return o
}

// Specifies the kind of backup, PHYSICAL or DEFAULT_SNAPSHOT.
func (o BackupRunOutput) BackupKind() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupRun) pulumi.StringOutput { return v.BackupKind }).(pulumi.StringOutput)
}

// The description of this run, only applicable to on-demand backups.
func (o BackupRunOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupRun) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Encryption configuration specific to a backup.
func (o BackupRunOutput) DiskEncryptionConfiguration() DiskEncryptionConfigurationResponseOutput {
	return o.ApplyT(func(v *BackupRun) DiskEncryptionConfigurationResponseOutput { return v.DiskEncryptionConfiguration }).(DiskEncryptionConfigurationResponseOutput)
}

// Encryption status specific to a backup.
func (o BackupRunOutput) DiskEncryptionStatus() DiskEncryptionStatusResponseOutput {
	return o.ApplyT(func(v *BackupRun) DiskEncryptionStatusResponseOutput { return v.DiskEncryptionStatus }).(DiskEncryptionStatusResponseOutput)
}

// The time the backup operation completed in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
func (o BackupRunOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupRun) pulumi.StringOutput { return v.EndTime }).(pulumi.StringOutput)
}

// The time the run was enqueued in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
func (o BackupRunOutput) EnqueuedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupRun) pulumi.StringOutput { return v.EnqueuedTime }).(pulumi.StringOutput)
}

// Information about why the backup operation failed. This is only present if the run has the FAILED status.
func (o BackupRunOutput) Error() OperationErrorResponseOutput {
	return o.ApplyT(func(v *BackupRun) OperationErrorResponseOutput { return v.Error }).(OperationErrorResponseOutput)
}

func (o BackupRunOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupRun) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
}

// This is always `sql#backupRun`.
func (o BackupRunOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupRun) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Location of the backups.
func (o BackupRunOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupRun) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o BackupRunOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupRun) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The URI of this resource.
func (o BackupRunOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupRun) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The time the backup operation actually started in UTC timezone in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
func (o BackupRunOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupRun) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

// The status of this run.
func (o BackupRunOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupRun) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Backup time zone to prevent restores to an instance with a different time zone. Now relevant only for SQL Server.
func (o BackupRunOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupRun) pulumi.StringOutput { return v.TimeZone }).(pulumi.StringOutput)
}

// The type of this run; can be either "AUTOMATED" or "ON_DEMAND" or "FINAL". This field defaults to "ON_DEMAND" and is ignored, when specified for insert requests.
func (o BackupRunOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupRun) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The start time of the backup window during which this the backup was attempted in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
func (o BackupRunOutput) WindowStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupRun) pulumi.StringOutput { return v.WindowStartTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupRunInput)(nil)).Elem(), &BackupRun{})
	pulumi.RegisterOutputType(BackupRunOutput{})
}
