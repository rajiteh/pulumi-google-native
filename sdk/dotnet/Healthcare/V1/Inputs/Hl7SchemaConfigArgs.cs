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
    /// Root config message for HL7v2 schema. This contains a schema structure of groups and segments, and filters that determine which messages to apply the schema structure to.
    /// </summary>
    public sealed class Hl7SchemaConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("messageSchemaConfigs")]
        private InputMap<string>? _messageSchemaConfigs;

        /// <summary>
        /// Map from each HL7v2 message type and trigger event pair, such as ADT_A04, to its schema configuration root group.
        /// </summary>
        public InputMap<string> MessageSchemaConfigs
        {
            get => _messageSchemaConfigs ?? (_messageSchemaConfigs = new InputMap<string>());
            set => _messageSchemaConfigs = value;
        }

        [Input("version")]
        private InputList<Inputs.VersionSourceArgs>? _version;

        /// <summary>
        /// Each VersionSource is tested and only if they all match is the schema used for the message.
        /// </summary>
        public InputList<Inputs.VersionSourceArgs> Version
        {
            get => _version ?? (_version = new InputList<Inputs.VersionSourceArgs>());
            set => _version = value;
        }

        public Hl7SchemaConfigArgs()
        {
        }
        public static new Hl7SchemaConfigArgs Empty => new Hl7SchemaConfigArgs();
    }
}
