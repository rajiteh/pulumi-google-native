// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Logging.V2
{
    /// <summary>
    /// Creates a sink that exports specified log entries to a destination. The export of newly-ingested log entries begins immediately, unless the sink's writer_identity is not permitted to write to the destination. A sink can export log entries only from the resource owning the sink.
    /// </summary>
    [GoogleNativeResourceType("google-native:logging/v2:FolderSink")]
    public partial class FolderSink : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. Options that affect sinks exporting data to BigQuery.
        /// </summary>
        [Output("bigqueryOptions")]
        public Output<Outputs.BigQueryOptionsResponse> BigqueryOptions { get; private set; } = null!;

        /// <summary>
        /// The creation timestamp of the sink.This field may not be present for older sinks.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. A description of this sink.The maximum length of the description is 8000 characters.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The export destination: "storage.googleapis.com/[GCS_BUCKET]" "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]" "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]" The sink's writer_identity, set when the sink is created, must have permission to write to the destination or else the log entries are not exported. For more information, see Exporting Logs with Sinks (https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
        /// </summary>
        [Output("destination")]
        public Output<string> Destination { get; private set; } = null!;

        /// <summary>
        /// Optional. If set to true, then this sink is disabled and it does not export any log entries.
        /// </summary>
        [Output("disabled")]
        public Output<bool> Disabled { get; private set; } = null!;

        /// <summary>
        /// Optional. Log entries that match any of these exclusion filters will not be exported.If a log entry is matched by both filter and one of exclusion_filters it will not be exported.
        /// </summary>
        [Output("exclusions")]
        public Output<ImmutableArray<Outputs.LogExclusionResponse>> Exclusions { get; private set; } = null!;

        /// <summary>
        /// Optional. An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries). The only exported log entries are those that are in the resource owning the sink and that match the filter.For example:logName="projects/[PROJECT_ID]/logs/[LOG_ID]" AND severity&gt;=ERROR
        /// </summary>
        [Output("filter")]
        public Output<string> Filter { get; private set; } = null!;

        [Output("folderId")]
        public Output<string> FolderId { get; private set; } = null!;

        /// <summary>
        /// Optional. This field applies only to sinks owned by organizations and folders. If the field is false, the default, only the logs owned by the sink's parent resource are available for export. If the field is true, then log entries from all the projects, folders, and billing accounts contained in the sink's parent resource are also available for export. Whether a particular log entry from the children is exported depends on the sink's filter expression.For example, if this field is true, then the filter resource.type=gce_instance would export all Compute Engine VM instance log entries from all projects in the sink's parent.To only export entries from certain child projects, filter on the project part of the log name:logName:("projects/test-project1/" OR "projects/test-project2/") AND resource.type=gce_instance
        /// </summary>
        [Output("includeChildren")]
        public Output<bool> IncludeChildren { get; private set; } = null!;

        /// <summary>
        /// The client-assigned sink identifier, unique within the project.For example: "my-syslog-errors-to-pubsub". Sink identifiers are limited to 100 characters and can include only the following characters: upper and lower-case alphanumeric characters, underscores, hyphens, and periods. First character has to be alphanumeric.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Deprecated. This field is unused.
        /// </summary>
        [Output("outputVersionFormat")]
        public Output<string> OutputVersionFormat { get; private set; } = null!;

        /// <summary>
        /// Optional. Determines the kind of IAM identity returned as writer_identity in the new sink. If this value is omitted or set to false, and if the sink's parent is a project, then the value returned as writer_identity is the same group or service account used by Cloud Logging before the addition of writer identities to this API. The sink's destination must be in the same project as the sink itself.If this field is set to true, or if the sink is owned by a non-project resource such as an organization, then the value of writer_identity will be a unique service account used only for exports from the new sink. For more information, see writer_identity in LogSink.
        /// </summary>
        [Output("uniqueWriterIdentity")]
        public Output<bool?> UniqueWriterIdentity { get; private set; } = null!;

        /// <summary>
        /// The last update timestamp of the sink.This field may not be present for older sinks.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// An IAM identity—a service account or group—under which Cloud Logging writes the exported log entries to the sink's destination. This field is either set by specifying custom_writer_identity or set automatically by sinks.create and sinks.update based on the value of unique_writer_identity in those methods.Until you grant this identity write-access to the destination, log entry exports from this sink will fail. For more information, see Granting Access for a Resource (https://cloud.google.com/iam/docs/granting-roles-to-service-accounts#granting_access_to_a_service_account_for_a_resource). Consult the destination service's documentation to determine the appropriate IAM roles to assign to the identity.Sinks that have a destination that is a log bucket in the same project as the sink cannot have a writer_identity and no additional permissions are required.
        /// </summary>
        [Output("writerIdentity")]
        public Output<string> WriterIdentity { get; private set; } = null!;


        /// <summary>
        /// Create a FolderSink resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FolderSink(string name, FolderSinkArgs args, CustomResourceOptions? options = null)
            : base("google-native:logging/v2:FolderSink", name, args ?? new FolderSinkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FolderSink(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:logging/v2:FolderSink", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "folderId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FolderSink resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FolderSink Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FolderSink(name, id, options);
        }
    }

    public sealed class FolderSinkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Options that affect sinks exporting data to BigQuery.
        /// </summary>
        [Input("bigqueryOptions")]
        public Input<Inputs.BigQueryOptionsArgs>? BigqueryOptions { get; set; }

        /// <summary>
        /// Optional. A description of this sink.The maximum length of the description is 8000 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The export destination: "storage.googleapis.com/[GCS_BUCKET]" "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]" "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]" The sink's writer_identity, set when the sink is created, must have permission to write to the destination or else the log entries are not exported. For more information, see Exporting Logs with Sinks (https://cloud.google.com/logging/docs/api/tasks/exporting-logs).
        /// </summary>
        [Input("destination", required: true)]
        public Input<string> Destination { get; set; } = null!;

        /// <summary>
        /// Optional. If set to true, then this sink is disabled and it does not export any log entries.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        [Input("exclusions")]
        private InputList<Inputs.LogExclusionArgs>? _exclusions;

        /// <summary>
        /// Optional. Log entries that match any of these exclusion filters will not be exported.If a log entry is matched by both filter and one of exclusion_filters it will not be exported.
        /// </summary>
        public InputList<Inputs.LogExclusionArgs> Exclusions
        {
            get => _exclusions ?? (_exclusions = new InputList<Inputs.LogExclusionArgs>());
            set => _exclusions = value;
        }

        /// <summary>
        /// Optional. An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries). The only exported log entries are those that are in the resource owning the sink and that match the filter.For example:logName="projects/[PROJECT_ID]/logs/[LOG_ID]" AND severity&gt;=ERROR
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        [Input("folderId", required: true)]
        public Input<string> FolderId { get; set; } = null!;

        /// <summary>
        /// Optional. This field applies only to sinks owned by organizations and folders. If the field is false, the default, only the logs owned by the sink's parent resource are available for export. If the field is true, then log entries from all the projects, folders, and billing accounts contained in the sink's parent resource are also available for export. Whether a particular log entry from the children is exported depends on the sink's filter expression.For example, if this field is true, then the filter resource.type=gce_instance would export all Compute Engine VM instance log entries from all projects in the sink's parent.To only export entries from certain child projects, filter on the project part of the log name:logName:("projects/test-project1/" OR "projects/test-project2/") AND resource.type=gce_instance
        /// </summary>
        [Input("includeChildren")]
        public Input<bool>? IncludeChildren { get; set; }

        /// <summary>
        /// The client-assigned sink identifier, unique within the project.For example: "my-syslog-errors-to-pubsub". Sink identifiers are limited to 100 characters and can include only the following characters: upper and lower-case alphanumeric characters, underscores, hyphens, and periods. First character has to be alphanumeric.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Deprecated. This field is unused.
        /// </summary>
        [Input("outputVersionFormat")]
        public Input<Pulumi.GoogleNative.Logging.V2.FolderSinkOutputVersionFormat>? OutputVersionFormat { get; set; }

        /// <summary>
        /// Optional. Determines the kind of IAM identity returned as writer_identity in the new sink. If this value is omitted or set to false, and if the sink's parent is a project, then the value returned as writer_identity is the same group or service account used by Cloud Logging before the addition of writer identities to this API. The sink's destination must be in the same project as the sink itself.If this field is set to true, or if the sink is owned by a non-project resource such as an organization, then the value of writer_identity will be a unique service account used only for exports from the new sink. For more information, see writer_identity in LogSink.
        /// </summary>
        [Input("uniqueWriterIdentity")]
        public Input<bool>? UniqueWriterIdentity { get; set; }

        public FolderSinkArgs()
        {
        }
        public static new FolderSinkArgs Empty => new FolderSinkArgs();
    }
}
