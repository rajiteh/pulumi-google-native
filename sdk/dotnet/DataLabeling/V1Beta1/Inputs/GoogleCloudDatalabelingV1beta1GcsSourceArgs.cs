// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataLabeling.V1Beta1.Inputs
{

    /// <summary>
    /// Source of the Cloud Storage file to be imported.
    /// </summary>
    public sealed class GoogleCloudDatalabelingV1beta1GcsSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The input URI of source file. This must be a Cloud Storage path (`gs://...`).
        /// </summary>
        [Input("inputUri", required: true)]
        public Input<string> InputUri { get; set; } = null!;

        /// <summary>
        /// The format of the source file. Only "text/csv" is supported.
        /// </summary>
        [Input("mimeType", required: true)]
        public Input<string> MimeType { get; set; } = null!;

        public GoogleCloudDatalabelingV1beta1GcsSourceArgs()
        {
        }
        public static new GoogleCloudDatalabelingV1beta1GcsSourceArgs Empty => new GoogleCloudDatalabelingV1beta1GcsSourceArgs();
    }
}
