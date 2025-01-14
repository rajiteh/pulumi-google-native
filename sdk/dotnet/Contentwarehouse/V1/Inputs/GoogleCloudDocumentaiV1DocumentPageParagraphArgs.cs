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
    /// A collection of lines that a human would perceive as a paragraph.
    /// </summary>
    public sealed class GoogleCloudDocumentaiV1DocumentPageParagraphArgs : global::Pulumi.ResourceArgs
    {
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
        /// Layout for Paragraph.
        /// </summary>
        [Input("layout")]
        public Input<Inputs.GoogleCloudDocumentaiV1DocumentPageLayoutArgs>? Layout { get; set; }

        /// <summary>
        /// The history of this annotation.
        /// </summary>
        [Input("provenance")]
        public Input<Inputs.GoogleCloudDocumentaiV1DocumentProvenanceArgs>? Provenance { get; set; }

        public GoogleCloudDocumentaiV1DocumentPageParagraphArgs()
        {
        }
        public static new GoogleCloudDocumentaiV1DocumentPageParagraphArgs Empty => new GoogleCloudDocumentaiV1DocumentPageParagraphArgs();
    }
}
