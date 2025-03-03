// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new data transfer configuration.
type TransferConfig struct {
	pulumi.CustomResourceState

	// Optional OAuth2 authorization code to use with this transfer configuration. This is required only if `transferConfig.dataSourceId` is 'youtube_channel' and new credentials are needed, as indicated by `CheckValidCreds`. In order to obtain authorization_code, make a request to the following URL: https://www.gstatic.com/bigquerydatatransfer/oauthz/auth?redirect_uri=urn:ietf:wg:oauth:2.0:oob&response_type=authorization_code&client_id=client_id&scope=data_source_scopes * The client_id is the OAuth client_id of the a data source as returned by ListDataSources method. * data_source_scopes are the scopes returned by ListDataSources method. Note that this should not be set when `service_account_name` is used to create the transfer config.
	AuthorizationCode pulumi.StringPtrOutput `pulumi:"authorizationCode"`
	// The number of days to look back to automatically refresh the data. For example, if `data_refresh_window_days = 10`, then every day BigQuery reingests data for [today-10, today-1], rather than ingesting data for just [today-1]. Only valid if the data source supports the feature. Set the value to 0 to use the default value.
	DataRefreshWindowDays pulumi.IntOutput `pulumi:"dataRefreshWindowDays"`
	// Data source ID. This cannot be changed once data transfer is created. The full list of available data source IDs can be returned through an API call: https://cloud.google.com/bigquery-transfer/docs/reference/datatransfer/rest/v1/projects.locations.dataSources/list
	DataSourceId pulumi.StringOutput `pulumi:"dataSourceId"`
	// Region in which BigQuery dataset is located.
	DatasetRegion pulumi.StringOutput `pulumi:"datasetRegion"`
	// The BigQuery target dataset id.
	DestinationDatasetId pulumi.StringOutput `pulumi:"destinationDatasetId"`
	// Is this config disabled. When set to true, no runs are scheduled for a given transfer.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// User specified display name for the data transfer.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Email notifications will be sent according to these preferences to the email address of the user who owns this transfer config.
	EmailPreferences EmailPreferencesResponseOutput `pulumi:"emailPreferences"`
	Location         pulumi.StringOutput            `pulumi:"location"`
	// The resource name of the transfer config. Transfer config names have the form `projects/{project_id}/locations/{region}/transferConfigs/{config_id}`. Where `config_id` is usually a uuid, even though it is not guaranteed or required. The name is ignored when creating a transfer config.
	Name pulumi.StringOutput `pulumi:"name"`
	// Next time when data transfer will run.
	NextRunTime pulumi.StringOutput `pulumi:"nextRunTime"`
	// Pub/Sub topic where notifications will be sent after transfer runs associated with this transfer config finish. The format for specifying a pubsub topic is: `projects/{project}/topics/{topic}`
	NotificationPubsubTopic pulumi.StringOutput `pulumi:"notificationPubsubTopic"`
	// Information about the user whose credentials are used to transfer data. Populated only for `transferConfigs.get` requests. In case the user information is not available, this field will not be populated.
	OwnerInfo UserInfoResponseOutput `pulumi:"ownerInfo"`
	// Parameters specific to each data source. For more information see the bq tab in the 'Setting up a data transfer' section for each data source. For example the parameters for Cloud Storage transfers are listed here: https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq
	Params  pulumi.StringMapOutput `pulumi:"params"`
	Project pulumi.StringOutput    `pulumi:"project"`
	// Data transfer schedule. If the data source does not support a custom schedule, this should be empty. If it is empty, the default value for the data source will be used. The specified times are in UTC. Examples of valid format: `1st,3rd monday of month 15:30`, `every wed,fri of jan,jun 13:15`, and `first sunday of quarter 00:00`. See more explanation about the format here: https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format NOTE: The minimum interval time between recurring transfers depends on the data source; refer to the documentation for your data source.
	Schedule pulumi.StringOutput `pulumi:"schedule"`
	// Options customizing the data transfer schedule.
	ScheduleOptions ScheduleOptionsResponseOutput `pulumi:"scheduleOptions"`
	// Optional service account email. If this field is set, the transfer config will be created with this service account's credentials. It requires that the requesting user calling this API has permissions to act as this service account. Note that not all data sources support service account credentials when creating a transfer config. For the latest list of data sources, read about [using service accounts](https://cloud.google.com/bigquery-transfer/docs/use-service-accounts).
	ServiceAccountName pulumi.StringPtrOutput `pulumi:"serviceAccountName"`
	// State of the most recently updated transfer run.
	State pulumi.StringOutput `pulumi:"state"`
	// Data transfer modification time. Ignored by server on input.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Deprecated. Unique ID of the user on whose behalf transfer is done.
	//
	// Deprecated: Deprecated. Unique ID of the user on whose behalf transfer is done.
	UserId pulumi.StringOutput `pulumi:"userId"`
	// Optional version info. This is required only if `transferConfig.dataSourceId` is not 'youtube_channel' and new credentials are needed, as indicated by `CheckValidCreds`. In order to obtain version info, make a request to the following URL: https://www.gstatic.com/bigquerydatatransfer/oauthz/auth?redirect_uri=urn:ietf:wg:oauth:2.0:oob&response_type=version_info&client_id=client_id&scope=data_source_scopes * The client_id is the OAuth client_id of the a data source as returned by ListDataSources method. * data_source_scopes are the scopes returned by ListDataSources method. Note that this should not be set when `service_account_name` is used to create the transfer config.
	VersionInfo pulumi.StringPtrOutput `pulumi:"versionInfo"`
}

// NewTransferConfig registers a new resource with the given unique name, arguments, and options.
func NewTransferConfig(ctx *pulumi.Context,
	name string, args *TransferConfigArgs, opts ...pulumi.ResourceOption) (*TransferConfig, error) {
	if args == nil {
		args = &TransferConfigArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource TransferConfig
	err := ctx.RegisterResource("google-native:bigquerydatatransfer/v1:TransferConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransferConfig gets an existing TransferConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransferConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransferConfigState, opts ...pulumi.ResourceOption) (*TransferConfig, error) {
	var resource TransferConfig
	err := ctx.ReadResource("google-native:bigquerydatatransfer/v1:TransferConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransferConfig resources.
type transferConfigState struct {
}

type TransferConfigState struct {
}

func (TransferConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*transferConfigState)(nil)).Elem()
}

type transferConfigArgs struct {
	// Optional OAuth2 authorization code to use with this transfer configuration. This is required only if `transferConfig.dataSourceId` is 'youtube_channel' and new credentials are needed, as indicated by `CheckValidCreds`. In order to obtain authorization_code, make a request to the following URL: https://www.gstatic.com/bigquerydatatransfer/oauthz/auth?redirect_uri=urn:ietf:wg:oauth:2.0:oob&response_type=authorization_code&client_id=client_id&scope=data_source_scopes * The client_id is the OAuth client_id of the a data source as returned by ListDataSources method. * data_source_scopes are the scopes returned by ListDataSources method. Note that this should not be set when `service_account_name` is used to create the transfer config.
	AuthorizationCode *string `pulumi:"authorizationCode"`
	// The number of days to look back to automatically refresh the data. For example, if `data_refresh_window_days = 10`, then every day BigQuery reingests data for [today-10, today-1], rather than ingesting data for just [today-1]. Only valid if the data source supports the feature. Set the value to 0 to use the default value.
	DataRefreshWindowDays *int `pulumi:"dataRefreshWindowDays"`
	// Data source ID. This cannot be changed once data transfer is created. The full list of available data source IDs can be returned through an API call: https://cloud.google.com/bigquery-transfer/docs/reference/datatransfer/rest/v1/projects.locations.dataSources/list
	DataSourceId *string `pulumi:"dataSourceId"`
	// The BigQuery target dataset id.
	DestinationDatasetId *string `pulumi:"destinationDatasetId"`
	// Is this config disabled. When set to true, no runs are scheduled for a given transfer.
	Disabled *bool `pulumi:"disabled"`
	// User specified display name for the data transfer.
	DisplayName *string `pulumi:"displayName"`
	// Email notifications will be sent according to these preferences to the email address of the user who owns this transfer config.
	EmailPreferences *EmailPreferences `pulumi:"emailPreferences"`
	Location         *string           `pulumi:"location"`
	// The resource name of the transfer config. Transfer config names have the form `projects/{project_id}/locations/{region}/transferConfigs/{config_id}`. Where `config_id` is usually a uuid, even though it is not guaranteed or required. The name is ignored when creating a transfer config.
	Name *string `pulumi:"name"`
	// Pub/Sub topic where notifications will be sent after transfer runs associated with this transfer config finish. The format for specifying a pubsub topic is: `projects/{project}/topics/{topic}`
	NotificationPubsubTopic *string `pulumi:"notificationPubsubTopic"`
	// Parameters specific to each data source. For more information see the bq tab in the 'Setting up a data transfer' section for each data source. For example the parameters for Cloud Storage transfers are listed here: https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq
	Params  map[string]string `pulumi:"params"`
	Project *string           `pulumi:"project"`
	// Data transfer schedule. If the data source does not support a custom schedule, this should be empty. If it is empty, the default value for the data source will be used. The specified times are in UTC. Examples of valid format: `1st,3rd monday of month 15:30`, `every wed,fri of jan,jun 13:15`, and `first sunday of quarter 00:00`. See more explanation about the format here: https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format NOTE: The minimum interval time between recurring transfers depends on the data source; refer to the documentation for your data source.
	Schedule *string `pulumi:"schedule"`
	// Options customizing the data transfer schedule.
	ScheduleOptions *ScheduleOptions `pulumi:"scheduleOptions"`
	// Optional service account email. If this field is set, the transfer config will be created with this service account's credentials. It requires that the requesting user calling this API has permissions to act as this service account. Note that not all data sources support service account credentials when creating a transfer config. For the latest list of data sources, read about [using service accounts](https://cloud.google.com/bigquery-transfer/docs/use-service-accounts).
	ServiceAccountName *string `pulumi:"serviceAccountName"`
	// Deprecated. Unique ID of the user on whose behalf transfer is done.
	//
	// Deprecated: Deprecated. Unique ID of the user on whose behalf transfer is done.
	UserId *string `pulumi:"userId"`
	// Optional version info. This is required only if `transferConfig.dataSourceId` is not 'youtube_channel' and new credentials are needed, as indicated by `CheckValidCreds`. In order to obtain version info, make a request to the following URL: https://www.gstatic.com/bigquerydatatransfer/oauthz/auth?redirect_uri=urn:ietf:wg:oauth:2.0:oob&response_type=version_info&client_id=client_id&scope=data_source_scopes * The client_id is the OAuth client_id of the a data source as returned by ListDataSources method. * data_source_scopes are the scopes returned by ListDataSources method. Note that this should not be set when `service_account_name` is used to create the transfer config.
	VersionInfo *string `pulumi:"versionInfo"`
}

// The set of arguments for constructing a TransferConfig resource.
type TransferConfigArgs struct {
	// Optional OAuth2 authorization code to use with this transfer configuration. This is required only if `transferConfig.dataSourceId` is 'youtube_channel' and new credentials are needed, as indicated by `CheckValidCreds`. In order to obtain authorization_code, make a request to the following URL: https://www.gstatic.com/bigquerydatatransfer/oauthz/auth?redirect_uri=urn:ietf:wg:oauth:2.0:oob&response_type=authorization_code&client_id=client_id&scope=data_source_scopes * The client_id is the OAuth client_id of the a data source as returned by ListDataSources method. * data_source_scopes are the scopes returned by ListDataSources method. Note that this should not be set when `service_account_name` is used to create the transfer config.
	AuthorizationCode pulumi.StringPtrInput
	// The number of days to look back to automatically refresh the data. For example, if `data_refresh_window_days = 10`, then every day BigQuery reingests data for [today-10, today-1], rather than ingesting data for just [today-1]. Only valid if the data source supports the feature. Set the value to 0 to use the default value.
	DataRefreshWindowDays pulumi.IntPtrInput
	// Data source ID. This cannot be changed once data transfer is created. The full list of available data source IDs can be returned through an API call: https://cloud.google.com/bigquery-transfer/docs/reference/datatransfer/rest/v1/projects.locations.dataSources/list
	DataSourceId pulumi.StringPtrInput
	// The BigQuery target dataset id.
	DestinationDatasetId pulumi.StringPtrInput
	// Is this config disabled. When set to true, no runs are scheduled for a given transfer.
	Disabled pulumi.BoolPtrInput
	// User specified display name for the data transfer.
	DisplayName pulumi.StringPtrInput
	// Email notifications will be sent according to these preferences to the email address of the user who owns this transfer config.
	EmailPreferences EmailPreferencesPtrInput
	Location         pulumi.StringPtrInput
	// The resource name of the transfer config. Transfer config names have the form `projects/{project_id}/locations/{region}/transferConfigs/{config_id}`. Where `config_id` is usually a uuid, even though it is not guaranteed or required. The name is ignored when creating a transfer config.
	Name pulumi.StringPtrInput
	// Pub/Sub topic where notifications will be sent after transfer runs associated with this transfer config finish. The format for specifying a pubsub topic is: `projects/{project}/topics/{topic}`
	NotificationPubsubTopic pulumi.StringPtrInput
	// Parameters specific to each data source. For more information see the bq tab in the 'Setting up a data transfer' section for each data source. For example the parameters for Cloud Storage transfers are listed here: https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq
	Params  pulumi.StringMapInput
	Project pulumi.StringPtrInput
	// Data transfer schedule. If the data source does not support a custom schedule, this should be empty. If it is empty, the default value for the data source will be used. The specified times are in UTC. Examples of valid format: `1st,3rd monday of month 15:30`, `every wed,fri of jan,jun 13:15`, and `first sunday of quarter 00:00`. See more explanation about the format here: https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format NOTE: The minimum interval time between recurring transfers depends on the data source; refer to the documentation for your data source.
	Schedule pulumi.StringPtrInput
	// Options customizing the data transfer schedule.
	ScheduleOptions ScheduleOptionsPtrInput
	// Optional service account email. If this field is set, the transfer config will be created with this service account's credentials. It requires that the requesting user calling this API has permissions to act as this service account. Note that not all data sources support service account credentials when creating a transfer config. For the latest list of data sources, read about [using service accounts](https://cloud.google.com/bigquery-transfer/docs/use-service-accounts).
	ServiceAccountName pulumi.StringPtrInput
	// Deprecated. Unique ID of the user on whose behalf transfer is done.
	//
	// Deprecated: Deprecated. Unique ID of the user on whose behalf transfer is done.
	UserId pulumi.StringPtrInput
	// Optional version info. This is required only if `transferConfig.dataSourceId` is not 'youtube_channel' and new credentials are needed, as indicated by `CheckValidCreds`. In order to obtain version info, make a request to the following URL: https://www.gstatic.com/bigquerydatatransfer/oauthz/auth?redirect_uri=urn:ietf:wg:oauth:2.0:oob&response_type=version_info&client_id=client_id&scope=data_source_scopes * The client_id is the OAuth client_id of the a data source as returned by ListDataSources method. * data_source_scopes are the scopes returned by ListDataSources method. Note that this should not be set when `service_account_name` is used to create the transfer config.
	VersionInfo pulumi.StringPtrInput
}

func (TransferConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transferConfigArgs)(nil)).Elem()
}

type TransferConfigInput interface {
	pulumi.Input

	ToTransferConfigOutput() TransferConfigOutput
	ToTransferConfigOutputWithContext(ctx context.Context) TransferConfigOutput
}

func (*TransferConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferConfig)(nil)).Elem()
}

func (i *TransferConfig) ToTransferConfigOutput() TransferConfigOutput {
	return i.ToTransferConfigOutputWithContext(context.Background())
}

func (i *TransferConfig) ToTransferConfigOutputWithContext(ctx context.Context) TransferConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferConfigOutput)
}

type TransferConfigOutput struct{ *pulumi.OutputState }

func (TransferConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferConfig)(nil)).Elem()
}

func (o TransferConfigOutput) ToTransferConfigOutput() TransferConfigOutput {
	return o
}

func (o TransferConfigOutput) ToTransferConfigOutputWithContext(ctx context.Context) TransferConfigOutput {
	return o
}

// Optional OAuth2 authorization code to use with this transfer configuration. This is required only if `transferConfig.dataSourceId` is 'youtube_channel' and new credentials are needed, as indicated by `CheckValidCreds`. In order to obtain authorization_code, make a request to the following URL: https://www.gstatic.com/bigquerydatatransfer/oauthz/auth?redirect_uri=urn:ietf:wg:oauth:2.0:oob&response_type=authorization_code&client_id=client_id&scope=data_source_scopes * The client_id is the OAuth client_id of the a data source as returned by ListDataSources method. * data_source_scopes are the scopes returned by ListDataSources method. Note that this should not be set when `service_account_name` is used to create the transfer config.
func (o TransferConfigOutput) AuthorizationCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransferConfig) pulumi.StringPtrOutput { return v.AuthorizationCode }).(pulumi.StringPtrOutput)
}

// The number of days to look back to automatically refresh the data. For example, if `data_refresh_window_days = 10`, then every day BigQuery reingests data for [today-10, today-1], rather than ingesting data for just [today-1]. Only valid if the data source supports the feature. Set the value to 0 to use the default value.
func (o TransferConfigOutput) DataRefreshWindowDays() pulumi.IntOutput {
	return o.ApplyT(func(v *TransferConfig) pulumi.IntOutput { return v.DataRefreshWindowDays }).(pulumi.IntOutput)
}

// Data source ID. This cannot be changed once data transfer is created. The full list of available data source IDs can be returned through an API call: https://cloud.google.com/bigquery-transfer/docs/reference/datatransfer/rest/v1/projects.locations.dataSources/list
func (o TransferConfigOutput) DataSourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferConfig) pulumi.StringOutput { return v.DataSourceId }).(pulumi.StringOutput)
}

// Region in which BigQuery dataset is located.
func (o TransferConfigOutput) DatasetRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferConfig) pulumi.StringOutput { return v.DatasetRegion }).(pulumi.StringOutput)
}

// The BigQuery target dataset id.
func (o TransferConfigOutput) DestinationDatasetId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferConfig) pulumi.StringOutput { return v.DestinationDatasetId }).(pulumi.StringOutput)
}

// Is this config disabled. When set to true, no runs are scheduled for a given transfer.
func (o TransferConfigOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *TransferConfig) pulumi.BoolOutput { return v.Disabled }).(pulumi.BoolOutput)
}

// User specified display name for the data transfer.
func (o TransferConfigOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferConfig) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Email notifications will be sent according to these preferences to the email address of the user who owns this transfer config.
func (o TransferConfigOutput) EmailPreferences() EmailPreferencesResponseOutput {
	return o.ApplyT(func(v *TransferConfig) EmailPreferencesResponseOutput { return v.EmailPreferences }).(EmailPreferencesResponseOutput)
}

func (o TransferConfigOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferConfig) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the transfer config. Transfer config names have the form `projects/{project_id}/locations/{region}/transferConfigs/{config_id}`. Where `config_id` is usually a uuid, even though it is not guaranteed or required. The name is ignored when creating a transfer config.
func (o TransferConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Next time when data transfer will run.
func (o TransferConfigOutput) NextRunTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferConfig) pulumi.StringOutput { return v.NextRunTime }).(pulumi.StringOutput)
}

// Pub/Sub topic where notifications will be sent after transfer runs associated with this transfer config finish. The format for specifying a pubsub topic is: `projects/{project}/topics/{topic}`
func (o TransferConfigOutput) NotificationPubsubTopic() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferConfig) pulumi.StringOutput { return v.NotificationPubsubTopic }).(pulumi.StringOutput)
}

// Information about the user whose credentials are used to transfer data. Populated only for `transferConfigs.get` requests. In case the user information is not available, this field will not be populated.
func (o TransferConfigOutput) OwnerInfo() UserInfoResponseOutput {
	return o.ApplyT(func(v *TransferConfig) UserInfoResponseOutput { return v.OwnerInfo }).(UserInfoResponseOutput)
}

// Parameters specific to each data source. For more information see the bq tab in the 'Setting up a data transfer' section for each data source. For example the parameters for Cloud Storage transfers are listed here: https://cloud.google.com/bigquery-transfer/docs/cloud-storage-transfer#bq
func (o TransferConfigOutput) Params() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TransferConfig) pulumi.StringMapOutput { return v.Params }).(pulumi.StringMapOutput)
}

func (o TransferConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Data transfer schedule. If the data source does not support a custom schedule, this should be empty. If it is empty, the default value for the data source will be used. The specified times are in UTC. Examples of valid format: `1st,3rd monday of month 15:30`, `every wed,fri of jan,jun 13:15`, and `first sunday of quarter 00:00`. See more explanation about the format here: https://cloud.google.com/appengine/docs/flexible/python/scheduling-jobs-with-cron-yaml#the_schedule_format NOTE: The minimum interval time between recurring transfers depends on the data source; refer to the documentation for your data source.
func (o TransferConfigOutput) Schedule() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferConfig) pulumi.StringOutput { return v.Schedule }).(pulumi.StringOutput)
}

// Options customizing the data transfer schedule.
func (o TransferConfigOutput) ScheduleOptions() ScheduleOptionsResponseOutput {
	return o.ApplyT(func(v *TransferConfig) ScheduleOptionsResponseOutput { return v.ScheduleOptions }).(ScheduleOptionsResponseOutput)
}

// Optional service account email. If this field is set, the transfer config will be created with this service account's credentials. It requires that the requesting user calling this API has permissions to act as this service account. Note that not all data sources support service account credentials when creating a transfer config. For the latest list of data sources, read about [using service accounts](https://cloud.google.com/bigquery-transfer/docs/use-service-accounts).
func (o TransferConfigOutput) ServiceAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransferConfig) pulumi.StringPtrOutput { return v.ServiceAccountName }).(pulumi.StringPtrOutput)
}

// State of the most recently updated transfer run.
func (o TransferConfigOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferConfig) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Data transfer modification time. Ignored by server on input.
func (o TransferConfigOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferConfig) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Deprecated. Unique ID of the user on whose behalf transfer is done.
//
// Deprecated: Deprecated. Unique ID of the user on whose behalf transfer is done.
func (o TransferConfigOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferConfig) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

// Optional version info. This is required only if `transferConfig.dataSourceId` is not 'youtube_channel' and new credentials are needed, as indicated by `CheckValidCreds`. In order to obtain version info, make a request to the following URL: https://www.gstatic.com/bigquerydatatransfer/oauthz/auth?redirect_uri=urn:ietf:wg:oauth:2.0:oob&response_type=version_info&client_id=client_id&scope=data_source_scopes * The client_id is the OAuth client_id of the a data source as returned by ListDataSources method. * data_source_scopes are the scopes returned by ListDataSources method. Note that this should not be set when `service_account_name` is used to create the transfer config.
func (o TransferConfigOutput) VersionInfo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransferConfig) pulumi.StringPtrOutput { return v.VersionInfo }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransferConfigInput)(nil)).Elem(), &TransferConfig{})
	pulumi.RegisterOutputType(TransferConfigOutput{})
}
