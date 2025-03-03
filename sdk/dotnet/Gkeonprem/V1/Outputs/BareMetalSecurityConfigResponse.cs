// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkeonprem.V1.Outputs
{

    /// <summary>
    /// Specifies the security related settings for the bare metal user cluster.
    /// </summary>
    [OutputType]
    public sealed class BareMetalSecurityConfigResponse
    {
        /// <summary>
        /// Configures user access to the user cluster.
        /// </summary>
        public readonly Outputs.AuthorizationResponse Authorization;

        [OutputConstructor]
        private BareMetalSecurityConfigResponse(Outputs.AuthorizationResponse authorization)
        {
            Authorization = authorization;
        }
    }
}
