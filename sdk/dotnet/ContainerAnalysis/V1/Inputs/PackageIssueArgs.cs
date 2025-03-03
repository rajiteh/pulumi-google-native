// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Inputs
{

    /// <summary>
    /// A detail for a distro and package this vulnerability occurrence was found in and its associated fix (if one is available).
    /// </summary>
    public sealed class PackageIssueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The [CPE URI](https://cpe.mitre.org/specification/) this vulnerability was found in.
        /// </summary>
        [Input("affectedCpeUri", required: true)]
        public Input<string> AffectedCpeUri { get; set; } = null!;

        /// <summary>
        /// The package this vulnerability was found in.
        /// </summary>
        [Input("affectedPackage", required: true)]
        public Input<string> AffectedPackage { get; set; } = null!;

        /// <summary>
        /// The version of the package that is installed on the resource affected by this vulnerability.
        /// </summary>
        [Input("affectedVersion", required: true)]
        public Input<Inputs.VersionArgs> AffectedVersion { get; set; } = null!;

        [Input("fileLocation")]
        private InputList<Inputs.GrafeasV1FileLocationArgs>? _fileLocation;

        /// <summary>
        /// The location at which this package was found.
        /// </summary>
        public InputList<Inputs.GrafeasV1FileLocationArgs> FileLocation
        {
            get => _fileLocation ?? (_fileLocation = new InputList<Inputs.GrafeasV1FileLocationArgs>());
            set => _fileLocation = value;
        }

        /// <summary>
        /// The [CPE URI](https://cpe.mitre.org/specification/) this vulnerability was fixed in. It is possible for this to be different from the affected_cpe_uri.
        /// </summary>
        [Input("fixedCpeUri")]
        public Input<string>? FixedCpeUri { get; set; }

        /// <summary>
        /// The package this vulnerability was fixed in. It is possible for this to be different from the affected_package.
        /// </summary>
        [Input("fixedPackage")]
        public Input<string>? FixedPackage { get; set; }

        /// <summary>
        /// The version of the package this vulnerability was fixed in. Setting this to VersionKind.MAXIMUM means no fix is yet available.
        /// </summary>
        [Input("fixedVersion", required: true)]
        public Input<Inputs.VersionArgs> FixedVersion { get; set; } = null!;

        /// <summary>
        /// The type of package (e.g. OS, MAVEN, GO).
        /// </summary>
        [Input("packageType")]
        public Input<string>? PackageType { get; set; }

        public PackageIssueArgs()
        {
        }
        public static new PackageIssueArgs Empty => new PackageIssueArgs();
    }
}
