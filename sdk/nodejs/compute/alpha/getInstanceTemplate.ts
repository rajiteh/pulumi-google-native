// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns the specified instance template.
 */
export function getInstanceTemplate(args: GetInstanceTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceTemplateResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/alpha:getInstanceTemplate", {
        "instanceTemplate": args.instanceTemplate,
        "project": args.project,
    }, opts);
}

export interface GetInstanceTemplateArgs {
    instanceTemplate: string;
    project?: string;
}

export interface GetInstanceTemplateResult {
    /**
     * The creation timestamp for this instance template in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * The resource type, which is always compute#instanceTemplate for instance templates.
     */
    readonly kind: string;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * The instance properties for this instance template.
     */
    readonly properties: outputs.compute.alpha.InstancePropertiesResponse;
    /**
     * URL of the region where the instance template resides. Only applicable for regional resources.
     */
    readonly region: string;
    /**
     * The URL for this instance template. The server defines this URL.
     */
    readonly selfLink: string;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    readonly selfLinkWithId: string;
    /**
     * The source instance used to create the template. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance 
     */
    readonly sourceInstance: string;
    /**
     * The source instance params to use to create this instance template.
     */
    readonly sourceInstanceParams: outputs.compute.alpha.SourceInstanceParamsResponse;
}
/**
 * Returns the specified instance template.
 */
export function getInstanceTemplateOutput(args: GetInstanceTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceTemplateResult> {
    return pulumi.output(args).apply((a: any) => getInstanceTemplate(a, opts))
}

export interface GetInstanceTemplateOutputArgs {
    instanceTemplate: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
