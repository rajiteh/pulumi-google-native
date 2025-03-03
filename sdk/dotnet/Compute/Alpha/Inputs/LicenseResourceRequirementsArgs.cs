// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    public sealed class LicenseResourceRequirementsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Minimum number of guest cpus required to use the Instance. Enforced at Instance creation and Instance start.
        /// </summary>
        [Input("minGuestCpuCount")]
        public Input<int>? MinGuestCpuCount { get; set; }

        /// <summary>
        /// Minimum memory required to use the Instance. Enforced at Instance creation and Instance start.
        /// </summary>
        [Input("minMemoryMb")]
        public Input<int>? MinMemoryMb { get; set; }

        public LicenseResourceRequirementsArgs()
        {
        }
        public static new LicenseResourceRequirementsArgs Empty => new LicenseResourceRequirementsArgs();
    }
}
