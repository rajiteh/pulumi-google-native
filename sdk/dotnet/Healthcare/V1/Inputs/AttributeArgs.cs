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
    /// An attribute value for a Consent or User data mapping. Each Attribute must have a corresponding AttributeDefinition in the consent store that defines the default and allowed values.
    /// </summary>
    public sealed class AttributeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates the name of an attribute defined in the consent store.
        /// </summary>
        [Input("attributeDefinitionId")]
        public Input<string>? AttributeDefinitionId { get; set; }

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// The value of the attribute. Must be an acceptable value as defined in the consent store. For example, if the consent store defines "data type" with acceptable values "questionnaire" and "step-count", when the attribute name is data type, this field must contain one of those values.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public AttributeArgs()
        {
        }
        public static new AttributeArgs Empty => new AttributeArgs();
    }
}
