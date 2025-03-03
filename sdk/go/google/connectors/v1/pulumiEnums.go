// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

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

// The type of authentication configured.
type AuthConfigAuthType string

const (
	// Authentication type not specified.
	AuthConfigAuthTypeAuthTypeUnspecified = AuthConfigAuthType("AUTH_TYPE_UNSPECIFIED")
	// Username and Password Authentication.
	AuthConfigAuthTypeUserPassword = AuthConfigAuthType("USER_PASSWORD")
	// JSON Web Token (JWT) Profile for Oauth 2.0 Authorization Grant based authentication
	AuthConfigAuthTypeOauth2JwtBearer = AuthConfigAuthType("OAUTH2_JWT_BEARER")
	// Oauth 2.0 Client Credentials Grant Authentication
	AuthConfigAuthTypeOauth2ClientCredentials = AuthConfigAuthType("OAUTH2_CLIENT_CREDENTIALS")
	// SSH Public Key Authentication
	AuthConfigAuthTypeSshPublicKey = AuthConfigAuthType("SSH_PUBLIC_KEY")
	// Oauth 2.0 Authorization Code Flow
	AuthConfigAuthTypeOauth2AuthCodeFlow = AuthConfigAuthType("OAUTH2_AUTH_CODE_FLOW")
)

func (AuthConfigAuthType) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthConfigAuthType)(nil)).Elem()
}

func (e AuthConfigAuthType) ToAuthConfigAuthTypeOutput() AuthConfigAuthTypeOutput {
	return pulumi.ToOutput(e).(AuthConfigAuthTypeOutput)
}

func (e AuthConfigAuthType) ToAuthConfigAuthTypeOutputWithContext(ctx context.Context) AuthConfigAuthTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AuthConfigAuthTypeOutput)
}

func (e AuthConfigAuthType) ToAuthConfigAuthTypePtrOutput() AuthConfigAuthTypePtrOutput {
	return e.ToAuthConfigAuthTypePtrOutputWithContext(context.Background())
}

func (e AuthConfigAuthType) ToAuthConfigAuthTypePtrOutputWithContext(ctx context.Context) AuthConfigAuthTypePtrOutput {
	return AuthConfigAuthType(e).ToAuthConfigAuthTypeOutputWithContext(ctx).ToAuthConfigAuthTypePtrOutputWithContext(ctx)
}

func (e AuthConfigAuthType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthConfigAuthType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthConfigAuthType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AuthConfigAuthType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AuthConfigAuthTypeOutput struct{ *pulumi.OutputState }

func (AuthConfigAuthTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthConfigAuthType)(nil)).Elem()
}

func (o AuthConfigAuthTypeOutput) ToAuthConfigAuthTypeOutput() AuthConfigAuthTypeOutput {
	return o
}

func (o AuthConfigAuthTypeOutput) ToAuthConfigAuthTypeOutputWithContext(ctx context.Context) AuthConfigAuthTypeOutput {
	return o
}

func (o AuthConfigAuthTypeOutput) ToAuthConfigAuthTypePtrOutput() AuthConfigAuthTypePtrOutput {
	return o.ToAuthConfigAuthTypePtrOutputWithContext(context.Background())
}

func (o AuthConfigAuthTypeOutput) ToAuthConfigAuthTypePtrOutputWithContext(ctx context.Context) AuthConfigAuthTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthConfigAuthType) *AuthConfigAuthType {
		return &v
	}).(AuthConfigAuthTypePtrOutput)
}

func (o AuthConfigAuthTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AuthConfigAuthTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthConfigAuthType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AuthConfigAuthTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthConfigAuthTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthConfigAuthType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AuthConfigAuthTypePtrOutput struct{ *pulumi.OutputState }

func (AuthConfigAuthTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthConfigAuthType)(nil)).Elem()
}

func (o AuthConfigAuthTypePtrOutput) ToAuthConfigAuthTypePtrOutput() AuthConfigAuthTypePtrOutput {
	return o
}

func (o AuthConfigAuthTypePtrOutput) ToAuthConfigAuthTypePtrOutputWithContext(ctx context.Context) AuthConfigAuthTypePtrOutput {
	return o
}

func (o AuthConfigAuthTypePtrOutput) Elem() AuthConfigAuthTypeOutput {
	return o.ApplyT(func(v *AuthConfigAuthType) AuthConfigAuthType {
		if v != nil {
			return *v
		}
		var ret AuthConfigAuthType
		return ret
	}).(AuthConfigAuthTypeOutput)
}

func (o AuthConfigAuthTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthConfigAuthTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AuthConfigAuthType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AuthConfigAuthTypeInput is an input type that accepts AuthConfigAuthTypeArgs and AuthConfigAuthTypeOutput values.
// You can construct a concrete instance of `AuthConfigAuthTypeInput` via:
//
//	AuthConfigAuthTypeArgs{...}
type AuthConfigAuthTypeInput interface {
	pulumi.Input

	ToAuthConfigAuthTypeOutput() AuthConfigAuthTypeOutput
	ToAuthConfigAuthTypeOutputWithContext(context.Context) AuthConfigAuthTypeOutput
}

var authConfigAuthTypePtrType = reflect.TypeOf((**AuthConfigAuthType)(nil)).Elem()

type AuthConfigAuthTypePtrInput interface {
	pulumi.Input

	ToAuthConfigAuthTypePtrOutput() AuthConfigAuthTypePtrOutput
	ToAuthConfigAuthTypePtrOutputWithContext(context.Context) AuthConfigAuthTypePtrOutput
}

type authConfigAuthTypePtr string

func AuthConfigAuthTypePtr(v string) AuthConfigAuthTypePtrInput {
	return (*authConfigAuthTypePtr)(&v)
}

func (*authConfigAuthTypePtr) ElementType() reflect.Type {
	return authConfigAuthTypePtrType
}

func (in *authConfigAuthTypePtr) ToAuthConfigAuthTypePtrOutput() AuthConfigAuthTypePtrOutput {
	return pulumi.ToOutput(in).(AuthConfigAuthTypePtrOutput)
}

func (in *authConfigAuthTypePtr) ToAuthConfigAuthTypePtrOutputWithContext(ctx context.Context) AuthConfigAuthTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AuthConfigAuthTypePtrOutput)
}

// Type of Client Cert (PEM/JKS/.. etc.)
type SslConfigClientCertType string

const (
	// Cert type unspecified.
	SslConfigClientCertTypeCertTypeUnspecified = SslConfigClientCertType("CERT_TYPE_UNSPECIFIED")
	// Privacy Enhanced Mail (PEM) Type
	SslConfigClientCertTypePem = SslConfigClientCertType("PEM")
)

func (SslConfigClientCertType) ElementType() reflect.Type {
	return reflect.TypeOf((*SslConfigClientCertType)(nil)).Elem()
}

func (e SslConfigClientCertType) ToSslConfigClientCertTypeOutput() SslConfigClientCertTypeOutput {
	return pulumi.ToOutput(e).(SslConfigClientCertTypeOutput)
}

func (e SslConfigClientCertType) ToSslConfigClientCertTypeOutputWithContext(ctx context.Context) SslConfigClientCertTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SslConfigClientCertTypeOutput)
}

func (e SslConfigClientCertType) ToSslConfigClientCertTypePtrOutput() SslConfigClientCertTypePtrOutput {
	return e.ToSslConfigClientCertTypePtrOutputWithContext(context.Background())
}

func (e SslConfigClientCertType) ToSslConfigClientCertTypePtrOutputWithContext(ctx context.Context) SslConfigClientCertTypePtrOutput {
	return SslConfigClientCertType(e).ToSslConfigClientCertTypeOutputWithContext(ctx).ToSslConfigClientCertTypePtrOutputWithContext(ctx)
}

func (e SslConfigClientCertType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslConfigClientCertType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslConfigClientCertType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SslConfigClientCertType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SslConfigClientCertTypeOutput struct{ *pulumi.OutputState }

func (SslConfigClientCertTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SslConfigClientCertType)(nil)).Elem()
}

func (o SslConfigClientCertTypeOutput) ToSslConfigClientCertTypeOutput() SslConfigClientCertTypeOutput {
	return o
}

func (o SslConfigClientCertTypeOutput) ToSslConfigClientCertTypeOutputWithContext(ctx context.Context) SslConfigClientCertTypeOutput {
	return o
}

func (o SslConfigClientCertTypeOutput) ToSslConfigClientCertTypePtrOutput() SslConfigClientCertTypePtrOutput {
	return o.ToSslConfigClientCertTypePtrOutputWithContext(context.Background())
}

func (o SslConfigClientCertTypeOutput) ToSslConfigClientCertTypePtrOutputWithContext(ctx context.Context) SslConfigClientCertTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SslConfigClientCertType) *SslConfigClientCertType {
		return &v
	}).(SslConfigClientCertTypePtrOutput)
}

func (o SslConfigClientCertTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SslConfigClientCertTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslConfigClientCertType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SslConfigClientCertTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslConfigClientCertTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslConfigClientCertType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SslConfigClientCertTypePtrOutput struct{ *pulumi.OutputState }

func (SslConfigClientCertTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslConfigClientCertType)(nil)).Elem()
}

func (o SslConfigClientCertTypePtrOutput) ToSslConfigClientCertTypePtrOutput() SslConfigClientCertTypePtrOutput {
	return o
}

func (o SslConfigClientCertTypePtrOutput) ToSslConfigClientCertTypePtrOutputWithContext(ctx context.Context) SslConfigClientCertTypePtrOutput {
	return o
}

func (o SslConfigClientCertTypePtrOutput) Elem() SslConfigClientCertTypeOutput {
	return o.ApplyT(func(v *SslConfigClientCertType) SslConfigClientCertType {
		if v != nil {
			return *v
		}
		var ret SslConfigClientCertType
		return ret
	}).(SslConfigClientCertTypeOutput)
}

func (o SslConfigClientCertTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslConfigClientCertTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SslConfigClientCertType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SslConfigClientCertTypeInput is an input type that accepts SslConfigClientCertTypeArgs and SslConfigClientCertTypeOutput values.
// You can construct a concrete instance of `SslConfigClientCertTypeInput` via:
//
//	SslConfigClientCertTypeArgs{...}
type SslConfigClientCertTypeInput interface {
	pulumi.Input

	ToSslConfigClientCertTypeOutput() SslConfigClientCertTypeOutput
	ToSslConfigClientCertTypeOutputWithContext(context.Context) SslConfigClientCertTypeOutput
}

var sslConfigClientCertTypePtrType = reflect.TypeOf((**SslConfigClientCertType)(nil)).Elem()

type SslConfigClientCertTypePtrInput interface {
	pulumi.Input

	ToSslConfigClientCertTypePtrOutput() SslConfigClientCertTypePtrOutput
	ToSslConfigClientCertTypePtrOutputWithContext(context.Context) SslConfigClientCertTypePtrOutput
}

type sslConfigClientCertTypePtr string

func SslConfigClientCertTypePtr(v string) SslConfigClientCertTypePtrInput {
	return (*sslConfigClientCertTypePtr)(&v)
}

func (*sslConfigClientCertTypePtr) ElementType() reflect.Type {
	return sslConfigClientCertTypePtrType
}

func (in *sslConfigClientCertTypePtr) ToSslConfigClientCertTypePtrOutput() SslConfigClientCertTypePtrOutput {
	return pulumi.ToOutput(in).(SslConfigClientCertTypePtrOutput)
}

func (in *sslConfigClientCertTypePtr) ToSslConfigClientCertTypePtrOutputWithContext(ctx context.Context) SslConfigClientCertTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SslConfigClientCertTypePtrOutput)
}

// Type of Server Cert (PEM/JKS/.. etc.)
type SslConfigServerCertType string

const (
	// Cert type unspecified.
	SslConfigServerCertTypeCertTypeUnspecified = SslConfigServerCertType("CERT_TYPE_UNSPECIFIED")
	// Privacy Enhanced Mail (PEM) Type
	SslConfigServerCertTypePem = SslConfigServerCertType("PEM")
)

func (SslConfigServerCertType) ElementType() reflect.Type {
	return reflect.TypeOf((*SslConfigServerCertType)(nil)).Elem()
}

func (e SslConfigServerCertType) ToSslConfigServerCertTypeOutput() SslConfigServerCertTypeOutput {
	return pulumi.ToOutput(e).(SslConfigServerCertTypeOutput)
}

func (e SslConfigServerCertType) ToSslConfigServerCertTypeOutputWithContext(ctx context.Context) SslConfigServerCertTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SslConfigServerCertTypeOutput)
}

func (e SslConfigServerCertType) ToSslConfigServerCertTypePtrOutput() SslConfigServerCertTypePtrOutput {
	return e.ToSslConfigServerCertTypePtrOutputWithContext(context.Background())
}

func (e SslConfigServerCertType) ToSslConfigServerCertTypePtrOutputWithContext(ctx context.Context) SslConfigServerCertTypePtrOutput {
	return SslConfigServerCertType(e).ToSslConfigServerCertTypeOutputWithContext(ctx).ToSslConfigServerCertTypePtrOutputWithContext(ctx)
}

func (e SslConfigServerCertType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslConfigServerCertType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslConfigServerCertType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SslConfigServerCertType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SslConfigServerCertTypeOutput struct{ *pulumi.OutputState }

func (SslConfigServerCertTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SslConfigServerCertType)(nil)).Elem()
}

func (o SslConfigServerCertTypeOutput) ToSslConfigServerCertTypeOutput() SslConfigServerCertTypeOutput {
	return o
}

func (o SslConfigServerCertTypeOutput) ToSslConfigServerCertTypeOutputWithContext(ctx context.Context) SslConfigServerCertTypeOutput {
	return o
}

func (o SslConfigServerCertTypeOutput) ToSslConfigServerCertTypePtrOutput() SslConfigServerCertTypePtrOutput {
	return o.ToSslConfigServerCertTypePtrOutputWithContext(context.Background())
}

func (o SslConfigServerCertTypeOutput) ToSslConfigServerCertTypePtrOutputWithContext(ctx context.Context) SslConfigServerCertTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SslConfigServerCertType) *SslConfigServerCertType {
		return &v
	}).(SslConfigServerCertTypePtrOutput)
}

func (o SslConfigServerCertTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SslConfigServerCertTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslConfigServerCertType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SslConfigServerCertTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslConfigServerCertTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslConfigServerCertType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SslConfigServerCertTypePtrOutput struct{ *pulumi.OutputState }

func (SslConfigServerCertTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslConfigServerCertType)(nil)).Elem()
}

func (o SslConfigServerCertTypePtrOutput) ToSslConfigServerCertTypePtrOutput() SslConfigServerCertTypePtrOutput {
	return o
}

func (o SslConfigServerCertTypePtrOutput) ToSslConfigServerCertTypePtrOutputWithContext(ctx context.Context) SslConfigServerCertTypePtrOutput {
	return o
}

func (o SslConfigServerCertTypePtrOutput) Elem() SslConfigServerCertTypeOutput {
	return o.ApplyT(func(v *SslConfigServerCertType) SslConfigServerCertType {
		if v != nil {
			return *v
		}
		var ret SslConfigServerCertType
		return ret
	}).(SslConfigServerCertTypeOutput)
}

func (o SslConfigServerCertTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslConfigServerCertTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SslConfigServerCertType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SslConfigServerCertTypeInput is an input type that accepts SslConfigServerCertTypeArgs and SslConfigServerCertTypeOutput values.
// You can construct a concrete instance of `SslConfigServerCertTypeInput` via:
//
//	SslConfigServerCertTypeArgs{...}
type SslConfigServerCertTypeInput interface {
	pulumi.Input

	ToSslConfigServerCertTypeOutput() SslConfigServerCertTypeOutput
	ToSslConfigServerCertTypeOutputWithContext(context.Context) SslConfigServerCertTypeOutput
}

var sslConfigServerCertTypePtrType = reflect.TypeOf((**SslConfigServerCertType)(nil)).Elem()

type SslConfigServerCertTypePtrInput interface {
	pulumi.Input

	ToSslConfigServerCertTypePtrOutput() SslConfigServerCertTypePtrOutput
	ToSslConfigServerCertTypePtrOutputWithContext(context.Context) SslConfigServerCertTypePtrOutput
}

type sslConfigServerCertTypePtr string

func SslConfigServerCertTypePtr(v string) SslConfigServerCertTypePtrInput {
	return (*sslConfigServerCertTypePtr)(&v)
}

func (*sslConfigServerCertTypePtr) ElementType() reflect.Type {
	return sslConfigServerCertTypePtrType
}

func (in *sslConfigServerCertTypePtr) ToSslConfigServerCertTypePtrOutput() SslConfigServerCertTypePtrOutput {
	return pulumi.ToOutput(in).(SslConfigServerCertTypePtrOutput)
}

func (in *sslConfigServerCertTypePtr) ToSslConfigServerCertTypePtrOutputWithContext(ctx context.Context) SslConfigServerCertTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SslConfigServerCertTypePtrOutput)
}

// Trust Model of the SSL connection
type SslConfigTrustModel string

const (
	// Public Trust Model. Takes the Default Java trust store.
	SslConfigTrustModelPublic = SslConfigTrustModel("PUBLIC")
	// Private Trust Model. Takes custom/private trust store.
	SslConfigTrustModelPrivate = SslConfigTrustModel("PRIVATE")
	// Insecure Trust Model. Accept all certificates.
	SslConfigTrustModelInsecure = SslConfigTrustModel("INSECURE")
)

func (SslConfigTrustModel) ElementType() reflect.Type {
	return reflect.TypeOf((*SslConfigTrustModel)(nil)).Elem()
}

func (e SslConfigTrustModel) ToSslConfigTrustModelOutput() SslConfigTrustModelOutput {
	return pulumi.ToOutput(e).(SslConfigTrustModelOutput)
}

func (e SslConfigTrustModel) ToSslConfigTrustModelOutputWithContext(ctx context.Context) SslConfigTrustModelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SslConfigTrustModelOutput)
}

func (e SslConfigTrustModel) ToSslConfigTrustModelPtrOutput() SslConfigTrustModelPtrOutput {
	return e.ToSslConfigTrustModelPtrOutputWithContext(context.Background())
}

func (e SslConfigTrustModel) ToSslConfigTrustModelPtrOutputWithContext(ctx context.Context) SslConfigTrustModelPtrOutput {
	return SslConfigTrustModel(e).ToSslConfigTrustModelOutputWithContext(ctx).ToSslConfigTrustModelPtrOutputWithContext(ctx)
}

func (e SslConfigTrustModel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslConfigTrustModel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslConfigTrustModel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SslConfigTrustModel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SslConfigTrustModelOutput struct{ *pulumi.OutputState }

func (SslConfigTrustModelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SslConfigTrustModel)(nil)).Elem()
}

func (o SslConfigTrustModelOutput) ToSslConfigTrustModelOutput() SslConfigTrustModelOutput {
	return o
}

func (o SslConfigTrustModelOutput) ToSslConfigTrustModelOutputWithContext(ctx context.Context) SslConfigTrustModelOutput {
	return o
}

func (o SslConfigTrustModelOutput) ToSslConfigTrustModelPtrOutput() SslConfigTrustModelPtrOutput {
	return o.ToSslConfigTrustModelPtrOutputWithContext(context.Background())
}

func (o SslConfigTrustModelOutput) ToSslConfigTrustModelPtrOutputWithContext(ctx context.Context) SslConfigTrustModelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SslConfigTrustModel) *SslConfigTrustModel {
		return &v
	}).(SslConfigTrustModelPtrOutput)
}

func (o SslConfigTrustModelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SslConfigTrustModelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslConfigTrustModel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SslConfigTrustModelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslConfigTrustModelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslConfigTrustModel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SslConfigTrustModelPtrOutput struct{ *pulumi.OutputState }

func (SslConfigTrustModelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslConfigTrustModel)(nil)).Elem()
}

func (o SslConfigTrustModelPtrOutput) ToSslConfigTrustModelPtrOutput() SslConfigTrustModelPtrOutput {
	return o
}

func (o SslConfigTrustModelPtrOutput) ToSslConfigTrustModelPtrOutputWithContext(ctx context.Context) SslConfigTrustModelPtrOutput {
	return o
}

func (o SslConfigTrustModelPtrOutput) Elem() SslConfigTrustModelOutput {
	return o.ApplyT(func(v *SslConfigTrustModel) SslConfigTrustModel {
		if v != nil {
			return *v
		}
		var ret SslConfigTrustModel
		return ret
	}).(SslConfigTrustModelOutput)
}

func (o SslConfigTrustModelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslConfigTrustModelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SslConfigTrustModel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SslConfigTrustModelInput is an input type that accepts SslConfigTrustModelArgs and SslConfigTrustModelOutput values.
// You can construct a concrete instance of `SslConfigTrustModelInput` via:
//
//	SslConfigTrustModelArgs{...}
type SslConfigTrustModelInput interface {
	pulumi.Input

	ToSslConfigTrustModelOutput() SslConfigTrustModelOutput
	ToSslConfigTrustModelOutputWithContext(context.Context) SslConfigTrustModelOutput
}

var sslConfigTrustModelPtrType = reflect.TypeOf((**SslConfigTrustModel)(nil)).Elem()

type SslConfigTrustModelPtrInput interface {
	pulumi.Input

	ToSslConfigTrustModelPtrOutput() SslConfigTrustModelPtrOutput
	ToSslConfigTrustModelPtrOutputWithContext(context.Context) SslConfigTrustModelPtrOutput
}

type sslConfigTrustModelPtr string

func SslConfigTrustModelPtr(v string) SslConfigTrustModelPtrInput {
	return (*sslConfigTrustModelPtr)(&v)
}

func (*sslConfigTrustModelPtr) ElementType() reflect.Type {
	return sslConfigTrustModelPtrType
}

func (in *sslConfigTrustModelPtr) ToSslConfigTrustModelPtrOutput() SslConfigTrustModelPtrOutput {
	return pulumi.ToOutput(in).(SslConfigTrustModelPtrOutput)
}

func (in *sslConfigTrustModelPtr) ToSslConfigTrustModelPtrOutputWithContext(ctx context.Context) SslConfigTrustModelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SslConfigTrustModelPtrOutput)
}

// Controls the ssl type for the given connector version.
type SslConfigType string

const (
	// No SSL configuration required.
	SslConfigTypeSslTypeUnspecified = SslConfigType("SSL_TYPE_UNSPECIFIED")
	// TLS Handshake
	SslConfigTypeTls = SslConfigType("TLS")
	// mutual TLS (MTLS) Handshake
	SslConfigTypeMtls = SslConfigType("MTLS")
)

func (SslConfigType) ElementType() reflect.Type {
	return reflect.TypeOf((*SslConfigType)(nil)).Elem()
}

func (e SslConfigType) ToSslConfigTypeOutput() SslConfigTypeOutput {
	return pulumi.ToOutput(e).(SslConfigTypeOutput)
}

func (e SslConfigType) ToSslConfigTypeOutputWithContext(ctx context.Context) SslConfigTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SslConfigTypeOutput)
}

func (e SslConfigType) ToSslConfigTypePtrOutput() SslConfigTypePtrOutput {
	return e.ToSslConfigTypePtrOutputWithContext(context.Background())
}

func (e SslConfigType) ToSslConfigTypePtrOutputWithContext(ctx context.Context) SslConfigTypePtrOutput {
	return SslConfigType(e).ToSslConfigTypeOutputWithContext(ctx).ToSslConfigTypePtrOutputWithContext(ctx)
}

func (e SslConfigType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslConfigType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslConfigType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SslConfigType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SslConfigTypeOutput struct{ *pulumi.OutputState }

func (SslConfigTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SslConfigType)(nil)).Elem()
}

func (o SslConfigTypeOutput) ToSslConfigTypeOutput() SslConfigTypeOutput {
	return o
}

func (o SslConfigTypeOutput) ToSslConfigTypeOutputWithContext(ctx context.Context) SslConfigTypeOutput {
	return o
}

func (o SslConfigTypeOutput) ToSslConfigTypePtrOutput() SslConfigTypePtrOutput {
	return o.ToSslConfigTypePtrOutputWithContext(context.Background())
}

func (o SslConfigTypeOutput) ToSslConfigTypePtrOutputWithContext(ctx context.Context) SslConfigTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SslConfigType) *SslConfigType {
		return &v
	}).(SslConfigTypePtrOutput)
}

func (o SslConfigTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SslConfigTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslConfigType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SslConfigTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslConfigTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslConfigType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SslConfigTypePtrOutput struct{ *pulumi.OutputState }

func (SslConfigTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslConfigType)(nil)).Elem()
}

func (o SslConfigTypePtrOutput) ToSslConfigTypePtrOutput() SslConfigTypePtrOutput {
	return o
}

func (o SslConfigTypePtrOutput) ToSslConfigTypePtrOutputWithContext(ctx context.Context) SslConfigTypePtrOutput {
	return o
}

func (o SslConfigTypePtrOutput) Elem() SslConfigTypeOutput {
	return o.ApplyT(func(v *SslConfigType) SslConfigType {
		if v != nil {
			return *v
		}
		var ret SslConfigType
		return ret
	}).(SslConfigTypeOutput)
}

func (o SslConfigTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslConfigTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SslConfigType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SslConfigTypeInput is an input type that accepts SslConfigTypeArgs and SslConfigTypeOutput values.
// You can construct a concrete instance of `SslConfigTypeInput` via:
//
//	SslConfigTypeArgs{...}
type SslConfigTypeInput interface {
	pulumi.Input

	ToSslConfigTypeOutput() SslConfigTypeOutput
	ToSslConfigTypeOutputWithContext(context.Context) SslConfigTypeOutput
}

var sslConfigTypePtrType = reflect.TypeOf((**SslConfigType)(nil)).Elem()

type SslConfigTypePtrInput interface {
	pulumi.Input

	ToSslConfigTypePtrOutput() SslConfigTypePtrOutput
	ToSslConfigTypePtrOutputWithContext(context.Context) SslConfigTypePtrOutput
}

type sslConfigTypePtr string

func SslConfigTypePtr(v string) SslConfigTypePtrInput {
	return (*sslConfigTypePtr)(&v)
}

func (*sslConfigTypePtr) ElementType() reflect.Type {
	return sslConfigTypePtrType
}

func (in *sslConfigTypePtr) ToSslConfigTypePtrOutput() SslConfigTypePtrOutput {
	return pulumi.ToOutput(in).(SslConfigTypePtrOutput)
}

func (in *sslConfigTypePtr) ToSslConfigTypePtrOutputWithContext(ctx context.Context) SslConfigTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SslConfigTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogConfigLogTypeInput)(nil)).Elem(), AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogConfigLogTypePtrInput)(nil)).Elem(), AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*AuthConfigAuthTypeInput)(nil)).Elem(), AuthConfigAuthType("AUTH_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*AuthConfigAuthTypePtrInput)(nil)).Elem(), AuthConfigAuthType("AUTH_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*SslConfigClientCertTypeInput)(nil)).Elem(), SslConfigClientCertType("CERT_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*SslConfigClientCertTypePtrInput)(nil)).Elem(), SslConfigClientCertType("CERT_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*SslConfigServerCertTypeInput)(nil)).Elem(), SslConfigServerCertType("CERT_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*SslConfigServerCertTypePtrInput)(nil)).Elem(), SslConfigServerCertType("CERT_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*SslConfigTrustModelInput)(nil)).Elem(), SslConfigTrustModel("PUBLIC"))
	pulumi.RegisterInputType(reflect.TypeOf((*SslConfigTrustModelPtrInput)(nil)).Elem(), SslConfigTrustModel("PUBLIC"))
	pulumi.RegisterInputType(reflect.TypeOf((*SslConfigTypeInput)(nil)).Elem(), SslConfigType("SSL_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*SslConfigTypePtrInput)(nil)).Elem(), SslConfigType("SSL_TYPE_UNSPECIFIED"))
	pulumi.RegisterOutputType(AuditLogConfigLogTypeOutput{})
	pulumi.RegisterOutputType(AuditLogConfigLogTypePtrOutput{})
	pulumi.RegisterOutputType(AuthConfigAuthTypeOutput{})
	pulumi.RegisterOutputType(AuthConfigAuthTypePtrOutput{})
	pulumi.RegisterOutputType(SslConfigClientCertTypeOutput{})
	pulumi.RegisterOutputType(SslConfigClientCertTypePtrOutput{})
	pulumi.RegisterOutputType(SslConfigServerCertTypeOutput{})
	pulumi.RegisterOutputType(SslConfigServerCertTypePtrOutput{})
	pulumi.RegisterOutputType(SslConfigTrustModelOutput{})
	pulumi.RegisterOutputType(SslConfigTrustModelPtrOutput{})
	pulumi.RegisterOutputType(SslConfigTypeOutput{})
	pulumi.RegisterOutputType(SslConfigTypePtrOutput{})
}
