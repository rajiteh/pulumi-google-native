// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an intent in the specified agent. Note: You should always train an agent prior to sending it queries. See the [training documentation](https://cloud.google.com/dialogflow/es/docs/training).
// Auto-naming is currently not supported for this resource.
type Intent struct {
	pulumi.CustomResourceState

	// Optional. The name of the action associated with the intent. Note: The action name must not contain whitespaces.
	Action pulumi.StringOutput `pulumi:"action"`
	// Optional. The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED (i.e. default platform).
	DefaultResponsePlatforms pulumi.StringArrayOutput `pulumi:"defaultResponsePlatforms"`
	// The name of this intent.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Optional. Indicates that this intent ends an interaction. Some integrations (e.g., Actions on Google or Dialogflow phone gateway) use this information to close interaction with an end user. Default is false.
	EndInteraction pulumi.BoolOutput `pulumi:"endInteraction"`
	// Optional. The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of the contexts must be present in the active user session for an event to trigger this intent. Event names are limited to 150 characters.
	Events pulumi.StringArrayOutput `pulumi:"events"`
	// Information about all followup intents that have this intent as a direct or indirect parent. We populate this field only in the output.
	FollowupIntentInfo GoogleCloudDialogflowV2beta1IntentFollowupIntentInfoResponseArrayOutput `pulumi:"followupIntentInfo"`
	// Optional. The list of context names required for this intent to be triggered. Formats: - `projects//agent/sessions/-/contexts/` - `projects//locations//agent/sessions/-/contexts/`
	InputContextNames pulumi.StringArrayOutput `pulumi:"inputContextNames"`
	// Optional. The resource view to apply to the returned intent.
	IntentView pulumi.StringPtrOutput `pulumi:"intentView"`
	// Optional. Indicates whether this is a fallback intent.
	IsFallback pulumi.BoolOutput `pulumi:"isFallback"`
	// Optional. The language used to access language-specific data. If not specified, the agent's default language is used. For more information, see [Multilingual intent and entity data](https://cloud.google.com/dialogflow/docs/agents-multilingual#intent-entity).
	LanguageCode pulumi.StringPtrOutput `pulumi:"languageCode"`
	// Optional. Indicates that a live agent should be brought in to handle the interaction with the user. In most cases, when you set this flag to true, you would also want to set end_interaction to true as well. Default is false.
	LiveAgentHandoff pulumi.BoolOutput   `pulumi:"liveAgentHandoff"`
	Location         pulumi.StringOutput `pulumi:"location"`
	// Optional. The collection of rich messages corresponding to the `Response` field in the Dialogflow console.
	Messages GoogleCloudDialogflowV2beta1IntentMessageResponseArrayOutput `pulumi:"messages"`
	// Optional. Indicates whether Machine Learning is disabled for the intent. Note: If `ml_disabled` setting is set to true, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off.
	MlDisabled pulumi.BoolOutput `pulumi:"mlDisabled"`
	// Optional. Indicates whether Machine Learning is enabled for the intent. Note: If `ml_enabled` setting is set to false, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off. DEPRECATED! Please use `ml_disabled` field instead. NOTE: If both `ml_enabled` and `ml_disabled` are either not set or false, then the default value is determined as follows: - Before April 15th, 2018 the default is: ml_enabled = false / ml_disabled = true. - After April 15th, 2018 the default is: ml_enabled = true / ml_disabled = false.
	//
	// Deprecated: Optional. Indicates whether Machine Learning is enabled for the intent. Note: If `ml_enabled` setting is set to false, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off. DEPRECATED! Please use `ml_disabled` field instead. NOTE: If both `ml_enabled` and `ml_disabled` are either not set or false, then the default value is determined as follows: - Before April 15th, 2018 the default is: ml_enabled = false / ml_disabled = true. - After April 15th, 2018 the default is: ml_enabled = true / ml_disabled = false.
	MlEnabled pulumi.BoolOutput `pulumi:"mlEnabled"`
	// Optional. The unique identifier of this intent. Required for Intents.UpdateIntent and Intents.BatchUpdateIntents methods. Supported formats: - `projects//agent/intents/` - `projects//locations//agent/intents/`
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. The collection of contexts that are activated when the intent is matched. Context messages in this collection should not set the parameters field. Setting the `lifespan_count` to 0 will reset the context when the intent is matched. Format: `projects//agent/sessions/-/contexts/`.
	OutputContexts GoogleCloudDialogflowV2beta1ContextResponseArrayOutput `pulumi:"outputContexts"`
	// Optional. The collection of parameters associated with the intent.
	Parameters GoogleCloudDialogflowV2beta1IntentParameterResponseArrayOutput `pulumi:"parameters"`
	// Optional. The unique identifier of the parent intent in the chain of followup intents. You can set this field when creating an intent, for example with CreateIntent or BatchUpdateIntents, in order to make this intent a followup intent. It identifies the parent followup intent. Format: `projects//agent/intents/`.
	ParentFollowupIntentName pulumi.StringOutput `pulumi:"parentFollowupIntentName"`
	// Optional. The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority pulumi.IntOutput    `pulumi:"priority"`
	Project  pulumi.StringOutput `pulumi:"project"`
	// Optional. Indicates whether to delete all contexts in the current session when this intent is matched.
	ResetContexts pulumi.BoolOutput `pulumi:"resetContexts"`
	// The unique identifier of the root intent in the chain of followup intents. It identifies the correct followup intents chain for this intent. Format: `projects//agent/intents/`.
	RootFollowupIntentName pulumi.StringOutput `pulumi:"rootFollowupIntentName"`
	// Optional. The collection of examples that the agent is trained on.
	TrainingPhrases GoogleCloudDialogflowV2beta1IntentTrainingPhraseResponseArrayOutput `pulumi:"trainingPhrases"`
	// Optional. Indicates whether webhooks are enabled for the intent.
	WebhookState pulumi.StringOutput `pulumi:"webhookState"`
}

// NewIntent registers a new resource with the given unique name, arguments, and options.
func NewIntent(ctx *pulumi.Context,
	name string, args *IntentArgs, opts ...pulumi.ResourceOption) (*Intent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource Intent
	err := ctx.RegisterResource("google-native:dialogflow/v2beta1:Intent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntent gets an existing Intent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntentState, opts ...pulumi.ResourceOption) (*Intent, error) {
	var resource Intent
	err := ctx.ReadResource("google-native:dialogflow/v2beta1:Intent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Intent resources.
type intentState struct {
}

type IntentState struct {
}

func (IntentState) ElementType() reflect.Type {
	return reflect.TypeOf((*intentState)(nil)).Elem()
}

type intentArgs struct {
	// Optional. The name of the action associated with the intent. Note: The action name must not contain whitespaces.
	Action *string `pulumi:"action"`
	// Optional. The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED (i.e. default platform).
	DefaultResponsePlatforms []IntentDefaultResponsePlatformsItem `pulumi:"defaultResponsePlatforms"`
	// The name of this intent.
	DisplayName string `pulumi:"displayName"`
	// Optional. Indicates that this intent ends an interaction. Some integrations (e.g., Actions on Google or Dialogflow phone gateway) use this information to close interaction with an end user. Default is false.
	EndInteraction *bool `pulumi:"endInteraction"`
	// Optional. The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of the contexts must be present in the active user session for an event to trigger this intent. Event names are limited to 150 characters.
	Events []string `pulumi:"events"`
	// Optional. The list of context names required for this intent to be triggered. Formats: - `projects//agent/sessions/-/contexts/` - `projects//locations//agent/sessions/-/contexts/`
	InputContextNames []string `pulumi:"inputContextNames"`
	// Optional. The resource view to apply to the returned intent.
	IntentView *string `pulumi:"intentView"`
	// Optional. Indicates whether this is a fallback intent.
	IsFallback *bool `pulumi:"isFallback"`
	// Optional. The language used to access language-specific data. If not specified, the agent's default language is used. For more information, see [Multilingual intent and entity data](https://cloud.google.com/dialogflow/docs/agents-multilingual#intent-entity).
	LanguageCode *string `pulumi:"languageCode"`
	// Optional. Indicates that a live agent should be brought in to handle the interaction with the user. In most cases, when you set this flag to true, you would also want to set end_interaction to true as well. Default is false.
	LiveAgentHandoff *bool   `pulumi:"liveAgentHandoff"`
	Location         *string `pulumi:"location"`
	// Optional. The collection of rich messages corresponding to the `Response` field in the Dialogflow console.
	Messages []GoogleCloudDialogflowV2beta1IntentMessage `pulumi:"messages"`
	// Optional. Indicates whether Machine Learning is disabled for the intent. Note: If `ml_disabled` setting is set to true, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off.
	MlDisabled *bool `pulumi:"mlDisabled"`
	// Optional. Indicates whether Machine Learning is enabled for the intent. Note: If `ml_enabled` setting is set to false, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off. DEPRECATED! Please use `ml_disabled` field instead. NOTE: If both `ml_enabled` and `ml_disabled` are either not set or false, then the default value is determined as follows: - Before April 15th, 2018 the default is: ml_enabled = false / ml_disabled = true. - After April 15th, 2018 the default is: ml_enabled = true / ml_disabled = false.
	//
	// Deprecated: Optional. Indicates whether Machine Learning is enabled for the intent. Note: If `ml_enabled` setting is set to false, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off. DEPRECATED! Please use `ml_disabled` field instead. NOTE: If both `ml_enabled` and `ml_disabled` are either not set or false, then the default value is determined as follows: - Before April 15th, 2018 the default is: ml_enabled = false / ml_disabled = true. - After April 15th, 2018 the default is: ml_enabled = true / ml_disabled = false.
	MlEnabled *bool `pulumi:"mlEnabled"`
	// Optional. The unique identifier of this intent. Required for Intents.UpdateIntent and Intents.BatchUpdateIntents methods. Supported formats: - `projects//agent/intents/` - `projects//locations//agent/intents/`
	Name *string `pulumi:"name"`
	// Optional. The collection of contexts that are activated when the intent is matched. Context messages in this collection should not set the parameters field. Setting the `lifespan_count` to 0 will reset the context when the intent is matched. Format: `projects//agent/sessions/-/contexts/`.
	OutputContexts []GoogleCloudDialogflowV2beta1Context `pulumi:"outputContexts"`
	// Optional. The collection of parameters associated with the intent.
	Parameters []GoogleCloudDialogflowV2beta1IntentParameter `pulumi:"parameters"`
	// Optional. The unique identifier of the parent intent in the chain of followup intents. You can set this field when creating an intent, for example with CreateIntent or BatchUpdateIntents, in order to make this intent a followup intent. It identifies the parent followup intent. Format: `projects//agent/intents/`.
	ParentFollowupIntentName *string `pulumi:"parentFollowupIntentName"`
	// Optional. The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority *int    `pulumi:"priority"`
	Project  *string `pulumi:"project"`
	// Optional. Indicates whether to delete all contexts in the current session when this intent is matched.
	ResetContexts *bool `pulumi:"resetContexts"`
	// Optional. The collection of examples that the agent is trained on.
	TrainingPhrases []GoogleCloudDialogflowV2beta1IntentTrainingPhrase `pulumi:"trainingPhrases"`
	// Optional. Indicates whether webhooks are enabled for the intent.
	WebhookState *IntentWebhookState `pulumi:"webhookState"`
}

// The set of arguments for constructing a Intent resource.
type IntentArgs struct {
	// Optional. The name of the action associated with the intent. Note: The action name must not contain whitespaces.
	Action pulumi.StringPtrInput
	// Optional. The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED (i.e. default platform).
	DefaultResponsePlatforms IntentDefaultResponsePlatformsItemArrayInput
	// The name of this intent.
	DisplayName pulumi.StringInput
	// Optional. Indicates that this intent ends an interaction. Some integrations (e.g., Actions on Google or Dialogflow phone gateway) use this information to close interaction with an end user. Default is false.
	EndInteraction pulumi.BoolPtrInput
	// Optional. The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of the contexts must be present in the active user session for an event to trigger this intent. Event names are limited to 150 characters.
	Events pulumi.StringArrayInput
	// Optional. The list of context names required for this intent to be triggered. Formats: - `projects//agent/sessions/-/contexts/` - `projects//locations//agent/sessions/-/contexts/`
	InputContextNames pulumi.StringArrayInput
	// Optional. The resource view to apply to the returned intent.
	IntentView pulumi.StringPtrInput
	// Optional. Indicates whether this is a fallback intent.
	IsFallback pulumi.BoolPtrInput
	// Optional. The language used to access language-specific data. If not specified, the agent's default language is used. For more information, see [Multilingual intent and entity data](https://cloud.google.com/dialogflow/docs/agents-multilingual#intent-entity).
	LanguageCode pulumi.StringPtrInput
	// Optional. Indicates that a live agent should be brought in to handle the interaction with the user. In most cases, when you set this flag to true, you would also want to set end_interaction to true as well. Default is false.
	LiveAgentHandoff pulumi.BoolPtrInput
	Location         pulumi.StringPtrInput
	// Optional. The collection of rich messages corresponding to the `Response` field in the Dialogflow console.
	Messages GoogleCloudDialogflowV2beta1IntentMessageArrayInput
	// Optional. Indicates whether Machine Learning is disabled for the intent. Note: If `ml_disabled` setting is set to true, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off.
	MlDisabled pulumi.BoolPtrInput
	// Optional. Indicates whether Machine Learning is enabled for the intent. Note: If `ml_enabled` setting is set to false, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off. DEPRECATED! Please use `ml_disabled` field instead. NOTE: If both `ml_enabled` and `ml_disabled` are either not set or false, then the default value is determined as follows: - Before April 15th, 2018 the default is: ml_enabled = false / ml_disabled = true. - After April 15th, 2018 the default is: ml_enabled = true / ml_disabled = false.
	//
	// Deprecated: Optional. Indicates whether Machine Learning is enabled for the intent. Note: If `ml_enabled` setting is set to false, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off. DEPRECATED! Please use `ml_disabled` field instead. NOTE: If both `ml_enabled` and `ml_disabled` are either not set or false, then the default value is determined as follows: - Before April 15th, 2018 the default is: ml_enabled = false / ml_disabled = true. - After April 15th, 2018 the default is: ml_enabled = true / ml_disabled = false.
	MlEnabled pulumi.BoolPtrInput
	// Optional. The unique identifier of this intent. Required for Intents.UpdateIntent and Intents.BatchUpdateIntents methods. Supported formats: - `projects//agent/intents/` - `projects//locations//agent/intents/`
	Name pulumi.StringPtrInput
	// Optional. The collection of contexts that are activated when the intent is matched. Context messages in this collection should not set the parameters field. Setting the `lifespan_count` to 0 will reset the context when the intent is matched. Format: `projects//agent/sessions/-/contexts/`.
	OutputContexts GoogleCloudDialogflowV2beta1ContextArrayInput
	// Optional. The collection of parameters associated with the intent.
	Parameters GoogleCloudDialogflowV2beta1IntentParameterArrayInput
	// Optional. The unique identifier of the parent intent in the chain of followup intents. You can set this field when creating an intent, for example with CreateIntent or BatchUpdateIntents, in order to make this intent a followup intent. It identifies the parent followup intent. Format: `projects//agent/intents/`.
	ParentFollowupIntentName pulumi.StringPtrInput
	// Optional. The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority pulumi.IntPtrInput
	Project  pulumi.StringPtrInput
	// Optional. Indicates whether to delete all contexts in the current session when this intent is matched.
	ResetContexts pulumi.BoolPtrInput
	// Optional. The collection of examples that the agent is trained on.
	TrainingPhrases GoogleCloudDialogflowV2beta1IntentTrainingPhraseArrayInput
	// Optional. Indicates whether webhooks are enabled for the intent.
	WebhookState IntentWebhookStatePtrInput
}

func (IntentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*intentArgs)(nil)).Elem()
}

type IntentInput interface {
	pulumi.Input

	ToIntentOutput() IntentOutput
	ToIntentOutputWithContext(ctx context.Context) IntentOutput
}

func (*Intent) ElementType() reflect.Type {
	return reflect.TypeOf((**Intent)(nil)).Elem()
}

func (i *Intent) ToIntentOutput() IntentOutput {
	return i.ToIntentOutputWithContext(context.Background())
}

func (i *Intent) ToIntentOutputWithContext(ctx context.Context) IntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntentOutput)
}

type IntentOutput struct{ *pulumi.OutputState }

func (IntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Intent)(nil)).Elem()
}

func (o IntentOutput) ToIntentOutput() IntentOutput {
	return o
}

func (o IntentOutput) ToIntentOutputWithContext(ctx context.Context) IntentOutput {
	return o
}

// Optional. The name of the action associated with the intent. Note: The action name must not contain whitespaces.
func (o IntentOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// Optional. The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED (i.e. default platform).
func (o IntentOutput) DefaultResponsePlatforms() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringArrayOutput { return v.DefaultResponsePlatforms }).(pulumi.StringArrayOutput)
}

// The name of this intent.
func (o IntentOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Optional. Indicates that this intent ends an interaction. Some integrations (e.g., Actions on Google or Dialogflow phone gateway) use this information to close interaction with an end user. Default is false.
func (o IntentOutput) EndInteraction() pulumi.BoolOutput {
	return o.ApplyT(func(v *Intent) pulumi.BoolOutput { return v.EndInteraction }).(pulumi.BoolOutput)
}

// Optional. The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of the contexts must be present in the active user session for an event to trigger this intent. Event names are limited to 150 characters.
func (o IntentOutput) Events() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringArrayOutput { return v.Events }).(pulumi.StringArrayOutput)
}

// Information about all followup intents that have this intent as a direct or indirect parent. We populate this field only in the output.
func (o IntentOutput) FollowupIntentInfo() GoogleCloudDialogflowV2beta1IntentFollowupIntentInfoResponseArrayOutput {
	return o.ApplyT(func(v *Intent) GoogleCloudDialogflowV2beta1IntentFollowupIntentInfoResponseArrayOutput {
		return v.FollowupIntentInfo
	}).(GoogleCloudDialogflowV2beta1IntentFollowupIntentInfoResponseArrayOutput)
}

// Optional. The list of context names required for this intent to be triggered. Formats: - `projects//agent/sessions/-/contexts/` - `projects//locations//agent/sessions/-/contexts/`
func (o IntentOutput) InputContextNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringArrayOutput { return v.InputContextNames }).(pulumi.StringArrayOutput)
}

// Optional. The resource view to apply to the returned intent.
func (o IntentOutput) IntentView() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringPtrOutput { return v.IntentView }).(pulumi.StringPtrOutput)
}

// Optional. Indicates whether this is a fallback intent.
func (o IntentOutput) IsFallback() pulumi.BoolOutput {
	return o.ApplyT(func(v *Intent) pulumi.BoolOutput { return v.IsFallback }).(pulumi.BoolOutput)
}

// Optional. The language used to access language-specific data. If not specified, the agent's default language is used. For more information, see [Multilingual intent and entity data](https://cloud.google.com/dialogflow/docs/agents-multilingual#intent-entity).
func (o IntentOutput) LanguageCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringPtrOutput { return v.LanguageCode }).(pulumi.StringPtrOutput)
}

// Optional. Indicates that a live agent should be brought in to handle the interaction with the user. In most cases, when you set this flag to true, you would also want to set end_interaction to true as well. Default is false.
func (o IntentOutput) LiveAgentHandoff() pulumi.BoolOutput {
	return o.ApplyT(func(v *Intent) pulumi.BoolOutput { return v.LiveAgentHandoff }).(pulumi.BoolOutput)
}

func (o IntentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Optional. The collection of rich messages corresponding to the `Response` field in the Dialogflow console.
func (o IntentOutput) Messages() GoogleCloudDialogflowV2beta1IntentMessageResponseArrayOutput {
	return o.ApplyT(func(v *Intent) GoogleCloudDialogflowV2beta1IntentMessageResponseArrayOutput { return v.Messages }).(GoogleCloudDialogflowV2beta1IntentMessageResponseArrayOutput)
}

// Optional. Indicates whether Machine Learning is disabled for the intent. Note: If `ml_disabled` setting is set to true, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off.
func (o IntentOutput) MlDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Intent) pulumi.BoolOutput { return v.MlDisabled }).(pulumi.BoolOutput)
}

// Optional. Indicates whether Machine Learning is enabled for the intent. Note: If `ml_enabled` setting is set to false, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off. DEPRECATED! Please use `ml_disabled` field instead. NOTE: If both `ml_enabled` and `ml_disabled` are either not set or false, then the default value is determined as follows: - Before April 15th, 2018 the default is: ml_enabled = false / ml_disabled = true. - After April 15th, 2018 the default is: ml_enabled = true / ml_disabled = false.
//
// Deprecated: Optional. Indicates whether Machine Learning is enabled for the intent. Note: If `ml_enabled` setting is set to false, then this intent is not taken into account during inference in `ML ONLY` match mode. Also, auto-markup in the UI is turned off. DEPRECATED! Please use `ml_disabled` field instead. NOTE: If both `ml_enabled` and `ml_disabled` are either not set or false, then the default value is determined as follows: - Before April 15th, 2018 the default is: ml_enabled = false / ml_disabled = true. - After April 15th, 2018 the default is: ml_enabled = true / ml_disabled = false.
func (o IntentOutput) MlEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Intent) pulumi.BoolOutput { return v.MlEnabled }).(pulumi.BoolOutput)
}

// Optional. The unique identifier of this intent. Required for Intents.UpdateIntent and Intents.BatchUpdateIntents methods. Supported formats: - `projects//agent/intents/` - `projects//locations//agent/intents/`
func (o IntentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional. The collection of contexts that are activated when the intent is matched. Context messages in this collection should not set the parameters field. Setting the `lifespan_count` to 0 will reset the context when the intent is matched. Format: `projects//agent/sessions/-/contexts/`.
func (o IntentOutput) OutputContexts() GoogleCloudDialogflowV2beta1ContextResponseArrayOutput {
	return o.ApplyT(func(v *Intent) GoogleCloudDialogflowV2beta1ContextResponseArrayOutput { return v.OutputContexts }).(GoogleCloudDialogflowV2beta1ContextResponseArrayOutput)
}

// Optional. The collection of parameters associated with the intent.
func (o IntentOutput) Parameters() GoogleCloudDialogflowV2beta1IntentParameterResponseArrayOutput {
	return o.ApplyT(func(v *Intent) GoogleCloudDialogflowV2beta1IntentParameterResponseArrayOutput { return v.Parameters }).(GoogleCloudDialogflowV2beta1IntentParameterResponseArrayOutput)
}

// Optional. The unique identifier of the parent intent in the chain of followup intents. You can set this field when creating an intent, for example with CreateIntent or BatchUpdateIntents, in order to make this intent a followup intent. It identifies the parent followup intent. Format: `projects//agent/intents/`.
func (o IntentOutput) ParentFollowupIntentName() pulumi.StringOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringOutput { return v.ParentFollowupIntentName }).(pulumi.StringOutput)
}

// Optional. The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
func (o IntentOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *Intent) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

func (o IntentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. Indicates whether to delete all contexts in the current session when this intent is matched.
func (o IntentOutput) ResetContexts() pulumi.BoolOutput {
	return o.ApplyT(func(v *Intent) pulumi.BoolOutput { return v.ResetContexts }).(pulumi.BoolOutput)
}

// The unique identifier of the root intent in the chain of followup intents. It identifies the correct followup intents chain for this intent. Format: `projects//agent/intents/`.
func (o IntentOutput) RootFollowupIntentName() pulumi.StringOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringOutput { return v.RootFollowupIntentName }).(pulumi.StringOutput)
}

// Optional. The collection of examples that the agent is trained on.
func (o IntentOutput) TrainingPhrases() GoogleCloudDialogflowV2beta1IntentTrainingPhraseResponseArrayOutput {
	return o.ApplyT(func(v *Intent) GoogleCloudDialogflowV2beta1IntentTrainingPhraseResponseArrayOutput {
		return v.TrainingPhrases
	}).(GoogleCloudDialogflowV2beta1IntentTrainingPhraseResponseArrayOutput)
}

// Optional. Indicates whether webhooks are enabled for the intent.
func (o IntentOutput) WebhookState() pulumi.StringOutput {
	return o.ApplyT(func(v *Intent) pulumi.StringOutput { return v.WebhookState }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntentInput)(nil)).Elem(), &Intent{})
	pulumi.RegisterOutputType(IntentOutput{})
}
