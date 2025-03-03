// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contactcenterinsights.V1
{
    /// <summary>
    /// Creates an issue model.
    /// </summary>
    [GoogleNativeResourceType("google-native:contactcenterinsights/v1:IssueModel")]
    public partial class IssueModel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The time at which this issue model was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The representative name for the issue model.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Configs for the input data that used to create the issue model.
        /// </summary>
        [Output("inputDataConfig")]
        public Output<Outputs.GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigResponse> InputDataConfig { get; private set; } = null!;

        /// <summary>
        /// Number of issues in this issue model.
        /// </summary>
        [Output("issueCount")]
        public Output<string> IssueCount { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Immutable. The resource name of the issue model. Format: projects/{project}/locations/{location}/issueModels/{issue_model}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// State of the model.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Immutable. The issue model's label statistics on its training data.
        /// </summary>
        [Output("trainingStats")]
        public Output<Outputs.GoogleCloudContactcenterinsightsV1IssueModelLabelStatsResponse> TrainingStats { get; private set; } = null!;

        /// <summary>
        /// The most recent time at which the issue model was updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a IssueModel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IssueModel(string name, IssueModelArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:contactcenterinsights/v1:IssueModel", name, args ?? new IssueModelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IssueModel(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:contactcenterinsights/v1:IssueModel", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing IssueModel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IssueModel Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IssueModel(name, id, options);
        }
    }

    public sealed class IssueModelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The representative name for the issue model.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Configs for the input data that used to create the issue model.
        /// </summary>
        [Input("inputDataConfig")]
        public Input<Inputs.GoogleCloudContactcenterinsightsV1IssueModelInputDataConfigArgs>? InputDataConfig { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Immutable. The resource name of the issue model. Format: projects/{project}/locations/{location}/issueModels/{issue_model}
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public IssueModelArgs()
        {
        }
        public static new IssueModelArgs Empty => new IssueModelArgs();
    }
}
