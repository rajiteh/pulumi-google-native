// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Spanner.V1
{
    public static class GetDatabase
    {
        /// <summary>
        /// Gets the state of a Cloud Spanner database.
        /// </summary>
        public static Task<GetDatabaseResult> InvokeAsync(GetDatabaseArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseResult>("google-native:spanner/v1:getDatabase", args ?? new GetDatabaseArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the state of a Cloud Spanner database.
        /// </summary>
        public static Output<GetDatabaseResult> Invoke(GetDatabaseInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabaseResult>("google-native:spanner/v1:getDatabase", args ?? new GetDatabaseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseArgs : global::Pulumi.InvokeArgs
    {
        [Input("databaseId", required: true)]
        public string DatabaseId { get; set; } = null!;

        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetDatabaseArgs()
        {
        }
        public static new GetDatabaseArgs Empty => new GetDatabaseArgs();
    }

    public sealed class GetDatabaseInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("databaseId", required: true)]
        public Input<string> DatabaseId { get; set; } = null!;

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetDatabaseInvokeArgs()
        {
        }
        public static new GetDatabaseInvokeArgs Empty => new GetDatabaseInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabaseResult
    {
        /// <summary>
        /// If exists, the time at which the database creation started.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The dialect of the Cloud Spanner Database.
        /// </summary>
        public readonly string DatabaseDialect;
        /// <summary>
        /// The read-write region which contains the database's leader replicas. This is the same as the value of default_leader database option set using DatabaseAdmin.CreateDatabase or DatabaseAdmin.UpdateDatabaseDdl. If not explicitly set, this is empty.
        /// </summary>
        public readonly string DefaultLeader;
        /// <summary>
        /// Earliest timestamp at which older versions of the data can be read. This value is continuously updated by Cloud Spanner and becomes stale the moment it is queried. If you are using this value to recover data, make sure to account for the time from the moment when the value is queried to the moment when you initiate the recovery.
        /// </summary>
        public readonly string EarliestVersionTime;
        /// <summary>
        /// Whether drop protection is enabled for this database. Defaults to false, if not set.
        /// </summary>
        public readonly bool EnableDropProtection;
        /// <summary>
        /// For databases that are using customer managed encryption, this field contains the encryption configuration for the database. For databases that are using Google default or other types of encryption, this field is empty.
        /// </summary>
        public readonly Outputs.EncryptionConfigResponse EncryptionConfig;
        /// <summary>
        /// For databases that are using customer managed encryption, this field contains the encryption information for the database, such as all Cloud KMS key versions that are in use. The `encryption_status' field inside of each `EncryptionInfo` is not populated. For databases that are using Google default or other types of encryption, this field is empty. This field is propagated lazily from the backend. There might be a delay from when a key version is being used and when it appears in this field.
        /// </summary>
        public readonly ImmutableArray<Outputs.EncryptionInfoResponse> EncryptionInfo;
        /// <summary>
        /// The name of the database. Values are of the form `projects//instances//databases/`, where `` is as specified in the `CREATE DATABASE` statement. This name can be passed to other API methods to identify the database.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// If true, the database is being updated. If false, there are no ongoing update operations for the database.
        /// </summary>
        public readonly bool Reconciling;
        /// <summary>
        /// Applicable only for restored databases. Contains information about the restore source.
        /// </summary>
        public readonly Outputs.RestoreInfoResponse RestoreInfo;
        /// <summary>
        /// The current database state.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The period in which Cloud Spanner retains all versions of data for the database. This is the same as the value of version_retention_period database option set using UpdateDatabaseDdl. Defaults to 1 hour, if not set.
        /// </summary>
        public readonly string VersionRetentionPeriod;

        [OutputConstructor]
        private GetDatabaseResult(
            string createTime,

            string databaseDialect,

            string defaultLeader,

            string earliestVersionTime,

            bool enableDropProtection,

            Outputs.EncryptionConfigResponse encryptionConfig,

            ImmutableArray<Outputs.EncryptionInfoResponse> encryptionInfo,

            string name,

            bool reconciling,

            Outputs.RestoreInfoResponse restoreInfo,

            string state,

            string versionRetentionPeriod)
        {
            CreateTime = createTime;
            DatabaseDialect = databaseDialect;
            DefaultLeader = defaultLeader;
            EarliestVersionTime = earliestVersionTime;
            EnableDropProtection = enableDropProtection;
            EncryptionConfig = encryptionConfig;
            EncryptionInfo = encryptionInfo;
            Name = name;
            Reconciling = reconciling;
            RestoreInfo = restoreInfo;
            State = state;
            VersionRetentionPeriod = versionRetentionPeriod;
        }
    }
}
