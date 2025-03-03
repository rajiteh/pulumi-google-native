// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a datasource. **Note:** This API requires an admin account to execute.
type DataSource struct {
	pulumi.CustomResourceState

	// If true, sets the datasource to read-only mode. In read-only mode, the Indexing API rejects any requests to index or delete items in this source. Enabling read-only mode does not stop the processing of previously accepted data.
	DisableModifications pulumi.BoolOutput `pulumi:"disableModifications"`
	// Disable serving any search or assist results.
	DisableServing pulumi.BoolOutput `pulumi:"disableServing"`
	// Display name of the datasource The maximum length is 300 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// List of service accounts that have indexing access.
	IndexingServiceAccounts pulumi.StringArrayOutput `pulumi:"indexingServiceAccounts"`
	// This field restricts visibility to items at the datasource level. Items within the datasource are restricted to the union of users and groups included in this field. Note that, this does not ensure access to a specific item, as users need to have ACL permissions on the contained items. This ensures a high level access on the entire datasource, and that the individual items are not shared outside this visibility.
	ItemsVisibility GSuitePrincipalResponseArrayOutput `pulumi:"itemsVisibility"`
	// The name of the datasource resource. Format: datasources/{source_id}. The name is ignored when creating a datasource.
	Name pulumi.StringOutput `pulumi:"name"`
	// IDs of the Long Running Operations (LROs) currently running for this schema.
	OperationIds pulumi.StringArrayOutput `pulumi:"operationIds"`
	// Can a user request to get thumbnail URI for Items indexed in this data source.
	ReturnThumbnailUrls pulumi.BoolOutput `pulumi:"returnThumbnailUrls"`
	// A short name or alias for the source. This value will be used to match the 'source' operator. For example, if the short name is *<value>* then queries like *source:<value>* will only return results for this source. The value must be unique across all datasources. The value must only contain alphanumeric characters (a-zA-Z0-9). The value cannot start with 'google' and cannot be one of the following: mail, gmail, docs, drive, groups, sites, calendar, hangouts, gplus, keep, people, teams. Its maximum length is 32 characters.
	ShortName pulumi.StringOutput `pulumi:"shortName"`
}

// NewDataSource registers a new resource with the given unique name, arguments, and options.
func NewDataSource(ctx *pulumi.Context,
	name string, args *DataSourceArgs, opts ...pulumi.ResourceOption) (*DataSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource DataSource
	err := ctx.RegisterResource("google-native:cloudsearch/v1:DataSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataSource gets an existing DataSource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataSourceState, opts ...pulumi.ResourceOption) (*DataSource, error) {
	var resource DataSource
	err := ctx.ReadResource("google-native:cloudsearch/v1:DataSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataSource resources.
type dataSourceState struct {
}

type DataSourceState struct {
}

func (DataSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceState)(nil)).Elem()
}

type dataSourceArgs struct {
	// If true, sets the datasource to read-only mode. In read-only mode, the Indexing API rejects any requests to index or delete items in this source. Enabling read-only mode does not stop the processing of previously accepted data.
	DisableModifications *bool `pulumi:"disableModifications"`
	// Disable serving any search or assist results.
	DisableServing *bool `pulumi:"disableServing"`
	// Display name of the datasource The maximum length is 300 characters.
	DisplayName string `pulumi:"displayName"`
	// List of service accounts that have indexing access.
	IndexingServiceAccounts []string `pulumi:"indexingServiceAccounts"`
	// This field restricts visibility to items at the datasource level. Items within the datasource are restricted to the union of users and groups included in this field. Note that, this does not ensure access to a specific item, as users need to have ACL permissions on the contained items. This ensures a high level access on the entire datasource, and that the individual items are not shared outside this visibility.
	ItemsVisibility []GSuitePrincipal `pulumi:"itemsVisibility"`
	// The name of the datasource resource. Format: datasources/{source_id}. The name is ignored when creating a datasource.
	Name *string `pulumi:"name"`
	// IDs of the Long Running Operations (LROs) currently running for this schema.
	OperationIds []string `pulumi:"operationIds"`
	// Can a user request to get thumbnail URI for Items indexed in this data source.
	ReturnThumbnailUrls *bool `pulumi:"returnThumbnailUrls"`
	// A short name or alias for the source. This value will be used to match the 'source' operator. For example, if the short name is *<value>* then queries like *source:<value>* will only return results for this source. The value must be unique across all datasources. The value must only contain alphanumeric characters (a-zA-Z0-9). The value cannot start with 'google' and cannot be one of the following: mail, gmail, docs, drive, groups, sites, calendar, hangouts, gplus, keep, people, teams. Its maximum length is 32 characters.
	ShortName *string `pulumi:"shortName"`
}

// The set of arguments for constructing a DataSource resource.
type DataSourceArgs struct {
	// If true, sets the datasource to read-only mode. In read-only mode, the Indexing API rejects any requests to index or delete items in this source. Enabling read-only mode does not stop the processing of previously accepted data.
	DisableModifications pulumi.BoolPtrInput
	// Disable serving any search or assist results.
	DisableServing pulumi.BoolPtrInput
	// Display name of the datasource The maximum length is 300 characters.
	DisplayName pulumi.StringInput
	// List of service accounts that have indexing access.
	IndexingServiceAccounts pulumi.StringArrayInput
	// This field restricts visibility to items at the datasource level. Items within the datasource are restricted to the union of users and groups included in this field. Note that, this does not ensure access to a specific item, as users need to have ACL permissions on the contained items. This ensures a high level access on the entire datasource, and that the individual items are not shared outside this visibility.
	ItemsVisibility GSuitePrincipalArrayInput
	// The name of the datasource resource. Format: datasources/{source_id}. The name is ignored when creating a datasource.
	Name pulumi.StringPtrInput
	// IDs of the Long Running Operations (LROs) currently running for this schema.
	OperationIds pulumi.StringArrayInput
	// Can a user request to get thumbnail URI for Items indexed in this data source.
	ReturnThumbnailUrls pulumi.BoolPtrInput
	// A short name or alias for the source. This value will be used to match the 'source' operator. For example, if the short name is *<value>* then queries like *source:<value>* will only return results for this source. The value must be unique across all datasources. The value must only contain alphanumeric characters (a-zA-Z0-9). The value cannot start with 'google' and cannot be one of the following: mail, gmail, docs, drive, groups, sites, calendar, hangouts, gplus, keep, people, teams. Its maximum length is 32 characters.
	ShortName pulumi.StringPtrInput
}

func (DataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataSourceArgs)(nil)).Elem()
}

type DataSourceInput interface {
	pulumi.Input

	ToDataSourceOutput() DataSourceOutput
	ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput
}

func (*DataSource) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSource)(nil)).Elem()
}

func (i *DataSource) ToDataSourceOutput() DataSourceOutput {
	return i.ToDataSourceOutputWithContext(context.Background())
}

func (i *DataSource) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataSourceOutput)
}

type DataSourceOutput struct{ *pulumi.OutputState }

func (DataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSource)(nil)).Elem()
}

func (o DataSourceOutput) ToDataSourceOutput() DataSourceOutput {
	return o
}

func (o DataSourceOutput) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return o
}

// If true, sets the datasource to read-only mode. In read-only mode, the Indexing API rejects any requests to index or delete items in this source. Enabling read-only mode does not stop the processing of previously accepted data.
func (o DataSourceOutput) DisableModifications() pulumi.BoolOutput {
	return o.ApplyT(func(v *DataSource) pulumi.BoolOutput { return v.DisableModifications }).(pulumi.BoolOutput)
}

// Disable serving any search or assist results.
func (o DataSourceOutput) DisableServing() pulumi.BoolOutput {
	return o.ApplyT(func(v *DataSource) pulumi.BoolOutput { return v.DisableServing }).(pulumi.BoolOutput)
}

// Display name of the datasource The maximum length is 300 characters.
func (o DataSourceOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// List of service accounts that have indexing access.
func (o DataSourceOutput) IndexingServiceAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringArrayOutput { return v.IndexingServiceAccounts }).(pulumi.StringArrayOutput)
}

// This field restricts visibility to items at the datasource level. Items within the datasource are restricted to the union of users and groups included in this field. Note that, this does not ensure access to a specific item, as users need to have ACL permissions on the contained items. This ensures a high level access on the entire datasource, and that the individual items are not shared outside this visibility.
func (o DataSourceOutput) ItemsVisibility() GSuitePrincipalResponseArrayOutput {
	return o.ApplyT(func(v *DataSource) GSuitePrincipalResponseArrayOutput { return v.ItemsVisibility }).(GSuitePrincipalResponseArrayOutput)
}

// The name of the datasource resource. Format: datasources/{source_id}. The name is ignored when creating a datasource.
func (o DataSourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// IDs of the Long Running Operations (LROs) currently running for this schema.
func (o DataSourceOutput) OperationIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringArrayOutput { return v.OperationIds }).(pulumi.StringArrayOutput)
}

// Can a user request to get thumbnail URI for Items indexed in this data source.
func (o DataSourceOutput) ReturnThumbnailUrls() pulumi.BoolOutput {
	return o.ApplyT(func(v *DataSource) pulumi.BoolOutput { return v.ReturnThumbnailUrls }).(pulumi.BoolOutput)
}

// A short name or alias for the source. This value will be used to match the 'source' operator. For example, if the short name is *<value>* then queries like *source:<value>* will only return results for this source. The value must be unique across all datasources. The value must only contain alphanumeric characters (a-zA-Z0-9). The value cannot start with 'google' and cannot be one of the following: mail, gmail, docs, drive, groups, sites, calendar, hangouts, gplus, keep, people, teams. Its maximum length is 32 characters.
func (o DataSourceOutput) ShortName() pulumi.StringOutput {
	return o.ApplyT(func(v *DataSource) pulumi.StringOutput { return v.ShortName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataSourceInput)(nil)).Elem(), &DataSource{})
	pulumi.RegisterOutputType(DataSourceOutput{})
}
