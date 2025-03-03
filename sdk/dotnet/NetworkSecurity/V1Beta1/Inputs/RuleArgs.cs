// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkSecurity.V1Beta1.Inputs
{

    /// <summary>
    /// Specification of rules.
    /// </summary>
    public sealed class RuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinations")]
        private InputList<Inputs.DestinationArgs>? _destinations;

        /// <summary>
        /// Optional. List of attributes for the traffic destination. All of the destinations must match. A destination is a match if a request matches all the specified hosts, ports, methods and headers. If not set, the action specified in the 'action' field will be applied without any rule checks for the destination.
        /// </summary>
        public InputList<Inputs.DestinationArgs> Destinations
        {
            get => _destinations ?? (_destinations = new InputList<Inputs.DestinationArgs>());
            set => _destinations = value;
        }

        [Input("sources")]
        private InputList<Inputs.SourceArgs>? _sources;

        /// <summary>
        /// Optional. List of attributes for the traffic source. All of the sources must match. A source is a match if both principals and ip_blocks match. If not set, the action specified in the 'action' field will be applied without any rule checks for the source.
        /// </summary>
        public InputList<Inputs.SourceArgs> Sources
        {
            get => _sources ?? (_sources = new InputList<Inputs.SourceArgs>());
            set => _sources = value;
        }

        public RuleArgs()
        {
        }
        public static new RuleArgs Empty => new RuleArgs();
    }
}
