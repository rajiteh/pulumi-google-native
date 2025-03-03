// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Composer.V1Beta1.Inputs
{

    /// <summary>
    /// The configuration of Cloud SQL instance that is used by the Apache Airflow software.
    /// </summary>
    public sealed class DatabaseConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Cloud SQL machine type used by Airflow database. It has to be one of: db-n1-standard-2, db-n1-standard-4, db-n1-standard-8 or db-n1-standard-16. If not specified, db-n1-standard-2 will be used. Supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.
        /// </summary>
        [Input("machineType")]
        public Input<string>? MachineType { get; set; }

        public DatabaseConfigArgs()
        {
        }
        public static new DatabaseConfigArgs Empty => new DatabaseConfigArgs();
    }
}
