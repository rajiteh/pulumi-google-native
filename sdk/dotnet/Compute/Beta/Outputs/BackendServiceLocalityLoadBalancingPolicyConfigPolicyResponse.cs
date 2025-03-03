// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// The configuration for a built-in load balancing policy.
    /// </summary>
    [OutputType]
    public sealed class BackendServiceLocalityLoadBalancingPolicyConfigPolicyResponse
    {
        /// <summary>
        /// The name of a locality load-balancing policy. Valid values include ROUND_ROBIN and, for Java clients, LEAST_REQUEST. For information about these values, see the description of localityLbPolicy. Do not specify the same policy more than once for a backend. If you do, the configuration is rejected.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private BackendServiceLocalityLoadBalancingPolicyConfigPolicyResponse(string name)
        {
            Name = name;
        }
    }
}
