// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Microsoft AD domain.
// Auto-naming is currently not supported for this resource.
type Domain struct {
	pulumi.CustomResourceState

	// Optional. The name of delegated administrator account used to perform Active Directory operations. If not specified, `setupadmin` will be used.
	Admin pulumi.StringOutput `pulumi:"admin"`
	// Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
	AuditLogsEnabled pulumi.BoolOutput `pulumi:"auditLogsEnabled"`
	// Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) the domain instance is connected to. Networks can be added using UpdateDomain. The domain is only available on networks listed in `authorized_networks`. If CIDR subnets overlap between networks, domain creation will fail.
	AuthorizedNetworks pulumi.StringArrayOutput `pulumi:"authorizedNetworks"`
	// The time the instance was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Required. A domain name, e.g. mydomain.myorg.com, with the following restrictions: * Must contain only lowercase letters, numbers, periods and hyphens. * Must start with a letter. * Must contain between 2-64 characters. * Must end with a number or a letter. * Must not start with period. * First segment length (mydomain form example above) shouldn't exceed 15 chars. * The last segment cannot be fully numeric. * Must be unique within the customer project.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The fully-qualified domain name of the exposed domain used by clients to connect to the service. Similar to what would be chosen for an Active Directory set up on an internal network.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// Optional. Resource labels that can contain user-provided metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	Locations pulumi.StringArrayOutput `pulumi:"locations"`
	// The unique name of the domain using the form: `projects/{project_id}/locations/global/domains/{domain_name}`.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
	ReservedIpRange pulumi.StringOutput `pulumi:"reservedIpRange"`
	// The current state of this domain.
	State pulumi.StringOutput `pulumi:"state"`
	// Additional information about the current status of this domain, if available.
	StatusMessage pulumi.StringOutput `pulumi:"statusMessage"`
	// The current trusts associated with the domain.
	Trusts TrustResponseArrayOutput `pulumi:"trusts"`
	// The last update time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.Locations == nil {
		return nil, errors.New("invalid value for required argument 'Locations'")
	}
	if args.ReservedIpRange == nil {
		return nil, errors.New("invalid value for required argument 'ReservedIpRange'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"domainName",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource Domain
	err := ctx.RegisterResource("google-native:managedidentities/v1beta1:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("google-native:managedidentities/v1beta1:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
}

type DomainState struct {
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// Optional. The name of delegated administrator account used to perform Active Directory operations. If not specified, `setupadmin` will be used.
	Admin *string `pulumi:"admin"`
	// Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
	AuditLogsEnabled *bool `pulumi:"auditLogsEnabled"`
	// Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) the domain instance is connected to. Networks can be added using UpdateDomain. The domain is only available on networks listed in `authorized_networks`. If CIDR subnets overlap between networks, domain creation will fail.
	AuthorizedNetworks []string `pulumi:"authorizedNetworks"`
	// Required. A domain name, e.g. mydomain.myorg.com, with the following restrictions: * Must contain only lowercase letters, numbers, periods and hyphens. * Must start with a letter. * Must contain between 2-64 characters. * Must end with a number or a letter. * Must not start with period. * First segment length (mydomain form example above) shouldn't exceed 15 chars. * The last segment cannot be fully numeric. * Must be unique within the customer project.
	DomainName string `pulumi:"domainName"`
	// Optional. Resource labels that can contain user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	Locations []string `pulumi:"locations"`
	Project   *string  `pulumi:"project"`
	// The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
	ReservedIpRange string `pulumi:"reservedIpRange"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// Optional. The name of delegated administrator account used to perform Active Directory operations. If not specified, `setupadmin` will be used.
	Admin pulumi.StringPtrInput
	// Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
	AuditLogsEnabled pulumi.BoolPtrInput
	// Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) the domain instance is connected to. Networks can be added using UpdateDomain. The domain is only available on networks listed in `authorized_networks`. If CIDR subnets overlap between networks, domain creation will fail.
	AuthorizedNetworks pulumi.StringArrayInput
	// Required. A domain name, e.g. mydomain.myorg.com, with the following restrictions: * Must contain only lowercase letters, numbers, periods and hyphens. * Must start with a letter. * Must contain between 2-64 characters. * Must end with a number or a letter. * Must not start with period. * First segment length (mydomain form example above) shouldn't exceed 15 chars. * The last segment cannot be fully numeric. * Must be unique within the customer project.
	DomainName pulumi.StringInput
	// Optional. Resource labels that can contain user-provided metadata.
	Labels pulumi.StringMapInput
	// Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	Locations pulumi.StringArrayInput
	Project   pulumi.StringPtrInput
	// The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
	ReservedIpRange pulumi.StringInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (*Domain) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

type DomainOutput struct{ *pulumi.OutputState }

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

// Optional. The name of delegated administrator account used to perform Active Directory operations. If not specified, `setupadmin` will be used.
func (o DomainOutput) Admin() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Admin }).(pulumi.StringOutput)
}

// Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
func (o DomainOutput) AuditLogsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Domain) pulumi.BoolOutput { return v.AuditLogsEnabled }).(pulumi.BoolOutput)
}

// Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) the domain instance is connected to. Networks can be added using UpdateDomain. The domain is only available on networks listed in `authorized_networks`. If CIDR subnets overlap between networks, domain creation will fail.
func (o DomainOutput) AuthorizedNetworks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringArrayOutput { return v.AuthorizedNetworks }).(pulumi.StringArrayOutput)
}

// The time the instance was created.
func (o DomainOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Required. A domain name, e.g. mydomain.myorg.com, with the following restrictions: * Must contain only lowercase letters, numbers, periods and hyphens. * Must start with a letter. * Must contain between 2-64 characters. * Must end with a number or a letter. * Must not start with period. * First segment length (mydomain form example above) shouldn't exceed 15 chars. * The last segment cannot be fully numeric. * Must be unique within the customer project.
func (o DomainOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.DomainName }).(pulumi.StringOutput)
}

// The fully-qualified domain name of the exposed domain used by clients to connect to the service. Similar to what would be chosen for an Active Directory set up on an internal network.
func (o DomainOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

// Optional. Resource labels that can contain user-provided metadata.
func (o DomainOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
func (o DomainOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringArrayOutput { return v.Locations }).(pulumi.StringArrayOutput)
}

// The unique name of the domain using the form: `projects/{project_id}/locations/global/domains/{domain_name}`.
func (o DomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DomainOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
func (o DomainOutput) ReservedIpRange() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.ReservedIpRange }).(pulumi.StringOutput)
}

// The current state of this domain.
func (o DomainOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Additional information about the current status of this domain, if available.
func (o DomainOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.StatusMessage }).(pulumi.StringOutput)
}

// The current trusts associated with the domain.
func (o DomainOutput) Trusts() TrustResponseArrayOutput {
	return o.ApplyT(func(v *Domain) TrustResponseArrayOutput { return v.Trusts }).(TrustResponseArrayOutput)
}

// The last update time.
func (o DomainOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainInput)(nil)).Elem(), &Domain{})
	pulumi.RegisterOutputType(DomainOutput{})
}
