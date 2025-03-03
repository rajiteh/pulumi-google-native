// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets a view on a log bucket..
 */
export function getOrganizationBucketView(args: GetOrganizationBucketViewArgs, opts?: pulumi.InvokeOptions): Promise<GetOrganizationBucketViewResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:logging/v2:getOrganizationBucketView", {
        "bucketId": args.bucketId,
        "location": args.location,
        "organizationId": args.organizationId,
        "viewId": args.viewId,
    }, opts);
}

export interface GetOrganizationBucketViewArgs {
    bucketId: string;
    location: string;
    organizationId: string;
    viewId: string;
}

export interface GetOrganizationBucketViewResult {
    /**
     * The creation timestamp of the view.
     */
    readonly createTime: string;
    /**
     * Describes this view.
     */
    readonly description: string;
    /**
     * Filter that restricts which log entries in a bucket are visible in this view.Filters are restricted to be a logical AND of ==/!= of any of the following: originating project/folder/organization/billing account. resource type log idFor example:SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
     */
    readonly filter: string;
    /**
     * The resource name of the view.For example:projects/my-project/locations/global/buckets/my-bucket/views/my-view
     */
    readonly name: string;
    /**
     * The last update timestamp of the view.
     */
    readonly updateTime: string;
}
/**
 * Gets a view on a log bucket..
 */
export function getOrganizationBucketViewOutput(args: GetOrganizationBucketViewOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrganizationBucketViewResult> {
    return pulumi.output(args).apply((a: any) => getOrganizationBucketView(a, opts))
}

export interface GetOrganizationBucketViewOutputArgs {
    bucketId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
    viewId: pulumi.Input<string>;
}
