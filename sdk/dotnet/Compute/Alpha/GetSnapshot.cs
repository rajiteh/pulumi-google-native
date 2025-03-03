// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    public static class GetSnapshot
    {
        /// <summary>
        /// Returns the specified Snapshot resource.
        /// </summary>
        public static Task<GetSnapshotResult> InvokeAsync(GetSnapshotArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotResult>("google-native:compute/alpha:getSnapshot", args ?? new GetSnapshotArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the specified Snapshot resource.
        /// </summary>
        public static Output<GetSnapshotResult> Invoke(GetSnapshotInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSnapshotResult>("google-native:compute/alpha:getSnapshot", args ?? new GetSnapshotInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSnapshotArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("snapshot", required: true)]
        public string Snapshot { get; set; } = null!;

        public GetSnapshotArgs()
        {
        }
        public static new GetSnapshotArgs Empty => new GetSnapshotArgs();
    }

    public sealed class GetSnapshotInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("snapshot", required: true)]
        public Input<string> Snapshot { get; set; } = null!;

        public GetSnapshotInvokeArgs()
        {
        }
        public static new GetSnapshotInvokeArgs Empty => new GetSnapshotInvokeArgs();
    }


    [OutputType]
    public sealed class GetSnapshotResult
    {
        /// <summary>
        /// The architecture of the snapshot. Valid values are ARM64 or X86_64.
        /// </summary>
        public readonly string Architecture;
        /// <summary>
        /// Set to true if snapshots are automatically created by applying resource policy on the target disk.
        /// </summary>
        public readonly bool AutoCreated;
        /// <summary>
        /// Creates the new snapshot in the snapshot chain labeled with the specified name. The chain name must be 1-63 characters long and comply with RFC1035. This is an uncommon option only for advanced service owners who needs to create separate snapshot chains, for example, for chargeback tracking. When you describe your snapshot resource, this field is visible only if it has a non-empty value.
        /// </summary>
        public readonly string ChainName;
        /// <summary>
        /// Size in bytes of the snapshot at creation time.
        /// </summary>
        public readonly string CreationSizeBytes;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Size of the source disk, specified in GB.
        /// </summary>
        public readonly string DiskSizeGb;
        /// <summary>
        /// Number of bytes downloaded to restore a snapshot to a disk.
        /// </summary>
        public readonly string DownloadBytes;
        /// <summary>
        /// Whether this snapshot is created from a confidential compute mode disk. see go/confidential-mode-in-arcus for details. [Output Only]: This field is not set by user, but from source disk.
        /// </summary>
        public readonly bool EnableConfidentialCompute;
        /// <summary>
        /// [Input Only] Whether to attempt an application consistent snapshot by informing the OS to prepare for the snapshot process.
        /// </summary>
        public readonly bool GuestFlush;
        /// <summary>
        /// A list of features to enable on the guest operating system. Applicable only for bootable images. Read Enabling guest operating system features to see a list of available options.
        /// </summary>
        public readonly ImmutableArray<Outputs.GuestOsFeatureResponse> GuestOsFeatures;
        /// <summary>
        /// Type of the resource. Always compute#snapshot for Snapshot resources.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// A fingerprint for the labels being applied to this snapshot, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a snapshot.
        /// </summary>
        public readonly string LabelFingerprint;
        /// <summary>
        /// Labels to apply to this snapshot. These can be later modified by the setLabels method. Label values may be empty.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Integer license codes indicating which licenses are attached to this snapshot.
        /// </summary>
        public readonly ImmutableArray<string> LicenseCodes;
        /// <summary>
        /// A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses attached (such as a Windows image).
        /// </summary>
        public readonly ImmutableArray<string> Licenses;
        /// <summary>
        /// An opaque location hint used to place the snapshot close to other resources. This field is for use by internal tools that use the public API.
        /// </summary>
        public readonly string LocationHint;
        /// <summary>
        /// Number of days the snapshot should be retained before being deleted automatically.
        /// </summary>
        public readonly int MaxRetentionDays;
        /// <summary>
        /// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Reserved for future use.
        /// </summary>
        public readonly bool SatisfiesPzs;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Server-defined URL for this resource's resource id.
        /// </summary>
        public readonly string SelfLinkWithId;
        /// <summary>
        /// Encrypts the snapshot using a customer-supplied encryption key. After you encrypt a snapshot using a customer-supplied key, you must provide the same key if you use the snapshot later. For example, you must provide the encryption key when you create a disk from the encrypted snapshot in a future request. Customer-supplied encryption keys do not protect access to metadata of the snapshot. If you do not provide an encryption key when creating the snapshot, then the snapshot will be encrypted using an automatically generated key and you do not need to provide a key to use the snapshot later.
        /// </summary>
        public readonly Outputs.CustomerEncryptionKeyResponse SnapshotEncryptionKey;
        /// <summary>
        /// Indicates the type of the snapshot.
        /// </summary>
        public readonly string SnapshotType;
        /// <summary>
        /// The source disk used to create this snapshot.
        /// </summary>
        public readonly string SourceDisk;
        /// <summary>
        /// The customer-supplied encryption key of the source disk. Required if the source disk is protected by a customer-supplied encryption key.
        /// </summary>
        public readonly Outputs.CustomerEncryptionKeyResponse SourceDiskEncryptionKey;
        /// <summary>
        /// The source disk whose recovery checkpoint will be used to create this snapshot.
        /// </summary>
        public readonly string SourceDiskForRecoveryCheckpoint;
        /// <summary>
        /// The ID value of the disk used to create this snapshot. This value may be used to determine whether the snapshot was taken from the current or a previous instance of a given disk name.
        /// </summary>
        public readonly string SourceDiskId;
        /// <summary>
        /// The source instant snapshot used to create this snapshot. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instantSnapshots/instantSnapshot - projects/project/zones/zone/instantSnapshots/instantSnapshot - zones/zone/instantSnapshots/instantSnapshot 
        /// </summary>
        public readonly string SourceInstantSnapshot;
        /// <summary>
        /// The unique ID of the instant snapshot used to create this snapshot. This value identifies the exact instant snapshot that was used to create this persistent disk. For example, if you created the persistent disk from an instant snapshot that was later deleted and recreated under the same name, the source instant snapshot ID would identify the exact instant snapshot that was used.
        /// </summary>
        public readonly string SourceInstantSnapshotId;
        /// <summary>
        /// URL of the resource policy which created this scheduled snapshot.
        /// </summary>
        public readonly string SourceSnapshotSchedulePolicy;
        /// <summary>
        /// ID of the resource policy which created this scheduled snapshot.
        /// </summary>
        public readonly string SourceSnapshotSchedulePolicyId;
        /// <summary>
        /// The status of the snapshot. This can be CREATING, DELETING, FAILED, READY, or UPLOADING.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A size of the storage used by the snapshot. As snapshots share storage, this number is expected to change with snapshot creation/deletion.
        /// </summary>
        public readonly string StorageBytes;
        /// <summary>
        /// An indicator whether storageBytes is in a stable state or it is being adjusted as a result of shared storage reallocation. This status can either be UPDATING, meaning the size of the snapshot is being updated, or UP_TO_DATE, meaning the size of the snapshot is up-to-date.
        /// </summary>
        public readonly string StorageBytesStatus;
        /// <summary>
        /// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
        /// </summary>
        public readonly ImmutableArray<string> StorageLocations;
        /// <summary>
        /// A list of user provided licenses represented by a list of URLs to the license resource.
        /// </summary>
        public readonly ImmutableArray<string> UserLicenses;

        [OutputConstructor]
        private GetSnapshotResult(
            string architecture,

            bool autoCreated,

            string chainName,

            string creationSizeBytes,

            string creationTimestamp,

            string description,

            string diskSizeGb,

            string downloadBytes,

            bool enableConfidentialCompute,

            bool guestFlush,

            ImmutableArray<Outputs.GuestOsFeatureResponse> guestOsFeatures,

            string kind,

            string labelFingerprint,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<string> licenseCodes,

            ImmutableArray<string> licenses,

            string locationHint,

            int maxRetentionDays,

            string name,

            bool satisfiesPzs,

            string selfLink,

            string selfLinkWithId,

            Outputs.CustomerEncryptionKeyResponse snapshotEncryptionKey,

            string snapshotType,

            string sourceDisk,

            Outputs.CustomerEncryptionKeyResponse sourceDiskEncryptionKey,

            string sourceDiskForRecoveryCheckpoint,

            string sourceDiskId,

            string sourceInstantSnapshot,

            string sourceInstantSnapshotId,

            string sourceSnapshotSchedulePolicy,

            string sourceSnapshotSchedulePolicyId,

            string status,

            string storageBytes,

            string storageBytesStatus,

            ImmutableArray<string> storageLocations,

            ImmutableArray<string> userLicenses)
        {
            Architecture = architecture;
            AutoCreated = autoCreated;
            ChainName = chainName;
            CreationSizeBytes = creationSizeBytes;
            CreationTimestamp = creationTimestamp;
            Description = description;
            DiskSizeGb = diskSizeGb;
            DownloadBytes = downloadBytes;
            EnableConfidentialCompute = enableConfidentialCompute;
            GuestFlush = guestFlush;
            GuestOsFeatures = guestOsFeatures;
            Kind = kind;
            LabelFingerprint = labelFingerprint;
            Labels = labels;
            LicenseCodes = licenseCodes;
            Licenses = licenses;
            LocationHint = locationHint;
            MaxRetentionDays = maxRetentionDays;
            Name = name;
            SatisfiesPzs = satisfiesPzs;
            SelfLink = selfLink;
            SelfLinkWithId = selfLinkWithId;
            SnapshotEncryptionKey = snapshotEncryptionKey;
            SnapshotType = snapshotType;
            SourceDisk = sourceDisk;
            SourceDiskEncryptionKey = sourceDiskEncryptionKey;
            SourceDiskForRecoveryCheckpoint = sourceDiskForRecoveryCheckpoint;
            SourceDiskId = sourceDiskId;
            SourceInstantSnapshot = sourceInstantSnapshot;
            SourceInstantSnapshotId = sourceInstantSnapshotId;
            SourceSnapshotSchedulePolicy = sourceSnapshotSchedulePolicy;
            SourceSnapshotSchedulePolicyId = sourceSnapshotSchedulePolicyId;
            Status = status;
            StorageBytes = storageBytes;
            StorageBytesStatus = storageBytesStatus;
            StorageLocations = storageLocations;
            UserLicenses = userLicenses;
        }
    }
}
