// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a PublicDelegatedPrefix in the specified project in the given region using the parameters that are included in the request.
type PublicDelegatedPrefix struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicDelegatedPrefix. An up-to-date fingerprint must be provided in order to update the PublicDelegatedPrefix, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a PublicDelegatedPrefix.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The IPv4 address range, in CIDR format, represented by this public delegated prefix.
	IpCidrRange pulumi.StringOutput `pulumi:"ipCidrRange"`
	// If true, the prefix will be live migrated.
	IsLiveMigration pulumi.BoolOutput `pulumi:"isLiveMigration"`
	// Type of the resource. Always compute#publicDelegatedPrefix for public delegated prefixes.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
	ParentPrefix pulumi.StringOutput `pulumi:"parentPrefix"`
	Project      pulumi.StringOutput `pulumi:"project"`
	// The list of sub public delegated prefixes that exist for this public delegated prefix.
	PublicDelegatedSubPrefixs PublicDelegatedPrefixPublicDelegatedSubPrefixResponseArrayOutput `pulumi:"publicDelegatedSubPrefixs"`
	Region                    pulumi.StringOutput                                              `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The status of the public delegated prefix, which can be one of following values: - `INITIALIZING` The public delegated prefix is being initialized and addresses cannot be created yet. - `READY_TO_ANNOUNCE` The public delegated prefix is a live migration prefix and is active. - `ANNOUNCED` The public delegated prefix is active. - `DELETING` The public delegated prefix is being deprovsioned.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewPublicDelegatedPrefix registers a new resource with the given unique name, arguments, and options.
func NewPublicDelegatedPrefix(ctx *pulumi.Context,
	name string, args *PublicDelegatedPrefixArgs, opts ...pulumi.ResourceOption) (*PublicDelegatedPrefix, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
		"region",
	})
	opts = append(opts, replaceOnChanges)
	var resource PublicDelegatedPrefix
	err := ctx.RegisterResource("google-native:compute/beta:PublicDelegatedPrefix", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicDelegatedPrefix gets an existing PublicDelegatedPrefix resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicDelegatedPrefix(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicDelegatedPrefixState, opts ...pulumi.ResourceOption) (*PublicDelegatedPrefix, error) {
	var resource PublicDelegatedPrefix
	err := ctx.ReadResource("google-native:compute/beta:PublicDelegatedPrefix", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicDelegatedPrefix resources.
type publicDelegatedPrefixState struct {
}

type PublicDelegatedPrefixState struct {
}

func (PublicDelegatedPrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicDelegatedPrefixState)(nil)).Elem()
}

type publicDelegatedPrefixArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// The IPv4 address range, in CIDR format, represented by this public delegated prefix.
	IpCidrRange *string `pulumi:"ipCidrRange"`
	// If true, the prefix will be live migrated.
	IsLiveMigration *bool `pulumi:"isLiveMigration"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
	ParentPrefix *string `pulumi:"parentPrefix"`
	Project      *string `pulumi:"project"`
	// The list of sub public delegated prefixes that exist for this public delegated prefix.
	PublicDelegatedSubPrefixs []PublicDelegatedPrefixPublicDelegatedSubPrefix `pulumi:"publicDelegatedSubPrefixs"`
	Region                    string                                          `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
}

// The set of arguments for constructing a PublicDelegatedPrefix resource.
type PublicDelegatedPrefixArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// The IPv4 address range, in CIDR format, represented by this public delegated prefix.
	IpCidrRange pulumi.StringPtrInput
	// If true, the prefix will be live migrated.
	IsLiveMigration pulumi.BoolPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
	ParentPrefix pulumi.StringPtrInput
	Project      pulumi.StringPtrInput
	// The list of sub public delegated prefixes that exist for this public delegated prefix.
	PublicDelegatedSubPrefixs PublicDelegatedPrefixPublicDelegatedSubPrefixArrayInput
	Region                    pulumi.StringInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
}

func (PublicDelegatedPrefixArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicDelegatedPrefixArgs)(nil)).Elem()
}

type PublicDelegatedPrefixInput interface {
	pulumi.Input

	ToPublicDelegatedPrefixOutput() PublicDelegatedPrefixOutput
	ToPublicDelegatedPrefixOutputWithContext(ctx context.Context) PublicDelegatedPrefixOutput
}

func (*PublicDelegatedPrefix) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicDelegatedPrefix)(nil)).Elem()
}

func (i *PublicDelegatedPrefix) ToPublicDelegatedPrefixOutput() PublicDelegatedPrefixOutput {
	return i.ToPublicDelegatedPrefixOutputWithContext(context.Background())
}

func (i *PublicDelegatedPrefix) ToPublicDelegatedPrefixOutputWithContext(ctx context.Context) PublicDelegatedPrefixOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicDelegatedPrefixOutput)
}

type PublicDelegatedPrefixOutput struct{ *pulumi.OutputState }

func (PublicDelegatedPrefixOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicDelegatedPrefix)(nil)).Elem()
}

func (o PublicDelegatedPrefixOutput) ToPublicDelegatedPrefixOutput() PublicDelegatedPrefixOutput {
	return o
}

func (o PublicDelegatedPrefixOutput) ToPublicDelegatedPrefixOutputWithContext(ctx context.Context) PublicDelegatedPrefixOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o PublicDelegatedPrefixOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicDelegatedPrefix) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o PublicDelegatedPrefixOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicDelegatedPrefix) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a new PublicDelegatedPrefix. An up-to-date fingerprint must be provided in order to update the PublicDelegatedPrefix, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a PublicDelegatedPrefix.
func (o PublicDelegatedPrefixOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicDelegatedPrefix) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// The IPv4 address range, in CIDR format, represented by this public delegated prefix.
func (o PublicDelegatedPrefixOutput) IpCidrRange() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicDelegatedPrefix) pulumi.StringOutput { return v.IpCidrRange }).(pulumi.StringOutput)
}

// If true, the prefix will be live migrated.
func (o PublicDelegatedPrefixOutput) IsLiveMigration() pulumi.BoolOutput {
	return o.ApplyT(func(v *PublicDelegatedPrefix) pulumi.BoolOutput { return v.IsLiveMigration }).(pulumi.BoolOutput)
}

// Type of the resource. Always compute#publicDelegatedPrefix for public delegated prefixes.
func (o PublicDelegatedPrefixOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicDelegatedPrefix) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o PublicDelegatedPrefixOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicDelegatedPrefix) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.
func (o PublicDelegatedPrefixOutput) ParentPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicDelegatedPrefix) pulumi.StringOutput { return v.ParentPrefix }).(pulumi.StringOutput)
}

func (o PublicDelegatedPrefixOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicDelegatedPrefix) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The list of sub public delegated prefixes that exist for this public delegated prefix.
func (o PublicDelegatedPrefixOutput) PublicDelegatedSubPrefixs() PublicDelegatedPrefixPublicDelegatedSubPrefixResponseArrayOutput {
	return o.ApplyT(func(v *PublicDelegatedPrefix) PublicDelegatedPrefixPublicDelegatedSubPrefixResponseArrayOutput {
		return v.PublicDelegatedSubPrefixs
	}).(PublicDelegatedPrefixPublicDelegatedSubPrefixResponseArrayOutput)
}

func (o PublicDelegatedPrefixOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicDelegatedPrefix) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
func (o PublicDelegatedPrefixOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicDelegatedPrefix) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// Server-defined URL for the resource.
func (o PublicDelegatedPrefixOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicDelegatedPrefix) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The status of the public delegated prefix, which can be one of following values: - `INITIALIZING` The public delegated prefix is being initialized and addresses cannot be created yet. - `READY_TO_ANNOUNCE` The public delegated prefix is a live migration prefix and is active. - `ANNOUNCED` The public delegated prefix is active. - `DELETING` The public delegated prefix is being deprovsioned.
func (o PublicDelegatedPrefixOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *PublicDelegatedPrefix) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PublicDelegatedPrefixInput)(nil)).Elem(), &PublicDelegatedPrefix{})
	pulumi.RegisterOutputType(PublicDelegatedPrefixOutput{})
}
