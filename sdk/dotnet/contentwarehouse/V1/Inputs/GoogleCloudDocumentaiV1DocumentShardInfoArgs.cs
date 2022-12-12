// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.contentwarehouse.V1.Inputs
{

    /// <summary>
    /// For a large document, sharding may be performed to produce several document shards. Each document shard contains this field to detail which shard it is.
    /// </summary>
    public sealed class GoogleCloudDocumentaiV1DocumentShardInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Total number of shards.
        /// </summary>
        [Input("shardCount")]
        public Input<string>? ShardCount { get; set; }

        /// <summary>
        /// The 0-based index of this shard.
        /// </summary>
        [Input("shardIndex")]
        public Input<string>? ShardIndex { get; set; }

        /// <summary>
        /// The index of the first character in Document.text in the overall document global text.
        /// </summary>
        [Input("textOffset")]
        public Input<string>? TextOffset { get; set; }

        public GoogleCloudDocumentaiV1DocumentShardInfoArgs()
        {
        }
        public static new GoogleCloudDocumentaiV1DocumentShardInfoArgs Empty => new GoogleCloudDocumentaiV1DocumentShardInfoArgs();
    }
}