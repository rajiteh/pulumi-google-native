// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3
{
    /// <summary>
    /// Creates a session entity type.
    /// </summary>
    [GoogleNativeResourceType("google-native:dialogflow/v3:SessionEntityType")]
    public partial class SessionEntityType : global::Pulumi.CustomResource
    {
        [Output("agentId")]
        public Output<string> AgentId { get; private set; } = null!;

        /// <summary>
        /// The collection of entities to override or supplement the custom entity type.
        /// </summary>
        [Output("entities")]
        public Output<ImmutableArray<Outputs.GoogleCloudDialogflowCxV3EntityTypeEntityResponse>> Entities { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the additional data should override or supplement the custom entity type definition.
        /// </summary>
        [Output("entityOverrideMode")]
        public Output<string> EntityOverrideMode { get; private set; } = null!;

        [Output("environmentId")]
        public Output<string> EnvironmentId { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("sessionId")]
        public Output<string> SessionId { get; private set; } = null!;


        /// <summary>
        /// Create a SessionEntityType resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SessionEntityType(string name, SessionEntityTypeArgs args, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v3:SessionEntityType", name, args ?? new SessionEntityTypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SessionEntityType(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dialogflow/v3:SessionEntityType", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "agentId",
                    "environmentId",
                    "location",
                    "project",
                    "sessionId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SessionEntityType resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SessionEntityType Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SessionEntityType(name, id, options);
        }
    }

    public sealed class SessionEntityTypeArgs : global::Pulumi.ResourceArgs
    {
        [Input("agentId", required: true)]
        public Input<string> AgentId { get; set; } = null!;

        [Input("entities", required: true)]
        private InputList<Inputs.GoogleCloudDialogflowCxV3EntityTypeEntityArgs>? _entities;

        /// <summary>
        /// The collection of entities to override or supplement the custom entity type.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowCxV3EntityTypeEntityArgs> Entities
        {
            get => _entities ?? (_entities = new InputList<Inputs.GoogleCloudDialogflowCxV3EntityTypeEntityArgs>());
            set => _entities = value;
        }

        /// <summary>
        /// Indicates whether the additional data should override or supplement the custom entity type definition.
        /// </summary>
        [Input("entityOverrideMode", required: true)]
        public Input<Pulumi.GoogleNative.Dialogflow.V3.SessionEntityTypeEntityOverrideMode> EntityOverrideMode { get; set; } = null!;

        [Input("environmentId", required: true)]
        public Input<string> EnvironmentId { get; set; } = null!;

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("sessionId", required: true)]
        public Input<string> SessionId { get; set; } = null!;

        public SessionEntityTypeArgs()
        {
        }
        public static new SessionEntityTypeArgs Empty => new SessionEntityTypeArgs();
    }
}
