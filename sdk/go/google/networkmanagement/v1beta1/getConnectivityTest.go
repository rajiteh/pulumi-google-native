// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of a specific Connectivity Test.
func LookupConnectivityTest(ctx *pulumi.Context, args *LookupConnectivityTestArgs, opts ...pulumi.InvokeOption) (*LookupConnectivityTestResult, error) {
	var rv LookupConnectivityTestResult
	err := ctx.Invoke("google-native:networkmanagement/v1beta1:getConnectivityTest", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConnectivityTestArgs struct {
	ConnectivityTestId string  `pulumi:"connectivityTestId"`
	Project            *string `pulumi:"project"`
}

type LookupConnectivityTestResult struct {
	// The time the test was created.
	CreateTime string `pulumi:"createTime"`
	// The user-supplied description of the Connectivity Test. Maximum of 512 characters.
	Description string `pulumi:"description"`
	// Destination specification of the Connectivity Test. You can use a combination of destination IP address, Compute Engine VM instance, or VPC network to uniquely identify the destination location. Even if the destination IP address is not unique, the source IP location is unique. Usually, the analysis can infer the destination endpoint from route information. If the destination you specify is a VM instance and the instance has multiple network interfaces, then you must also specify either a destination IP address or VPC network to identify the destination interface. A reachability analysis proceeds even if the destination location is ambiguous. However, the result can include endpoints that you don't intend to test.
	Destination EndpointResponse `pulumi:"destination"`
	// The display name of a Connectivity Test.
	DisplayName string `pulumi:"displayName"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// Unique name of the resource using the form: `projects/{project_id}/locations/global/connectivityTests/{test}`
	Name string `pulumi:"name"`
	// The probing details of this test from the latest run, present for applicable tests only. The details are updated when creating a new test, updating an existing test, or triggering a one-time rerun of an existing test.
	ProbingDetails ProbingDetailsResponse `pulumi:"probingDetails"`
	// IP Protocol of the test. When not provided, "TCP" is assumed.
	Protocol string `pulumi:"protocol"`
	// The reachability details of this test from the latest run. The details are updated when creating a new test, updating an existing test, or triggering a one-time rerun of an existing test.
	ReachabilityDetails ReachabilityDetailsResponse `pulumi:"reachabilityDetails"`
	// Other projects that may be relevant for reachability analysis. This is applicable to scenarios where a test can cross project boundaries.
	RelatedProjects []string `pulumi:"relatedProjects"`
	// Source specification of the Connectivity Test. You can use a combination of source IP address, virtual machine (VM) instance, or Compute Engine network to uniquely identify the source location. Examples: If the source IP address is an internal IP address within a Google Cloud Virtual Private Cloud (VPC) network, then you must also specify the VPC network. Otherwise, specify the VM instance, which already contains its internal IP address and VPC network information. If the source of the test is within an on-premises network, then you must provide the destination VPC network. If the source endpoint is a Compute Engine VM instance with multiple network interfaces, the instance itself is not sufficient to identify the endpoint. So, you must also specify the source IP address or VPC network. A reachability analysis proceeds even if the source location is ambiguous. However, the test result may include endpoints that you don't intend to test.
	Source EndpointResponse `pulumi:"source"`
	// The time the test's configuration was updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupConnectivityTestOutput(ctx *pulumi.Context, args LookupConnectivityTestOutputArgs, opts ...pulumi.InvokeOption) LookupConnectivityTestResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConnectivityTestResult, error) {
			args := v.(LookupConnectivityTestArgs)
			r, err := LookupConnectivityTest(ctx, &args, opts...)
			var s LookupConnectivityTestResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConnectivityTestResultOutput)
}

type LookupConnectivityTestOutputArgs struct {
	ConnectivityTestId pulumi.StringInput    `pulumi:"connectivityTestId"`
	Project            pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupConnectivityTestOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectivityTestArgs)(nil)).Elem()
}

type LookupConnectivityTestResultOutput struct{ *pulumi.OutputState }

func (LookupConnectivityTestResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectivityTestResult)(nil)).Elem()
}

func (o LookupConnectivityTestResultOutput) ToLookupConnectivityTestResultOutput() LookupConnectivityTestResultOutput {
	return o
}

func (o LookupConnectivityTestResultOutput) ToLookupConnectivityTestResultOutputWithContext(ctx context.Context) LookupConnectivityTestResultOutput {
	return o
}

// The time the test was created.
func (o LookupConnectivityTestResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectivityTestResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The user-supplied description of the Connectivity Test. Maximum of 512 characters.
func (o LookupConnectivityTestResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectivityTestResult) string { return v.Description }).(pulumi.StringOutput)
}

// Destination specification of the Connectivity Test. You can use a combination of destination IP address, Compute Engine VM instance, or VPC network to uniquely identify the destination location. Even if the destination IP address is not unique, the source IP location is unique. Usually, the analysis can infer the destination endpoint from route information. If the destination you specify is a VM instance and the instance has multiple network interfaces, then you must also specify either a destination IP address or VPC network to identify the destination interface. A reachability analysis proceeds even if the destination location is ambiguous. However, the result can include endpoints that you don't intend to test.
func (o LookupConnectivityTestResultOutput) Destination() EndpointResponseOutput {
	return o.ApplyT(func(v LookupConnectivityTestResult) EndpointResponse { return v.Destination }).(EndpointResponseOutput)
}

// The display name of a Connectivity Test.
func (o LookupConnectivityTestResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectivityTestResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Resource labels to represent user-provided metadata.
func (o LookupConnectivityTestResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectivityTestResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Unique name of the resource using the form: `projects/{project_id}/locations/global/connectivityTests/{test}`
func (o LookupConnectivityTestResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectivityTestResult) string { return v.Name }).(pulumi.StringOutput)
}

// The probing details of this test from the latest run, present for applicable tests only. The details are updated when creating a new test, updating an existing test, or triggering a one-time rerun of an existing test.
func (o LookupConnectivityTestResultOutput) ProbingDetails() ProbingDetailsResponseOutput {
	return o.ApplyT(func(v LookupConnectivityTestResult) ProbingDetailsResponse { return v.ProbingDetails }).(ProbingDetailsResponseOutput)
}

// IP Protocol of the test. When not provided, "TCP" is assumed.
func (o LookupConnectivityTestResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectivityTestResult) string { return v.Protocol }).(pulumi.StringOutput)
}

// The reachability details of this test from the latest run. The details are updated when creating a new test, updating an existing test, or triggering a one-time rerun of an existing test.
func (o LookupConnectivityTestResultOutput) ReachabilityDetails() ReachabilityDetailsResponseOutput {
	return o.ApplyT(func(v LookupConnectivityTestResult) ReachabilityDetailsResponse { return v.ReachabilityDetails }).(ReachabilityDetailsResponseOutput)
}

// Other projects that may be relevant for reachability analysis. This is applicable to scenarios where a test can cross project boundaries.
func (o LookupConnectivityTestResultOutput) RelatedProjects() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupConnectivityTestResult) []string { return v.RelatedProjects }).(pulumi.StringArrayOutput)
}

// Source specification of the Connectivity Test. You can use a combination of source IP address, virtual machine (VM) instance, or Compute Engine network to uniquely identify the source location. Examples: If the source IP address is an internal IP address within a Google Cloud Virtual Private Cloud (VPC) network, then you must also specify the VPC network. Otherwise, specify the VM instance, which already contains its internal IP address and VPC network information. If the source of the test is within an on-premises network, then you must provide the destination VPC network. If the source endpoint is a Compute Engine VM instance with multiple network interfaces, the instance itself is not sufficient to identify the endpoint. So, you must also specify the source IP address or VPC network. A reachability analysis proceeds even if the source location is ambiguous. However, the test result may include endpoints that you don't intend to test.
func (o LookupConnectivityTestResultOutput) Source() EndpointResponseOutput {
	return o.ApplyT(func(v LookupConnectivityTestResult) EndpointResponse { return v.Source }).(EndpointResponseOutput)
}

// The time the test's configuration was updated.
func (o LookupConnectivityTestResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectivityTestResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectivityTestResultOutput{})
}
