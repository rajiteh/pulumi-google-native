// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1.Outputs
{

    /// <summary>
    /// Container message for hash values.
    /// </summary>
    [OutputType]
    public sealed class HashResponse
    {
        /// <summary>
        /// The type of hash that was performed.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The hash value.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private HashResponse(
            string type,

            string value)
        {
            Type = type;
            Value = value;
        }
    }
}
