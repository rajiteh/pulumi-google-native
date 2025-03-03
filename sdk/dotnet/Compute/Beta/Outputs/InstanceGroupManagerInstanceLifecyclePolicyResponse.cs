// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    [OutputType]
    public sealed class InstanceGroupManagerInstanceLifecyclePolicyResponse
    {
        /// <summary>
        /// A bit indicating whether to forcefully apply the group's latest configuration when repairing a VM. Valid options are: - NO (default): If configuration updates are available, they are not forcefully applied during repair. Instead, configuration updates are applied according to the group's update policy. - YES: If configuration updates are available, they are applied during repair. 
        /// </summary>
        public readonly string ForceUpdateOnRepair;

        [OutputConstructor]
        private InstanceGroupManagerInstanceLifecyclePolicyResponse(string forceUpdateOnRepair)
        {
            ForceUpdateOnRepair = forceUpdateOnRepair;
        }
    }
}
