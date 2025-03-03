// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a content resource.
func LookupContentitem(ctx *pulumi.Context, args *LookupContentitemArgs, opts ...pulumi.InvokeOption) (*LookupContentitemResult, error) {
	var rv LookupContentitemResult
	err := ctx.Invoke("google-native:dataplex/v1:getContentitem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContentitemArgs struct {
	ContentitemId string  `pulumi:"contentitemId"`
	LakeId        string  `pulumi:"lakeId"`
	Location      string  `pulumi:"location"`
	Project       *string `pulumi:"project"`
	View          *string `pulumi:"view"`
}

type LookupContentitemResult struct {
	// Content creation time.
	CreateTime string `pulumi:"createTime"`
	// Content data in string format.
	DataText string `pulumi:"dataText"`
	// Optional. Description of the content.
	Description string `pulumi:"description"`
	// Optional. User defined labels for the content.
	Labels map[string]string `pulumi:"labels"`
	// The relative resource name of the content, of the form: projects/{project_id}/locations/{location_id}/lakes/{lake_id}/content/{content_id}
	Name string `pulumi:"name"`
	// Notebook related configurations.
	Notebook GoogleCloudDataplexV1ContentNotebookResponse `pulumi:"notebook"`
	// The path for the Content file, represented as directory structure. Unique within a lake. Limited to alphanumerics, hyphens, underscores, dots and slashes.
	Path string `pulumi:"path"`
	// Sql Script related configurations.
	SqlScript GoogleCloudDataplexV1ContentSqlScriptResponse `pulumi:"sqlScript"`
	// System generated globally unique ID for the content. This ID will be different if the content is deleted and re-created with the same name.
	Uid string `pulumi:"uid"`
	// The time when the content was last updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupContentitemOutput(ctx *pulumi.Context, args LookupContentitemOutputArgs, opts ...pulumi.InvokeOption) LookupContentitemResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContentitemResult, error) {
			args := v.(LookupContentitemArgs)
			r, err := LookupContentitem(ctx, &args, opts...)
			var s LookupContentitemResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContentitemResultOutput)
}

type LookupContentitemOutputArgs struct {
	ContentitemId pulumi.StringInput    `pulumi:"contentitemId"`
	LakeId        pulumi.StringInput    `pulumi:"lakeId"`
	Location      pulumi.StringInput    `pulumi:"location"`
	Project       pulumi.StringPtrInput `pulumi:"project"`
	View          pulumi.StringPtrInput `pulumi:"view"`
}

func (LookupContentitemOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContentitemArgs)(nil)).Elem()
}

type LookupContentitemResultOutput struct{ *pulumi.OutputState }

func (LookupContentitemResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContentitemResult)(nil)).Elem()
}

func (o LookupContentitemResultOutput) ToLookupContentitemResultOutput() LookupContentitemResultOutput {
	return o
}

func (o LookupContentitemResultOutput) ToLookupContentitemResultOutputWithContext(ctx context.Context) LookupContentitemResultOutput {
	return o
}

// Content creation time.
func (o LookupContentitemResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentitemResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Content data in string format.
func (o LookupContentitemResultOutput) DataText() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentitemResult) string { return v.DataText }).(pulumi.StringOutput)
}

// Optional. Description of the content.
func (o LookupContentitemResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentitemResult) string { return v.Description }).(pulumi.StringOutput)
}

// Optional. User defined labels for the content.
func (o LookupContentitemResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupContentitemResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The relative resource name of the content, of the form: projects/{project_id}/locations/{location_id}/lakes/{lake_id}/content/{content_id}
func (o LookupContentitemResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentitemResult) string { return v.Name }).(pulumi.StringOutput)
}

// Notebook related configurations.
func (o LookupContentitemResultOutput) Notebook() GoogleCloudDataplexV1ContentNotebookResponseOutput {
	return o.ApplyT(func(v LookupContentitemResult) GoogleCloudDataplexV1ContentNotebookResponse { return v.Notebook }).(GoogleCloudDataplexV1ContentNotebookResponseOutput)
}

// The path for the Content file, represented as directory structure. Unique within a lake. Limited to alphanumerics, hyphens, underscores, dots and slashes.
func (o LookupContentitemResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentitemResult) string { return v.Path }).(pulumi.StringOutput)
}

// Sql Script related configurations.
func (o LookupContentitemResultOutput) SqlScript() GoogleCloudDataplexV1ContentSqlScriptResponseOutput {
	return o.ApplyT(func(v LookupContentitemResult) GoogleCloudDataplexV1ContentSqlScriptResponse { return v.SqlScript }).(GoogleCloudDataplexV1ContentSqlScriptResponseOutput)
}

// System generated globally unique ID for the content. This ID will be different if the content is deleted and re-created with the same name.
func (o LookupContentitemResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentitemResult) string { return v.Uid }).(pulumi.StringOutput)
}

// The time when the content was last updated.
func (o LookupContentitemResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContentitemResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContentitemResultOutput{})
}
