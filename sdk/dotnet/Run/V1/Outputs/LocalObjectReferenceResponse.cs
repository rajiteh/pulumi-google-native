// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.Run.V1.Outputs
{

    [OutputType]
    public sealed class LocalObjectReferenceResponse
    {
        /// <summary>
        /// (Optional) Cloud Run fully managed: not supported Cloud Run for Anthos: supported Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private LocalObjectReferenceResponse(string name)
        {
            Name = name;
        }
    }
}