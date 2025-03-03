// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Inputs
{

    /// <summary>
    /// Configuration for the legacy Attribute Based Access Control authorization mode.
    /// </summary>
    public sealed class LegacyAbacArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the ABAC authorizer is enabled for this cluster. When enabled, identities in the system, including service accounts, nodes, and controllers, will have statically granted permissions beyond those provided by the RBAC configuration or IAM.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public LegacyAbacArgs()
        {
        }
        public static new LegacyAbacArgs Empty => new LegacyAbacArgs();
    }
}
