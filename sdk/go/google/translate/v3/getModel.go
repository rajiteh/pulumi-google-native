// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a model.
func LookupModel(ctx *pulumi.Context, args *LookupModelArgs, opts ...pulumi.InvokeOption) (*LookupModelResult, error) {
	var rv LookupModelResult
	err := ctx.Invoke("google-native:translate/v3:getModel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupModelArgs struct {
	Location string  `pulumi:"location"`
	ModelId  string  `pulumi:"modelId"`
	Project  *string `pulumi:"project"`
}

type LookupModelResult struct {
	// Timestamp when the model resource was created, which is also when the training started.
	CreateTime string `pulumi:"createTime"`
	// The dataset from which the model is trained, in form of `projects/{project-number-or-id}/locations/{location_id}/datasets/{dataset_id}`
	Dataset string `pulumi:"dataset"`
	// The name of the model to show in the interface. The name can be up to 32 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscores (_), and ASCII digits 0-9.
	DisplayName string `pulumi:"displayName"`
	// The resource name of the model, in form of `projects/{project-number-or-id}/locations/{location_id}/models/{model_id}`
	Name string `pulumi:"name"`
	// The BCP-47 language code of the source language.
	SourceLanguageCode string `pulumi:"sourceLanguageCode"`
	// The BCP-47 language code of the target language.
	TargetLanguageCode string `pulumi:"targetLanguageCode"`
	// Number of examples (sentence pairs) used to test the model.
	TestExampleCount int `pulumi:"testExampleCount"`
	// Number of examples (sentence pairs) used to train the model.
	TrainExampleCount int `pulumi:"trainExampleCount"`
	// Timestamp when this model was last updated.
	UpdateTime string `pulumi:"updateTime"`
	// Number of examples (sentence pairs) used to validate the model.
	ValidateExampleCount int `pulumi:"validateExampleCount"`
}

func LookupModelOutput(ctx *pulumi.Context, args LookupModelOutputArgs, opts ...pulumi.InvokeOption) LookupModelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupModelResult, error) {
			args := v.(LookupModelArgs)
			r, err := LookupModel(ctx, &args, opts...)
			var s LookupModelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupModelResultOutput)
}

type LookupModelOutputArgs struct {
	Location pulumi.StringInput    `pulumi:"location"`
	ModelId  pulumi.StringInput    `pulumi:"modelId"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupModelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelArgs)(nil)).Elem()
}

type LookupModelResultOutput struct{ *pulumi.OutputState }

func (LookupModelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelResult)(nil)).Elem()
}

func (o LookupModelResultOutput) ToLookupModelResultOutput() LookupModelResultOutput {
	return o
}

func (o LookupModelResultOutput) ToLookupModelResultOutputWithContext(ctx context.Context) LookupModelResultOutput {
	return o
}

// Timestamp when the model resource was created, which is also when the training started.
func (o LookupModelResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The dataset from which the model is trained, in form of `projects/{project-number-or-id}/locations/{location_id}/datasets/{dataset_id}`
func (o LookupModelResultOutput) Dataset() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelResult) string { return v.Dataset }).(pulumi.StringOutput)
}

// The name of the model to show in the interface. The name can be up to 32 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscores (_), and ASCII digits 0-9.
func (o LookupModelResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The resource name of the model, in form of `projects/{project-number-or-id}/locations/{location_id}/models/{model_id}`
func (o LookupModelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelResult) string { return v.Name }).(pulumi.StringOutput)
}

// The BCP-47 language code of the source language.
func (o LookupModelResultOutput) SourceLanguageCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelResult) string { return v.SourceLanguageCode }).(pulumi.StringOutput)
}

// The BCP-47 language code of the target language.
func (o LookupModelResultOutput) TargetLanguageCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelResult) string { return v.TargetLanguageCode }).(pulumi.StringOutput)
}

// Number of examples (sentence pairs) used to test the model.
func (o LookupModelResultOutput) TestExampleCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupModelResult) int { return v.TestExampleCount }).(pulumi.IntOutput)
}

// Number of examples (sentence pairs) used to train the model.
func (o LookupModelResultOutput) TrainExampleCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupModelResult) int { return v.TrainExampleCount }).(pulumi.IntOutput)
}

// Timestamp when this model was last updated.
func (o LookupModelResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

// Number of examples (sentence pairs) used to validate the model.
func (o LookupModelResultOutput) ValidateExampleCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupModelResult) int { return v.ValidateExampleCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupModelResultOutput{})
}
