// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets a PerfSampleSeries. May return any of the following error code(s): - NOT_FOUND - The specified PerfSampleSeries does not exist
 */
export function getPerfSampleSeries(args: GetPerfSampleSeriesArgs, opts?: pulumi.InvokeOptions): Promise<GetPerfSampleSeriesResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:toolresults/v1beta3:getPerfSampleSeries", {
        "executionId": args.executionId,
        "historyId": args.historyId,
        "project": args.project,
        "sampleSeriesId": args.sampleSeriesId,
        "stepId": args.stepId,
    }, opts);
}

export interface GetPerfSampleSeriesArgs {
    executionId: string;
    historyId: string;
    project?: string;
    sampleSeriesId: string;
    stepId: string;
}

export interface GetPerfSampleSeriesResult {
    /**
     * Basic series represented by a line chart
     */
    readonly basicPerfSampleSeries: outputs.toolresults.v1beta3.BasicPerfSampleSeriesResponse;
    /**
     * A tool results execution ID. 
     */
    readonly executionId: string;
    /**
     * A tool results history ID. 
     */
    readonly historyId: string;
    /**
     * The cloud project 
     */
    readonly project: string;
    /**
     * A sample series id 
     */
    readonly sampleSeriesId: string;
    /**
     * A tool results step ID. 
     */
    readonly stepId: string;
}
/**
 * Gets a PerfSampleSeries. May return any of the following error code(s): - NOT_FOUND - The specified PerfSampleSeries does not exist
 */
export function getPerfSampleSeriesOutput(args: GetPerfSampleSeriesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPerfSampleSeriesResult> {
    return pulumi.output(args).apply((a: any) => getPerfSampleSeries(a, opts))
}

export interface GetPerfSampleSeriesOutputArgs {
    executionId: pulumi.Input<string>;
    historyId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    sampleSeriesId: pulumi.Input<string>;
    stepId: pulumi.Input<string>;
}
