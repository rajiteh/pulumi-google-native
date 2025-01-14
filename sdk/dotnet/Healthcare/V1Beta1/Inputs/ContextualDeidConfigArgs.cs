// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1.Inputs
{

    /// <summary>
    /// The fields that aren't marked `Keep` or `CleanText` in the `BASIC` profile are collected into a contextual phrase list. For fields marked `CleanText`, the process attempts to transform phrases matching these contextual entries. These contextual phrases are replaced with the token "[CTX]". This feature uses an additional InfoType during inspection.
    /// </summary>
    public sealed class ContextualDeidConfigArgs : global::Pulumi.ResourceArgs
    {
        public ContextualDeidConfigArgs()
        {
        }
        public static new ContextualDeidConfigArgs Empty => new ContextualDeidConfigArgs();
    }
}
