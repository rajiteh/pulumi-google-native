// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.TPU.V2Alpha1.Outputs
{

    /// <summary>
    /// QueuedResourceState defines the details of the QueuedResource request.
    /// </summary>
    [OutputType]
    public sealed class QueuedResourceStateResponse
    {
        /// <summary>
        /// Further data for the accepted state.
        /// </summary>
        public readonly Outputs.AcceptedDataResponse AcceptedData;
        /// <summary>
        /// Further data for the active state.
        /// </summary>
        public readonly Outputs.ActiveDataResponse ActiveData;
        /// <summary>
        /// Further data for the creating state.
        /// </summary>
        public readonly Outputs.CreatingDataResponse CreatingData;
        /// <summary>
        /// Further data for the deleting state.
        /// </summary>
        public readonly Outputs.DeletingDataResponse DeletingData;
        /// <summary>
        /// Further data for the failed state.
        /// </summary>
        public readonly Outputs.FailedDataResponse FailedData;
        /// <summary>
        /// Further data for the provisioning state.
        /// </summary>
        public readonly Outputs.ProvisioningDataResponse ProvisioningData;
        /// <summary>
        /// State of the QueuedResource request.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Further data for the suspended state.
        /// </summary>
        public readonly Outputs.SuspendedDataResponse SuspendedData;
        /// <summary>
        /// Further data for the suspending state.
        /// </summary>
        public readonly Outputs.SuspendingDataResponse SuspendingData;

        [OutputConstructor]
        private QueuedResourceStateResponse(
            Outputs.AcceptedDataResponse acceptedData,

            Outputs.ActiveDataResponse activeData,

            Outputs.CreatingDataResponse creatingData,

            Outputs.DeletingDataResponse deletingData,

            Outputs.FailedDataResponse failedData,

            Outputs.ProvisioningDataResponse provisioningData,

            string state,

            Outputs.SuspendedDataResponse suspendedData,

            Outputs.SuspendingDataResponse suspendingData)
        {
            AcceptedData = acceptedData;
            ActiveData = activeData;
            CreatingData = creatingData;
            DeletingData = deletingData;
            FailedData = failedData;
            ProvisioningData = provisioningData;
            State = state;
            SuspendedData = suspendedData;
            SuspendingData = suspendingData;
        }
    }
}
