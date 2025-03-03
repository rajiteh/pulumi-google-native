// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns an IntegrationTemplateVersion in the specified project.
func LookupTemplatesVersion(ctx *pulumi.Context, args *LookupTemplatesVersionArgs, opts ...pulumi.InvokeOption) (*LookupTemplatesVersionResult, error) {
	var rv LookupTemplatesVersionResult
	err := ctx.Invoke("google-native:integrations/v1alpha:getTemplatesVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTemplatesVersionArgs struct {
	IntegrationtemplateId string  `pulumi:"integrationtemplateId"`
	Location              string  `pulumi:"location"`
	ProductId             string  `pulumi:"productId"`
	Project               *string `pulumi:"project"`
	VersionId             string  `pulumi:"versionId"`
}

type LookupTemplatesVersionResult struct {
	// Auto-generated.
	CreateTime string `pulumi:"createTime"`
	// Optional. Flag to disable database persistence for execution data, including event execution info, execution export info, execution metadata index and execution param index.
	DatabasePersistencePolicy string `pulumi:"databasePersistencePolicy"`
	// Optional. The templateversion description. Permitted format is alphanumeric with underscores and no spaces.
	Description string `pulumi:"description"`
	// Optional. Error Catch Task configuration for the IntegrationTemplateVersion. It's optional.
	ErrorCatcherConfigs []GoogleCloudIntegrationsV1alphaErrorCatcherConfigResponse `pulumi:"errorCatcherConfigs"`
	// Optional. The last modifier's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
	LastModifierEmail string `pulumi:"lastModifierEmail"`
	// Auto-generated primary key. Format: projects/{project}/locations/{location}/products/{product}/integrationtemplates/{integrationtemplate}/versions/{version}
	Name string `pulumi:"name"`
	// Optional. ID of the IntegrationVersion that was used to create this IntegrationTemplateVersion
	ParentIntegrationVersionId string `pulumi:"parentIntegrationVersionId"`
	// An increasing sequence that is set when a new snapshot is created.
	SnapshotNumber string `pulumi:"snapshotNumber"`
	// Optional. Generated by eventbus. User should not set it as an input.
	Status string `pulumi:"status"`
	// Optional. Task configuration for the IntegrationTemplateVersion. It's optional, but the IntegrationTemplateVersion doesn't do anything without task_configs.
	TaskConfigs []EnterpriseCrmFrontendsEventbusProtoTaskConfigResponse `pulumi:"taskConfigs"`
	// Optional. Contains a graph of tasks that will be executed before putting the event in a terminal state (SUCCEEDED/FAILED/FATAL), regardless of success or failure, similar to "finally" in code.
	Teardown EnterpriseCrmEventbusProtoTeardownResponse `pulumi:"teardown"`
	// Optional. Parameters that are expected to be passed to the IntegrationTemplateVersion when an event is triggered. This consists of all the parameters that are expected in the IntegrationTemplateVersion execution. This gives the user the ability to provide default values, add information like PII and also provide data types of each parameter.
	TemplateParameters EnterpriseCrmFrontendsEventbusProtoWorkflowParametersResponse `pulumi:"templateParameters"`
	// Optional. Trigger configurations.
	TriggerConfigs []EnterpriseCrmFrontendsEventbusProtoTriggerConfigResponse `pulumi:"triggerConfigs"`
	// Auto-generated.
	UpdateTime string `pulumi:"updateTime"`
	// Optional. A user-defined label that annotates an integration version. Typically, this is only set when the integration version is created.
	UserLabel string `pulumi:"userLabel"`
}

func LookupTemplatesVersionOutput(ctx *pulumi.Context, args LookupTemplatesVersionOutputArgs, opts ...pulumi.InvokeOption) LookupTemplatesVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTemplatesVersionResult, error) {
			args := v.(LookupTemplatesVersionArgs)
			r, err := LookupTemplatesVersion(ctx, &args, opts...)
			var s LookupTemplatesVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTemplatesVersionResultOutput)
}

type LookupTemplatesVersionOutputArgs struct {
	IntegrationtemplateId pulumi.StringInput    `pulumi:"integrationtemplateId"`
	Location              pulumi.StringInput    `pulumi:"location"`
	ProductId             pulumi.StringInput    `pulumi:"productId"`
	Project               pulumi.StringPtrInput `pulumi:"project"`
	VersionId             pulumi.StringInput    `pulumi:"versionId"`
}

func (LookupTemplatesVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplatesVersionArgs)(nil)).Elem()
}

type LookupTemplatesVersionResultOutput struct{ *pulumi.OutputState }

func (LookupTemplatesVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTemplatesVersionResult)(nil)).Elem()
}

func (o LookupTemplatesVersionResultOutput) ToLookupTemplatesVersionResultOutput() LookupTemplatesVersionResultOutput {
	return o
}

func (o LookupTemplatesVersionResultOutput) ToLookupTemplatesVersionResultOutputWithContext(ctx context.Context) LookupTemplatesVersionResultOutput {
	return o
}

// Auto-generated.
func (o LookupTemplatesVersionResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplatesVersionResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Flag to disable database persistence for execution data, including event execution info, execution export info, execution metadata index and execution param index.
func (o LookupTemplatesVersionResultOutput) DatabasePersistencePolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplatesVersionResult) string { return v.DatabasePersistencePolicy }).(pulumi.StringOutput)
}

// Optional. The templateversion description. Permitted format is alphanumeric with underscores and no spaces.
func (o LookupTemplatesVersionResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplatesVersionResult) string { return v.Description }).(pulumi.StringOutput)
}

// Optional. Error Catch Task configuration for the IntegrationTemplateVersion. It's optional.
func (o LookupTemplatesVersionResultOutput) ErrorCatcherConfigs() GoogleCloudIntegrationsV1alphaErrorCatcherConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupTemplatesVersionResult) []GoogleCloudIntegrationsV1alphaErrorCatcherConfigResponse {
		return v.ErrorCatcherConfigs
	}).(GoogleCloudIntegrationsV1alphaErrorCatcherConfigResponseArrayOutput)
}

// Optional. The last modifier's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
func (o LookupTemplatesVersionResultOutput) LastModifierEmail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplatesVersionResult) string { return v.LastModifierEmail }).(pulumi.StringOutput)
}

// Auto-generated primary key. Format: projects/{project}/locations/{location}/products/{product}/integrationtemplates/{integrationtemplate}/versions/{version}
func (o LookupTemplatesVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplatesVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional. ID of the IntegrationVersion that was used to create this IntegrationTemplateVersion
func (o LookupTemplatesVersionResultOutput) ParentIntegrationVersionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplatesVersionResult) string { return v.ParentIntegrationVersionId }).(pulumi.StringOutput)
}

// An increasing sequence that is set when a new snapshot is created.
func (o LookupTemplatesVersionResultOutput) SnapshotNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplatesVersionResult) string { return v.SnapshotNumber }).(pulumi.StringOutput)
}

// Optional. Generated by eventbus. User should not set it as an input.
func (o LookupTemplatesVersionResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplatesVersionResult) string { return v.Status }).(pulumi.StringOutput)
}

// Optional. Task configuration for the IntegrationTemplateVersion. It's optional, but the IntegrationTemplateVersion doesn't do anything without task_configs.
func (o LookupTemplatesVersionResultOutput) TaskConfigs() EnterpriseCrmFrontendsEventbusProtoTaskConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupTemplatesVersionResult) []EnterpriseCrmFrontendsEventbusProtoTaskConfigResponse {
		return v.TaskConfigs
	}).(EnterpriseCrmFrontendsEventbusProtoTaskConfigResponseArrayOutput)
}

// Optional. Contains a graph of tasks that will be executed before putting the event in a terminal state (SUCCEEDED/FAILED/FATAL), regardless of success or failure, similar to "finally" in code.
func (o LookupTemplatesVersionResultOutput) Teardown() EnterpriseCrmEventbusProtoTeardownResponseOutput {
	return o.ApplyT(func(v LookupTemplatesVersionResult) EnterpriseCrmEventbusProtoTeardownResponse { return v.Teardown }).(EnterpriseCrmEventbusProtoTeardownResponseOutput)
}

// Optional. Parameters that are expected to be passed to the IntegrationTemplateVersion when an event is triggered. This consists of all the parameters that are expected in the IntegrationTemplateVersion execution. This gives the user the ability to provide default values, add information like PII and also provide data types of each parameter.
func (o LookupTemplatesVersionResultOutput) TemplateParameters() EnterpriseCrmFrontendsEventbusProtoWorkflowParametersResponseOutput {
	return o.ApplyT(func(v LookupTemplatesVersionResult) EnterpriseCrmFrontendsEventbusProtoWorkflowParametersResponse {
		return v.TemplateParameters
	}).(EnterpriseCrmFrontendsEventbusProtoWorkflowParametersResponseOutput)
}

// Optional. Trigger configurations.
func (o LookupTemplatesVersionResultOutput) TriggerConfigs() EnterpriseCrmFrontendsEventbusProtoTriggerConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupTemplatesVersionResult) []EnterpriseCrmFrontendsEventbusProtoTriggerConfigResponse {
		return v.TriggerConfigs
	}).(EnterpriseCrmFrontendsEventbusProtoTriggerConfigResponseArrayOutput)
}

// Auto-generated.
func (o LookupTemplatesVersionResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplatesVersionResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

// Optional. A user-defined label that annotates an integration version. Typically, this is only set when the integration version is created.
func (o LookupTemplatesVersionResultOutput) UserLabel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTemplatesVersionResult) string { return v.UserLabel }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTemplatesVersionResultOutput{})
}
