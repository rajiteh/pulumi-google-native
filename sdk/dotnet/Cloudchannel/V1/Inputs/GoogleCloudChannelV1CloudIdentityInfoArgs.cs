// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleCloud.Cloudchannel.V1.Inputs
{

    /// <summary>
    /// Cloud Identity information for the Cloud Channel Customer.
    /// </summary>
    public sealed class GoogleCloudChannelV1CloudIdentityInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Output only. URI of Customer's Admin console dashboard.
        /// </summary>
        [Input("adminConsoleUri")]
        public Input<string>? AdminConsoleUri { get; set; }

        /// <summary>
        /// The alternate email.
        /// </summary>
        [Input("alternateEmail")]
        public Input<string>? AlternateEmail { get; set; }

        /// <summary>
        /// CustomerType indicates verification type needed for using services.
        /// </summary>
        [Input("customerType")]
        public Input<string>? CustomerType { get; set; }

        /// <summary>
        /// Edu information about the customer.
        /// </summary>
        [Input("eduData")]
        public Input<Inputs.GoogleCloudChannelV1EduDataArgs>? EduData { get; set; }

        /// <summary>
        /// Whether the domain is verified.
        /// </summary>
        [Input("isDomainVerified")]
        public Input<bool>? IsDomainVerified { get; set; }

        /// <summary>
        /// Language code.
        /// </summary>
        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        /// <summary>
        /// Phone number associated with the Cloud Identity.
        /// </summary>
        [Input("phoneNumber")]
        public Input<string>? PhoneNumber { get; set; }

        /// <summary>
        /// Output only. The primary domain name.
        /// </summary>
        [Input("primaryDomain")]
        public Input<string>? PrimaryDomain { get; set; }

        public GoogleCloudChannelV1CloudIdentityInfoArgs()
        {
        }
    }
}