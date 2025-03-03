// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a sink that exports specified log entries to a destination. The export of newly-ingested log entries begins immediately, unless the sink's writer_identity is not permitted to write to the destination. A sink can export log entries only from the resource owning the sink.
type BillingAccountSink struct {
	pulumi.CustomResourceState

	// Optional. Options that affect sinks exporting data to BigQuery.
	BigqueryOptions  BigQueryOptionsResponseOutput `pulumi:"bigqueryOptions"`
	BillingAccountId pulumi.StringOutput           `pulumi:"billingAccountId"`
	// The creation timestamp of the sink.This field may not be present for older sinks.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. A description of this sink.The maximum length of the description is 8000 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// The export destination: "storage.googleapis.com/[GCS_BUCKET]" "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]" "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]" The sink's writer_identity, set when the sink is created, must have permission to write to the destination or else the log entries are not exported. For more information, see Exporting Logs with Sinks (https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
	Destination pulumi.StringOutput `pulumi:"destination"`
	// Optional. If set to true, then this sink is disabled and it does not export any log entries.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// Optional. Log entries that match any of these exclusion filters will not be exported.If a log entry is matched by both filter and one of exclusion_filters it will not be exported.
	Exclusions LogExclusionResponseArrayOutput `pulumi:"exclusions"`
	// Optional. An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries). The only exported log entries are those that are in the resource owning the sink and that match the filter.For example:logName="projects/[PROJECT_ID]/logs/[LOG_ID]" AND severity>=ERROR
	Filter pulumi.StringOutput `pulumi:"filter"`
	// Optional. This field applies only to sinks owned by organizations and folders. If the field is false, the default, only the logs owned by the sink's parent resource are available for export. If the field is true, then log entries from all the projects, folders, and billing accounts contained in the sink's parent resource are also available for export. Whether a particular log entry from the children is exported depends on the sink's filter expression.For example, if this field is true, then the filter resource.type=gce_instance would export all Compute Engine VM instance log entries from all projects in the sink's parent.To only export entries from certain child projects, filter on the project part of the log name:logName:("projects/test-project1/" OR "projects/test-project2/") AND resource.type=gce_instance
	IncludeChildren pulumi.BoolOutput `pulumi:"includeChildren"`
	// The client-assigned sink identifier, unique within the project.For example: "my-syslog-errors-to-pubsub". Sink identifiers are limited to 100 characters and can include only the following characters: upper and lower-case alphanumeric characters, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name pulumi.StringOutput `pulumi:"name"`
	// Deprecated. This field is unused.
	//
	// Deprecated: Deprecated. This field is unused.
	OutputVersionFormat pulumi.StringOutput `pulumi:"outputVersionFormat"`
	// Optional. Determines the kind of IAM identity returned as writer_identity in the new sink. If this value is omitted or set to false, and if the sink's parent is a project, then the value returned as writer_identity is the same group or service account used by Cloud Logging before the addition of writer identities to this API. The sink's destination must be in the same project as the sink itself.If this field is set to true, or if the sink is owned by a non-project resource such as an organization, then the value of writer_identity will be a unique service account used only for exports from the new sink. For more information, see writer_identity in LogSink.
	UniqueWriterIdentity pulumi.BoolPtrOutput `pulumi:"uniqueWriterIdentity"`
	// The last update timestamp of the sink.This field may not be present for older sinks.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// An IAM identity—a service account or group—under which Cloud Logging writes the exported log entries to the sink's destination. This field is either set by specifying custom_writer_identity or set automatically by sinks.create and sinks.update based on the value of unique_writer_identity in those methods.Until you grant this identity write-access to the destination, log entry exports from this sink will fail. For more information, see Granting Access for a Resource (https://cloud.google.com/iam/docs/granting-roles-to-service-accounts#granting_access_to_a_service_account_for_a_resource). Consult the destination service's documentation to determine the appropriate IAM roles to assign to the identity.Sinks that have a destination that is a log bucket in the same project as the sink cannot have a writer_identity and no additional permissions are required.
	WriterIdentity pulumi.StringOutput `pulumi:"writerIdentity"`
}

// NewBillingAccountSink registers a new resource with the given unique name, arguments, and options.
func NewBillingAccountSink(ctx *pulumi.Context,
	name string, args *BillingAccountSinkArgs, opts ...pulumi.ResourceOption) (*BillingAccountSink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountId == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountId'")
	}
	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"billingAccountId",
	})
	opts = append(opts, replaceOnChanges)
	var resource BillingAccountSink
	err := ctx.RegisterResource("google-native:logging/v2:BillingAccountSink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBillingAccountSink gets an existing BillingAccountSink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBillingAccountSink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BillingAccountSinkState, opts ...pulumi.ResourceOption) (*BillingAccountSink, error) {
	var resource BillingAccountSink
	err := ctx.ReadResource("google-native:logging/v2:BillingAccountSink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BillingAccountSink resources.
type billingAccountSinkState struct {
}

type BillingAccountSinkState struct {
}

func (BillingAccountSinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*billingAccountSinkState)(nil)).Elem()
}

type billingAccountSinkArgs struct {
	// Optional. Options that affect sinks exporting data to BigQuery.
	BigqueryOptions  *BigQueryOptions `pulumi:"bigqueryOptions"`
	BillingAccountId string           `pulumi:"billingAccountId"`
	// Optional. A description of this sink.The maximum length of the description is 8000 characters.
	Description *string `pulumi:"description"`
	// The export destination: "storage.googleapis.com/[GCS_BUCKET]" "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]" "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]" The sink's writer_identity, set when the sink is created, must have permission to write to the destination or else the log entries are not exported. For more information, see Exporting Logs with Sinks (https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
	Destination string `pulumi:"destination"`
	// Optional. If set to true, then this sink is disabled and it does not export any log entries.
	Disabled *bool `pulumi:"disabled"`
	// Optional. Log entries that match any of these exclusion filters will not be exported.If a log entry is matched by both filter and one of exclusion_filters it will not be exported.
	Exclusions []LogExclusion `pulumi:"exclusions"`
	// Optional. An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries). The only exported log entries are those that are in the resource owning the sink and that match the filter.For example:logName="projects/[PROJECT_ID]/logs/[LOG_ID]" AND severity>=ERROR
	Filter *string `pulumi:"filter"`
	// Optional. This field applies only to sinks owned by organizations and folders. If the field is false, the default, only the logs owned by the sink's parent resource are available for export. If the field is true, then log entries from all the projects, folders, and billing accounts contained in the sink's parent resource are also available for export. Whether a particular log entry from the children is exported depends on the sink's filter expression.For example, if this field is true, then the filter resource.type=gce_instance would export all Compute Engine VM instance log entries from all projects in the sink's parent.To only export entries from certain child projects, filter on the project part of the log name:logName:("projects/test-project1/" OR "projects/test-project2/") AND resource.type=gce_instance
	IncludeChildren *bool `pulumi:"includeChildren"`
	// The client-assigned sink identifier, unique within the project.For example: "my-syslog-errors-to-pubsub". Sink identifiers are limited to 100 characters and can include only the following characters: upper and lower-case alphanumeric characters, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name *string `pulumi:"name"`
	// Deprecated. This field is unused.
	//
	// Deprecated: Deprecated. This field is unused.
	OutputVersionFormat *BillingAccountSinkOutputVersionFormat `pulumi:"outputVersionFormat"`
	// Optional. Determines the kind of IAM identity returned as writer_identity in the new sink. If this value is omitted or set to false, and if the sink's parent is a project, then the value returned as writer_identity is the same group or service account used by Cloud Logging before the addition of writer identities to this API. The sink's destination must be in the same project as the sink itself.If this field is set to true, or if the sink is owned by a non-project resource such as an organization, then the value of writer_identity will be a unique service account used only for exports from the new sink. For more information, see writer_identity in LogSink.
	UniqueWriterIdentity *bool `pulumi:"uniqueWriterIdentity"`
}

// The set of arguments for constructing a BillingAccountSink resource.
type BillingAccountSinkArgs struct {
	// Optional. Options that affect sinks exporting data to BigQuery.
	BigqueryOptions  BigQueryOptionsPtrInput
	BillingAccountId pulumi.StringInput
	// Optional. A description of this sink.The maximum length of the description is 8000 characters.
	Description pulumi.StringPtrInput
	// The export destination: "storage.googleapis.com/[GCS_BUCKET]" "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]" "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]" The sink's writer_identity, set when the sink is created, must have permission to write to the destination or else the log entries are not exported. For more information, see Exporting Logs with Sinks (https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
	Destination pulumi.StringInput
	// Optional. If set to true, then this sink is disabled and it does not export any log entries.
	Disabled pulumi.BoolPtrInput
	// Optional. Log entries that match any of these exclusion filters will not be exported.If a log entry is matched by both filter and one of exclusion_filters it will not be exported.
	Exclusions LogExclusionArrayInput
	// Optional. An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries). The only exported log entries are those that are in the resource owning the sink and that match the filter.For example:logName="projects/[PROJECT_ID]/logs/[LOG_ID]" AND severity>=ERROR
	Filter pulumi.StringPtrInput
	// Optional. This field applies only to sinks owned by organizations and folders. If the field is false, the default, only the logs owned by the sink's parent resource are available for export. If the field is true, then log entries from all the projects, folders, and billing accounts contained in the sink's parent resource are also available for export. Whether a particular log entry from the children is exported depends on the sink's filter expression.For example, if this field is true, then the filter resource.type=gce_instance would export all Compute Engine VM instance log entries from all projects in the sink's parent.To only export entries from certain child projects, filter on the project part of the log name:logName:("projects/test-project1/" OR "projects/test-project2/") AND resource.type=gce_instance
	IncludeChildren pulumi.BoolPtrInput
	// The client-assigned sink identifier, unique within the project.For example: "my-syslog-errors-to-pubsub". Sink identifiers are limited to 100 characters and can include only the following characters: upper and lower-case alphanumeric characters, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name pulumi.StringPtrInput
	// Deprecated. This field is unused.
	//
	// Deprecated: Deprecated. This field is unused.
	OutputVersionFormat BillingAccountSinkOutputVersionFormatPtrInput
	// Optional. Determines the kind of IAM identity returned as writer_identity in the new sink. If this value is omitted or set to false, and if the sink's parent is a project, then the value returned as writer_identity is the same group or service account used by Cloud Logging before the addition of writer identities to this API. The sink's destination must be in the same project as the sink itself.If this field is set to true, or if the sink is owned by a non-project resource such as an organization, then the value of writer_identity will be a unique service account used only for exports from the new sink. For more information, see writer_identity in LogSink.
	UniqueWriterIdentity pulumi.BoolPtrInput
}

func (BillingAccountSinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*billingAccountSinkArgs)(nil)).Elem()
}

type BillingAccountSinkInput interface {
	pulumi.Input

	ToBillingAccountSinkOutput() BillingAccountSinkOutput
	ToBillingAccountSinkOutputWithContext(ctx context.Context) BillingAccountSinkOutput
}

func (*BillingAccountSink) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingAccountSink)(nil)).Elem()
}

func (i *BillingAccountSink) ToBillingAccountSinkOutput() BillingAccountSinkOutput {
	return i.ToBillingAccountSinkOutputWithContext(context.Background())
}

func (i *BillingAccountSink) ToBillingAccountSinkOutputWithContext(ctx context.Context) BillingAccountSinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingAccountSinkOutput)
}

type BillingAccountSinkOutput struct{ *pulumi.OutputState }

func (BillingAccountSinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingAccountSink)(nil)).Elem()
}

func (o BillingAccountSinkOutput) ToBillingAccountSinkOutput() BillingAccountSinkOutput {
	return o
}

func (o BillingAccountSinkOutput) ToBillingAccountSinkOutputWithContext(ctx context.Context) BillingAccountSinkOutput {
	return o
}

// Optional. Options that affect sinks exporting data to BigQuery.
func (o BillingAccountSinkOutput) BigqueryOptions() BigQueryOptionsResponseOutput {
	return o.ApplyT(func(v *BillingAccountSink) BigQueryOptionsResponseOutput { return v.BigqueryOptions }).(BigQueryOptionsResponseOutput)
}

func (o BillingAccountSinkOutput) BillingAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountSink) pulumi.StringOutput { return v.BillingAccountId }).(pulumi.StringOutput)
}

// The creation timestamp of the sink.This field may not be present for older sinks.
func (o BillingAccountSinkOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountSink) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. A description of this sink.The maximum length of the description is 8000 characters.
func (o BillingAccountSinkOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountSink) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The export destination: "storage.googleapis.com/[GCS_BUCKET]" "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]" "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]" The sink's writer_identity, set when the sink is created, must have permission to write to the destination or else the log entries are not exported. For more information, see Exporting Logs with Sinks (https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
func (o BillingAccountSinkOutput) Destination() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountSink) pulumi.StringOutput { return v.Destination }).(pulumi.StringOutput)
}

// Optional. If set to true, then this sink is disabled and it does not export any log entries.
func (o BillingAccountSinkOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *BillingAccountSink) pulumi.BoolOutput { return v.Disabled }).(pulumi.BoolOutput)
}

// Optional. Log entries that match any of these exclusion filters will not be exported.If a log entry is matched by both filter and one of exclusion_filters it will not be exported.
func (o BillingAccountSinkOutput) Exclusions() LogExclusionResponseArrayOutput {
	return o.ApplyT(func(v *BillingAccountSink) LogExclusionResponseArrayOutput { return v.Exclusions }).(LogExclusionResponseArrayOutput)
}

// Optional. An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries). The only exported log entries are those that are in the resource owning the sink and that match the filter.For example:logName="projects/[PROJECT_ID]/logs/[LOG_ID]" AND severity>=ERROR
func (o BillingAccountSinkOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountSink) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

// Optional. This field applies only to sinks owned by organizations and folders. If the field is false, the default, only the logs owned by the sink's parent resource are available for export. If the field is true, then log entries from all the projects, folders, and billing accounts contained in the sink's parent resource are also available for export. Whether a particular log entry from the children is exported depends on the sink's filter expression.For example, if this field is true, then the filter resource.type=gce_instance would export all Compute Engine VM instance log entries from all projects in the sink's parent.To only export entries from certain child projects, filter on the project part of the log name:logName:("projects/test-project1/" OR "projects/test-project2/") AND resource.type=gce_instance
func (o BillingAccountSinkOutput) IncludeChildren() pulumi.BoolOutput {
	return o.ApplyT(func(v *BillingAccountSink) pulumi.BoolOutput { return v.IncludeChildren }).(pulumi.BoolOutput)
}

// The client-assigned sink identifier, unique within the project.For example: "my-syslog-errors-to-pubsub". Sink identifiers are limited to 100 characters and can include only the following characters: upper and lower-case alphanumeric characters, underscores, hyphens, and periods. First character has to be alphanumeric.
func (o BillingAccountSinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountSink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Deprecated. This field is unused.
//
// Deprecated: Deprecated. This field is unused.
func (o BillingAccountSinkOutput) OutputVersionFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountSink) pulumi.StringOutput { return v.OutputVersionFormat }).(pulumi.StringOutput)
}

// Optional. Determines the kind of IAM identity returned as writer_identity in the new sink. If this value is omitted or set to false, and if the sink's parent is a project, then the value returned as writer_identity is the same group or service account used by Cloud Logging before the addition of writer identities to this API. The sink's destination must be in the same project as the sink itself.If this field is set to true, or if the sink is owned by a non-project resource such as an organization, then the value of writer_identity will be a unique service account used only for exports from the new sink. For more information, see writer_identity in LogSink.
func (o BillingAccountSinkOutput) UniqueWriterIdentity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BillingAccountSink) pulumi.BoolPtrOutput { return v.UniqueWriterIdentity }).(pulumi.BoolPtrOutput)
}

// The last update timestamp of the sink.This field may not be present for older sinks.
func (o BillingAccountSinkOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountSink) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// An IAM identity—a service account or group—under which Cloud Logging writes the exported log entries to the sink's destination. This field is either set by specifying custom_writer_identity or set automatically by sinks.create and sinks.update based on the value of unique_writer_identity in those methods.Until you grant this identity write-access to the destination, log entry exports from this sink will fail. For more information, see Granting Access for a Resource (https://cloud.google.com/iam/docs/granting-roles-to-service-accounts#granting_access_to_a_service_account_for_a_resource). Consult the destination service's documentation to determine the appropriate IAM roles to assign to the identity.Sinks that have a destination that is a log bucket in the same project as the sink cannot have a writer_identity and no additional permissions are required.
func (o BillingAccountSinkOutput) WriterIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingAccountSink) pulumi.StringOutput { return v.WriterIdentity }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BillingAccountSinkInput)(nil)).Elem(), &BillingAccountSink{})
	pulumi.RegisterOutputType(BillingAccountSinkOutput{})
}
