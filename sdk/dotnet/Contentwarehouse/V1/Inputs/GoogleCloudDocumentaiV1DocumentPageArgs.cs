// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contentwarehouse.V1.Inputs
{

    /// <summary>
    /// A page in a Document.
    /// </summary>
    public sealed class GoogleCloudDocumentaiV1DocumentPageArgs : global::Pulumi.ResourceArgs
    {
        [Input("blocks")]
        private InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageBlockArgs>? _blocks;

        /// <summary>
        /// A list of visually detected text blocks on the page. A block has a set of lines (collected into paragraphs) that have a common line-spacing and orientation.
        /// </summary>
        public InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageBlockArgs> Blocks
        {
            get => _blocks ?? (_blocks = new InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageBlockArgs>());
            set => _blocks = value;
        }

        [Input("detectedBarcodes")]
        private InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageDetectedBarcodeArgs>? _detectedBarcodes;

        /// <summary>
        /// A list of detected barcodes.
        /// </summary>
        public InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageDetectedBarcodeArgs> DetectedBarcodes
        {
            get => _detectedBarcodes ?? (_detectedBarcodes = new InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageDetectedBarcodeArgs>());
            set => _detectedBarcodes = value;
        }

        [Input("detectedLanguages")]
        private InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageDetectedLanguageArgs>? _detectedLanguages;

        /// <summary>
        /// A list of detected languages together with confidence.
        /// </summary>
        public InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageDetectedLanguageArgs> DetectedLanguages
        {
            get => _detectedLanguages ?? (_detectedLanguages = new InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageDetectedLanguageArgs>());
            set => _detectedLanguages = value;
        }

        /// <summary>
        /// Physical dimension of the page.
        /// </summary>
        [Input("dimension")]
        public Input<Inputs.GoogleCloudDocumentaiV1DocumentPageDimensionArgs>? Dimension { get; set; }

        [Input("formFields")]
        private InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageFormFieldArgs>? _formFields;

        /// <summary>
        /// A list of visually detected form fields on the page.
        /// </summary>
        public InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageFormFieldArgs> FormFields
        {
            get => _formFields ?? (_formFields = new InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageFormFieldArgs>());
            set => _formFields = value;
        }

        /// <summary>
        /// Rendered image for this page. This image is preprocessed to remove any skew, rotation, and distortions such that the annotation bounding boxes can be upright and axis-aligned.
        /// </summary>
        [Input("image")]
        public Input<Inputs.GoogleCloudDocumentaiV1DocumentPageImageArgs>? Image { get; set; }

        /// <summary>
        /// Image Quality Scores.
        /// </summary>
        [Input("imageQualityScores")]
        public Input<Inputs.GoogleCloudDocumentaiV1DocumentPageImageQualityScoresArgs>? ImageQualityScores { get; set; }

        /// <summary>
        /// Layout for the page.
        /// </summary>
        [Input("layout")]
        public Input<Inputs.GoogleCloudDocumentaiV1DocumentPageLayoutArgs>? Layout { get; set; }

        [Input("lines")]
        private InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageLineArgs>? _lines;

        /// <summary>
        /// A list of visually detected text lines on the page. A collection of tokens that a human would perceive as a line.
        /// </summary>
        public InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageLineArgs> Lines
        {
            get => _lines ?? (_lines = new InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageLineArgs>());
            set => _lines = value;
        }

        /// <summary>
        /// 1-based index for current Page in a parent Document. Useful when a page is taken out of a Document for individual processing.
        /// </summary>
        [Input("pageNumber")]
        public Input<int>? PageNumber { get; set; }

        [Input("paragraphs")]
        private InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageParagraphArgs>? _paragraphs;

        /// <summary>
        /// A list of visually detected text paragraphs on the page. A collection of lines that a human would perceive as a paragraph.
        /// </summary>
        public InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageParagraphArgs> Paragraphs
        {
            get => _paragraphs ?? (_paragraphs = new InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageParagraphArgs>());
            set => _paragraphs = value;
        }

        /// <summary>
        /// The history of this page.
        /// </summary>
        [Input("provenance")]
        public Input<Inputs.GoogleCloudDocumentaiV1DocumentProvenanceArgs>? Provenance { get; set; }

        [Input("symbols")]
        private InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageSymbolArgs>? _symbols;

        /// <summary>
        /// A list of visually detected symbols on the page.
        /// </summary>
        public InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageSymbolArgs> Symbols
        {
            get => _symbols ?? (_symbols = new InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageSymbolArgs>());
            set => _symbols = value;
        }

        [Input("tables")]
        private InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageTableArgs>? _tables;

        /// <summary>
        /// A list of visually detected tables on the page.
        /// </summary>
        public InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageTableArgs> Tables
        {
            get => _tables ?? (_tables = new InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageTableArgs>());
            set => _tables = value;
        }

        [Input("tokens")]
        private InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageTokenArgs>? _tokens;

        /// <summary>
        /// A list of visually detected tokens on the page.
        /// </summary>
        public InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageTokenArgs> Tokens
        {
            get => _tokens ?? (_tokens = new InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageTokenArgs>());
            set => _tokens = value;
        }

        [Input("transforms")]
        private InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageMatrixArgs>? _transforms;

        /// <summary>
        /// Transformation matrices that were applied to the original document image to produce Page.image.
        /// </summary>
        public InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageMatrixArgs> Transforms
        {
            get => _transforms ?? (_transforms = new InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageMatrixArgs>());
            set => _transforms = value;
        }

        [Input("visualElements")]
        private InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageVisualElementArgs>? _visualElements;

        /// <summary>
        /// A list of detected non-text visual elements e.g. checkbox, signature etc. on the page.
        /// </summary>
        public InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageVisualElementArgs> VisualElements
        {
            get => _visualElements ?? (_visualElements = new InputList<Inputs.GoogleCloudDocumentaiV1DocumentPageVisualElementArgs>());
            set => _visualElements = value;
        }

        public GoogleCloudDocumentaiV1DocumentPageArgs()
        {
        }
        public static new GoogleCloudDocumentaiV1DocumentPageArgs Empty => new GoogleCloudDocumentaiV1DocumentPageArgs();
    }
}
