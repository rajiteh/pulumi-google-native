// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkSecurity.V1
{
    /// <summary>
    /// Creates a new ClientTlsPolicy in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:networksecurity/v1:ClientTlsPolicy")]
    public partial class ClientTlsPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
        /// </summary>
        [Output("clientCertificate")]
        public Output<Outputs.GoogleCloudNetworksecurityV1CertificateProviderResponse> ClientCertificate { get; private set; } = null!;

        /// <summary>
        /// Required. Short name of the ClientTlsPolicy resource to be created. This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. "client_mtls_policy".
        /// </summary>
        [Output("clientTlsPolicyId")]
        public Output<string> ClientTlsPolicyId { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the resource was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. Free-text description of the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Optional. Set of label tags associated with the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Name of the ClientTlsPolicy resource. It matches the pattern `projects/*/locations/{location}/clientTlsPolicies/{client_tls_policy}`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Optional. Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty, client does not validate the server certificate.
        /// </summary>
        [Output("serverValidationCa")]
        public Output<ImmutableArray<Outputs.ValidationCAResponse>> ServerValidationCa { get; private set; } = null!;

        /// <summary>
        /// Optional. Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".
        /// </summary>
        [Output("sni")]
        public Output<string> Sni { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the resource was updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a ClientTlsPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClientTlsPolicy(string name, ClientTlsPolicyArgs args, CustomResourceOptions? options = null)
            : base("google-native:networksecurity/v1:ClientTlsPolicy", name, args ?? new ClientTlsPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClientTlsPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:networksecurity/v1:ClientTlsPolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "clientTlsPolicyId",
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
        /// Get an existing ClientTlsPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClientTlsPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ClientTlsPolicy(name, id, options);
        }
    }

    public sealed class ClientTlsPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Defines a mechanism to provision client identity (public and private keys) for peer to peer authentication. The presence of this dictates mTLS.
        /// </summary>
        [Input("clientCertificate")]
        public Input<Inputs.GoogleCloudNetworksecurityV1CertificateProviderArgs>? ClientCertificate { get; set; }

        /// <summary>
        /// Required. Short name of the ClientTlsPolicy resource to be created. This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. "client_mtls_policy".
        /// </summary>
        [Input("clientTlsPolicyId", required: true)]
        public Input<string> ClientTlsPolicyId { get; set; } = null!;

        /// <summary>
        /// Optional. Free-text description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. Set of label tags associated with the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the ClientTlsPolicy resource. It matches the pattern `projects/*/locations/{location}/clientTlsPolicies/{client_tls_policy}`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("serverValidationCa")]
        private InputList<Inputs.ValidationCAArgs>? _serverValidationCa;

        /// <summary>
        /// Optional. Defines the mechanism to obtain the Certificate Authority certificate to validate the server certificate. If empty, client does not validate the server certificate.
        /// </summary>
        public InputList<Inputs.ValidationCAArgs> ServerValidationCa
        {
            get => _serverValidationCa ?? (_serverValidationCa = new InputList<Inputs.ValidationCAArgs>());
            set => _serverValidationCa = value;
        }

        /// <summary>
        /// Optional. Server Name Indication string to present to the server during TLS handshake. E.g: "secure.example.com".
        /// </summary>
        [Input("sni")]
        public Input<string>? Sni { get; set; }

        public ClientTlsPolicyArgs()
        {
        }
        public static new ClientTlsPolicyArgs Empty => new ClientTlsPolicyArgs();
    }
}
