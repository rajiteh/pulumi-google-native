// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AccessContextManager.V1.Inputs
{

    /// <summary>
    /// Defines the conditions under which an EgressPolicy matches a request. Conditions based on information about the source of the request. Note that if the destination of the request is also protected by a ServicePerimeter, then that ServicePerimeter must have an IngressPolicy which allows access in order for this request to succeed.
    /// </summary>
    public sealed class EgressFromArgs : global::Pulumi.ResourceArgs
    {
        [Input("identities")]
        private InputList<string>? _identities;

        /// <summary>
        /// A list of identities that are allowed access through this [EgressPolicy]. Should be in the format of email address. The email address should represent individual user or service account only.
        /// </summary>
        public InputList<string> Identities
        {
            get => _identities ?? (_identities = new InputList<string>());
            set => _identities = value;
        }

        /// <summary>
        /// Specifies the type of identities that are allowed access to outside the perimeter. If left unspecified, then members of `identities` field will be allowed access.
        /// </summary>
        [Input("identityType")]
        public Input<Pulumi.GoogleNative.AccessContextManager.V1.EgressFromIdentityType>? IdentityType { get; set; }

        public EgressFromArgs()
        {
        }
        public static new EgressFromArgs Empty => new EgressFromArgs();
    }
}
