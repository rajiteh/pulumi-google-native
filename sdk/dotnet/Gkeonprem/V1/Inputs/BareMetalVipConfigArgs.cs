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
    /// Specifies the VIP config for the bare metal load balancer.
    /// </summary>
    public sealed class BareMetalVipConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The VIP which you previously set aside for the Kubernetes API of this bare metal user cluster.
        /// </summary>
        [Input("controlPlaneVip")]
        public Input<string>? ControlPlaneVip { get; set; }

        /// <summary>
        /// The VIP which you previously set aside for ingress traffic into this bare metal user cluster.
        /// </summary>
        [Input("ingressVip")]
        public Input<string>? IngressVip { get; set; }

        public BareMetalVipConfigArgs()
        {
        }
        public static new BareMetalVipConfigArgs Empty => new BareMetalVipConfigArgs();
    }
}