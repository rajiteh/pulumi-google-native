// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.IAM.V1.Inputs
{

    /// <summary>
    /// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
    /// </summary>
    public sealed class ConditionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of the expression. This is a longer text which describes the expression, e.g., when hovering over it in a UI.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Textual representation of an expression in Common Expression Language syntax.
        /// </summary>
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        /// <summary>
        /// A title for the expression, i.e. a short string describing its purpose.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public ConditionArgs()
        {
        }
        public static new ConditionArgs Empty => new ConditionArgs();
    }
}