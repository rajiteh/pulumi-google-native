// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an agent pool.
func LookupAgentPool(ctx *pulumi.Context, args *LookupAgentPoolArgs, opts ...pulumi.InvokeOption) (*LookupAgentPoolResult, error) {
	var rv LookupAgentPoolResult
	err := ctx.Invoke("google-native:storagetransfer/v1:getAgentPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAgentPoolArgs struct {
	AgentPoolId string  `pulumi:"agentPoolId"`
	Project     *string `pulumi:"project"`
}

type LookupAgentPoolResult struct {
	// Specifies the bandwidth limit details. If this field is unspecified, the default value is set as 'No Limit'.
	BandwidthLimit BandwidthLimitResponse `pulumi:"bandwidthLimit"`
	// Specifies the client-specified AgentPool description.
	DisplayName string `pulumi:"displayName"`
	// Specifies a unique string that identifies the agent pool. Format: `projects/{project_id}/agentPools/{agent_pool_id}`
	Name string `pulumi:"name"`
	// Specifies the state of the AgentPool.
	State string `pulumi:"state"`
}

func LookupAgentPoolOutput(ctx *pulumi.Context, args LookupAgentPoolOutputArgs, opts ...pulumi.InvokeOption) LookupAgentPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAgentPoolResult, error) {
			args := v.(LookupAgentPoolArgs)
			r, err := LookupAgentPool(ctx, &args, opts...)
			var s LookupAgentPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAgentPoolResultOutput)
}

type LookupAgentPoolOutputArgs struct {
	AgentPoolId pulumi.StringInput    `pulumi:"agentPoolId"`
	Project     pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupAgentPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgentPoolArgs)(nil)).Elem()
}

type LookupAgentPoolResultOutput struct{ *pulumi.OutputState }

func (LookupAgentPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAgentPoolResult)(nil)).Elem()
}

func (o LookupAgentPoolResultOutput) ToLookupAgentPoolResultOutput() LookupAgentPoolResultOutput {
	return o
}

func (o LookupAgentPoolResultOutput) ToLookupAgentPoolResultOutputWithContext(ctx context.Context) LookupAgentPoolResultOutput {
	return o
}

// Specifies the bandwidth limit details. If this field is unspecified, the default value is set as 'No Limit'.
func (o LookupAgentPoolResultOutput) BandwidthLimit() BandwidthLimitResponseOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) BandwidthLimitResponse { return v.BandwidthLimit }).(BandwidthLimitResponseOutput)
}

// Specifies the client-specified AgentPool description.
func (o LookupAgentPoolResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Specifies a unique string that identifies the agent pool. Format: `projects/{project_id}/agentPools/{agent_pool_id}`
func (o LookupAgentPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the state of the AgentPool.
func (o LookupAgentPoolResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAgentPoolResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAgentPoolResultOutput{})
}
