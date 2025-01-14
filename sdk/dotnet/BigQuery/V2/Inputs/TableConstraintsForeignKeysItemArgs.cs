// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2.Inputs
{

    public sealed class TableConstraintsForeignKeysItemArgs : global::Pulumi.ResourceArgs
    {
        [Input("columnReferences")]
        private InputList<Inputs.TableConstraintsForeignKeysItemColumnReferencesItemArgs>? _columnReferences;
        public InputList<Inputs.TableConstraintsForeignKeysItemColumnReferencesItemArgs> ColumnReferences
        {
            get => _columnReferences ?? (_columnReferences = new InputList<Inputs.TableConstraintsForeignKeysItemColumnReferencesItemArgs>());
            set => _columnReferences = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("referencedTable")]
        public Input<Inputs.TableConstraintsForeignKeysItemReferencedTableArgs>? ReferencedTable { get; set; }

        public TableConstraintsForeignKeysItemArgs()
        {
        }
        public static new TableConstraintsForeignKeysItemArgs Empty => new TableConstraintsForeignKeysItemArgs();
    }
}
