// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets environment details.
 */
export function getEnvironment(args: GetEnvironmentArgs, opts?: pulumi.InvokeOptions): Promise<GetEnvironmentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:apigee/v1:getEnvironment", {
        "environmentId": args.environmentId,
        "organizationId": args.organizationId,
    }, opts);
}

export interface GetEnvironmentArgs {
    environmentId: string;
    organizationId: string;
}

export interface GetEnvironmentResult {
    /**
     * Creation time of this environment as milliseconds since epoch.
     */
    readonly createdAt: string;
    /**
     * Optional. Description of the environment.
     */
    readonly description: string;
    /**
     * Optional. Display name for this environment.
     */
    readonly displayName: string;
    /**
     * Last modification time of this environment as milliseconds since epoch.
     */
    readonly lastModifiedAt: string;
    /**
     * Required. Name of the environment. Values must match the regular expression `^[.\\p{Alnum}-_]{1,255}$`
     */
    readonly name: string;
    /**
     * Optional. Key-value pairs that may be used for customizing the environment.
     */
    readonly properties: outputs.apigee.v1.GoogleCloudApigeeV1PropertiesResponse;
    /**
     * State of the environment. Values other than ACTIVE means the resource is not ready to use.
     */
    readonly state: string;
}