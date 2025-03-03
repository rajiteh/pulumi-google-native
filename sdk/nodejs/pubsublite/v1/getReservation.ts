// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Returns the reservation configuration.
 */
export function getReservation(args: GetReservationArgs, opts?: pulumi.InvokeOptions): Promise<GetReservationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:pubsublite/v1:getReservation", {
        "location": args.location,
        "project": args.project,
        "reservationId": args.reservationId,
    }, opts);
}

export interface GetReservationArgs {
    location: string;
    project?: string;
    reservationId: string;
}

export interface GetReservationResult {
    /**
     * The name of the reservation. Structured like: projects/{project_number}/locations/{location}/reservations/{reservation_id}
     */
    readonly name: string;
    /**
     * The reserved throughput capacity. Every unit of throughput capacity is equivalent to 1 MiB/s of published messages or 2 MiB/s of subscribed messages. Any topics which are declared as using capacity from a Reservation will consume resources from this reservation instead of being charged individually.
     */
    readonly throughputCapacity: string;
}
/**
 * Returns the reservation configuration.
 */
export function getReservationOutput(args: GetReservationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReservationResult> {
    return pulumi.output(args).apply((a: any) => getReservation(a, opts))
}

export interface GetReservationOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    reservationId: pulumi.Input<string>;
}
