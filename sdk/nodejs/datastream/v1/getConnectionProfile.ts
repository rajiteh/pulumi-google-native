// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Use this method to get details about a connection profile.
 */
export function getConnectionProfile(args: GetConnectionProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetConnectionProfileResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:datastream/v1:getConnectionProfile", {
        "connectionProfileId": args.connectionProfileId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetConnectionProfileArgs {
    connectionProfileId: string;
    location: string;
    project?: string;
}

export interface GetConnectionProfileResult {
    /**
     * BigQuery Connection Profile configuration.
     */
    readonly bigqueryProfile: outputs.datastream.v1.BigQueryProfileResponse;
    /**
     * The create time of the resource.
     */
    readonly createTime: string;
    /**
     * Display name.
     */
    readonly displayName: string;
    /**
     * Forward SSH tunnel connectivity.
     */
    readonly forwardSshConnectivity: outputs.datastream.v1.ForwardSshTunnelConnectivityResponse;
    /**
     * Cloud Storage ConnectionProfile configuration.
     */
    readonly gcsProfile: outputs.datastream.v1.GcsProfileResponse;
    /**
     * Labels.
     */
    readonly labels: {[key: string]: string};
    /**
     * MySQL ConnectionProfile configuration.
     */
    readonly mysqlProfile: outputs.datastream.v1.MysqlProfileResponse;
    /**
     * The resource's name.
     */
    readonly name: string;
    /**
     * Oracle ConnectionProfile configuration.
     */
    readonly oracleProfile: outputs.datastream.v1.OracleProfileResponse;
    /**
     * PostgreSQL Connection Profile configuration.
     */
    readonly postgresqlProfile: outputs.datastream.v1.PostgresqlProfileResponse;
    /**
     * Private connectivity.
     */
    readonly privateConnectivity: outputs.datastream.v1.PrivateConnectivityResponse;
    /**
     * Static Service IP connectivity.
     */
    readonly staticServiceIpConnectivity: outputs.datastream.v1.StaticServiceIpConnectivityResponse;
    /**
     * The update time of the resource.
     */
    readonly updateTime: string;
}
/**
 * Use this method to get details about a connection profile.
 */
export function getConnectionProfileOutput(args: GetConnectionProfileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConnectionProfileResult> {
    return pulumi.output(args).apply((a: any) => getConnectionProfile(a, opts))
}

export interface GetConnectionProfileOutputArgs {
    connectionProfileId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
