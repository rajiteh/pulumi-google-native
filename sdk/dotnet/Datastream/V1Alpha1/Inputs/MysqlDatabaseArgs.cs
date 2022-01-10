// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1Alpha1.Inputs
{

    /// <summary>
    /// MySQL database.
    /// </summary>
    public sealed class MysqlDatabaseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Database name.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        [Input("mysqlTables")]
        private InputList<Inputs.MysqlTableArgs>? _mysqlTables;

        /// <summary>
        /// Tables in the database.
        /// </summary>
        public InputList<Inputs.MysqlTableArgs> MysqlTables
        {
            get => _mysqlTables ?? (_mysqlTables = new InputList<Inputs.MysqlTableArgs>());
            set => _mysqlTables = value;
        }

        public MysqlDatabaseArgs()
        {
        }
    }
}