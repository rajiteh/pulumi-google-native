// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.CloudChannel.V1
{
    /// <summary>
    /// Registers a service account with subscriber privileges on the Cloud Pub/Sub topic for this Channel Services account. After you create a subscriber, you get the events through SubscriberEvent Possible error codes: * PERMISSION_DENIED: The reseller account making the request and the provided reseller account are different, or the impersonated user is not a super admin. * INVALID_ARGUMENT: Required request parameters are missing or invalid. * INTERNAL: Any non-user error related to a technical issue in the backend. Contact Cloud Channel support. * UNKNOWN: Any non-user error related to a technical issue in the backend. Contact Cloud Channel support. Return value: The topic name with the registered service email address.
    /// </summary>
    [GcpNativeResourceType("gcp-native:cloudchannel/v1:Account")]
    public partial class Account : Pulumi.CustomResource
    {
        /// <summary>
        /// Create a Account resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Account(string name, AccountArgs args, CustomResourceOptions? options = null)
            : base("gcp-native:cloudchannel/v1:Account", name, args ?? new AccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Account(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("gcp-native:cloudchannel/v1:Account", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Account resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Account Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Account(name, id, options);
        }
    }

    public sealed class AccountArgs : Pulumi.ResourceArgs
    {
        [Input("accountsId", required: true)]
        public Input<string> AccountsId { get; set; } = null!;

        /// <summary>
        /// Required. Service account that provides subscriber access to the registered topic.
        /// </summary>
        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        public AccountArgs()
        {
        }
    }
}