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
    /// Dataproc Node Group. The Dataproc NodeGroup resource is not related to the Dataproc NodeGroupAffinity resource.
    /// </summary>
    public sealed class NodeGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. Node group labels. Label keys must consist of from 1 to 63 characters and conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values can be empty. If specified, they must consist of from 1 to 63 characters and conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). The node group must have no more than 32 labelsn.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The Node group resource name (https://aip.dev/122).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Optional. The node group instance group configuration.
        /// </summary>
        [Input("nodeGroupConfig")]
        public Input<Inputs.InstanceGroupConfigArgs>? NodeGroupConfig { get; set; }

        [Input("roles", required: true)]
        private InputList<Pulumi.GoogleNative.Dataproc.V1.NodeGroupRolesItem>? _roles;

        /// <summary>
        /// Node group roles.
        /// </summary>
        public InputList<Pulumi.GoogleNative.Dataproc.V1.NodeGroupRolesItem> Roles
        {
            get => _roles ?? (_roles = new InputList<Pulumi.GoogleNative.Dataproc.V1.NodeGroupRolesItem>());
            set => _roles = value;
        }

        public NodeGroupArgs()
        {
        }
        public static new NodeGroupArgs Empty => new NodeGroupArgs();
    }
}