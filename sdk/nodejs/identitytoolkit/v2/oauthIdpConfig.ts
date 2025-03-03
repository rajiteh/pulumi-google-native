// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Create an Oidc Idp configuration for an Identity Toolkit project.
 */
export class OauthIdpConfig extends pulumi.CustomResource {
    /**
     * Get an existing OauthIdpConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OauthIdpConfig {
        return new OauthIdpConfig(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:identitytoolkit/v2:OauthIdpConfig';

    /**
     * Returns true if the given object is an instance of OauthIdpConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OauthIdpConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OauthIdpConfig.__pulumiType;
    }

    /**
     * The client id of an OAuth client.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * The client secret of the OAuth client, to enable OIDC code flow.
     */
    public readonly clientSecret!: pulumi.Output<string>;
    /**
     * The config's display name set by developers.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * True if allows the user to sign in with the provider.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * For OIDC Idps, the issuer identifier.
     */
    public readonly issuer!: pulumi.Output<string>;
    /**
     * The name of the OAuthIdpConfig resource, for example: 'projects/my-awesome-project/oauthIdpConfigs/oauth-config-id'. Ignored during create requests.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The id to use for this config.
     */
    public readonly oauthIdpConfigId!: pulumi.Output<string | undefined>;
    public readonly project!: pulumi.Output<string>;
    /**
     * The response type to request for in the OAuth authorization flow. You can set either `id_token` or `code` to true, but not both. Setting both types to be simultaneously true (`{code: true, id_token: true}`) is not yet supported.
     */
    public readonly responseType!: pulumi.Output<outputs.identitytoolkit.v2.GoogleCloudIdentitytoolkitAdminV2OAuthResponseTypeResponse>;
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a OauthIdpConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OauthIdpConfigArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientSecret"] = args ? args.clientSecret : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["issuer"] = args ? args.issuer : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["oauthIdpConfigId"] = args ? args.oauthIdpConfigId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["responseType"] = args ? args.responseType : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
        } else {
            resourceInputs["clientId"] = undefined /*out*/;
            resourceInputs["clientSecret"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["enabled"] = undefined /*out*/;
            resourceInputs["issuer"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["oauthIdpConfigId"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["responseType"] = undefined /*out*/;
            resourceInputs["tenantId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["project", "tenantId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(OauthIdpConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a OauthIdpConfig resource.
 */
export interface OauthIdpConfigArgs {
    /**
     * The client id of an OAuth client.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The client secret of the OAuth client, to enable OIDC code flow.
     */
    clientSecret?: pulumi.Input<string>;
    /**
     * The config's display name set by developers.
     */
    displayName?: pulumi.Input<string>;
    /**
     * True if allows the user to sign in with the provider.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * For OIDC Idps, the issuer identifier.
     */
    issuer?: pulumi.Input<string>;
    /**
     * The name of the OAuthIdpConfig resource, for example: 'projects/my-awesome-project/oauthIdpConfigs/oauth-config-id'. Ignored during create requests.
     */
    name?: pulumi.Input<string>;
    /**
     * The id to use for this config.
     */
    oauthIdpConfigId?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The response type to request for in the OAuth authorization flow. You can set either `id_token` or `code` to true, but not both. Setting both types to be simultaneously true (`{code: true, id_token: true}`) is not yet supported.
     */
    responseType?: pulumi.Input<inputs.identitytoolkit.v2.GoogleCloudIdentitytoolkitAdminV2OAuthResponseTypeArgs>;
    tenantId: pulumi.Input<string>;
}
