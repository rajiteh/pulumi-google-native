// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BareMetalSolution.V2.Inputs
{

    /// <summary>
    /// A storage volume.
    /// </summary>
    public sealed class VolumeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The size, in GiB, that this storage volume has expanded as a result of an auto grow policy. In the absence of auto-grow, the value is 0.
        /// </summary>
        [Input("autoGrownSizeGib")]
        public Input<string>? AutoGrownSizeGib { get; set; }

        /// <summary>
        /// The current size of this storage volume, in GiB, including space reserved for snapshots. This size might be different than the requested size if the storage volume has been configured with auto grow or auto shrink.
        /// </summary>
        [Input("currentSizeGib")]
        public Input<string>? CurrentSizeGib { get; set; }

        /// <summary>
        /// Additional emergency size that was requested for this Volume, in GiB. current_size_gib includes this value.
        /// </summary>
        [Input("emergencySizeGib")]
        public Input<string>? EmergencySizeGib { get; set; }

        /// <summary>
        /// An identifier for the `Volume`, generated by the backend.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels as key value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Maximum size volume can be expanded to in case of evergency, in GiB.
        /// </summary>
        [Input("maxSizeGib")]
        public Input<string>? MaxSizeGib { get; set; }

        /// <summary>
        /// Originally requested size, in GiB.
        /// </summary>
        [Input("originallyRequestedSizeGib")]
        public Input<string>? OriginallyRequestedSizeGib { get; set; }

        /// <summary>
        /// Immutable. Pod name.
        /// </summary>
        [Input("pod")]
        public Input<string>? Pod { get; set; }

        /// <summary>
        /// The space remaining in the storage volume for new LUNs, in GiB, excluding space reserved for snapshots.
        /// </summary>
        [Input("remainingSpaceGib")]
        public Input<string>? RemainingSpaceGib { get; set; }

        /// <summary>
        /// The requested size of this storage volume, in GiB.
        /// </summary>
        [Input("requestedSizeGib")]
        public Input<string>? RequestedSizeGib { get; set; }

        /// <summary>
        /// The behavior to use when snapshot reserved space is full.
        /// </summary>
        [Input("snapshotAutoDeleteBehavior")]
        public Input<Pulumi.GoogleNative.BareMetalSolution.V2.VolumeSnapshotAutoDeleteBehavior>? SnapshotAutoDeleteBehavior { get; set; }

        /// <summary>
        /// Whether snapshots are enabled.
        /// </summary>
        [Input("snapshotEnabled")]
        public Input<bool>? SnapshotEnabled { get; set; }

        /// <summary>
        /// Details about snapshot space reservation and usage on the storage volume.
        /// </summary>
        [Input("snapshotReservationDetail")]
        public Input<Inputs.SnapshotReservationDetailArgs>? SnapshotReservationDetail { get; set; }

        /// <summary>
        /// The name of the snapshot schedule policy in use for this volume, if any.
        /// </summary>
        [Input("snapshotSchedulePolicy")]
        public Input<string>? SnapshotSchedulePolicy { get; set; }

        /// <summary>
        /// The state of this storage volume.
        /// </summary>
        [Input("state")]
        public Input<Pulumi.GoogleNative.BareMetalSolution.V2.VolumeState>? State { get; set; }

        /// <summary>
        /// The storage type for this volume.
        /// </summary>
        [Input("storageType")]
        public Input<Pulumi.GoogleNative.BareMetalSolution.V2.VolumeStorageType>? StorageType { get; set; }

        public VolumeArgs()
        {
        }
        public static new VolumeArgs Empty => new VolumeArgs();
    }
}