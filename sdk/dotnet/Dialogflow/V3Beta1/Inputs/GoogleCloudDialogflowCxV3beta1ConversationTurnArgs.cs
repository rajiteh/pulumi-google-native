// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3Beta1.Inputs
{

    /// <summary>
    /// One interaction between a human and virtual agent. The human provides some input and the virtual agent provides a response.
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3beta1ConversationTurnArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The user input.
        /// </summary>
        [Input("userInput")]
        public Input<Inputs.GoogleCloudDialogflowCxV3beta1ConversationTurnUserInputArgs>? UserInput { get; set; }

        /// <summary>
        /// The virtual agent output.
        /// </summary>
        [Input("virtualAgentOutput")]
        public Input<Inputs.GoogleCloudDialogflowCxV3beta1ConversationTurnVirtualAgentOutputArgs>? VirtualAgentOutput { get; set; }

        public GoogleCloudDialogflowCxV3beta1ConversationTurnArgs()
        {
        }
        public static new GoogleCloudDialogflowCxV3beta1ConversationTurnArgs Empty => new GoogleCloudDialogflowCxV3beta1ConversationTurnArgs();
    }
}
