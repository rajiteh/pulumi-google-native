// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1.Inputs
{

    /// <summary>
    /// Describes JSON data format.
    /// </summary>
    public sealed class GoogleCloudDataplexV1StorageFormatJsonOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The character encoding of the data. Accepts "US-ASCII", "UTF-8" and "ISO-8859-1". Defaults to UTF-8 if not specified.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        public GoogleCloudDataplexV1StorageFormatJsonOptionsArgs()
        {
        }
        public static new GoogleCloudDataplexV1StorageFormatJsonOptionsArgs Empty => new GoogleCloudDataplexV1StorageFormatJsonOptionsArgs();
    }
}
