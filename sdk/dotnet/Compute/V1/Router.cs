// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    /// <summary>
    /// Creates a Router resource in the specified project and region using the data included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/v1:Router")]
    public partial class Router : global::Pulumi.CustomResource
    {
        /// <summary>
        /// BGP information specific to this router.
        /// </summary>
        [Output("bgp")]
        public Output<Outputs.RouterBgpResponse> Bgp { get; private set; } = null!;

        /// <summary>
        /// BGP information that must be configured into the routing stack to establish BGP peering. This information must specify the peer ASN and either the interface name, IP address, or peer IP address. Please refer to RFC4273.
        /// </summary>
        [Output("bgpPeers")]
        public Output<ImmutableArray<Outputs.RouterBgpPeerResponse>> BgpPeers { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Indicates if a router is dedicated for use with encrypted VLAN attachments (interconnectAttachments).
        /// </summary>
        [Output("encryptedInterconnectRouter")]
        public Output<bool> EncryptedInterconnectRouter { get; private set; } = null!;

        /// <summary>
        /// Router interfaces. Each interface requires either one linked resource, (for example, linkedVpnTunnel), or IP address and IP address range (for example, ipRange), or both.
        /// </summary>
        [Output("interfaces")]
        public Output<ImmutableArray<Outputs.RouterInterfaceResponse>> Interfaces { get; private set; } = null!;

        /// <summary>
        /// Type of resource. Always compute#router for routers.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Keys used for MD5 authentication.
        /// </summary>
        [Output("md5AuthenticationKeys")]
        public Output<ImmutableArray<Outputs.RouterMd5AuthenticationKeyResponse>> Md5AuthenticationKeys { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of NAT services created in this router.
        /// </summary>
        [Output("nats")]
        public Output<ImmutableArray<Outputs.RouterNatResponse>> Nats { get; private set; } = null!;

        /// <summary>
        /// URI of the network to which this router belongs.
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;


        /// <summary>
        /// Create a Router resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Router(string name, RouterArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:Router", name, args ?? new RouterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Router(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:Router", name, null, MakeResourceOptions(options, id))
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
                    "region",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Router resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Router Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Router(name, id, options);
        }
    }

    public sealed class RouterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// BGP information specific to this router.
        /// </summary>
        [Input("bgp")]
        public Input<Inputs.RouterBgpArgs>? Bgp { get; set; }

        [Input("bgpPeers")]
        private InputList<Inputs.RouterBgpPeerArgs>? _bgpPeers;

        /// <summary>
        /// BGP information that must be configured into the routing stack to establish BGP peering. This information must specify the peer ASN and either the interface name, IP address, or peer IP address. Please refer to RFC4273.
        /// </summary>
        public InputList<Inputs.RouterBgpPeerArgs> BgpPeers
        {
            get => _bgpPeers ?? (_bgpPeers = new InputList<Inputs.RouterBgpPeerArgs>());
            set => _bgpPeers = value;
        }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates if a router is dedicated for use with encrypted VLAN attachments (interconnectAttachments).
        /// </summary>
        [Input("encryptedInterconnectRouter")]
        public Input<bool>? EncryptedInterconnectRouter { get; set; }

        [Input("interfaces")]
        private InputList<Inputs.RouterInterfaceArgs>? _interfaces;

        /// <summary>
        /// Router interfaces. Each interface requires either one linked resource, (for example, linkedVpnTunnel), or IP address and IP address range (for example, ipRange), or both.
        /// </summary>
        public InputList<Inputs.RouterInterfaceArgs> Interfaces
        {
            get => _interfaces ?? (_interfaces = new InputList<Inputs.RouterInterfaceArgs>());
            set => _interfaces = value;
        }

        [Input("md5AuthenticationKeys")]
        private InputList<Inputs.RouterMd5AuthenticationKeyArgs>? _md5AuthenticationKeys;

        /// <summary>
        /// Keys used for MD5 authentication.
        /// </summary>
        public InputList<Inputs.RouterMd5AuthenticationKeyArgs> Md5AuthenticationKeys
        {
            get => _md5AuthenticationKeys ?? (_md5AuthenticationKeys = new InputList<Inputs.RouterMd5AuthenticationKeyArgs>());
            set => _md5AuthenticationKeys = value;
        }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nats")]
        private InputList<Inputs.RouterNatArgs>? _nats;

        /// <summary>
        /// A list of NAT services created in this router.
        /// </summary>
        public InputList<Inputs.RouterNatArgs> Nats
        {
            get => _nats ?? (_nats = new InputList<Inputs.RouterNatArgs>());
            set => _nats = value;
        }

        /// <summary>
        /// URI of the network to which this router belongs.
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        public RouterArgs()
        {
        }
        public static new RouterArgs Empty => new RouterArgs();
    }
}
