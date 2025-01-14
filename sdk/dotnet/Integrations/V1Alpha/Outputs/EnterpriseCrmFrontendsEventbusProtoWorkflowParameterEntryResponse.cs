// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Outputs
{

    [OutputType]
    public sealed class EnterpriseCrmFrontendsEventbusProtoWorkflowParameterEntryResponse
    {
        /// <summary>
        /// Metadata information about the parameters.
        /// </summary>
        public readonly Outputs.EnterpriseCrmEventbusProtoAttributesResponse Attributes;
        /// <summary>
        /// Child parameters nested within this parameter. This field only applies to protobuf parameters
        /// </summary>
        public readonly ImmutableArray<Outputs.EnterpriseCrmFrontendsEventbusProtoWorkflowParameterEntryResponse> Children;
        /// <summary>
        /// The data type of the parameter.
        /// </summary>
        public readonly string DataType;
        /// <summary>
        /// Default values for the defined keys. Each value can either be string, int, double or any proto message or a serialized object.
        /// </summary>
        public readonly Outputs.EnterpriseCrmFrontendsEventbusProtoParameterValueTypeResponse DefaultValue;
        /// <summary>
        /// Specifies the input/output type for the parameter.
        /// </summary>
        public readonly string InOutType;
        /// <summary>
        /// Whether this parameter is a transient parameter.
        /// </summary>
        public readonly bool IsTransient;
        /// <summary>
        /// This schema will be used to validate runtime JSON-typed values of this parameter.
        /// </summary>
        public readonly string JsonSchema;
        /// <summary>
        /// Key is used to retrieve the corresponding parameter value. This should be unique for a given fired event. These parameters must be predefined in the workflow definition.
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// The name (without prefix) to be displayed in the UI for this parameter. E.g. if the key is "foo.bar.myName", then the name would be "myName".
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The identifier of the node (TaskConfig/TriggerConfig) this parameter was produced by, if it is a transient param or a copy of an input param.
        /// </summary>
        public readonly Outputs.EnterpriseCrmEventbusProtoNodeIdentifierResponse ProducedBy;
        public readonly string Producer;
        /// <summary>
        /// The name of the protobuf type if the parameter has a protobuf data type.
        /// </summary>
        public readonly string ProtoDefName;
        /// <summary>
        /// If the data type is of type proto or proto array, this field needs to be populated with the fully qualified proto name. This message, for example, would be "enterprise.crm.frontends.eventbus.proto.WorkflowParameterEntry".
        /// </summary>
        public readonly string ProtoDefPath;

        [OutputConstructor]
        private EnterpriseCrmFrontendsEventbusProtoWorkflowParameterEntryResponse(
            Outputs.EnterpriseCrmEventbusProtoAttributesResponse attributes,

            ImmutableArray<Outputs.EnterpriseCrmFrontendsEventbusProtoWorkflowParameterEntryResponse> children,

            string dataType,

            Outputs.EnterpriseCrmFrontendsEventbusProtoParameterValueTypeResponse defaultValue,

            string inOutType,

            bool isTransient,

            string jsonSchema,

            string key,

            string name,

            Outputs.EnterpriseCrmEventbusProtoNodeIdentifierResponse producedBy,

            string producer,

            string protoDefName,

            string protoDefPath)
        {
            Attributes = attributes;
            Children = children;
            DataType = dataType;
            DefaultValue = defaultValue;
            InOutType = inOutType;
            IsTransient = isTransient;
            JsonSchema = jsonSchema;
            Key = key;
            Name = name;
            ProducedBy = producedBy;
            Producer = producer;
            ProtoDefName = protoDefName;
            ProtoDefPath = protoDefPath;
        }
    }
}
