// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Composer.V1Beta1.Outputs
{

    /// <summary>
    /// Configuration for Cloud Data Lineage integration.
    /// </summary>
    [OutputType]
    public sealed class CloudDataLineageIntegrationResponse
    {
        /// <summary>
        /// Optional. Whether or not Cloud Data Lineage integration is enabled.
        /// </summary>
        public readonly bool Enabled;

        [OutputConstructor]
        private CloudDataLineageIntegrationResponse(bool enabled)
        {
            Enabled = enabled;
        }
    }
}
