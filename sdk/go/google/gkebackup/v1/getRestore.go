// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the details of a single Restore.
func LookupRestore(ctx *pulumi.Context, args *LookupRestoreArgs, opts ...pulumi.InvokeOption) (*LookupRestoreResult, error) {
	var rv LookupRestoreResult
	err := ctx.Invoke("google-native:gkebackup/v1:getRestore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRestoreArgs struct {
	Location      string  `pulumi:"location"`
	Project       *string `pulumi:"project"`
	RestoreId     string  `pulumi:"restoreId"`
	RestorePlanId string  `pulumi:"restorePlanId"`
}

type LookupRestoreResult struct {
	// Immutable. A reference to the Backup used as the source from which this Restore will restore. Note that this Backup must be a sub-resource of the RestorePlan's backup_plan. Format: `projects/*/locations/*/backupPlans/*/backups/*`.
	Backup string `pulumi:"backup"`
	// The target cluster into which this Restore will restore data. Valid formats: - `projects/*/locations/*/clusters/*` - `projects/*/zones/*/clusters/*` Inherited from parent RestorePlan's cluster value.
	Cluster string `pulumi:"cluster"`
	// Timestamp of when the restore operation completed.
	CompleteTime string `pulumi:"completeTime"`
	// The timestamp when this Restore resource was created.
	CreateTime string `pulumi:"createTime"`
	// User specified descriptive string for this Restore.
	Description string `pulumi:"description"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a restore from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform restore updates in order to avoid race conditions: An `etag` is returned in the response to `GetRestore`, and systems are expected to put that etag in the request to `UpdateRestore` or `DeleteRestore` to ensure that their change will be applied to the same version of the resource.
	Etag string `pulumi:"etag"`
	// A set of custom labels supplied by user.
	Labels map[string]string `pulumi:"labels"`
	// The full name of the Restore resource. Format: `projects/*/locations/*/restorePlans/*/restores/*`
	Name string `pulumi:"name"`
	// Number of resources excluded during the restore execution.
	ResourcesExcludedCount int `pulumi:"resourcesExcludedCount"`
	// Number of resources that failed to be restored during the restore execution.
	ResourcesFailedCount int `pulumi:"resourcesFailedCount"`
	// Number of resources restored during the restore execution.
	ResourcesRestoredCount int `pulumi:"resourcesRestoredCount"`
	// Configuration of the Restore. Inherited from parent RestorePlan's restore_config.
	RestoreConfig RestoreConfigResponse `pulumi:"restoreConfig"`
	// The current state of the Restore.
	State string `pulumi:"state"`
	// Human-readable description of why the Restore is in its current state.
	StateReason string `pulumi:"stateReason"`
	// Server generated global unique identifier of [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier) format.
	Uid string `pulumi:"uid"`
	// The timestamp when this Restore resource was last updated.
	UpdateTime string `pulumi:"updateTime"`
	// Number of volumes restored during the restore execution.
	VolumesRestoredCount int `pulumi:"volumesRestoredCount"`
}

func LookupRestoreOutput(ctx *pulumi.Context, args LookupRestoreOutputArgs, opts ...pulumi.InvokeOption) LookupRestoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRestoreResult, error) {
			args := v.(LookupRestoreArgs)
			r, err := LookupRestore(ctx, &args, opts...)
			var s LookupRestoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRestoreResultOutput)
}

type LookupRestoreOutputArgs struct {
	Location      pulumi.StringInput    `pulumi:"location"`
	Project       pulumi.StringPtrInput `pulumi:"project"`
	RestoreId     pulumi.StringInput    `pulumi:"restoreId"`
	RestorePlanId pulumi.StringInput    `pulumi:"restorePlanId"`
}

func (LookupRestoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRestoreArgs)(nil)).Elem()
}

type LookupRestoreResultOutput struct{ *pulumi.OutputState }

func (LookupRestoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRestoreResult)(nil)).Elem()
}

func (o LookupRestoreResultOutput) ToLookupRestoreResultOutput() LookupRestoreResultOutput {
	return o
}

func (o LookupRestoreResultOutput) ToLookupRestoreResultOutputWithContext(ctx context.Context) LookupRestoreResultOutput {
	return o
}

// Immutable. A reference to the Backup used as the source from which this Restore will restore. Note that this Backup must be a sub-resource of the RestorePlan's backup_plan. Format: `projects/*/locations/*/backupPlans/*/backups/*`.
func (o LookupRestoreResultOutput) Backup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestoreResult) string { return v.Backup }).(pulumi.StringOutput)
}

// The target cluster into which this Restore will restore data. Valid formats: - `projects/*/locations/*/clusters/*` - `projects/*/zones/*/clusters/*` Inherited from parent RestorePlan's cluster value.
func (o LookupRestoreResultOutput) Cluster() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestoreResult) string { return v.Cluster }).(pulumi.StringOutput)
}

// Timestamp of when the restore operation completed.
func (o LookupRestoreResultOutput) CompleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestoreResult) string { return v.CompleteTime }).(pulumi.StringOutput)
}

// The timestamp when this Restore resource was created.
func (o LookupRestoreResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestoreResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// User specified descriptive string for this Restore.
func (o LookupRestoreResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestoreResult) string { return v.Description }).(pulumi.StringOutput)
}

// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a restore from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform restore updates in order to avoid race conditions: An `etag` is returned in the response to `GetRestore`, and systems are expected to put that etag in the request to `UpdateRestore` or `DeleteRestore` to ensure that their change will be applied to the same version of the resource.
func (o LookupRestoreResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestoreResult) string { return v.Etag }).(pulumi.StringOutput)
}

// A set of custom labels supplied by user.
func (o LookupRestoreResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRestoreResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The full name of the Restore resource. Format: `projects/*/locations/*/restorePlans/*/restores/*`
func (o LookupRestoreResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestoreResult) string { return v.Name }).(pulumi.StringOutput)
}

// Number of resources excluded during the restore execution.
func (o LookupRestoreResultOutput) ResourcesExcludedCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRestoreResult) int { return v.ResourcesExcludedCount }).(pulumi.IntOutput)
}

// Number of resources that failed to be restored during the restore execution.
func (o LookupRestoreResultOutput) ResourcesFailedCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRestoreResult) int { return v.ResourcesFailedCount }).(pulumi.IntOutput)
}

// Number of resources restored during the restore execution.
func (o LookupRestoreResultOutput) ResourcesRestoredCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRestoreResult) int { return v.ResourcesRestoredCount }).(pulumi.IntOutput)
}

// Configuration of the Restore. Inherited from parent RestorePlan's restore_config.
func (o LookupRestoreResultOutput) RestoreConfig() RestoreConfigResponseOutput {
	return o.ApplyT(func(v LookupRestoreResult) RestoreConfigResponse { return v.RestoreConfig }).(RestoreConfigResponseOutput)
}

// The current state of the Restore.
func (o LookupRestoreResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestoreResult) string { return v.State }).(pulumi.StringOutput)
}

// Human-readable description of why the Restore is in its current state.
func (o LookupRestoreResultOutput) StateReason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestoreResult) string { return v.StateReason }).(pulumi.StringOutput)
}

// Server generated global unique identifier of [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier) format.
func (o LookupRestoreResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestoreResult) string { return v.Uid }).(pulumi.StringOutput)
}

// The timestamp when this Restore resource was last updated.
func (o LookupRestoreResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRestoreResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

// Number of volumes restored during the restore execution.
func (o LookupRestoreResultOutput) VolumesRestoredCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRestoreResult) int { return v.VolumesRestoredCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRestoreResultOutput{})
}
