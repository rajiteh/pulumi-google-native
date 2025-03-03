// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Workstations.V1Beta.Outputs
{

    /// <summary>
    /// A set of Compute Engine Shielded instance options.
    /// </summary>
    [OutputType]
    public sealed class GceShieldedInstanceConfigResponse
    {
        /// <summary>
        /// Whether the instance has integrity monitoring enabled.
        /// </summary>
        public readonly bool EnableIntegrityMonitoring;
        /// <summary>
        /// Whether the instance has Secure Boot enabled.
        /// </summary>
        public readonly bool EnableSecureBoot;
        /// <summary>
        /// Whether the instance has the vTPM enabled.
        /// </summary>
        public readonly bool EnableVtpm;

        [OutputConstructor]
        private GceShieldedInstanceConfigResponse(
            bool enableIntegrityMonitoring,

            bool enableSecureBoot,

            bool enableVtpm)
        {
            EnableIntegrityMonitoring = enableIntegrityMonitoring;
            EnableSecureBoot = enableSecureBoot;
            EnableVtpm = enableVtpm;
        }
    }
}
