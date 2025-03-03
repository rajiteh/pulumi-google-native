// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Updates an IAM policy for the specified bucket.
 * Note - this resource's API doesn't support deletion. When deleted, the resource will persist
 * on Google Cloud even though it will be deleted from Pulumi state.
 */
export class BucketIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing BucketIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BucketIamPolicy {
        return new BucketIamPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:storage/v1:BucketIamPolicy';

    /**
     * Returns true if the given object is an instance of BucketIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BucketIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BucketIamPolicy.__pulumiType;
    }

    /**
     * An association between a role, which comes with a set of permissions, and members who may assume that role.
     */
    public readonly bindings!: pulumi.Output<outputs.storage.v1.BucketIamPolicyBindingsItemResponse[]>;
    public readonly bucket!: pulumi.Output<string>;
    /**
     * HTTP 1.1  Entity tag for the policy.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * The kind of item this is. For policies, this is always storage#policy. This field is ignored on input.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The ID of the resource to which this policy belongs. Will be of the form projects/_/buckets/bucket for buckets, and projects/_/buckets/bucket/objects/object for objects. A specific generation may be specified by appending #generationNumber to the end of the object name, e.g. projects/_/buckets/my-bucket/objects/data.txt#17. The current generation can be denoted with #0. This field is ignored on input.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * The project to be billed for this request. Required for Requester Pays buckets.
     */
    public readonly userProject!: pulumi.Output<string | undefined>;
    /**
     * The IAM policy format version.
     */
    public readonly version!: pulumi.Output<number>;

    /**
     * Create a BucketIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketIamPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            resourceInputs["bindings"] = args ? args.bindings : undefined;
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["userProject"] = args ? args.userProject : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
        } else {
            resourceInputs["bindings"] = undefined /*out*/;
            resourceInputs["bucket"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["resourceId"] = undefined /*out*/;
            resourceInputs["userProject"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["bucket"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(BucketIamPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a BucketIamPolicy resource.
 */
export interface BucketIamPolicyArgs {
    /**
     * An association between a role, which comes with a set of permissions, and members who may assume that role.
     */
    bindings?: pulumi.Input<pulumi.Input<inputs.storage.v1.BucketIamPolicyBindingsItemArgs>[]>;
    bucket: pulumi.Input<string>;
    /**
     * HTTP 1.1  Entity tag for the policy.
     */
    etag?: pulumi.Input<string>;
    /**
     * The kind of item this is. For policies, this is always storage#policy. This field is ignored on input.
     */
    kind?: pulumi.Input<string>;
    /**
     * The ID of the resource to which this policy belongs. Will be of the form projects/_/buckets/bucket for buckets, and projects/_/buckets/bucket/objects/object for objects. A specific generation may be specified by appending #generationNumber to the end of the object name, e.g. projects/_/buckets/my-bucket/objects/data.txt#17. The current generation can be denoted with #0. This field is ignored on input.
     */
    resourceId?: pulumi.Input<string>;
    /**
     * The project to be billed for this request. Required for Requester Pays buckets.
     */
    userProject?: pulumi.Input<string>;
    /**
     * The IAM policy format version.
     */
    version?: pulumi.Input<number>;
}
