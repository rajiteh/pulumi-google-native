// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Auto scaling settings.
type Autoscale struct {
	// Number of slots to be scaled when needed.
	MaxSlots *string `pulumi:"maxSlots"`
}

// AutoscaleInput is an input type that accepts AutoscaleArgs and AutoscaleOutput values.
// You can construct a concrete instance of `AutoscaleInput` via:
//
//	AutoscaleArgs{...}
type AutoscaleInput interface {
	pulumi.Input

	ToAutoscaleOutput() AutoscaleOutput
	ToAutoscaleOutputWithContext(context.Context) AutoscaleOutput
}

// Auto scaling settings.
type AutoscaleArgs struct {
	// Number of slots to be scaled when needed.
	MaxSlots pulumi.StringPtrInput `pulumi:"maxSlots"`
}

func (AutoscaleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Autoscale)(nil)).Elem()
}

func (i AutoscaleArgs) ToAutoscaleOutput() AutoscaleOutput {
	return i.ToAutoscaleOutputWithContext(context.Background())
}

func (i AutoscaleArgs) ToAutoscaleOutputWithContext(ctx context.Context) AutoscaleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleOutput)
}

func (i AutoscaleArgs) ToAutoscalePtrOutput() AutoscalePtrOutput {
	return i.ToAutoscalePtrOutputWithContext(context.Background())
}

func (i AutoscaleArgs) ToAutoscalePtrOutputWithContext(ctx context.Context) AutoscalePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleOutput).ToAutoscalePtrOutputWithContext(ctx)
}

// AutoscalePtrInput is an input type that accepts AutoscaleArgs, AutoscalePtr and AutoscalePtrOutput values.
// You can construct a concrete instance of `AutoscalePtrInput` via:
//
//	        AutoscaleArgs{...}
//
//	or:
//
//	        nil
type AutoscalePtrInput interface {
	pulumi.Input

	ToAutoscalePtrOutput() AutoscalePtrOutput
	ToAutoscalePtrOutputWithContext(context.Context) AutoscalePtrOutput
}

type autoscalePtrType AutoscaleArgs

func AutoscalePtr(v *AutoscaleArgs) AutoscalePtrInput {
	return (*autoscalePtrType)(v)
}

func (*autoscalePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Autoscale)(nil)).Elem()
}

func (i *autoscalePtrType) ToAutoscalePtrOutput() AutoscalePtrOutput {
	return i.ToAutoscalePtrOutputWithContext(context.Background())
}

func (i *autoscalePtrType) ToAutoscalePtrOutputWithContext(ctx context.Context) AutoscalePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscalePtrOutput)
}

// Auto scaling settings.
type AutoscaleOutput struct{ *pulumi.OutputState }

func (AutoscaleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Autoscale)(nil)).Elem()
}

func (o AutoscaleOutput) ToAutoscaleOutput() AutoscaleOutput {
	return o
}

func (o AutoscaleOutput) ToAutoscaleOutputWithContext(ctx context.Context) AutoscaleOutput {
	return o
}

func (o AutoscaleOutput) ToAutoscalePtrOutput() AutoscalePtrOutput {
	return o.ToAutoscalePtrOutputWithContext(context.Background())
}

func (o AutoscaleOutput) ToAutoscalePtrOutputWithContext(ctx context.Context) AutoscalePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Autoscale) *Autoscale {
		return &v
	}).(AutoscalePtrOutput)
}

// Number of slots to be scaled when needed.
func (o AutoscaleOutput) MaxSlots() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Autoscale) *string { return v.MaxSlots }).(pulumi.StringPtrOutput)
}

type AutoscalePtrOutput struct{ *pulumi.OutputState }

func (AutoscalePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Autoscale)(nil)).Elem()
}

func (o AutoscalePtrOutput) ToAutoscalePtrOutput() AutoscalePtrOutput {
	return o
}

func (o AutoscalePtrOutput) ToAutoscalePtrOutputWithContext(ctx context.Context) AutoscalePtrOutput {
	return o
}

func (o AutoscalePtrOutput) Elem() AutoscaleOutput {
	return o.ApplyT(func(v *Autoscale) Autoscale {
		if v != nil {
			return *v
		}
		var ret Autoscale
		return ret
	}).(AutoscaleOutput)
}

// Number of slots to be scaled when needed.
func (o AutoscalePtrOutput) MaxSlots() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Autoscale) *string {
		if v == nil {
			return nil
		}
		return v.MaxSlots
	}).(pulumi.StringPtrOutput)
}

// Auto scaling settings.
type AutoscaleResponse struct {
	// The slot capacity added to this reservation when autoscale happens. Will be between [0, max_slots].
	CurrentSlots string `pulumi:"currentSlots"`
	// Number of slots to be scaled when needed.
	MaxSlots string `pulumi:"maxSlots"`
}

// Auto scaling settings.
type AutoscaleResponseOutput struct{ *pulumi.OutputState }

func (AutoscaleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleResponse)(nil)).Elem()
}

func (o AutoscaleResponseOutput) ToAutoscaleResponseOutput() AutoscaleResponseOutput {
	return o
}

func (o AutoscaleResponseOutput) ToAutoscaleResponseOutputWithContext(ctx context.Context) AutoscaleResponseOutput {
	return o
}

// The slot capacity added to this reservation when autoscale happens. Will be between [0, max_slots].
func (o AutoscaleResponseOutput) CurrentSlots() pulumi.StringOutput {
	return o.ApplyT(func(v AutoscaleResponse) string { return v.CurrentSlots }).(pulumi.StringOutput)
}

// Number of slots to be scaled when needed.
func (o AutoscaleResponseOutput) MaxSlots() pulumi.StringOutput {
	return o.ApplyT(func(v AutoscaleResponse) string { return v.MaxSlots }).(pulumi.StringOutput)
}

// The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
type StatusResponse struct {
	// The status code, which should be an enum value of google.rpc.Code.
	Code int `pulumi:"code"`
	// A list of messages that carry the error details. There is a common set of message types for APIs to use.
	Details []map[string]string `pulumi:"details"`
	// A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.
	Message string `pulumi:"message"`
}

// The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
type StatusResponseOutput struct{ *pulumi.OutputState }

func (StatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusResponse)(nil)).Elem()
}

func (o StatusResponseOutput) ToStatusResponseOutput() StatusResponseOutput {
	return o
}

func (o StatusResponseOutput) ToStatusResponseOutputWithContext(ctx context.Context) StatusResponseOutput {
	return o
}

// The status code, which should be an enum value of google.rpc.Code.
func (o StatusResponseOutput) Code() pulumi.IntOutput {
	return o.ApplyT(func(v StatusResponse) int { return v.Code }).(pulumi.IntOutput)
}

// A list of messages that carry the error details. There is a common set of message types for APIs to use.
func (o StatusResponseOutput) Details() pulumi.StringMapArrayOutput {
	return o.ApplyT(func(v StatusResponse) []map[string]string { return v.Details }).(pulumi.StringMapArrayOutput)
}

// A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the google.rpc.Status.details field, or localized by the client.
func (o StatusResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v StatusResponse) string { return v.Message }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutoscaleInput)(nil)).Elem(), AutoscaleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoscalePtrInput)(nil)).Elem(), AutoscaleArgs{})
	pulumi.RegisterOutputType(AutoscaleOutput{})
	pulumi.RegisterOutputType(AutoscalePtrOutput{})
	pulumi.RegisterOutputType(AutoscaleResponseOutput{})
	pulumi.RegisterOutputType(StatusResponseOutput{})
}
