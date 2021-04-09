// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.ServiceNetworking.V1
{
    /// <summary>
    /// Creates a private connection that establishes a VPC Network Peering connection to a VPC network in the service producer's organization. The administrator of the service consumer's VPC network invokes this method. The administrator must assign one or more allocated IP ranges for provisioning subnetworks in the service producer's VPC network. This connection is used for all supported services in the service producer's organization, so it only needs to be invoked once.
    /// </summary>
    [GcpNativeResourceType("gcp-native:servicenetworking/v1:ServiceConnection")]
    public partial class ServiceConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// Create a ServiceConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceConnection(string name, ServiceConnectionArgs args, CustomResourceOptions? options = null)
            : base("gcp-native:servicenetworking/v1:ServiceConnection", name, args ?? new ServiceConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("gcp-native:servicenetworking/v1:ServiceConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServiceConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServiceConnection(name, id, options);
        }
    }

    public sealed class ServiceConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of service consumer's VPC network that's connected with service producer network, in the following format: `projects/{project}/global/networks/{network}`. `{project}` is a project number, such as in `12345` that includes the VPC service consumer's VPC network. `{network}` is the name of the service consumer's VPC network.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        [Input("reservedPeeringRanges")]
        private InputList<string>? _reservedPeeringRanges;

        /// <summary>
        /// The name of one or more allocated IP address ranges for this service producer of type `PEERING`. Note that invoking CreateConnection method with a different range when connection is already established will not modify already provisioned service producer subnetworks. If CreateConnection method is invoked repeatedly to reconnect when peering connection had been disconnected on the consumer side, leaving this field empty will restore previously allocated IP ranges.
        /// </summary>
        public InputList<string> ReservedPeeringRanges
        {
            get => _reservedPeeringRanges ?? (_reservedPeeringRanges = new InputList<string>());
            set => _reservedPeeringRanges = value;
        }

        [Input("servicesId", required: true)]
        public Input<string> ServicesId { get; set; } = null!;

        public ServiceConnectionArgs()
        {
        }
    }
}