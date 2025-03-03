// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single Connection.
func LookupConnection(ctx *pulumi.Context, args *LookupConnectionArgs, opts ...pulumi.InvokeOption) (*LookupConnectionResult, error) {
	var rv LookupConnectionResult
	err := ctx.Invoke("google-native:connectors/v1:getConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectionArgs struct {
	ConnectionId string  `pulumi:"connectionId"`
	Location     string  `pulumi:"location"`
	Project      *string `pulumi:"project"`
	View         *string `pulumi:"view"`
}

type LookupConnectionResult struct {
	// Optional. Configuration for establishing the connection's authentication with an external system.
	AuthConfig AuthConfigResponse `pulumi:"authConfig"`
	// Optional. Configuration for configuring the connection with an external system.
	ConfigVariables []ConfigVariableResponse `pulumi:"configVariables"`
	// Connector version on which the connection is created. The format is: projects/*/locations/*/providers/*/connectors/*/versions/* Only global location is supported for ConnectorVersion resource.
	ConnectorVersion string `pulumi:"connectorVersion"`
	// Created time.
	CreateTime string `pulumi:"createTime"`
	// Optional. Description of the resource.
	Description string `pulumi:"description"`
	// Optional. Configuration of the Connector's destination. Only accepted for Connectors that accepts user defined destination(s).
	DestinationConfigs []DestinationConfigResponse `pulumi:"destinationConfigs"`
	// GCR location where the envoy image is stored. formatted like: gcr.io/{bucketName}/{imageName}
	EnvoyImageLocation string `pulumi:"envoyImageLocation"`
	// GCR location where the runtime image is stored. formatted like: gcr.io/{bucketName}/{imageName}
	ImageLocation string `pulumi:"imageLocation"`
	// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
	Labels map[string]string `pulumi:"labels"`
	// Optional. Configuration that indicates whether or not the Connection can be edited.
	LockConfig LockConfigResponse `pulumi:"lockConfig"`
	// Optional. Log configuration for the connection.
	LogConfig ConnectorsLogConfigResponse `pulumi:"logConfig"`
	// Resource name of the Connection. Format: projects/{project}/locations/{location}/connections/{connection}
	Name string `pulumi:"name"`
	// Optional. Node configuration for the connection.
	NodeConfig NodeConfigResponse `pulumi:"nodeConfig"`
	// Optional. Service account needed for runtime plane to access GCP resources.
	ServiceAccount string `pulumi:"serviceAccount"`
	// The name of the Service Directory service name. Used for Private Harpoon to resolve the ILB address. e.g. "projects/cloud-connectors-e2e-testing/locations/us-central1/namespaces/istio-system/services/istio-ingressgateway-connectors"
	ServiceDirectory string `pulumi:"serviceDirectory"`
	// Optional. Ssl config of a connection
	SslConfig SslConfigResponse `pulumi:"sslConfig"`
	// Current status of the connection.
	Status ConnectionStatusResponse `pulumi:"status"`
	// This subscription type enum states the subscription type of the project.
	SubscriptionType string `pulumi:"subscriptionType"`
	// Optional. Suspended indicates if a user has suspended a connection or not.
	Suspended bool `pulumi:"suspended"`
	// Updated time.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupConnectionOutput(ctx *pulumi.Context, args LookupConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectionResult, error) {
			args := v.(LookupConnectionArgs)
			r, err := LookupConnection(ctx, &args, opts...)
			var s LookupConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectionResultOutput)
}

type LookupConnectionOutputArgs struct {
	ConnectionId pulumi.StringInput    `pulumi:"connectionId"`
	Location     pulumi.StringInput    `pulumi:"location"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
	View         pulumi.StringPtrInput `pulumi:"view"`
}

func (LookupConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionArgs)(nil)).Elem()
}

type LookupConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionResult)(nil)).Elem()
}

func (o LookupConnectionResultOutput) ToLookupConnectionResultOutput() LookupConnectionResultOutput {
	return o
}

func (o LookupConnectionResultOutput) ToLookupConnectionResultOutputWithContext(ctx context.Context) LookupConnectionResultOutput {
	return o
}

// Optional. Configuration for establishing the connection's authentication with an external system.
func (o LookupConnectionResultOutput) AuthConfig() AuthConfigResponseOutput {
	return o.ApplyT(func(v LookupConnectionResult) AuthConfigResponse { return v.AuthConfig }).(AuthConfigResponseOutput)
}

// Optional. Configuration for configuring the connection with an external system.
func (o LookupConnectionResultOutput) ConfigVariables() ConfigVariableResponseArrayOutput {
	return o.ApplyT(func(v LookupConnectionResult) []ConfigVariableResponse { return v.ConfigVariables }).(ConfigVariableResponseArrayOutput)
}

// Connector version on which the connection is created. The format is: projects/*/locations/*/providers/*/connectors/*/versions/* Only global location is supported for ConnectorVersion resource.
func (o LookupConnectionResultOutput) ConnectorVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.ConnectorVersion }).(pulumi.StringOutput)
}

// Created time.
func (o LookupConnectionResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Description of the resource.
func (o LookupConnectionResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.Description }).(pulumi.StringOutput)
}

// Optional. Configuration of the Connector's destination. Only accepted for Connectors that accepts user defined destination(s).
func (o LookupConnectionResultOutput) DestinationConfigs() DestinationConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupConnectionResult) []DestinationConfigResponse { return v.DestinationConfigs }).(DestinationConfigResponseArrayOutput)
}

// GCR location where the envoy image is stored. formatted like: gcr.io/{bucketName}/{imageName}
func (o LookupConnectionResultOutput) EnvoyImageLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.EnvoyImageLocation }).(pulumi.StringOutput)
}

// GCR location where the runtime image is stored. formatted like: gcr.io/{bucketName}/{imageName}
func (o LookupConnectionResultOutput) ImageLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.ImageLocation }).(pulumi.StringOutput)
}

// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
func (o LookupConnectionResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectionResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Optional. Configuration that indicates whether or not the Connection can be edited.
func (o LookupConnectionResultOutput) LockConfig() LockConfigResponseOutput {
	return o.ApplyT(func(v LookupConnectionResult) LockConfigResponse { return v.LockConfig }).(LockConfigResponseOutput)
}

// Optional. Log configuration for the connection.
func (o LookupConnectionResultOutput) LogConfig() ConnectorsLogConfigResponseOutput {
	return o.ApplyT(func(v LookupConnectionResult) ConnectorsLogConfigResponse { return v.LogConfig }).(ConnectorsLogConfigResponseOutput)
}

// Resource name of the Connection. Format: projects/{project}/locations/{location}/connections/{connection}
func (o LookupConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional. Node configuration for the connection.
func (o LookupConnectionResultOutput) NodeConfig() NodeConfigResponseOutput {
	return o.ApplyT(func(v LookupConnectionResult) NodeConfigResponse { return v.NodeConfig }).(NodeConfigResponseOutput)
}

// Optional. Service account needed for runtime plane to access GCP resources.
func (o LookupConnectionResultOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.ServiceAccount }).(pulumi.StringOutput)
}

// The name of the Service Directory service name. Used for Private Harpoon to resolve the ILB address. e.g. "projects/cloud-connectors-e2e-testing/locations/us-central1/namespaces/istio-system/services/istio-ingressgateway-connectors"
func (o LookupConnectionResultOutput) ServiceDirectory() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.ServiceDirectory }).(pulumi.StringOutput)
}

// Optional. Ssl config of a connection
func (o LookupConnectionResultOutput) SslConfig() SslConfigResponseOutput {
	return o.ApplyT(func(v LookupConnectionResult) SslConfigResponse { return v.SslConfig }).(SslConfigResponseOutput)
}

// Current status of the connection.
func (o LookupConnectionResultOutput) Status() ConnectionStatusResponseOutput {
	return o.ApplyT(func(v LookupConnectionResult) ConnectionStatusResponse { return v.Status }).(ConnectionStatusResponseOutput)
}

// This subscription type enum states the subscription type of the project.
func (o LookupConnectionResultOutput) SubscriptionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.SubscriptionType }).(pulumi.StringOutput)
}

// Optional. Suspended indicates if a user has suspended a connection or not.
func (o LookupConnectionResultOutput) Suspended() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupConnectionResult) bool { return v.Suspended }).(pulumi.BoolOutput)
}

// Updated time.
func (o LookupConnectionResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectionResultOutput{})
}
