// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ScanConfigTargetPlatformsItem string

const (
	// The target platform is unknown. Requests with this enum value will be rejected with INVALID_ARGUMENT error.
	ScanConfigTargetPlatformsItemTargetPlatformUnspecified = ScanConfigTargetPlatformsItem("TARGET_PLATFORM_UNSPECIFIED")
	// Google App Engine service.
	ScanConfigTargetPlatformsItemAppEngine = ScanConfigTargetPlatformsItem("APP_ENGINE")
	// Google Compute Engine service.
	ScanConfigTargetPlatformsItemCompute = ScanConfigTargetPlatformsItem("COMPUTE")
	// Google Cloud Run service.
	ScanConfigTargetPlatformsItemCloudRun = ScanConfigTargetPlatformsItem("CLOUD_RUN")
	// Google Cloud Function service.
	ScanConfigTargetPlatformsItemCloudFunctions = ScanConfigTargetPlatformsItem("CLOUD_FUNCTIONS")
)

func (ScanConfigTargetPlatformsItem) ElementType() reflect.Type {
	return reflect.TypeOf((*ScanConfigTargetPlatformsItem)(nil)).Elem()
}

func (e ScanConfigTargetPlatformsItem) ToScanConfigTargetPlatformsItemOutput() ScanConfigTargetPlatformsItemOutput {
	return pulumi.ToOutput(e).(ScanConfigTargetPlatformsItemOutput)
}

func (e ScanConfigTargetPlatformsItem) ToScanConfigTargetPlatformsItemOutputWithContext(ctx context.Context) ScanConfigTargetPlatformsItemOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScanConfigTargetPlatformsItemOutput)
}

func (e ScanConfigTargetPlatformsItem) ToScanConfigTargetPlatformsItemPtrOutput() ScanConfigTargetPlatformsItemPtrOutput {
	return e.ToScanConfigTargetPlatformsItemPtrOutputWithContext(context.Background())
}

func (e ScanConfigTargetPlatformsItem) ToScanConfigTargetPlatformsItemPtrOutputWithContext(ctx context.Context) ScanConfigTargetPlatformsItemPtrOutput {
	return ScanConfigTargetPlatformsItem(e).ToScanConfigTargetPlatformsItemOutputWithContext(ctx).ToScanConfigTargetPlatformsItemPtrOutputWithContext(ctx)
}

func (e ScanConfigTargetPlatformsItem) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScanConfigTargetPlatformsItem) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScanConfigTargetPlatformsItem) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScanConfigTargetPlatformsItem) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScanConfigTargetPlatformsItemOutput struct{ *pulumi.OutputState }

func (ScanConfigTargetPlatformsItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScanConfigTargetPlatformsItem)(nil)).Elem()
}

func (o ScanConfigTargetPlatformsItemOutput) ToScanConfigTargetPlatformsItemOutput() ScanConfigTargetPlatformsItemOutput {
	return o
}

func (o ScanConfigTargetPlatformsItemOutput) ToScanConfigTargetPlatformsItemOutputWithContext(ctx context.Context) ScanConfigTargetPlatformsItemOutput {
	return o
}

func (o ScanConfigTargetPlatformsItemOutput) ToScanConfigTargetPlatformsItemPtrOutput() ScanConfigTargetPlatformsItemPtrOutput {
	return o.ToScanConfigTargetPlatformsItemPtrOutputWithContext(context.Background())
}

func (o ScanConfigTargetPlatformsItemOutput) ToScanConfigTargetPlatformsItemPtrOutputWithContext(ctx context.Context) ScanConfigTargetPlatformsItemPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScanConfigTargetPlatformsItem) *ScanConfigTargetPlatformsItem {
		return &v
	}).(ScanConfigTargetPlatformsItemPtrOutput)
}

func (o ScanConfigTargetPlatformsItemOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScanConfigTargetPlatformsItemOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScanConfigTargetPlatformsItem) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScanConfigTargetPlatformsItemOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScanConfigTargetPlatformsItemOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScanConfigTargetPlatformsItem) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScanConfigTargetPlatformsItemPtrOutput struct{ *pulumi.OutputState }

func (ScanConfigTargetPlatformsItemPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScanConfigTargetPlatformsItem)(nil)).Elem()
}

func (o ScanConfigTargetPlatformsItemPtrOutput) ToScanConfigTargetPlatformsItemPtrOutput() ScanConfigTargetPlatformsItemPtrOutput {
	return o
}

func (o ScanConfigTargetPlatformsItemPtrOutput) ToScanConfigTargetPlatformsItemPtrOutputWithContext(ctx context.Context) ScanConfigTargetPlatformsItemPtrOutput {
	return o
}

func (o ScanConfigTargetPlatformsItemPtrOutput) Elem() ScanConfigTargetPlatformsItemOutput {
	return o.ApplyT(func(v *ScanConfigTargetPlatformsItem) ScanConfigTargetPlatformsItem {
		if v != nil {
			return *v
		}
		var ret ScanConfigTargetPlatformsItem
		return ret
	}).(ScanConfigTargetPlatformsItemOutput)
}

func (o ScanConfigTargetPlatformsItemPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScanConfigTargetPlatformsItemPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScanConfigTargetPlatformsItem) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ScanConfigTargetPlatformsItemInput is an input type that accepts ScanConfigTargetPlatformsItemArgs and ScanConfigTargetPlatformsItemOutput values.
// You can construct a concrete instance of `ScanConfigTargetPlatformsItemInput` via:
//
//	ScanConfigTargetPlatformsItemArgs{...}
type ScanConfigTargetPlatformsItemInput interface {
	pulumi.Input

	ToScanConfigTargetPlatformsItemOutput() ScanConfigTargetPlatformsItemOutput
	ToScanConfigTargetPlatformsItemOutputWithContext(context.Context) ScanConfigTargetPlatformsItemOutput
}

var scanConfigTargetPlatformsItemPtrType = reflect.TypeOf((**ScanConfigTargetPlatformsItem)(nil)).Elem()

type ScanConfigTargetPlatformsItemPtrInput interface {
	pulumi.Input

	ToScanConfigTargetPlatformsItemPtrOutput() ScanConfigTargetPlatformsItemPtrOutput
	ToScanConfigTargetPlatformsItemPtrOutputWithContext(context.Context) ScanConfigTargetPlatformsItemPtrOutput
}

type scanConfigTargetPlatformsItemPtr string

func ScanConfigTargetPlatformsItemPtr(v string) ScanConfigTargetPlatformsItemPtrInput {
	return (*scanConfigTargetPlatformsItemPtr)(&v)
}

func (*scanConfigTargetPlatformsItemPtr) ElementType() reflect.Type {
	return scanConfigTargetPlatformsItemPtrType
}

func (in *scanConfigTargetPlatformsItemPtr) ToScanConfigTargetPlatformsItemPtrOutput() ScanConfigTargetPlatformsItemPtrOutput {
	return pulumi.ToOutput(in).(ScanConfigTargetPlatformsItemPtrOutput)
}

func (in *scanConfigTargetPlatformsItemPtr) ToScanConfigTargetPlatformsItemPtrOutputWithContext(ctx context.Context) ScanConfigTargetPlatformsItemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScanConfigTargetPlatformsItemPtrOutput)
}

// ScanConfigTargetPlatformsItemArrayInput is an input type that accepts ScanConfigTargetPlatformsItemArray and ScanConfigTargetPlatformsItemArrayOutput values.
// You can construct a concrete instance of `ScanConfigTargetPlatformsItemArrayInput` via:
//
//	ScanConfigTargetPlatformsItemArray{ ScanConfigTargetPlatformsItemArgs{...} }
type ScanConfigTargetPlatformsItemArrayInput interface {
	pulumi.Input

	ToScanConfigTargetPlatformsItemArrayOutput() ScanConfigTargetPlatformsItemArrayOutput
	ToScanConfigTargetPlatformsItemArrayOutputWithContext(context.Context) ScanConfigTargetPlatformsItemArrayOutput
}

type ScanConfigTargetPlatformsItemArray []ScanConfigTargetPlatformsItem

func (ScanConfigTargetPlatformsItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScanConfigTargetPlatformsItem)(nil)).Elem()
}

func (i ScanConfigTargetPlatformsItemArray) ToScanConfigTargetPlatformsItemArrayOutput() ScanConfigTargetPlatformsItemArrayOutput {
	return i.ToScanConfigTargetPlatformsItemArrayOutputWithContext(context.Background())
}

func (i ScanConfigTargetPlatformsItemArray) ToScanConfigTargetPlatformsItemArrayOutputWithContext(ctx context.Context) ScanConfigTargetPlatformsItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScanConfigTargetPlatformsItemArrayOutput)
}

type ScanConfigTargetPlatformsItemArrayOutput struct{ *pulumi.OutputState }

func (ScanConfigTargetPlatformsItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScanConfigTargetPlatformsItem)(nil)).Elem()
}

func (o ScanConfigTargetPlatformsItemArrayOutput) ToScanConfigTargetPlatformsItemArrayOutput() ScanConfigTargetPlatformsItemArrayOutput {
	return o
}

func (o ScanConfigTargetPlatformsItemArrayOutput) ToScanConfigTargetPlatformsItemArrayOutputWithContext(ctx context.Context) ScanConfigTargetPlatformsItemArrayOutput {
	return o
}

func (o ScanConfigTargetPlatformsItemArrayOutput) Index(i pulumi.IntInput) ScanConfigTargetPlatformsItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScanConfigTargetPlatformsItem {
		return vs[0].([]ScanConfigTargetPlatformsItem)[vs[1].(int)]
	}).(ScanConfigTargetPlatformsItemOutput)
}

// The user agent used during scanning.
type ScanConfigUserAgent string

const (
	// The user agent is unknown. Service will default to CHROME_LINUX.
	ScanConfigUserAgentUserAgentUnspecified = ScanConfigUserAgent("USER_AGENT_UNSPECIFIED")
	// Chrome on Linux. This is the service default if unspecified.
	ScanConfigUserAgentChromeLinux = ScanConfigUserAgent("CHROME_LINUX")
	// Chrome on Android.
	ScanConfigUserAgentChromeAndroid = ScanConfigUserAgent("CHROME_ANDROID")
	// Safari on IPhone.
	ScanConfigUserAgentSafariIphone = ScanConfigUserAgent("SAFARI_IPHONE")
)

func (ScanConfigUserAgent) ElementType() reflect.Type {
	return reflect.TypeOf((*ScanConfigUserAgent)(nil)).Elem()
}

func (e ScanConfigUserAgent) ToScanConfigUserAgentOutput() ScanConfigUserAgentOutput {
	return pulumi.ToOutput(e).(ScanConfigUserAgentOutput)
}

func (e ScanConfigUserAgent) ToScanConfigUserAgentOutputWithContext(ctx context.Context) ScanConfigUserAgentOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScanConfigUserAgentOutput)
}

func (e ScanConfigUserAgent) ToScanConfigUserAgentPtrOutput() ScanConfigUserAgentPtrOutput {
	return e.ToScanConfigUserAgentPtrOutputWithContext(context.Background())
}

func (e ScanConfigUserAgent) ToScanConfigUserAgentPtrOutputWithContext(ctx context.Context) ScanConfigUserAgentPtrOutput {
	return ScanConfigUserAgent(e).ToScanConfigUserAgentOutputWithContext(ctx).ToScanConfigUserAgentPtrOutputWithContext(ctx)
}

func (e ScanConfigUserAgent) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScanConfigUserAgent) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScanConfigUserAgent) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScanConfigUserAgent) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScanConfigUserAgentOutput struct{ *pulumi.OutputState }

func (ScanConfigUserAgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScanConfigUserAgent)(nil)).Elem()
}

func (o ScanConfigUserAgentOutput) ToScanConfigUserAgentOutput() ScanConfigUserAgentOutput {
	return o
}

func (o ScanConfigUserAgentOutput) ToScanConfigUserAgentOutputWithContext(ctx context.Context) ScanConfigUserAgentOutput {
	return o
}

func (o ScanConfigUserAgentOutput) ToScanConfigUserAgentPtrOutput() ScanConfigUserAgentPtrOutput {
	return o.ToScanConfigUserAgentPtrOutputWithContext(context.Background())
}

func (o ScanConfigUserAgentOutput) ToScanConfigUserAgentPtrOutputWithContext(ctx context.Context) ScanConfigUserAgentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScanConfigUserAgent) *ScanConfigUserAgent {
		return &v
	}).(ScanConfigUserAgentPtrOutput)
}

func (o ScanConfigUserAgentOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScanConfigUserAgentOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScanConfigUserAgent) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScanConfigUserAgentOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScanConfigUserAgentOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScanConfigUserAgent) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScanConfigUserAgentPtrOutput struct{ *pulumi.OutputState }

func (ScanConfigUserAgentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScanConfigUserAgent)(nil)).Elem()
}

func (o ScanConfigUserAgentPtrOutput) ToScanConfigUserAgentPtrOutput() ScanConfigUserAgentPtrOutput {
	return o
}

func (o ScanConfigUserAgentPtrOutput) ToScanConfigUserAgentPtrOutputWithContext(ctx context.Context) ScanConfigUserAgentPtrOutput {
	return o
}

func (o ScanConfigUserAgentPtrOutput) Elem() ScanConfigUserAgentOutput {
	return o.ApplyT(func(v *ScanConfigUserAgent) ScanConfigUserAgent {
		if v != nil {
			return *v
		}
		var ret ScanConfigUserAgent
		return ret
	}).(ScanConfigUserAgentOutput)
}

func (o ScanConfigUserAgentPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScanConfigUserAgentPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScanConfigUserAgent) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ScanConfigUserAgentInput is an input type that accepts ScanConfigUserAgentArgs and ScanConfigUserAgentOutput values.
// You can construct a concrete instance of `ScanConfigUserAgentInput` via:
//
//	ScanConfigUserAgentArgs{...}
type ScanConfigUserAgentInput interface {
	pulumi.Input

	ToScanConfigUserAgentOutput() ScanConfigUserAgentOutput
	ToScanConfigUserAgentOutputWithContext(context.Context) ScanConfigUserAgentOutput
}

var scanConfigUserAgentPtrType = reflect.TypeOf((**ScanConfigUserAgent)(nil)).Elem()

type ScanConfigUserAgentPtrInput interface {
	pulumi.Input

	ToScanConfigUserAgentPtrOutput() ScanConfigUserAgentPtrOutput
	ToScanConfigUserAgentPtrOutputWithContext(context.Context) ScanConfigUserAgentPtrOutput
}

type scanConfigUserAgentPtr string

func ScanConfigUserAgentPtr(v string) ScanConfigUserAgentPtrInput {
	return (*scanConfigUserAgentPtr)(&v)
}

func (*scanConfigUserAgentPtr) ElementType() reflect.Type {
	return scanConfigUserAgentPtrType
}

func (in *scanConfigUserAgentPtr) ToScanConfigUserAgentPtrOutput() ScanConfigUserAgentPtrOutput {
	return pulumi.ToOutput(in).(ScanConfigUserAgentPtrOutput)
}

func (in *scanConfigUserAgentPtr) ToScanConfigUserAgentPtrOutputWithContext(ctx context.Context) ScanConfigUserAgentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScanConfigUserAgentPtrOutput)
}

// The execution state of the ScanRun.
type ScanRunExecutionState string

const (
	// Represents an invalid state caused by internal server error. This value should never be returned.
	ScanRunExecutionStateExecutionStateUnspecified = ScanRunExecutionState("EXECUTION_STATE_UNSPECIFIED")
	// The scan is waiting in the queue.
	ScanRunExecutionStateQueued = ScanRunExecutionState("QUEUED")
	// The scan is in progress.
	ScanRunExecutionStateScanning = ScanRunExecutionState("SCANNING")
	// The scan is either finished or stopped by user.
	ScanRunExecutionStateFinished = ScanRunExecutionState("FINISHED")
)

func (ScanRunExecutionState) ElementType() reflect.Type {
	return reflect.TypeOf((*ScanRunExecutionState)(nil)).Elem()
}

func (e ScanRunExecutionState) ToScanRunExecutionStateOutput() ScanRunExecutionStateOutput {
	return pulumi.ToOutput(e).(ScanRunExecutionStateOutput)
}

func (e ScanRunExecutionState) ToScanRunExecutionStateOutputWithContext(ctx context.Context) ScanRunExecutionStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScanRunExecutionStateOutput)
}

func (e ScanRunExecutionState) ToScanRunExecutionStatePtrOutput() ScanRunExecutionStatePtrOutput {
	return e.ToScanRunExecutionStatePtrOutputWithContext(context.Background())
}

func (e ScanRunExecutionState) ToScanRunExecutionStatePtrOutputWithContext(ctx context.Context) ScanRunExecutionStatePtrOutput {
	return ScanRunExecutionState(e).ToScanRunExecutionStateOutputWithContext(ctx).ToScanRunExecutionStatePtrOutputWithContext(ctx)
}

func (e ScanRunExecutionState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScanRunExecutionState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScanRunExecutionState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScanRunExecutionState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScanRunExecutionStateOutput struct{ *pulumi.OutputState }

func (ScanRunExecutionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScanRunExecutionState)(nil)).Elem()
}

func (o ScanRunExecutionStateOutput) ToScanRunExecutionStateOutput() ScanRunExecutionStateOutput {
	return o
}

func (o ScanRunExecutionStateOutput) ToScanRunExecutionStateOutputWithContext(ctx context.Context) ScanRunExecutionStateOutput {
	return o
}

func (o ScanRunExecutionStateOutput) ToScanRunExecutionStatePtrOutput() ScanRunExecutionStatePtrOutput {
	return o.ToScanRunExecutionStatePtrOutputWithContext(context.Background())
}

func (o ScanRunExecutionStateOutput) ToScanRunExecutionStatePtrOutputWithContext(ctx context.Context) ScanRunExecutionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScanRunExecutionState) *ScanRunExecutionState {
		return &v
	}).(ScanRunExecutionStatePtrOutput)
}

func (o ScanRunExecutionStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScanRunExecutionStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScanRunExecutionState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScanRunExecutionStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScanRunExecutionStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScanRunExecutionState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScanRunExecutionStatePtrOutput struct{ *pulumi.OutputState }

func (ScanRunExecutionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScanRunExecutionState)(nil)).Elem()
}

func (o ScanRunExecutionStatePtrOutput) ToScanRunExecutionStatePtrOutput() ScanRunExecutionStatePtrOutput {
	return o
}

func (o ScanRunExecutionStatePtrOutput) ToScanRunExecutionStatePtrOutputWithContext(ctx context.Context) ScanRunExecutionStatePtrOutput {
	return o
}

func (o ScanRunExecutionStatePtrOutput) Elem() ScanRunExecutionStateOutput {
	return o.ApplyT(func(v *ScanRunExecutionState) ScanRunExecutionState {
		if v != nil {
			return *v
		}
		var ret ScanRunExecutionState
		return ret
	}).(ScanRunExecutionStateOutput)
}

func (o ScanRunExecutionStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScanRunExecutionStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScanRunExecutionState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ScanRunExecutionStateInput is an input type that accepts ScanRunExecutionStateArgs and ScanRunExecutionStateOutput values.
// You can construct a concrete instance of `ScanRunExecutionStateInput` via:
//
//	ScanRunExecutionStateArgs{...}
type ScanRunExecutionStateInput interface {
	pulumi.Input

	ToScanRunExecutionStateOutput() ScanRunExecutionStateOutput
	ToScanRunExecutionStateOutputWithContext(context.Context) ScanRunExecutionStateOutput
}

var scanRunExecutionStatePtrType = reflect.TypeOf((**ScanRunExecutionState)(nil)).Elem()

type ScanRunExecutionStatePtrInput interface {
	pulumi.Input

	ToScanRunExecutionStatePtrOutput() ScanRunExecutionStatePtrOutput
	ToScanRunExecutionStatePtrOutputWithContext(context.Context) ScanRunExecutionStatePtrOutput
}

type scanRunExecutionStatePtr string

func ScanRunExecutionStatePtr(v string) ScanRunExecutionStatePtrInput {
	return (*scanRunExecutionStatePtr)(&v)
}

func (*scanRunExecutionStatePtr) ElementType() reflect.Type {
	return scanRunExecutionStatePtrType
}

func (in *scanRunExecutionStatePtr) ToScanRunExecutionStatePtrOutput() ScanRunExecutionStatePtrOutput {
	return pulumi.ToOutput(in).(ScanRunExecutionStatePtrOutput)
}

func (in *scanRunExecutionStatePtr) ToScanRunExecutionStatePtrOutputWithContext(ctx context.Context) ScanRunExecutionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScanRunExecutionStatePtrOutput)
}

// The result state of the ScanRun. This field is only available after the execution state reaches "FINISHED".
type ScanRunResultState string

const (
	// Default value. This value is returned when the ScanRun is not yet finished.
	ScanRunResultStateResultStateUnspecified = ScanRunResultState("RESULT_STATE_UNSPECIFIED")
	// The scan finished without errors.
	ScanRunResultStateSuccess = ScanRunResultState("SUCCESS")
	// The scan finished with errors.
	ScanRunResultStateError = ScanRunResultState("ERROR")
	// The scan was terminated by user.
	ScanRunResultStateKilled = ScanRunResultState("KILLED")
)

func (ScanRunResultState) ElementType() reflect.Type {
	return reflect.TypeOf((*ScanRunResultState)(nil)).Elem()
}

func (e ScanRunResultState) ToScanRunResultStateOutput() ScanRunResultStateOutput {
	return pulumi.ToOutput(e).(ScanRunResultStateOutput)
}

func (e ScanRunResultState) ToScanRunResultStateOutputWithContext(ctx context.Context) ScanRunResultStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScanRunResultStateOutput)
}

func (e ScanRunResultState) ToScanRunResultStatePtrOutput() ScanRunResultStatePtrOutput {
	return e.ToScanRunResultStatePtrOutputWithContext(context.Background())
}

func (e ScanRunResultState) ToScanRunResultStatePtrOutputWithContext(ctx context.Context) ScanRunResultStatePtrOutput {
	return ScanRunResultState(e).ToScanRunResultStateOutputWithContext(ctx).ToScanRunResultStatePtrOutputWithContext(ctx)
}

func (e ScanRunResultState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScanRunResultState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScanRunResultState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScanRunResultState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScanRunResultStateOutput struct{ *pulumi.OutputState }

func (ScanRunResultStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScanRunResultState)(nil)).Elem()
}

func (o ScanRunResultStateOutput) ToScanRunResultStateOutput() ScanRunResultStateOutput {
	return o
}

func (o ScanRunResultStateOutput) ToScanRunResultStateOutputWithContext(ctx context.Context) ScanRunResultStateOutput {
	return o
}

func (o ScanRunResultStateOutput) ToScanRunResultStatePtrOutput() ScanRunResultStatePtrOutput {
	return o.ToScanRunResultStatePtrOutputWithContext(context.Background())
}

func (o ScanRunResultStateOutput) ToScanRunResultStatePtrOutputWithContext(ctx context.Context) ScanRunResultStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScanRunResultState) *ScanRunResultState {
		return &v
	}).(ScanRunResultStatePtrOutput)
}

func (o ScanRunResultStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScanRunResultStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScanRunResultState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScanRunResultStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScanRunResultStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScanRunResultState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScanRunResultStatePtrOutput struct{ *pulumi.OutputState }

func (ScanRunResultStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScanRunResultState)(nil)).Elem()
}

func (o ScanRunResultStatePtrOutput) ToScanRunResultStatePtrOutput() ScanRunResultStatePtrOutput {
	return o
}

func (o ScanRunResultStatePtrOutput) ToScanRunResultStatePtrOutputWithContext(ctx context.Context) ScanRunResultStatePtrOutput {
	return o
}

func (o ScanRunResultStatePtrOutput) Elem() ScanRunResultStateOutput {
	return o.ApplyT(func(v *ScanRunResultState) ScanRunResultState {
		if v != nil {
			return *v
		}
		var ret ScanRunResultState
		return ret
	}).(ScanRunResultStateOutput)
}

func (o ScanRunResultStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScanRunResultStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScanRunResultState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ScanRunResultStateInput is an input type that accepts ScanRunResultStateArgs and ScanRunResultStateOutput values.
// You can construct a concrete instance of `ScanRunResultStateInput` via:
//
//	ScanRunResultStateArgs{...}
type ScanRunResultStateInput interface {
	pulumi.Input

	ToScanRunResultStateOutput() ScanRunResultStateOutput
	ToScanRunResultStateOutputWithContext(context.Context) ScanRunResultStateOutput
}

var scanRunResultStatePtrType = reflect.TypeOf((**ScanRunResultState)(nil)).Elem()

type ScanRunResultStatePtrInput interface {
	pulumi.Input

	ToScanRunResultStatePtrOutput() ScanRunResultStatePtrOutput
	ToScanRunResultStatePtrOutputWithContext(context.Context) ScanRunResultStatePtrOutput
}

type scanRunResultStatePtr string

func ScanRunResultStatePtr(v string) ScanRunResultStatePtrInput {
	return (*scanRunResultStatePtr)(&v)
}

func (*scanRunResultStatePtr) ElementType() reflect.Type {
	return scanRunResultStatePtrType
}

func (in *scanRunResultStatePtr) ToScanRunResultStatePtrOutput() ScanRunResultStatePtrOutput {
	return pulumi.ToOutput(in).(ScanRunResultStatePtrOutput)
}

func (in *scanRunResultStatePtr) ToScanRunResultStatePtrOutputWithContext(ctx context.Context) ScanRunResultStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScanRunResultStatePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScanConfigTargetPlatformsItemInput)(nil)).Elem(), ScanConfigTargetPlatformsItem("TARGET_PLATFORM_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScanConfigTargetPlatformsItemPtrInput)(nil)).Elem(), ScanConfigTargetPlatformsItem("TARGET_PLATFORM_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScanConfigTargetPlatformsItemArrayInput)(nil)).Elem(), ScanConfigTargetPlatformsItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScanConfigUserAgentInput)(nil)).Elem(), ScanConfigUserAgent("USER_AGENT_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScanConfigUserAgentPtrInput)(nil)).Elem(), ScanConfigUserAgent("USER_AGENT_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScanRunExecutionStateInput)(nil)).Elem(), ScanRunExecutionState("EXECUTION_STATE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScanRunExecutionStatePtrInput)(nil)).Elem(), ScanRunExecutionState("EXECUTION_STATE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScanRunResultStateInput)(nil)).Elem(), ScanRunResultState("RESULT_STATE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ScanRunResultStatePtrInput)(nil)).Elem(), ScanRunResultState("RESULT_STATE_UNSPECIFIED"))
	pulumi.RegisterOutputType(ScanConfigTargetPlatformsItemOutput{})
	pulumi.RegisterOutputType(ScanConfigTargetPlatformsItemPtrOutput{})
	pulumi.RegisterOutputType(ScanConfigTargetPlatformsItemArrayOutput{})
	pulumi.RegisterOutputType(ScanConfigUserAgentOutput{})
	pulumi.RegisterOutputType(ScanConfigUserAgentPtrOutput{})
	pulumi.RegisterOutputType(ScanRunExecutionStateOutput{})
	pulumi.RegisterOutputType(ScanRunExecutionStatePtrOutput{})
	pulumi.RegisterOutputType(ScanRunResultStateOutput{})
	pulumi.RegisterOutputType(ScanRunResultStatePtrOutput{})
}
