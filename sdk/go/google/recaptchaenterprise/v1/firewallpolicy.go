// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new FirewallPolicy, specifying conditions at which reCAPTCHA Enterprise actions can be executed. A project may have a maximum of 1000 policies.
type Firewallpolicy struct {
	pulumi.CustomResourceState

	// The actions that the caller should take regarding user access. There should be at most one terminal action. A terminal action is any action that forces a response, such as AllowAction, BlockAction or SubstituteAction. Zero or more non-terminal actions such as SetHeader might be specified. A single policy can contain up to 16 actions.
	Actions GoogleCloudRecaptchaenterpriseV1FirewallActionResponseArrayOutput `pulumi:"actions"`
	// A CEL (Common Expression Language) conditional expression that specifies if this policy applies to an incoming user request. If this condition evaluates to true and the requested path matched the path pattern, the associated actions should be executed by the caller. The condition string is checked for CEL syntax correctness on creation. For more information, see the [CEL spec](https://github.com/google/cel-spec) and its [language definition](https://github.com/google/cel-spec/blob/master/doc/langdef.md). A condition has a max length of 500 characters.
	Condition pulumi.StringOutput `pulumi:"condition"`
	// A description of what this policy aims to achieve, for convenience purposes. The description can at most include 256 UTF-8 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// The resource name for the FirewallPolicy in the format "projects/{project}/firewallpolicies/{firewallpolicy}".
	Name pulumi.StringOutput `pulumi:"name"`
	// The path for which this policy applies, specified as a glob pattern. For more information on glob, see the [manual page](https://man7.org/linux/man-pages/man7/glob.7.html). A path has a max length of 200 characters.
	Path    pulumi.StringOutput `pulumi:"path"`
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewFirewallpolicy registers a new resource with the given unique name, arguments, and options.
func NewFirewallpolicy(ctx *pulumi.Context,
	name string, args *FirewallpolicyArgs, opts ...pulumi.ResourceOption) (*Firewallpolicy, error) {
	if args == nil {
		args = &FirewallpolicyArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource Firewallpolicy
	err := ctx.RegisterResource("google-native:recaptchaenterprise/v1:Firewallpolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallpolicy gets an existing Firewallpolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallpolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallpolicyState, opts ...pulumi.ResourceOption) (*Firewallpolicy, error) {
	var resource Firewallpolicy
	err := ctx.ReadResource("google-native:recaptchaenterprise/v1:Firewallpolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Firewallpolicy resources.
type firewallpolicyState struct {
}

type FirewallpolicyState struct {
}

func (FirewallpolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallpolicyState)(nil)).Elem()
}

type firewallpolicyArgs struct {
	// The actions that the caller should take regarding user access. There should be at most one terminal action. A terminal action is any action that forces a response, such as AllowAction, BlockAction or SubstituteAction. Zero or more non-terminal actions such as SetHeader might be specified. A single policy can contain up to 16 actions.
	Actions []GoogleCloudRecaptchaenterpriseV1FirewallAction `pulumi:"actions"`
	// A CEL (Common Expression Language) conditional expression that specifies if this policy applies to an incoming user request. If this condition evaluates to true and the requested path matched the path pattern, the associated actions should be executed by the caller. The condition string is checked for CEL syntax correctness on creation. For more information, see the [CEL spec](https://github.com/google/cel-spec) and its [language definition](https://github.com/google/cel-spec/blob/master/doc/langdef.md). A condition has a max length of 500 characters.
	Condition *string `pulumi:"condition"`
	// A description of what this policy aims to achieve, for convenience purposes. The description can at most include 256 UTF-8 characters.
	Description *string `pulumi:"description"`
	// The resource name for the FirewallPolicy in the format "projects/{project}/firewallpolicies/{firewallpolicy}".
	Name *string `pulumi:"name"`
	// The path for which this policy applies, specified as a glob pattern. For more information on glob, see the [manual page](https://man7.org/linux/man-pages/man7/glob.7.html). A path has a max length of 200 characters.
	Path    *string `pulumi:"path"`
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Firewallpolicy resource.
type FirewallpolicyArgs struct {
	// The actions that the caller should take regarding user access. There should be at most one terminal action. A terminal action is any action that forces a response, such as AllowAction, BlockAction or SubstituteAction. Zero or more non-terminal actions such as SetHeader might be specified. A single policy can contain up to 16 actions.
	Actions GoogleCloudRecaptchaenterpriseV1FirewallActionArrayInput
	// A CEL (Common Expression Language) conditional expression that specifies if this policy applies to an incoming user request. If this condition evaluates to true and the requested path matched the path pattern, the associated actions should be executed by the caller. The condition string is checked for CEL syntax correctness on creation. For more information, see the [CEL spec](https://github.com/google/cel-spec) and its [language definition](https://github.com/google/cel-spec/blob/master/doc/langdef.md). A condition has a max length of 500 characters.
	Condition pulumi.StringPtrInput
	// A description of what this policy aims to achieve, for convenience purposes. The description can at most include 256 UTF-8 characters.
	Description pulumi.StringPtrInput
	// The resource name for the FirewallPolicy in the format "projects/{project}/firewallpolicies/{firewallpolicy}".
	Name pulumi.StringPtrInput
	// The path for which this policy applies, specified as a glob pattern. For more information on glob, see the [manual page](https://man7.org/linux/man-pages/man7/glob.7.html). A path has a max length of 200 characters.
	Path    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
}

func (FirewallpolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallpolicyArgs)(nil)).Elem()
}

type FirewallpolicyInput interface {
	pulumi.Input

	ToFirewallpolicyOutput() FirewallpolicyOutput
	ToFirewallpolicyOutputWithContext(ctx context.Context) FirewallpolicyOutput
}

func (*Firewallpolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**Firewallpolicy)(nil)).Elem()
}

func (i *Firewallpolicy) ToFirewallpolicyOutput() FirewallpolicyOutput {
	return i.ToFirewallpolicyOutputWithContext(context.Background())
}

func (i *Firewallpolicy) ToFirewallpolicyOutputWithContext(ctx context.Context) FirewallpolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallpolicyOutput)
}

type FirewallpolicyOutput struct{ *pulumi.OutputState }

func (FirewallpolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Firewallpolicy)(nil)).Elem()
}

func (o FirewallpolicyOutput) ToFirewallpolicyOutput() FirewallpolicyOutput {
	return o
}

func (o FirewallpolicyOutput) ToFirewallpolicyOutputWithContext(ctx context.Context) FirewallpolicyOutput {
	return o
}

// The actions that the caller should take regarding user access. There should be at most one terminal action. A terminal action is any action that forces a response, such as AllowAction, BlockAction or SubstituteAction. Zero or more non-terminal actions such as SetHeader might be specified. A single policy can contain up to 16 actions.
func (o FirewallpolicyOutput) Actions() GoogleCloudRecaptchaenterpriseV1FirewallActionResponseArrayOutput {
	return o.ApplyT(func(v *Firewallpolicy) GoogleCloudRecaptchaenterpriseV1FirewallActionResponseArrayOutput {
		return v.Actions
	}).(GoogleCloudRecaptchaenterpriseV1FirewallActionResponseArrayOutput)
}

// A CEL (Common Expression Language) conditional expression that specifies if this policy applies to an incoming user request. If this condition evaluates to true and the requested path matched the path pattern, the associated actions should be executed by the caller. The condition string is checked for CEL syntax correctness on creation. For more information, see the [CEL spec](https://github.com/google/cel-spec) and its [language definition](https://github.com/google/cel-spec/blob/master/doc/langdef.md). A condition has a max length of 500 characters.
func (o FirewallpolicyOutput) Condition() pulumi.StringOutput {
	return o.ApplyT(func(v *Firewallpolicy) pulumi.StringOutput { return v.Condition }).(pulumi.StringOutput)
}

// A description of what this policy aims to achieve, for convenience purposes. The description can at most include 256 UTF-8 characters.
func (o FirewallpolicyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Firewallpolicy) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The resource name for the FirewallPolicy in the format "projects/{project}/firewallpolicies/{firewallpolicy}".
func (o FirewallpolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Firewallpolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The path for which this policy applies, specified as a glob pattern. For more information on glob, see the [manual page](https://man7.org/linux/man-pages/man7/glob.7.html). A path has a max length of 200 characters.
func (o FirewallpolicyOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *Firewallpolicy) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

func (o FirewallpolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Firewallpolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallpolicyInput)(nil)).Elem(), &Firewallpolicy{})
	pulumi.RegisterOutputType(FirewallpolicyOutput{})
}
