// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a job trigger to run DLP actions such as scanning storage for sensitive information on a set schedule. See https://cloud.google.com/dlp/docs/creating-job-triggers to learn more.
type JobTrigger struct {
	pulumi.CustomResourceState

	// The creation timestamp of a triggeredJob.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// User provided description (max 256 chars)
	Description pulumi.StringOutput `pulumi:"description"`
	// Display name (max 100 chars)
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// A stream of errors encountered when the trigger was activated. Repeated errors may result in the JobTrigger automatically being paused. Will return the last 100 errors. Whenever the JobTrigger is modified this list will be cleared.
	Errors GooglePrivacyDlpV2ErrorResponseArrayOutput `pulumi:"errors"`
	// For inspect jobs, a snapshot of the configuration.
	InspectJob GooglePrivacyDlpV2InspectJobConfigResponseOutput `pulumi:"inspectJob"`
	// The timestamp of the last time this trigger executed.
	LastRunTime pulumi.StringOutput `pulumi:"lastRunTime"`
	Location    pulumi.StringOutput `pulumi:"location"`
	// Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// A status for this trigger.
	Status pulumi.StringOutput `pulumi:"status"`
	// A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
	Triggers GooglePrivacyDlpV2TriggerResponseArrayOutput `pulumi:"triggers"`
	// The last update timestamp of a triggeredJob.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewJobTrigger registers a new resource with the given unique name, arguments, and options.
func NewJobTrigger(ctx *pulumi.Context,
	name string, args *JobTriggerArgs, opts ...pulumi.ResourceOption) (*JobTrigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Status == nil {
		return nil, errors.New("invalid value for required argument 'Status'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource JobTrigger
	err := ctx.RegisterResource("google-native:dlp/v2:JobTrigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobTrigger gets an existing JobTrigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobTriggerState, opts ...pulumi.ResourceOption) (*JobTrigger, error) {
	var resource JobTrigger
	err := ctx.ReadResource("google-native:dlp/v2:JobTrigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobTrigger resources.
type jobTriggerState struct {
}

type JobTriggerState struct {
}

func (JobTriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobTriggerState)(nil)).Elem()
}

type jobTriggerArgs struct {
	// User provided description (max 256 chars)
	Description *string `pulumi:"description"`
	// Display name (max 100 chars)
	DisplayName *string `pulumi:"displayName"`
	// For inspect jobs, a snapshot of the configuration.
	InspectJob *GooglePrivacyDlpV2InspectJobConfig `pulumi:"inspectJob"`
	// Deprecated. This field has no effect.
	//
	// Deprecated: Deprecated. This field has no effect.
	Location *string `pulumi:"location"`
	// Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// A status for this trigger.
	Status JobTriggerStatus `pulumi:"status"`
	// The trigger id can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
	TriggerId *string `pulumi:"triggerId"`
	// A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
	Triggers []GooglePrivacyDlpV2Trigger `pulumi:"triggers"`
}

// The set of arguments for constructing a JobTrigger resource.
type JobTriggerArgs struct {
	// User provided description (max 256 chars)
	Description pulumi.StringPtrInput
	// Display name (max 100 chars)
	DisplayName pulumi.StringPtrInput
	// For inspect jobs, a snapshot of the configuration.
	InspectJob GooglePrivacyDlpV2InspectJobConfigPtrInput
	// Deprecated. This field has no effect.
	//
	// Deprecated: Deprecated. This field has no effect.
	Location pulumi.StringPtrInput
	// Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// A status for this trigger.
	Status JobTriggerStatusInput
	// The trigger id can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
	TriggerId pulumi.StringPtrInput
	// A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
	Triggers GooglePrivacyDlpV2TriggerArrayInput
}

func (JobTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobTriggerArgs)(nil)).Elem()
}

type JobTriggerInput interface {
	pulumi.Input

	ToJobTriggerOutput() JobTriggerOutput
	ToJobTriggerOutputWithContext(ctx context.Context) JobTriggerOutput
}

func (*JobTrigger) ElementType() reflect.Type {
	return reflect.TypeOf((**JobTrigger)(nil)).Elem()
}

func (i *JobTrigger) ToJobTriggerOutput() JobTriggerOutput {
	return i.ToJobTriggerOutputWithContext(context.Background())
}

func (i *JobTrigger) ToJobTriggerOutputWithContext(ctx context.Context) JobTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTriggerOutput)
}

type JobTriggerOutput struct{ *pulumi.OutputState }

func (JobTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobTrigger)(nil)).Elem()
}

func (o JobTriggerOutput) ToJobTriggerOutput() JobTriggerOutput {
	return o
}

func (o JobTriggerOutput) ToJobTriggerOutputWithContext(ctx context.Context) JobTriggerOutput {
	return o
}

// The creation timestamp of a triggeredJob.
func (o JobTriggerOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *JobTrigger) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// User provided description (max 256 chars)
func (o JobTriggerOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *JobTrigger) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Display name (max 100 chars)
func (o JobTriggerOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *JobTrigger) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// A stream of errors encountered when the trigger was activated. Repeated errors may result in the JobTrigger automatically being paused. Will return the last 100 errors. Whenever the JobTrigger is modified this list will be cleared.
func (o JobTriggerOutput) Errors() GooglePrivacyDlpV2ErrorResponseArrayOutput {
	return o.ApplyT(func(v *JobTrigger) GooglePrivacyDlpV2ErrorResponseArrayOutput { return v.Errors }).(GooglePrivacyDlpV2ErrorResponseArrayOutput)
}

// For inspect jobs, a snapshot of the configuration.
func (o JobTriggerOutput) InspectJob() GooglePrivacyDlpV2InspectJobConfigResponseOutput {
	return o.ApplyT(func(v *JobTrigger) GooglePrivacyDlpV2InspectJobConfigResponseOutput { return v.InspectJob }).(GooglePrivacyDlpV2InspectJobConfigResponseOutput)
}

// The timestamp of the last time this trigger executed.
func (o JobTriggerOutput) LastRunTime() pulumi.StringOutput {
	return o.ApplyT(func(v *JobTrigger) pulumi.StringOutput { return v.LastRunTime }).(pulumi.StringOutput)
}

func (o JobTriggerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *JobTrigger) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
func (o JobTriggerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *JobTrigger) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o JobTriggerOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *JobTrigger) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A status for this trigger.
func (o JobTriggerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *JobTrigger) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
func (o JobTriggerOutput) Triggers() GooglePrivacyDlpV2TriggerResponseArrayOutput {
	return o.ApplyT(func(v *JobTrigger) GooglePrivacyDlpV2TriggerResponseArrayOutput { return v.Triggers }).(GooglePrivacyDlpV2TriggerResponseArrayOutput)
}

// The last update timestamp of a triggeredJob.
func (o JobTriggerOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *JobTrigger) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobTriggerInput)(nil)).Elem(), &JobTrigger{})
	pulumi.RegisterOutputType(JobTriggerOutput{})
}
