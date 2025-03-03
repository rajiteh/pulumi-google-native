// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1Beta4.Outputs
{

    /// <summary>
    /// Database instance backup configuration.
    /// </summary>
    [OutputType]
    public sealed class BackupConfigurationResponse
    {
        /// <summary>
        /// Backup retention settings.
        /// </summary>
        public readonly Outputs.BackupRetentionSettingsResponse BackupRetentionSettings;
        /// <summary>
        /// (MySQL only) Whether binary log is enabled. If backup configuration is disabled, binarylog must be disabled as well.
        /// </summary>
        public readonly bool BinaryLogEnabled;
        /// <summary>
        /// Whether this configuration is enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// This is always `sql#backupConfiguration`.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// Location of the backup
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// (Postgres only) Whether point in time recovery is enabled.
        /// </summary>
        public readonly bool PointInTimeRecoveryEnabled;
        /// <summary>
        /// Reserved for future use.
        /// </summary>
        public readonly bool ReplicationLogArchivingEnabled;
        /// <summary>
        /// Start time for the daily backup configuration in UTC timezone in the 24 hour format - `HH:MM`.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// The number of days of transaction logs we retain for point in time restore, from 1-7.
        /// </summary>
        public readonly int TransactionLogRetentionDays;

        [OutputConstructor]
        private BackupConfigurationResponse(
            Outputs.BackupRetentionSettingsResponse backupRetentionSettings,

            bool binaryLogEnabled,

            bool enabled,

            string kind,

            string location,

            bool pointInTimeRecoveryEnabled,

            bool replicationLogArchivingEnabled,

            string startTime,

            int transactionLogRetentionDays)
        {
            BackupRetentionSettings = backupRetentionSettings;
            BinaryLogEnabled = binaryLogEnabled;
            Enabled = enabled;
            Kind = kind;
            Location = location;
            PointInTimeRecoveryEnabled = pointInTimeRecoveryEnabled;
            ReplicationLogArchivingEnabled = replicationLogArchivingEnabled;
            StartTime = startTime;
            TransactionLogRetentionDays = transactionLogRetentionDays;
        }
    }
}
