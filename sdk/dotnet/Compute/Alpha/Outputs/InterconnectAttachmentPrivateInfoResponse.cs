// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    /// <summary>
    /// Information for an interconnect attachment when this belongs to an interconnect of type DEDICATED.
    /// </summary>
    [OutputType]
    public sealed class InterconnectAttachmentPrivateInfoResponse
    {
        /// <summary>
        /// 802.1q encapsulation tag to be used for traffic between Google and the customer, going to and from this network and region.
        /// </summary>
        public readonly int Tag8021q;

        [OutputConstructor]
        private InterconnectAttachmentPrivateInfoResponse(int tag8021q)
        {
            Tag8021q = tag8021q;
        }
    }
}
