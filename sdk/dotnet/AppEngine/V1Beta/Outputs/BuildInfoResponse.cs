// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1Beta.Outputs
{

    /// <summary>
    /// Google Cloud Build information.
    /// </summary>
    [OutputType]
    public sealed class BuildInfoResponse
    {
        /// <summary>
        /// The Google Cloud Build id. Example: "f966068f-08b2-42c8-bdfe-74137dff2bf9"
        /// </summary>
        public readonly string CloudBuildId;

        [OutputConstructor]
        private BuildInfoResponse(string cloudBuildId)
        {
            CloudBuildId = cloudBuildId;
        }
    }
}
