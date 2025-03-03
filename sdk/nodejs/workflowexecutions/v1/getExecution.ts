// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns an execution of the given name.
 */
export function getExecution(args: GetExecutionArgs, opts?: pulumi.InvokeOptions): Promise<GetExecutionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:workflowexecutions/v1:getExecution", {
        "executionId": args.executionId,
        "location": args.location,
        "project": args.project,
        "view": args.view,
        "workflowId": args.workflowId,
    }, opts);
}

export interface GetExecutionArgs {
    executionId: string;
    location: string;
    project?: string;
    view?: string;
    workflowId: string;
}

export interface GetExecutionResult {
    /**
     * Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\"firstName\":\"FIRST\",\"lastName\":\"LAST\"}"}'`
     */
    readonly argument: string;
    /**
     * The call logging level associated to this execution.
     */
    readonly callLogLevel: string;
    /**
     * Measures the duration of the execution.
     */
    readonly duration: string;
    /**
     * Marks the end of execution, successful or not.
     */
    readonly endTime: string;
    /**
     * The error which caused the execution to finish prematurely. The value is only present if the execution's state is `FAILED` or `CANCELLED`.
     */
    readonly error: outputs.workflowexecutions.v1.ErrorResponse;
    /**
     * Labels associated with this execution. Labels can contain at most 64 entries. Keys and values can be no longer than 63 characters and can only contain lowercase letters, numeric characters, underscores, and dashes. Label keys must start with a letter. International characters are allowed. By default, labels are inherited from the workflow but are overridden by any labels associated with the execution.
     */
    readonly labels: {[key: string]: string};
    /**
     * The resource name of the execution. Format: projects/{project}/locations/{location}/workflows/{workflow}/executions/{execution}
     */
    readonly name: string;
    /**
     * Output of the execution represented as a JSON string. The value can only be present if the execution's state is `SUCCEEDED`.
     */
    readonly result: string;
    /**
     * Marks the beginning of execution.
     */
    readonly startTime: string;
    /**
     * Current state of the execution.
     */
    readonly state: string;
    /**
     * Error regarding the state of the Execution resource. For example, this field will have error details if the Execution data is unavailable due to revoked KMS key permissions.
     */
    readonly stateError: outputs.workflowexecutions.v1.StateErrorResponse;
    /**
     * Status tracks the current steps and progress data of this execution.
     */
    readonly status: outputs.workflowexecutions.v1.StatusResponse;
    /**
     * Revision of the workflow this execution is using.
     */
    readonly workflowRevisionId: string;
}
/**
 * Returns an execution of the given name.
 */
export function getExecutionOutput(args: GetExecutionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetExecutionResult> {
    return pulumi.output(args).apply((a: any) => getExecution(a, opts))
}

export interface GetExecutionOutputArgs {
    executionId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    view?: pulumi.Input<string>;
    workflowId: pulumi.Input<string>;
}
