// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkeonprem.V1.Inputs
{

    /// <summary>
    /// Represents VMware user cluster node's network configuration.
    /// </summary>
    public sealed class VmwareHostIpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Hostname of the machine. VM's name will be used if this field is empty.
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// IP could be an IP address (like 1.2.3.4) or a CIDR (like 1.2.3.0/24).
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        public VmwareHostIpArgs()
        {
        }
        public static new VmwareHostIpArgs Empty => new VmwareHostIpArgs();
    }
}