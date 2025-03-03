// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Inputs
{

    /// <summary>
    /// Add-on configurations for the Apigee organization.
    /// </summary>
    public sealed class GoogleCloudApigeeV1AddonsConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration for the Advanced API Ops add-on.
        /// </summary>
        [Input("advancedApiOpsConfig")]
        public Input<Inputs.GoogleCloudApigeeV1AdvancedApiOpsConfigArgs>? AdvancedApiOpsConfig { get; set; }

        /// <summary>
        /// Configuration for the API Security add-on.
        /// </summary>
        [Input("apiSecurityConfig")]
        public Input<Inputs.GoogleCloudApigeeV1ApiSecurityConfigArgs>? ApiSecurityConfig { get; set; }

        /// <summary>
        /// Configuration for the Connectors Platform add-on.
        /// </summary>
        [Input("connectorsPlatformConfig")]
        public Input<Inputs.GoogleCloudApigeeV1ConnectorsPlatformConfigArgs>? ConnectorsPlatformConfig { get; set; }

        /// <summary>
        /// Configuration for the Integration add-on.
        /// </summary>
        [Input("integrationConfig")]
        public Input<Inputs.GoogleCloudApigeeV1IntegrationConfigArgs>? IntegrationConfig { get; set; }

        /// <summary>
        /// Configuration for the Monetization add-on.
        /// </summary>
        [Input("monetizationConfig")]
        public Input<Inputs.GoogleCloudApigeeV1MonetizationConfigArgs>? MonetizationConfig { get; set; }

        public GoogleCloudApigeeV1AddonsConfigArgs()
        {
        }
        public static new GoogleCloudApigeeV1AddonsConfigArgs Empty => new GoogleCloudApigeeV1AddonsConfigArgs();
    }
}
