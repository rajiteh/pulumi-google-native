// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets the details of a `Registration` resource.
 */
export function getRegistration(args: GetRegistrationArgs, opts?: pulumi.InvokeOptions): Promise<GetRegistrationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:domains/v1beta1:getRegistration", {
        "location": args.location,
        "project": args.project,
        "registrationId": args.registrationId,
    }, opts);
}

export interface GetRegistrationArgs {
    location: string;
    project?: string;
    registrationId: string;
}

export interface GetRegistrationResult {
    /**
     * Settings for contact information linked to the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureContactSettings` method.
     */
    readonly contactSettings: outputs.domains.v1beta1.ContactSettingsResponse;
    /**
     * The creation timestamp of the `Registration` resource.
     */
    readonly createTime: string;
    /**
     * Settings controlling the DNS configuration of the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureDnsSettings` method.
     */
    readonly dnsSettings: outputs.domains.v1beta1.DnsSettingsResponse;
    /**
     * Immutable. The domain name. Unicode domain names must be expressed in Punycode format.
     */
    readonly domainName: string;
    /**
     * The expiration timestamp of the `Registration`.
     */
    readonly expireTime: string;
    /**
     * The set of issues with the `Registration` that require attention.
     */
    readonly issues: string[];
    /**
     * Set of labels associated with the `Registration`.
     */
    readonly labels: {[key: string]: string};
    /**
     * Settings for management of the `Registration`, including renewal, billing, and transfer. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureManagementSettings` method.
     */
    readonly managementSettings: outputs.domains.v1beta1.ManagementSettingsResponse;
    /**
     * Name of the `Registration` resource, in the format `projects/*&#47;locations/*&#47;registrations/`.
     */
    readonly name: string;
    /**
     * Pending contact settings for the `Registration`. Updates to the `contact_settings` field that change its `registrant_contact` or `privacy` fields require email confirmation by the `registrant_contact` before taking effect. This field is set only if there are pending updates to the `contact_settings` that have not been confirmed. To confirm the changes, the `registrant_contact` must follow the instructions in the email they receive.
     */
    readonly pendingContactSettings: outputs.domains.v1beta1.ContactSettingsResponse;
    /**
     * The reason the domain registration failed. Only set for domains in REGISTRATION_FAILED state.
     */
    readonly registerFailureReason: string;
    /**
     * The state of the `Registration`
     */
    readonly state: string;
    /**
     * Set of options for the `contact_settings.privacy` field that this `Registration` supports.
     */
    readonly supportedPrivacy: string[];
    /**
     * The reason the domain transfer failed. Only set for domains in TRANSFER_FAILED state.
     */
    readonly transferFailureReason: string;
}
/**
 * Gets the details of a `Registration` resource.
 */
export function getRegistrationOutput(args: GetRegistrationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegistrationResult> {
    return pulumi.output(args).apply((a: any) => getRegistration(a, opts))
}

export interface GetRegistrationOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    registrationId: pulumi.Input<string>;
}
