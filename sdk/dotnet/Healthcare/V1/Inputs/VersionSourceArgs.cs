// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1.Inputs
{

    /// <summary>
    /// Describes a selector for extracting and matching an MSH field to a value.
    /// </summary>
    public sealed class VersionSourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The field to extract from the MSH segment. For example, "3.1" or "18[1].1".
        /// </summary>
        [Input("mshField")]
        public Input<string>? MshField { get; set; }

        /// <summary>
        /// The value to match with the field. For example, "My Application Name" or "2.3".
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public VersionSourceArgs()
        {
        }
        public static new VersionSourceArgs Empty => new VersionSourceArgs();
    }
}
