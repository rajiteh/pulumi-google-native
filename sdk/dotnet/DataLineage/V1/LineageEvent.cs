// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataLineage.V1
{
    /// <summary>
    /// Creates a new lineage event.
    /// </summary>
    [GoogleNativeResourceType("google-native:datalineage/v1:LineageEvent")]
    public partial class LineageEvent : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. The end of the transformation which resulted in this lineage event. For streaming scenarios, it should be the end of the period from which the lineage is being reported.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// Optional. List of source-target pairs. Can't contain more than 100 tuples.
        /// </summary>
        [Output("links")]
        public Output<ImmutableArray<Outputs.GoogleCloudDatacatalogLineageV1EventLinkResponse>> Links { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Immutable. The resource name of the lineage event. Format: `projects/{project}/locations/{location}/processes/{process}/runs/{run}/lineageEvents/{lineage_event}`. Can be specified or auto-assigned. {lineage_event} must be not longer than 200 characters and only contain characters in a set: `a-zA-Z0-9_-:.`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("processId")]
        public Output<string> ProcessId { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// A unique identifier for this request. Restricted to 36 ASCII characters. A random UUID is recommended. This request is idempotent only if a `request_id` is provided.
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        [Output("runId")]
        public Output<string> RunId { get; private set; } = null!;

        /// <summary>
        /// Optional. The beginning of the transformation which resulted in this lineage event. For streaming scenarios, it should be the beginning of the period from which the lineage is being reported.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;


        /// <summary>
        /// Create a LineageEvent resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LineageEvent(string name, LineageEventArgs args, CustomResourceOptions? options = null)
            : base("google-native:datalineage/v1:LineageEvent", name, args ?? new LineageEventArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LineageEvent(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:datalineage/v1:LineageEvent", name, null, MakeResourceOptions(options, id))
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
                    "processId",
                    "project",
                    "runId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LineageEvent resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LineageEvent Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LineageEvent(name, id, options);
        }
    }

    public sealed class LineageEventArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The end of the transformation which resulted in this lineage event. For streaming scenarios, it should be the end of the period from which the lineage is being reported.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        [Input("links")]
        private InputList<Inputs.GoogleCloudDatacatalogLineageV1EventLinkArgs>? _links;

        /// <summary>
        /// Optional. List of source-target pairs. Can't contain more than 100 tuples.
        /// </summary>
        public InputList<Inputs.GoogleCloudDatacatalogLineageV1EventLinkArgs> Links
        {
            get => _links ?? (_links = new InputList<Inputs.GoogleCloudDatacatalogLineageV1EventLinkArgs>());
            set => _links = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Immutable. The resource name of the lineage event. Format: `projects/{project}/locations/{location}/processes/{process}/runs/{run}/lineageEvents/{lineage_event}`. Can be specified or auto-assigned. {lineage_event} must be not longer than 200 characters and only contain characters in a set: `a-zA-Z0-9_-:.`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("processId", required: true)]
        public Input<string> ProcessId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A unique identifier for this request. Restricted to 36 ASCII characters. A random UUID is recommended. This request is idempotent only if a `request_id` is provided.
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        [Input("runId", required: true)]
        public Input<string> RunId { get; set; } = null!;

        /// <summary>
        /// Optional. The beginning of the transformation which resulted in this lineage event. For streaming scenarios, it should be the beginning of the period from which the lineage is being reported.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public LineageEventArgs()
        {
        }
        public static new LineageEventArgs Empty => new LineageEventArgs();
    }
}
