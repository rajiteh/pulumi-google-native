// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1
{
    /// <summary>
    /// Creates an asset resource.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:dataplex/v1:Asset")]
    public partial class Asset : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Required. Asset identifier. This ID will be used to generate names such as table names when publishing metadata to Hive Metastore and BigQuery. * Must contain only lowercase letters, numbers and hyphens. * Must start with a letter. * Must end with a number or a letter. * Must be between 1-63 characters. * Must be unique within the zone.
        /// </summary>
        [Output("assetId")]
        public Output<string> AssetId { get; private set; } = null!;

        /// <summary>
        /// The time when the asset was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. Description of the asset.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Optional. Specification of the discovery feature applied to data referenced by this asset. When this spec is left unset, the asset will use the spec set on the parent zone.
        /// </summary>
        [Output("discoverySpec")]
        public Output<Outputs.GoogleCloudDataplexV1AssetDiscoverySpecResponse> DiscoverySpec { get; private set; } = null!;

        /// <summary>
        /// Status of the discovery feature applied to data referenced by this asset.
        /// </summary>
        [Output("discoveryStatus")]
        public Output<Outputs.GoogleCloudDataplexV1AssetDiscoveryStatusResponse> DiscoveryStatus { get; private set; } = null!;

        /// <summary>
        /// Optional. User friendly display name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Optional. User defined labels for the asset.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("lakeId")]
        public Output<string> LakeId { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The relative resource name of the asset, of the form: projects/{project_number}/locations/{location_id}/lakes/{lake_id}/zones/{zone_id}/assets/{asset_id}.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Specification of the resource that is referenced by this asset.
        /// </summary>
        [Output("resourceSpec")]
        public Output<Outputs.GoogleCloudDataplexV1AssetResourceSpecResponse> ResourceSpec { get; private set; } = null!;

        /// <summary>
        /// Status of the resource referenced by this asset.
        /// </summary>
        [Output("resourceStatus")]
        public Output<Outputs.GoogleCloudDataplexV1AssetResourceStatusResponse> ResourceStatus { get; private set; } = null!;

        /// <summary>
        /// Status of the security policy applied to resource referenced by this asset.
        /// </summary>
        [Output("securityStatus")]
        public Output<Outputs.GoogleCloudDataplexV1AssetSecurityStatusResponse> SecurityStatus { get; private set; } = null!;

        /// <summary>
        /// Current state of the asset.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// System generated globally unique ID for the asset. This ID will be different if the asset is deleted and re-created with the same name.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// The time when the asset was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a Asset resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Asset(string name, AssetArgs args, CustomResourceOptions? options = null)
            : base("google-native:dataplex/v1:Asset", name, args ?? new AssetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Asset(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dataplex/v1:Asset", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "assetId",
                    "lakeId",
                    "location",
                    "project",
                    "zone",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Asset resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Asset Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Asset(name, id, options);
        }
    }

    public sealed class AssetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. Asset identifier. This ID will be used to generate names such as table names when publishing metadata to Hive Metastore and BigQuery. * Must contain only lowercase letters, numbers and hyphens. * Must start with a letter. * Must end with a number or a letter. * Must be between 1-63 characters. * Must be unique within the zone.
        /// </summary>
        [Input("assetId", required: true)]
        public Input<string> AssetId { get; set; } = null!;

        /// <summary>
        /// Optional. Description of the asset.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Optional. Specification of the discovery feature applied to data referenced by this asset. When this spec is left unset, the asset will use the spec set on the parent zone.
        /// </summary>
        [Input("discoverySpec")]
        public Input<Inputs.GoogleCloudDataplexV1AssetDiscoverySpecArgs>? DiscoverySpec { get; set; }

        /// <summary>
        /// Optional. User friendly display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. User defined labels for the asset.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("lakeId", required: true)]
        public Input<string> LakeId { get; set; } = null!;

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Specification of the resource that is referenced by this asset.
        /// </summary>
        [Input("resourceSpec", required: true)]
        public Input<Inputs.GoogleCloudDataplexV1AssetResourceSpecArgs> ResourceSpec { get; set; } = null!;

        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public AssetArgs()
        {
        }
        public static new AssetArgs Empty => new AssetArgs();
    }
}
