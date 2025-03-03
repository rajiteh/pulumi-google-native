// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contactcenterinsights.V1
{
    public static class GetView
    {
        /// <summary>
        /// Gets a view.
        /// </summary>
        public static Task<GetViewResult> InvokeAsync(GetViewArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetViewResult>("google-native:contactcenterinsights/v1:getView", args ?? new GetViewArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a view.
        /// </summary>
        public static Output<GetViewResult> Invoke(GetViewInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetViewResult>("google-native:contactcenterinsights/v1:getView", args ?? new GetViewInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetViewArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("viewId", required: true)]
        public string ViewId { get; set; } = null!;

        public GetViewArgs()
        {
        }
        public static new GetViewArgs Empty => new GetViewArgs();
    }

    public sealed class GetViewInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("viewId", required: true)]
        public Input<string> ViewId { get; set; } = null!;

        public GetViewInvokeArgs()
        {
        }
        public static new GetViewInvokeArgs Empty => new GetViewInvokeArgs();
    }


    [OutputType]
    public sealed class GetViewResult
    {
        /// <summary>
        /// The time at which this view was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The human-readable display name of the view.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Immutable. The resource name of the view. Format: projects/{project}/locations/{location}/views/{view}
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The most recent time at which the view was updated.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// String with specific view properties, must be non-empty.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetViewResult(
            string createTime,

            string displayName,

            string name,

            string updateTime,

            string value)
        {
            CreateTime = createTime;
            DisplayName = displayName;
            Name = name;
            UpdateTime = updateTime;
            Value = value;
        }
    }
}
