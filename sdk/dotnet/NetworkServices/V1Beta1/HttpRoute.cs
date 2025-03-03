// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1Beta1
{
    /// <summary>
    /// Creates a new HttpRoute in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:networkservices/v1beta1:HttpRoute")]
    public partial class HttpRoute : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The timestamp when the resource was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. A free-text description of the resource. Max length 1024 characters.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Optional. Gateways defines a list of gateways this HttpRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
        /// </summary>
        [Output("gateways")]
        public Output<ImmutableArray<string>> Gateways { get; private set; } = null!;

        /// <summary>
        /// Hostnames define a set of hosts that should match against the HTTP host header to select a HttpRoute to process the request. Hostname is the fully qualified domain name of a network host, as defined by RFC 1123 with the exception that: - IPs are not allowed. - A hostname may be prefixed with a wildcard label (`*.`). The wildcard label must appear by itself as the first label. Hostname can be "precise" which is a domain name without the terminating dot of a network host (e.g. `foo.example.com`) or "wildcard", which is a domain name prefixed with a single wildcard label (e.g. `*.example.com`). Note that as per RFC1035 and RFC1123, a label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character. No other punctuation is allowed. The routes associated with a Mesh or Gateways must have unique hostnames. If you attempt to attach multiple routes with conflicting hostnames, the configuration will be rejected. For example, while it is acceptable for routes for the hostnames `*.foo.bar.com` and `*.bar.com` to be associated with the same Mesh (or Gateways under the same scope), it is not possible to associate two routes both with `*.bar.com` or both with `bar.com`.
        /// </summary>
        [Output("hostnames")]
        public Output<ImmutableArray<string>> Hostnames { get; private set; } = null!;

        /// <summary>
        /// Required. Short name of the HttpRoute resource to be created.
        /// </summary>
        [Output("httpRouteId")]
        public Output<string> HttpRouteId { get; private set; } = null!;

        /// <summary>
        /// Optional. Set of label tags associated with the HttpRoute resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Optional. Meshes defines a list of meshes this HttpRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/` The attached Mesh should be of a type SIDECAR
        /// </summary>
        [Output("meshes")]
        public Output<ImmutableArray<string>> Meshes { get; private set; } = null!;

        /// <summary>
        /// Name of the HttpRoute resource. It matches pattern `projects/*/locations/global/httpRoutes/http_route_name&gt;`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Rules that define how traffic is routed and handled. Rules will be matched sequentially based on the RouteMatch specified for the rule.
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.HttpRouteRouteRuleResponse>> Rules { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL of this resource
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the resource was updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a HttpRoute resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HttpRoute(string name, HttpRouteArgs args, CustomResourceOptions? options = null)
            : base("google-native:networkservices/v1beta1:HttpRoute", name, args ?? new HttpRouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HttpRoute(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:networkservices/v1beta1:HttpRoute", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "httpRouteId",
                    "location",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing HttpRoute resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HttpRoute Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new HttpRoute(name, id, options);
        }
    }

    public sealed class HttpRouteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. A free-text description of the resource. Max length 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("gateways")]
        private InputList<string>? _gateways;

        /// <summary>
        /// Optional. Gateways defines a list of gateways this HttpRoute is attached to, as one of the routing rules to route the requests served by the gateway. Each gateway reference should match the pattern: `projects/*/locations/global/gateways/`
        /// </summary>
        public InputList<string> Gateways
        {
            get => _gateways ?? (_gateways = new InputList<string>());
            set => _gateways = value;
        }

        [Input("hostnames", required: true)]
        private InputList<string>? _hostnames;

        /// <summary>
        /// Hostnames define a set of hosts that should match against the HTTP host header to select a HttpRoute to process the request. Hostname is the fully qualified domain name of a network host, as defined by RFC 1123 with the exception that: - IPs are not allowed. - A hostname may be prefixed with a wildcard label (`*.`). The wildcard label must appear by itself as the first label. Hostname can be "precise" which is a domain name without the terminating dot of a network host (e.g. `foo.example.com`) or "wildcard", which is a domain name prefixed with a single wildcard label (e.g. `*.example.com`). Note that as per RFC1035 and RFC1123, a label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character. No other punctuation is allowed. The routes associated with a Mesh or Gateways must have unique hostnames. If you attempt to attach multiple routes with conflicting hostnames, the configuration will be rejected. For example, while it is acceptable for routes for the hostnames `*.foo.bar.com` and `*.bar.com` to be associated with the same Mesh (or Gateways under the same scope), it is not possible to associate two routes both with `*.bar.com` or both with `bar.com`.
        /// </summary>
        public InputList<string> Hostnames
        {
            get => _hostnames ?? (_hostnames = new InputList<string>());
            set => _hostnames = value;
        }

        /// <summary>
        /// Required. Short name of the HttpRoute resource to be created.
        /// </summary>
        [Input("httpRouteId", required: true)]
        public Input<string> HttpRouteId { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. Set of label tags associated with the HttpRoute resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("meshes")]
        private InputList<string>? _meshes;

        /// <summary>
        /// Optional. Meshes defines a list of meshes this HttpRoute is attached to, as one of the routing rules to route the requests served by the mesh. Each mesh reference should match the pattern: `projects/*/locations/global/meshes/` The attached Mesh should be of a type SIDECAR
        /// </summary>
        public InputList<string> Meshes
        {
            get => _meshes ?? (_meshes = new InputList<string>());
            set => _meshes = value;
        }

        /// <summary>
        /// Name of the HttpRoute resource. It matches pattern `projects/*/locations/global/httpRoutes/http_route_name&gt;`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("rules", required: true)]
        private InputList<Inputs.HttpRouteRouteRuleArgs>? _rules;

        /// <summary>
        /// Rules that define how traffic is routed and handled. Rules will be matched sequentially based on the RouteMatch specified for the rule.
        /// </summary>
        public InputList<Inputs.HttpRouteRouteRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.HttpRouteRouteRuleArgs>());
            set => _rules = value;
        }

        public HttpRouteArgs()
        {
        }
        public static new HttpRouteArgs Empty => new HttpRouteArgs();
    }
}
