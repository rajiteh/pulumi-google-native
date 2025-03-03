// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkeonprem.V1
{
    public static class GetBareMetalAdminCluster
    {
        /// <summary>
        /// Gets details of a single bare metal admin cluster.
        /// </summary>
        public static Task<GetBareMetalAdminClusterResult> InvokeAsync(GetBareMetalAdminClusterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBareMetalAdminClusterResult>("google-native:gkeonprem/v1:getBareMetalAdminCluster", args ?? new GetBareMetalAdminClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single bare metal admin cluster.
        /// </summary>
        public static Output<GetBareMetalAdminClusterResult> Invoke(GetBareMetalAdminClusterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBareMetalAdminClusterResult>("google-native:gkeonprem/v1:getBareMetalAdminCluster", args ?? new GetBareMetalAdminClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBareMetalAdminClusterArgs : global::Pulumi.InvokeArgs
    {
        [Input("bareMetalAdminClusterId", required: true)]
        public string BareMetalAdminClusterId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("view")]
        public string? View { get; set; }

        public GetBareMetalAdminClusterArgs()
        {
        }
        public static new GetBareMetalAdminClusterArgs Empty => new GetBareMetalAdminClusterArgs();
    }

    public sealed class GetBareMetalAdminClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("bareMetalAdminClusterId", required: true)]
        public Input<string> BareMetalAdminClusterId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("view")]
        public Input<string>? View { get; set; }

        public GetBareMetalAdminClusterInvokeArgs()
        {
        }
        public static new GetBareMetalAdminClusterInvokeArgs Empty => new GetBareMetalAdminClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetBareMetalAdminClusterResult
    {
        /// <summary>
        /// Annotations on the bare metal admin cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Annotations;
        /// <summary>
        /// The Anthos clusters on bare metal version for the bare metal admin cluster.
        /// </summary>
        public readonly string BareMetalVersion;
        /// <summary>
        /// Cluster operations configuration.
        /// </summary>
        public readonly Outputs.BareMetalAdminClusterOperationsConfigResponse ClusterOperations;
        /// <summary>
        /// Control plane configuration.
        /// </summary>
        public readonly Outputs.BareMetalAdminControlPlaneConfigResponse ControlPlane;
        /// <summary>
        /// The time at which this bare metal admin cluster was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The time at which this bare metal admin cluster was deleted. If the resource is not deleted, this must be empty
        /// </summary>
        public readonly string DeleteTime;
        /// <summary>
        /// A human readable description of this bare metal admin cluster.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The IP address name of bare metal admin cluster's API server.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Allows clients to perform consistent read-modify-writes through optimistic concurrency control.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Fleet configuration for the cluster.
        /// </summary>
        public readonly Outputs.FleetResponse Fleet;
        /// <summary>
        /// Load balancer configuration.
        /// </summary>
        public readonly Outputs.BareMetalAdminLoadBalancerConfigResponse LoadBalancer;
        /// <summary>
        /// The object name of the bare metal cluster custom resource. This field is used to support conflicting names when enrolling existing clusters to the API. When used as a part of cluster enrollment, this field will differ from the ID in the resource name. For new clusters, this field will match the user provided cluster name and be visible in the last component of the resource name. It is not modifiable. All users should use this name to access their cluster using gkectl or kubectl and should expect to see the local name when viewing admin cluster controller logs.
        /// </summary>
        public readonly string LocalName;
        /// <summary>
        /// Maintenance configuration.
        /// </summary>
        public readonly Outputs.BareMetalAdminMaintenanceConfigResponse MaintenanceConfig;
        /// <summary>
        /// MaintenanceStatus representing state of maintenance.
        /// </summary>
        public readonly Outputs.BareMetalAdminMaintenanceStatusResponse MaintenanceStatus;
        /// <summary>
        /// Immutable. The bare metal admin cluster resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Network configuration.
        /// </summary>
        public readonly Outputs.BareMetalAdminNetworkConfigResponse NetworkConfig;
        /// <summary>
        /// Node access related configurations.
        /// </summary>
        public readonly Outputs.BareMetalAdminNodeAccessConfigResponse NodeAccessConfig;
        /// <summary>
        /// Workload node configuration.
        /// </summary>
        public readonly Outputs.BareMetalAdminWorkloadNodeConfigResponse NodeConfig;
        /// <summary>
        /// OS environment related configurations.
        /// </summary>
        public readonly Outputs.BareMetalAdminOsEnvironmentConfigResponse OsEnvironmentConfig;
        /// <summary>
        /// Proxy configuration.
        /// </summary>
        public readonly Outputs.BareMetalAdminProxyConfigResponse Proxy;
        /// <summary>
        /// If set, there are currently changes in flight to the bare metal Admin Cluster.
        /// </summary>
        public readonly bool Reconciling;
        /// <summary>
        /// Security related configuration.
        /// </summary>
        public readonly Outputs.BareMetalAdminSecurityConfigResponse SecurityConfig;
        /// <summary>
        /// The current state of the bare metal admin cluster.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// ResourceStatus representing detailed cluster status.
        /// </summary>
        public readonly Outputs.ResourceStatusResponse Status;
        /// <summary>
        /// Storage configuration.
        /// </summary>
        public readonly Outputs.BareMetalAdminStorageConfigResponse Storage;
        /// <summary>
        /// The unique identifier of the bare metal admin cluster.
        /// </summary>
        public readonly string Uid;
        /// <summary>
        /// The time at which this bare metal admin cluster was last updated.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// ValidationCheck representing the result of the preflight check.
        /// </summary>
        public readonly Outputs.ValidationCheckResponse ValidationCheck;

        [OutputConstructor]
        private GetBareMetalAdminClusterResult(
            ImmutableDictionary<string, string> annotations,

            string bareMetalVersion,

            Outputs.BareMetalAdminClusterOperationsConfigResponse clusterOperations,

            Outputs.BareMetalAdminControlPlaneConfigResponse controlPlane,

            string createTime,

            string deleteTime,

            string description,

            string endpoint,

            string etag,

            Outputs.FleetResponse fleet,

            Outputs.BareMetalAdminLoadBalancerConfigResponse loadBalancer,

            string localName,

            Outputs.BareMetalAdminMaintenanceConfigResponse maintenanceConfig,

            Outputs.BareMetalAdminMaintenanceStatusResponse maintenanceStatus,

            string name,

            Outputs.BareMetalAdminNetworkConfigResponse networkConfig,

            Outputs.BareMetalAdminNodeAccessConfigResponse nodeAccessConfig,

            Outputs.BareMetalAdminWorkloadNodeConfigResponse nodeConfig,

            Outputs.BareMetalAdminOsEnvironmentConfigResponse osEnvironmentConfig,

            Outputs.BareMetalAdminProxyConfigResponse proxy,

            bool reconciling,

            Outputs.BareMetalAdminSecurityConfigResponse securityConfig,

            string state,

            Outputs.ResourceStatusResponse status,

            Outputs.BareMetalAdminStorageConfigResponse storage,

            string uid,

            string updateTime,

            Outputs.ValidationCheckResponse validationCheck)
        {
            Annotations = annotations;
            BareMetalVersion = bareMetalVersion;
            ClusterOperations = clusterOperations;
            ControlPlane = controlPlane;
            CreateTime = createTime;
            DeleteTime = deleteTime;
            Description = description;
            Endpoint = endpoint;
            Etag = etag;
            Fleet = fleet;
            LoadBalancer = loadBalancer;
            LocalName = localName;
            MaintenanceConfig = maintenanceConfig;
            MaintenanceStatus = maintenanceStatus;
            Name = name;
            NetworkConfig = networkConfig;
            NodeAccessConfig = nodeAccessConfig;
            NodeConfig = nodeConfig;
            OsEnvironmentConfig = osEnvironmentConfig;
            Proxy = proxy;
            Reconciling = reconciling;
            SecurityConfig = securityConfig;
            State = state;
            Status = status;
            Storage = storage;
            Uid = uid;
            UpdateTime = updateTime;
            ValidationCheck = validationCheck;
        }
    }
}
