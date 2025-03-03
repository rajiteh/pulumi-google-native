// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Inputs
{

    /// <summary>
    /// Other properties of the build.
    /// </summary>
    public sealed class MetadataArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The timestamp of when the build completed.
        /// </summary>
        [Input("buildFinishedOn")]
        public Input<string>? BuildFinishedOn { get; set; }

        /// <summary>
        /// Identifies the particular build invocation, which can be useful for finding associated logs or other ad-hoc analysis. The value SHOULD be globally unique, per in-toto Provenance spec.
        /// </summary>
        [Input("buildInvocationId")]
        public Input<string>? BuildInvocationId { get; set; }

        /// <summary>
        /// The timestamp of when the build started.
        /// </summary>
        [Input("buildStartedOn")]
        public Input<string>? BuildStartedOn { get; set; }

        /// <summary>
        /// Indicates that the builder claims certain fields in this message to be complete.
        /// </summary>
        [Input("completeness")]
        public Input<Inputs.CompletenessArgs>? Completeness { get; set; }

        /// <summary>
        /// If true, the builder claims that running the recipe on materials will produce bit-for-bit identical output.
        /// </summary>
        [Input("reproducible")]
        public Input<bool>? Reproducible { get; set; }

        public MetadataArgs()
        {
        }
        public static new MetadataArgs Empty => new MetadataArgs();
    }
}
