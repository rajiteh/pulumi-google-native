// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The log type that this config enables.
type AuditLogConfigLogType string

const (
	// Default case. Should never be this.
	AuditLogConfigLogTypeLogTypeUnspecified = AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED")
	// Admin reads. Example: CloudIAM getIamPolicy
	AuditLogConfigLogTypeAdminRead = AuditLogConfigLogType("ADMIN_READ")
	// Data writes. Example: CloudSQL Users create
	AuditLogConfigLogTypeDataWrite = AuditLogConfigLogType("DATA_WRITE")
	// Data reads. Example: CloudSQL Users list
	AuditLogConfigLogTypeDataRead = AuditLogConfigLogType("DATA_READ")
)

func (AuditLogConfigLogType) ElementType() reflect.Type {
	return reflect.TypeOf((*AuditLogConfigLogType)(nil)).Elem()
}

func (e AuditLogConfigLogType) ToAuditLogConfigLogTypeOutput() AuditLogConfigLogTypeOutput {
	return pulumi.ToOutput(e).(AuditLogConfigLogTypeOutput)
}

func (e AuditLogConfigLogType) ToAuditLogConfigLogTypeOutputWithContext(ctx context.Context) AuditLogConfigLogTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AuditLogConfigLogTypeOutput)
}

func (e AuditLogConfigLogType) ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput {
	return e.ToAuditLogConfigLogTypePtrOutputWithContext(context.Background())
}

func (e AuditLogConfigLogType) ToAuditLogConfigLogTypePtrOutputWithContext(ctx context.Context) AuditLogConfigLogTypePtrOutput {
	return AuditLogConfigLogType(e).ToAuditLogConfigLogTypeOutputWithContext(ctx).ToAuditLogConfigLogTypePtrOutputWithContext(ctx)
}

func (e AuditLogConfigLogType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuditLogConfigLogType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuditLogConfigLogType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AuditLogConfigLogType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AuditLogConfigLogTypeOutput struct{ *pulumi.OutputState }

func (AuditLogConfigLogTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuditLogConfigLogType)(nil)).Elem()
}

func (o AuditLogConfigLogTypeOutput) ToAuditLogConfigLogTypeOutput() AuditLogConfigLogTypeOutput {
	return o
}

func (o AuditLogConfigLogTypeOutput) ToAuditLogConfigLogTypeOutputWithContext(ctx context.Context) AuditLogConfigLogTypeOutput {
	return o
}

func (o AuditLogConfigLogTypeOutput) ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput {
	return o.ToAuditLogConfigLogTypePtrOutputWithContext(context.Background())
}

func (o AuditLogConfigLogTypeOutput) ToAuditLogConfigLogTypePtrOutputWithContext(ctx context.Context) AuditLogConfigLogTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuditLogConfigLogType) *AuditLogConfigLogType {
		return &v
	}).(AuditLogConfigLogTypePtrOutput)
}

func (o AuditLogConfigLogTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AuditLogConfigLogTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuditLogConfigLogType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AuditLogConfigLogTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuditLogConfigLogTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuditLogConfigLogType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AuditLogConfigLogTypePtrOutput struct{ *pulumi.OutputState }

func (AuditLogConfigLogTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuditLogConfigLogType)(nil)).Elem()
}

func (o AuditLogConfigLogTypePtrOutput) ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput {
	return o
}

func (o AuditLogConfigLogTypePtrOutput) ToAuditLogConfigLogTypePtrOutputWithContext(ctx context.Context) AuditLogConfigLogTypePtrOutput {
	return o
}

func (o AuditLogConfigLogTypePtrOutput) Elem() AuditLogConfigLogTypeOutput {
	return o.ApplyT(func(v *AuditLogConfigLogType) AuditLogConfigLogType {
		if v != nil {
			return *v
		}
		var ret AuditLogConfigLogType
		return ret
	}).(AuditLogConfigLogTypeOutput)
}

func (o AuditLogConfigLogTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuditLogConfigLogTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AuditLogConfigLogType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AuditLogConfigLogTypeInput is an input type that accepts AuditLogConfigLogTypeArgs and AuditLogConfigLogTypeOutput values.
// You can construct a concrete instance of `AuditLogConfigLogTypeInput` via:
//
//	AuditLogConfigLogTypeArgs{...}
type AuditLogConfigLogTypeInput interface {
	pulumi.Input

	ToAuditLogConfigLogTypeOutput() AuditLogConfigLogTypeOutput
	ToAuditLogConfigLogTypeOutputWithContext(context.Context) AuditLogConfigLogTypeOutput
}

var auditLogConfigLogTypePtrType = reflect.TypeOf((**AuditLogConfigLogType)(nil)).Elem()

type AuditLogConfigLogTypePtrInput interface {
	pulumi.Input

	ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput
	ToAuditLogConfigLogTypePtrOutputWithContext(context.Context) AuditLogConfigLogTypePtrOutput
}

type auditLogConfigLogTypePtr string

func AuditLogConfigLogTypePtr(v string) AuditLogConfigLogTypePtrInput {
	return (*auditLogConfigLogTypePtr)(&v)
}

func (*auditLogConfigLogTypePtr) ElementType() reflect.Type {
	return auditLogConfigLogTypePtrType
}

func (in *auditLogConfigLogTypePtr) ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput {
	return pulumi.ToOutput(in).(AuditLogConfigLogTypePtrOutput)
}

func (in *auditLogConfigLogTypePtr) ToAuditLogConfigLogTypePtrOutputWithContext(ctx context.Context) AuditLogConfigLogTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AuditLogConfigLogTypePtrOutput)
}

// Immutable. The type of storage used by this cluster to serve its parent instance's tables, unless explicitly overridden.
type ClusterDefaultStorageType string

const (
	// The user did not specify a storage type.
	ClusterDefaultStorageTypeStorageTypeUnspecified = ClusterDefaultStorageType("STORAGE_TYPE_UNSPECIFIED")
	// Flash (SSD) storage should be used.
	ClusterDefaultStorageTypeSsd = ClusterDefaultStorageType("SSD")
	// Magnetic drive (HDD) storage should be used.
	ClusterDefaultStorageTypeHdd = ClusterDefaultStorageType("HDD")
)

func (ClusterDefaultStorageType) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterDefaultStorageType)(nil)).Elem()
}

func (e ClusterDefaultStorageType) ToClusterDefaultStorageTypeOutput() ClusterDefaultStorageTypeOutput {
	return pulumi.ToOutput(e).(ClusterDefaultStorageTypeOutput)
}

func (e ClusterDefaultStorageType) ToClusterDefaultStorageTypeOutputWithContext(ctx context.Context) ClusterDefaultStorageTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ClusterDefaultStorageTypeOutput)
}

func (e ClusterDefaultStorageType) ToClusterDefaultStorageTypePtrOutput() ClusterDefaultStorageTypePtrOutput {
	return e.ToClusterDefaultStorageTypePtrOutputWithContext(context.Background())
}

func (e ClusterDefaultStorageType) ToClusterDefaultStorageTypePtrOutputWithContext(ctx context.Context) ClusterDefaultStorageTypePtrOutput {
	return ClusterDefaultStorageType(e).ToClusterDefaultStorageTypeOutputWithContext(ctx).ToClusterDefaultStorageTypePtrOutputWithContext(ctx)
}

func (e ClusterDefaultStorageType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterDefaultStorageType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterDefaultStorageType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ClusterDefaultStorageType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ClusterDefaultStorageTypeOutput struct{ *pulumi.OutputState }

func (ClusterDefaultStorageTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterDefaultStorageType)(nil)).Elem()
}

func (o ClusterDefaultStorageTypeOutput) ToClusterDefaultStorageTypeOutput() ClusterDefaultStorageTypeOutput {
	return o
}

func (o ClusterDefaultStorageTypeOutput) ToClusterDefaultStorageTypeOutputWithContext(ctx context.Context) ClusterDefaultStorageTypeOutput {
	return o
}

func (o ClusterDefaultStorageTypeOutput) ToClusterDefaultStorageTypePtrOutput() ClusterDefaultStorageTypePtrOutput {
	return o.ToClusterDefaultStorageTypePtrOutputWithContext(context.Background())
}

func (o ClusterDefaultStorageTypeOutput) ToClusterDefaultStorageTypePtrOutputWithContext(ctx context.Context) ClusterDefaultStorageTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterDefaultStorageType) *ClusterDefaultStorageType {
		return &v
	}).(ClusterDefaultStorageTypePtrOutput)
}

func (o ClusterDefaultStorageTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ClusterDefaultStorageTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterDefaultStorageType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ClusterDefaultStorageTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterDefaultStorageTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterDefaultStorageType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ClusterDefaultStorageTypePtrOutput struct{ *pulumi.OutputState }

func (ClusterDefaultStorageTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterDefaultStorageType)(nil)).Elem()
}

func (o ClusterDefaultStorageTypePtrOutput) ToClusterDefaultStorageTypePtrOutput() ClusterDefaultStorageTypePtrOutput {
	return o
}

func (o ClusterDefaultStorageTypePtrOutput) ToClusterDefaultStorageTypePtrOutputWithContext(ctx context.Context) ClusterDefaultStorageTypePtrOutput {
	return o
}

func (o ClusterDefaultStorageTypePtrOutput) Elem() ClusterDefaultStorageTypeOutput {
	return o.ApplyT(func(v *ClusterDefaultStorageType) ClusterDefaultStorageType {
		if v != nil {
			return *v
		}
		var ret ClusterDefaultStorageType
		return ret
	}).(ClusterDefaultStorageTypeOutput)
}

func (o ClusterDefaultStorageTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterDefaultStorageTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ClusterDefaultStorageType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ClusterDefaultStorageTypeInput is an input type that accepts ClusterDefaultStorageTypeArgs and ClusterDefaultStorageTypeOutput values.
// You can construct a concrete instance of `ClusterDefaultStorageTypeInput` via:
//
//	ClusterDefaultStorageTypeArgs{...}
type ClusterDefaultStorageTypeInput interface {
	pulumi.Input

	ToClusterDefaultStorageTypeOutput() ClusterDefaultStorageTypeOutput
	ToClusterDefaultStorageTypeOutputWithContext(context.Context) ClusterDefaultStorageTypeOutput
}

var clusterDefaultStorageTypePtrType = reflect.TypeOf((**ClusterDefaultStorageType)(nil)).Elem()

type ClusterDefaultStorageTypePtrInput interface {
	pulumi.Input

	ToClusterDefaultStorageTypePtrOutput() ClusterDefaultStorageTypePtrOutput
	ToClusterDefaultStorageTypePtrOutputWithContext(context.Context) ClusterDefaultStorageTypePtrOutput
}

type clusterDefaultStorageTypePtr string

func ClusterDefaultStorageTypePtr(v string) ClusterDefaultStorageTypePtrInput {
	return (*clusterDefaultStorageTypePtr)(&v)
}

func (*clusterDefaultStorageTypePtr) ElementType() reflect.Type {
	return clusterDefaultStorageTypePtrType
}

func (in *clusterDefaultStorageTypePtr) ToClusterDefaultStorageTypePtrOutput() ClusterDefaultStorageTypePtrOutput {
	return pulumi.ToOutput(in).(ClusterDefaultStorageTypePtrOutput)
}

func (in *clusterDefaultStorageTypePtr) ToClusterDefaultStorageTypePtrOutputWithContext(ctx context.Context) ClusterDefaultStorageTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ClusterDefaultStorageTypePtrOutput)
}

// The type of the instance. Defaults to `PRODUCTION`.
type InstanceType string

const (
	// The type of the instance is unspecified. If set when creating an instance, a `PRODUCTION` instance will be created. If set when updating an instance, the type will be left unchanged.
	InstanceTypeTypeUnspecified = InstanceType("TYPE_UNSPECIFIED")
	// An instance meant for production use. `serve_nodes` must be set on the cluster.
	InstanceTypeProduction = InstanceType("PRODUCTION")
	// DEPRECATED: Prefer PRODUCTION for all use cases, as it no longer enforces a higher minimum node count than DEVELOPMENT.
	InstanceTypeDevelopment = InstanceType("DEVELOPMENT")
)

func (InstanceType) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceType)(nil)).Elem()
}

func (e InstanceType) ToInstanceTypeOutput() InstanceTypeOutput {
	return pulumi.ToOutput(e).(InstanceTypeOutput)
}

func (e InstanceType) ToInstanceTypeOutputWithContext(ctx context.Context) InstanceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InstanceTypeOutput)
}

func (e InstanceType) ToInstanceTypePtrOutput() InstanceTypePtrOutput {
	return e.ToInstanceTypePtrOutputWithContext(context.Background())
}

func (e InstanceType) ToInstanceTypePtrOutputWithContext(ctx context.Context) InstanceTypePtrOutput {
	return InstanceType(e).ToInstanceTypeOutputWithContext(ctx).ToInstanceTypePtrOutputWithContext(ctx)
}

func (e InstanceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InstanceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InstanceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InstanceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InstanceTypeOutput struct{ *pulumi.OutputState }

func (InstanceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceType)(nil)).Elem()
}

func (o InstanceTypeOutput) ToInstanceTypeOutput() InstanceTypeOutput {
	return o
}

func (o InstanceTypeOutput) ToInstanceTypeOutputWithContext(ctx context.Context) InstanceTypeOutput {
	return o
}

func (o InstanceTypeOutput) ToInstanceTypePtrOutput() InstanceTypePtrOutput {
	return o.ToInstanceTypePtrOutputWithContext(context.Background())
}

func (o InstanceTypeOutput) ToInstanceTypePtrOutputWithContext(ctx context.Context) InstanceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstanceType) *InstanceType {
		return &v
	}).(InstanceTypePtrOutput)
}

func (o InstanceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InstanceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InstanceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InstanceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InstanceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InstanceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InstanceTypePtrOutput struct{ *pulumi.OutputState }

func (InstanceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceType)(nil)).Elem()
}

func (o InstanceTypePtrOutput) ToInstanceTypePtrOutput() InstanceTypePtrOutput {
	return o
}

func (o InstanceTypePtrOutput) ToInstanceTypePtrOutputWithContext(ctx context.Context) InstanceTypePtrOutput {
	return o
}

func (o InstanceTypePtrOutput) Elem() InstanceTypeOutput {
	return o.ApplyT(func(v *InstanceType) InstanceType {
		if v != nil {
			return *v
		}
		var ret InstanceType
		return ret
	}).(InstanceTypeOutput)
}

func (o InstanceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InstanceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InstanceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// InstanceTypeInput is an input type that accepts InstanceTypeArgs and InstanceTypeOutput values.
// You can construct a concrete instance of `InstanceTypeInput` via:
//
//	InstanceTypeArgs{...}
type InstanceTypeInput interface {
	pulumi.Input

	ToInstanceTypeOutput() InstanceTypeOutput
	ToInstanceTypeOutputWithContext(context.Context) InstanceTypeOutput
}

var instanceTypePtrType = reflect.TypeOf((**InstanceType)(nil)).Elem()

type InstanceTypePtrInput interface {
	pulumi.Input

	ToInstanceTypePtrOutput() InstanceTypePtrOutput
	ToInstanceTypePtrOutputWithContext(context.Context) InstanceTypePtrOutput
}

type instanceTypePtr string

func InstanceTypePtr(v string) InstanceTypePtrInput {
	return (*instanceTypePtr)(&v)
}

func (*instanceTypePtr) ElementType() reflect.Type {
	return instanceTypePtrType
}

func (in *instanceTypePtr) ToInstanceTypePtrOutput() InstanceTypePtrOutput {
	return pulumi.ToOutput(in).(InstanceTypePtrOutput)
}

func (in *instanceTypePtr) ToInstanceTypePtrOutputWithContext(ctx context.Context) InstanceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InstanceTypePtrOutput)
}

// Immutable. The granularity (i.e. `MILLIS`) at which timestamps are stored in this table. Timestamps not matching the granularity will be rejected. If unspecified at creation time, the value will be set to `MILLIS`. Views: `SCHEMA_VIEW`, `FULL`.
type TableGranularity string

const (
	// The user did not specify a granularity. Should not be returned. When specified during table creation, MILLIS will be used.
	TableGranularityTimestampGranularityUnspecified = TableGranularity("TIMESTAMP_GRANULARITY_UNSPECIFIED")
	// The table keeps data versioned at a granularity of 1ms.
	TableGranularityMillis = TableGranularity("MILLIS")
)

func (TableGranularity) ElementType() reflect.Type {
	return reflect.TypeOf((*TableGranularity)(nil)).Elem()
}

func (e TableGranularity) ToTableGranularityOutput() TableGranularityOutput {
	return pulumi.ToOutput(e).(TableGranularityOutput)
}

func (e TableGranularity) ToTableGranularityOutputWithContext(ctx context.Context) TableGranularityOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TableGranularityOutput)
}

func (e TableGranularity) ToTableGranularityPtrOutput() TableGranularityPtrOutput {
	return e.ToTableGranularityPtrOutputWithContext(context.Background())
}

func (e TableGranularity) ToTableGranularityPtrOutputWithContext(ctx context.Context) TableGranularityPtrOutput {
	return TableGranularity(e).ToTableGranularityOutputWithContext(ctx).ToTableGranularityPtrOutputWithContext(ctx)
}

func (e TableGranularity) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TableGranularity) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TableGranularity) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TableGranularity) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TableGranularityOutput struct{ *pulumi.OutputState }

func (TableGranularityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableGranularity)(nil)).Elem()
}

func (o TableGranularityOutput) ToTableGranularityOutput() TableGranularityOutput {
	return o
}

func (o TableGranularityOutput) ToTableGranularityOutputWithContext(ctx context.Context) TableGranularityOutput {
	return o
}

func (o TableGranularityOutput) ToTableGranularityPtrOutput() TableGranularityPtrOutput {
	return o.ToTableGranularityPtrOutputWithContext(context.Background())
}

func (o TableGranularityOutput) ToTableGranularityPtrOutputWithContext(ctx context.Context) TableGranularityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TableGranularity) *TableGranularity {
		return &v
	}).(TableGranularityPtrOutput)
}

func (o TableGranularityOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TableGranularityOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TableGranularity) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TableGranularityOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TableGranularityOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TableGranularity) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TableGranularityPtrOutput struct{ *pulumi.OutputState }

func (TableGranularityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableGranularity)(nil)).Elem()
}

func (o TableGranularityPtrOutput) ToTableGranularityPtrOutput() TableGranularityPtrOutput {
	return o
}

func (o TableGranularityPtrOutput) ToTableGranularityPtrOutputWithContext(ctx context.Context) TableGranularityPtrOutput {
	return o
}

func (o TableGranularityPtrOutput) Elem() TableGranularityOutput {
	return o.ApplyT(func(v *TableGranularity) TableGranularity {
		if v != nil {
			return *v
		}
		var ret TableGranularity
		return ret
	}).(TableGranularityOutput)
}

func (o TableGranularityPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TableGranularityPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TableGranularity) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// TableGranularityInput is an input type that accepts TableGranularityArgs and TableGranularityOutput values.
// You can construct a concrete instance of `TableGranularityInput` via:
//
//	TableGranularityArgs{...}
type TableGranularityInput interface {
	pulumi.Input

	ToTableGranularityOutput() TableGranularityOutput
	ToTableGranularityOutputWithContext(context.Context) TableGranularityOutput
}

var tableGranularityPtrType = reflect.TypeOf((**TableGranularity)(nil)).Elem()

type TableGranularityPtrInput interface {
	pulumi.Input

	ToTableGranularityPtrOutput() TableGranularityPtrOutput
	ToTableGranularityPtrOutputWithContext(context.Context) TableGranularityPtrOutput
}

type tableGranularityPtr string

func TableGranularityPtr(v string) TableGranularityPtrInput {
	return (*tableGranularityPtr)(&v)
}

func (*tableGranularityPtr) ElementType() reflect.Type {
	return tableGranularityPtrType
}

func (in *tableGranularityPtr) ToTableGranularityPtrOutput() TableGranularityPtrOutput {
	return pulumi.ToOutput(in).(TableGranularityPtrOutput)
}

func (in *tableGranularityPtr) ToTableGranularityPtrOutputWithContext(ctx context.Context) TableGranularityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TableGranularityPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogConfigLogTypeInput)(nil)).Elem(), AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogConfigLogTypePtrInput)(nil)).Elem(), AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterDefaultStorageTypeInput)(nil)).Elem(), ClusterDefaultStorageType("STORAGE_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterDefaultStorageTypePtrInput)(nil)).Elem(), ClusterDefaultStorageType("STORAGE_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceTypeInput)(nil)).Elem(), InstanceType("TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceTypePtrInput)(nil)).Elem(), InstanceType("TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*TableGranularityInput)(nil)).Elem(), TableGranularity("TIMESTAMP_GRANULARITY_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*TableGranularityPtrInput)(nil)).Elem(), TableGranularity("TIMESTAMP_GRANULARITY_UNSPECIFIED"))
	pulumi.RegisterOutputType(AuditLogConfigLogTypeOutput{})
	pulumi.RegisterOutputType(AuditLogConfigLogTypePtrOutput{})
	pulumi.RegisterOutputType(ClusterDefaultStorageTypeOutput{})
	pulumi.RegisterOutputType(ClusterDefaultStorageTypePtrOutput{})
	pulumi.RegisterOutputType(InstanceTypeOutput{})
	pulumi.RegisterOutputType(InstanceTypePtrOutput{})
	pulumi.RegisterOutputType(TableGranularityOutput{})
	pulumi.RegisterOutputType(TableGranularityPtrOutput{})
}
