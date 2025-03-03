// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Inputs
{

    /// <summary>
    /// Node group identification and configuration information.
    /// </summary>
    public sealed class AuxiliaryNodeGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Node group configuration.
        /// </summary>
        [Input("nodeGroup", required: true)]
        public Input<Inputs.NodeGroupArgs> NodeGroup { get; set; } = null!;

        /// <summary>
        /// Optional. A node group ID. Generated if not specified.The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of from 3 to 33 characters.
        /// </summary>
        [Input("nodeGroupId")]
        public Input<string>? NodeGroupId { get; set; }

        public AuxiliaryNodeGroupArgs()
        {
        }
        public static new AuxiliaryNodeGroupArgs Empty => new AuxiliaryNodeGroupArgs();
    }
}
