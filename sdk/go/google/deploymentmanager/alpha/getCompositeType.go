// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a specific composite type.
func LookupCompositeType(ctx *pulumi.Context, args *LookupCompositeTypeArgs, opts ...pulumi.InvokeOption) (*LookupCompositeTypeResult, error) {
	var rv LookupCompositeTypeResult
	err := ctx.Invoke("google-native:deploymentmanager/alpha:getCompositeType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCompositeTypeArgs struct {
	CompositeType string  `pulumi:"compositeType"`
	Project       *string `pulumi:"project"`
}

type LookupCompositeTypeResult struct {
	// An optional textual description of the resource; provided by the client when the resource is created.
	Description string `pulumi:"description"`
	// Creation timestamp in RFC3339 text format.
	InsertTime string `pulumi:"insertTime"`
	// Map of labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
	Labels []CompositeTypeLabelEntryResponse `pulumi:"labels"`
	// Name of the composite type, must follow the expression: `[a-z]([-a-z0-9_.]{0,61}[a-z0-9])?`.
	Name string `pulumi:"name"`
	// The Operation that most recently ran, or is currently running, on this composite type.
	Operation OperationResponse `pulumi:"operation"`
	// Server defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	Status   string `pulumi:"status"`
	// Files for the template type.
	TemplateContents TemplateContentsResponse `pulumi:"templateContents"`
}

func LookupCompositeTypeOutput(ctx *pulumi.Context, args LookupCompositeTypeOutputArgs, opts ...pulumi.InvokeOption) LookupCompositeTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCompositeTypeResult, error) {
			args := v.(LookupCompositeTypeArgs)
			r, err := LookupCompositeType(ctx, &args, opts...)
			var s LookupCompositeTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCompositeTypeResultOutput)
}

type LookupCompositeTypeOutputArgs struct {
	CompositeType pulumi.StringInput    `pulumi:"compositeType"`
	Project       pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupCompositeTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCompositeTypeArgs)(nil)).Elem()
}

type LookupCompositeTypeResultOutput struct{ *pulumi.OutputState }

func (LookupCompositeTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCompositeTypeResult)(nil)).Elem()
}

func (o LookupCompositeTypeResultOutput) ToLookupCompositeTypeResultOutput() LookupCompositeTypeResultOutput {
	return o
}

func (o LookupCompositeTypeResultOutput) ToLookupCompositeTypeResultOutputWithContext(ctx context.Context) LookupCompositeTypeResultOutput {
	return o
}

// An optional textual description of the resource; provided by the client when the resource is created.
func (o LookupCompositeTypeResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCompositeTypeResult) string { return v.Description }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupCompositeTypeResultOutput) InsertTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCompositeTypeResult) string { return v.InsertTime }).(pulumi.StringOutput)
}

// Map of labels; provided by the client when the resource is created or updated. Specifically: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: `[a-z]([-a-z0-9]*[a-z0-9])?` Label values must be between 0 and 63 characters long and must conform to the regular expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`.
func (o LookupCompositeTypeResultOutput) Labels() CompositeTypeLabelEntryResponseArrayOutput {
	return o.ApplyT(func(v LookupCompositeTypeResult) []CompositeTypeLabelEntryResponse { return v.Labels }).(CompositeTypeLabelEntryResponseArrayOutput)
}

// Name of the composite type, must follow the expression: `[a-z]([-a-z0-9_.]{0,61}[a-z0-9])?`.
func (o LookupCompositeTypeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCompositeTypeResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Operation that most recently ran, or is currently running, on this composite type.
func (o LookupCompositeTypeResultOutput) Operation() OperationResponseOutput {
	return o.ApplyT(func(v LookupCompositeTypeResult) OperationResponse { return v.Operation }).(OperationResponseOutput)
}

// Server defined URL for the resource.
func (o LookupCompositeTypeResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCompositeTypeResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func (o LookupCompositeTypeResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCompositeTypeResult) string { return v.Status }).(pulumi.StringOutput)
}

// Files for the template type.
func (o LookupCompositeTypeResultOutput) TemplateContents() TemplateContentsResponseOutput {
	return o.ApplyT(func(v LookupCompositeTypeResult) TemplateContentsResponse { return v.TemplateContents }).(TemplateContentsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCompositeTypeResultOutput{})
}
