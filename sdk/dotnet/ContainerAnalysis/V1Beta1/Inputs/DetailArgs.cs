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
    /// Identifies all appearances of this vulnerability in the package for a specific distro/location. For example: glibc in cpe:/o:debian:debian_linux:8 for versions 2.1 - 2.2
    /// </summary>
    public sealed class DetailArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The CPE URI in [cpe format](https://cpe.mitre.org/specification/) in which the vulnerability manifests. Examples include distro or storage location for vulnerable jar.
        /// </summary>
        [Input("cpeUri", required: true)]
        public Input<string> CpeUri { get; set; } = null!;

        /// <summary>
        /// A vendor-specific description of this note.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The fix for this specific package version.
        /// </summary>
        [Input("fixedLocation")]
        public Input<Inputs.VulnerabilityLocationArgs>? FixedLocation { get; set; }

        /// <summary>
        /// Whether this detail is obsolete. Occurrences are expected not to point to obsolete details.
        /// </summary>
        [Input("isObsolete")]
        public Input<bool>? IsObsolete { get; set; }

        /// <summary>
        /// The max version of the package in which the vulnerability exists.
        /// </summary>
        [Input("maxAffectedVersion")]
        public Input<Inputs.VersionArgs>? MaxAffectedVersion { get; set; }

        /// <summary>
        /// The min version of the package in which the vulnerability exists.
        /// </summary>
        [Input("minAffectedVersion")]
        public Input<Inputs.VersionArgs>? MinAffectedVersion { get; set; }

        /// <summary>
        /// The name of the package where the vulnerability was found.
        /// </summary>
        [Input("package", required: true)]
        public Input<string> Package { get; set; } = null!;

        /// <summary>
        /// The type of package; whether native or non native(ruby gems, node.js packages etc).
        /// </summary>
        [Input("packageType")]
        public Input<string>? PackageType { get; set; }

        /// <summary>
        /// The severity (eg: distro assigned severity) for this vulnerability.
        /// </summary>
        [Input("severityName")]
        public Input<string>? SeverityName { get; set; }

        /// <summary>
        /// The source from which the information in this Detail was obtained.
        /// </summary>
        [Input("source")]
        public Input<string>? Source { get; set; }

        /// <summary>
        /// The time this information was last changed at the source. This is an upstream timestamp from the underlying information source - e.g. Ubuntu security tracker.
        /// </summary>
        [Input("sourceUpdateTime")]
        public Input<string>? SourceUpdateTime { get; set; }

        /// <summary>
        /// The name of the vendor of the product.
        /// </summary>
        [Input("vendor")]
        public Input<string>? Vendor { get; set; }

        public DetailArgs()
        {
        }
        public static new DetailArgs Empty => new DetailArgs();
    }
}
