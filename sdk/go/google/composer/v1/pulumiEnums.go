// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The current state of the environment.
type EnvironmentStateEnum pulumi.String

const (
	// The state of the environment is unknown.
	EnvironmentStateEnumStateUnspecified = EnvironmentStateEnum("STATE_UNSPECIFIED")
	// The environment is in the process of being created.
	EnvironmentStateEnumCreating = EnvironmentStateEnum("CREATING")
	// The environment is currently running and healthy. It is ready for use.
	EnvironmentStateEnumRunning = EnvironmentStateEnum("RUNNING")
	// The environment is being updated. It remains usable but cannot receive additional update requests or be deleted at this time.
	EnvironmentStateEnumUpdating = EnvironmentStateEnum("UPDATING")
	// The environment is undergoing deletion. It cannot be used.
	EnvironmentStateEnumDeleting = EnvironmentStateEnum("DELETING")
	// The environment has encountered an error and cannot be used.
	EnvironmentStateEnumError = EnvironmentStateEnum("ERROR")
)

func (EnvironmentStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e EnvironmentStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnvironmentStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnvironmentStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EnvironmentStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}