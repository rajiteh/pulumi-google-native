// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the requested workstation.
func LookupWorkstation(ctx *pulumi.Context, args *LookupWorkstationArgs, opts ...pulumi.InvokeOption) (*LookupWorkstationResult, error) {
	var rv LookupWorkstationResult
	err := ctx.Invoke("google-native:workstations/v1beta:getWorkstation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkstationArgs struct {
	Location             string  `pulumi:"location"`
	Project              *string `pulumi:"project"`
	WorkstationClusterId string  `pulumi:"workstationClusterId"`
	WorkstationConfigId  string  `pulumi:"workstationConfigId"`
	WorkstationId        string  `pulumi:"workstationId"`
}

type LookupWorkstationResult struct {
	// Client-specified annotations.
	Annotations map[string]string `pulumi:"annotations"`
	// Time when this resource was created.
	CreateTime string `pulumi:"createTime"`
	// Time when this resource was soft-deleted.
	DeleteTime string `pulumi:"deleteTime"`
	// Human-readable name for this resource.
	DisplayName string `pulumi:"displayName"`
	// Environment variables passed to the workstation container's entrypoint.
	Env map[string]string `pulumi:"env"`
	// Checksum computed by the server. May be sent on update and delete requests to make sure that the client has an up-to-date value before proceeding.
	Etag string `pulumi:"etag"`
	// Host to which clients can send HTTPS traffic that will be received by the workstation. Authorized traffic will be received to the workstation as HTTP on port 80. To send traffic to a different port, clients may prefix the host with the destination port in the format `{port}-{host}`.
	Host string `pulumi:"host"`
	// Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
	Labels map[string]string `pulumi:"labels"`
	// Full name of this resource.
	Name string `pulumi:"name"`
	// Indicates whether this resource is currently being updated to match its intended state.
	Reconciling bool `pulumi:"reconciling"`
	// Current state of the workstation.
	State string `pulumi:"state"`
	// A system-assigned unique identified for this resource.
	Uid string `pulumi:"uid"`
	// Time when this resource was most recently updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupWorkstationOutput(ctx *pulumi.Context, args LookupWorkstationOutputArgs, opts ...pulumi.InvokeOption) LookupWorkstationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkstationResult, error) {
			args := v.(LookupWorkstationArgs)
			r, err := LookupWorkstation(ctx, &args, opts...)
			var s LookupWorkstationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkstationResultOutput)
}

type LookupWorkstationOutputArgs struct {
	Location             pulumi.StringInput    `pulumi:"location"`
	Project              pulumi.StringPtrInput `pulumi:"project"`
	WorkstationClusterId pulumi.StringInput    `pulumi:"workstationClusterId"`
	WorkstationConfigId  pulumi.StringInput    `pulumi:"workstationConfigId"`
	WorkstationId        pulumi.StringInput    `pulumi:"workstationId"`
}

func (LookupWorkstationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkstationArgs)(nil)).Elem()
}

type LookupWorkstationResultOutput struct{ *pulumi.OutputState }

func (LookupWorkstationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkstationResult)(nil)).Elem()
}

func (o LookupWorkstationResultOutput) ToLookupWorkstationResultOutput() LookupWorkstationResultOutput {
	return o
}

func (o LookupWorkstationResultOutput) ToLookupWorkstationResultOutputWithContext(ctx context.Context) LookupWorkstationResultOutput {
	return o
}

// Client-specified annotations.
func (o LookupWorkstationResultOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkstationResult) map[string]string { return v.Annotations }).(pulumi.StringMapOutput)
}

// Time when this resource was created.
func (o LookupWorkstationResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Time when this resource was soft-deleted.
func (o LookupWorkstationResultOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationResult) string { return v.DeleteTime }).(pulumi.StringOutput)
}

// Human-readable name for this resource.
func (o LookupWorkstationResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Environment variables passed to the workstation container's entrypoint.
func (o LookupWorkstationResultOutput) Env() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkstationResult) map[string]string { return v.Env }).(pulumi.StringMapOutput)
}

// Checksum computed by the server. May be sent on update and delete requests to make sure that the client has an up-to-date value before proceeding.
func (o LookupWorkstationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Host to which clients can send HTTPS traffic that will be received by the workstation. Authorized traffic will be received to the workstation as HTTP on port 80. To send traffic to a different port, clients may prefix the host with the destination port in the format `{port}-{host}`.
func (o LookupWorkstationResultOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationResult) string { return v.Host }).(pulumi.StringOutput)
}

// Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
func (o LookupWorkstationResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkstationResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Full name of this resource.
func (o LookupWorkstationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Indicates whether this resource is currently being updated to match its intended state.
func (o LookupWorkstationResultOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWorkstationResult) bool { return v.Reconciling }).(pulumi.BoolOutput)
}

// Current state of the workstation.
func (o LookupWorkstationResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationResult) string { return v.State }).(pulumi.StringOutput)
}

// A system-assigned unique identified for this resource.
func (o LookupWorkstationResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationResult) string { return v.Uid }).(pulumi.StringOutput)
}

// Time when this resource was most recently updated.
func (o LookupWorkstationResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkstationResultOutput{})
}
