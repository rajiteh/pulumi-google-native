// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1
{
    /// <summary>
    /// Creates a global PublicDelegatedPrefix in the specified project using the parameters that are included in the request.
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/v1:GlobalPublicDelegatedPrefix")]
    public partial class GlobalPublicDelegatedPrefix : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicDelegatedPrefix. An up-to-date fingerprint must be provided in order to update the PublicDelegatedPrefix, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a PublicDelegatedPrefix.
        /// </summary>
        [Output("fingerprint")]
        public Output<string> Fingerprint { get; private set; } = null!;

        /// <summary>
        /// The IPv4 address range, in CIDR format, represented by this public delegated prefix.
        /// </summary>
        [Output("ipCidrRange")]
        public Output<string> IpCidrRange { get; private set; } = null!;

        /// <summary>
        /// If true, the prefix will be live migrated.
        /// </summary>
        [Output("isLiveMigration")]
        public Output<bool> IsLiveMigration { get; private set; } = null!;

        /// <summary>
        /// Type of the resource. Always compute#publicDelegatedPrefix for public delegated prefixes.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
        /// </summary>
        [Output("parentPrefix")]
        public Output<string> ParentPrefix { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The list of sub public delegated prefixes that exist for this public delegated prefix.
        /// </summary>
        [Output("publicDelegatedSubPrefixs")]
        public Output<ImmutableArray<Outputs.PublicDelegatedPrefixPublicDelegatedSubPrefixResponse>> PublicDelegatedSubPrefixs { get; private set; } = null!;

        /// <summary>
        /// URL of the region where the public delegated prefix resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// The status of the public delegated prefix, which can be one of following values: - `INITIALIZING` The public delegated prefix is being initialized and addresses cannot be created yet. - `READY_TO_ANNOUNCE` The public delegated prefix is a live migration prefix and is active. - `ANNOUNCED` The public delegated prefix is active. - `DELETING` The public delegated prefix is being deprovsioned. 
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a GlobalPublicDelegatedPrefix resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GlobalPublicDelegatedPrefix(string name, GlobalPublicDelegatedPrefixArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:GlobalPublicDelegatedPrefix", name, args ?? new GlobalPublicDelegatedPrefixArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GlobalPublicDelegatedPrefix(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/v1:GlobalPublicDelegatedPrefix", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GlobalPublicDelegatedPrefix resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GlobalPublicDelegatedPrefix Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GlobalPublicDelegatedPrefix(name, id, options);
        }
    }

    public sealed class GlobalPublicDelegatedPrefixArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The IPv4 address range, in CIDR format, represented by this public delegated prefix.
        /// </summary>
        [Input("ipCidrRange")]
        public Input<string>? IpCidrRange { get; set; }

        /// <summary>
        /// If true, the prefix will be live migrated.
        /// </summary>
        [Input("isLiveMigration")]
        public Input<bool>? IsLiveMigration { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
        /// </summary>
        [Input("parentPrefix")]
        public Input<string>? ParentPrefix { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("publicDelegatedSubPrefixs")]
        private InputList<Inputs.PublicDelegatedPrefixPublicDelegatedSubPrefixArgs>? _publicDelegatedSubPrefixs;

        /// <summary>
        /// The list of sub public delegated prefixes that exist for this public delegated prefix.
        /// </summary>
        public InputList<Inputs.PublicDelegatedPrefixPublicDelegatedSubPrefixArgs> PublicDelegatedSubPrefixs
        {
            get => _publicDelegatedSubPrefixs ?? (_publicDelegatedSubPrefixs = new InputList<Inputs.PublicDelegatedPrefixPublicDelegatedSubPrefixArgs>());
            set => _publicDelegatedSubPrefixs = value;
        }

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        public GlobalPublicDelegatedPrefixArgs()
        {
        }
        public static new GlobalPublicDelegatedPrefixArgs Empty => new GlobalPublicDelegatedPrefixArgs();
    }
}
