// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkConnectivity.V1
{
    public static class GetServiceConnectionPolicy
    {
        /// <summary>
        /// Gets details of a single ServiceConnectionPolicy.
        /// </summary>
        public static Task<GetServiceConnectionPolicyResult> InvokeAsync(GetServiceConnectionPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServiceConnectionPolicyResult>("google-native:networkconnectivity/v1:getServiceConnectionPolicy", args ?? new GetServiceConnectionPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single ServiceConnectionPolicy.
        /// </summary>
        public static Output<GetServiceConnectionPolicyResult> Invoke(GetServiceConnectionPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServiceConnectionPolicyResult>("google-native:networkconnectivity/v1:getServiceConnectionPolicy", args ?? new GetServiceConnectionPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServiceConnectionPolicyArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("serviceConnectionPolicyId", required: true)]
        public string ServiceConnectionPolicyId { get; set; } = null!;

        public GetServiceConnectionPolicyArgs()
        {
        }
        public static new GetServiceConnectionPolicyArgs Empty => new GetServiceConnectionPolicyArgs();
    }

    public sealed class GetServiceConnectionPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("serviceConnectionPolicyId", required: true)]
        public Input<string> ServiceConnectionPolicyId { get; set; } = null!;

        public GetServiceConnectionPolicyInvokeArgs()
        {
        }
        public static new GetServiceConnectionPolicyInvokeArgs Empty => new GetServiceConnectionPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetServiceConnectionPolicyResult
    {
        /// <summary>
        /// Time when the ServiceConnectionMap was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// A description of this resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The type of underlying resources used to create the connection.
        /// </summary>
        public readonly string Infrastructure;
        /// <summary>
        /// User-defined labels.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Immutable. The name of a ServiceConnectionPolicy. Format: projects/{project}/locations/{location}/serviceConnectionPolicies/{service_connection_policy} See: https://google.aip.dev/122#fields-representing-resource-names
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The resource path of the consumer network. Example: - projects/{projectNumOrId}/global/networks/{resourceId}.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// Configuration used for Private Service Connect connections. Used when Infrastructure is PSC.
        /// </summary>
        public readonly Outputs.PscConfigResponse PscConfig;
        /// <summary>
        /// [Output only] Information about each Private Service Connect connection.
        /// </summary>
        public readonly ImmutableArray<Outputs.PscConnectionResponse> PscConnections;
        /// <summary>
        /// The service class identifier for which this ServiceConnectionPolicy is for. The service class identifier is a unique, symbolic representation of a ServiceClass. It is provided by the Service Producer. Google services have a prefix of gcp. For example, gcp-cloud-sql. 3rd party services do not. For example, test-service-a3dfcx.
        /// </summary>
        public readonly string ServiceClass;
        /// <summary>
        /// Time when the ServiceConnectionMap was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetServiceConnectionPolicyResult(
            string createTime,

            string description,

            string infrastructure,

            ImmutableDictionary<string, string> labels,

            string name,

            string network,

            Outputs.PscConfigResponse pscConfig,

            ImmutableArray<Outputs.PscConnectionResponse> pscConnections,

            string serviceClass,

            string updateTime)
        {
            CreateTime = createTime;
            Description = description;
            Infrastructure = infrastructure;
            Labels = labels;
            Name = name;
            Network = network;
            PscConfig = pscConfig;
            PscConnections = pscConnections;
            ServiceClass = serviceClass;
            UpdateTime = updateTime;
        }
    }
}
