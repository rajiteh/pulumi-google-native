// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new migration job in a given project and location.
 */
export class MigrationJob extends pulumi.CustomResource {
    /**
     * Get an existing MigrationJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): MigrationJob {
        return new MigrationJob(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:datamigration/v1beta1:MigrationJob';

    /**
     * Returns true if the given object is an instance of MigrationJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MigrationJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MigrationJob.__pulumiType;
    }

    /**
     * The timestamp when the migration job resource was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The resource name (URI) of the destination connection profile.
     */
    public readonly destination!: pulumi.Output<string>;
    /**
     * The database engine type and provider of the destination.
     */
    public readonly destinationDatabase!: pulumi.Output<outputs.datamigration.v1beta1.DatabaseTypeResponse>;
    /**
     * The migration job display name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The path to the dump file in Google Cloud Storage, in the format: (gs://[BUCKET_NAME]/[OBJECT_NAME]).
     */
    public readonly dumpPath!: pulumi.Output<string>;
    /**
     * The duration of the migration job (in seconds). A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
     */
    public /*out*/ readonly duration!: pulumi.Output<string>;
    /**
     * If the migration job is completed, the time when it was completed.
     */
    public /*out*/ readonly endTime!: pulumi.Output<string>;
    /**
     * The error details in case of state FAILED.
     */
    public /*out*/ readonly error!: pulumi.Output<outputs.datamigration.v1beta1.StatusResponse>;
    /**
     * The resource labels for migration job to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Required. The ID of the instance to create.
     */
    public readonly migrationJobId!: pulumi.Output<string>;
    /**
     * The name (URI) of this migration job resource, in the form of: projects/{project}/locations/{location}/migrationJobs/{migrationJob}.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The current migration job phase.
     */
    public /*out*/ readonly phase!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * A unique id used to identify the request. If the server receives two requests with the same id, then the second request will be ignored. It is recommended to always set this value to a UUID. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
     */
    public readonly requestId!: pulumi.Output<string | undefined>;
    /**
     * The details needed to communicate to the source over Reverse SSH tunnel connectivity.
     */
    public readonly reverseSshConnectivity!: pulumi.Output<outputs.datamigration.v1beta1.ReverseSshConnectivityResponse>;
    /**
     * The resource name (URI) of the source connection profile.
     */
    public readonly source!: pulumi.Output<string>;
    /**
     * The database engine type and provider of the source.
     */
    public readonly sourceDatabase!: pulumi.Output<outputs.datamigration.v1beta1.DatabaseTypeResponse>;
    /**
     * The current migration job state.
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * static ip connectivity data (default, no additional details needed).
     */
    public readonly staticIpConnectivity!: pulumi.Output<outputs.datamigration.v1beta1.StaticIpConnectivityResponse>;
    /**
     * The migration job type.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The timestamp when the migration job resource was last updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * The details of the VPC network that the source database is located in.
     */
    public readonly vpcPeeringConnectivity!: pulumi.Output<outputs.datamigration.v1beta1.VpcPeeringConnectivityResponse>;

    /**
     * Create a MigrationJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MigrationJobArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.destination === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destination'");
            }
            if ((!args || args.migrationJobId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'migrationJobId'");
            }
            if ((!args || args.source === undefined) && !opts.urn) {
                throw new Error("Missing required property 'source'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["destination"] = args ? args.destination : undefined;
            resourceInputs["destinationDatabase"] = args ? args.destinationDatabase : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["dumpPath"] = args ? args.dumpPath : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["migrationJobId"] = args ? args.migrationJobId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["reverseSshConnectivity"] = args ? args.reverseSshConnectivity : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["sourceDatabase"] = args ? args.sourceDatabase : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["staticIpConnectivity"] = args ? args.staticIpConnectivity : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vpcPeeringConnectivity"] = args ? args.vpcPeeringConnectivity : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["duration"] = undefined /*out*/;
            resourceInputs["endTime"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["phase"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["destination"] = undefined /*out*/;
            resourceInputs["destinationDatabase"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["dumpPath"] = undefined /*out*/;
            resourceInputs["duration"] = undefined /*out*/;
            resourceInputs["endTime"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["migrationJobId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["phase"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["reverseSshConnectivity"] = undefined /*out*/;
            resourceInputs["source"] = undefined /*out*/;
            resourceInputs["sourceDatabase"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["staticIpConnectivity"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["vpcPeeringConnectivity"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "migrationJobId", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(MigrationJob.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a MigrationJob resource.
 */
export interface MigrationJobArgs {
    /**
     * The resource name (URI) of the destination connection profile.
     */
    destination: pulumi.Input<string>;
    /**
     * The database engine type and provider of the destination.
     */
    destinationDatabase?: pulumi.Input<inputs.datamigration.v1beta1.DatabaseTypeArgs>;
    /**
     * The migration job display name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The path to the dump file in Google Cloud Storage, in the format: (gs://[BUCKET_NAME]/[OBJECT_NAME]).
     */
    dumpPath?: pulumi.Input<string>;
    /**
     * The resource labels for migration job to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Required. The ID of the instance to create.
     */
    migrationJobId: pulumi.Input<string>;
    /**
     * The name (URI) of this migration job resource, in the form of: projects/{project}/locations/{location}/migrationJobs/{migrationJob}.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * A unique id used to identify the request. If the server receives two requests with the same id, then the second request will be ignored. It is recommended to always set this value to a UUID. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
     */
    requestId?: pulumi.Input<string>;
    /**
     * The details needed to communicate to the source over Reverse SSH tunnel connectivity.
     */
    reverseSshConnectivity?: pulumi.Input<inputs.datamigration.v1beta1.ReverseSshConnectivityArgs>;
    /**
     * The resource name (URI) of the source connection profile.
     */
    source: pulumi.Input<string>;
    /**
     * The database engine type and provider of the source.
     */
    sourceDatabase?: pulumi.Input<inputs.datamigration.v1beta1.DatabaseTypeArgs>;
    /**
     * The current migration job state.
     */
    state?: pulumi.Input<enums.datamigration.v1beta1.MigrationJobState>;
    /**
     * static ip connectivity data (default, no additional details needed).
     */
    staticIpConnectivity?: pulumi.Input<inputs.datamigration.v1beta1.StaticIpConnectivityArgs>;
    /**
     * The migration job type.
     */
    type: pulumi.Input<enums.datamigration.v1beta1.MigrationJobType>;
    /**
     * The details of the VPC network that the source database is located in.
     */
    vpcPeeringConnectivity?: pulumi.Input<inputs.datamigration.v1beta1.VpcPeeringConnectivityArgs>;
}
