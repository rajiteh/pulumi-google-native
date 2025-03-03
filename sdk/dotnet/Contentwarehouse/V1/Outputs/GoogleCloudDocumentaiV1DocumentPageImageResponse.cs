// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contentwarehouse.V1.Outputs
{

    /// <summary>
    /// Rendered image contents for this page.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDocumentaiV1DocumentPageImageResponse
    {
        /// <summary>
        /// Raw byte content of the image.
        /// </summary>
        public readonly string Content;
        /// <summary>
        /// Height of the image in pixels.
        /// </summary>
        public readonly int Height;
        /// <summary>
        /// Encoding mime type for the image.
        /// </summary>
        public readonly string MimeType;
        /// <summary>
        /// Width of the image in pixels.
        /// </summary>
        public readonly int Width;

        [OutputConstructor]
        private GoogleCloudDocumentaiV1DocumentPageImageResponse(
            string content,

            int height,

            string mimeType,

            int width)
        {
            Content = content;
            Height = height;
            MimeType = mimeType;
            Width = width;
        }
    }
}
