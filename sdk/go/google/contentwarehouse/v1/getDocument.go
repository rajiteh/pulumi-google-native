// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a document. Returns NOT_FOUND if the document does not exist.
func LookupDocument(ctx *pulumi.Context, args *LookupDocumentArgs, opts ...pulumi.InvokeOption) (*LookupDocumentResult, error) {
	var rv LookupDocumentResult
	err := ctx.Invoke("google-native:contentwarehouse/v1:getDocument", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDocumentArgs struct {
	DocumentId string  `pulumi:"documentId"`
	Location   string  `pulumi:"location"`
	Project    *string `pulumi:"project"`
}

type LookupDocumentResult struct {
	// Document AI format to save the structured content, including OCR.
	CloudAiDocument GoogleCloudDocumentaiV1DocumentResponse `pulumi:"cloudAiDocument"`
	// Indicates the category (image, audio, video etc.) of the original content.
	ContentCategory string `pulumi:"contentCategory"`
	// The time when the document is created.
	CreateTime string `pulumi:"createTime"`
	// The user who creates the document.
	Creator string `pulumi:"creator"`
	// Display name of the document given by the user. This name will be displayed in the UI. Customer can populate this field with the name of the document. This differs from the 'title' field as 'title' is optional and stores the top heading in the document.
	DisplayName string `pulumi:"displayName"`
	// Uri to display the document, for example, in the UI.
	DisplayUri string `pulumi:"displayUri"`
	// If linked to a Collection with RetentionPolicy, the date when the document becomes mutable.
	DispositionTime string `pulumi:"dispositionTime"`
	// The Document schema name. Format: projects/{project_number}/locations/{location}/documentSchemas/{document_schema_id}.
	DocumentSchemaName string `pulumi:"documentSchemaName"`
	// Raw document content.
	InlineRawDocument string `pulumi:"inlineRawDocument"`
	// The resource name of the document. Format: projects/{project_number}/locations/{location}/documents/{document_id}. The name is ignored when creating a document.
	Name string `pulumi:"name"`
	// Other document format, such as PPTX, XLXS
	PlainText string `pulumi:"plainText"`
	// List of values that are user supplied metadata.
	Properties []GoogleCloudContentwarehouseV1PropertyResponse `pulumi:"properties"`
	// This is used when DocAI was not used to load the document and parsing/ extracting is needed for the inline_raw_document. For example, if inline_raw_document is the byte representation of a PDF file, then this should be set to: RAW_DOCUMENT_FILE_TYPE_PDF.
	RawDocumentFileType string `pulumi:"rawDocumentFileType"`
	// Raw document file in Cloud Storage path.
	RawDocumentPath string `pulumi:"rawDocumentPath"`
	// The reference ID set by customers. Must be unique per project and location.
	ReferenceId string `pulumi:"referenceId"`
	// If true, text extraction will not be performed.
	TextExtractionDisabled bool `pulumi:"textExtractionDisabled"`
	// If true, text extraction will be performed.
	TextExtractionEnabled bool `pulumi:"textExtractionEnabled"`
	// Title that describes the document. This can be the top heading or text that describes the document.
	Title string `pulumi:"title"`
	// The time when the document is last updated.
	UpdateTime string `pulumi:"updateTime"`
	// The user who lastly updates the document.
	Updater string `pulumi:"updater"`
}

func LookupDocumentOutput(ctx *pulumi.Context, args LookupDocumentOutputArgs, opts ...pulumi.InvokeOption) LookupDocumentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDocumentResult, error) {
			args := v.(LookupDocumentArgs)
			r, err := LookupDocument(ctx, &args, opts...)
			var s LookupDocumentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDocumentResultOutput)
}

type LookupDocumentOutputArgs struct {
	DocumentId pulumi.StringInput    `pulumi:"documentId"`
	Location   pulumi.StringInput    `pulumi:"location"`
	Project    pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupDocumentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDocumentArgs)(nil)).Elem()
}

type LookupDocumentResultOutput struct{ *pulumi.OutputState }

func (LookupDocumentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDocumentResult)(nil)).Elem()
}

func (o LookupDocumentResultOutput) ToLookupDocumentResultOutput() LookupDocumentResultOutput {
	return o
}

func (o LookupDocumentResultOutput) ToLookupDocumentResultOutputWithContext(ctx context.Context) LookupDocumentResultOutput {
	return o
}

// Document AI format to save the structured content, including OCR.
func (o LookupDocumentResultOutput) CloudAiDocument() GoogleCloudDocumentaiV1DocumentResponseOutput {
	return o.ApplyT(func(v LookupDocumentResult) GoogleCloudDocumentaiV1DocumentResponse { return v.CloudAiDocument }).(GoogleCloudDocumentaiV1DocumentResponseOutput)
}

// Indicates the category (image, audio, video etc.) of the original content.
func (o LookupDocumentResultOutput) ContentCategory() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.ContentCategory }).(pulumi.StringOutput)
}

// The time when the document is created.
func (o LookupDocumentResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The user who creates the document.
func (o LookupDocumentResultOutput) Creator() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.Creator }).(pulumi.StringOutput)
}

// Display name of the document given by the user. This name will be displayed in the UI. Customer can populate this field with the name of the document. This differs from the 'title' field as 'title' is optional and stores the top heading in the document.
func (o LookupDocumentResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Uri to display the document, for example, in the UI.
func (o LookupDocumentResultOutput) DisplayUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.DisplayUri }).(pulumi.StringOutput)
}

// If linked to a Collection with RetentionPolicy, the date when the document becomes mutable.
func (o LookupDocumentResultOutput) DispositionTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.DispositionTime }).(pulumi.StringOutput)
}

// The Document schema name. Format: projects/{project_number}/locations/{location}/documentSchemas/{document_schema_id}.
func (o LookupDocumentResultOutput) DocumentSchemaName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.DocumentSchemaName }).(pulumi.StringOutput)
}

// Raw document content.
func (o LookupDocumentResultOutput) InlineRawDocument() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.InlineRawDocument }).(pulumi.StringOutput)
}

// The resource name of the document. Format: projects/{project_number}/locations/{location}/documents/{document_id}. The name is ignored when creating a document.
func (o LookupDocumentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Other document format, such as PPTX, XLXS
func (o LookupDocumentResultOutput) PlainText() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.PlainText }).(pulumi.StringOutput)
}

// List of values that are user supplied metadata.
func (o LookupDocumentResultOutput) Properties() GoogleCloudContentwarehouseV1PropertyResponseArrayOutput {
	return o.ApplyT(func(v LookupDocumentResult) []GoogleCloudContentwarehouseV1PropertyResponse { return v.Properties }).(GoogleCloudContentwarehouseV1PropertyResponseArrayOutput)
}

// This is used when DocAI was not used to load the document and parsing/ extracting is needed for the inline_raw_document. For example, if inline_raw_document is the byte representation of a PDF file, then this should be set to: RAW_DOCUMENT_FILE_TYPE_PDF.
func (o LookupDocumentResultOutput) RawDocumentFileType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.RawDocumentFileType }).(pulumi.StringOutput)
}

// Raw document file in Cloud Storage path.
func (o LookupDocumentResultOutput) RawDocumentPath() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.RawDocumentPath }).(pulumi.StringOutput)
}

// The reference ID set by customers. Must be unique per project and location.
func (o LookupDocumentResultOutput) ReferenceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.ReferenceId }).(pulumi.StringOutput)
}

// If true, text extraction will not be performed.
func (o LookupDocumentResultOutput) TextExtractionDisabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDocumentResult) bool { return v.TextExtractionDisabled }).(pulumi.BoolOutput)
}

// If true, text extraction will be performed.
func (o LookupDocumentResultOutput) TextExtractionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDocumentResult) bool { return v.TextExtractionEnabled }).(pulumi.BoolOutput)
}

// Title that describes the document. This can be the top heading or text that describes the document.
func (o LookupDocumentResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.Title }).(pulumi.StringOutput)
}

// The time when the document is last updated.
func (o LookupDocumentResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

// The user who lastly updates the document.
func (o LookupDocumentResultOutput) Updater() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDocumentResult) string { return v.Updater }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDocumentResultOutput{})
}
