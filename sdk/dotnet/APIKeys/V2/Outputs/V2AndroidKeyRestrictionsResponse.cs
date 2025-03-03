// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.APIKeys.V2.Outputs
{

    /// <summary>
    /// The Android apps that are allowed to use the key.
    /// </summary>
    [OutputType]
    public sealed class V2AndroidKeyRestrictionsResponse
    {
        /// <summary>
        /// A list of Android applications that are allowed to make API calls with this key.
        /// </summary>
        public readonly ImmutableArray<Outputs.V2AndroidApplicationResponse> AllowedApplications;

        [OutputConstructor]
        private V2AndroidKeyRestrictionsResponse(ImmutableArray<Outputs.V2AndroidApplicationResponse> allowedApplications)
        {
            AllowedApplications = allowedApplications;
        }
    }
}
