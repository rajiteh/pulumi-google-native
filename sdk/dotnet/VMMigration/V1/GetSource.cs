// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.VMMigration.V1
{
    public static class GetSource
    {
        /// <summary>
        /// Gets details of a single Source.
        /// </summary>
        public static Task<GetSourceResult> InvokeAsync(GetSourceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSourceResult>("google-native:vmmigration/v1:getSource", args ?? new GetSourceArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single Source.
        /// </summary>
        public static Output<GetSourceResult> Invoke(GetSourceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSourceResult>("google-native:vmmigration/v1:getSource", args ?? new GetSourceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSourceArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetSourceArgs()
        {
        }
        public static new GetSourceArgs Empty => new GetSourceArgs();
    }

    public sealed class GetSourceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetSourceInvokeArgs()
        {
        }
        public static new GetSourceInvokeArgs Empty => new GetSourceInvokeArgs();
    }


    [OutputType]
    public sealed class GetSourceResult
    {
        /// <summary>
        /// AWS type source details.
        /// </summary>
        public readonly Outputs.AwsSourceDetailsResponse Aws;
        /// <summary>
        /// The create time timestamp.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// User-provided description of the source.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The labels of the source.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The Source name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The update time timestamp.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// Vmware type source details.
        /// </summary>
        public readonly Outputs.VmwareSourceDetailsResponse Vmware;

        [OutputConstructor]
        private GetSourceResult(
            Outputs.AwsSourceDetailsResponse aws,

            string createTime,

            string description,

            ImmutableDictionary<string, string> labels,

            string name,

            string updateTime,

            Outputs.VmwareSourceDetailsResponse vmware)
        {
            Aws = aws;
            CreateTime = createTime;
            Description = description;
            Labels = labels;
            Name = name;
            UpdateTime = updateTime;
            Vmware = vmware;
        }
    }
}
