// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.PolicySimulator.V1Beta1.Inputs
{

    /// <summary>
    /// The configuration used for a Replay.
    /// </summary>
    public sealed class GoogleCloudPolicysimulatorV1beta1ReplayConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The logs to use as input for the Replay.
        /// </summary>
        [Input("logSource")]
        public Input<Pulumi.GoogleNative.PolicySimulator.V1Beta1.GoogleCloudPolicysimulatorV1beta1ReplayConfigLogSource>? LogSource { get; set; }

        [Input("policyOverlay")]
        private InputMap<string>? _policyOverlay;

        /// <summary>
        /// A mapping of the resources that you want to simulate policies for and the policies that you want to simulate. Keys are the full resource names for the resources. For example, `//cloudresourcemanager.googleapis.com/projects/my-project`. For examples of full resource names for Google Cloud services, see https://cloud.google.com/iam/help/troubleshooter/full-resource-names. Values are Policy objects representing the policies that you want to simulate. Replays automatically take into account any IAM policies inherited through the resource hierarchy, and any policies set on descendant resources. You do not need to include these policies in the policy overlay.
        /// </summary>
        public InputMap<string> PolicyOverlay
        {
            get => _policyOverlay ?? (_policyOverlay = new InputMap<string>());
            set => _policyOverlay = value;
        }

        public GoogleCloudPolicysimulatorV1beta1ReplayConfigArgs()
        {
        }
        public static new GoogleCloudPolicysimulatorV1beta1ReplayConfigArgs Empty => new GoogleCloudPolicysimulatorV1beta1ReplayConfigArgs();
    }
}
