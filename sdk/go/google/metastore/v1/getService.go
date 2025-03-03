// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of a single service.
func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("google-native:metastore/v1:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceArgs struct {
	Location  string  `pulumi:"location"`
	Project   *string `pulumi:"project"`
	ServiceId string  `pulumi:"serviceId"`
}

type LookupServiceResult struct {
	// A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
	ArtifactGcsUri string `pulumi:"artifactGcsUri"`
	// The time when the metastore service was created.
	CreateTime string `pulumi:"createTime"`
	// Immutable. The database type that the Metastore service stores its data.
	DatabaseType string `pulumi:"databaseType"`
	// Immutable. Information used to configure the Dataproc Metastore service to encrypt customer data at rest. Cannot be updated.
	EncryptionConfig EncryptionConfigResponse `pulumi:"encryptionConfig"`
	// The URI of the endpoint used to access the metastore service.
	EndpointUri string `pulumi:"endpointUri"`
	// Configuration information specific to running Hive metastore software as the metastore service.
	HiveMetastoreConfig HiveMetastoreConfigResponse `pulumi:"hiveMetastoreConfig"`
	// User-defined labels for the metastore service.
	Labels map[string]string `pulumi:"labels"`
	// The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time. Maintenance window is not needed for services with the SPANNER database type.
	MaintenanceWindow MaintenanceWindowResponse `pulumi:"maintenanceWindow"`
	// The metadata management activities of the metastore service.
	MetadataManagementActivity MetadataManagementActivityResponse `pulumi:"metadataManagementActivity"`
	// Immutable. The relative resource name of the metastore service, in the following format:projects/{project_number}/locations/{location_id}/services/{service_id}.
	Name string `pulumi:"name"`
	// Immutable. The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:projects/{project_number}/global/networks/{network_id}.
	Network string `pulumi:"network"`
	// The configuration specifying the network settings for the Dataproc Metastore service.
	NetworkConfig NetworkConfigResponse `pulumi:"networkConfig"`
	// The TCP port at which the metastore service is reached. Default: 9083.
	Port int `pulumi:"port"`
	// Immutable. The release channel of the service. If unspecified, defaults to STABLE.
	ReleaseChannel string `pulumi:"releaseChannel"`
	// Scaling configuration of the metastore service.
	ScalingConfig ScalingConfigResponse `pulumi:"scalingConfig"`
	// The current state of the metastore service.
	State string `pulumi:"state"`
	// Additional information about the current state of the metastore service, if available.
	StateMessage string `pulumi:"stateMessage"`
	// The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
	TelemetryConfig TelemetryConfigResponse `pulumi:"telemetryConfig"`
	// The tier of the service.
	Tier string `pulumi:"tier"`
	// The globally unique resource identifier of the metastore service.
	Uid string `pulumi:"uid"`
	// The time when the metastore service was last updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupServiceOutput(ctx *pulumi.Context, args LookupServiceOutputArgs, opts ...pulumi.InvokeOption) LookupServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServiceResult, error) {
			args := v.(LookupServiceArgs)
			r, err := LookupService(ctx, &args, opts...)
			var s LookupServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServiceResultOutput)
}

type LookupServiceOutputArgs struct {
	Location  pulumi.StringInput    `pulumi:"location"`
	Project   pulumi.StringPtrInput `pulumi:"project"`
	ServiceId pulumi.StringInput    `pulumi:"serviceId"`
}

func (LookupServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceArgs)(nil)).Elem()
}

type LookupServiceResultOutput struct{ *pulumi.OutputState }

func (LookupServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceResult)(nil)).Elem()
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutput() LookupServiceResultOutput {
	return o
}

func (o LookupServiceResultOutput) ToLookupServiceResultOutputWithContext(ctx context.Context) LookupServiceResultOutput {
	return o
}

// A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
func (o LookupServiceResultOutput) ArtifactGcsUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.ArtifactGcsUri }).(pulumi.StringOutput)
}

// The time when the metastore service was created.
func (o LookupServiceResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Immutable. The database type that the Metastore service stores its data.
func (o LookupServiceResultOutput) DatabaseType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.DatabaseType }).(pulumi.StringOutput)
}

// Immutable. Information used to configure the Dataproc Metastore service to encrypt customer data at rest. Cannot be updated.
func (o LookupServiceResultOutput) EncryptionConfig() EncryptionConfigResponseOutput {
	return o.ApplyT(func(v LookupServiceResult) EncryptionConfigResponse { return v.EncryptionConfig }).(EncryptionConfigResponseOutput)
}

// The URI of the endpoint used to access the metastore service.
func (o LookupServiceResultOutput) EndpointUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.EndpointUri }).(pulumi.StringOutput)
}

// Configuration information specific to running Hive metastore software as the metastore service.
func (o LookupServiceResultOutput) HiveMetastoreConfig() HiveMetastoreConfigResponseOutput {
	return o.ApplyT(func(v LookupServiceResult) HiveMetastoreConfigResponse { return v.HiveMetastoreConfig }).(HiveMetastoreConfigResponseOutput)
}

// User-defined labels for the metastore service.
func (o LookupServiceResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServiceResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time. Maintenance window is not needed for services with the SPANNER database type.
func (o LookupServiceResultOutput) MaintenanceWindow() MaintenanceWindowResponseOutput {
	return o.ApplyT(func(v LookupServiceResult) MaintenanceWindowResponse { return v.MaintenanceWindow }).(MaintenanceWindowResponseOutput)
}

// The metadata management activities of the metastore service.
func (o LookupServiceResultOutput) MetadataManagementActivity() MetadataManagementActivityResponseOutput {
	return o.ApplyT(func(v LookupServiceResult) MetadataManagementActivityResponse { return v.MetadataManagementActivity }).(MetadataManagementActivityResponseOutput)
}

// Immutable. The relative resource name of the metastore service, in the following format:projects/{project_number}/locations/{location_id}/services/{service_id}.
func (o LookupServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Immutable. The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:projects/{project_number}/global/networks/{network_id}.
func (o LookupServiceResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Network }).(pulumi.StringOutput)
}

// The configuration specifying the network settings for the Dataproc Metastore service.
func (o LookupServiceResultOutput) NetworkConfig() NetworkConfigResponseOutput {
	return o.ApplyT(func(v LookupServiceResult) NetworkConfigResponse { return v.NetworkConfig }).(NetworkConfigResponseOutput)
}

// The TCP port at which the metastore service is reached. Default: 9083.
func (o LookupServiceResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupServiceResult) int { return v.Port }).(pulumi.IntOutput)
}

// Immutable. The release channel of the service. If unspecified, defaults to STABLE.
func (o LookupServiceResultOutput) ReleaseChannel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.ReleaseChannel }).(pulumi.StringOutput)
}

// Scaling configuration of the metastore service.
func (o LookupServiceResultOutput) ScalingConfig() ScalingConfigResponseOutput {
	return o.ApplyT(func(v LookupServiceResult) ScalingConfigResponse { return v.ScalingConfig }).(ScalingConfigResponseOutput)
}

// The current state of the metastore service.
func (o LookupServiceResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.State }).(pulumi.StringOutput)
}

// Additional information about the current state of the metastore service, if available.
func (o LookupServiceResultOutput) StateMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.StateMessage }).(pulumi.StringOutput)
}

// The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
func (o LookupServiceResultOutput) TelemetryConfig() TelemetryConfigResponseOutput {
	return o.ApplyT(func(v LookupServiceResult) TelemetryConfigResponse { return v.TelemetryConfig }).(TelemetryConfigResponseOutput)
}

// The tier of the service.
func (o LookupServiceResultOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Tier }).(pulumi.StringOutput)
}

// The globally unique resource identifier of the metastore service.
func (o LookupServiceResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.Uid }).(pulumi.StringOutput)
}

// The time when the metastore service was last updated.
func (o LookupServiceResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceResultOutput{})
}
