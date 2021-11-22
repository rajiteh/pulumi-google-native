// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1.Outputs
{

    /// <summary>
    /// An MSI package. MSI packages only support INSTALLED state.
    /// </summary>
    [OutputType]
    public sealed class OSPolicyResourcePackageResourceMSIResponse
    {
        /// <summary>
        /// Additional properties to use during installation. This should be in the format of Property=Setting. Appended to the defaults of `ACTION=INSTALL REBOOT=ReallySuppress`.
        /// </summary>
        public readonly ImmutableArray<string> Properties;
        /// <summary>
        /// The MSI package.
        /// </summary>
        public readonly Outputs.OSPolicyResourceFileResponse Source;

        [OutputConstructor]
        private OSPolicyResourcePackageResourceMSIResponse(
            ImmutableArray<string> properties,

            Outputs.OSPolicyResourceFileResponse source)
        {
            Properties = properties;
            Source = source;
        }
    }
}