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
    /// Authorization defines the On-Prem cluster authorization configuration to bootstrap onto the admin cluster.
    /// </summary>
    [OutputType]
    public sealed class AuthorizationResponse
    {
        /// <summary>
        /// For VMware user, bare metal user and standalone clusters, users that will be granted the cluster-admin role on the cluster, providing full access to the cluster. For bare metal Admin cluster, users will be granted the view role, which is a view only access.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterUserResponse> AdminUsers;

        [OutputConstructor]
        private AuthorizationResponse(ImmutableArray<Outputs.ClusterUserResponse> adminUsers)
        {
            AdminUsers = adminUsers;
        }
    }
}
