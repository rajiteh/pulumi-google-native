// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V2.Outputs
{

    /// <summary>
    /// Represents the observed state of a single `TrafficTarget` entry.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRunV2TrafficTargetStatusResponse
    {
        /// <summary>
        /// Specifies percent of the traffic to this Revision.
        /// </summary>
        public readonly int Percent;
        /// <summary>
        /// Revision to which this traffic is sent.
        /// </summary>
        public readonly string Revision;
        /// <summary>
        /// Indicates the string used in the URI to exclusively reference this target.
        /// </summary>
        public readonly string Tag;
        /// <summary>
        /// The allocation type for this traffic target.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Displays the target URI.
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private GoogleCloudRunV2TrafficTargetStatusResponse(
            int percent,

            string revision,

            string tag,

            string type,

            string uri)
        {
            Percent = percent;
            Revision = revision;
            Tag = tag;
            Type = type;
            Uri = uri;
        }
    }
}
