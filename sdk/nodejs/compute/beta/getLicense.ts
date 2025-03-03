// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns the specified License resource. *Caution* This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.
 */
export function getLicense(args: GetLicenseArgs, opts?: pulumi.InvokeOptions): Promise<GetLicenseResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/beta:getLicense", {
        "license": args.license,
        "project": args.project,
    }, opts);
}

export interface GetLicenseArgs {
    license: string;
    project?: string;
}

export interface GetLicenseResult {
    /**
     * Deprecated. This field no longer reflects whether a license charges a usage fee.
     *
     * @deprecated [Output Only] Deprecated. This field no longer reflects whether a license charges a usage fee.
     */
    readonly chargesUseFee: boolean;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional textual description of the resource; provided by the client when the resource is created.
     */
    readonly description: string;
    /**
     * Type of resource. Always compute#license for licenses.
     */
    readonly kind: string;
    /**
     * The unique code used to attach this license to images, snapshots, and disks.
     */
    readonly licenseCode: string;
    /**
     * Name of the resource. The name must be 1-63 characters long and comply with RFC1035.
     */
    readonly name: string;
    readonly resourceRequirements: outputs.compute.beta.LicenseResourceRequirementsResponse;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * If false, licenses will not be copied from the source resource when creating an image from a disk, disk from snapshot, or snapshot from disk.
     */
    readonly transferable: boolean;
}
/**
 * Returns the specified License resource. *Caution* This resource is intended for use only by third-party partners who are creating Cloud Marketplace images.
 */
export function getLicenseOutput(args: GetLicenseOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLicenseResult> {
    return pulumi.output(args).apply((a: any) => getLicense(a, opts))
}

export interface GetLicenseOutputArgs {
    license: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
