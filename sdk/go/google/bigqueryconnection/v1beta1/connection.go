// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new connection.
type Connection struct {
	pulumi.CustomResourceState

	// Cloud SQL properties.
	CloudSql CloudSqlPropertiesResponseOutput `pulumi:"cloudSql"`
	// Optional. Connection id that should be assigned to the created connection.
	ConnectionId pulumi.StringPtrOutput `pulumi:"connectionId"`
	// The creation timestamp of the connection.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// User provided description.
	Description pulumi.StringOutput `pulumi:"description"`
	// User provided display name for the connection.
	FriendlyName pulumi.StringOutput `pulumi:"friendlyName"`
	// True, if credential is configured for this connection.
	HasCredential pulumi.BoolOutput `pulumi:"hasCredential"`
	// The last update timestamp of the connection.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	Location         pulumi.StringOutput `pulumi:"location"`
	// The resource name of the connection in the form of: `projects/{project_id}/locations/{location_id}/connections/{connection_id}`
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil {
		args = &ConnectionArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource Connection
	err := ctx.RegisterResource("google-native:bigqueryconnection/v1beta1:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("google-native:bigqueryconnection/v1beta1:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
}

type ConnectionState struct {
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// Cloud SQL properties.
	CloudSql *CloudSqlProperties `pulumi:"cloudSql"`
	// Optional. Connection id that should be assigned to the created connection.
	ConnectionId *string `pulumi:"connectionId"`
	// User provided description.
	Description *string `pulumi:"description"`
	// User provided display name for the connection.
	FriendlyName *string `pulumi:"friendlyName"`
	Location     *string `pulumi:"location"`
	// The resource name of the connection in the form of: `projects/{project_id}/locations/{location_id}/connections/{connection_id}`
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// Cloud SQL properties.
	CloudSql CloudSqlPropertiesPtrInput
	// Optional. Connection id that should be assigned to the created connection.
	ConnectionId pulumi.StringPtrInput
	// User provided description.
	Description pulumi.StringPtrInput
	// User provided display name for the connection.
	FriendlyName pulumi.StringPtrInput
	Location     pulumi.StringPtrInput
	// The resource name of the connection in the form of: `projects/{project_id}/locations/{location_id}/connections/{connection_id}`
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}

type ConnectionInput interface {
	pulumi.Input

	ToConnectionOutput() ConnectionOutput
	ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput
}

func (*Connection) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (i *Connection) ToConnectionOutput() ConnectionOutput {
	return i.ToConnectionOutputWithContext(context.Background())
}

func (i *Connection) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionOutput)
}

type ConnectionOutput struct{ *pulumi.OutputState }

func (ConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Connection)(nil)).Elem()
}

func (o ConnectionOutput) ToConnectionOutput() ConnectionOutput {
	return o
}

func (o ConnectionOutput) ToConnectionOutputWithContext(ctx context.Context) ConnectionOutput {
	return o
}

// Cloud SQL properties.
func (o ConnectionOutput) CloudSql() CloudSqlPropertiesResponseOutput {
	return o.ApplyT(func(v *Connection) CloudSqlPropertiesResponseOutput { return v.CloudSql }).(CloudSqlPropertiesResponseOutput)
}

// Optional. Connection id that should be assigned to the created connection.
func (o ConnectionOutput) ConnectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringPtrOutput { return v.ConnectionId }).(pulumi.StringPtrOutput)
}

// The creation timestamp of the connection.
func (o ConnectionOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// User provided description.
func (o ConnectionOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// User provided display name for the connection.
func (o ConnectionOutput) FriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.FriendlyName }).(pulumi.StringOutput)
}

// True, if credential is configured for this connection.
func (o ConnectionOutput) HasCredential() pulumi.BoolOutput {
	return o.ApplyT(func(v *Connection) pulumi.BoolOutput { return v.HasCredential }).(pulumi.BoolOutput)
}

// The last update timestamp of the connection.
func (o ConnectionOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o ConnectionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the connection in the form of: `projects/{project_id}/locations/{location_id}/connections/{connection_id}`
func (o ConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConnectionOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Connection) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionInput)(nil)).Elem(), &Connection{})
	pulumi.RegisterOutputType(ConnectionOutput{})
}
