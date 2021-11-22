// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Connectors.V1.Outputs
{

    /// <summary>
    /// JWT claims used for the jwt-bearer authorization grant.
    /// </summary>
    [OutputType]
    public sealed class JwtClaimsResponse
    {
        /// <summary>
        /// Value for the "aud" claim.
        /// </summary>
        public readonly string Audience;
        /// <summary>
        /// Value for the "iss" claim.
        /// </summary>
        public readonly string Issuer;
        /// <summary>
        /// Value for the "sub" claim.
        /// </summary>
        public readonly string Subject;

        [OutputConstructor]
        private JwtClaimsResponse(
            string audience,

            string issuer,

            string subject)
        {
            Audience = audience;
            Issuer = issuer;
            Subject = subject;
        }
    }
}