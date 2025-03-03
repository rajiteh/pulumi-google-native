// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Inputs
{

    /// <summary>
    /// An occurrence of a particular package installation found within a system's filesystem. E.g., glibc was found in `/var/lib/dpkg/status`.
    /// </summary>
    public sealed class LocationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Deprecated. The CPE URI in [CPE format](https://cpe.mitre.org/specification/) denoting the package manager version distributing a package.
        /// </summary>
        [Input("cpeUri")]
        public Input<string>? CpeUri { get; set; }

        /// <summary>
        /// The path from which we gathered that this package/version is installed.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Deprecated. The version installed at this location.
        /// </summary>
        [Input("version")]
        public Input<Inputs.VersionArgs>? Version { get; set; }

        public LocationArgs()
        {
        }
        public static new LocationArgs Empty => new LocationArgs();
    }
}
