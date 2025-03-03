// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new HL7v2 store within the parent dataset.
type Hl7V2Store struct {
	pulumi.CustomResourceState

	DatasetId pulumi.StringOutput `pulumi:"datasetId"`
	// The ID of the HL7v2 store that is being created. The string must match the following regex: `[\p{L}\p{N}_\-\.]{1,256}`.
	Hl7V2StoreId pulumi.StringPtrOutput `pulumi:"hl7V2StoreId"`
	// User-supplied key-value pairs used to organize HL7v2 stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// Resource name of the HL7v2 store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/hl7V2Stores/{hl7v2_store_id}`.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of notification configs. Each configuration uses a filter to determine whether to publish a message (both Ingest & Create) on the corresponding notification destination. Only the message name is sent as part of the notification. Supplied by the client.
	NotificationConfigs Hl7V2NotificationConfigResponseArrayOutput `pulumi:"notificationConfigs"`
	// The configuration for the parser. It determines how the server parses the messages.
	ParserConfig ParserConfigResponseOutput `pulumi:"parserConfig"`
	Project      pulumi.StringOutput        `pulumi:"project"`
	// Determines whether to reject duplicate messages. A duplicate message is a message with the same raw bytes as a message that has already been ingested/created in this HL7v2 store. The default value is false, meaning that the store accepts the duplicate messages and it also returns the same ACK message in the IngestMessageResponse as has been returned previously. Note that only one resource is created in the store. When this field is set to true, CreateMessage/IngestMessage requests with a duplicate message will be rejected by the store, and IngestMessageErrorDetail returns a NACK message upon rejection.
	RejectDuplicateMessage pulumi.BoolOutput `pulumi:"rejectDuplicateMessage"`
}

// NewHl7V2Store registers a new resource with the given unique name, arguments, and options.
func NewHl7V2Store(ctx *pulumi.Context,
	name string, args *Hl7V2StoreArgs, opts ...pulumi.ResourceOption) (*Hl7V2Store, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"datasetId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource Hl7V2Store
	err := ctx.RegisterResource("google-native:healthcare/v1:Hl7V2Store", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHl7V2Store gets an existing Hl7V2Store resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHl7V2Store(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Hl7V2StoreState, opts ...pulumi.ResourceOption) (*Hl7V2Store, error) {
	var resource Hl7V2Store
	err := ctx.ReadResource("google-native:healthcare/v1:Hl7V2Store", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Hl7V2Store resources.
type hl7V2StoreState struct {
}

type Hl7V2StoreState struct {
}

func (Hl7V2StoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*hl7V2StoreState)(nil)).Elem()
}

type hl7V2StoreArgs struct {
	DatasetId string `pulumi:"datasetId"`
	// The ID of the HL7v2 store that is being created. The string must match the following regex: `[\p{L}\p{N}_\-\.]{1,256}`.
	Hl7V2StoreId *string `pulumi:"hl7V2StoreId"`
	// User-supplied key-value pairs used to organize HL7v2 stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Resource name of the HL7v2 store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/hl7V2Stores/{hl7v2_store_id}`.
	Name *string `pulumi:"name"`
	// A list of notification configs. Each configuration uses a filter to determine whether to publish a message (both Ingest & Create) on the corresponding notification destination. Only the message name is sent as part of the notification. Supplied by the client.
	NotificationConfigs []Hl7V2NotificationConfig `pulumi:"notificationConfigs"`
	// The configuration for the parser. It determines how the server parses the messages.
	ParserConfig *ParserConfig `pulumi:"parserConfig"`
	Project      *string       `pulumi:"project"`
	// Determines whether to reject duplicate messages. A duplicate message is a message with the same raw bytes as a message that has already been ingested/created in this HL7v2 store. The default value is false, meaning that the store accepts the duplicate messages and it also returns the same ACK message in the IngestMessageResponse as has been returned previously. Note that only one resource is created in the store. When this field is set to true, CreateMessage/IngestMessage requests with a duplicate message will be rejected by the store, and IngestMessageErrorDetail returns a NACK message upon rejection.
	RejectDuplicateMessage *bool `pulumi:"rejectDuplicateMessage"`
}

// The set of arguments for constructing a Hl7V2Store resource.
type Hl7V2StoreArgs struct {
	DatasetId pulumi.StringInput
	// The ID of the HL7v2 store that is being created. The string must match the following regex: `[\p{L}\p{N}_\-\.]{1,256}`.
	Hl7V2StoreId pulumi.StringPtrInput
	// User-supplied key-value pairs used to organize HL7v2 stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Resource name of the HL7v2 store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/hl7V2Stores/{hl7v2_store_id}`.
	Name pulumi.StringPtrInput
	// A list of notification configs. Each configuration uses a filter to determine whether to publish a message (both Ingest & Create) on the corresponding notification destination. Only the message name is sent as part of the notification. Supplied by the client.
	NotificationConfigs Hl7V2NotificationConfigArrayInput
	// The configuration for the parser. It determines how the server parses the messages.
	ParserConfig ParserConfigPtrInput
	Project      pulumi.StringPtrInput
	// Determines whether to reject duplicate messages. A duplicate message is a message with the same raw bytes as a message that has already been ingested/created in this HL7v2 store. The default value is false, meaning that the store accepts the duplicate messages and it also returns the same ACK message in the IngestMessageResponse as has been returned previously. Note that only one resource is created in the store. When this field is set to true, CreateMessage/IngestMessage requests with a duplicate message will be rejected by the store, and IngestMessageErrorDetail returns a NACK message upon rejection.
	RejectDuplicateMessage pulumi.BoolPtrInput
}

func (Hl7V2StoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hl7V2StoreArgs)(nil)).Elem()
}

type Hl7V2StoreInput interface {
	pulumi.Input

	ToHl7V2StoreOutput() Hl7V2StoreOutput
	ToHl7V2StoreOutputWithContext(ctx context.Context) Hl7V2StoreOutput
}

func (*Hl7V2Store) ElementType() reflect.Type {
	return reflect.TypeOf((**Hl7V2Store)(nil)).Elem()
}

func (i *Hl7V2Store) ToHl7V2StoreOutput() Hl7V2StoreOutput {
	return i.ToHl7V2StoreOutputWithContext(context.Background())
}

func (i *Hl7V2Store) ToHl7V2StoreOutputWithContext(ctx context.Context) Hl7V2StoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Hl7V2StoreOutput)
}

type Hl7V2StoreOutput struct{ *pulumi.OutputState }

func (Hl7V2StoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Hl7V2Store)(nil)).Elem()
}

func (o Hl7V2StoreOutput) ToHl7V2StoreOutput() Hl7V2StoreOutput {
	return o
}

func (o Hl7V2StoreOutput) ToHl7V2StoreOutputWithContext(ctx context.Context) Hl7V2StoreOutput {
	return o
}

func (o Hl7V2StoreOutput) DatasetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Hl7V2Store) pulumi.StringOutput { return v.DatasetId }).(pulumi.StringOutput)
}

// The ID of the HL7v2 store that is being created. The string must match the following regex: `[\p{L}\p{N}_\-\.]{1,256}`.
func (o Hl7V2StoreOutput) Hl7V2StoreId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Hl7V2Store) pulumi.StringPtrOutput { return v.Hl7V2StoreId }).(pulumi.StringPtrOutput)
}

// User-supplied key-value pairs used to organize HL7v2 stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
func (o Hl7V2StoreOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Hl7V2Store) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o Hl7V2StoreOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Hl7V2Store) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name of the HL7v2 store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/hl7V2Stores/{hl7v2_store_id}`.
func (o Hl7V2StoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Hl7V2Store) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of notification configs. Each configuration uses a filter to determine whether to publish a message (both Ingest & Create) on the corresponding notification destination. Only the message name is sent as part of the notification. Supplied by the client.
func (o Hl7V2StoreOutput) NotificationConfigs() Hl7V2NotificationConfigResponseArrayOutput {
	return o.ApplyT(func(v *Hl7V2Store) Hl7V2NotificationConfigResponseArrayOutput { return v.NotificationConfigs }).(Hl7V2NotificationConfigResponseArrayOutput)
}

// The configuration for the parser. It determines how the server parses the messages.
func (o Hl7V2StoreOutput) ParserConfig() ParserConfigResponseOutput {
	return o.ApplyT(func(v *Hl7V2Store) ParserConfigResponseOutput { return v.ParserConfig }).(ParserConfigResponseOutput)
}

func (o Hl7V2StoreOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Hl7V2Store) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Determines whether to reject duplicate messages. A duplicate message is a message with the same raw bytes as a message that has already been ingested/created in this HL7v2 store. The default value is false, meaning that the store accepts the duplicate messages and it also returns the same ACK message in the IngestMessageResponse as has been returned previously. Note that only one resource is created in the store. When this field is set to true, CreateMessage/IngestMessage requests with a duplicate message will be rejected by the store, and IngestMessageErrorDetail returns a NACK message upon rejection.
func (o Hl7V2StoreOutput) RejectDuplicateMessage() pulumi.BoolOutput {
	return o.ApplyT(func(v *Hl7V2Store) pulumi.BoolOutput { return v.RejectDuplicateMessage }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Hl7V2StoreInput)(nil)).Elem(), &Hl7V2Store{})
	pulumi.RegisterOutputType(Hl7V2StoreOutput{})
}
