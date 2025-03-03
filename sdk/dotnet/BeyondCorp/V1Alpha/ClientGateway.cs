// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BeyondCorp.V1Alpha
{
    /// <summary>
    /// Creates a new ClientGateway in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:beyondcorp/v1alpha:ClientGateway")]
    public partial class ClientGateway : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The client connector service name that the client gateway is associated to. Client Connector Services, named as follows: `projects/{project_id}/locations/{location_id}/client_connector_services/{client_connector_service_id}`.
        /// </summary>
        [Output("clientConnectorService")]
        public Output<string> ClientConnectorService { get; private set; } = null!;

        /// <summary>
        /// Optional. User-settable client gateway resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
        /// </summary>
        [Output("clientGatewayId")]
        public Output<string?> ClientGatewayId { get; private set; } = null!;

        /// <summary>
        /// [Output only] Create time stamp.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// name of resource. The name is ignored during creation.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// The operational state of the gateway.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// [Output only] Update time stamp.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a ClientGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClientGateway(string name, ClientGatewayArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:beyondcorp/v1alpha:ClientGateway", name, args ?? new ClientGatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClientGateway(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:beyondcorp/v1alpha:ClientGateway", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ClientGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClientGateway Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ClientGateway(name, id, options);
        }
    }

    public sealed class ClientGatewayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. User-settable client gateway resource ID. * Must start with a letter. * Must contain between 4-63 characters from `/a-z-/`. * Must end with a number or a letter.
        /// </summary>
        [Input("clientGatewayId")]
        public Input<string>? ClientGatewayId { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// name of resource. The name is ignored during creation.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        public ClientGatewayArgs()
        {
        }
        public static new ClientGatewayArgs Empty => new ClientGatewayArgs();
    }
}
