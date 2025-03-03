// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3.Outputs
{

    /// <summary>
    /// A transition route specifies a intent that can be matched and/or a data condition that can be evaluated during a session. When a specified transition is matched, the following actions are taken in order: * If there is a `trigger_fulfillment` associated with the transition, it will be called. * If there is a `target_page` associated with the transition, the session will transition into the specified page. * If there is a `target_flow` associated with the transition, the session will transition into the specified flow.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3TransitionRouteResponse
    {
        /// <summary>
        /// The condition to evaluate against form parameters or session parameters. See the [conditions reference](https://cloud.google.com/dialogflow/cx/docs/reference/condition). At least one of `intent` or `condition` must be specified. When both `intent` and `condition` are specified, the transition can only happen when both are fulfilled.
        /// </summary>
        public readonly string Condition;
        /// <summary>
        /// The unique identifier of an Intent. Format: `projects//locations//agents//intents/`. Indicates that the transition can only happen when the given intent is matched. At least one of `intent` or `condition` must be specified. When both `intent` and `condition` are specified, the transition can only happen when both are fulfilled.
        /// </summary>
        public readonly string Intent;
        /// <summary>
        /// The unique identifier of this transition route.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The target flow to transition to. Format: `projects//locations//agents//flows/`.
        /// </summary>
        public readonly string TargetFlow;
        /// <summary>
        /// The target page to transition to. Format: `projects//locations//agents//flows//pages/`.
        /// </summary>
        public readonly string TargetPage;
        /// <summary>
        /// The fulfillment to call when the condition is satisfied. At least one of `trigger_fulfillment` and `target` must be specified. When both are defined, `trigger_fulfillment` is executed first.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowCxV3FulfillmentResponse TriggerFulfillment;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3TransitionRouteResponse(
            string condition,

            string intent,

            string name,

            string targetFlow,

            string targetPage,

            Outputs.GoogleCloudDialogflowCxV3FulfillmentResponse triggerFulfillment)
        {
            Condition = condition;
            Intent = intent;
            Name = name;
            TargetFlow = targetFlow;
            TargetPage = targetPage;
            TriggerFulfillment = triggerFulfillment;
        }
    }
}
