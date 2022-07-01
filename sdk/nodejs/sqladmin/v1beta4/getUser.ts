// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves a resource containing information about a user.
 */
export function getUser(args: GetUserArgs, opts?: pulumi.InvokeOptions): Promise<GetUserResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:sqladmin/v1beta4:getUser", {
        "instance": args.instance,
        "name": args.name,
        "project": args.project,
    }, opts);
}

export interface GetUserArgs {
    instance: string;
    name: string;
    project?: string;
}

export interface GetUserResult {
    /**
     * Dual password status for the user.
     */
    readonly dualPasswordType: string;
    /**
     * This field is deprecated and will be removed from a future version of the API.
     *
     * @deprecated This field is deprecated and will be removed from a future version of the API.
     */
    readonly etag: string;
    /**
     * Optional. The host from which the user can connect. For `insert` operations, host defaults to an empty string. For `update` operations, host is specified as part of the request URL. The host name cannot be updated after insertion. For a MySQL instance, it's required; for a PostgreSQL or SQL Server instance, it's optional.
     */
    readonly host: string;
    /**
     * The name of the Cloud SQL instance. This does not include the project ID. Can be omitted for *update* because it is already specified on the URL.
     */
    readonly instance: string;
    /**
     * This is always `sql#user`.
     */
    readonly kind: string;
    /**
     * The name of the user in the Cloud SQL instance. Can be omitted for `update` because it is already specified in the URL.
     */
    readonly name: string;
    /**
     * The password for the user.
     */
    readonly password: string;
    /**
     * User level password validation policy.
     */
    readonly passwordPolicy: outputs.sqladmin.v1beta4.UserPasswordValidationPolicyResponse;
    /**
     * The project ID of the project containing the Cloud SQL database. The Google apps domain is prefixed if applicable. Can be omitted for *update* because it is already specified on the URL.
     */
    readonly project: string;
    readonly sqlserverUserDetails: outputs.sqladmin.v1beta4.SqlServerUserDetailsResponse;
    /**
     * The user type. It determines the method to authenticate the user during login. The default is the database's built-in user type.
     */
    readonly type: string;
}

export function getUserOutput(args: GetUserOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserResult> {
    return pulumi.output(args).apply(a => getUser(a, opts))
}

export interface GetUserOutputArgs {
    instance: pulumi.Input<string>;
    name: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}