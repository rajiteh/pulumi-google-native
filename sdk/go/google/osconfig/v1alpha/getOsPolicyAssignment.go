// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve an existing OS policy assignment. This method always returns the latest revision. In order to retrieve a previous revision of the assignment, also provide the revision ID in the `name` parameter.
func LookupOsPolicyAssignment(ctx *pulumi.Context, args *LookupOsPolicyAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupOsPolicyAssignmentResult, error) {
	var rv LookupOsPolicyAssignmentResult
	err := ctx.Invoke("google-native:osconfig/v1alpha:getOsPolicyAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOsPolicyAssignmentArgs struct {
	Location             string  `pulumi:"location"`
	OsPolicyAssignmentId string  `pulumi:"osPolicyAssignmentId"`
	Project              *string `pulumi:"project"`
}

type LookupOsPolicyAssignmentResult struct {
	// Indicates that this revision has been successfully rolled out in this zone and new VMs will be assigned OS policies from this revision. For a given OS policy assignment, there is only one revision with a value of `true` for this field.
	Baseline bool `pulumi:"baseline"`
	// Indicates that this revision deletes the OS policy assignment.
	Deleted bool `pulumi:"deleted"`
	// OS policy assignment description. Length of the description is limited to 1024 characters.
	Description string `pulumi:"description"`
	// The etag for this OS policy assignment. If this is provided on update, it must match the server's etag.
	Etag string `pulumi:"etag"`
	// Filter to select VMs.
	InstanceFilter OSPolicyAssignmentInstanceFilterResponse `pulumi:"instanceFilter"`
	// Resource name. Format: `projects/{project_number}/locations/{location}/osPolicyAssignments/{os_policy_assignment_id}` This field is ignored when you create an OS policy assignment.
	Name string `pulumi:"name"`
	// List of OS policies to be applied to the VMs.
	OsPolicies []OSPolicyResponse `pulumi:"osPolicies"`
	// Indicates that reconciliation is in progress for the revision. This value is `true` when the `rollout_state` is one of: * IN_PROGRESS * CANCELLING
	Reconciling bool `pulumi:"reconciling"`
	// The timestamp that the revision was created.
	RevisionCreateTime string `pulumi:"revisionCreateTime"`
	// The assignment revision ID A new revision is committed whenever a rollout is triggered for a OS policy assignment
	RevisionId string `pulumi:"revisionId"`
	// Rollout to deploy the OS policy assignment. A rollout is triggered in the following situations: 1) OSPolicyAssignment is created. 2) OSPolicyAssignment is updated and the update contains changes to one of the following fields: - instance_filter - os_policies 3) OSPolicyAssignment is deleted.
	Rollout OSPolicyAssignmentRolloutResponse `pulumi:"rollout"`
	// OS policy assignment rollout state
	RolloutState string `pulumi:"rolloutState"`
	// Server generated unique id for the OS policy assignment resource.
	Uid string `pulumi:"uid"`
}

func LookupOsPolicyAssignmentOutput(ctx *pulumi.Context, args LookupOsPolicyAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupOsPolicyAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOsPolicyAssignmentResult, error) {
			args := v.(LookupOsPolicyAssignmentArgs)
			r, err := LookupOsPolicyAssignment(ctx, &args, opts...)
			var s LookupOsPolicyAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOsPolicyAssignmentResultOutput)
}

type LookupOsPolicyAssignmentOutputArgs struct {
	Location             pulumi.StringInput    `pulumi:"location"`
	OsPolicyAssignmentId pulumi.StringInput    `pulumi:"osPolicyAssignmentId"`
	Project              pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupOsPolicyAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOsPolicyAssignmentArgs)(nil)).Elem()
}

type LookupOsPolicyAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupOsPolicyAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOsPolicyAssignmentResult)(nil)).Elem()
}

func (o LookupOsPolicyAssignmentResultOutput) ToLookupOsPolicyAssignmentResultOutput() LookupOsPolicyAssignmentResultOutput {
	return o
}

func (o LookupOsPolicyAssignmentResultOutput) ToLookupOsPolicyAssignmentResultOutputWithContext(ctx context.Context) LookupOsPolicyAssignmentResultOutput {
	return o
}

// Indicates that this revision has been successfully rolled out in this zone and new VMs will be assigned OS policies from this revision. For a given OS policy assignment, there is only one revision with a value of `true` for this field.
func (o LookupOsPolicyAssignmentResultOutput) Baseline() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOsPolicyAssignmentResult) bool { return v.Baseline }).(pulumi.BoolOutput)
}

// Indicates that this revision deletes the OS policy assignment.
func (o LookupOsPolicyAssignmentResultOutput) Deleted() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOsPolicyAssignmentResult) bool { return v.Deleted }).(pulumi.BoolOutput)
}

// OS policy assignment description. Length of the description is limited to 1024 characters.
func (o LookupOsPolicyAssignmentResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOsPolicyAssignmentResult) string { return v.Description }).(pulumi.StringOutput)
}

// The etag for this OS policy assignment. If this is provided on update, it must match the server's etag.
func (o LookupOsPolicyAssignmentResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOsPolicyAssignmentResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Filter to select VMs.
func (o LookupOsPolicyAssignmentResultOutput) InstanceFilter() OSPolicyAssignmentInstanceFilterResponseOutput {
	return o.ApplyT(func(v LookupOsPolicyAssignmentResult) OSPolicyAssignmentInstanceFilterResponse {
		return v.InstanceFilter
	}).(OSPolicyAssignmentInstanceFilterResponseOutput)
}

// Resource name. Format: `projects/{project_number}/locations/{location}/osPolicyAssignments/{os_policy_assignment_id}` This field is ignored when you create an OS policy assignment.
func (o LookupOsPolicyAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOsPolicyAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of OS policies to be applied to the VMs.
func (o LookupOsPolicyAssignmentResultOutput) OsPolicies() OSPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupOsPolicyAssignmentResult) []OSPolicyResponse { return v.OsPolicies }).(OSPolicyResponseArrayOutput)
}

// Indicates that reconciliation is in progress for the revision. This value is `true` when the `rollout_state` is one of: * IN_PROGRESS * CANCELLING
func (o LookupOsPolicyAssignmentResultOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOsPolicyAssignmentResult) bool { return v.Reconciling }).(pulumi.BoolOutput)
}

// The timestamp that the revision was created.
func (o LookupOsPolicyAssignmentResultOutput) RevisionCreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOsPolicyAssignmentResult) string { return v.RevisionCreateTime }).(pulumi.StringOutput)
}

// The assignment revision ID A new revision is committed whenever a rollout is triggered for a OS policy assignment
func (o LookupOsPolicyAssignmentResultOutput) RevisionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOsPolicyAssignmentResult) string { return v.RevisionId }).(pulumi.StringOutput)
}

// Rollout to deploy the OS policy assignment. A rollout is triggered in the following situations: 1) OSPolicyAssignment is created. 2) OSPolicyAssignment is updated and the update contains changes to one of the following fields: - instance_filter - os_policies 3) OSPolicyAssignment is deleted.
func (o LookupOsPolicyAssignmentResultOutput) Rollout() OSPolicyAssignmentRolloutResponseOutput {
	return o.ApplyT(func(v LookupOsPolicyAssignmentResult) OSPolicyAssignmentRolloutResponse { return v.Rollout }).(OSPolicyAssignmentRolloutResponseOutput)
}

// OS policy assignment rollout state
func (o LookupOsPolicyAssignmentResultOutput) RolloutState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOsPolicyAssignmentResult) string { return v.RolloutState }).(pulumi.StringOutput)
}

// Server generated unique id for the OS policy assignment resource.
func (o LookupOsPolicyAssignmentResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOsPolicyAssignmentResult) string { return v.Uid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOsPolicyAssignmentResultOutput{})
}
