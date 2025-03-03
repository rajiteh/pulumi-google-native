// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    public static class GetTargetHttpsProxy
    {
        /// <summary>
        /// Returns the specified TargetHttpsProxy resource.
        /// </summary>
        public static Task<GetTargetHttpsProxyResult> InvokeAsync(GetTargetHttpsProxyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTargetHttpsProxyResult>("google-native:compute/v1:getTargetHttpsProxy", args ?? new GetTargetHttpsProxyArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified TargetHttpsProxy resource.
        /// </summary>
        public static Output<GetTargetHttpsProxyResult> Invoke(GetTargetHttpsProxyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTargetHttpsProxyResult>("google-native:compute/v1:getTargetHttpsProxy", args ?? new GetTargetHttpsProxyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTargetHttpsProxyArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("targetHttpsProxy", required: true)]
        public string TargetHttpsProxy { get; set; } = null!;

        public GetTargetHttpsProxyArgs()
        {
        }
        public static new GetTargetHttpsProxyArgs Empty => new GetTargetHttpsProxyArgs();
    }

    public sealed class GetTargetHttpsProxyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("targetHttpsProxy", required: true)]
        public Input<string> TargetHttpsProxy { get; set; } = null!;

        public GetTargetHttpsProxyInvokeArgs()
        {
        }
        public static new GetTargetHttpsProxyInvokeArgs Empty => new GetTargetHttpsProxyInvokeArgs();
    }


    [OutputType]
    public sealed class GetTargetHttpsProxyResult
    {
        /// <summary>
        /// Optional. A URL referring to a networksecurity.AuthorizationPolicy resource that describes how the proxy should authorize inbound traffic. If left blank, access will not be restricted by an authorization policy. Refer to the AuthorizationPolicy resource for additional details. authorizationPolicy only applies to a global TargetHttpsProxy attached to globalForwardingRules with the loadBalancingScheme set to INTERNAL_SELF_MANAGED. Note: This field currently has no impact.
        /// </summary>
        public readonly string AuthorizationPolicy;
        /// <summary>
        /// URL of a certificate map that identifies a certificate map associated with the given target proxy. This field can only be set for global target proxies. If set, sslCertificates will be ignored. Accepted format is //certificatemanager.googleapis.com/projects/{project }/locations/{location}/certificateMaps/{resourceName}.
        /// </summary>
        public readonly string CertificateMap;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetHttpsProxy. An up-to-date fingerprint must be provided in order to patch the TargetHttpsProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetHttpsProxy.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// Type of resource. Always compute#targetHttpsProxy for target HTTPS proxies.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED. When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them. The default is false.
        /// </summary>
        public readonly bool ProxyBind;
        /// <summary>
        /// Specifies the QUIC override policy for this TargetHttpsProxy resource. This setting determines whether the load balancer attempts to negotiate QUIC with clients. You can specify NONE, ENABLE, or DISABLE. - When quic-override is set to NONE, Google manages whether QUIC is used. - When quic-override is set to ENABLE, the load balancer uses QUIC when possible. - When quic-override is set to DISABLE, the load balancer doesn't use QUIC. - If the quic-override flag is not specified, NONE is implied. 
        /// </summary>
        public readonly string QuicOverride;
        /// <summary>
        /// URL of the region where the regional TargetHttpsProxy resides. This field is not applicable to global TargetHttpsProxies.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Optional. A URL referring to a networksecurity.ServerTlsPolicy resource that describes how the proxy should authenticate inbound traffic. serverTlsPolicy only applies to a global TargetHttpsProxy attached to globalForwardingRules with the loadBalancingScheme set to INTERNAL_SELF_MANAGED or EXTERNAL or EXTERNAL_MANAGED. For details which ServerTlsPolicy resources are accepted with INTERNAL_SELF_MANAGED and which with EXTERNAL, EXTERNAL_MANAGED loadBalancingScheme consult ServerTlsPolicy documentation. If left blank, communications are not encrypted.
        /// </summary>
        public readonly string ServerTlsPolicy;
        /// <summary>
        /// URLs to SslCertificate resources that are used to authenticate connections between users and the load balancer. At least one SSL certificate must be specified. Currently, you may specify up to 15 SSL certificates. sslCertificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.
        /// </summary>
        public readonly ImmutableArray<string> SslCertificates;
        /// <summary>
        /// URL of SslPolicy resource that will be associated with the TargetHttpsProxy resource. If not set, the TargetHttpsProxy resource has no SSL policy configured.
        /// </summary>
        public readonly string SslPolicy;
        /// <summary>
        /// A fully-qualified or valid partial URL to the UrlMap resource that defines the mapping from URL to the BackendService. For example, the following are all valid URLs for specifying a URL map: - https://www.googleapis.compute/v1/projects/project/global/urlMaps/ url-map - projects/project/global/urlMaps/url-map - global/urlMaps/url-map 
        /// </summary>
        public readonly string UrlMap;

        [OutputConstructor]
        private GetTargetHttpsProxyResult(
            string authorizationPolicy,

            string certificateMap,

            string creationTimestamp,

            string description,

            string fingerprint,

            string kind,

            string name,

            bool proxyBind,

            string quicOverride,

            string region,

            string selfLink,

            string serverTlsPolicy,

            ImmutableArray<string> sslCertificates,

            string sslPolicy,

            string urlMap)
        {
            AuthorizationPolicy = authorizationPolicy;
            CertificateMap = certificateMap;
            CreationTimestamp = creationTimestamp;
            Description = description;
            Fingerprint = fingerprint;
            Kind = kind;
            Name = name;
            ProxyBind = proxyBind;
            QuicOverride = quicOverride;
            Region = region;
            SelfLink = selfLink;
            ServerTlsPolicy = serverTlsPolicy;
            SslCertificates = sslCertificates;
            SslPolicy = sslPolicy;
            UrlMap = urlMap;
        }
    }
}
