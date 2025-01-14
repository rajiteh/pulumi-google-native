// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Inputs
{

    /// <summary>
    /// The LogSettings define the logging attributes for an event property. These attributes are used to map the property to the parameter in the log proto. Also used to define scrubbing/truncation behavior and PII information.
    /// </summary>
    public sealed class EnterpriseCrmEventbusProtoLogSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of corresponding logging field of the event property. If omitted, assumes the same name as the event property key.
        /// </summary>
        [Input("logFieldName")]
        public Input<string>? LogFieldName { get; set; }

        /// <summary>
        /// Contains the scrubbing options, such as whether to scrub, obfuscate, etc.
        /// </summary>
        [Input("sanitizeOptions")]
        public Input<Inputs.EnterpriseCrmLoggingGwsSanitizeOptionsArgs>? SanitizeOptions { get; set; }

        [Input("seedPeriod")]
        public Input<Pulumi.GoogleNative.Integrations.V1Alpha.EnterpriseCrmEventbusProtoLogSettingsSeedPeriod>? SeedPeriod { get; set; }

        [Input("seedScope")]
        public Input<Pulumi.GoogleNative.Integrations.V1Alpha.EnterpriseCrmEventbusProtoLogSettingsSeedScope>? SeedScope { get; set; }

        /// <summary>
        /// Contains the field limits for shortening, such as max string length and max array length.
        /// </summary>
        [Input("shorteningLimits")]
        public Input<Inputs.EnterpriseCrmLoggingGwsFieldLimitsArgs>? ShorteningLimits { get; set; }

        public EnterpriseCrmEventbusProtoLogSettingsArgs()
        {
        }
        public static new EnterpriseCrmEventbusProtoLogSettingsArgs Empty => new EnterpriseCrmEventbusProtoLogSettingsArgs();
    }
}
