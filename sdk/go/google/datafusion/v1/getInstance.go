// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single Data Fusion instance.
func LookupInstance(ctx *pulumi.Context, args *LookupInstanceArgs, opts ...pulumi.InvokeOption) (*LookupInstanceResult, error) {
	var rv LookupInstanceResult
	err := ctx.Invoke("google-native:datafusion/v1:getInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceArgs struct {
	InstanceId string  `pulumi:"instanceId"`
	Location   string  `pulumi:"location"`
	Project    *string `pulumi:"project"`
}

type LookupInstanceResult struct {
	// List of accelerators enabled for this CDF instance.
	Accelerators []AcceleratorResponse `pulumi:"accelerators"`
	// Endpoint on which the REST APIs is accessible.
	ApiEndpoint string `pulumi:"apiEndpoint"`
	// Available versions that the instance can be upgraded to using UpdateInstanceRequest.
	AvailableVersion []VersionResponse `pulumi:"availableVersion"`
	// The time the instance was created.
	CreateTime string `pulumi:"createTime"`
	// The crypto key configuration. This field is used by the Customer-Managed Encryption Keys (CMEK) feature.
	CryptoKeyConfig CryptoKeyConfigResponse `pulumi:"cryptoKeyConfig"`
	// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.
	DataprocServiceAccount string `pulumi:"dataprocServiceAccount"`
	// A description of this instance.
	Description string `pulumi:"description"`
	// If the instance state is DISABLED, the reason for disabling the instance.
	DisabledReason []string `pulumi:"disabledReason"`
	// Display name for an instance.
	DisplayName string `pulumi:"displayName"`
	// Option to enable granular role-based access control.
	EnableRbac bool `pulumi:"enableRbac"`
	// Option to enable Stackdriver Logging.
	EnableStackdriverLogging bool `pulumi:"enableStackdriverLogging"`
	// Option to enable Stackdriver Monitoring.
	EnableStackdriverMonitoring bool `pulumi:"enableStackdriverMonitoring"`
	// Option to enable granular zone separation.
	EnableZoneSeparation bool `pulumi:"enableZoneSeparation"`
	// Option to enable and pass metadata for event publishing.
	EventPublishConfig EventPublishConfigResponse `pulumi:"eventPublishConfig"`
	// Cloud Storage bucket generated by Data Fusion in the customer project.
	GcsBucket string `pulumi:"gcsBucket"`
	// The resource labels for instance to use to annotate any related underlying resources such as Compute Engine VMs. The character '=' is not allowed to be used within the labels.
	Labels map[string]string `pulumi:"labels"`
	// The name of this instance is in the form of projects/{project}/locations/{location}/instances/{instance}.
	Name string `pulumi:"name"`
	// Network configuration options. These are required when a private Data Fusion instance is to be created.
	NetworkConfig NetworkConfigResponse `pulumi:"networkConfig"`
	// Map of additional options used to configure the behavior of Data Fusion instance.
	Options map[string]string `pulumi:"options"`
	// P4 service account for the customer project.
	P4ServiceAccount string `pulumi:"p4ServiceAccount"`
	// Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
	PrivateInstance bool `pulumi:"privateInstance"`
	// Reserved for future use.
	SatisfiesPzs bool `pulumi:"satisfiesPzs"`
	// Deprecated. Use tenant_project_id instead to extract the tenant project ID.
	//
	// Deprecated: Output only. Deprecated. Use tenant_project_id instead to extract the tenant project ID.
	ServiceAccount string `pulumi:"serviceAccount"`
	// Endpoint on which the Data Fusion UI is accessible.
	ServiceEndpoint string `pulumi:"serviceEndpoint"`
	// The current state of this Data Fusion instance.
	State string `pulumi:"state"`
	// Additional information about the current state of this Data Fusion instance if available.
	StateMessage string `pulumi:"stateMessage"`
	// The name of the tenant project.
	TenantProjectId string `pulumi:"tenantProjectId"`
	// Instance type.
	Type string `pulumi:"type"`
	// The time the instance was last updated.
	UpdateTime string `pulumi:"updateTime"`
	// Current version of the Data Fusion. Only specifiable in Update.
	Version string `pulumi:"version"`
	// Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
	Zone string `pulumi:"zone"`
}

func LookupInstanceOutput(ctx *pulumi.Context, args LookupInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceResult, error) {
			args := v.(LookupInstanceArgs)
			r, err := LookupInstance(ctx, &args, opts...)
			var s LookupInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceResultOutput)
}

type LookupInstanceOutputArgs struct {
	InstanceId pulumi.StringInput    `pulumi:"instanceId"`
	Location   pulumi.StringInput    `pulumi:"location"`
	Project    pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceArgs)(nil)).Elem()
}

type LookupInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceResult)(nil)).Elem()
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutput() LookupInstanceResultOutput {
	return o
}

func (o LookupInstanceResultOutput) ToLookupInstanceResultOutputWithContext(ctx context.Context) LookupInstanceResultOutput {
	return o
}

// List of accelerators enabled for this CDF instance.
func (o LookupInstanceResultOutput) Accelerators() AcceleratorResponseArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []AcceleratorResponse { return v.Accelerators }).(AcceleratorResponseArrayOutput)
}

// Endpoint on which the REST APIs is accessible.
func (o LookupInstanceResultOutput) ApiEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.ApiEndpoint }).(pulumi.StringOutput)
}

// Available versions that the instance can be upgraded to using UpdateInstanceRequest.
func (o LookupInstanceResultOutput) AvailableVersion() VersionResponseArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []VersionResponse { return v.AvailableVersion }).(VersionResponseArrayOutput)
}

// The time the instance was created.
func (o LookupInstanceResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The crypto key configuration. This field is used by the Customer-Managed Encryption Keys (CMEK) feature.
func (o LookupInstanceResultOutput) CryptoKeyConfig() CryptoKeyConfigResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) CryptoKeyConfigResponse { return v.CryptoKeyConfig }).(CryptoKeyConfigResponseOutput)
}

// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.
func (o LookupInstanceResultOutput) DataprocServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.DataprocServiceAccount }).(pulumi.StringOutput)
}

// A description of this instance.
func (o LookupInstanceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Description }).(pulumi.StringOutput)
}

// If the instance state is DISABLED, the reason for disabling the instance.
func (o LookupInstanceResultOutput) DisabledReason() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceResult) []string { return v.DisabledReason }).(pulumi.StringArrayOutput)
}

// Display name for an instance.
func (o LookupInstanceResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Option to enable granular role-based access control.
func (o LookupInstanceResultOutput) EnableRbac() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.EnableRbac }).(pulumi.BoolOutput)
}

// Option to enable Stackdriver Logging.
func (o LookupInstanceResultOutput) EnableStackdriverLogging() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.EnableStackdriverLogging }).(pulumi.BoolOutput)
}

// Option to enable Stackdriver Monitoring.
func (o LookupInstanceResultOutput) EnableStackdriverMonitoring() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.EnableStackdriverMonitoring }).(pulumi.BoolOutput)
}

// Option to enable granular zone separation.
func (o LookupInstanceResultOutput) EnableZoneSeparation() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.EnableZoneSeparation }).(pulumi.BoolOutput)
}

// Option to enable and pass metadata for event publishing.
func (o LookupInstanceResultOutput) EventPublishConfig() EventPublishConfigResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) EventPublishConfigResponse { return v.EventPublishConfig }).(EventPublishConfigResponseOutput)
}

// Cloud Storage bucket generated by Data Fusion in the customer project.
func (o LookupInstanceResultOutput) GcsBucket() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.GcsBucket }).(pulumi.StringOutput)
}

// The resource labels for instance to use to annotate any related underlying resources such as Compute Engine VMs. The character '=' is not allowed to be used within the labels.
func (o LookupInstanceResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The name of this instance is in the form of projects/{project}/locations/{location}/instances/{instance}.
func (o LookupInstanceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Network configuration options. These are required when a private Data Fusion instance is to be created.
func (o LookupInstanceResultOutput) NetworkConfig() NetworkConfigResponseOutput {
	return o.ApplyT(func(v LookupInstanceResult) NetworkConfigResponse { return v.NetworkConfig }).(NetworkConfigResponseOutput)
}

// Map of additional options used to configure the behavior of Data Fusion instance.
func (o LookupInstanceResultOutput) Options() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInstanceResult) map[string]string { return v.Options }).(pulumi.StringMapOutput)
}

// P4 service account for the customer project.
func (o LookupInstanceResultOutput) P4ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.P4ServiceAccount }).(pulumi.StringOutput)
}

// Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
func (o LookupInstanceResultOutput) PrivateInstance() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.PrivateInstance }).(pulumi.BoolOutput)
}

// Reserved for future use.
func (o LookupInstanceResultOutput) SatisfiesPzs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInstanceResult) bool { return v.SatisfiesPzs }).(pulumi.BoolOutput)
}

// Deprecated. Use tenant_project_id instead to extract the tenant project ID.
//
// Deprecated: Output only. Deprecated. Use tenant_project_id instead to extract the tenant project ID.
func (o LookupInstanceResultOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.ServiceAccount }).(pulumi.StringOutput)
}

// Endpoint on which the Data Fusion UI is accessible.
func (o LookupInstanceResultOutput) ServiceEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.ServiceEndpoint }).(pulumi.StringOutput)
}

// The current state of this Data Fusion instance.
func (o LookupInstanceResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.State }).(pulumi.StringOutput)
}

// Additional information about the current state of this Data Fusion instance if available.
func (o LookupInstanceResultOutput) StateMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.StateMessage }).(pulumi.StringOutput)
}

// The name of the tenant project.
func (o LookupInstanceResultOutput) TenantProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.TenantProjectId }).(pulumi.StringOutput)
}

// Instance type.
func (o LookupInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

// The time the instance was last updated.
func (o LookupInstanceResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

// Current version of the Data Fusion. Only specifiable in Update.
func (o LookupInstanceResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Version }).(pulumi.StringOutput)
}

// Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
func (o LookupInstanceResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceResultOutput{})
}
