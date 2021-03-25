// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Dialogflow.V2
{
    /// <summary>
    /// Creates a knowledge base.
    /// </summary>
    [GoogleCloudResourceType("google-cloud:dialogflow/v2:GoogleCloudDialogflowV2KnowledgeBase")]
    public partial class GoogleCloudDialogflowV2KnowledgeBase : Pulumi.CustomResource
    {
        /// <summary>
        /// Create a GoogleCloudDialogflowV2KnowledgeBase resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GoogleCloudDialogflowV2KnowledgeBase(string name, GoogleCloudDialogflowV2KnowledgeBaseArgs args, CustomResourceOptions? options = null)
            : base("google-cloud:dialogflow/v2:GoogleCloudDialogflowV2KnowledgeBase", name, args ?? new GoogleCloudDialogflowV2KnowledgeBaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GoogleCloudDialogflowV2KnowledgeBase(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-cloud:dialogflow/v2:GoogleCloudDialogflowV2KnowledgeBase", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing GoogleCloudDialogflowV2KnowledgeBase resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GoogleCloudDialogflowV2KnowledgeBase Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GoogleCloudDialogflowV2KnowledgeBase(name, id, options);
        }
    }

    public sealed class GoogleCloudDialogflowV2KnowledgeBaseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The display name of the knowledge base. The name must be 1024 bytes or less; otherwise, the creation request fails.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Language which represents the KnowledgeBase. When the KnowledgeBase is created/updated, expect this to be present for non en-us languages. When unspecified, the default language code en-us applies.
        /// </summary>
        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        /// <summary>
        /// The knowledge base resource name. The name must be empty when creating a knowledge base. Format: `projects//locations//knowledgeBases/`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Required. The project to create a knowledge base for. Format: `projects//locations/`.
        /// </summary>
        [Input("parent", required: true)]
        public Input<string> Parent { get; set; } = null!;

        public GoogleCloudDialogflowV2KnowledgeBaseArgs()
        {
        }
    }
}