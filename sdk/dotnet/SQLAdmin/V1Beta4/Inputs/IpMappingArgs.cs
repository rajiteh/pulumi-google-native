// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1Beta4.Inputs
{

    /// <summary>
    /// Database instance IP Mapping.
    /// </summary>
    public sealed class IpMappingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP address assigned.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The due time for this IP to be retired in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`. This field is only available when the IP is scheduled to be retired.
        /// </summary>
        [Input("timeToRetire")]
        public Input<string>? TimeToRetire { get; set; }

        /// <summary>
        /// The type of this IP address. A `PRIMARY` address is a public address that can accept incoming connections. A `PRIVATE` address is a private address that can accept incoming connections. An `OUTGOING` address is the source address of connections originating from the instance, if supported.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.SQLAdmin.V1Beta4.IpMappingType>? Type { get; set; }

        public IpMappingArgs()
        {
        }
        public static new IpMappingArgs Empty => new IpMappingArgs();
    }
}
