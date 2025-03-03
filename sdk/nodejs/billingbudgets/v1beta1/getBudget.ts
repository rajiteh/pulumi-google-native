// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns a budget. WARNING: There are some fields exposed on the Google Cloud Console that aren't available on this API. When reading from the API, you will not see these fields in the return value, though they may have been set in the Cloud Console.
 */
export function getBudget(args: GetBudgetArgs, opts?: pulumi.InvokeOptions): Promise<GetBudgetResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:billingbudgets/v1beta1:getBudget", {
        "billingAccountId": args.billingAccountId,
        "budgetId": args.budgetId,
    }, opts);
}

export interface GetBudgetArgs {
    billingAccountId: string;
    budgetId: string;
}

export interface GetBudgetResult {
    /**
     * Optional. Rules to apply to notifications sent based on budget spend and thresholds.
     */
    readonly allUpdatesRule: outputs.billingbudgets.v1beta1.GoogleCloudBillingBudgetsV1beta1AllUpdatesRuleResponse;
    /**
     * Budgeted amount.
     */
    readonly amount: outputs.billingbudgets.v1beta1.GoogleCloudBillingBudgetsV1beta1BudgetAmountResponse;
    /**
     * Optional. Filters that define which resources are used to compute the actual spend against the budget amount, such as projects, services, and the budget's time period, as well as other filters.
     */
    readonly budgetFilter: outputs.billingbudgets.v1beta1.GoogleCloudBillingBudgetsV1beta1FilterResponse;
    /**
     * User data for display name in UI. Validation: <= 60 chars.
     */
    readonly displayName: string;
    /**
     * Optional. Etag to validate that the object is unchanged for a read-modify-write operation. An empty etag will cause an update to overwrite other changes.
     */
    readonly etag: string;
    /**
     * Resource name of the budget. The resource name implies the scope of a budget. Values are of the form `billingAccounts/{billingAccountId}/budgets/{budgetId}`.
     */
    readonly name: string;
    /**
     * Optional. Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of the budget. Optional for `pubsubTopic` notifications. Required if using email notifications.
     */
    readonly thresholdRules: outputs.billingbudgets.v1beta1.GoogleCloudBillingBudgetsV1beta1ThresholdRuleResponse[];
}
/**
 * Returns a budget. WARNING: There are some fields exposed on the Google Cloud Console that aren't available on this API. When reading from the API, you will not see these fields in the return value, though they may have been set in the Cloud Console.
 */
export function getBudgetOutput(args: GetBudgetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBudgetResult> {
    return pulumi.output(args).apply((a: any) => getBudget(a, opts))
}

export interface GetBudgetOutputArgs {
    billingAccountId: pulumi.Input<string>;
    budgetId: pulumi.Input<string>;
}
