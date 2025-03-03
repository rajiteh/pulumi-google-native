// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a queued resource.
func LookupQueuedResource(ctx *pulumi.Context, args *LookupQueuedResourceArgs, opts ...pulumi.InvokeOption) (*LookupQueuedResourceResult, error) {
	var rv LookupQueuedResourceResult
	err := ctx.Invoke("google-native:tpu/v2alpha1:getQueuedResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueuedResourceArgs struct {
	Location         string  `pulumi:"location"`
	Project          *string `pulumi:"project"`
	QueuedResourceId string  `pulumi:"queuedResourceId"`
}

type LookupQueuedResourceResult struct {
	// The BestEffort tier.
	BestEffort BestEffortResponse `pulumi:"bestEffort"`
	// The Guaranteed tier
	Guaranteed GuaranteedResponse `pulumi:"guaranteed"`
	// Immutable. The name of the QueuedResource.
	Name string `pulumi:"name"`
	// The queueing policy of the QueuedRequest.
	QueueingPolicy QueueingPolicyResponse `pulumi:"queueingPolicy"`
	// Name of the reservation in which the resource should be provisioned. Format: projects/{project}/locations/{zone}/reservations/{reservation}
	ReservationName string `pulumi:"reservationName"`
	// State of the QueuedResource request.
	State QueuedResourceStateResponse `pulumi:"state"`
	// Defines a TPU resource.
	Tpu TpuResponse `pulumi:"tpu"`
}

func LookupQueuedResourceOutput(ctx *pulumi.Context, args LookupQueuedResourceOutputArgs, opts ...pulumi.InvokeOption) LookupQueuedResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQueuedResourceResult, error) {
			args := v.(LookupQueuedResourceArgs)
			r, err := LookupQueuedResource(ctx, &args, opts...)
			var s LookupQueuedResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupQueuedResourceResultOutput)
}

type LookupQueuedResourceOutputArgs struct {
	Location         pulumi.StringInput    `pulumi:"location"`
	Project          pulumi.StringPtrInput `pulumi:"project"`
	QueuedResourceId pulumi.StringInput    `pulumi:"queuedResourceId"`
}

func (LookupQueuedResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueuedResourceArgs)(nil)).Elem()
}

type LookupQueuedResourceResultOutput struct{ *pulumi.OutputState }

func (LookupQueuedResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueuedResourceResult)(nil)).Elem()
}

func (o LookupQueuedResourceResultOutput) ToLookupQueuedResourceResultOutput() LookupQueuedResourceResultOutput {
	return o
}

func (o LookupQueuedResourceResultOutput) ToLookupQueuedResourceResultOutputWithContext(ctx context.Context) LookupQueuedResourceResultOutput {
	return o
}

// The BestEffort tier.
func (o LookupQueuedResourceResultOutput) BestEffort() BestEffortResponseOutput {
	return o.ApplyT(func(v LookupQueuedResourceResult) BestEffortResponse { return v.BestEffort }).(BestEffortResponseOutput)
}

// The Guaranteed tier
func (o LookupQueuedResourceResultOutput) Guaranteed() GuaranteedResponseOutput {
	return o.ApplyT(func(v LookupQueuedResourceResult) GuaranteedResponse { return v.Guaranteed }).(GuaranteedResponseOutput)
}

// Immutable. The name of the QueuedResource.
func (o LookupQueuedResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueuedResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The queueing policy of the QueuedRequest.
func (o LookupQueuedResourceResultOutput) QueueingPolicy() QueueingPolicyResponseOutput {
	return o.ApplyT(func(v LookupQueuedResourceResult) QueueingPolicyResponse { return v.QueueingPolicy }).(QueueingPolicyResponseOutput)
}

// Name of the reservation in which the resource should be provisioned. Format: projects/{project}/locations/{zone}/reservations/{reservation}
func (o LookupQueuedResourceResultOutput) ReservationName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueuedResourceResult) string { return v.ReservationName }).(pulumi.StringOutput)
}

// State of the QueuedResource request.
func (o LookupQueuedResourceResultOutput) State() QueuedResourceStateResponseOutput {
	return o.ApplyT(func(v LookupQueuedResourceResult) QueuedResourceStateResponse { return v.State }).(QueuedResourceStateResponseOutput)
}

// Defines a TPU resource.
func (o LookupQueuedResourceResultOutput) Tpu() TpuResponseOutput {
	return o.ApplyT(func(v LookupQueuedResourceResult) TpuResponse { return v.Tpu }).(TpuResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQueuedResourceResultOutput{})
}
