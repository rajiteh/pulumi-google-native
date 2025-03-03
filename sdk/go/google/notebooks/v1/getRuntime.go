// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single Runtime. The location must be a regional endpoint rather than zonal.
func LookupRuntime(ctx *pulumi.Context, args *LookupRuntimeArgs, opts ...pulumi.InvokeOption) (*LookupRuntimeResult, error) {
	var rv LookupRuntimeResult
	err := ctx.Invoke("google-native:notebooks/v1:getRuntime", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRuntimeArgs struct {
	Location  string  `pulumi:"location"`
	Project   *string `pulumi:"project"`
	RuntimeId string  `pulumi:"runtimeId"`
}

type LookupRuntimeResult struct {
	// The config settings for accessing runtime.
	AccessConfig RuntimeAccessConfigResponse `pulumi:"accessConfig"`
	// Runtime creation time.
	CreateTime string `pulumi:"createTime"`
	// Runtime health_state.
	HealthState string `pulumi:"healthState"`
	// Optional. The labels to associate with this Managed Notebook or Runtime. Label **keys** must contain 1 to 63 characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be empty, but, if present, must contain 1 to 63 characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a cluster.
	Labels map[string]string `pulumi:"labels"`
	// Contains Runtime daemon metrics such as Service status and JupyterLab stats.
	Metrics RuntimeMetricsResponse `pulumi:"metrics"`
	// The resource name of the runtime. Format: `projects/{project}/locations/{location}/runtimes/{runtimeId}`
	Name string `pulumi:"name"`
	// The config settings for software inside the runtime.
	SoftwareConfig RuntimeSoftwareConfigResponse `pulumi:"softwareConfig"`
	// Runtime state.
	State string `pulumi:"state"`
	// Runtime update time.
	UpdateTime string `pulumi:"updateTime"`
	// Use a Compute Engine VM image to start the managed notebook instance.
	VirtualMachine VirtualMachineResponse `pulumi:"virtualMachine"`
}

func LookupRuntimeOutput(ctx *pulumi.Context, args LookupRuntimeOutputArgs, opts ...pulumi.InvokeOption) LookupRuntimeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRuntimeResult, error) {
			args := v.(LookupRuntimeArgs)
			r, err := LookupRuntime(ctx, &args, opts...)
			var s LookupRuntimeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRuntimeResultOutput)
}

type LookupRuntimeOutputArgs struct {
	Location  pulumi.StringInput    `pulumi:"location"`
	Project   pulumi.StringPtrInput `pulumi:"project"`
	RuntimeId pulumi.StringInput    `pulumi:"runtimeId"`
}

func (LookupRuntimeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuntimeArgs)(nil)).Elem()
}

type LookupRuntimeResultOutput struct{ *pulumi.OutputState }

func (LookupRuntimeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuntimeResult)(nil)).Elem()
}

func (o LookupRuntimeResultOutput) ToLookupRuntimeResultOutput() LookupRuntimeResultOutput {
	return o
}

func (o LookupRuntimeResultOutput) ToLookupRuntimeResultOutputWithContext(ctx context.Context) LookupRuntimeResultOutput {
	return o
}

// The config settings for accessing runtime.
func (o LookupRuntimeResultOutput) AccessConfig() RuntimeAccessConfigResponseOutput {
	return o.ApplyT(func(v LookupRuntimeResult) RuntimeAccessConfigResponse { return v.AccessConfig }).(RuntimeAccessConfigResponseOutput)
}

// Runtime creation time.
func (o LookupRuntimeResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuntimeResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Runtime health_state.
func (o LookupRuntimeResultOutput) HealthState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuntimeResult) string { return v.HealthState }).(pulumi.StringOutput)
}

// Optional. The labels to associate with this Managed Notebook or Runtime. Label **keys** must contain 1 to 63 characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be empty, but, if present, must contain 1 to 63 characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a cluster.
func (o LookupRuntimeResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRuntimeResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Contains Runtime daemon metrics such as Service status and JupyterLab stats.
func (o LookupRuntimeResultOutput) Metrics() RuntimeMetricsResponseOutput {
	return o.ApplyT(func(v LookupRuntimeResult) RuntimeMetricsResponse { return v.Metrics }).(RuntimeMetricsResponseOutput)
}

// The resource name of the runtime. Format: `projects/{project}/locations/{location}/runtimes/{runtimeId}`
func (o LookupRuntimeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuntimeResult) string { return v.Name }).(pulumi.StringOutput)
}

// The config settings for software inside the runtime.
func (o LookupRuntimeResultOutput) SoftwareConfig() RuntimeSoftwareConfigResponseOutput {
	return o.ApplyT(func(v LookupRuntimeResult) RuntimeSoftwareConfigResponse { return v.SoftwareConfig }).(RuntimeSoftwareConfigResponseOutput)
}

// Runtime state.
func (o LookupRuntimeResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuntimeResult) string { return v.State }).(pulumi.StringOutput)
}

// Runtime update time.
func (o LookupRuntimeResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuntimeResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

// Use a Compute Engine VM image to start the managed notebook instance.
func (o LookupRuntimeResultOutput) VirtualMachine() VirtualMachineResponseOutput {
	return o.ApplyT(func(v LookupRuntimeResult) VirtualMachineResponse { return v.VirtualMachine }).(VirtualMachineResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRuntimeResultOutput{})
}
