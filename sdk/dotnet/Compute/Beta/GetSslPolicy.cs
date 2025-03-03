// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta
{
    public static class GetSslPolicy
    {
        /// <summary>
        /// Lists all of the ordered rules present in a single specified policy.
        /// </summary>
        public static Task<GetSslPolicyResult> InvokeAsync(GetSslPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSslPolicyResult>("google-native:compute/beta:getSslPolicy", args ?? new GetSslPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Lists all of the ordered rules present in a single specified policy.
        /// </summary>
        public static Output<GetSslPolicyResult> Invoke(GetSslPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSslPolicyResult>("google-native:compute/beta:getSslPolicy", args ?? new GetSslPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSslPolicyArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("sslPolicy", required: true)]
        public string SslPolicy { get; set; } = null!;

        public GetSslPolicyArgs()
        {
        }
        public static new GetSslPolicyArgs Empty => new GetSslPolicyArgs();
    }

    public sealed class GetSslPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("sslPolicy", required: true)]
        public Input<string> SslPolicy { get; set; } = null!;

        public GetSslPolicyInvokeArgs()
        {
        }
        public static new GetSslPolicyInvokeArgs Empty => new GetSslPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetSslPolicyResult
    {
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// A list of features enabled when the selected profile is CUSTOM. The method returns the set of features that can be specified in this list. This field must be empty if the profile is not CUSTOM.
        /// </summary>
        public readonly ImmutableArray<string> CustomFeatures;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The list of features enabled in the SSL policy.
        /// </summary>
        public readonly ImmutableArray<string> EnabledFeatures;
        /// <summary>
        /// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a SslPolicy. An up-to-date fingerprint must be provided in order to update the SslPolicy, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an SslPolicy.
        /// </summary>
        public readonly string Fingerprint;
        /// <summary>
        /// [Output only] Type of the resource. Always compute#sslPolicyfor SSL policies.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of TLS_1_0, TLS_1_1, TLS_1_2.
        /// </summary>
        public readonly string MinTlsVersion;
        /// <summary>
        /// Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field.
        /// </summary>
        public readonly string Profile;
        /// <summary>
        /// URL of the region where the regional SSL policy resides. This field is not applicable to global SSL policies.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Server-defined URL for the resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// If potential misconfigurations are detected for this SSL policy, this field will be populated with warning messages.
        /// </summary>
        public readonly ImmutableArray<Outputs.SslPolicyWarningsItemResponse> Warnings;

        [OutputConstructor]
        private GetSslPolicyResult(
            string creationTimestamp,

            ImmutableArray<string> customFeatures,

            string description,

            ImmutableArray<string> enabledFeatures,

            string fingerprint,

            string kind,

            string minTlsVersion,

            string name,

            string profile,

            string region,

            string selfLink,

            ImmutableArray<Outputs.SslPolicyWarningsItemResponse> warnings)
        {
            CreationTimestamp = creationTimestamp;
            CustomFeatures = customFeatures;
            Description = description;
            EnabledFeatures = enabledFeatures;
            Fingerprint = fingerprint;
            Kind = kind;
            MinTlsVersion = minTlsVersion;
            Name = name;
            Profile = profile;
            Region = region;
            SelfLink = selfLink;
            Warnings = warnings;
        }
    }
}
