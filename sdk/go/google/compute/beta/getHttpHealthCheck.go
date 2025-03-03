// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified HttpHealthCheck resource.
func LookupHttpHealthCheck(ctx *pulumi.Context, args *LookupHttpHealthCheckArgs, opts ...pulumi.InvokeOption) (*LookupHttpHealthCheckResult, error) {
	var rv LookupHttpHealthCheckResult
	err := ctx.Invoke("google-native:compute/beta:getHttpHealthCheck", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHttpHealthCheckArgs struct {
	HttpHealthCheck string  `pulumi:"httpHealthCheck"`
	Project         *string `pulumi:"project"`
}

type LookupHttpHealthCheckResult struct {
	// How often (in seconds) to send a health check. The default value is 5 seconds.
	CheckIntervalSec int `pulumi:"checkIntervalSec"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
	HealthyThreshold int `pulumi:"healthyThreshold"`
	// The value of the host header in the HTTP health check request. If left empty (default value), the public IP on behalf of which this health check is performed will be used.
	Host string `pulumi:"host"`
	// Type of the resource. Always compute#httpHealthCheck for HTTP health checks.
	Kind string `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// The TCP port number for the HTTP health check request. The default value is 80.
	Port int `pulumi:"port"`
	// The request path of the HTTP health check request. The default value is /. This field does not support query parameters. Must comply with RFC3986.
	RequestPath string `pulumi:"requestPath"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.
	TimeoutSec int `pulumi:"timeoutSec"`
	// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
	UnhealthyThreshold int `pulumi:"unhealthyThreshold"`
}

func LookupHttpHealthCheckOutput(ctx *pulumi.Context, args LookupHttpHealthCheckOutputArgs, opts ...pulumi.InvokeOption) LookupHttpHealthCheckResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHttpHealthCheckResult, error) {
			args := v.(LookupHttpHealthCheckArgs)
			r, err := LookupHttpHealthCheck(ctx, &args, opts...)
			var s LookupHttpHealthCheckResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHttpHealthCheckResultOutput)
}

type LookupHttpHealthCheckOutputArgs struct {
	HttpHealthCheck pulumi.StringInput    `pulumi:"httpHealthCheck"`
	Project         pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupHttpHealthCheckOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHttpHealthCheckArgs)(nil)).Elem()
}

type LookupHttpHealthCheckResultOutput struct{ *pulumi.OutputState }

func (LookupHttpHealthCheckResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHttpHealthCheckResult)(nil)).Elem()
}

func (o LookupHttpHealthCheckResultOutput) ToLookupHttpHealthCheckResultOutput() LookupHttpHealthCheckResultOutput {
	return o
}

func (o LookupHttpHealthCheckResultOutput) ToLookupHttpHealthCheckResultOutputWithContext(ctx context.Context) LookupHttpHealthCheckResultOutput {
	return o
}

// How often (in seconds) to send a health check. The default value is 5 seconds.
func (o LookupHttpHealthCheckResultOutput) CheckIntervalSec() pulumi.IntOutput {
	return o.ApplyT(func(v LookupHttpHealthCheckResult) int { return v.CheckIntervalSec }).(pulumi.IntOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupHttpHealthCheckResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHttpHealthCheckResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupHttpHealthCheckResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHttpHealthCheckResult) string { return v.Description }).(pulumi.StringOutput)
}

// A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
func (o LookupHttpHealthCheckResultOutput) HealthyThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v LookupHttpHealthCheckResult) int { return v.HealthyThreshold }).(pulumi.IntOutput)
}

// The value of the host header in the HTTP health check request. If left empty (default value), the public IP on behalf of which this health check is performed will be used.
func (o LookupHttpHealthCheckResultOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHttpHealthCheckResult) string { return v.Host }).(pulumi.StringOutput)
}

// Type of the resource. Always compute#httpHealthCheck for HTTP health checks.
func (o LookupHttpHealthCheckResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHttpHealthCheckResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupHttpHealthCheckResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHttpHealthCheckResult) string { return v.Name }).(pulumi.StringOutput)
}

// The TCP port number for the HTTP health check request. The default value is 80.
func (o LookupHttpHealthCheckResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupHttpHealthCheckResult) int { return v.Port }).(pulumi.IntOutput)
}

// The request path of the HTTP health check request. The default value is /. This field does not support query parameters. Must comply with RFC3986.
func (o LookupHttpHealthCheckResultOutput) RequestPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHttpHealthCheckResult) string { return v.RequestPath }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupHttpHealthCheckResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHttpHealthCheckResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.
func (o LookupHttpHealthCheckResultOutput) TimeoutSec() pulumi.IntOutput {
	return o.ApplyT(func(v LookupHttpHealthCheckResult) int { return v.TimeoutSec }).(pulumi.IntOutput)
}

// A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
func (o LookupHttpHealthCheckResultOutput) UnhealthyThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v LookupHttpHealthCheckResult) int { return v.UnhealthyThreshold }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHttpHealthCheckResultOutput{})
}
