// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets an attachment. **Note:** Not supported for Apigee hybrid.
func LookupInstanceAttachment(ctx *pulumi.Context, args *LookupInstanceAttachmentArgs, opts ...pulumi.InvokeOption) (*LookupInstanceAttachmentResult, error) {
	var rv LookupInstanceAttachmentResult
	err := ctx.Invoke("google-native:apigee/v1:getInstanceAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceAttachmentArgs struct {
	AttachmentId   string `pulumi:"attachmentId"`
	InstanceId     string `pulumi:"instanceId"`
	OrganizationId string `pulumi:"organizationId"`
}

type LookupInstanceAttachmentResult struct {
	// Time the attachment was created in milliseconds since epoch.
	CreatedAt string `pulumi:"createdAt"`
	// ID of the attached environment.
	Environment string `pulumi:"environment"`
	// ID of the attachment.
	Name string `pulumi:"name"`
}

func LookupInstanceAttachmentOutput(ctx *pulumi.Context, args LookupInstanceAttachmentOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceAttachmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceAttachmentResult, error) {
			args := v.(LookupInstanceAttachmentArgs)
			r, err := LookupInstanceAttachment(ctx, &args, opts...)
			var s LookupInstanceAttachmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupInstanceAttachmentResultOutput)
}

type LookupInstanceAttachmentOutputArgs struct {
	AttachmentId   pulumi.StringInput `pulumi:"attachmentId"`
	InstanceId     pulumi.StringInput `pulumi:"instanceId"`
	OrganizationId pulumi.StringInput `pulumi:"organizationId"`
}

func (LookupInstanceAttachmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceAttachmentArgs)(nil)).Elem()
}

type LookupInstanceAttachmentResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceAttachmentResult)(nil)).Elem()
}

func (o LookupInstanceAttachmentResultOutput) ToLookupInstanceAttachmentResultOutput() LookupInstanceAttachmentResultOutput {
	return o
}

func (o LookupInstanceAttachmentResultOutput) ToLookupInstanceAttachmentResultOutputWithContext(ctx context.Context) LookupInstanceAttachmentResultOutput {
	return o
}

// Time the attachment was created in milliseconds since epoch.
func (o LookupInstanceAttachmentResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceAttachmentResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// ID of the attached environment.
func (o LookupInstanceAttachmentResultOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceAttachmentResult) string { return v.Environment }).(pulumi.StringOutput)
}

// ID of the attachment.
func (o LookupInstanceAttachmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInstanceAttachmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceAttachmentResultOutput{})
}
