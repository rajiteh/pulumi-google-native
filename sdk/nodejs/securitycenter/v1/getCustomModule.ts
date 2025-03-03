// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Retrieves a SecurityHealthAnalyticsCustomModule.
 */
export function getCustomModule(args: GetCustomModuleArgs, opts?: pulumi.InvokeOptions): Promise<GetCustomModuleResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:securitycenter/v1:getCustomModule", {
        "customModuleId": args.customModuleId,
        "project": args.project,
    }, opts);
}

export interface GetCustomModuleArgs {
    customModuleId: string;
    project?: string;
}

export interface GetCustomModuleResult {
    /**
     * If empty, indicates that the custom module was created in the organization, folder, or project in which you are viewing the custom module. Otherwise, `ancestor_module` specifies the organization or folder from which the custom module is inherited.
     */
    readonly ancestorModule: string;
    /**
     * The user specified custom configuration for the module.
     */
    readonly customConfig: outputs.securitycenter.v1.GoogleCloudSecuritycenterV1CustomConfigResponse;
    /**
     * The display name of the Security Health Analytics custom module. This display name becomes the finding category for all findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a lowercase letter, and contain alphanumeric characters or underscores only.
     */
    readonly displayName: string;
    /**
     * The enablement state of the custom module.
     */
    readonly enablementState: string;
    /**
     * The editor that last updated the custom module.
     */
    readonly lastEditor: string;
    /**
     * Immutable. The resource name of the custom module. Its format is "organizations/{organization}/securityHealthAnalyticsSettings/customModules/{customModule}", or "folders/{folder}/securityHealthAnalyticsSettings/customModules/{customModule}", or "projects/{project}/securityHealthAnalyticsSettings/customModules/{customModule}" The id {customModule} is server-generated and is not user settable. It will be a numeric id containing 1-20 digits.
     */
    readonly name: string;
    /**
     * The time at which the custom module was last updated.
     */
    readonly updateTime: string;
}
/**
 * Retrieves a SecurityHealthAnalyticsCustomModule.
 */
export function getCustomModuleOutput(args: GetCustomModuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCustomModuleResult> {
    return pulumi.output(args).apply((a: any) => getCustomModule(a, opts))
}

export interface GetCustomModuleOutputArgs {
    customModuleId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
