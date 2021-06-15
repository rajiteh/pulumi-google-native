// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves autoscaling policy.
func LookupAutoscalingPolicy(ctx *pulumi.Context, args *LookupAutoscalingPolicyArgs, opts ...pulumi.InvokeOption) (*LookupAutoscalingPolicyResult, error) {
	var rv LookupAutoscalingPolicyResult
	err := ctx.Invoke("google-native:dataproc/v1beta2:getAutoscalingPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAutoscalingPolicyArgs struct {
	AutoscalingPolicyId string `pulumi:"autoscalingPolicyId"`
	Location            string `pulumi:"location"`
	Project             string `pulumi:"project"`
}

type LookupAutoscalingPolicyResult struct {
	BasicAlgorithm BasicAutoscalingAlgorithmResponse `pulumi:"basicAlgorithm"`
	// The "resource name" of the autoscaling policy, as described in https://cloud.google.com/apis/design/resource_names. For projects.regions.autoscalingPolicies, the resource name of the policy has the following format: projects/{project_id}/regions/{region}/autoscalingPolicies/{policy_id} For projects.locations.autoscalingPolicies, the resource name of the policy has the following format: projects/{project_id}/locations/{location}/autoscalingPolicies/{policy_id}
	Name string `pulumi:"name"`
	// Optional. Describes how the autoscaler will operate for secondary workers.
	SecondaryWorkerConfig InstanceGroupAutoscalingPolicyConfigResponse `pulumi:"secondaryWorkerConfig"`
	// Required. Describes how the autoscaler will operate for primary workers.
	WorkerConfig InstanceGroupAutoscalingPolicyConfigResponse `pulumi:"workerConfig"`
}