// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the subscription configuration.
func LookupSubscription(ctx *pulumi.Context, args *LookupSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupSubscriptionResult, error) {
	var rv LookupSubscriptionResult
	err := ctx.Invoke("google-native:pubsublite/v1:getSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubscriptionArgs struct {
	Location       string  `pulumi:"location"`
	Project        *string `pulumi:"project"`
	SubscriptionId string  `pulumi:"subscriptionId"`
}

type LookupSubscriptionResult struct {
	// The settings for this subscription's message delivery.
	DeliveryConfig DeliveryConfigResponse `pulumi:"deliveryConfig"`
	// If present, messages are automatically written from the Pub/Sub Lite topic associated with this subscription to a destination.
	ExportConfig ExportConfigResponse `pulumi:"exportConfig"`
	// The name of the subscription. Structured like: projects/{project_number}/locations/{location}/subscriptions/{subscription_id}
	Name string `pulumi:"name"`
	// The name of the topic this subscription is attached to. Structured like: projects/{project_number}/locations/{location}/topics/{topic_id}
	Topic string `pulumi:"topic"`
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
	Location       pulumi.StringInput    `pulumi:"location"`
	Project        pulumi.StringPtrInput `pulumi:"project"`
	SubscriptionId pulumi.StringInput    `pulumi:"subscriptionId"`
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

// The settings for this subscription's message delivery.
func (o LookupSubscriptionResultOutput) DeliveryConfig() DeliveryConfigResponseOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) DeliveryConfigResponse { return v.DeliveryConfig }).(DeliveryConfigResponseOutput)
}

// If present, messages are automatically written from the Pub/Sub Lite topic associated with this subscription to a destination.
func (o LookupSubscriptionResultOutput) ExportConfig() ExportConfigResponseOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) ExportConfigResponse { return v.ExportConfig }).(ExportConfigResponseOutput)
}

// The name of the subscription. Structured like: projects/{project_number}/locations/{location}/subscriptions/{subscription_id}
func (o LookupSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The name of the topic this subscription is attached to. Structured like: projects/{project_number}/locations/{location}/topics/{topic_id}
func (o LookupSubscriptionResultOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Topic }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubscriptionResultOutput{})
}
