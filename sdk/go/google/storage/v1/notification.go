// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a notification subscription for a given bucket.
// Auto-naming is currently not supported for this resource.
type Notification struct {
	pulumi.CustomResourceState

	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// An optional list of additional attributes to attach to each Cloud PubSub message published for this notification subscription.
	CustomAttributes pulumi.StringMapOutput `pulumi:"customAttributes"`
	// HTTP 1.1 Entity tag for this subscription notification.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// If present, only send notifications about listed event types. If empty, sent notifications for all event types.
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// The kind of item this is. For notifications, this is always storage#notification.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// If present, only apply this notification configuration to object names that begin with this prefix.
	ObjectNamePrefix pulumi.StringOutput `pulumi:"objectNamePrefix"`
	// The desired content of the Payload.
	PayloadFormat pulumi.StringOutput `pulumi:"payloadFormat"`
	// The canonical URL of this notification.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The Cloud PubSub topic to which this subscription publishes. Formatted as: '//pubsub.googleapis.com/projects/{project-identifier}/topics/{my-topic}'
	Topic pulumi.StringOutput `pulumi:"topic"`
	// The project to be billed for this request. Required for Requester Pays buckets.
	UserProject pulumi.StringPtrOutput `pulumi:"userProject"`
}

// NewNotification registers a new resource with the given unique name, arguments, and options.
func NewNotification(ctx *pulumi.Context,
	name string, args *NotificationArgs, opts ...pulumi.ResourceOption) (*Notification, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"bucket",
	})
	opts = append(opts, replaceOnChanges)
	var resource Notification
	err := ctx.RegisterResource("google-native:storage/v1:Notification", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotification gets an existing Notification resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotification(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationState, opts ...pulumi.ResourceOption) (*Notification, error) {
	var resource Notification
	err := ctx.ReadResource("google-native:storage/v1:Notification", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Notification resources.
type notificationState struct {
}

type NotificationState struct {
}

func (NotificationState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationState)(nil)).Elem()
}

type notificationArgs struct {
	Bucket string `pulumi:"bucket"`
	// An optional list of additional attributes to attach to each Cloud PubSub message published for this notification subscription.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// HTTP 1.1 Entity tag for this subscription notification.
	Etag *string `pulumi:"etag"`
	// If present, only send notifications about listed event types. If empty, sent notifications for all event types.
	EventTypes []string `pulumi:"eventTypes"`
	// The ID of the notification.
	Id *string `pulumi:"id"`
	// The kind of item this is. For notifications, this is always storage#notification.
	Kind *string `pulumi:"kind"`
	// If present, only apply this notification configuration to object names that begin with this prefix.
	ObjectNamePrefix *string `pulumi:"objectNamePrefix"`
	// The desired content of the Payload.
	PayloadFormat *string `pulumi:"payloadFormat"`
	// The canonical URL of this notification.
	SelfLink *string `pulumi:"selfLink"`
	// The Cloud PubSub topic to which this subscription publishes. Formatted as: '//pubsub.googleapis.com/projects/{project-identifier}/topics/{my-topic}'
	Topic *string `pulumi:"topic"`
	// The project to be billed for this request. Required for Requester Pays buckets.
	UserProject *string `pulumi:"userProject"`
}

// The set of arguments for constructing a Notification resource.
type NotificationArgs struct {
	Bucket pulumi.StringInput
	// An optional list of additional attributes to attach to each Cloud PubSub message published for this notification subscription.
	CustomAttributes pulumi.StringMapInput
	// HTTP 1.1 Entity tag for this subscription notification.
	Etag pulumi.StringPtrInput
	// If present, only send notifications about listed event types. If empty, sent notifications for all event types.
	EventTypes pulumi.StringArrayInput
	// The ID of the notification.
	Id pulumi.StringPtrInput
	// The kind of item this is. For notifications, this is always storage#notification.
	Kind pulumi.StringPtrInput
	// If present, only apply this notification configuration to object names that begin with this prefix.
	ObjectNamePrefix pulumi.StringPtrInput
	// The desired content of the Payload.
	PayloadFormat pulumi.StringPtrInput
	// The canonical URL of this notification.
	SelfLink pulumi.StringPtrInput
	// The Cloud PubSub topic to which this subscription publishes. Formatted as: '//pubsub.googleapis.com/projects/{project-identifier}/topics/{my-topic}'
	Topic pulumi.StringPtrInput
	// The project to be billed for this request. Required for Requester Pays buckets.
	UserProject pulumi.StringPtrInput
}

func (NotificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationArgs)(nil)).Elem()
}

type NotificationInput interface {
	pulumi.Input

	ToNotificationOutput() NotificationOutput
	ToNotificationOutputWithContext(ctx context.Context) NotificationOutput
}

func (*Notification) ElementType() reflect.Type {
	return reflect.TypeOf((**Notification)(nil)).Elem()
}

func (i *Notification) ToNotificationOutput() NotificationOutput {
	return i.ToNotificationOutputWithContext(context.Background())
}

func (i *Notification) ToNotificationOutputWithContext(ctx context.Context) NotificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationOutput)
}

type NotificationOutput struct{ *pulumi.OutputState }

func (NotificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Notification)(nil)).Elem()
}

func (o NotificationOutput) ToNotificationOutput() NotificationOutput {
	return o
}

func (o NotificationOutput) ToNotificationOutputWithContext(ctx context.Context) NotificationOutput {
	return o
}

func (o NotificationOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// An optional list of additional attributes to attach to each Cloud PubSub message published for this notification subscription.
func (o NotificationOutput) CustomAttributes() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringMapOutput { return v.CustomAttributes }).(pulumi.StringMapOutput)
}

// HTTP 1.1 Entity tag for this subscription notification.
func (o NotificationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// If present, only send notifications about listed event types. If empty, sent notifications for all event types.
func (o NotificationOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringArrayOutput { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// The kind of item this is. For notifications, this is always storage#notification.
func (o NotificationOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// If present, only apply this notification configuration to object names that begin with this prefix.
func (o NotificationOutput) ObjectNamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringOutput { return v.ObjectNamePrefix }).(pulumi.StringOutput)
}

// The desired content of the Payload.
func (o NotificationOutput) PayloadFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringOutput { return v.PayloadFormat }).(pulumi.StringOutput)
}

// The canonical URL of this notification.
func (o NotificationOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The Cloud PubSub topic to which this subscription publishes. Formatted as: '//pubsub.googleapis.com/projects/{project-identifier}/topics/{my-topic}'
func (o NotificationOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringOutput { return v.Topic }).(pulumi.StringOutput)
}

// The project to be billed for this request. Required for Requester Pays buckets.
func (o NotificationOutput) UserProject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringPtrOutput { return v.UserProject }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationInput)(nil)).Elem(), &Notification{})
	pulumi.RegisterOutputType(NotificationOutput{})
}
