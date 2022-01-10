// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V2.Inputs
{

    /// <summary>
    /// VPC Access settings. For more information on creating a VPC Connector, visit https://cloud.google.com/vpc/docs/configure-serverless-vpc-access For information on how to configure Cloud Run with an existing VPC Connector, visit https://cloud.google.com/run/docs/configuring/connecting-vpc
    /// </summary>
    public sealed class GoogleCloudRunOpV2VpcAccessArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// VPC Access connector name. Format: projects/{project}/locations/{location}/connectors/{connector}
        /// </summary>
        [Input("connector")]
        public Input<string>? Connector { get; set; }

        /// <summary>
        /// Traffic VPC egress settings.
        /// </summary>
        [Input("egress")]
        public Input<Pulumi.GoogleNative.Run.V2.GoogleCloudRunOpV2VpcAccessEgress>? Egress { get; set; }

        public GoogleCloudRunOpV2VpcAccessArgs()
        {
        }
    }
}