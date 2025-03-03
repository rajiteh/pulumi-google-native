// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    /// <summary>
    /// Creates a Route resource in the specified project using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/alpha:Route")]
    public partial class Route : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether this route can conflict with existing subnetworks. Setting this to true allows this route to conflict with subnetworks that have already been configured on the corresponding network.
        /// </summary>
        [Output("allowConflictingSubnetworks")]
        public Output<bool> AllowConflictingSubnetworks { get; private set; } = null!;

        /// <summary>
        /// AS path.
        /// </summary>
        [Output("asPaths")]
        public Output<ImmutableArray<Outputs.RouteAsPathResponse>> AsPaths { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this field when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The destination range of outgoing packets that this route applies to. Both IPv4 and IPv6 are supported.
        /// </summary>
        [Output("destRange")]
        public Output<string> DestRange { get; private set; } = null!;

        /// <summary>
        /// ILB route behavior when ILB is deemed unhealthy based on user specified threshold on the Backend Service of the internal load balancing.
        /// </summary>
        [Output("ilbRouteBehaviorOnUnhealthy")]
        public Output<string> IlbRouteBehaviorOnUnhealthy { get; private set; } = null!;

        /// <summary>
        /// Type of this resource. Always compute#routes for Route resources.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Fully-qualified URL of the network that this route applies to.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// The URL to a gateway that should handle matching packets. You can only specify the internet gateway using a full or partial valid URL: projects/ project/global/gateways/default-internet-gateway
        /// </summary>
        [Output("nextHopGateway")]
        public Output<string> NextHopGateway { get; private set; } = null!;

        /// <summary>
        /// The full resource name of the network connectivity center hub that should handle matching packets.
        /// </summary>
        [Output("nextHopHub")]
        public Output<string> NextHopHub { get; private set; } = null!;

        /// <summary>
        /// The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets or the IP address of the forwarding Rule. For example, the following are all valid URLs: - 10.128.0.56 - https://www.googleapis.com/compute/v1/projects/project/regions/region /forwardingRules/forwardingRule - regions/region/forwardingRules/forwardingRule 
        /// </summary>
        [Output("nextHopIlb")]
        public Output<string> NextHopIlb { get; private set; } = null!;

        /// <summary>
        /// The URL to an instance that should handle matching packets. You can specify this as a full or partial URL. For example: https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/
        /// </summary>
        [Output("nextHopInstance")]
        public Output<string> NextHopInstance { get; private set; } = null!;

        /// <summary>
        /// The URL to an InterconnectAttachment which is the next hop for the route. This field will only be populated for the dynamic routes generated by Cloud Router with a linked interconnectAttachment.
        /// </summary>
        [Output("nextHopInterconnectAttachment")]
        public Output<string> NextHopInterconnectAttachment { get; private set; } = null!;

        /// <summary>
        /// The network IP address of an instance that should handle matching packets. Only IPv4 is supported.
        /// </summary>
        [Output("nextHopIp")]
        public Output<string> NextHopIp { get; private set; } = null!;

        /// <summary>
        /// The URL of the local network if it should handle matching packets.
        /// </summary>
        [Output("nextHopNetwork")]
        public Output<string> NextHopNetwork { get; private set; } = null!;

        /// <summary>
        /// The network peering name that should handle matching packets, which should conform to RFC1035.
        /// </summary>
        [Output("nextHopPeering")]
        public Output<string> NextHopPeering { get; private set; } = null!;

        /// <summary>
        /// The URL to a VpnTunnel that should handle matching packets.
        /// </summary>
        [Output("nextHopVpnTunnel")]
        public Output<string> NextHopVpnTunnel { get; private set; } = null!;

        /// <summary>
        /// The priority of this route. Priority is used to break ties in cases where there is more than one matching route of equal prefix length. In cases where multiple routes have equal prefix length, the one with the lowest-numbered priority value wins. The default value is `1000`. The priority value must be from `0` to `65535`, inclusive.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// [Output only] The status of the route.
        /// </summary>
        [Output("routeStatus")]
        public Output<string> RouteStatus { get; private set; } = null!;

        /// <summary>
        /// The type of this route, which can be one of the following values: - 'TRANSIT' for a transit route that this router learned from another Cloud Router and will readvertise to one of its BGP peers - 'SUBNET' for a route from a subnet of the VPC - 'BGP' for a route learned from a BGP peer of this router - 'STATIC' for a static route
        /// </summary>
        [Output("routeType")]
        public Output<string> RouteType { get; private set; } = null!;

        /// <summary>
        /// Server-defined fully-qualified URL for this resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL for this resource with the resource id.
        /// </summary>
        [Output("selfLinkWithId")]
        public Output<string> SelfLinkWithId { get; private set; } = null!;

        /// <summary>
        /// A list of instance tags to which this route applies.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// If potential misconfigurations are detected for this route, this field will be populated with warning messages.
        /// </summary>
        [Output("warnings")]
        public Output<ImmutableArray<Outputs.RouteWarningsItemResponse>> Warnings { get; private set; } = null!;


        /// <summary>
        /// Create a Route resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Route(string name, RouteArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:compute/alpha:Route", name, args ?? new RouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Route(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/alpha:Route", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Route resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Route Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Route(name, id, options);
        }
    }

    public sealed class RouteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether this route can conflict with existing subnetworks. Setting this to true allows this route to conflict with subnetworks that have already been configured on the corresponding network.
        /// </summary>
        [Input("allowConflictingSubnetworks")]
        public Input<bool>? AllowConflictingSubnetworks { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this field when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The destination range of outgoing packets that this route applies to. Both IPv4 and IPv6 are supported.
        /// </summary>
        [Input("destRange")]
        public Input<string>? DestRange { get; set; }

        /// <summary>
        /// ILB route behavior when ILB is deemed unhealthy based on user specified threshold on the Backend Service of the internal load balancing.
        /// </summary>
        [Input("ilbRouteBehaviorOnUnhealthy")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.RouteIlbRouteBehaviorOnUnhealthy>? IlbRouteBehaviorOnUnhealthy { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Fully-qualified URL of the network that this route applies to.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The URL to a gateway that should handle matching packets. You can only specify the internet gateway using a full or partial valid URL: projects/ project/global/gateways/default-internet-gateway
        /// </summary>
        [Input("nextHopGateway")]
        public Input<string>? NextHopGateway { get; set; }

        /// <summary>
        /// The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets or the IP address of the forwarding Rule. For example, the following are all valid URLs: - 10.128.0.56 - https://www.googleapis.com/compute/v1/projects/project/regions/region /forwardingRules/forwardingRule - regions/region/forwardingRules/forwardingRule 
        /// </summary>
        [Input("nextHopIlb")]
        public Input<string>? NextHopIlb { get; set; }

        /// <summary>
        /// The URL to an instance that should handle matching packets. You can specify this as a full or partial URL. For example: https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/
        /// </summary>
        [Input("nextHopInstance")]
        public Input<string>? NextHopInstance { get; set; }

        /// <summary>
        /// The network IP address of an instance that should handle matching packets. Only IPv4 is supported.
        /// </summary>
        [Input("nextHopIp")]
        public Input<string>? NextHopIp { get; set; }

        /// <summary>
        /// The URL of the local network if it should handle matching packets.
        /// </summary>
        [Input("nextHopNetwork")]
        public Input<string>? NextHopNetwork { get; set; }

        /// <summary>
        /// The URL to a VpnTunnel that should handle matching packets.
        /// </summary>
        [Input("nextHopVpnTunnel")]
        public Input<string>? NextHopVpnTunnel { get; set; }

        /// <summary>
        /// The priority of this route. Priority is used to break ties in cases where there is more than one matching route of equal prefix length. In cases where multiple routes have equal prefix length, the one with the lowest-numbered priority value wins. The default value is `1000`. The priority value must be from `0` to `65535`, inclusive.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of instance tags to which this route applies.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public RouteArgs()
        {
        }
        public static new RouteArgs Empty => new RouteArgs();
    }
}
