// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a metastore service in a project and location.
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:metastore/v1:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
     */
    public /*out*/ readonly artifactGcsUri!: pulumi.Output<string>;
    /**
     * The time when the metastore service was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Immutable. The database type that the Metastore service stores its data.
     */
    public readonly databaseType!: pulumi.Output<string>;
    /**
     * Immutable. Information used to configure the Dataproc Metastore service to encrypt customer data at rest. Cannot be updated.
     */
    public readonly encryptionConfig!: pulumi.Output<outputs.metastore.v1.EncryptionConfigResponse>;
    /**
     * The URI of the endpoint used to access the metastore service.
     */
    public /*out*/ readonly endpointUri!: pulumi.Output<string>;
    /**
     * Configuration information specific to running Hive metastore software as the metastore service.
     */
    public readonly hiveMetastoreConfig!: pulumi.Output<outputs.metastore.v1.HiveMetastoreConfigResponse>;
    /**
     * User-defined labels for the metastore service.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time. Maintenance window is not needed for services with the SPANNER database type.
     */
    public readonly maintenanceWindow!: pulumi.Output<outputs.metastore.v1.MaintenanceWindowResponse>;
    /**
     * The metadata management activities of the metastore service.
     */
    public /*out*/ readonly metadataManagementActivity!: pulumi.Output<outputs.metastore.v1.MetadataManagementActivityResponse>;
    /**
     * Immutable. The relative resource name of the metastore service, in the following format:projects/{project_number}/locations/{location_id}/services/{service_id}.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Immutable. The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:projects/{project_number}/global/networks/{network_id}.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * The configuration specifying the network settings for the Dataproc Metastore service.
     */
    public readonly networkConfig!: pulumi.Output<outputs.metastore.v1.NetworkConfigResponse>;
    /**
     * The TCP port at which the metastore service is reached. Default: 9083.
     */
    public readonly port!: pulumi.Output<number>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Immutable. The release channel of the service. If unspecified, defaults to STABLE.
     */
    public readonly releaseChannel!: pulumi.Output<string>;
    /**
     * Optional. A request ID. Specify a unique request ID to allow the server to ignore the request if it has completed. The server will ignore subsequent requests that provide a duplicate request ID for at least 60 minutes after the first request.For example, if an initial request times out, followed by another request with the same request ID, the server ignores the second request to prevent the creation of duplicate commitments.The request ID must be a valid UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier#Format) A zero UUID (00000000-0000-0000-0000-000000000000) is not supported.
     */
    public readonly requestId!: pulumi.Output<string | undefined>;
    /**
     * Scaling configuration of the metastore service.
     */
    public readonly scalingConfig!: pulumi.Output<outputs.metastore.v1.ScalingConfigResponse>;
    /**
     * Required. The ID of the metastore service, which is used as the final component of the metastore service's name.This value must be between 2 and 63 characters long inclusive, begin with a letter, end with a letter or number, and consist of alpha-numeric ASCII characters or hyphens.
     */
    public readonly serviceId!: pulumi.Output<string>;
    /**
     * The current state of the metastore service.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Additional information about the current state of the metastore service, if available.
     */
    public /*out*/ readonly stateMessage!: pulumi.Output<string>;
    /**
     * The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
     */
    public readonly telemetryConfig!: pulumi.Output<outputs.metastore.v1.TelemetryConfigResponse>;
    /**
     * The tier of the service.
     */
    public readonly tier!: pulumi.Output<string>;
    /**
     * The globally unique resource identifier of the metastore service.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * The time when the metastore service was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.serviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceId'");
            }
            resourceInputs["databaseType"] = args ? args.databaseType : undefined;
            resourceInputs["encryptionConfig"] = args ? args.encryptionConfig : undefined;
            resourceInputs["hiveMetastoreConfig"] = args ? args.hiveMetastoreConfig : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["maintenanceWindow"] = args ? args.maintenanceWindow : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["networkConfig"] = args ? args.networkConfig : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["releaseChannel"] = args ? args.releaseChannel : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["scalingConfig"] = args ? args.scalingConfig : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["telemetryConfig"] = args ? args.telemetryConfig : undefined;
            resourceInputs["tier"] = args ? args.tier : undefined;
            resourceInputs["artifactGcsUri"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["endpointUri"] = undefined /*out*/;
            resourceInputs["metadataManagementActivity"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateMessage"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["artifactGcsUri"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["databaseType"] = undefined /*out*/;
            resourceInputs["encryptionConfig"] = undefined /*out*/;
            resourceInputs["endpointUri"] = undefined /*out*/;
            resourceInputs["hiveMetastoreConfig"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["maintenanceWindow"] = undefined /*out*/;
            resourceInputs["metadataManagementActivity"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["network"] = undefined /*out*/;
            resourceInputs["networkConfig"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["releaseChannel"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["scalingConfig"] = undefined /*out*/;
            resourceInputs["serviceId"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateMessage"] = undefined /*out*/;
            resourceInputs["telemetryConfig"] = undefined /*out*/;
            resourceInputs["tier"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "project", "serviceId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Service.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * Immutable. The database type that the Metastore service stores its data.
     */
    databaseType?: pulumi.Input<enums.metastore.v1.ServiceDatabaseType>;
    /**
     * Immutable. Information used to configure the Dataproc Metastore service to encrypt customer data at rest. Cannot be updated.
     */
    encryptionConfig?: pulumi.Input<inputs.metastore.v1.EncryptionConfigArgs>;
    /**
     * Configuration information specific to running Hive metastore software as the metastore service.
     */
    hiveMetastoreConfig?: pulumi.Input<inputs.metastore.v1.HiveMetastoreConfigArgs>;
    /**
     * User-defined labels for the metastore service.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time. Maintenance window is not needed for services with the SPANNER database type.
     */
    maintenanceWindow?: pulumi.Input<inputs.metastore.v1.MaintenanceWindowArgs>;
    /**
     * Immutable. The relative resource name of the metastore service, in the following format:projects/{project_number}/locations/{location_id}/services/{service_id}.
     */
    name?: pulumi.Input<string>;
    /**
     * Immutable. The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:projects/{project_number}/global/networks/{network_id}.
     */
    network?: pulumi.Input<string>;
    /**
     * The configuration specifying the network settings for the Dataproc Metastore service.
     */
    networkConfig?: pulumi.Input<inputs.metastore.v1.NetworkConfigArgs>;
    /**
     * The TCP port at which the metastore service is reached. Default: 9083.
     */
    port?: pulumi.Input<number>;
    project?: pulumi.Input<string>;
    /**
     * Immutable. The release channel of the service. If unspecified, defaults to STABLE.
     */
    releaseChannel?: pulumi.Input<enums.metastore.v1.ServiceReleaseChannel>;
    /**
     * Optional. A request ID. Specify a unique request ID to allow the server to ignore the request if it has completed. The server will ignore subsequent requests that provide a duplicate request ID for at least 60 minutes after the first request.For example, if an initial request times out, followed by another request with the same request ID, the server ignores the second request to prevent the creation of duplicate commitments.The request ID must be a valid UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier#Format) A zero UUID (00000000-0000-0000-0000-000000000000) is not supported.
     */
    requestId?: pulumi.Input<string>;
    /**
     * Scaling configuration of the metastore service.
     */
    scalingConfig?: pulumi.Input<inputs.metastore.v1.ScalingConfigArgs>;
    /**
     * Required. The ID of the metastore service, which is used as the final component of the metastore service's name.This value must be between 2 and 63 characters long inclusive, begin with a letter, end with a letter or number, and consist of alpha-numeric ASCII characters or hyphens.
     */
    serviceId: pulumi.Input<string>;
    /**
     * The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
     */
    telemetryConfig?: pulumi.Input<inputs.metastore.v1.TelemetryConfigArgs>;
    /**
     * The tier of the service.
     */
    tier?: pulumi.Input<enums.metastore.v1.ServiceTier>;
}