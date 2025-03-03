// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V3.Outputs
{

    /// <summary>
    /// GKE Namespace. The field names correspond to the resource metadata labels on monitored resources that fall under a namespace (for example, k8s_container or k8s_pod).
    /// </summary>
    [OutputType]
    public sealed class GkeNamespaceResponse
    {
        /// <summary>
        /// The name of the parent cluster.
        /// </summary>
        public readonly string ClusterName;
        /// <summary>
        /// The location of the parent cluster. This may be a zone or region.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of this namespace.
        /// </summary>
        public readonly string NamespaceName;
        /// <summary>
        /// The project this resource lives in. For legacy services migrated from the Custom type, this may be a distinct project from the one parenting the service itself.
        /// </summary>
        public readonly string Project;

        [OutputConstructor]
        private GkeNamespaceResponse(
            string clusterName,

            string location,

            string namespaceName,

            string project)
        {
            ClusterName = clusterName;
            Location = location;
            NamespaceName = namespaceName;
            Project = project;
        }
    }
}
