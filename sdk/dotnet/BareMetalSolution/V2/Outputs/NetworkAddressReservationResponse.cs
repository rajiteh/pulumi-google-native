// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BareMetalSolution.V2.Outputs
{

    /// <summary>
    /// A reservation of one or more addresses in a network.
    /// </summary>
    [OutputType]
    public sealed class NetworkAddressReservationResponse
    {
        /// <summary>
        /// The last address of this reservation block, inclusive. I.e., for cases when reservations are only single addresses, end_address and start_address will be the same. Must be specified as a single IPv4 address, e.g. 10.1.2.2.
        /// </summary>
        public readonly string EndAddress;
        /// <summary>
        /// A note about this reservation, intended for human consumption.
        /// </summary>
        public readonly string Note;
        /// <summary>
        /// The first address of this reservation block. Must be specified as a single IPv4 address, e.g. 10.1.2.2.
        /// </summary>
        public readonly string StartAddress;

        [OutputConstructor]
        private NetworkAddressReservationResponse(
            string endAddress,

            string note,

            string startAddress)
        {
            EndAddress = endAddress;
            Note = note;
            StartAddress = startAddress;
        }
    }
}
