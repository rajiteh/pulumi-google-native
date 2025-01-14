// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.IdentityToolkit.V2
{
    /// <summary>
    /// Create an inbound SAML configuration for an Identity Toolkit project.
    /// </summary>
    [GoogleNativeResourceType("google-native:identitytoolkit/v2:InboundSamlConfig")]
    public partial class InboundSamlConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The config's display name set by developers.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// True if allows the user to sign in with the provider.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// The SAML IdP (Identity Provider) configuration when the project acts as the relying party.
        /// </summary>
        [Output("idpConfig")]
        public Output<Outputs.GoogleCloudIdentitytoolkitAdminV2IdpConfigResponse> IdpConfig { get; private set; } = null!;

        /// <summary>
        /// The id to use for this config.
        /// </summary>
        [Output("inboundSamlConfigId")]
        public Output<string?> InboundSamlConfigId { get; private set; } = null!;

        /// <summary>
        /// The name of the InboundSamlConfig resource, for example: 'projects/my-awesome-project/inboundSamlConfigs/my-config-id'. Ignored during create requests.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The SAML SP (Service Provider) configuration when the project acts as the relying party to receive and accept an authentication assertion issued by a SAML identity provider.
        /// </summary>
        [Output("spConfig")]
        public Output<Outputs.GoogleCloudIdentitytoolkitAdminV2SpConfigResponse> SpConfig { get; private set; } = null!;

        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a InboundSamlConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InboundSamlConfig(string name, InboundSamlConfigArgs args, CustomResourceOptions? options = null)
            : base("google-native:identitytoolkit/v2:InboundSamlConfig", name, args ?? new InboundSamlConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InboundSamlConfig(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:identitytoolkit/v2:InboundSamlConfig", name, null, MakeResourceOptions(options, id))
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
                    "tenantId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing InboundSamlConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InboundSamlConfig Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new InboundSamlConfig(name, id, options);
        }
    }

    public sealed class InboundSamlConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The config's display name set by developers.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// True if allows the user to sign in with the provider.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The SAML IdP (Identity Provider) configuration when the project acts as the relying party.
        /// </summary>
        [Input("idpConfig")]
        public Input<Inputs.GoogleCloudIdentitytoolkitAdminV2IdpConfigArgs>? IdpConfig { get; set; }

        /// <summary>
        /// The id to use for this config.
        /// </summary>
        [Input("inboundSamlConfigId")]
        public Input<string>? InboundSamlConfigId { get; set; }

        /// <summary>
        /// The name of the InboundSamlConfig resource, for example: 'projects/my-awesome-project/inboundSamlConfigs/my-config-id'. Ignored during create requests.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The SAML SP (Service Provider) configuration when the project acts as the relying party to receive and accept an authentication assertion issued by a SAML identity provider.
        /// </summary>
        [Input("spConfig")]
        public Input<Inputs.GoogleCloudIdentitytoolkitAdminV2SpConfigArgs>? SpConfig { get; set; }

        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

        public InboundSamlConfigArgs()
        {
        }
        public static new InboundSamlConfigArgs Empty => new InboundSamlConfigArgs();
    }
}
