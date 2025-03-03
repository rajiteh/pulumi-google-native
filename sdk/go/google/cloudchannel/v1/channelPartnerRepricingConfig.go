// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a ChannelPartnerRepricingConfig. Call this method to set modifications for a specific ChannelPartner's bill. You can only create configs if the RepricingConfig.effective_invoice_month is a future month. If needed, you can create a config for the current month, with some restrictions. When creating a config for a future month, make sure there are no existing configs for that RepricingConfig.effective_invoice_month. The following restrictions are for creating configs in the current month. * This functionality is reserved for recovering from an erroneous config, and should not be used for regular business cases. * The new config will not modify exports used with other configs. Changes to the config may be immediate, but may take up to 24 hours. * There is a limit of ten configs for any ChannelPartner or RepricingConfig.effective_invoice_month. * The contained ChannelPartnerRepricingConfig.repricing_config vaule must be different from the value used in the current config for a ChannelPartner. Possible Error Codes: * PERMISSION_DENIED: If the account making the request and the account being queried are different. * INVALID_ARGUMENT: Missing or invalid required parameters in the request. Also displays if the updated config is for the current month or past months. * NOT_FOUND: The ChannelPartnerRepricingConfig specified does not exist or is not associated with the given account. * INTERNAL: Any non-user error related to technical issues in the backend. In this case, contact Cloud Channel support. Return Value: If successful, the updated ChannelPartnerRepricingConfig resource, otherwise returns an error.
// Auto-naming is currently not supported for this resource.
type ChannelPartnerRepricingConfig struct {
	pulumi.CustomResourceState

	AccountId            pulumi.StringOutput `pulumi:"accountId"`
	ChannelPartnerLinkId pulumi.StringOutput `pulumi:"channelPartnerLinkId"`
	// Resource name of the ChannelPartnerRepricingConfig. Format: accounts/{account_id}/channelPartnerLinks/{channel_partner_id}/channelPartnerRepricingConfigs/{id}.
	Name pulumi.StringOutput `pulumi:"name"`
	// The configuration for bill modifications made by a reseller before sending it to ChannelPartner.
	RepricingConfig GoogleCloudChannelV1RepricingConfigResponseOutput `pulumi:"repricingConfig"`
	// Timestamp of an update to the repricing rule. If `update_time` is after RepricingConfig.effective_invoice_month then it indicates this was set mid-month.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewChannelPartnerRepricingConfig registers a new resource with the given unique name, arguments, and options.
func NewChannelPartnerRepricingConfig(ctx *pulumi.Context,
	name string, args *ChannelPartnerRepricingConfigArgs, opts ...pulumi.ResourceOption) (*ChannelPartnerRepricingConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	if args.ChannelPartnerLinkId == nil {
		return nil, errors.New("invalid value for required argument 'ChannelPartnerLinkId'")
	}
	if args.RepricingConfig == nil {
		return nil, errors.New("invalid value for required argument 'RepricingConfig'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"accountId",
		"channelPartnerLinkId",
	})
	opts = append(opts, replaceOnChanges)
	var resource ChannelPartnerRepricingConfig
	err := ctx.RegisterResource("google-native:cloudchannel/v1:ChannelPartnerRepricingConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChannelPartnerRepricingConfig gets an existing ChannelPartnerRepricingConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChannelPartnerRepricingConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelPartnerRepricingConfigState, opts ...pulumi.ResourceOption) (*ChannelPartnerRepricingConfig, error) {
	var resource ChannelPartnerRepricingConfig
	err := ctx.ReadResource("google-native:cloudchannel/v1:ChannelPartnerRepricingConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ChannelPartnerRepricingConfig resources.
type channelPartnerRepricingConfigState struct {
}

type ChannelPartnerRepricingConfigState struct {
}

func (ChannelPartnerRepricingConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*channelPartnerRepricingConfigState)(nil)).Elem()
}

type channelPartnerRepricingConfigArgs struct {
	AccountId            string `pulumi:"accountId"`
	ChannelPartnerLinkId string `pulumi:"channelPartnerLinkId"`
	// The configuration for bill modifications made by a reseller before sending it to ChannelPartner.
	RepricingConfig GoogleCloudChannelV1RepricingConfig `pulumi:"repricingConfig"`
}

// The set of arguments for constructing a ChannelPartnerRepricingConfig resource.
type ChannelPartnerRepricingConfigArgs struct {
	AccountId            pulumi.StringInput
	ChannelPartnerLinkId pulumi.StringInput
	// The configuration for bill modifications made by a reseller before sending it to ChannelPartner.
	RepricingConfig GoogleCloudChannelV1RepricingConfigInput
}

func (ChannelPartnerRepricingConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*channelPartnerRepricingConfigArgs)(nil)).Elem()
}

type ChannelPartnerRepricingConfigInput interface {
	pulumi.Input

	ToChannelPartnerRepricingConfigOutput() ChannelPartnerRepricingConfigOutput
	ToChannelPartnerRepricingConfigOutputWithContext(ctx context.Context) ChannelPartnerRepricingConfigOutput
}

func (*ChannelPartnerRepricingConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ChannelPartnerRepricingConfig)(nil)).Elem()
}

func (i *ChannelPartnerRepricingConfig) ToChannelPartnerRepricingConfigOutput() ChannelPartnerRepricingConfigOutput {
	return i.ToChannelPartnerRepricingConfigOutputWithContext(context.Background())
}

func (i *ChannelPartnerRepricingConfig) ToChannelPartnerRepricingConfigOutputWithContext(ctx context.Context) ChannelPartnerRepricingConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelPartnerRepricingConfigOutput)
}

type ChannelPartnerRepricingConfigOutput struct{ *pulumi.OutputState }

func (ChannelPartnerRepricingConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChannelPartnerRepricingConfig)(nil)).Elem()
}

func (o ChannelPartnerRepricingConfigOutput) ToChannelPartnerRepricingConfigOutput() ChannelPartnerRepricingConfigOutput {
	return o
}

func (o ChannelPartnerRepricingConfigOutput) ToChannelPartnerRepricingConfigOutputWithContext(ctx context.Context) ChannelPartnerRepricingConfigOutput {
	return o
}

func (o ChannelPartnerRepricingConfigOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *ChannelPartnerRepricingConfig) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

func (o ChannelPartnerRepricingConfigOutput) ChannelPartnerLinkId() pulumi.StringOutput {
	return o.ApplyT(func(v *ChannelPartnerRepricingConfig) pulumi.StringOutput { return v.ChannelPartnerLinkId }).(pulumi.StringOutput)
}

// Resource name of the ChannelPartnerRepricingConfig. Format: accounts/{account_id}/channelPartnerLinks/{channel_partner_id}/channelPartnerRepricingConfigs/{id}.
func (o ChannelPartnerRepricingConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ChannelPartnerRepricingConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The configuration for bill modifications made by a reseller before sending it to ChannelPartner.
func (o ChannelPartnerRepricingConfigOutput) RepricingConfig() GoogleCloudChannelV1RepricingConfigResponseOutput {
	return o.ApplyT(func(v *ChannelPartnerRepricingConfig) GoogleCloudChannelV1RepricingConfigResponseOutput {
		return v.RepricingConfig
	}).(GoogleCloudChannelV1RepricingConfigResponseOutput)
}

// Timestamp of an update to the repricing rule. If `update_time` is after RepricingConfig.effective_invoice_month then it indicates this was set mid-month.
func (o ChannelPartnerRepricingConfigOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ChannelPartnerRepricingConfig) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelPartnerRepricingConfigInput)(nil)).Elem(), &ChannelPartnerRepricingConfig{})
	pulumi.RegisterOutputType(ChannelPartnerRepricingConfigOutput{})
}
