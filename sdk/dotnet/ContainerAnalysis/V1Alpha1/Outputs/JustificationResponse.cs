// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Outputs
{

    /// <summary>
    /// Justification provides the justification when the state of the assessment if NOT_AFFECTED.
    /// </summary>
    [OutputType]
    public sealed class JustificationResponse
    {
        /// <summary>
        /// Additional details on why this justification was chosen.
        /// </summary>
        public readonly string Details;
        /// <summary>
        /// The justification type for this vulnerability.
        /// </summary>
        public readonly string JustificationType;

        [OutputConstructor]
        private JustificationResponse(
            string details,

            string justificationType)
        {
            Details = details;
            JustificationType = justificationType;
        }
    }
}