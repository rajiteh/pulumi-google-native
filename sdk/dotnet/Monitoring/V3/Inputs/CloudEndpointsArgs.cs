// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V3.Inputs
{

    /// <summary>
    /// Cloud Endpoints service. Learn more at https://cloud.google.com/endpoints.
    /// </summary>
    public sealed class CloudEndpointsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Cloud Endpoints service underlying this service. Corresponds to the service resource label in the api monitored resource (https://cloud.google.com/monitoring/api/resources#tag_api).
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        public CloudEndpointsArgs()
        {
        }
        public static new CloudEndpointsArgs Empty => new CloudEndpointsArgs();
    }
}
