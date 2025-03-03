// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Create a set of phrase hints. Each item in the set can be a single word or a multi-word phrase. The items in the PhraseSet are favored by the recognition model when you send a call that includes the PhraseSet.
 */
export class PhraseSet extends pulumi.CustomResource {
    /**
     * Get an existing PhraseSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PhraseSet {
        return new PhraseSet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:speech/v1:PhraseSet';

    /**
     * Returns true if the given object is an instance of PhraseSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PhraseSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PhraseSet.__pulumiType;
    }

    /**
     * Hint Boost. Positive value will increase the probability that a specific phrase will be recognized over other similar sounding phrases. The higher the boost, the higher the chance of false positive recognition as well. Negative boost values would correspond to anti-biasing. Anti-biasing is not enabled, so negative boost will simply be ignored. Though `boost` can accept a wide range of positive values, most use cases are best served with values between 0 (exclusive) and 20. We recommend using a binary search approach to finding the optimal value for your use case as well as adding phrases both with and without boost to your requests.
     */
    public readonly boost!: pulumi.Output<number>;
    /**
     * The [KMS key name](https://cloud.google.com/kms/docs/resource-hierarchy#keys) with which the content of the PhraseSet is encrypted. The expected format is `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
     */
    public /*out*/ readonly kmsKeyName!: pulumi.Output<string>;
    /**
     * The [KMS key version name](https://cloud.google.com/kms/docs/resource-hierarchy#key_versions) with which content of the PhraseSet is encrypted. The expected format is `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}/cryptoKeyVersions/{crypto_key_version}`.
     */
    public /*out*/ readonly kmsKeyVersionName!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * The resource name of the phrase set.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of word and phrases.
     */
    public readonly phrases!: pulumi.Output<outputs.speech.v1.PhraseResponse[]>;
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a PhraseSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PhraseSetArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.phraseSetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'phraseSetId'");
            }
            resourceInputs["boost"] = args ? args.boost : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["phraseSetId"] = args ? args.phraseSetId : undefined;
            resourceInputs["phrases"] = args ? args.phrases : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["kmsKeyName"] = undefined /*out*/;
            resourceInputs["kmsKeyVersionName"] = undefined /*out*/;
        } else {
            resourceInputs["boost"] = undefined /*out*/;
            resourceInputs["kmsKeyName"] = undefined /*out*/;
            resourceInputs["kmsKeyVersionName"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["phrases"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(PhraseSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PhraseSet resource.
 */
export interface PhraseSetArgs {
    /**
     * Hint Boost. Positive value will increase the probability that a specific phrase will be recognized over other similar sounding phrases. The higher the boost, the higher the chance of false positive recognition as well. Negative boost values would correspond to anti-biasing. Anti-biasing is not enabled, so negative boost will simply be ignored. Though `boost` can accept a wide range of positive values, most use cases are best served with values between 0 (exclusive) and 20. We recommend using a binary search approach to finding the optimal value for your use case as well as adding phrases both with and without boost to your requests.
     */
    boost?: pulumi.Input<number>;
    location?: pulumi.Input<string>;
    /**
     * The resource name of the phrase set.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID to use for the phrase set, which will become the final component of the phrase set's resource name. This value should restrict to letters, numbers, and hyphens, with the first character a letter, the last a letter or a number, and be 4-63 characters.
     */
    phraseSetId: pulumi.Input<string>;
    /**
     * A list of word and phrases.
     */
    phrases?: pulumi.Input<pulumi.Input<inputs.speech.v1.PhraseArgs>[]>;
    project?: pulumi.Input<string>;
}
