// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets information about a billing account. The current authenticated user must be a [viewer of the billing account](https://cloud.google.com/billing/docs/how-to/billing-access).
 */
export function getBillingAccount(args: GetBillingAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetBillingAccountResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:cloudbilling/v1:getBillingAccount", {
        "billingAccountId": args.billingAccountId,
    }, opts);
}

export interface GetBillingAccountArgs {
    billingAccountId: string;
}

export interface GetBillingAccountResult {
    /**
     * The display name given to the billing account, such as `My Billing Account`. This name is displayed in the Google Cloud Console.
     */
    readonly displayName: string;
    /**
     * If this account is a [subaccount](https://cloud.google.com/billing/docs/concepts), then this will be the resource name of the parent billing account that it is being resold through. Otherwise this will be empty.
     */
    readonly masterBillingAccount: string;
    /**
     * The resource name of the billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, `billingAccounts/012345-567890-ABCDEF` would be the resource name for billing account `012345-567890-ABCDEF`.
     */
    readonly name: string;
    /**
     * True if the billing account is open, and will therefore be charged for any usage on associated projects. False if the billing account is closed, and therefore projects associated with it will be unable to use paid services.
     */
    readonly open: boolean;
}
/**
 * Gets information about a billing account. The current authenticated user must be a [viewer of the billing account](https://cloud.google.com/billing/docs/how-to/billing-access).
 */
export function getBillingAccountOutput(args: GetBillingAccountOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBillingAccountResult> {
    return pulumi.output(args).apply((a: any) => getBillingAccount(a, opts))
}

export interface GetBillingAccountOutputArgs {
    billingAccountId: pulumi.Input<string>;
}
