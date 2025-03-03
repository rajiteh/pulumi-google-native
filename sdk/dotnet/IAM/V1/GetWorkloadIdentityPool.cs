// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.IAM.V1
{
    public static class GetWorkloadIdentityPool
    {
        /// <summary>
        /// Gets an individual WorkloadIdentityPool.
        /// </summary>
        public static Task<GetWorkloadIdentityPoolResult> InvokeAsync(GetWorkloadIdentityPoolArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWorkloadIdentityPoolResult>("google-native:iam/v1:getWorkloadIdentityPool", args ?? new GetWorkloadIdentityPoolArgs(), options.WithDefaults());

        /// <summary>
        /// Gets an individual WorkloadIdentityPool.
        /// </summary>
        public static Output<GetWorkloadIdentityPoolResult> Invoke(GetWorkloadIdentityPoolInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWorkloadIdentityPoolResult>("google-native:iam/v1:getWorkloadIdentityPool", args ?? new GetWorkloadIdentityPoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWorkloadIdentityPoolArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("workloadIdentityPoolId", required: true)]
        public string WorkloadIdentityPoolId { get; set; } = null!;

        public GetWorkloadIdentityPoolArgs()
        {
        }
        public static new GetWorkloadIdentityPoolArgs Empty => new GetWorkloadIdentityPoolArgs();
    }

    public sealed class GetWorkloadIdentityPoolInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("workloadIdentityPoolId", required: true)]
        public Input<string> WorkloadIdentityPoolId { get; set; } = null!;

        public GetWorkloadIdentityPoolInvokeArgs()
        {
        }
        public static new GetWorkloadIdentityPoolInvokeArgs Empty => new GetWorkloadIdentityPoolInvokeArgs();
    }


    [OutputType]
    public sealed class GetWorkloadIdentityPoolResult
    {
        /// <summary>
        /// A description of the pool. Cannot exceed 256 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use existing tokens to access resources. If the pool is re-enabled, existing tokens grant access again.
        /// </summary>
        public readonly bool Disabled;
        /// <summary>
        /// A display name for the pool. Cannot exceed 32 characters.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The resource name of the pool.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The state of the pool.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetWorkloadIdentityPoolResult(
            string description,

            bool disabled,

            string displayName,

            string name,

            string state)
        {
            Description = description;
            Disabled = disabled;
            DisplayName = displayName;
            Name = name;
            State = state;
        }
    }
}
