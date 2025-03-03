// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves all information of the specified resource policy.
func LookupResourcePolicy(ctx *pulumi.Context, args *LookupResourcePolicyArgs, opts ...pulumi.InvokeOption) (*LookupResourcePolicyResult, error) {
	var rv LookupResourcePolicyResult
	err := ctx.Invoke("google-native:compute/beta:getResourcePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourcePolicyArgs struct {
	Project        *string `pulumi:"project"`
	Region         string  `pulumi:"region"`
	ResourcePolicy string  `pulumi:"resourcePolicy"`
}

type LookupResourcePolicyResult struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	Description       string `pulumi:"description"`
	// Resource policy for disk consistency groups.
	DiskConsistencyGroupPolicy ResourcePolicyDiskConsistencyGroupPolicyResponse `pulumi:"diskConsistencyGroupPolicy"`
	// Resource policy for instances for placement configuration.
	GroupPlacementPolicy ResourcePolicyGroupPlacementPolicyResponse `pulumi:"groupPlacementPolicy"`
	// Resource policy for scheduling instance operations.
	InstanceSchedulePolicy ResourcePolicyInstanceSchedulePolicyResponse `pulumi:"instanceSchedulePolicy"`
	// Type of the resource. Always compute#resource_policies for resource policies.
	Kind string `pulumi:"kind"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name   string `pulumi:"name"`
	Region string `pulumi:"region"`
	// The system status of the resource policy.
	ResourceStatus ResourcePolicyResourceStatusResponse `pulumi:"resourceStatus"`
	// Server-defined fully-qualified URL for this resource.
	SelfLink string `pulumi:"selfLink"`
	// Resource policy for persistent disks for creating snapshots.
	SnapshotSchedulePolicy ResourcePolicySnapshotSchedulePolicyResponse `pulumi:"snapshotSchedulePolicy"`
	// The status of resource policy creation.
	Status string `pulumi:"status"`
}

func LookupResourcePolicyOutput(ctx *pulumi.Context, args LookupResourcePolicyOutputArgs, opts ...pulumi.InvokeOption) LookupResourcePolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourcePolicyResult, error) {
			args := v.(LookupResourcePolicyArgs)
			r, err := LookupResourcePolicy(ctx, &args, opts...)
			var s LookupResourcePolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourcePolicyResultOutput)
}

type LookupResourcePolicyOutputArgs struct {
	Project        pulumi.StringPtrInput `pulumi:"project"`
	Region         pulumi.StringInput    `pulumi:"region"`
	ResourcePolicy pulumi.StringInput    `pulumi:"resourcePolicy"`
}

func (LookupResourcePolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourcePolicyArgs)(nil)).Elem()
}

type LookupResourcePolicyResultOutput struct{ *pulumi.OutputState }

func (LookupResourcePolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourcePolicyResult)(nil)).Elem()
}

func (o LookupResourcePolicyResultOutput) ToLookupResourcePolicyResultOutput() LookupResourcePolicyResultOutput {
	return o
}

func (o LookupResourcePolicyResultOutput) ToLookupResourcePolicyResultOutputWithContext(ctx context.Context) LookupResourcePolicyResultOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o LookupResourcePolicyResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePolicyResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

func (o LookupResourcePolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// Resource policy for disk consistency groups.
func (o LookupResourcePolicyResultOutput) DiskConsistencyGroupPolicy() ResourcePolicyDiskConsistencyGroupPolicyResponseOutput {
	return o.ApplyT(func(v LookupResourcePolicyResult) ResourcePolicyDiskConsistencyGroupPolicyResponse {
		return v.DiskConsistencyGroupPolicy
	}).(ResourcePolicyDiskConsistencyGroupPolicyResponseOutput)
}

// Resource policy for instances for placement configuration.
func (o LookupResourcePolicyResultOutput) GroupPlacementPolicy() ResourcePolicyGroupPlacementPolicyResponseOutput {
	return o.ApplyT(func(v LookupResourcePolicyResult) ResourcePolicyGroupPlacementPolicyResponse {
		return v.GroupPlacementPolicy
	}).(ResourcePolicyGroupPlacementPolicyResponseOutput)
}

// Resource policy for scheduling instance operations.
func (o LookupResourcePolicyResultOutput) InstanceSchedulePolicy() ResourcePolicyInstanceSchedulePolicyResponseOutput {
	return o.ApplyT(func(v LookupResourcePolicyResult) ResourcePolicyInstanceSchedulePolicyResponse {
		return v.InstanceSchedulePolicy
	}).(ResourcePolicyInstanceSchedulePolicyResponseOutput)
}

// Type of the resource. Always compute#resource_policies for resource policies.
func (o LookupResourcePolicyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePolicyResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupResourcePolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupResourcePolicyResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePolicyResult) string { return v.Region }).(pulumi.StringOutput)
}

// The system status of the resource policy.
func (o LookupResourcePolicyResultOutput) ResourceStatus() ResourcePolicyResourceStatusResponseOutput {
	return o.ApplyT(func(v LookupResourcePolicyResult) ResourcePolicyResourceStatusResponse { return v.ResourceStatus }).(ResourcePolicyResourceStatusResponseOutput)
}

// Server-defined fully-qualified URL for this resource.
func (o LookupResourcePolicyResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePolicyResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Resource policy for persistent disks for creating snapshots.
func (o LookupResourcePolicyResultOutput) SnapshotSchedulePolicy() ResourcePolicySnapshotSchedulePolicyResponseOutput {
	return o.ApplyT(func(v LookupResourcePolicyResult) ResourcePolicySnapshotSchedulePolicyResponse {
		return v.SnapshotSchedulePolicy
	}).(ResourcePolicySnapshotSchedulePolicyResponseOutput)
}

// The status of resource policy creation.
func (o LookupResourcePolicyResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourcePolicyResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourcePolicyResultOutput{})
}
