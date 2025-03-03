// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    public sealed class InstanceGroupManagerAutoHealingPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL for the health check that signals autohealing.
        /// </summary>
        [Input("healthCheck")]
        public Input<string>? HealthCheck { get; set; }

        /// <summary>
        /// The number of seconds that the managed instance group waits before it applies autohealing policies to new instances or recently recreated instances. This initial delay allows instances to initialize and run their startup scripts before the instance group determines that they are UNHEALTHY. This prevents the managed instance group from recreating its instances prematurely. This value must be from range [0, 3600].
        /// </summary>
        [Input("initialDelaySec")]
        public Input<int>? InitialDelaySec { get; set; }

        public InstanceGroupManagerAutoHealingPolicyArgs()
        {
        }
        public static new InstanceGroupManagerAutoHealingPolicyArgs Empty => new InstanceGroupManagerAutoHealingPolicyArgs();
    }
}
