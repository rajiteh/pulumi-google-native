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
    public sealed class EnterpriseCrmFrontendsEventbusProtoDoubleParameterArrayResponse
    {
        public readonly ImmutableArray<double> DoubleValues;

        [OutputConstructor]
        private EnterpriseCrmFrontendsEventbusProtoDoubleParameterArrayResponse(ImmutableArray<double> doubleValues)
        {
            DoubleValues = doubleValues;
        }
    }
}
