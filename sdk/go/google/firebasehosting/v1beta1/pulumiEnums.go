// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Required. The redirect status code.
type DomainRedirectType string

const (
	// The default redirect type; should not be intentionlly used.
	DomainRedirectTypeRedirectTypeUnspecified = DomainRedirectType("REDIRECT_TYPE_UNSPECIFIED")
	// The redirect will respond with an HTTP status code of `301 Moved Permanently`.
	DomainRedirectTypeMovedPermanently = DomainRedirectType("MOVED_PERMANENTLY")
)

func (DomainRedirectType) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainRedirectType)(nil)).Elem()
}

func (e DomainRedirectType) ToDomainRedirectTypeOutput() DomainRedirectTypeOutput {
	return pulumi.ToOutput(e).(DomainRedirectTypeOutput)
}

func (e DomainRedirectType) ToDomainRedirectTypeOutputWithContext(ctx context.Context) DomainRedirectTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DomainRedirectTypeOutput)
}

func (e DomainRedirectType) ToDomainRedirectTypePtrOutput() DomainRedirectTypePtrOutput {
	return e.ToDomainRedirectTypePtrOutputWithContext(context.Background())
}

func (e DomainRedirectType) ToDomainRedirectTypePtrOutputWithContext(ctx context.Context) DomainRedirectTypePtrOutput {
	return DomainRedirectType(e).ToDomainRedirectTypeOutputWithContext(ctx).ToDomainRedirectTypePtrOutputWithContext(ctx)
}

func (e DomainRedirectType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DomainRedirectType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DomainRedirectType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DomainRedirectType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DomainRedirectTypeOutput struct{ *pulumi.OutputState }

func (DomainRedirectTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DomainRedirectType)(nil)).Elem()
}

func (o DomainRedirectTypeOutput) ToDomainRedirectTypeOutput() DomainRedirectTypeOutput {
	return o
}

func (o DomainRedirectTypeOutput) ToDomainRedirectTypeOutputWithContext(ctx context.Context) DomainRedirectTypeOutput {
	return o
}

func (o DomainRedirectTypeOutput) ToDomainRedirectTypePtrOutput() DomainRedirectTypePtrOutput {
	return o.ToDomainRedirectTypePtrOutputWithContext(context.Background())
}

func (o DomainRedirectTypeOutput) ToDomainRedirectTypePtrOutputWithContext(ctx context.Context) DomainRedirectTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DomainRedirectType) *DomainRedirectType {
		return &v
	}).(DomainRedirectTypePtrOutput)
}

func (o DomainRedirectTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DomainRedirectTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DomainRedirectType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DomainRedirectTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DomainRedirectTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DomainRedirectType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DomainRedirectTypePtrOutput struct{ *pulumi.OutputState }

func (DomainRedirectTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainRedirectType)(nil)).Elem()
}

func (o DomainRedirectTypePtrOutput) ToDomainRedirectTypePtrOutput() DomainRedirectTypePtrOutput {
	return o
}

func (o DomainRedirectTypePtrOutput) ToDomainRedirectTypePtrOutputWithContext(ctx context.Context) DomainRedirectTypePtrOutput {
	return o
}

func (o DomainRedirectTypePtrOutput) Elem() DomainRedirectTypeOutput {
	return o.ApplyT(func(v *DomainRedirectType) DomainRedirectType {
		if v != nil {
			return *v
		}
		var ret DomainRedirectType
		return ret
	}).(DomainRedirectTypeOutput)
}

func (o DomainRedirectTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DomainRedirectTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DomainRedirectType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DomainRedirectTypeInput is an input type that accepts DomainRedirectTypeArgs and DomainRedirectTypeOutput values.
// You can construct a concrete instance of `DomainRedirectTypeInput` via:
//
//	DomainRedirectTypeArgs{...}
type DomainRedirectTypeInput interface {
	pulumi.Input

	ToDomainRedirectTypeOutput() DomainRedirectTypeOutput
	ToDomainRedirectTypeOutputWithContext(context.Context) DomainRedirectTypeOutput
}

var domainRedirectTypePtrType = reflect.TypeOf((**DomainRedirectType)(nil)).Elem()

type DomainRedirectTypePtrInput interface {
	pulumi.Input

	ToDomainRedirectTypePtrOutput() DomainRedirectTypePtrOutput
	ToDomainRedirectTypePtrOutputWithContext(context.Context) DomainRedirectTypePtrOutput
}

type domainRedirectTypePtr string

func DomainRedirectTypePtr(v string) DomainRedirectTypePtrInput {
	return (*domainRedirectTypePtr)(&v)
}

func (*domainRedirectTypePtr) ElementType() reflect.Type {
	return domainRedirectTypePtrType
}

func (in *domainRedirectTypePtr) ToDomainRedirectTypePtrOutput() DomainRedirectTypePtrOutput {
	return pulumi.ToOutput(in).(DomainRedirectTypePtrOutput)
}

func (in *domainRedirectTypePtr) ToDomainRedirectTypePtrOutputWithContext(ctx context.Context) DomainRedirectTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DomainRedirectTypePtrOutput)
}

// Explains the reason for the release. Specify a value for this field only when creating a `SITE_DISABLE` type release.
type ReleaseType string

const (
	// An unspecified type. Indicates that a version was released. This is the default value when no other `type` is explicitly specified.
	ReleaseTypeTypeUnspecified = ReleaseType("TYPE_UNSPECIFIED")
	// A version was uploaded to Firebase Hosting and released.
	ReleaseTypeDeploy = ReleaseType("DEPLOY")
	// The release points back to a previously deployed version.
	ReleaseTypeRollback = ReleaseType("ROLLBACK")
	// The release prevents the site from serving content. Firebase Hosting acts as if the site never existed.
	ReleaseTypeSiteDisable = ReleaseType("SITE_DISABLE")
)

func (ReleaseType) ElementType() reflect.Type {
	return reflect.TypeOf((*ReleaseType)(nil)).Elem()
}

func (e ReleaseType) ToReleaseTypeOutput() ReleaseTypeOutput {
	return pulumi.ToOutput(e).(ReleaseTypeOutput)
}

func (e ReleaseType) ToReleaseTypeOutputWithContext(ctx context.Context) ReleaseTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ReleaseTypeOutput)
}

func (e ReleaseType) ToReleaseTypePtrOutput() ReleaseTypePtrOutput {
	return e.ToReleaseTypePtrOutputWithContext(context.Background())
}

func (e ReleaseType) ToReleaseTypePtrOutputWithContext(ctx context.Context) ReleaseTypePtrOutput {
	return ReleaseType(e).ToReleaseTypeOutputWithContext(ctx).ToReleaseTypePtrOutputWithContext(ctx)
}

func (e ReleaseType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReleaseType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReleaseType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ReleaseType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ReleaseTypeOutput struct{ *pulumi.OutputState }

func (ReleaseTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReleaseType)(nil)).Elem()
}

func (o ReleaseTypeOutput) ToReleaseTypeOutput() ReleaseTypeOutput {
	return o
}

func (o ReleaseTypeOutput) ToReleaseTypeOutputWithContext(ctx context.Context) ReleaseTypeOutput {
	return o
}

func (o ReleaseTypeOutput) ToReleaseTypePtrOutput() ReleaseTypePtrOutput {
	return o.ToReleaseTypePtrOutputWithContext(context.Background())
}

func (o ReleaseTypeOutput) ToReleaseTypePtrOutputWithContext(ctx context.Context) ReleaseTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReleaseType) *ReleaseType {
		return &v
	}).(ReleaseTypePtrOutput)
}

func (o ReleaseTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ReleaseTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReleaseType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ReleaseTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReleaseTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReleaseType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ReleaseTypePtrOutput struct{ *pulumi.OutputState }

func (ReleaseTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReleaseType)(nil)).Elem()
}

func (o ReleaseTypePtrOutput) ToReleaseTypePtrOutput() ReleaseTypePtrOutput {
	return o
}

func (o ReleaseTypePtrOutput) ToReleaseTypePtrOutputWithContext(ctx context.Context) ReleaseTypePtrOutput {
	return o
}

func (o ReleaseTypePtrOutput) Elem() ReleaseTypeOutput {
	return o.ApplyT(func(v *ReleaseType) ReleaseType {
		if v != nil {
			return *v
		}
		var ret ReleaseType
		return ret
	}).(ReleaseTypeOutput)
}

func (o ReleaseTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReleaseTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ReleaseType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ReleaseTypeInput is an input type that accepts ReleaseTypeArgs and ReleaseTypeOutput values.
// You can construct a concrete instance of `ReleaseTypeInput` via:
//
//	ReleaseTypeArgs{...}
type ReleaseTypeInput interface {
	pulumi.Input

	ToReleaseTypeOutput() ReleaseTypeOutput
	ToReleaseTypeOutputWithContext(context.Context) ReleaseTypeOutput
}

var releaseTypePtrType = reflect.TypeOf((**ReleaseType)(nil)).Elem()

type ReleaseTypePtrInput interface {
	pulumi.Input

	ToReleaseTypePtrOutput() ReleaseTypePtrOutput
	ToReleaseTypePtrOutputWithContext(context.Context) ReleaseTypePtrOutput
}

type releaseTypePtr string

func ReleaseTypePtr(v string) ReleaseTypePtrInput {
	return (*releaseTypePtr)(&v)
}

func (*releaseTypePtr) ElementType() reflect.Type {
	return releaseTypePtrType
}

func (in *releaseTypePtr) ToReleaseTypePtrOutput() ReleaseTypePtrOutput {
	return pulumi.ToOutput(in).(ReleaseTypePtrOutput)
}

func (in *releaseTypePtr) ToReleaseTypePtrOutputWithContext(ctx context.Context) ReleaseTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ReleaseTypePtrOutput)
}

// How to handle well known App Association files.
type ServingConfigAppAssociation string

const (
	// The app association files will be automatically created from the apps that exist in the Firebase project.
	ServingConfigAppAssociationAuto = ServingConfigAppAssociation("AUTO")
	// No special handling of the app association files will occur, these paths will result in a 404 unless caught with a Rewrite.
	ServingConfigAppAssociationNone = ServingConfigAppAssociation("NONE")
)

func (ServingConfigAppAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*ServingConfigAppAssociation)(nil)).Elem()
}

func (e ServingConfigAppAssociation) ToServingConfigAppAssociationOutput() ServingConfigAppAssociationOutput {
	return pulumi.ToOutput(e).(ServingConfigAppAssociationOutput)
}

func (e ServingConfigAppAssociation) ToServingConfigAppAssociationOutputWithContext(ctx context.Context) ServingConfigAppAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServingConfigAppAssociationOutput)
}

func (e ServingConfigAppAssociation) ToServingConfigAppAssociationPtrOutput() ServingConfigAppAssociationPtrOutput {
	return e.ToServingConfigAppAssociationPtrOutputWithContext(context.Background())
}

func (e ServingConfigAppAssociation) ToServingConfigAppAssociationPtrOutputWithContext(ctx context.Context) ServingConfigAppAssociationPtrOutput {
	return ServingConfigAppAssociation(e).ToServingConfigAppAssociationOutputWithContext(ctx).ToServingConfigAppAssociationPtrOutputWithContext(ctx)
}

func (e ServingConfigAppAssociation) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServingConfigAppAssociation) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServingConfigAppAssociation) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServingConfigAppAssociation) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServingConfigAppAssociationOutput struct{ *pulumi.OutputState }

func (ServingConfigAppAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServingConfigAppAssociation)(nil)).Elem()
}

func (o ServingConfigAppAssociationOutput) ToServingConfigAppAssociationOutput() ServingConfigAppAssociationOutput {
	return o
}

func (o ServingConfigAppAssociationOutput) ToServingConfigAppAssociationOutputWithContext(ctx context.Context) ServingConfigAppAssociationOutput {
	return o
}

func (o ServingConfigAppAssociationOutput) ToServingConfigAppAssociationPtrOutput() ServingConfigAppAssociationPtrOutput {
	return o.ToServingConfigAppAssociationPtrOutputWithContext(context.Background())
}

func (o ServingConfigAppAssociationOutput) ToServingConfigAppAssociationPtrOutputWithContext(ctx context.Context) ServingConfigAppAssociationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServingConfigAppAssociation) *ServingConfigAppAssociation {
		return &v
	}).(ServingConfigAppAssociationPtrOutput)
}

func (o ServingConfigAppAssociationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServingConfigAppAssociationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServingConfigAppAssociation) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServingConfigAppAssociationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServingConfigAppAssociationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServingConfigAppAssociation) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServingConfigAppAssociationPtrOutput struct{ *pulumi.OutputState }

func (ServingConfigAppAssociationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServingConfigAppAssociation)(nil)).Elem()
}

func (o ServingConfigAppAssociationPtrOutput) ToServingConfigAppAssociationPtrOutput() ServingConfigAppAssociationPtrOutput {
	return o
}

func (o ServingConfigAppAssociationPtrOutput) ToServingConfigAppAssociationPtrOutputWithContext(ctx context.Context) ServingConfigAppAssociationPtrOutput {
	return o
}

func (o ServingConfigAppAssociationPtrOutput) Elem() ServingConfigAppAssociationOutput {
	return o.ApplyT(func(v *ServingConfigAppAssociation) ServingConfigAppAssociation {
		if v != nil {
			return *v
		}
		var ret ServingConfigAppAssociation
		return ret
	}).(ServingConfigAppAssociationOutput)
}

func (o ServingConfigAppAssociationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServingConfigAppAssociationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServingConfigAppAssociation) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ServingConfigAppAssociationInput is an input type that accepts ServingConfigAppAssociationArgs and ServingConfigAppAssociationOutput values.
// You can construct a concrete instance of `ServingConfigAppAssociationInput` via:
//
//	ServingConfigAppAssociationArgs{...}
type ServingConfigAppAssociationInput interface {
	pulumi.Input

	ToServingConfigAppAssociationOutput() ServingConfigAppAssociationOutput
	ToServingConfigAppAssociationOutputWithContext(context.Context) ServingConfigAppAssociationOutput
}

var servingConfigAppAssociationPtrType = reflect.TypeOf((**ServingConfigAppAssociation)(nil)).Elem()

type ServingConfigAppAssociationPtrInput interface {
	pulumi.Input

	ToServingConfigAppAssociationPtrOutput() ServingConfigAppAssociationPtrOutput
	ToServingConfigAppAssociationPtrOutputWithContext(context.Context) ServingConfigAppAssociationPtrOutput
}

type servingConfigAppAssociationPtr string

func ServingConfigAppAssociationPtr(v string) ServingConfigAppAssociationPtrInput {
	return (*servingConfigAppAssociationPtr)(&v)
}

func (*servingConfigAppAssociationPtr) ElementType() reflect.Type {
	return servingConfigAppAssociationPtrType
}

func (in *servingConfigAppAssociationPtr) ToServingConfigAppAssociationPtrOutput() ServingConfigAppAssociationPtrOutput {
	return pulumi.ToOutput(in).(ServingConfigAppAssociationPtrOutput)
}

func (in *servingConfigAppAssociationPtr) ToServingConfigAppAssociationPtrOutputWithContext(ctx context.Context) ServingConfigAppAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServingConfigAppAssociationPtrOutput)
}

// Defines how to handle a trailing slash in the URL path.
type ServingConfigTrailingSlashBehavior string

const (
	// No behavior is specified. Files are served at their exact location only, and trailing slashes are only added to directory indexes.
	ServingConfigTrailingSlashBehaviorTrailingSlashBehaviorUnspecified = ServingConfigTrailingSlashBehavior("TRAILING_SLASH_BEHAVIOR_UNSPECIFIED")
	// Trailing slashes are _added_ to directory indexes as well as to any URL path not ending in a file extension.
	ServingConfigTrailingSlashBehaviorAdd = ServingConfigTrailingSlashBehavior("ADD")
	// Trailing slashes are _removed_ from directory indexes as well as from any URL path not ending in a file extension.
	ServingConfigTrailingSlashBehaviorRemove = ServingConfigTrailingSlashBehavior("REMOVE")
)

func (ServingConfigTrailingSlashBehavior) ElementType() reflect.Type {
	return reflect.TypeOf((*ServingConfigTrailingSlashBehavior)(nil)).Elem()
}

func (e ServingConfigTrailingSlashBehavior) ToServingConfigTrailingSlashBehaviorOutput() ServingConfigTrailingSlashBehaviorOutput {
	return pulumi.ToOutput(e).(ServingConfigTrailingSlashBehaviorOutput)
}

func (e ServingConfigTrailingSlashBehavior) ToServingConfigTrailingSlashBehaviorOutputWithContext(ctx context.Context) ServingConfigTrailingSlashBehaviorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServingConfigTrailingSlashBehaviorOutput)
}

func (e ServingConfigTrailingSlashBehavior) ToServingConfigTrailingSlashBehaviorPtrOutput() ServingConfigTrailingSlashBehaviorPtrOutput {
	return e.ToServingConfigTrailingSlashBehaviorPtrOutputWithContext(context.Background())
}

func (e ServingConfigTrailingSlashBehavior) ToServingConfigTrailingSlashBehaviorPtrOutputWithContext(ctx context.Context) ServingConfigTrailingSlashBehaviorPtrOutput {
	return ServingConfigTrailingSlashBehavior(e).ToServingConfigTrailingSlashBehaviorOutputWithContext(ctx).ToServingConfigTrailingSlashBehaviorPtrOutputWithContext(ctx)
}

func (e ServingConfigTrailingSlashBehavior) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServingConfigTrailingSlashBehavior) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServingConfigTrailingSlashBehavior) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServingConfigTrailingSlashBehavior) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServingConfigTrailingSlashBehaviorOutput struct{ *pulumi.OutputState }

func (ServingConfigTrailingSlashBehaviorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServingConfigTrailingSlashBehavior)(nil)).Elem()
}

func (o ServingConfigTrailingSlashBehaviorOutput) ToServingConfigTrailingSlashBehaviorOutput() ServingConfigTrailingSlashBehaviorOutput {
	return o
}

func (o ServingConfigTrailingSlashBehaviorOutput) ToServingConfigTrailingSlashBehaviorOutputWithContext(ctx context.Context) ServingConfigTrailingSlashBehaviorOutput {
	return o
}

func (o ServingConfigTrailingSlashBehaviorOutput) ToServingConfigTrailingSlashBehaviorPtrOutput() ServingConfigTrailingSlashBehaviorPtrOutput {
	return o.ToServingConfigTrailingSlashBehaviorPtrOutputWithContext(context.Background())
}

func (o ServingConfigTrailingSlashBehaviorOutput) ToServingConfigTrailingSlashBehaviorPtrOutputWithContext(ctx context.Context) ServingConfigTrailingSlashBehaviorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServingConfigTrailingSlashBehavior) *ServingConfigTrailingSlashBehavior {
		return &v
	}).(ServingConfigTrailingSlashBehaviorPtrOutput)
}

func (o ServingConfigTrailingSlashBehaviorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServingConfigTrailingSlashBehaviorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServingConfigTrailingSlashBehavior) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServingConfigTrailingSlashBehaviorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServingConfigTrailingSlashBehaviorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServingConfigTrailingSlashBehavior) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServingConfigTrailingSlashBehaviorPtrOutput struct{ *pulumi.OutputState }

func (ServingConfigTrailingSlashBehaviorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServingConfigTrailingSlashBehavior)(nil)).Elem()
}

func (o ServingConfigTrailingSlashBehaviorPtrOutput) ToServingConfigTrailingSlashBehaviorPtrOutput() ServingConfigTrailingSlashBehaviorPtrOutput {
	return o
}

func (o ServingConfigTrailingSlashBehaviorPtrOutput) ToServingConfigTrailingSlashBehaviorPtrOutputWithContext(ctx context.Context) ServingConfigTrailingSlashBehaviorPtrOutput {
	return o
}

func (o ServingConfigTrailingSlashBehaviorPtrOutput) Elem() ServingConfigTrailingSlashBehaviorOutput {
	return o.ApplyT(func(v *ServingConfigTrailingSlashBehavior) ServingConfigTrailingSlashBehavior {
		if v != nil {
			return *v
		}
		var ret ServingConfigTrailingSlashBehavior
		return ret
	}).(ServingConfigTrailingSlashBehaviorOutput)
}

func (o ServingConfigTrailingSlashBehaviorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServingConfigTrailingSlashBehaviorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServingConfigTrailingSlashBehavior) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ServingConfigTrailingSlashBehaviorInput is an input type that accepts ServingConfigTrailingSlashBehaviorArgs and ServingConfigTrailingSlashBehaviorOutput values.
// You can construct a concrete instance of `ServingConfigTrailingSlashBehaviorInput` via:
//
//	ServingConfigTrailingSlashBehaviorArgs{...}
type ServingConfigTrailingSlashBehaviorInput interface {
	pulumi.Input

	ToServingConfigTrailingSlashBehaviorOutput() ServingConfigTrailingSlashBehaviorOutput
	ToServingConfigTrailingSlashBehaviorOutputWithContext(context.Context) ServingConfigTrailingSlashBehaviorOutput
}

var servingConfigTrailingSlashBehaviorPtrType = reflect.TypeOf((**ServingConfigTrailingSlashBehavior)(nil)).Elem()

type ServingConfigTrailingSlashBehaviorPtrInput interface {
	pulumi.Input

	ToServingConfigTrailingSlashBehaviorPtrOutput() ServingConfigTrailingSlashBehaviorPtrOutput
	ToServingConfigTrailingSlashBehaviorPtrOutputWithContext(context.Context) ServingConfigTrailingSlashBehaviorPtrOutput
}

type servingConfigTrailingSlashBehaviorPtr string

func ServingConfigTrailingSlashBehaviorPtr(v string) ServingConfigTrailingSlashBehaviorPtrInput {
	return (*servingConfigTrailingSlashBehaviorPtr)(&v)
}

func (*servingConfigTrailingSlashBehaviorPtr) ElementType() reflect.Type {
	return servingConfigTrailingSlashBehaviorPtrType
}

func (in *servingConfigTrailingSlashBehaviorPtr) ToServingConfigTrailingSlashBehaviorPtrOutput() ServingConfigTrailingSlashBehaviorPtrOutput {
	return pulumi.ToOutput(in).(ServingConfigTrailingSlashBehaviorPtrOutput)
}

func (in *servingConfigTrailingSlashBehaviorPtr) ToServingConfigTrailingSlashBehaviorPtrOutputWithContext(ctx context.Context) ServingConfigTrailingSlashBehaviorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServingConfigTrailingSlashBehaviorPtrOutput)
}

// The deploy status of the version. For a successful deploy, call [`CreateVersion`](sites.versions/create) to make a new version (`CREATED` status), [upload all desired files](sites.versions/populateFiles) to the version, then [update](sites.versions/patch) the version to the `FINALIZED` status. Note that if you leave the version in the `CREATED` state for more than 12 hours, the system will automatically mark the version as `ABANDONED`. You can also change the status of a version to `DELETED` by calling [`DeleteVersion`](sites.versions/delete).
type VersionStatus string

const (
	// The default status; should not be intentionally used.
	VersionStatusVersionStatusUnspecified = VersionStatus("VERSION_STATUS_UNSPECIFIED")
	// The version has been created, and content is currently being added to the version.
	VersionStatusCreated = VersionStatus("CREATED")
	// All content has been added to the version, and the version can no longer be changed.
	VersionStatusFinalized = VersionStatus("FINALIZED")
	// The version has been deleted.
	VersionStatusDeleted = VersionStatus("DELETED")
	// The version was not updated to `FINALIZED` within 12 hours and was automatically deleted.
	VersionStatusAbandoned = VersionStatus("ABANDONED")
	// The version is outside the site-configured limit for the number of retained versions, so the version's content is scheduled for deletion.
	VersionStatusExpired = VersionStatus("EXPIRED")
	// The version is being cloned from another version. All content is still being copied over.
	VersionStatusCloning = VersionStatus("CLONING")
)

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainRedirectTypeInput)(nil)).Elem(), DomainRedirectType("REDIRECT_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*DomainRedirectTypePtrInput)(nil)).Elem(), DomainRedirectType("REDIRECT_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseTypeInput)(nil)).Elem(), ReleaseType("TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseTypePtrInput)(nil)).Elem(), ReleaseType("TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ServingConfigAppAssociationInput)(nil)).Elem(), ServingConfigAppAssociation("AUTO"))
	pulumi.RegisterInputType(reflect.TypeOf((*ServingConfigAppAssociationPtrInput)(nil)).Elem(), ServingConfigAppAssociation("AUTO"))
	pulumi.RegisterInputType(reflect.TypeOf((*ServingConfigTrailingSlashBehaviorInput)(nil)).Elem(), ServingConfigTrailingSlashBehavior("TRAILING_SLASH_BEHAVIOR_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ServingConfigTrailingSlashBehaviorPtrInput)(nil)).Elem(), ServingConfigTrailingSlashBehavior("TRAILING_SLASH_BEHAVIOR_UNSPECIFIED"))
	pulumi.RegisterOutputType(DomainRedirectTypeOutput{})
	pulumi.RegisterOutputType(DomainRedirectTypePtrOutput{})
	pulumi.RegisterOutputType(ReleaseTypeOutput{})
	pulumi.RegisterOutputType(ReleaseTypePtrOutput{})
	pulumi.RegisterOutputType(ServingConfigAppAssociationOutput{})
	pulumi.RegisterOutputType(ServingConfigAppAssociationPtrOutput{})
	pulumi.RegisterOutputType(ServingConfigTrailingSlashBehaviorOutput{})
	pulumi.RegisterOutputType(ServingConfigTrailingSlashBehaviorPtrOutput{})
}
