// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets a job.
 */
export function getJob(args: GetJobArgs, opts?: pulumi.InvokeOptions): Promise<GetJobResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:cloudscheduler/v1beta1:getJob", {
        "jobId": args.jobId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetJobArgs {
    jobId: string;
    location: string;
    project?: string;
}

export interface GetJobResult {
    /**
     * App Engine HTTP target.
     */
    readonly appEngineHttpTarget: outputs.cloudscheduler.v1beta1.AppEngineHttpTargetResponse;
    /**
     * The deadline for job attempts. If the request handler does not respond by this deadline then the request is cancelled and the attempt is marked as a `DEADLINE_EXCEEDED` failure. The failed attempt can be viewed in execution logs. Cloud Scheduler will retry the job according to the RetryConfig. The default and the allowed values depend on the type of target: * For HTTP targets, the default is 3 minutes. The deadline must be in the interval [15 seconds, 30 minutes]. * For App Engine HTTP targets, 0 indicates that the request has the default deadline. The default deadline depends on the scaling type of the service: 10 minutes for standard apps with automatic scaling, 24 hours for standard apps with manual and basic scaling, and 60 minutes for flex apps. If the request deadline is set, it must be in the interval [15 seconds, 24 hours 15 seconds]. * For Pub/Sub targets, this field is ignored.
     */
    readonly attemptDeadline: string;
    /**
     * Optionally caller-specified in CreateJob or UpdateJob. A human-readable description for the job. This string must not contain more than 500 characters.
     */
    readonly description: string;
    /**
     * HTTP target.
     */
    readonly httpTarget: outputs.cloudscheduler.v1beta1.HttpTargetResponse;
    /**
     * The time the last job attempt started.
     */
    readonly lastAttemptTime: string;
    /**
     * Immutable. This field is used to manage the legacy App Engine Cron jobs using the Cloud Scheduler API. If the field is set to true, the job will be considered a legacy job. Note that App Engine Cron jobs have fewer features than Cloud Scheduler jobs, e.g., are only limited to App Engine targets.
     */
    readonly legacyAppEngineCron: boolean;
    /**
     * Optionally caller-specified in CreateJob, after which it becomes output only. The job name. For example: `projects/PROJECT_ID/locations/LOCATION_ID/jobs/JOB_ID`. * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), or periods (.). For more information, see [Identifying projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects) * `LOCATION_ID` is the canonical ID for the job's location. The list of available locations can be obtained by calling ListLocations. For more information, see https://cloud.google.com/about/locations/. * `JOB_ID` can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), or underscores (_). The maximum length is 500 characters.
     */
    readonly name: string;
    /**
     * Pub/Sub target.
     */
    readonly pubsubTarget: outputs.cloudscheduler.v1beta1.PubsubTargetResponse;
    /**
     * Settings that determine the retry behavior.
     */
    readonly retryConfig: outputs.cloudscheduler.v1beta1.RetryConfigResponse;
    /**
     * Required, except when used with UpdateJob. Describes the schedule on which the job will be executed. The schedule can be either of the following types: * [Crontab](https://en.wikipedia.org/wiki/Cron#Overview) * English-like [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules) As a general rule, execution `n + 1` of a job will not begin until execution `n` has finished. Cloud Scheduler will never allow two simultaneously outstanding executions. For example, this implies that if the `n+1`th execution is scheduled to run at 16:00 but the `n`th execution takes until 16:15, the `n+1`th execution will not start until `16:15`. A scheduled start time will be delayed if the previous execution has not ended when its scheduled time occurs. If retry_count > 0 and a job attempt fails, the job will be tried a total of retry_count times, with exponential backoff, until the next scheduled start time.
     */
    readonly schedule: string;
    /**
     * The next time the job is scheduled. Note that this may be a retry of a previously failed attempt or the next execution time according to the schedule.
     */
    readonly scheduleTime: string;
    /**
     * State of the job.
     */
    readonly state: string;
    /**
     * The response from the target for the last attempted execution.
     */
    readonly status: outputs.cloudscheduler.v1beta1.StatusResponse;
    /**
     * Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the [tz database](http://en.wikipedia.org/wiki/Tz_database). Note that some time zones include a provision for daylight savings time. The rules for daylight saving time are determined by the chosen tz. For UTC use the string "utc". If a time zone is not specified, the default will be in UTC (also known as GMT).
     */
    readonly timeZone: string;
    /**
     * The creation time of the job.
     */
    readonly userUpdateTime: string;
}
/**
 * Gets a job.
 */
export function getJobOutput(args: GetJobOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetJobResult> {
    return pulumi.output(args).apply((a: any) => getJob(a, opts))
}

export interface GetJobOutputArgs {
    jobId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
