// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Inputs
{

    public sealed class EnterpriseCrmEventbusProtoIntParameterArrayArgs : global::Pulumi.ResourceArgs
    {
        [Input("intValues")]
        private InputList<string>? _intValues;
        public InputList<string> IntValues
        {
            get => _intValues ?? (_intValues = new InputList<string>());
            set => _intValues = value;
        }

        public EnterpriseCrmEventbusProtoIntParameterArrayArgs()
        {
        }
        public static new EnterpriseCrmEventbusProtoIntParameterArrayArgs Empty => new EnterpriseCrmEventbusProtoIntParameterArrayArgs();
    }
}
