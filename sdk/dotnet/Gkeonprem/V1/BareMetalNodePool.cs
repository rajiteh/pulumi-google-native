// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkeonprem.V1
{
    /// <summary>
    /// Creates a new bare metal node pool in a given project, location and Bare Metal cluster.
    /// </summary>
    [GoogleNativeResourceType("google-native:gkeonprem/v1:BareMetalNodePool")]
    public partial class BareMetalNodePool : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Annotations on the bare metal node pool. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
        /// </summary>
        [Output("annotations")]
        public Output<ImmutableDictionary<string, string>> Annotations { get; private set; } = null!;

        [Output("bareMetalClusterId")]
        public Output<string> BareMetalClusterId { get; private set; } = null!;

        /// <summary>
        /// The ID to use for the node pool, which will become the final component of the node pool's resource name. This value must be up to 63 characters, and valid characters are /a-z-/. The value must not be permitted to be a UUID (or UUID-like: anything matching /^[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$/i).
        /// </summary>
        [Output("bareMetalNodePoolId")]
        public Output<string?> BareMetalNodePoolId { get; private set; } = null!;

        /// <summary>
        /// The time at which this bare metal node pool was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The time at which this bare metal node pool was deleted. If the resource is not deleted, this must be empty
        /// </summary>
        [Output("deleteTime")]
        public Output<string> DeleteTime { get; private set; } = null!;

        /// <summary>
        /// The display name for the bare metal node pool.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Allows clients to perform consistent read-modify-writes through optimistic concurrency control.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Immutable. The bare metal node pool resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Node pool configuration.
        /// </summary>
        [Output("nodePoolConfig")]
        public Output<Outputs.BareMetalNodePoolConfigResponse> NodePoolConfig { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// If set, there are currently changes in flight to the bare metal node pool.
        /// </summary>
        [Output("reconciling")]
        public Output<bool> Reconciling { get; private set; } = null!;

        /// <summary>
        /// The current state of the bare metal node pool.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// ResourceStatus representing the detailed node pool status.
        /// </summary>
        [Output("status")]
        public Output<Outputs.ResourceStatusResponse> Status { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the bare metal node pool.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// The time at which this bare metal node pool was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a BareMetalNodePool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BareMetalNodePool(string name, BareMetalNodePoolArgs args, CustomResourceOptions? options = null)
            : base("google-native:gkeonprem/v1:BareMetalNodePool", name, args ?? new BareMetalNodePoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BareMetalNodePool(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:gkeonprem/v1:BareMetalNodePool", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "bareMetalClusterId",
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
        /// Get an existing BareMetalNodePool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BareMetalNodePool Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BareMetalNodePool(name, id, options);
        }
    }

    public sealed class BareMetalNodePoolArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// Annotations on the bare metal node pool. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
        /// </summary>
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        [Input("bareMetalClusterId", required: true)]
        public Input<string> BareMetalClusterId { get; set; } = null!;

        /// <summary>
        /// The ID to use for the node pool, which will become the final component of the node pool's resource name. This value must be up to 63 characters, and valid characters are /a-z-/. The value must not be permitted to be a UUID (or UUID-like: anything matching /^[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}$/i).
        /// </summary>
        [Input("bareMetalNodePoolId")]
        public Input<string>? BareMetalNodePoolId { get; set; }

        /// <summary>
        /// The display name for the bare metal node pool.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Allows clients to perform consistent read-modify-writes through optimistic concurrency control.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Immutable. The bare metal node pool resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Node pool configuration.
        /// </summary>
        [Input("nodePoolConfig", required: true)]
        public Input<Inputs.BareMetalNodePoolConfigArgs> NodePoolConfig { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public BareMetalNodePoolArgs()
        {
        }
        public static new BareMetalNodePoolArgs Empty => new BareMetalNodePoolArgs();
    }
}