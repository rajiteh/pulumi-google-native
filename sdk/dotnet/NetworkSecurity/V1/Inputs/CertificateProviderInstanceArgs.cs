// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkSecurity.V1.Inputs
{

    /// <summary>
    /// Specification of a TLS certificate provider instance. Workloads may have one or more CertificateProvider instances (plugins) and one of them is enabled and configured by specifying this message. Workloads use the values from this message to locate and load the CertificateProvider instance configuration.
    /// </summary>
    public sealed class CertificateProviderInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Plugin instance name, used to locate and load CertificateProvider instance configuration. Set to "google_cloud_private_spiffe" to use Certificate Authority Service certificate provider instance.
        /// </summary>
        [Input("pluginInstance", required: true)]
        public Input<string> PluginInstance { get; set; } = null!;

        public CertificateProviderInstanceArgs()
        {
        }
    }
}