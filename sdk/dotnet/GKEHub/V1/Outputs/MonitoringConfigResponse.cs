// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1.Outputs
{

    /// <summary>
    /// This field informs Fleet-based applications/services/UIs with the necessary information for where each underlying Cluster reports its metrics.
    /// </summary>
    [OutputType]
    public sealed class MonitoringConfigResponse
    {
        /// <summary>
        /// Immutable. Cluster name used to report metrics. For Anthos on VMWare/Baremetal, it would be in format `memberClusters/cluster_name`; And for Anthos on MultiCloud, it would be in format `{azureClusters, awsClusters}/cluster_name`.
        /// </summary>
        public readonly string Cluster;
        /// <summary>
        /// Immutable. Cluster hash, this is a unique string generated by google code, which does not contain any PII, which we can use to reference the cluster. This is expected to be created by the monitoring stack and persisted into the Cluster object as well as to GKE-Hub.
        /// </summary>
        public readonly string ClusterHash;
        /// <summary>
        /// Kubernetes system metrics, if available, are written to this prefix. This defaults to kubernetes.io for GKE, and kubernetes.io/anthos for Anthos eventually. Noted: Anthos MultiCloud will have kubernetes.io prefix today but will migration to be under kubernetes.io/anthos
        /// </summary>
        public readonly string KubernetesMetricsPrefix;
        /// <summary>
        /// Immutable. Location used to report Metrics
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Immutable. Project used to report Metrics
        /// </summary>
        public readonly string Project;

        [OutputConstructor]
        private MonitoringConfigResponse(
            string cluster,

            string clusterHash,

            string kubernetesMetricsPrefix,

            string location,

            string project)
        {
            Cluster = cluster;
            ClusterHash = clusterHash;
            KubernetesMetricsPrefix = kubernetesMetricsPrefix;
            Location = location;
            Project = project;
        }
    }
}