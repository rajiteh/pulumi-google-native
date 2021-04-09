// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.Eventarc.V1.Outputs
{

    [OutputType]
    public sealed class EventFilterResponse
    {
        /// <summary>
        /// Required. The name of a CloudEvents attribute. Currently, only a subset of attributes are supported for filtering. All triggers MUST provide a filter for the 'type' attribute.
        /// </summary>
        public readonly string Attribute;
        /// <summary>
        /// Required. The value for the attribute.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private EventFilterResponse(
            string attribute,

            string value)
        {
            Attribute = attribute;
            Value = value;
        }
    }
}