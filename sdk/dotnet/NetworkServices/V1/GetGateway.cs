// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1
{
    public static class GetGateway
    {
        /// <summary>
        /// Gets details of a single Gateway.
        /// </summary>
        public static Task<GetGatewayResult> InvokeAsync(GetGatewayArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGatewayResult>("google-native:networkservices/v1:getGateway", args ?? new GetGatewayArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single Gateway.
        /// </summary>
        public static Output<GetGatewayResult> Invoke(GetGatewayInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGatewayResult>("google-native:networkservices/v1:getGateway", args ?? new GetGatewayInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGatewayArgs : global::Pulumi.InvokeArgs
    {
        [Input("gatewayId", required: true)]
        public string GatewayId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetGatewayArgs()
        {
        }
        public static new GetGatewayArgs Empty => new GetGatewayArgs();
    }

    public sealed class GetGatewayInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("gatewayId", required: true)]
        public Input<string> GatewayId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetGatewayInvokeArgs()
        {
        }
        public static new GetGatewayInvokeArgs Empty => new GetGatewayInvokeArgs();
    }


    [OutputType]
    public sealed class GetGatewayResult
    {
        /// <summary>
        /// Optional. Zero or one IPv4-address on which the Gateway will receive the traffic. When no address is provided, an IP from the subnetwork is allocated This field only applies to gateways of type 'SECURE_WEB_GATEWAY'. Gateways of type 'OPEN_MESH' listen on 0.0.0.0.
        /// </summary>
        public readonly ImmutableArray<string> Addresses;
        /// <summary>
        /// Optional. A fully-qualified Certificates URL reference. The proxy presents a Certificate (selected based on SNI) when establishing a TLS connection. This feature only applies to gateways of type 'SECURE_WEB_GATEWAY'.
        /// </summary>
        public readonly ImmutableArray<string> CertificateUrls;
        /// <summary>
        /// The timestamp when the resource was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. A free-text description of the resource. Max length 1024 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. A fully-qualified GatewaySecurityPolicy URL reference. Defines how a server should apply security policy to inbound (VM to Proxy) initiated connections. For example: `projects/*/locations/*/gatewaySecurityPolicies/swg-policy`. This policy is specific to gateways of type 'SECURE_WEB_GATEWAY'.
        /// </summary>
        public readonly string GatewaySecurityPolicy;
        /// <summary>
        /// Optional. Set of label tags associated with the Gateway resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Name of the Gateway resource. It matches pattern `projects/*/locations/*/gateways/`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Optional. The relative resource name identifying the VPC network that is using this configuration. For example: `projects/*/global/networks/network-1`. Currently, this field is specific to gateways of type 'SECURE_WEB_GATEWAY'.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// One or more port numbers (1-65535), on which the Gateway will receive traffic. The proxy binds to the specified ports. Gateways of type 'SECURE_WEB_GATEWAY' are limited to 1 port. Gateways of type 'OPEN_MESH' listen on 0.0.0.0 and support multiple ports.
        /// </summary>
        public readonly ImmutableArray<int> Ports;
        /// <summary>
        /// Optional. Scope determines how configuration across multiple Gateway instances are merged. The configuration for multiple Gateway instances with the same scope will be merged as presented as a single coniguration to the proxy/load balancer. Max length 64 characters. Scope should start with a letter and can only have letters, numbers, hyphens.
        /// </summary>
        public readonly string Scope;
        /// <summary>
        /// Server-defined URL of this resource
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Optional. A fully-qualified ServerTLSPolicy URL reference. Specifies how TLS traffic is terminated. If empty, TLS termination is disabled.
        /// </summary>
        public readonly string ServerTlsPolicy;
        /// <summary>
        /// Optional. The relative resource name identifying the subnetwork in which this SWG is allocated. For example: `projects/*/regions/us-central1/subnetworks/network-1` Currently, this field is specific to gateways of type 'SECURE_WEB_GATEWAY".
        /// </summary>
        public readonly string Subnetwork;
        /// <summary>
        /// Immutable. The type of the customer managed gateway. This field is required. If unspecified, an error is returned.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The timestamp when the resource was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetGatewayResult(
            ImmutableArray<string> addresses,

            ImmutableArray<string> certificateUrls,

            string createTime,

            string description,

            string gatewaySecurityPolicy,

            ImmutableDictionary<string, string> labels,

            string name,

            string network,

            ImmutableArray<int> ports,

            string scope,

            string selfLink,

            string serverTlsPolicy,

            string subnetwork,

            string type,

            string updateTime)
        {
            Addresses = addresses;
            CertificateUrls = certificateUrls;
            CreateTime = createTime;
            Description = description;
            GatewaySecurityPolicy = gatewaySecurityPolicy;
            Labels = labels;
            Name = name;
            Network = network;
            Ports = ports;
            Scope = scope;
            SelfLink = selfLink;
            ServerTlsPolicy = serverTlsPolicy;
            Subnetwork = subnetwork;
            Type = type;
            UpdateTime = updateTime;
        }
    }
}
