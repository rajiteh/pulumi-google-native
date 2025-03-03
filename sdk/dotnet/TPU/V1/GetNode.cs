// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.TPU.V1
{
    public static class GetNode
    {
        /// <summary>
        /// Gets the details of a node.
        /// </summary>
        public static Task<GetNodeResult> InvokeAsync(GetNodeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNodeResult>("google-native:tpu/v1:getNode", args ?? new GetNodeArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of a node.
        /// </summary>
        public static Output<GetNodeResult> Invoke(GetNodeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNodeResult>("google-native:tpu/v1:getNode", args ?? new GetNodeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNodeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("nodeId", required: true)]
        public string NodeId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetNodeArgs()
        {
        }
        public static new GetNodeArgs Empty => new GetNodeArgs();
    }

    public sealed class GetNodeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("nodeId", required: true)]
        public Input<string> NodeId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetNodeInvokeArgs()
        {
        }
        public static new GetNodeInvokeArgs Empty => new GetNodeInvokeArgs();
    }


    [OutputType]
    public sealed class GetNodeResult
    {
        /// <summary>
        /// The type of hardware accelerators associated with this node.
        /// </summary>
        public readonly string AcceleratorType;
        /// <summary>
        /// The API version that created this Node.
        /// </summary>
        public readonly string ApiVersion;
        /// <summary>
        /// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network that is using that CIDR block.
        /// </summary>
        public readonly string CidrBlock;
        /// <summary>
        /// The time when the node was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The user-supplied description of the TPU. Maximum of 512 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The health status of the TPU node.
        /// </summary>
        public readonly string Health;
        /// <summary>
        /// If this field is populated, it contains a description of why the TPU Node is unhealthy.
        /// </summary>
        public readonly string HealthDescription;
        /// <summary>
        /// DEPRECATED! Use network_endpoints instead. The network address for the TPU Node as visible to Compute Engine instances.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// Resource labels to represent user-provided metadata.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Immutable. The name of the TPU
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The name of a network they wish to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on which this API has been activated. If none is provided, "default" will be used.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the node reach out to the 0th entry in this map first.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkEndpointResponse> NetworkEndpoints;
        /// <summary>
        /// DEPRECATED! Use network_endpoints instead. The network port for the TPU Node as visible to Compute Engine instances.
        /// </summary>
        public readonly string Port;
        /// <summary>
        /// The scheduling options for this node.
        /// </summary>
        public readonly Outputs.SchedulingConfigResponse SchedulingConfig;
        /// <summary>
        /// The service account used to run the tensor flow services within the node. To share resources, including Google Cloud Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// The current state for the TPU Node.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The Symptoms that have occurred to the TPU Node.
        /// </summary>
        public readonly ImmutableArray<Outputs.SymptomResponse> Symptoms;
        /// <summary>
        /// The version of Tensorflow running in the Node.
        /// </summary>
        public readonly string TensorflowVersion;
        /// <summary>
        /// Whether the VPC peering for the node is set up through Service Networking API. The VPC Peering should be set up before provisioning the node. If this field is set, cidr_block field should not be specified. If the network, that you want to peer the TPU Node to, is Shared VPC networks, the node must be created with this this field enabled.
        /// </summary>
        public readonly bool UseServiceNetworking;

        [OutputConstructor]
        private GetNodeResult(
            string acceleratorType,

            string apiVersion,

            string cidrBlock,

            string createTime,

            string description,

            string health,

            string healthDescription,

            string ipAddress,

            ImmutableDictionary<string, string> labels,

            string name,

            string network,

            ImmutableArray<Outputs.NetworkEndpointResponse> networkEndpoints,

            string port,

            Outputs.SchedulingConfigResponse schedulingConfig,

            string serviceAccount,

            string state,

            ImmutableArray<Outputs.SymptomResponse> symptoms,

            string tensorflowVersion,

            bool useServiceNetworking)
        {
            AcceleratorType = acceleratorType;
            ApiVersion = apiVersion;
            CidrBlock = cidrBlock;
            CreateTime = createTime;
            Description = description;
            Health = health;
            HealthDescription = healthDescription;
            IpAddress = ipAddress;
            Labels = labels;
            Name = name;
            Network = network;
            NetworkEndpoints = networkEndpoints;
            Port = port;
            SchedulingConfig = schedulingConfig;
            ServiceAccount = serviceAccount;
            State = state;
            Symptoms = symptoms;
            TensorflowVersion = tensorflowVersion;
            UseServiceNetworking = useServiceNetworking;
        }
    }
}
