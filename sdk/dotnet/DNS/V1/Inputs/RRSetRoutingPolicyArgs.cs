// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1.Inputs
{

    /// <summary>
    /// A RRSetRoutingPolicy represents ResourceRecordSet data that is returned dynamically with the response varying based on configured properties such as geolocation or by weighted random selection.
    /// </summary>
    public sealed class RRSetRoutingPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("geo")]
        public Input<Inputs.RRSetRoutingPolicyGeoPolicyArgs>? Geo { get; set; }

        [Input("kind")]
        public Input<string>? Kind { get; set; }

        [Input("primaryBackup")]
        public Input<Inputs.RRSetRoutingPolicyPrimaryBackupPolicyArgs>? PrimaryBackup { get; set; }

        [Input("wrr")]
        public Input<Inputs.RRSetRoutingPolicyWrrPolicyArgs>? Wrr { get; set; }

        public RRSetRoutingPolicyArgs()
        {
        }
        public static new RRSetRoutingPolicyArgs Empty => new RRSetRoutingPolicyArgs();
    }
}
