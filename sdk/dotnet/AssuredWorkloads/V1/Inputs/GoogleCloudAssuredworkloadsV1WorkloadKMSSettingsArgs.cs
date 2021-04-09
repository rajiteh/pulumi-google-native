// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.AssuredWorkloads.V1.Inputs
{

    /// <summary>
    /// Settings specific to the Key Management Service.
    /// </summary>
    public sealed class GoogleCloudAssuredworkloadsV1WorkloadKMSSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. Input only. Immutable. The time at which the Key Management Service will automatically create a new version of the crypto key and mark it as the primary.
        /// </summary>
        [Input("nextRotationTime")]
        public Input<string>? NextRotationTime { get; set; }

        /// <summary>
        /// Required. Input only. Immutable. [next_rotation_time] will be advanced by this period when the Key Management Service automatically rotates a key. Must be at least 24 hours and at most 876,000 hours.
        /// </summary>
        [Input("rotationPeriod")]
        public Input<string>? RotationPeriod { get; set; }

        public GoogleCloudAssuredworkloadsV1WorkloadKMSSettingsArgs()
        {
        }
    }
}