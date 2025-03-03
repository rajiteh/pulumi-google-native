// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Inputs
{

    /// <summary>
    /// Other properties of the build.
    /// </summary>
    public sealed class GrafeasV1SlsaProvenanceZeroTwoSlsaMetadataArgs : global::Pulumi.ResourceArgs
    {
        [Input("buildFinishedOn")]
        public Input<string>? BuildFinishedOn { get; set; }

        [Input("buildInvocationId")]
        public Input<string>? BuildInvocationId { get; set; }

        [Input("buildStartedOn")]
        public Input<string>? BuildStartedOn { get; set; }

        [Input("completeness")]
        public Input<Inputs.GrafeasV1SlsaProvenanceZeroTwoSlsaCompletenessArgs>? Completeness { get; set; }

        [Input("reproducible")]
        public Input<bool>? Reproducible { get; set; }

        public GrafeasV1SlsaProvenanceZeroTwoSlsaMetadataArgs()
        {
        }
        public static new GrafeasV1SlsaProvenanceZeroTwoSlsaMetadataArgs Empty => new GrafeasV1SlsaProvenanceZeroTwoSlsaMetadataArgs();
    }
}
