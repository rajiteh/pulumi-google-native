// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Privateca.V1.Inputs
{

    /// <summary>
    /// Subject describes parts of a distinguished name that, in turn, describes the subject of the certificate.
    /// </summary>
    public sealed class SubjectArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The "common name" of the subject.
        /// </summary>
        [Input("commonName")]
        public Input<string>? CommonName { get; set; }

        /// <summary>
        /// The country code of the subject.
        /// </summary>
        [Input("countryCode")]
        public Input<string>? CountryCode { get; set; }

        /// <summary>
        /// The locality or city of the subject.
        /// </summary>
        [Input("locality")]
        public Input<string>? Locality { get; set; }

        /// <summary>
        /// The organization of the subject.
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        /// <summary>
        /// The organizational_unit of the subject.
        /// </summary>
        [Input("organizationalUnit")]
        public Input<string>? OrganizationalUnit { get; set; }

        /// <summary>
        /// The postal code of the subject.
        /// </summary>
        [Input("postalCode")]
        public Input<string>? PostalCode { get; set; }

        /// <summary>
        /// The province, territory, or regional state of the subject.
        /// </summary>
        [Input("province")]
        public Input<string>? Province { get; set; }

        /// <summary>
        /// The street address of the subject.
        /// </summary>
        [Input("streetAddress")]
        public Input<string>? StreetAddress { get; set; }

        public SubjectArgs()
        {
        }
    }
}