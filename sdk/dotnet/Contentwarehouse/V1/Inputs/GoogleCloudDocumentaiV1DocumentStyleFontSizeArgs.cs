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
    /// Font size with unit.
    /// </summary>
    public sealed class GoogleCloudDocumentaiV1DocumentStyleFontSizeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Font size for the text.
        /// </summary>
        [Input("size")]
        public Input<double>? Size { get; set; }

        /// <summary>
        /// Unit for the font size. Follows CSS naming (in, px, pt, etc.).
        /// </summary>
        [Input("unit")]
        public Input<string>? Unit { get; set; }

        public GoogleCloudDocumentaiV1DocumentStyleFontSizeArgs()
        {
        }
        public static new GoogleCloudDocumentaiV1DocumentStyleFontSizeArgs Empty => new GoogleCloudDocumentaiV1DocumentStyleFontSizeArgs();
    }
}
