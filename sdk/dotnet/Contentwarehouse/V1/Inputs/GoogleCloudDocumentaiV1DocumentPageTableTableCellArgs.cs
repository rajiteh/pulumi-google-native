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
    /// A cell representation inside the table.
    /// </summary>
    public sealed class GoogleCloudDocumentaiV1DocumentPageTableTableCellArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// How many columns this cell spans.
        /// </summary>
        [Input("colSpan")]
        public Input<int>? ColSpan { get; set; }

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
        /// Layout for TableCell.
        /// </summary>
        [Input("layout")]
        public Input<Inputs.GoogleCloudDocumentaiV1DocumentPageLayoutArgs>? Layout { get; set; }

        /// <summary>
        /// How many rows this cell spans.
        /// </summary>
        [Input("rowSpan")]
        public Input<int>? RowSpan { get; set; }

        public GoogleCloudDocumentaiV1DocumentPageTableTableCellArgs()
        {
        }
        public static new GoogleCloudDocumentaiV1DocumentPageTableTableCellArgs Empty => new GoogleCloudDocumentaiV1DocumentPageTableTableCellArgs();
    }
}
