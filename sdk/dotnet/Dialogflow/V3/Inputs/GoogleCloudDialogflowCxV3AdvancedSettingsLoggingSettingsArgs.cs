// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3.Inputs
{

    /// <summary>
    /// Define behaviors on logging.
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3AdvancedSettingsLoggingSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, DF Interaction logging is currently enabled.
        /// </summary>
        [Input("enableInteractionLogging")]
        public Input<bool>? EnableInteractionLogging { get; set; }

        /// <summary>
        /// If true, StackDriver logging is currently enabled.
        /// </summary>
        [Input("enableStackdriverLogging")]
        public Input<bool>? EnableStackdriverLogging { get; set; }

        public GoogleCloudDialogflowCxV3AdvancedSettingsLoggingSettingsArgs()
        {
        }
        public static new GoogleCloudDialogflowCxV3AdvancedSettingsLoggingSettingsArgs Empty => new GoogleCloudDialogflowCxV3AdvancedSettingsLoggingSettingsArgs();
    }
}
