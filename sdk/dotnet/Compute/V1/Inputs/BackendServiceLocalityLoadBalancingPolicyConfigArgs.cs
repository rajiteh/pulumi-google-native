// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// Container for either a built-in LB policy supported by gRPC or Envoy or a custom one implemented by the end user.
    /// </summary>
    public sealed class BackendServiceLocalityLoadBalancingPolicyConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("customPolicy")]
        public Input<Inputs.BackendServiceLocalityLoadBalancingPolicyConfigCustomPolicyArgs>? CustomPolicy { get; set; }

        [Input("policy")]
        public Input<Inputs.BackendServiceLocalityLoadBalancingPolicyConfigPolicyArgs>? Policy { get; set; }

        public BackendServiceLocalityLoadBalancingPolicyConfigArgs()
        {
        }
        public static new BackendServiceLocalityLoadBalancingPolicyConfigArgs Empty => new BackendServiceLocalityLoadBalancingPolicyConfigArgs();
    }
}
