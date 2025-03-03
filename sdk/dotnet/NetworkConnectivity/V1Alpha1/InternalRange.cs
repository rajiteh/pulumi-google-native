// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkConnectivity.V1Alpha1
{
    /// <summary>
    /// Creates a new internal range in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:networkconnectivity/v1alpha1:InternalRange")]
    public partial class InternalRange : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Time when the internal range was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// A description of this resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Optional. Resource ID (i.e. 'foo' in '[...]/projects/p/locations/l/internalRanges/foo') See https://google.aip.dev/122#resource-id-segments Unique per location.
        /// </summary>
        [Output("internalRangeId")]
        public Output<string?> InternalRangeId { get; private set; } = null!;

        /// <summary>
        /// IP range that this internal range defines.
        /// </summary>
        [Output("ipCidrRange")]
        public Output<string> IpCidrRange { get; private set; } = null!;

        /// <summary>
        /// User-defined labels.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Immutable. The name of an internal range. Format: projects/{project}/locations/{location}/internalRanges/{internal_range} See: https://google.aip.dev/122#fields-representing-resource-names
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The URL or resource ID of the network in which to reserve the internal range. The network cannot be deleted if there are any reserved internal ranges referring to it. Legacy networks are not supported. This can only be specified for a global internal address. Example: - URL: /compute/v1/projects/{project}/global/networks/{resourceId} - ID: network123
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// Optional. Types of resources that are allowed to overlap with the current internal range.
        /// </summary>
        [Output("overlaps")]
        public Output<ImmutableArray<string>> Overlaps { get; private set; } = null!;

        /// <summary>
        /// The type of peering set for this internal range.
        /// </summary>
        [Output("peering")]
        public Output<string> Peering { get; private set; } = null!;

        /// <summary>
        /// An alternative to ip_cidr_range. Can be set when trying to create a reservation that automatically finds a free range of the given size. If both ip_cidr_range and prefix_length are set, there is an error if the range sizes do not match. Can also be used during updates to change the range size.
        /// </summary>
        [Output("prefixLength")]
        public Output<int> PrefixLength { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if the original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// Optional. Can be set to narrow down or pick a different address space while searching for a free range. If not set, defaults to the "10.0.0.0/8" address space. This can be used to search in other rfc-1918 address spaces like "172.16.0.0/12" and "192.168.0.0/16" or non-rfc-1918 address spaces used in the VPC.
        /// </summary>
        [Output("targetCidrRange")]
        public Output<ImmutableArray<string>> TargetCidrRange { get; private set; } = null!;

        /// <summary>
        /// Time when the internal range was updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// The type of usage set for this internal range.
        /// </summary>
        [Output("usage")]
        public Output<string> Usage { get; private set; } = null!;

        /// <summary>
        /// The list of resources that refer to this internal range. Resources that use the internal range for their range allocation are referred to as users of the range. Other resources mark themselves as users while doing so by creating a reference to this internal range. Having a user, based on this reference, prevents deletion of the internal range that is referred to. Can be empty.
        /// </summary>
        [Output("users")]
        public Output<ImmutableArray<string>> Users { get; private set; } = null!;


        /// <summary>
        /// Create a InternalRange resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InternalRange(string name, InternalRangeArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:networkconnectivity/v1alpha1:InternalRange", name, args ?? new InternalRangeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InternalRange(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:networkconnectivity/v1alpha1:InternalRange", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
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
        /// Get an existing InternalRange resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InternalRange Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new InternalRange(name, id, options);
        }
    }

    public sealed class InternalRangeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Time when the internal range was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// A description of this resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Optional. Resource ID (i.e. 'foo' in '[...]/projects/p/locations/l/internalRanges/foo') See https://google.aip.dev/122#resource-id-segments Unique per location.
        /// </summary>
        [Input("internalRangeId")]
        public Input<string>? InternalRangeId { get; set; }

        /// <summary>
        /// IP range that this internal range defines.
        /// </summary>
        [Input("ipCidrRange")]
        public Input<string>? IpCidrRange { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// User-defined labels.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Immutable. The name of an internal range. Format: projects/{project}/locations/{location}/internalRanges/{internal_range} See: https://google.aip.dev/122#fields-representing-resource-names
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The URL or resource ID of the network in which to reserve the internal range. The network cannot be deleted if there are any reserved internal ranges referring to it. Legacy networks are not supported. This can only be specified for a global internal address. Example: - URL: /compute/v1/projects/{project}/global/networks/{resourceId} - ID: network123
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        [Input("overlaps")]
        private InputList<Pulumi.GoogleNative.NetworkConnectivity.V1Alpha1.InternalRangeOverlapsItem>? _overlaps;

        /// <summary>
        /// Optional. Types of resources that are allowed to overlap with the current internal range.
        /// </summary>
        public InputList<Pulumi.GoogleNative.NetworkConnectivity.V1Alpha1.InternalRangeOverlapsItem> Overlaps
        {
            get => _overlaps ?? (_overlaps = new InputList<Pulumi.GoogleNative.NetworkConnectivity.V1Alpha1.InternalRangeOverlapsItem>());
            set => _overlaps = value;
        }

        /// <summary>
        /// The type of peering set for this internal range.
        /// </summary>
        [Input("peering")]
        public Input<Pulumi.GoogleNative.NetworkConnectivity.V1Alpha1.InternalRangePeering>? Peering { get; set; }

        /// <summary>
        /// An alternative to ip_cidr_range. Can be set when trying to create a reservation that automatically finds a free range of the given size. If both ip_cidr_range and prefix_length are set, there is an error if the range sizes do not match. Can also be used during updates to change the range size.
        /// </summary>
        [Input("prefixLength")]
        public Input<int>? PrefixLength { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if the original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        [Input("targetCidrRange")]
        private InputList<string>? _targetCidrRange;

        /// <summary>
        /// Optional. Can be set to narrow down or pick a different address space while searching for a free range. If not set, defaults to the "10.0.0.0/8" address space. This can be used to search in other rfc-1918 address spaces like "172.16.0.0/12" and "192.168.0.0/16" or non-rfc-1918 address spaces used in the VPC.
        /// </summary>
        public InputList<string> TargetCidrRange
        {
            get => _targetCidrRange ?? (_targetCidrRange = new InputList<string>());
            set => _targetCidrRange = value;
        }

        /// <summary>
        /// Time when the internal range was updated.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// The type of usage set for this internal range.
        /// </summary>
        [Input("usage")]
        public Input<Pulumi.GoogleNative.NetworkConnectivity.V1Alpha1.InternalRangeUsage>? Usage { get; set; }

        public InternalRangeArgs()
        {
        }
        public static new InternalRangeArgs Empty => new InternalRangeArgs();
    }
}
