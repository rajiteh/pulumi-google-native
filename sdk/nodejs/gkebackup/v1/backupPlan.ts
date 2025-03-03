// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new BackupPlan in a given location.
 * Auto-naming is currently not supported for this resource.
 */
export class BackupPlan extends pulumi.CustomResource {
    /**
     * Get an existing BackupPlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BackupPlan {
        return new BackupPlan(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:gkebackup/v1:BackupPlan';

    /**
     * Returns true if the given object is an instance of BackupPlan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackupPlan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackupPlan.__pulumiType;
    }

    /**
     * Defines the configuration of Backups created via this BackupPlan.
     */
    public readonly backupConfig!: pulumi.Output<outputs.gkebackup.v1.BackupConfigResponse>;
    /**
     * Required. The client-provided short name for the BackupPlan resource. This name must: - be between 1 and 63 characters long (inclusive) - consist of only lower-case ASCII letters, numbers, and dashes - start with a lower-case letter - end with a lower-case letter or number - be unique within the set of BackupPlans in this location
     */
    public readonly backupPlanId!: pulumi.Output<string>;
    /**
     * Defines a schedule for automatic Backup creation via this BackupPlan.
     */
    public readonly backupSchedule!: pulumi.Output<outputs.gkebackup.v1.ScheduleResponse>;
    /**
     * Immutable. The source cluster from which Backups will be created via this BackupPlan. Valid formats: - `projects/*&#47;locations/*&#47;clusters/*` - `projects/*&#47;zones/*&#47;clusters/*`
     */
    public readonly cluster!: pulumi.Output<string>;
    /**
     * The timestamp when this BackupPlan resource was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * This flag indicates whether this BackupPlan has been deactivated. Setting this field to True locks the BackupPlan such that no further updates will be allowed (except deletes), including the deactivated field itself. It also prevents any new Backups from being created via this BackupPlan (including scheduled Backups). Default: False
     */
    public readonly deactivated!: pulumi.Output<boolean>;
    /**
     * User specified descriptive string for this BackupPlan.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a backup plan from overwriting each other. It is strongly suggested that systems make use of the 'etag' in the read-modify-write cycle to perform BackupPlan updates in order to avoid race conditions: An `etag` is returned in the response to `GetBackupPlan`, and systems are expected to put that etag in the request to `UpdateBackupPlan` or `DeleteBackupPlan` to ensure that their change will be applied to the same version of the resource.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * A set of custom labels supplied by user.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * The full name of the BackupPlan resource. Format: `projects/*&#47;locations/*&#47;backupPlans/*`
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * The number of Kubernetes Pods backed up in the last successful Backup created via this BackupPlan.
     */
    public /*out*/ readonly protectedPodCount!: pulumi.Output<number>;
    /**
     * RetentionPolicy governs lifecycle of Backups created under this plan.
     */
    public readonly retentionPolicy!: pulumi.Output<outputs.gkebackup.v1.RetentionPolicyResponse>;
    /**
     * Server generated global unique identifier of [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier) format.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * The timestamp when this BackupPlan resource was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a BackupPlan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupPlanArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.backupPlanId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backupPlanId'");
            }
            if ((!args || args.cluster === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cluster'");
            }
            resourceInputs["backupConfig"] = args ? args.backupConfig : undefined;
            resourceInputs["backupPlanId"] = args ? args.backupPlanId : undefined;
            resourceInputs["backupSchedule"] = args ? args.backupSchedule : undefined;
            resourceInputs["cluster"] = args ? args.cluster : undefined;
            resourceInputs["deactivated"] = args ? args.deactivated : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["retentionPolicy"] = args ? args.retentionPolicy : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["protectedPodCount"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["backupConfig"] = undefined /*out*/;
            resourceInputs["backupPlanId"] = undefined /*out*/;
            resourceInputs["backupSchedule"] = undefined /*out*/;
            resourceInputs["cluster"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deactivated"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["protectedPodCount"] = undefined /*out*/;
            resourceInputs["retentionPolicy"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["backupPlanId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(BackupPlan.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a BackupPlan resource.
 */
export interface BackupPlanArgs {
    /**
     * Defines the configuration of Backups created via this BackupPlan.
     */
    backupConfig?: pulumi.Input<inputs.gkebackup.v1.BackupConfigArgs>;
    /**
     * Required. The client-provided short name for the BackupPlan resource. This name must: - be between 1 and 63 characters long (inclusive) - consist of only lower-case ASCII letters, numbers, and dashes - start with a lower-case letter - end with a lower-case letter or number - be unique within the set of BackupPlans in this location
     */
    backupPlanId: pulumi.Input<string>;
    /**
     * Defines a schedule for automatic Backup creation via this BackupPlan.
     */
    backupSchedule?: pulumi.Input<inputs.gkebackup.v1.ScheduleArgs>;
    /**
     * Immutable. The source cluster from which Backups will be created via this BackupPlan. Valid formats: - `projects/*&#47;locations/*&#47;clusters/*` - `projects/*&#47;zones/*&#47;clusters/*`
     */
    cluster: pulumi.Input<string>;
    /**
     * This flag indicates whether this BackupPlan has been deactivated. Setting this field to True locks the BackupPlan such that no further updates will be allowed (except deletes), including the deactivated field itself. It also prevents any new Backups from being created via this BackupPlan (including scheduled Backups). Default: False
     */
    deactivated?: pulumi.Input<boolean>;
    /**
     * User specified descriptive string for this BackupPlan.
     */
    description?: pulumi.Input<string>;
    /**
     * A set of custom labels supplied by user.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * RetentionPolicy governs lifecycle of Backups created under this plan.
     */
    retentionPolicy?: pulumi.Input<inputs.gkebackup.v1.RetentionPolicyArgs>;
}
