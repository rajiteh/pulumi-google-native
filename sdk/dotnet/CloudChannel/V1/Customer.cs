// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudChannel.V1
{
    /// <summary>
    /// Creates a new Customer resource under the reseller or distributor account. Possible error codes: * PERMISSION_DENIED: The reseller account making the request is different from the reseller account in the API request. * INVALID_ARGUMENT: * Required request parameters are missing or invalid. * Domain field value doesn't match the primary email domain. Return value: The newly created Customer resource.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:cloudchannel/v1:Customer")]
    public partial class Customer : global::Pulumi.CustomResource
    {
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// Secondary contact email. You need to provide an alternate email to create different domains if a primary contact email already exists. Users will receive a notification with credentials when you create an admin.google.com account. Secondary emails are also recovery email addresses. Alternate emails are optional when you create Team customers.
        /// </summary>
        [Output("alternateEmail")]
        public Output<string> AlternateEmail { get; private set; } = null!;

        /// <summary>
        /// Cloud Identity ID of the customer's channel partner. Populated only if a channel partner exists for this customer.
        /// </summary>
        [Output("channelPartnerId")]
        public Output<string> ChannelPartnerId { get; private set; } = null!;

        [Output("channelPartnerLinkId")]
        public Output<string> ChannelPartnerLinkId { get; private set; } = null!;

        /// <summary>
        /// The customer's Cloud Identity ID if the customer has a Cloud Identity resource.
        /// </summary>
        [Output("cloudIdentityId")]
        public Output<string> CloudIdentityId { get; private set; } = null!;

        /// <summary>
        /// Cloud Identity information for the customer. Populated only if a Cloud Identity account exists for this customer.
        /// </summary>
        [Output("cloudIdentityInfo")]
        public Output<Outputs.GoogleCloudChannelV1CloudIdentityInfoResponse> CloudIdentityInfo { get; private set; } = null!;

        /// <summary>
        /// Optional. External CRM ID for the customer. Populated only if a CRM ID exists for this customer.
        /// </summary>
        [Output("correlationId")]
        public Output<string> CorrelationId { get; private set; } = null!;

        /// <summary>
        /// Time when the customer was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The customer's primary domain. Must match the primary contact email's domain.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Optional. The BCP-47 language code, such as "en-US" or "sr-Latn". For more information, see https://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
        /// </summary>
        [Output("languageCode")]
        public Output<string> LanguageCode { get; private set; } = null!;

        /// <summary>
        /// Resource name of the customer. Format: accounts/{account_id}/customers/{customer_id}
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Name of the organization that the customer entity represents.
        /// </summary>
        [Output("orgDisplayName")]
        public Output<string> OrgDisplayName { get; private set; } = null!;

        /// <summary>
        /// The organization address for the customer. To enforce US laws and embargoes, we require a region and zip code. You must provide valid addresses for every customer. To set the customer's language, use the Customer-level language code.
        /// </summary>
        [Output("orgPostalAddress")]
        public Output<Outputs.GoogleTypePostalAddressResponse> OrgPostalAddress { get; private set; } = null!;

        /// <summary>
        /// Primary contact info.
        /// </summary>
        [Output("primaryContactInfo")]
        public Output<Outputs.GoogleCloudChannelV1ContactInfoResponse> PrimaryContactInfo { get; private set; } = null!;

        /// <summary>
        /// Time when the customer was updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Customer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Customer(string name, CustomerArgs args, CustomResourceOptions? options = null)
            : base("google-native:cloudchannel/v1:Customer", name, args ?? new CustomerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Customer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:cloudchannel/v1:Customer", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "accountId",
                    "channelPartnerLinkId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Customer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Customer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Customer(name, id, options);
        }
    }

    public sealed class CustomerArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        /// <summary>
        /// Secondary contact email. You need to provide an alternate email to create different domains if a primary contact email already exists. Users will receive a notification with credentials when you create an admin.google.com account. Secondary emails are also recovery email addresses. Alternate emails are optional when you create Team customers.
        /// </summary>
        [Input("alternateEmail")]
        public Input<string>? AlternateEmail { get; set; }

        /// <summary>
        /// Cloud Identity ID of the customer's channel partner. Populated only if a channel partner exists for this customer.
        /// </summary>
        [Input("channelPartnerId")]
        public Input<string>? ChannelPartnerId { get; set; }

        [Input("channelPartnerLinkId", required: true)]
        public Input<string> ChannelPartnerLinkId { get; set; } = null!;

        /// <summary>
        /// Optional. External CRM ID for the customer. Populated only if a CRM ID exists for this customer.
        /// </summary>
        [Input("correlationId")]
        public Input<string>? CorrelationId { get; set; }

        /// <summary>
        /// The customer's primary domain. Must match the primary contact email's domain.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Optional. The BCP-47 language code, such as "en-US" or "sr-Latn". For more information, see https://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
        /// </summary>
        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        /// <summary>
        /// Name of the organization that the customer entity represents.
        /// </summary>
        [Input("orgDisplayName", required: true)]
        public Input<string> OrgDisplayName { get; set; } = null!;

        /// <summary>
        /// The organization address for the customer. To enforce US laws and embargoes, we require a region and zip code. You must provide valid addresses for every customer. To set the customer's language, use the Customer-level language code.
        /// </summary>
        [Input("orgPostalAddress", required: true)]
        public Input<Inputs.GoogleTypePostalAddressArgs> OrgPostalAddress { get; set; } = null!;

        /// <summary>
        /// Primary contact info.
        /// </summary>
        [Input("primaryContactInfo")]
        public Input<Inputs.GoogleCloudChannelV1ContactInfoArgs>? PrimaryContactInfo { get; set; }

        public CustomerArgs()
        {
        }
        public static new CustomerArgs Empty => new CustomerArgs();
    }
}
