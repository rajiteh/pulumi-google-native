// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a webhook in the specified agent.
type Webhook struct {
	pulumi.CustomResourceState

	// Indicates whether the webhook is disabled.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// Required. The human-readable name of the webhook, unique within the agent.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Configuration for a generic web service.
	GenericWebService GoogleCloudDialogflowCxV3WebhookGenericWebServiceResponseOutput `pulumi:"genericWebService"`
	// The unique identifier of the webhook. Required for the Webhooks.UpdateWebhook method. Webhooks.CreateWebhook populates the name automatically. Format: `projects//locations//agents//webhooks/`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Webhook execution timeout. Execution is considered failed if Dialogflow doesn't receive a response from webhook at the end of the timeout period. Defaults to 5 seconds, maximum allowed timeout is 30 seconds.
	Timeout pulumi.StringOutput `pulumi:"timeout"`
}

// NewWebhook registers a new resource with the given unique name, arguments, and options.
func NewWebhook(ctx *pulumi.Context,
	name string, args *WebhookArgs, opts ...pulumi.ResourceOption) (*Webhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentId == nil {
		return nil, errors.New("invalid value for required argument 'AgentId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Webhook
	err := ctx.RegisterResource("google-native:dialogflow/v3:Webhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebhook gets an existing Webhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebhookState, opts ...pulumi.ResourceOption) (*Webhook, error) {
	var resource Webhook
	err := ctx.ReadResource("google-native:dialogflow/v3:Webhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Webhook resources.
type webhookState struct {
	// Indicates whether the webhook is disabled.
	Disabled *bool `pulumi:"disabled"`
	// Required. The human-readable name of the webhook, unique within the agent.
	DisplayName *string `pulumi:"displayName"`
	// Configuration for a generic web service.
	GenericWebService *GoogleCloudDialogflowCxV3WebhookGenericWebServiceResponse `pulumi:"genericWebService"`
	// The unique identifier of the webhook. Required for the Webhooks.UpdateWebhook method. Webhooks.CreateWebhook populates the name automatically. Format: `projects//locations//agents//webhooks/`.
	Name *string `pulumi:"name"`
	// Webhook execution timeout. Execution is considered failed if Dialogflow doesn't receive a response from webhook at the end of the timeout period. Defaults to 5 seconds, maximum allowed timeout is 30 seconds.
	Timeout *string `pulumi:"timeout"`
}

type WebhookState struct {
	// Indicates whether the webhook is disabled.
	Disabled pulumi.BoolPtrInput
	// Required. The human-readable name of the webhook, unique within the agent.
	DisplayName pulumi.StringPtrInput
	// Configuration for a generic web service.
	GenericWebService GoogleCloudDialogflowCxV3WebhookGenericWebServiceResponsePtrInput
	// The unique identifier of the webhook. Required for the Webhooks.UpdateWebhook method. Webhooks.CreateWebhook populates the name automatically. Format: `projects//locations//agents//webhooks/`.
	Name pulumi.StringPtrInput
	// Webhook execution timeout. Execution is considered failed if Dialogflow doesn't receive a response from webhook at the end of the timeout period. Defaults to 5 seconds, maximum allowed timeout is 30 seconds.
	Timeout pulumi.StringPtrInput
}

func (WebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookState)(nil)).Elem()
}

type webhookArgs struct {
	AgentId string `pulumi:"agentId"`
	// Indicates whether the webhook is disabled.
	Disabled *bool `pulumi:"disabled"`
	// Required. The human-readable name of the webhook, unique within the agent.
	DisplayName *string `pulumi:"displayName"`
	// Configuration for a generic web service.
	GenericWebService *GoogleCloudDialogflowCxV3WebhookGenericWebService `pulumi:"genericWebService"`
	Location          string                                             `pulumi:"location"`
	// The unique identifier of the webhook. Required for the Webhooks.UpdateWebhook method. Webhooks.CreateWebhook populates the name automatically. Format: `projects//locations//agents//webhooks/`.
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	// Webhook execution timeout. Execution is considered failed if Dialogflow doesn't receive a response from webhook at the end of the timeout period. Defaults to 5 seconds, maximum allowed timeout is 30 seconds.
	Timeout *string `pulumi:"timeout"`
}

// The set of arguments for constructing a Webhook resource.
type WebhookArgs struct {
	AgentId pulumi.StringInput
	// Indicates whether the webhook is disabled.
	Disabled pulumi.BoolPtrInput
	// Required. The human-readable name of the webhook, unique within the agent.
	DisplayName pulumi.StringPtrInput
	// Configuration for a generic web service.
	GenericWebService GoogleCloudDialogflowCxV3WebhookGenericWebServicePtrInput
	Location          pulumi.StringInput
	// The unique identifier of the webhook. Required for the Webhooks.UpdateWebhook method. Webhooks.CreateWebhook populates the name automatically. Format: `projects//locations//agents//webhooks/`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	// Webhook execution timeout. Execution is considered failed if Dialogflow doesn't receive a response from webhook at the end of the timeout period. Defaults to 5 seconds, maximum allowed timeout is 30 seconds.
	Timeout pulumi.StringPtrInput
}

func (WebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webhookArgs)(nil)).Elem()
}

type WebhookInput interface {
	pulumi.Input

	ToWebhookOutput() WebhookOutput
	ToWebhookOutputWithContext(ctx context.Context) WebhookOutput
}

func (*Webhook) ElementType() reflect.Type {
	return reflect.TypeOf((*Webhook)(nil))
}

func (i *Webhook) ToWebhookOutput() WebhookOutput {
	return i.ToWebhookOutputWithContext(context.Background())
}

func (i *Webhook) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookOutput)
}

type WebhookOutput struct {
	*pulumi.OutputState
}

func (WebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Webhook)(nil))
}

func (o WebhookOutput) ToWebhookOutput() WebhookOutput {
	return o
}

func (o WebhookOutput) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebhookOutput{})
}