// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single Subscription.
func LookupSubscription(ctx *pulumi.Context, args *LookupSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupSubscriptionResult, error) {
	var rv LookupSubscriptionResult
	err := ctx.Invoke("google-native:beyondcorp/v1alpha:getSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubscriptionArgs struct {
	Location       string `pulumi:"location"`
	OrganizationId string `pulumi:"organizationId"`
	SubscriptionId string `pulumi:"subscriptionId"`
}

type LookupSubscriptionResult struct {
	// Represents that, if subscription will renew or end when the term ends.
	AutoRenewEnabled bool `pulumi:"autoRenewEnabled"`
	// Create time of the subscription.
	CreateTime string `pulumi:"createTime"`
	// End time of the subscription.
	EndTime string `pulumi:"endTime"`
	// Unique resource name of the Subscription. The name is ignored when creating a subscription.
	Name string `pulumi:"name"`
	// Number of seats in the subscription.
	SeatCount string `pulumi:"seatCount"`
	// SKU of subscription.
	Sku string `pulumi:"sku"`
	// Start time of the subscription.
	StartTime string `pulumi:"startTime"`
	// The current state of the subscription.
	State string `pulumi:"state"`
	// Type of subscription.
	Type string `pulumi:"type"`
}

func LookupSubscriptionOutput(ctx *pulumi.Context, args LookupSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubscriptionResult, error) {
			args := v.(LookupSubscriptionArgs)
			r, err := LookupSubscription(ctx, &args, opts...)
			var s LookupSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubscriptionResultOutput)
}

type LookupSubscriptionOutputArgs struct {
	Location       pulumi.StringInput `pulumi:"location"`
	OrganizationId pulumi.StringInput `pulumi:"organizationId"`
	SubscriptionId pulumi.StringInput `pulumi:"subscriptionId"`
}

func (LookupSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionArgs)(nil)).Elem()
}

type LookupSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionResult)(nil)).Elem()
}

func (o LookupSubscriptionResultOutput) ToLookupSubscriptionResultOutput() LookupSubscriptionResultOutput {
	return o
}

func (o LookupSubscriptionResultOutput) ToLookupSubscriptionResultOutputWithContext(ctx context.Context) LookupSubscriptionResultOutput {
	return o
}

// Represents that, if subscription will renew or end when the term ends.
func (o LookupSubscriptionResultOutput) AutoRenewEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) bool { return v.AutoRenewEnabled }).(pulumi.BoolOutput)
}

// Create time of the subscription.
func (o LookupSubscriptionResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// End time of the subscription.
func (o LookupSubscriptionResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// Unique resource name of the Subscription. The name is ignored when creating a subscription.
func (o LookupSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Number of seats in the subscription.
func (o LookupSubscriptionResultOutput) SeatCount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.SeatCount }).(pulumi.StringOutput)
}

// SKU of subscription.
func (o LookupSubscriptionResultOutput) Sku() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Sku }).(pulumi.StringOutput)
}

// Start time of the subscription.
func (o LookupSubscriptionResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// The current state of the subscription.
func (o LookupSubscriptionResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.State }).(pulumi.StringOutput)
}

// Type of subscription.
func (o LookupSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubscriptionResultOutput{})
}
