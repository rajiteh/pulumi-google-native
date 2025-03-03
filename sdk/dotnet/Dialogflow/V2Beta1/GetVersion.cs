// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1
{
    public static class GetVersion
    {
        /// <summary>
        /// Retrieves the specified agent version.
        /// </summary>
        public static Task<GetVersionResult> InvokeAsync(GetVersionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVersionResult>("google-native:dialogflow/v2beta1:getVersion", args ?? new GetVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the specified agent version.
        /// </summary>
        public static Output<GetVersionResult> Invoke(GetVersionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVersionResult>("google-native:dialogflow/v2beta1:getVersion", args ?? new GetVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVersionArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("versionId", required: true)]
        public string VersionId { get; set; } = null!;

        public GetVersionArgs()
        {
        }
        public static new GetVersionArgs Empty => new GetVersionArgs();
    }

    public sealed class GetVersionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("versionId", required: true)]
        public Input<string> VersionId { get; set; } = null!;

        public GetVersionInvokeArgs()
        {
        }
        public static new GetVersionInvokeArgs Empty => new GetVersionInvokeArgs();
    }


    [OutputType]
    public sealed class GetVersionResult
    {
        /// <summary>
        /// The creation time of this version. This field is read-only, i.e., it cannot be set by create and update methods.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. The developer-provided description of this version.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The unique identifier of this agent version. Supported formats: - `projects//agent/versions/` - `projects//locations//agent/versions/`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The status of this version. This field is read-only and cannot be set by create and update methods.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The sequential number of this version. This field is read-only which means it cannot be set by create and update methods.
        /// </summary>
        public readonly int VersionNumber;

        [OutputConstructor]
        private GetVersionResult(
            string createTime,

            string description,

            string name,

            string status,

            int versionNumber)
        {
            CreateTime = createTime;
            Description = description;
            Name = name;
            Status = status;
            VersionNumber = versionNumber;
        }
    }
}
