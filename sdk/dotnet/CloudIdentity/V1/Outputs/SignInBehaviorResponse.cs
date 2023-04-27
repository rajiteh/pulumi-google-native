// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudIdentity.V1.Outputs
{

    /// <summary>
    /// Controls sign-in behavior.
    /// </summary>
    [OutputType]
    public sealed class SignInBehaviorResponse
    {
        /// <summary>
        /// When to redirect sign-ins to the IdP.
        /// </summary>
        public readonly string RedirectCondition;

        [OutputConstructor]
        private SignInBehaviorResponse(string redirectCondition)
        {
            RedirectCondition = redirectCondition;
        }
    }
}