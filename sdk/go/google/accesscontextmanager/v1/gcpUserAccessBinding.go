// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a GcpUserAccessBinding. If the client specifies a name, the server ignores it. Fails if a resource already exists with the same group_key. Completion of this long-running operation does not necessarily signify that the new binding is deployed onto all affected users, which may take more time.
type GcpUserAccessBinding struct {
	pulumi.CustomResourceState

	// Optional. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	AccessLevels pulumi.StringArrayOutput `pulumi:"accessLevels"`
	// Optional. Dry run access level that will be evaluated but will not be enforced. The access denial based on dry run policy will be logged. Only one access level is supported, not multiple. This list must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	DryRunAccessLevels pulumi.StringArrayOutput `pulumi:"dryRunAccessLevels"`
	// Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the [G Suite Directory API's Groups resource] (https://developers.google.com/admin-sdk/directory/v1/reference/groups#resource). If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht"
	GroupKey pulumi.StringOutput `pulumi:"groupKey"`
	// Immutable. Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved characters (as defined by [RFC 3986 Section 2.3](https://tools.ietf.org/html/rfc3986#section-2.3)). Should not be specified by the client during creation. Example: "organizations/256/gcpUserAccessBindings/b3-BhcX_Ud5N"
	Name           pulumi.StringOutput `pulumi:"name"`
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
}

// NewGcpUserAccessBinding registers a new resource with the given unique name, arguments, and options.
func NewGcpUserAccessBinding(ctx *pulumi.Context,
	name string, args *GcpUserAccessBindingArgs, opts ...pulumi.ResourceOption) (*GcpUserAccessBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupKey == nil {
		return nil, errors.New("invalid value for required argument 'GroupKey'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"organizationId",
	})
	opts = append(opts, replaceOnChanges)
	var resource GcpUserAccessBinding
	err := ctx.RegisterResource("google-native:accesscontextmanager/v1:GcpUserAccessBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGcpUserAccessBinding gets an existing GcpUserAccessBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGcpUserAccessBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GcpUserAccessBindingState, opts ...pulumi.ResourceOption) (*GcpUserAccessBinding, error) {
	var resource GcpUserAccessBinding
	err := ctx.ReadResource("google-native:accesscontextmanager/v1:GcpUserAccessBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GcpUserAccessBinding resources.
type gcpUserAccessBindingState struct {
}

type GcpUserAccessBindingState struct {
}

func (GcpUserAccessBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*gcpUserAccessBindingState)(nil)).Elem()
}

type gcpUserAccessBindingArgs struct {
	// Optional. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	AccessLevels []string `pulumi:"accessLevels"`
	// Optional. Dry run access level that will be evaluated but will not be enforced. The access denial based on dry run policy will be logged. Only one access level is supported, not multiple. This list must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	DryRunAccessLevels []string `pulumi:"dryRunAccessLevels"`
	// Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the [G Suite Directory API's Groups resource] (https://developers.google.com/admin-sdk/directory/v1/reference/groups#resource). If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht"
	GroupKey string `pulumi:"groupKey"`
	// Immutable. Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved characters (as defined by [RFC 3986 Section 2.3](https://tools.ietf.org/html/rfc3986#section-2.3)). Should not be specified by the client during creation. Example: "organizations/256/gcpUserAccessBindings/b3-BhcX_Ud5N"
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
}

// The set of arguments for constructing a GcpUserAccessBinding resource.
type GcpUserAccessBindingArgs struct {
	// Optional. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	AccessLevels pulumi.StringArrayInput
	// Optional. Dry run access level that will be evaluated but will not be enforced. The access denial based on dry run policy will be logged. Only one access level is supported, not multiple. This list must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	DryRunAccessLevels pulumi.StringArrayInput
	// Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the [G Suite Directory API's Groups resource] (https://developers.google.com/admin-sdk/directory/v1/reference/groups#resource). If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht"
	GroupKey pulumi.StringInput
	// Immutable. Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved characters (as defined by [RFC 3986 Section 2.3](https://tools.ietf.org/html/rfc3986#section-2.3)). Should not be specified by the client during creation. Example: "organizations/256/gcpUserAccessBindings/b3-BhcX_Ud5N"
	Name           pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
}

func (GcpUserAccessBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gcpUserAccessBindingArgs)(nil)).Elem()
}

type GcpUserAccessBindingInput interface {
	pulumi.Input

	ToGcpUserAccessBindingOutput() GcpUserAccessBindingOutput
	ToGcpUserAccessBindingOutputWithContext(ctx context.Context) GcpUserAccessBindingOutput
}

func (*GcpUserAccessBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**GcpUserAccessBinding)(nil)).Elem()
}

func (i *GcpUserAccessBinding) ToGcpUserAccessBindingOutput() GcpUserAccessBindingOutput {
	return i.ToGcpUserAccessBindingOutputWithContext(context.Background())
}

func (i *GcpUserAccessBinding) ToGcpUserAccessBindingOutputWithContext(ctx context.Context) GcpUserAccessBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GcpUserAccessBindingOutput)
}

type GcpUserAccessBindingOutput struct{ *pulumi.OutputState }

func (GcpUserAccessBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GcpUserAccessBinding)(nil)).Elem()
}

func (o GcpUserAccessBindingOutput) ToGcpUserAccessBindingOutput() GcpUserAccessBindingOutput {
	return o
}

func (o GcpUserAccessBindingOutput) ToGcpUserAccessBindingOutputWithContext(ctx context.Context) GcpUserAccessBindingOutput {
	return o
}

// Optional. Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
func (o GcpUserAccessBindingOutput) AccessLevels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GcpUserAccessBinding) pulumi.StringArrayOutput { return v.AccessLevels }).(pulumi.StringArrayOutput)
}

// Optional. Dry run access level that will be evaluated but will not be enforced. The access denial based on dry run policy will be logged. Only one access level is supported, not multiple. This list must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
func (o GcpUserAccessBindingOutput) DryRunAccessLevels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GcpUserAccessBinding) pulumi.StringArrayOutput { return v.DryRunAccessLevels }).(pulumi.StringArrayOutput)
}

// Immutable. Google Group id whose members are subject to this binding's restrictions. See "id" in the [G Suite Directory API's Groups resource] (https://developers.google.com/admin-sdk/directory/v1/reference/groups#resource). If a group's email address/alias is changed, this resource will continue to point at the changed group. This field does not accept group email addresses or aliases. Example: "01d520gv4vjcrht"
func (o GcpUserAccessBindingOutput) GroupKey() pulumi.StringOutput {
	return o.ApplyT(func(v *GcpUserAccessBinding) pulumi.StringOutput { return v.GroupKey }).(pulumi.StringOutput)
}

// Immutable. Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved characters (as defined by [RFC 3986 Section 2.3](https://tools.ietf.org/html/rfc3986#section-2.3)). Should not be specified by the client during creation. Example: "organizations/256/gcpUserAccessBindings/b3-BhcX_Ud5N"
func (o GcpUserAccessBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GcpUserAccessBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GcpUserAccessBindingOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *GcpUserAccessBinding) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GcpUserAccessBindingInput)(nil)).Elem(), &GcpUserAccessBinding{})
	pulumi.RegisterOutputType(GcpUserAccessBindingOutput{})
}
