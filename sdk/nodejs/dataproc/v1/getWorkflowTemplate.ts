// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Retrieves the latest workflow template.Can retrieve previously instantiated template by specifying optional version parameter.
 */
export function getWorkflowTemplate(args: GetWorkflowTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkflowTemplateResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:dataproc/v1:getWorkflowTemplate", {
        "location": args.location,
        "project": args.project,
        "version": args.version,
        "workflowTemplateId": args.workflowTemplateId,
    }, opts);
}

export interface GetWorkflowTemplateArgs {
    location: string;
    project?: string;
    version?: number;
    workflowTemplateId: string;
}

export interface GetWorkflowTemplateResult {
    /**
     * The time template was created.
     */
    readonly createTime: string;
    /**
     * Optional. Timeout duration for the DAG of jobs, expressed in seconds (see JSON representation of duration (https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10 minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a managed cluster, the cluster is deleted.
     */
    readonly dagTimeout: string;
    /**
     * The Directed Acyclic Graph of Jobs to submit.
     */
    readonly jobs: outputs.dataproc.v1.OrderedJobResponse[];
    /**
     * Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance.Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).No more than 32 labels can be associated with a template.
     */
    readonly labels: {[key: string]: string};
    /**
     * The resource name of the workflow template, as described in https://cloud.google.com/apis/design/resource_names. For projects.regions.workflowTemplates, the resource name of the template has the following format: projects/{project_id}/regions/{region}/workflowTemplates/{template_id} For projects.locations.workflowTemplates, the resource name of the template has the following format: projects/{project_id}/locations/{location}/workflowTemplates/{template_id}
     */
    readonly name: string;
    /**
     * Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
     */
    readonly parameters: outputs.dataproc.v1.TemplateParameterResponse[];
    /**
     * WorkflowTemplate scheduling information.
     */
    readonly placement: outputs.dataproc.v1.WorkflowTemplatePlacementResponse;
    /**
     * The time template was last updated.
     */
    readonly updateTime: string;
    /**
     * Optional. Used to perform a consistent read-modify-write.This field should be left blank for a CreateWorkflowTemplate request. It is required for an UpdateWorkflowTemplate request, and must match the current server version. A typical update template flow would fetch the current template with a GetWorkflowTemplate request, which will return the current template with the version field filled in with the current server version. The user updates other fields in the template, then returns it as part of the UpdateWorkflowTemplate request.
     */
    readonly version: number;
}
/**
 * Retrieves the latest workflow template.Can retrieve previously instantiated template by specifying optional version parameter.
 */
export function getWorkflowTemplateOutput(args: GetWorkflowTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWorkflowTemplateResult> {
    return pulumi.output(args).apply((a: any) => getWorkflowTemplate(a, opts))
}

export interface GetWorkflowTemplateOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    version?: pulumi.Input<number>;
    workflowTemplateId: pulumi.Input<string>;
}
