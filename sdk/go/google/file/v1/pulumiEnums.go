// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The service tier of the instance.
type InstanceTier string

const (
	// Not set.
	InstanceTierTierUnspecified = InstanceTier("TIER_UNSPECIFIED")
	// STANDARD tier. BASIC_HDD is the preferred term for this tier.
	InstanceTierStandard = InstanceTier("STANDARD")
	// PREMIUM tier. BASIC_SSD is the preferred term for this tier.
	InstanceTierPremium = InstanceTier("PREMIUM")
	// BASIC instances offer a maximum capacity of 63.9 TB. BASIC_HDD is an alias for STANDARD Tier, offering economical performance backed by HDD.
	InstanceTierBasicHdd = InstanceTier("BASIC_HDD")
	// BASIC instances offer a maximum capacity of 63.9 TB. BASIC_SSD is an alias for PREMIUM Tier, and offers improved performance backed by SSD.
	InstanceTierBasicSsd = InstanceTier("BASIC_SSD")
	// HIGH_SCALE instances offer expanded capacity and performance scaling capabilities.
	InstanceTierHighScaleSsd = InstanceTier("HIGH_SCALE_SSD")
	// ENTERPRISE instances offer the features and availability needed for mission-critical workloads.
	InstanceTierEnterprise = InstanceTier("ENTERPRISE")
	// ZONAL instances offer expanded capacity and performance scaling capabilities.
	InstanceTierZonal = InstanceTier("ZONAL")
)

func (InstanceTier) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceTier)(nil)).Elem()
}

func (e InstanceTier) ToInstanceTierOutput() InstanceTierOutput {
	return pulumi.ToOutput(e).(InstanceTierOutput)
}

func (e InstanceTier) ToInstanceTierOutputWithContext(ctx context.Context) InstanceTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InstanceTierOutput)
}

func (e InstanceTier) ToInstanceTierPtrOutput() InstanceTierPtrOutput {
	return e.ToInstanceTierPtrOutputWithContext(context.Background())
}

func (e InstanceTier) ToInstanceTierPtrOutputWithContext(ctx context.Context) InstanceTierPtrOutput {
	return InstanceTier(e).ToInstanceTierOutputWithContext(ctx).ToInstanceTierPtrOutputWithContext(ctx)
}

func (e InstanceTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InstanceTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InstanceTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InstanceTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InstanceTierOutput struct{ *pulumi.OutputState }

func (InstanceTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceTier)(nil)).Elem()
}

func (o InstanceTierOutput) ToInstanceTierOutput() InstanceTierOutput {
	return o
}

func (o InstanceTierOutput) ToInstanceTierOutputWithContext(ctx context.Context) InstanceTierOutput {
	return o
}

func (o InstanceTierOutput) ToInstanceTierPtrOutput() InstanceTierPtrOutput {
	return o.ToInstanceTierPtrOutputWithContext(context.Background())
}

func (o InstanceTierOutput) ToInstanceTierPtrOutputWithContext(ctx context.Context) InstanceTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstanceTier) *InstanceTier {
		return &v
	}).(InstanceTierPtrOutput)
}

func (o InstanceTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InstanceTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InstanceTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InstanceTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InstanceTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InstanceTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InstanceTierPtrOutput struct{ *pulumi.OutputState }

func (InstanceTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceTier)(nil)).Elem()
}

func (o InstanceTierPtrOutput) ToInstanceTierPtrOutput() InstanceTierPtrOutput {
	return o
}

func (o InstanceTierPtrOutput) ToInstanceTierPtrOutputWithContext(ctx context.Context) InstanceTierPtrOutput {
	return o
}

func (o InstanceTierPtrOutput) Elem() InstanceTierOutput {
	return o.ApplyT(func(v *InstanceTier) InstanceTier {
		if v != nil {
			return *v
		}
		var ret InstanceTier
		return ret
	}).(InstanceTierOutput)
}

func (o InstanceTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InstanceTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InstanceTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// InstanceTierInput is an input type that accepts InstanceTierArgs and InstanceTierOutput values.
// You can construct a concrete instance of `InstanceTierInput` via:
//
//	InstanceTierArgs{...}
type InstanceTierInput interface {
	pulumi.Input

	ToInstanceTierOutput() InstanceTierOutput
	ToInstanceTierOutputWithContext(context.Context) InstanceTierOutput
}

var instanceTierPtrType = reflect.TypeOf((**InstanceTier)(nil)).Elem()

type InstanceTierPtrInput interface {
	pulumi.Input

	ToInstanceTierPtrOutput() InstanceTierPtrOutput
	ToInstanceTierPtrOutputWithContext(context.Context) InstanceTierPtrOutput
}

type instanceTierPtr string

func InstanceTierPtr(v string) InstanceTierPtrInput {
	return (*instanceTierPtr)(&v)
}

func (*instanceTierPtr) ElementType() reflect.Type {
	return instanceTierPtrType
}

func (in *instanceTierPtr) ToInstanceTierPtrOutput() InstanceTierPtrOutput {
	return pulumi.ToOutput(in).(InstanceTierPtrOutput)
}

func (in *instanceTierPtr) ToInstanceTierPtrOutputWithContext(ctx context.Context) InstanceTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InstanceTierPtrOutput)
}

// The network connect mode of the Filestore instance. If not provided, the connect mode defaults to DIRECT_PEERING.
type NetworkConfigConnectMode string

const (
	// Not set.
	NetworkConfigConnectModeConnectModeUnspecified = NetworkConfigConnectMode("CONNECT_MODE_UNSPECIFIED")
	// Connect via direct peering to the Filestore service.
	NetworkConfigConnectModeDirectPeering = NetworkConfigConnectMode("DIRECT_PEERING")
	// Connect to your Filestore instance using Private Service Access. Private services access provides an IP address range for multiple Google Cloud services, including Filestore.
	NetworkConfigConnectModePrivateServiceAccess = NetworkConfigConnectMode("PRIVATE_SERVICE_ACCESS")
)

func (NetworkConfigConnectMode) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkConfigConnectMode)(nil)).Elem()
}

func (e NetworkConfigConnectMode) ToNetworkConfigConnectModeOutput() NetworkConfigConnectModeOutput {
	return pulumi.ToOutput(e).(NetworkConfigConnectModeOutput)
}

func (e NetworkConfigConnectMode) ToNetworkConfigConnectModeOutputWithContext(ctx context.Context) NetworkConfigConnectModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkConfigConnectModeOutput)
}

func (e NetworkConfigConnectMode) ToNetworkConfigConnectModePtrOutput() NetworkConfigConnectModePtrOutput {
	return e.ToNetworkConfigConnectModePtrOutputWithContext(context.Background())
}

func (e NetworkConfigConnectMode) ToNetworkConfigConnectModePtrOutputWithContext(ctx context.Context) NetworkConfigConnectModePtrOutput {
	return NetworkConfigConnectMode(e).ToNetworkConfigConnectModeOutputWithContext(ctx).ToNetworkConfigConnectModePtrOutputWithContext(ctx)
}

func (e NetworkConfigConnectMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkConfigConnectMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkConfigConnectMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkConfigConnectMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkConfigConnectModeOutput struct{ *pulumi.OutputState }

func (NetworkConfigConnectModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkConfigConnectMode)(nil)).Elem()
}

func (o NetworkConfigConnectModeOutput) ToNetworkConfigConnectModeOutput() NetworkConfigConnectModeOutput {
	return o
}

func (o NetworkConfigConnectModeOutput) ToNetworkConfigConnectModeOutputWithContext(ctx context.Context) NetworkConfigConnectModeOutput {
	return o
}

func (o NetworkConfigConnectModeOutput) ToNetworkConfigConnectModePtrOutput() NetworkConfigConnectModePtrOutput {
	return o.ToNetworkConfigConnectModePtrOutputWithContext(context.Background())
}

func (o NetworkConfigConnectModeOutput) ToNetworkConfigConnectModePtrOutputWithContext(ctx context.Context) NetworkConfigConnectModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkConfigConnectMode) *NetworkConfigConnectMode {
		return &v
	}).(NetworkConfigConnectModePtrOutput)
}

func (o NetworkConfigConnectModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkConfigConnectModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkConfigConnectMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkConfigConnectModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkConfigConnectModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkConfigConnectMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkConfigConnectModePtrOutput struct{ *pulumi.OutputState }

func (NetworkConfigConnectModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkConfigConnectMode)(nil)).Elem()
}

func (o NetworkConfigConnectModePtrOutput) ToNetworkConfigConnectModePtrOutput() NetworkConfigConnectModePtrOutput {
	return o
}

func (o NetworkConfigConnectModePtrOutput) ToNetworkConfigConnectModePtrOutputWithContext(ctx context.Context) NetworkConfigConnectModePtrOutput {
	return o
}

func (o NetworkConfigConnectModePtrOutput) Elem() NetworkConfigConnectModeOutput {
	return o.ApplyT(func(v *NetworkConfigConnectMode) NetworkConfigConnectMode {
		if v != nil {
			return *v
		}
		var ret NetworkConfigConnectMode
		return ret
	}).(NetworkConfigConnectModeOutput)
}

func (o NetworkConfigConnectModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkConfigConnectModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkConfigConnectMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// NetworkConfigConnectModeInput is an input type that accepts NetworkConfigConnectModeArgs and NetworkConfigConnectModeOutput values.
// You can construct a concrete instance of `NetworkConfigConnectModeInput` via:
//
//	NetworkConfigConnectModeArgs{...}
type NetworkConfigConnectModeInput interface {
	pulumi.Input

	ToNetworkConfigConnectModeOutput() NetworkConfigConnectModeOutput
	ToNetworkConfigConnectModeOutputWithContext(context.Context) NetworkConfigConnectModeOutput
}

var networkConfigConnectModePtrType = reflect.TypeOf((**NetworkConfigConnectMode)(nil)).Elem()

type NetworkConfigConnectModePtrInput interface {
	pulumi.Input

	ToNetworkConfigConnectModePtrOutput() NetworkConfigConnectModePtrOutput
	ToNetworkConfigConnectModePtrOutputWithContext(context.Context) NetworkConfigConnectModePtrOutput
}

type networkConfigConnectModePtr string

func NetworkConfigConnectModePtr(v string) NetworkConfigConnectModePtrInput {
	return (*networkConfigConnectModePtr)(&v)
}

func (*networkConfigConnectModePtr) ElementType() reflect.Type {
	return networkConfigConnectModePtrType
}

func (in *networkConfigConnectModePtr) ToNetworkConfigConnectModePtrOutput() NetworkConfigConnectModePtrOutput {
	return pulumi.ToOutput(in).(NetworkConfigConnectModePtrOutput)
}

func (in *networkConfigConnectModePtr) ToNetworkConfigConnectModePtrOutputWithContext(ctx context.Context) NetworkConfigConnectModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkConfigConnectModePtrOutput)
}

type NetworkConfigModesItem string

const (
	// Internet protocol not set.
	NetworkConfigModesItemAddressModeUnspecified = NetworkConfigModesItem("ADDRESS_MODE_UNSPECIFIED")
	// Use the IPv4 internet protocol.
	NetworkConfigModesItemModeIpv4 = NetworkConfigModesItem("MODE_IPV4")
)

func (NetworkConfigModesItem) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkConfigModesItem)(nil)).Elem()
}

func (e NetworkConfigModesItem) ToNetworkConfigModesItemOutput() NetworkConfigModesItemOutput {
	return pulumi.ToOutput(e).(NetworkConfigModesItemOutput)
}

func (e NetworkConfigModesItem) ToNetworkConfigModesItemOutputWithContext(ctx context.Context) NetworkConfigModesItemOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkConfigModesItemOutput)
}

func (e NetworkConfigModesItem) ToNetworkConfigModesItemPtrOutput() NetworkConfigModesItemPtrOutput {
	return e.ToNetworkConfigModesItemPtrOutputWithContext(context.Background())
}

func (e NetworkConfigModesItem) ToNetworkConfigModesItemPtrOutputWithContext(ctx context.Context) NetworkConfigModesItemPtrOutput {
	return NetworkConfigModesItem(e).ToNetworkConfigModesItemOutputWithContext(ctx).ToNetworkConfigModesItemPtrOutputWithContext(ctx)
}

func (e NetworkConfigModesItem) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkConfigModesItem) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkConfigModesItem) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkConfigModesItem) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkConfigModesItemOutput struct{ *pulumi.OutputState }

func (NetworkConfigModesItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkConfigModesItem)(nil)).Elem()
}

func (o NetworkConfigModesItemOutput) ToNetworkConfigModesItemOutput() NetworkConfigModesItemOutput {
	return o
}

func (o NetworkConfigModesItemOutput) ToNetworkConfigModesItemOutputWithContext(ctx context.Context) NetworkConfigModesItemOutput {
	return o
}

func (o NetworkConfigModesItemOutput) ToNetworkConfigModesItemPtrOutput() NetworkConfigModesItemPtrOutput {
	return o.ToNetworkConfigModesItemPtrOutputWithContext(context.Background())
}

func (o NetworkConfigModesItemOutput) ToNetworkConfigModesItemPtrOutputWithContext(ctx context.Context) NetworkConfigModesItemPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkConfigModesItem) *NetworkConfigModesItem {
		return &v
	}).(NetworkConfigModesItemPtrOutput)
}

func (o NetworkConfigModesItemOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkConfigModesItemOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkConfigModesItem) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkConfigModesItemOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkConfigModesItemOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkConfigModesItem) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkConfigModesItemPtrOutput struct{ *pulumi.OutputState }

func (NetworkConfigModesItemPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkConfigModesItem)(nil)).Elem()
}

func (o NetworkConfigModesItemPtrOutput) ToNetworkConfigModesItemPtrOutput() NetworkConfigModesItemPtrOutput {
	return o
}

func (o NetworkConfigModesItemPtrOutput) ToNetworkConfigModesItemPtrOutputWithContext(ctx context.Context) NetworkConfigModesItemPtrOutput {
	return o
}

func (o NetworkConfigModesItemPtrOutput) Elem() NetworkConfigModesItemOutput {
	return o.ApplyT(func(v *NetworkConfigModesItem) NetworkConfigModesItem {
		if v != nil {
			return *v
		}
		var ret NetworkConfigModesItem
		return ret
	}).(NetworkConfigModesItemOutput)
}

func (o NetworkConfigModesItemPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkConfigModesItemPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkConfigModesItem) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// NetworkConfigModesItemInput is an input type that accepts NetworkConfigModesItemArgs and NetworkConfigModesItemOutput values.
// You can construct a concrete instance of `NetworkConfigModesItemInput` via:
//
//	NetworkConfigModesItemArgs{...}
type NetworkConfigModesItemInput interface {
	pulumi.Input

	ToNetworkConfigModesItemOutput() NetworkConfigModesItemOutput
	ToNetworkConfigModesItemOutputWithContext(context.Context) NetworkConfigModesItemOutput
}

var networkConfigModesItemPtrType = reflect.TypeOf((**NetworkConfigModesItem)(nil)).Elem()

type NetworkConfigModesItemPtrInput interface {
	pulumi.Input

	ToNetworkConfigModesItemPtrOutput() NetworkConfigModesItemPtrOutput
	ToNetworkConfigModesItemPtrOutputWithContext(context.Context) NetworkConfigModesItemPtrOutput
}

type networkConfigModesItemPtr string

func NetworkConfigModesItemPtr(v string) NetworkConfigModesItemPtrInput {
	return (*networkConfigModesItemPtr)(&v)
}

func (*networkConfigModesItemPtr) ElementType() reflect.Type {
	return networkConfigModesItemPtrType
}

func (in *networkConfigModesItemPtr) ToNetworkConfigModesItemPtrOutput() NetworkConfigModesItemPtrOutput {
	return pulumi.ToOutput(in).(NetworkConfigModesItemPtrOutput)
}

func (in *networkConfigModesItemPtr) ToNetworkConfigModesItemPtrOutputWithContext(ctx context.Context) NetworkConfigModesItemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkConfigModesItemPtrOutput)
}

// NetworkConfigModesItemArrayInput is an input type that accepts NetworkConfigModesItemArray and NetworkConfigModesItemArrayOutput values.
// You can construct a concrete instance of `NetworkConfigModesItemArrayInput` via:
//
//	NetworkConfigModesItemArray{ NetworkConfigModesItemArgs{...} }
type NetworkConfigModesItemArrayInput interface {
	pulumi.Input

	ToNetworkConfigModesItemArrayOutput() NetworkConfigModesItemArrayOutput
	ToNetworkConfigModesItemArrayOutputWithContext(context.Context) NetworkConfigModesItemArrayOutput
}

type NetworkConfigModesItemArray []NetworkConfigModesItem

func (NetworkConfigModesItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkConfigModesItem)(nil)).Elem()
}

func (i NetworkConfigModesItemArray) ToNetworkConfigModesItemArrayOutput() NetworkConfigModesItemArrayOutput {
	return i.ToNetworkConfigModesItemArrayOutputWithContext(context.Background())
}

func (i NetworkConfigModesItemArray) ToNetworkConfigModesItemArrayOutputWithContext(ctx context.Context) NetworkConfigModesItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkConfigModesItemArrayOutput)
}

type NetworkConfigModesItemArrayOutput struct{ *pulumi.OutputState }

func (NetworkConfigModesItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkConfigModesItem)(nil)).Elem()
}

func (o NetworkConfigModesItemArrayOutput) ToNetworkConfigModesItemArrayOutput() NetworkConfigModesItemArrayOutput {
	return o
}

func (o NetworkConfigModesItemArrayOutput) ToNetworkConfigModesItemArrayOutputWithContext(ctx context.Context) NetworkConfigModesItemArrayOutput {
	return o
}

func (o NetworkConfigModesItemArrayOutput) Index(i pulumi.IntInput) NetworkConfigModesItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkConfigModesItem {
		return vs[0].([]NetworkConfigModesItem)[vs[1].(int)]
	}).(NetworkConfigModesItemOutput)
}

// Either READ_ONLY, for allowing only read requests on the exported directory, or READ_WRITE, for allowing both read and write requests. The default is READ_WRITE.
type NfsExportOptionsAccessMode string

const (
	// AccessMode not set.
	NfsExportOptionsAccessModeAccessModeUnspecified = NfsExportOptionsAccessMode("ACCESS_MODE_UNSPECIFIED")
	// The client can only read the file share.
	NfsExportOptionsAccessModeReadOnly = NfsExportOptionsAccessMode("READ_ONLY")
	// The client can read and write the file share (default).
	NfsExportOptionsAccessModeReadWrite = NfsExportOptionsAccessMode("READ_WRITE")
)

func (NfsExportOptionsAccessMode) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsExportOptionsAccessMode)(nil)).Elem()
}

func (e NfsExportOptionsAccessMode) ToNfsExportOptionsAccessModeOutput() NfsExportOptionsAccessModeOutput {
	return pulumi.ToOutput(e).(NfsExportOptionsAccessModeOutput)
}

func (e NfsExportOptionsAccessMode) ToNfsExportOptionsAccessModeOutputWithContext(ctx context.Context) NfsExportOptionsAccessModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NfsExportOptionsAccessModeOutput)
}

func (e NfsExportOptionsAccessMode) ToNfsExportOptionsAccessModePtrOutput() NfsExportOptionsAccessModePtrOutput {
	return e.ToNfsExportOptionsAccessModePtrOutputWithContext(context.Background())
}

func (e NfsExportOptionsAccessMode) ToNfsExportOptionsAccessModePtrOutputWithContext(ctx context.Context) NfsExportOptionsAccessModePtrOutput {
	return NfsExportOptionsAccessMode(e).ToNfsExportOptionsAccessModeOutputWithContext(ctx).ToNfsExportOptionsAccessModePtrOutputWithContext(ctx)
}

func (e NfsExportOptionsAccessMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NfsExportOptionsAccessMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NfsExportOptionsAccessMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NfsExportOptionsAccessMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NfsExportOptionsAccessModeOutput struct{ *pulumi.OutputState }

func (NfsExportOptionsAccessModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsExportOptionsAccessMode)(nil)).Elem()
}

func (o NfsExportOptionsAccessModeOutput) ToNfsExportOptionsAccessModeOutput() NfsExportOptionsAccessModeOutput {
	return o
}

func (o NfsExportOptionsAccessModeOutput) ToNfsExportOptionsAccessModeOutputWithContext(ctx context.Context) NfsExportOptionsAccessModeOutput {
	return o
}

func (o NfsExportOptionsAccessModeOutput) ToNfsExportOptionsAccessModePtrOutput() NfsExportOptionsAccessModePtrOutput {
	return o.ToNfsExportOptionsAccessModePtrOutputWithContext(context.Background())
}

func (o NfsExportOptionsAccessModeOutput) ToNfsExportOptionsAccessModePtrOutputWithContext(ctx context.Context) NfsExportOptionsAccessModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NfsExportOptionsAccessMode) *NfsExportOptionsAccessMode {
		return &v
	}).(NfsExportOptionsAccessModePtrOutput)
}

func (o NfsExportOptionsAccessModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NfsExportOptionsAccessModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NfsExportOptionsAccessMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NfsExportOptionsAccessModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NfsExportOptionsAccessModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NfsExportOptionsAccessMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NfsExportOptionsAccessModePtrOutput struct{ *pulumi.OutputState }

func (NfsExportOptionsAccessModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NfsExportOptionsAccessMode)(nil)).Elem()
}

func (o NfsExportOptionsAccessModePtrOutput) ToNfsExportOptionsAccessModePtrOutput() NfsExportOptionsAccessModePtrOutput {
	return o
}

func (o NfsExportOptionsAccessModePtrOutput) ToNfsExportOptionsAccessModePtrOutputWithContext(ctx context.Context) NfsExportOptionsAccessModePtrOutput {
	return o
}

func (o NfsExportOptionsAccessModePtrOutput) Elem() NfsExportOptionsAccessModeOutput {
	return o.ApplyT(func(v *NfsExportOptionsAccessMode) NfsExportOptionsAccessMode {
		if v != nil {
			return *v
		}
		var ret NfsExportOptionsAccessMode
		return ret
	}).(NfsExportOptionsAccessModeOutput)
}

func (o NfsExportOptionsAccessModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NfsExportOptionsAccessModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NfsExportOptionsAccessMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// NfsExportOptionsAccessModeInput is an input type that accepts NfsExportOptionsAccessModeArgs and NfsExportOptionsAccessModeOutput values.
// You can construct a concrete instance of `NfsExportOptionsAccessModeInput` via:
//
//	NfsExportOptionsAccessModeArgs{...}
type NfsExportOptionsAccessModeInput interface {
	pulumi.Input

	ToNfsExportOptionsAccessModeOutput() NfsExportOptionsAccessModeOutput
	ToNfsExportOptionsAccessModeOutputWithContext(context.Context) NfsExportOptionsAccessModeOutput
}

var nfsExportOptionsAccessModePtrType = reflect.TypeOf((**NfsExportOptionsAccessMode)(nil)).Elem()

type NfsExportOptionsAccessModePtrInput interface {
	pulumi.Input

	ToNfsExportOptionsAccessModePtrOutput() NfsExportOptionsAccessModePtrOutput
	ToNfsExportOptionsAccessModePtrOutputWithContext(context.Context) NfsExportOptionsAccessModePtrOutput
}

type nfsExportOptionsAccessModePtr string

func NfsExportOptionsAccessModePtr(v string) NfsExportOptionsAccessModePtrInput {
	return (*nfsExportOptionsAccessModePtr)(&v)
}

func (*nfsExportOptionsAccessModePtr) ElementType() reflect.Type {
	return nfsExportOptionsAccessModePtrType
}

func (in *nfsExportOptionsAccessModePtr) ToNfsExportOptionsAccessModePtrOutput() NfsExportOptionsAccessModePtrOutput {
	return pulumi.ToOutput(in).(NfsExportOptionsAccessModePtrOutput)
}

func (in *nfsExportOptionsAccessModePtr) ToNfsExportOptionsAccessModePtrOutputWithContext(ctx context.Context) NfsExportOptionsAccessModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NfsExportOptionsAccessModePtrOutput)
}

// Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH, for not allowing root access. The default is NO_ROOT_SQUASH.
type NfsExportOptionsSquashMode string

const (
	// SquashMode not set.
	NfsExportOptionsSquashModeSquashModeUnspecified = NfsExportOptionsSquashMode("SQUASH_MODE_UNSPECIFIED")
	// The Root user has root access to the file share (default).
	NfsExportOptionsSquashModeNoRootSquash = NfsExportOptionsSquashMode("NO_ROOT_SQUASH")
	// The Root user has squashed access to the anonymous uid/gid.
	NfsExportOptionsSquashModeRootSquash = NfsExportOptionsSquashMode("ROOT_SQUASH")
)

func (NfsExportOptionsSquashMode) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsExportOptionsSquashMode)(nil)).Elem()
}

func (e NfsExportOptionsSquashMode) ToNfsExportOptionsSquashModeOutput() NfsExportOptionsSquashModeOutput {
	return pulumi.ToOutput(e).(NfsExportOptionsSquashModeOutput)
}

func (e NfsExportOptionsSquashMode) ToNfsExportOptionsSquashModeOutputWithContext(ctx context.Context) NfsExportOptionsSquashModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NfsExportOptionsSquashModeOutput)
}

func (e NfsExportOptionsSquashMode) ToNfsExportOptionsSquashModePtrOutput() NfsExportOptionsSquashModePtrOutput {
	return e.ToNfsExportOptionsSquashModePtrOutputWithContext(context.Background())
}

func (e NfsExportOptionsSquashMode) ToNfsExportOptionsSquashModePtrOutputWithContext(ctx context.Context) NfsExportOptionsSquashModePtrOutput {
	return NfsExportOptionsSquashMode(e).ToNfsExportOptionsSquashModeOutputWithContext(ctx).ToNfsExportOptionsSquashModePtrOutputWithContext(ctx)
}

func (e NfsExportOptionsSquashMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NfsExportOptionsSquashMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NfsExportOptionsSquashMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NfsExportOptionsSquashMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NfsExportOptionsSquashModeOutput struct{ *pulumi.OutputState }

func (NfsExportOptionsSquashModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NfsExportOptionsSquashMode)(nil)).Elem()
}

func (o NfsExportOptionsSquashModeOutput) ToNfsExportOptionsSquashModeOutput() NfsExportOptionsSquashModeOutput {
	return o
}

func (o NfsExportOptionsSquashModeOutput) ToNfsExportOptionsSquashModeOutputWithContext(ctx context.Context) NfsExportOptionsSquashModeOutput {
	return o
}

func (o NfsExportOptionsSquashModeOutput) ToNfsExportOptionsSquashModePtrOutput() NfsExportOptionsSquashModePtrOutput {
	return o.ToNfsExportOptionsSquashModePtrOutputWithContext(context.Background())
}

func (o NfsExportOptionsSquashModeOutput) ToNfsExportOptionsSquashModePtrOutputWithContext(ctx context.Context) NfsExportOptionsSquashModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NfsExportOptionsSquashMode) *NfsExportOptionsSquashMode {
		return &v
	}).(NfsExportOptionsSquashModePtrOutput)
}

func (o NfsExportOptionsSquashModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NfsExportOptionsSquashModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NfsExportOptionsSquashMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NfsExportOptionsSquashModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NfsExportOptionsSquashModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NfsExportOptionsSquashMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NfsExportOptionsSquashModePtrOutput struct{ *pulumi.OutputState }

func (NfsExportOptionsSquashModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NfsExportOptionsSquashMode)(nil)).Elem()
}

func (o NfsExportOptionsSquashModePtrOutput) ToNfsExportOptionsSquashModePtrOutput() NfsExportOptionsSquashModePtrOutput {
	return o
}

func (o NfsExportOptionsSquashModePtrOutput) ToNfsExportOptionsSquashModePtrOutputWithContext(ctx context.Context) NfsExportOptionsSquashModePtrOutput {
	return o
}

func (o NfsExportOptionsSquashModePtrOutput) Elem() NfsExportOptionsSquashModeOutput {
	return o.ApplyT(func(v *NfsExportOptionsSquashMode) NfsExportOptionsSquashMode {
		if v != nil {
			return *v
		}
		var ret NfsExportOptionsSquashMode
		return ret
	}).(NfsExportOptionsSquashModeOutput)
}

func (o NfsExportOptionsSquashModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NfsExportOptionsSquashModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NfsExportOptionsSquashMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// NfsExportOptionsSquashModeInput is an input type that accepts NfsExportOptionsSquashModeArgs and NfsExportOptionsSquashModeOutput values.
// You can construct a concrete instance of `NfsExportOptionsSquashModeInput` via:
//
//	NfsExportOptionsSquashModeArgs{...}
type NfsExportOptionsSquashModeInput interface {
	pulumi.Input

	ToNfsExportOptionsSquashModeOutput() NfsExportOptionsSquashModeOutput
	ToNfsExportOptionsSquashModeOutputWithContext(context.Context) NfsExportOptionsSquashModeOutput
}

var nfsExportOptionsSquashModePtrType = reflect.TypeOf((**NfsExportOptionsSquashMode)(nil)).Elem()

type NfsExportOptionsSquashModePtrInput interface {
	pulumi.Input

	ToNfsExportOptionsSquashModePtrOutput() NfsExportOptionsSquashModePtrOutput
	ToNfsExportOptionsSquashModePtrOutputWithContext(context.Context) NfsExportOptionsSquashModePtrOutput
}

type nfsExportOptionsSquashModePtr string

func NfsExportOptionsSquashModePtr(v string) NfsExportOptionsSquashModePtrInput {
	return (*nfsExportOptionsSquashModePtr)(&v)
}

func (*nfsExportOptionsSquashModePtr) ElementType() reflect.Type {
	return nfsExportOptionsSquashModePtrType
}

func (in *nfsExportOptionsSquashModePtr) ToNfsExportOptionsSquashModePtrOutput() NfsExportOptionsSquashModePtrOutput {
	return pulumi.ToOutput(in).(NfsExportOptionsSquashModePtrOutput)
}

func (in *nfsExportOptionsSquashModePtr) ToNfsExportOptionsSquashModePtrOutputWithContext(ctx context.Context) NfsExportOptionsSquashModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NfsExportOptionsSquashModePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceTierInput)(nil)).Elem(), InstanceTier("TIER_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceTierPtrInput)(nil)).Elem(), InstanceTier("TIER_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkConfigConnectModeInput)(nil)).Elem(), NetworkConfigConnectMode("CONNECT_MODE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkConfigConnectModePtrInput)(nil)).Elem(), NetworkConfigConnectMode("CONNECT_MODE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkConfigModesItemInput)(nil)).Elem(), NetworkConfigModesItem("ADDRESS_MODE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkConfigModesItemPtrInput)(nil)).Elem(), NetworkConfigModesItem("ADDRESS_MODE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkConfigModesItemArrayInput)(nil)).Elem(), NetworkConfigModesItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NfsExportOptionsAccessModeInput)(nil)).Elem(), NfsExportOptionsAccessMode("ACCESS_MODE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*NfsExportOptionsAccessModePtrInput)(nil)).Elem(), NfsExportOptionsAccessMode("ACCESS_MODE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*NfsExportOptionsSquashModeInput)(nil)).Elem(), NfsExportOptionsSquashMode("SQUASH_MODE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*NfsExportOptionsSquashModePtrInput)(nil)).Elem(), NfsExportOptionsSquashMode("SQUASH_MODE_UNSPECIFIED"))
	pulumi.RegisterOutputType(InstanceTierOutput{})
	pulumi.RegisterOutputType(InstanceTierPtrOutput{})
	pulumi.RegisterOutputType(NetworkConfigConnectModeOutput{})
	pulumi.RegisterOutputType(NetworkConfigConnectModePtrOutput{})
	pulumi.RegisterOutputType(NetworkConfigModesItemOutput{})
	pulumi.RegisterOutputType(NetworkConfigModesItemPtrOutput{})
	pulumi.RegisterOutputType(NetworkConfigModesItemArrayOutput{})
	pulumi.RegisterOutputType(NfsExportOptionsAccessModeOutput{})
	pulumi.RegisterOutputType(NfsExportOptionsAccessModePtrOutput{})
	pulumi.RegisterOutputType(NfsExportOptionsSquashModeOutput{})
	pulumi.RegisterOutputType(NfsExportOptionsSquashModePtrOutput{})
}
