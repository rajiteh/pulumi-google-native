// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// The TLS settings for the server.
    /// </summary>
    public sealed class ServerTlsSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configures the mechanism to obtain security certificates and identity information.
        /// </summary>
        [Input("proxyTlsContext")]
        public Input<Inputs.TlsContextArgs>? ProxyTlsContext { get; set; }

        [Input("subjectAltNames")]
        private InputList<string>? _subjectAltNames;

        /// <summary>
        /// A list of alternate names to verify the subject identity in the certificate presented by the client.
        /// </summary>
        public InputList<string> SubjectAltNames
        {
            get => _subjectAltNames ?? (_subjectAltNames = new InputList<string>());
            set => _subjectAltNames = value;
        }

        /// <summary>
        /// Indicates whether connections should be secured using TLS. The value of this field determines how TLS is enforced. This field can be set to one of the following: - SIMPLE Secure connections with standard TLS semantics. - MUTUAL Secure connections to the backends using mutual TLS by presenting client certificates for authentication. 
        /// </summary>
        [Input("tlsMode")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.ServerTlsSettingsTlsMode>? TlsMode { get; set; }

        public ServerTlsSettingsArgs()
        {
        }
        public static new ServerTlsSettingsArgs Empty => new ServerTlsSettingsArgs();
    }
}
