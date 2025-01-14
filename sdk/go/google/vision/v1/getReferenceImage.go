// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information associated with a ReferenceImage. Possible errors: * Returns NOT_FOUND if the specified image does not exist.
func LookupReferenceImage(ctx *pulumi.Context, args *LookupReferenceImageArgs, opts ...pulumi.InvokeOption) (*LookupReferenceImageResult, error) {
	var rv LookupReferenceImageResult
	err := ctx.Invoke("google-native:vision/v1:getReferenceImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReferenceImageArgs struct {
	Location         string  `pulumi:"location"`
	ProductId        string  `pulumi:"productId"`
	Project          *string `pulumi:"project"`
	ReferenceImageId string  `pulumi:"referenceImageId"`
}

type LookupReferenceImageResult struct {
	// Optional. Bounding polygons around the areas of interest in the reference image. If this field is empty, the system will try to detect regions of interest. At most 10 bounding polygons will be used. The provided shape is converted into a non-rotated rectangle. Once converted, the small edge of the rectangle must be greater than or equal to 300 pixels. The aspect ratio must be 1:4 or less (i.e. 1:3 is ok; 1:5 is not).
	BoundingPolys []BoundingPolyResponse `pulumi:"boundingPolys"`
	// The resource name of the reference image. Format is: `projects/PROJECT_ID/locations/LOC_ID/products/PRODUCT_ID/referenceImages/IMAGE_ID`. This field is ignored when creating a reference image.
	Name string `pulumi:"name"`
	// The Google Cloud Storage URI of the reference image. The URI must start with `gs://`.
	Uri string `pulumi:"uri"`
}

func LookupReferenceImageOutput(ctx *pulumi.Context, args LookupReferenceImageOutputArgs, opts ...pulumi.InvokeOption) LookupReferenceImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReferenceImageResult, error) {
			args := v.(LookupReferenceImageArgs)
			r, err := LookupReferenceImage(ctx, &args, opts...)
			var s LookupReferenceImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReferenceImageResultOutput)
}

type LookupReferenceImageOutputArgs struct {
	Location         pulumi.StringInput    `pulumi:"location"`
	ProductId        pulumi.StringInput    `pulumi:"productId"`
	Project          pulumi.StringPtrInput `pulumi:"project"`
	ReferenceImageId pulumi.StringInput    `pulumi:"referenceImageId"`
}

func (LookupReferenceImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReferenceImageArgs)(nil)).Elem()
}

type LookupReferenceImageResultOutput struct{ *pulumi.OutputState }

func (LookupReferenceImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReferenceImageResult)(nil)).Elem()
}

func (o LookupReferenceImageResultOutput) ToLookupReferenceImageResultOutput() LookupReferenceImageResultOutput {
	return o
}

func (o LookupReferenceImageResultOutput) ToLookupReferenceImageResultOutputWithContext(ctx context.Context) LookupReferenceImageResultOutput {
	return o
}

// Optional. Bounding polygons around the areas of interest in the reference image. If this field is empty, the system will try to detect regions of interest. At most 10 bounding polygons will be used. The provided shape is converted into a non-rotated rectangle. Once converted, the small edge of the rectangle must be greater than or equal to 300 pixels. The aspect ratio must be 1:4 or less (i.e. 1:3 is ok; 1:5 is not).
func (o LookupReferenceImageResultOutput) BoundingPolys() BoundingPolyResponseArrayOutput {
	return o.ApplyT(func(v LookupReferenceImageResult) []BoundingPolyResponse { return v.BoundingPolys }).(BoundingPolyResponseArrayOutput)
}

// The resource name of the reference image. Format is: `projects/PROJECT_ID/locations/LOC_ID/products/PRODUCT_ID/referenceImages/IMAGE_ID`. This field is ignored when creating a reference image.
func (o LookupReferenceImageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReferenceImageResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Google Cloud Storage URI of the reference image. The URI must start with `gs://`.
func (o LookupReferenceImageResultOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReferenceImageResult) string { return v.Uri }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReferenceImageResultOutput{})
}
