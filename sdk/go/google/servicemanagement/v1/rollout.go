// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new service configuration rollout. Based on rollout, the Google Service Management will roll out the service configurations to different backend services. For example, the logging configuration will be pushed to Google Cloud Logging. Please note that any previous pending and running Rollouts and associated Operations will be automatically cancelled so that the latest Rollout will not be blocked by previous Rollouts. Only the 100 most recent (in any state) and the last 10 successful (if not already part of the set of 100 most recent) rollouts are kept for each service. The rest will be deleted eventually. Operation
// Auto-naming is currently not supported for this resource.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type Rollout struct {
	pulumi.CustomResourceState

	// Creation time of the rollout. Readonly.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The user who created the Rollout. Readonly.
	CreatedBy pulumi.StringOutput `pulumi:"createdBy"`
	// The strategy associated with a rollout to delete a `ManagedService`. Readonly.
	DeleteServiceStrategy DeleteServiceStrategyResponseOutput `pulumi:"deleteServiceStrategy"`
	// Optional. Unique identifier of this Rollout. Must be no longer than 63 characters and only lower case letters, digits, '.', '_' and '-' are allowed. If not specified by client, the server will generate one. The generated id will have the form of , where "date" is the create date in ISO 8601 format. "revision number" is a monotonically increasing positive number that is reset every day for each service. An example of the generated rollout_id is '2016-02-16r1'
	RolloutId   pulumi.StringOutput `pulumi:"rolloutId"`
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// The status of this rollout. Readonly. In case of a failed rollout, the system will automatically rollback to the current Rollout version. Readonly.
	Status pulumi.StringOutput `pulumi:"status"`
	// Google Service Control selects service configurations based on traffic percentage.
	TrafficPercentStrategy TrafficPercentStrategyResponseOutput `pulumi:"trafficPercentStrategy"`
}

// NewRollout registers a new resource with the given unique name, arguments, and options.
func NewRollout(ctx *pulumi.Context,
	name string, args *RolloutArgs, opts ...pulumi.ResourceOption) (*Rollout, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"serviceName",
	})
	opts = append(opts, replaceOnChanges)
	var resource Rollout
	err := ctx.RegisterResource("google-native:servicemanagement/v1:Rollout", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRollout gets an existing Rollout resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRollout(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RolloutState, opts ...pulumi.ResourceOption) (*Rollout, error) {
	var resource Rollout
	err := ctx.ReadResource("google-native:servicemanagement/v1:Rollout", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rollout resources.
type rolloutState struct {
}

type RolloutState struct {
}

func (RolloutState) ElementType() reflect.Type {
	return reflect.TypeOf((*rolloutState)(nil)).Elem()
}

type rolloutArgs struct {
	// Creation time of the rollout. Readonly.
	CreateTime *string `pulumi:"createTime"`
	// The user who created the Rollout. Readonly.
	CreatedBy *string `pulumi:"createdBy"`
	// The strategy associated with a rollout to delete a `ManagedService`. Readonly.
	DeleteServiceStrategy *DeleteServiceStrategy `pulumi:"deleteServiceStrategy"`
	// Optional. Unique identifier of this Rollout. Must be no longer than 63 characters and only lower case letters, digits, '.', '_' and '-' are allowed. If not specified by client, the server will generate one. The generated id will have the form of , where "date" is the create date in ISO 8601 format. "revision number" is a monotonically increasing positive number that is reset every day for each service. An example of the generated rollout_id is '2016-02-16r1'
	RolloutId *string `pulumi:"rolloutId"`
	// The name of the service associated with this Rollout.
	ServiceName string `pulumi:"serviceName"`
	// Google Service Control selects service configurations based on traffic percentage.
	TrafficPercentStrategy *TrafficPercentStrategy `pulumi:"trafficPercentStrategy"`
}

// The set of arguments for constructing a Rollout resource.
type RolloutArgs struct {
	// Creation time of the rollout. Readonly.
	CreateTime pulumi.StringPtrInput
	// The user who created the Rollout. Readonly.
	CreatedBy pulumi.StringPtrInput
	// The strategy associated with a rollout to delete a `ManagedService`. Readonly.
	DeleteServiceStrategy DeleteServiceStrategyPtrInput
	// Optional. Unique identifier of this Rollout. Must be no longer than 63 characters and only lower case letters, digits, '.', '_' and '-' are allowed. If not specified by client, the server will generate one. The generated id will have the form of , where "date" is the create date in ISO 8601 format. "revision number" is a monotonically increasing positive number that is reset every day for each service. An example of the generated rollout_id is '2016-02-16r1'
	RolloutId pulumi.StringPtrInput
	// The name of the service associated with this Rollout.
	ServiceName pulumi.StringInput
	// Google Service Control selects service configurations based on traffic percentage.
	TrafficPercentStrategy TrafficPercentStrategyPtrInput
}

func (RolloutArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rolloutArgs)(nil)).Elem()
}

type RolloutInput interface {
	pulumi.Input

	ToRolloutOutput() RolloutOutput
	ToRolloutOutputWithContext(ctx context.Context) RolloutOutput
}

func (*Rollout) ElementType() reflect.Type {
	return reflect.TypeOf((**Rollout)(nil)).Elem()
}

func (i *Rollout) ToRolloutOutput() RolloutOutput {
	return i.ToRolloutOutputWithContext(context.Background())
}

func (i *Rollout) ToRolloutOutputWithContext(ctx context.Context) RolloutOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RolloutOutput)
}

type RolloutOutput struct{ *pulumi.OutputState }

func (RolloutOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rollout)(nil)).Elem()
}

func (o RolloutOutput) ToRolloutOutput() RolloutOutput {
	return o
}

func (o RolloutOutput) ToRolloutOutputWithContext(ctx context.Context) RolloutOutput {
	return o
}

// Creation time of the rollout. Readonly.
func (o RolloutOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Rollout) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The user who created the Rollout. Readonly.
func (o RolloutOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *Rollout) pulumi.StringOutput { return v.CreatedBy }).(pulumi.StringOutput)
}

// The strategy associated with a rollout to delete a `ManagedService`. Readonly.
func (o RolloutOutput) DeleteServiceStrategy() DeleteServiceStrategyResponseOutput {
	return o.ApplyT(func(v *Rollout) DeleteServiceStrategyResponseOutput { return v.DeleteServiceStrategy }).(DeleteServiceStrategyResponseOutput)
}

// Optional. Unique identifier of this Rollout. Must be no longer than 63 characters and only lower case letters, digits, '.', '_' and '-' are allowed. If not specified by client, the server will generate one. The generated id will have the form of , where "date" is the create date in ISO 8601 format. "revision number" is a monotonically increasing positive number that is reset every day for each service. An example of the generated rollout_id is '2016-02-16r1'
func (o RolloutOutput) RolloutId() pulumi.StringOutput {
	return o.ApplyT(func(v *Rollout) pulumi.StringOutput { return v.RolloutId }).(pulumi.StringOutput)
}

func (o RolloutOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Rollout) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// The status of this rollout. Readonly. In case of a failed rollout, the system will automatically rollback to the current Rollout version. Readonly.
func (o RolloutOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Rollout) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Google Service Control selects service configurations based on traffic percentage.
func (o RolloutOutput) TrafficPercentStrategy() TrafficPercentStrategyResponseOutput {
	return o.ApplyT(func(v *Rollout) TrafficPercentStrategyResponseOutput { return v.TrafficPercentStrategy }).(TrafficPercentStrategyResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RolloutInput)(nil)).Elem(), &Rollout{})
	pulumi.RegisterOutputType(RolloutOutput{})
}
