// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkeonprem.V1
{
    public static class GetVmwareCluster
    {
        /// <summary>
        /// Gets details of a single VMware Cluster.
        /// </summary>
        public static Task<GetVmwareClusterResult> InvokeAsync(GetVmwareClusterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVmwareClusterResult>("google-native:gkeonprem/v1:getVmwareCluster", args ?? new GetVmwareClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single VMware Cluster.
        /// </summary>
        public static Output<GetVmwareClusterResult> Invoke(GetVmwareClusterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVmwareClusterResult>("google-native:gkeonprem/v1:getVmwareCluster", args ?? new GetVmwareClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVmwareClusterArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("view")]
        public string? View { get; set; }

        [Input("vmwareClusterId", required: true)]
        public string VmwareClusterId { get; set; } = null!;

        public GetVmwareClusterArgs()
        {
        }
        public static new GetVmwareClusterArgs Empty => new GetVmwareClusterArgs();
    }

    public sealed class GetVmwareClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("view")]
        public Input<string>? View { get; set; }

        [Input("vmwareClusterId", required: true)]
        public Input<string> VmwareClusterId { get; set; } = null!;

        public GetVmwareClusterInvokeArgs()
        {
        }
        public static new GetVmwareClusterInvokeArgs Empty => new GetVmwareClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetVmwareClusterResult
    {
        /// <summary>
        /// The admin cluster this VMware user cluster belongs to. This is the full resource name of the admin cluster's fleet membership. In the future, references to other resource types might be allowed if admin clusters are modeled as their own resources.
        /// </summary>
        public readonly string AdminClusterMembership;
        /// <summary>
        /// The resource name of the VMware admin cluster hosting this user cluster.
        /// </summary>
        public readonly string AdminClusterName;
        /// <summary>
        /// Annotations on the VMware user cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Annotations;
        /// <summary>
        /// AAGConfig specifies whether to spread VMware user cluster nodes across at least three physical hosts in the datacenter.
        /// </summary>
        public readonly Outputs.VmwareAAGConfigResponse AntiAffinityGroups;
        /// <summary>
        /// RBAC policy that will be applied and managed by the Anthos On-Prem API.
        /// </summary>
        public readonly Outputs.AuthorizationResponse Authorization;
        /// <summary>
        /// Configuration for auto repairing.
        /// </summary>
        public readonly Outputs.VmwareAutoRepairConfigResponse AutoRepairConfig;
        /// <summary>
        /// VMware user cluster control plane nodes must have either 1 or 3 replicas.
        /// </summary>
        public readonly Outputs.VmwareControlPlaneNodeConfigResponse ControlPlaneNode;
        /// <summary>
        /// The time at which VMware user cluster was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// VmwareDataplaneV2Config specifies configuration for Dataplane V2.
        /// </summary>
        public readonly Outputs.VmwareDataplaneV2ConfigResponse DataplaneV2;
        /// <summary>
        /// The time at which VMware user cluster was deleted.
        /// </summary>
        public readonly string DeleteTime;
        /// <summary>
        /// A human readable description of this VMware user cluster.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Enable control plane V2. Default to false.
        /// </summary>
        public readonly bool EnableControlPlaneV2;
        /// <summary>
        /// The DNS name of VMware user cluster's API server.
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
        public readonly Outputs.VmwareLoadBalancerConfigResponse LoadBalancer;
        /// <summary>
        /// The object name of the VMware OnPremUserCluster custom resource on the associated admin cluster. This field is used to support conflicting names when enrolling existing clusters to the API. When used as a part of cluster enrollment, this field will differ from the ID in the resource name. For new clusters, this field will match the user provided cluster name and be visible in the last component of the resource name. It is not modifiable. All users should use this name to access their cluster using gkectl or kubectl and should expect to see the local name when viewing admin cluster controller logs.
        /// </summary>
        public readonly string LocalName;
        /// <summary>
        /// Immutable. The VMware user cluster resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The VMware user cluster network configuration.
        /// </summary>
        public readonly Outputs.VmwareNetworkConfigResponse NetworkConfig;
        /// <summary>
        /// The Anthos clusters on the VMware version for your user cluster. Defaults to the admin cluster version.
        /// </summary>
        public readonly string OnPremVersion;
        /// <summary>
        /// If set, there are currently changes in flight to the VMware user cluster.
        /// </summary>
        public readonly bool Reconciling;
        /// <summary>
        /// The current state of VMware user cluster.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// ResourceStatus representing detailed cluster state.
        /// </summary>
        public readonly Outputs.ResourceStatusResponse Status;
        /// <summary>
        /// Storage configuration.
        /// </summary>
        public readonly Outputs.VmwareStorageConfigResponse Storage;
        /// <summary>
        /// The unique identifier of the VMware user cluster.
        /// </summary>
        public readonly string Uid;
        /// <summary>
        /// The time at which VMware user cluster was last updated.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// ValidationCheck represents the result of the preflight check job.
        /// </summary>
        public readonly Outputs.ValidationCheckResponse ValidationCheck;
        /// <summary>
        /// VmwareVCenterConfig specifies vCenter config for the user cluster. Inherited from the admin cluster.
        /// </summary>
        public readonly Outputs.VmwareVCenterConfigResponse Vcenter;
        /// <summary>
        /// Enable VM tracking.
        /// </summary>
        public readonly bool VmTrackingEnabled;

        [OutputConstructor]
        private GetVmwareClusterResult(
            string adminClusterMembership,

            string adminClusterName,

            ImmutableDictionary<string, string> annotations,

            Outputs.VmwareAAGConfigResponse antiAffinityGroups,

            Outputs.AuthorizationResponse authorization,

            Outputs.VmwareAutoRepairConfigResponse autoRepairConfig,

            Outputs.VmwareControlPlaneNodeConfigResponse controlPlaneNode,

            string createTime,

            Outputs.VmwareDataplaneV2ConfigResponse dataplaneV2,

            string deleteTime,

            string description,

            bool enableControlPlaneV2,

            string endpoint,

            string etag,

            Outputs.FleetResponse fleet,

            Outputs.VmwareLoadBalancerConfigResponse loadBalancer,

            string localName,

            string name,

            Outputs.VmwareNetworkConfigResponse networkConfig,

            string onPremVersion,

            bool reconciling,

            string state,

            Outputs.ResourceStatusResponse status,

            Outputs.VmwareStorageConfigResponse storage,

            string uid,

            string updateTime,

            Outputs.ValidationCheckResponse validationCheck,

            Outputs.VmwareVCenterConfigResponse vcenter,

            bool vmTrackingEnabled)
        {
            AdminClusterMembership = adminClusterMembership;
            AdminClusterName = adminClusterName;
            Annotations = annotations;
            AntiAffinityGroups = antiAffinityGroups;
            Authorization = authorization;
            AutoRepairConfig = autoRepairConfig;
            ControlPlaneNode = controlPlaneNode;
            CreateTime = createTime;
            DataplaneV2 = dataplaneV2;
            DeleteTime = deleteTime;
            Description = description;
            EnableControlPlaneV2 = enableControlPlaneV2;
            Endpoint = endpoint;
            Etag = etag;
            Fleet = fleet;
            LoadBalancer = loadBalancer;
            LocalName = localName;
            Name = name;
            NetworkConfig = networkConfig;
            OnPremVersion = onPremVersion;
            Reconciling = reconciling;
            State = state;
            Status = status;
            Storage = storage;
            Uid = uid;
            UpdateTime = updateTime;
            ValidationCheck = validationCheck;
            Vcenter = vcenter;
            VmTrackingEnabled = vmTrackingEnabled;
        }
    }
}
