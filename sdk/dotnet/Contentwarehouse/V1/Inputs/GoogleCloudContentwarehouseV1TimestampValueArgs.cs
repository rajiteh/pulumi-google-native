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
    /// Timestamp value type.
    /// </summary>
    public sealed class GoogleCloudContentwarehouseV1TimestampValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The string must represent a valid instant in UTC and is parsed using java.time.format.DateTimeFormatter.ISO_INSTANT. e.g. "2013-09-29T18:46:19Z"
        /// </summary>
        [Input("textValue")]
        public Input<string>? TextValue { get; set; }

        /// <summary>
        /// Timestamp value
        /// </summary>
        [Input("timestampValue")]
        public Input<string>? TimestampValue { get; set; }

        public GoogleCloudContentwarehouseV1TimestampValueArgs()
        {
        }
        public static new GoogleCloudContentwarehouseV1TimestampValueArgs Empty => new GoogleCloudContentwarehouseV1TimestampValueArgs();
    }
}
