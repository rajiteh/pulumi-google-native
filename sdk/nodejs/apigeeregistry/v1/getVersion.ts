// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Returns a specified version.
 */
export function getVersion(args: GetVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetVersionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:apigeeregistry/v1:getVersion", {
        "apiId": args.apiId,
        "location": args.location,
        "project": args.project,
        "versionId": args.versionId,
    }, opts);
}

export interface GetVersionArgs {
    apiId: string;
    location: string;
    project?: string;
    versionId: string;
}

export interface GetVersionResult {
    /**
     * Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
     */
    readonly annotations: {[key: string]: string};
    /**
     * Creation timestamp.
     */
    readonly createTime: string;
    /**
     * A detailed description.
     */
    readonly description: string;
    /**
     * Human-meaningful name.
     */
    readonly displayName: string;
    /**
     * Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with `apigeeregistry.googleapis.com/` and cannot be changed.
     */
    readonly labels: {[key: string]: string};
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The primary spec for this version. Format: projects/{project}/locations/{location}/apis/{api}/versions/{version}/specs/{spec}
     */
    readonly primarySpec: string;
    /**
     * A user-definable description of the lifecycle phase of this API version. Format: free-form, but we expect single words that describe API maturity, e.g., "CONCEPT", "DESIGN", "DEVELOPMENT", "STAGING", "PRODUCTION", "DEPRECATED", "RETIRED".
     */
    readonly state: string;
    /**
     * Last update timestamp.
     */
    readonly updateTime: string;
}
/**
 * Returns a specified version.
 */
export function getVersionOutput(args: GetVersionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVersionResult> {
    return pulumi.output(args).apply((a: any) => getVersion(a, opts))
}

export interface GetVersionOutputArgs {
    apiId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    versionId: pulumi.Input<string>;
}
