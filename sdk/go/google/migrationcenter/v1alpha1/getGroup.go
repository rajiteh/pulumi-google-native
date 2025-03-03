// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of a group.
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	var rv LookupGroupResult
	err := ctx.Invoke("google-native:migrationcenter/v1alpha1:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGroupArgs struct {
	GroupId  string  `pulumi:"groupId"`
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
}

type LookupGroupResult struct {
	// The timestamp when the group was created.
	CreateTime string `pulumi:"createTime"`
	// The description of the resource.
	Description string `pulumi:"description"`
	// User-friendly display name.
	DisplayName string `pulumi:"displayName"`
	// Labels as key value pairs.
	Labels map[string]string `pulumi:"labels"`
	// The name of the group.
	Name string `pulumi:"name"`
	// The timestamp when the group was last updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupGroupOutput(ctx *pulumi.Context, args LookupGroupOutputArgs, opts ...pulumi.InvokeOption) LookupGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGroupResult, error) {
			args := v.(LookupGroupArgs)
			r, err := LookupGroup(ctx, &args, opts...)
			var s LookupGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGroupResultOutput)
}

type LookupGroupOutputArgs struct {
	GroupId  pulumi.StringInput    `pulumi:"groupId"`
	Location pulumi.StringInput    `pulumi:"location"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupArgs)(nil)).Elem()
}

type LookupGroupResultOutput struct{ *pulumi.OutputState }

func (LookupGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupResult)(nil)).Elem()
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutput() LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutputWithContext(ctx context.Context) LookupGroupResultOutput {
	return o
}

// The timestamp when the group was created.
func (o LookupGroupResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The description of the resource.
func (o LookupGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// User-friendly display name.
func (o LookupGroupResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Labels as key value pairs.
func (o LookupGroupResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGroupResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The name of the group.
func (o LookupGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// The timestamp when the group was last updated.
func (o LookupGroupResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGroupResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupResultOutput{})
}
