// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns the specified HealthCheck resource.
 */
export function getRegionHealthCheck(args: GetRegionHealthCheckArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionHealthCheckResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/beta:getRegionHealthCheck", {
        "healthCheck": args.healthCheck,
        "project": args.project,
        "region": args.region,
    }, opts);
}

export interface GetRegionHealthCheckArgs {
    healthCheck: string;
    project?: string;
    region: string;
}

export interface GetRegionHealthCheckResult {
    /**
     * How often (in seconds) to send a health check. The default value is 5 seconds.
     */
    readonly checkIntervalSec: number;
    /**
     * Creation timestamp in 3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    readonly grpcHealthCheck: outputs.compute.beta.GRPCHealthCheckResponse;
    /**
     * A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
     */
    readonly healthyThreshold: number;
    readonly http2HealthCheck: outputs.compute.beta.HTTP2HealthCheckResponse;
    readonly httpHealthCheck: outputs.compute.beta.HTTPHealthCheckResponse;
    readonly httpsHealthCheck: outputs.compute.beta.HTTPSHealthCheckResponse;
    /**
     * Type of the resource.
     */
    readonly kind: string;
    /**
     * Configure logging on this health check.
     */
    readonly logConfig: outputs.compute.beta.HealthCheckLogConfigResponse;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. For example, a name that is 1-63 characters long, matches the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`, and otherwise complies with RFC1035. This regular expression describes a name where the first character is a lowercase letter, and all following characters are a dash, lowercase letter, or digit, except the last character, which isn't a dash.
     */
    readonly name: string;
    /**
     * Region where the health check resides. Not applicable to global health checks.
     */
    readonly region: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    readonly sslHealthCheck: outputs.compute.beta.SSLHealthCheckResponse;
    readonly tcpHealthCheck: outputs.compute.beta.TCPHealthCheckResponse;
    /**
     * How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.
     */
    readonly timeoutSec: number;
    /**
     * Specifies the type of the healthCheck, either TCP, SSL, HTTP, HTTPS, HTTP2 or GRPC. Exactly one of the protocol-specific health check fields must be specified, which must match type field.
     */
    readonly type: string;
    /**
     * A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
     */
    readonly unhealthyThreshold: number;
}
/**
 * Returns the specified HealthCheck resource.
 */
export function getRegionHealthCheckOutput(args: GetRegionHealthCheckOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegionHealthCheckResult> {
    return pulumi.output(args).apply((a: any) => getRegionHealthCheck(a, opts))
}

export interface GetRegionHealthCheckOutputArgs {
    healthCheck: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
}
