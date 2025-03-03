// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Outputs
{

    /// <summary>
    /// A set of properties that uniquely identify a given Docker image.
    /// </summary>
    [OutputType]
    public sealed class FingerprintResponse
    {
        /// <summary>
        /// The layer ID of the final layer in the Docker image's v1 representation.
        /// </summary>
        public readonly string V1Name;
        /// <summary>
        /// The ordered list of v2 blobs that represent a given image.
        /// </summary>
        public readonly ImmutableArray<string> V2Blob;
        /// <summary>
        /// The name of the image's v2 blobs computed via: [bottom] := v2_blobbottom := sha256(v2_blob[N] + " " + v2_name[N+1]) Only the name of the final blob is kept.
        /// </summary>
        public readonly string V2Name;

        [OutputConstructor]
        private FingerprintResponse(
            string v1Name,

            ImmutableArray<string> v2Blob,

            string v2Name)
        {
            V1Name = v1Name;
            V2Blob = v2Blob;
            V2Name = v2Name;
        }
    }
}
