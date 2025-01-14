// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Outputs
{

    /// <summary>
    /// Configuration for the legacy Attribute Based Access Control authorization mode.
    /// </summary>
    [OutputType]
    public sealed class LegacyAbacResponse
    {
        /// <summary>
        /// Whether the ABAC authorizer is enabled for this cluster. When enabled, identities in the system, including service accounts, nodes, and controllers, will have statically granted permissions beyond those provided by the RBAC configuration or IAM.
        /// </summary>
        public readonly bool Enabled;

        [OutputConstructor]
        private LegacyAbacResponse(bool enabled)
        {
            Enabled = enabled;
        }
    }
}
