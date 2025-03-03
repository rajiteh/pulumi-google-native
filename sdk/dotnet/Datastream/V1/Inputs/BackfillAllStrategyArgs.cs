// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1.Inputs
{

    /// <summary>
    /// Backfill strategy to automatically backfill the Stream's objects. Specific objects can be excluded.
    /// </summary>
    public sealed class BackfillAllStrategyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// MySQL data source objects to avoid backfilling.
        /// </summary>
        [Input("mysqlExcludedObjects")]
        public Input<Inputs.MysqlRdbmsArgs>? MysqlExcludedObjects { get; set; }

        /// <summary>
        /// Oracle data source objects to avoid backfilling.
        /// </summary>
        [Input("oracleExcludedObjects")]
        public Input<Inputs.OracleRdbmsArgs>? OracleExcludedObjects { get; set; }

        /// <summary>
        /// PostgreSQL data source objects to avoid backfilling.
        /// </summary>
        [Input("postgresqlExcludedObjects")]
        public Input<Inputs.PostgresqlRdbmsArgs>? PostgresqlExcludedObjects { get; set; }

        public BackfillAllStrategyArgs()
        {
        }
        public static new BackfillAllStrategyArgs Empty => new BackfillAllStrategyArgs();
    }
}
