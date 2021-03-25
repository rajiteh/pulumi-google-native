// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a `WorkerPool` to run the builds, and returns the new worker pool. NOTE: As of now, this method returns an `Operation` that is always complete.
type WorkerPool struct {
	pulumi.CustomResourceState
}

// NewWorkerPool registers a new resource with the given unique name, arguments, and options.
func NewWorkerPool(ctx *pulumi.Context,
	name string, args *WorkerPoolArgs, opts ...pulumi.ResourceOption) (*WorkerPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	var resource WorkerPool
	err := ctx.RegisterResource("google-cloud:cloudbuild/v1beta1:WorkerPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkerPool gets an existing WorkerPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkerPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkerPoolState, opts ...pulumi.ResourceOption) (*WorkerPool, error) {
	var resource WorkerPool
	err := ctx.ReadResource("google-cloud:cloudbuild/v1beta1:WorkerPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkerPool resources.
type workerPoolState struct {
}

type WorkerPoolState struct {
}

func (WorkerPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*workerPoolState)(nil)).Elem()
}

type workerPoolArgs struct {
	// Output only. Time at which the request to create the `WorkerPool` was received.
	CreateTime *string `pulumi:"createTime"`
	// Output only. Time at which the request to delete the `WorkerPool` was received.
	DeleteTime *string `pulumi:"deleteTime"`
	// Output only. The resource name of the `WorkerPool`, with format `projects/{project}/locations/{location}/workerPools/{worker_pool}`. The value of `{worker_pool}` is provided by `worker_pool_id` in `CreateWorkerPool` request and the value of `{location}` is determined by the endpoint accessed.
	Name *string `pulumi:"name"`
	// Network configuration for the `WorkerPool`.
	NetworkConfig *NetworkConfig `pulumi:"networkConfig"`
	// Required. The parent resource where this worker pool will be created. Format: `projects/{project}/locations/{location}`.
	Parent string `pulumi:"parent"`
	// Output only. `WorkerPool` state.
	State *string `pulumi:"state"`
	// Output only. Time at which the request to update the `WorkerPool` was received.
	UpdateTime *string `pulumi:"updateTime"`
	// Worker configuration for the `WorkerPool`.
	WorkerConfig *WorkerConfig `pulumi:"workerConfig"`
	// Required. Immutable. The ID to use for the `WorkerPool`, which will become the final component of the resource name. This value should be 1-63 characters, and valid characters are /a-z-/.
	WorkerPoolId *string `pulumi:"workerPoolId"`
}

// The set of arguments for constructing a WorkerPool resource.
type WorkerPoolArgs struct {
	// Output only. Time at which the request to create the `WorkerPool` was received.
	CreateTime pulumi.StringPtrInput
	// Output only. Time at which the request to delete the `WorkerPool` was received.
	DeleteTime pulumi.StringPtrInput
	// Output only. The resource name of the `WorkerPool`, with format `projects/{project}/locations/{location}/workerPools/{worker_pool}`. The value of `{worker_pool}` is provided by `worker_pool_id` in `CreateWorkerPool` request and the value of `{location}` is determined by the endpoint accessed.
	Name pulumi.StringPtrInput
	// Network configuration for the `WorkerPool`.
	NetworkConfig NetworkConfigPtrInput
	// Required. The parent resource where this worker pool will be created. Format: `projects/{project}/locations/{location}`.
	Parent pulumi.StringInput
	// Output only. `WorkerPool` state.
	State pulumi.StringPtrInput
	// Output only. Time at which the request to update the `WorkerPool` was received.
	UpdateTime pulumi.StringPtrInput
	// Worker configuration for the `WorkerPool`.
	WorkerConfig WorkerConfigPtrInput
	// Required. Immutable. The ID to use for the `WorkerPool`, which will become the final component of the resource name. This value should be 1-63 characters, and valid characters are /a-z-/.
	WorkerPoolId pulumi.StringPtrInput
}

func (WorkerPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workerPoolArgs)(nil)).Elem()
}

type WorkerPoolInput interface {
	pulumi.Input

	ToWorkerPoolOutput() WorkerPoolOutput
	ToWorkerPoolOutputWithContext(ctx context.Context) WorkerPoolOutput
}

func (*WorkerPool) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkerPool)(nil))
}

func (i *WorkerPool) ToWorkerPoolOutput() WorkerPoolOutput {
	return i.ToWorkerPoolOutputWithContext(context.Background())
}

func (i *WorkerPool) ToWorkerPoolOutputWithContext(ctx context.Context) WorkerPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkerPoolOutput)
}

type WorkerPoolOutput struct {
	*pulumi.OutputState
}

func (WorkerPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkerPool)(nil))
}

func (o WorkerPoolOutput) ToWorkerPoolOutput() WorkerPoolOutput {
	return o
}

func (o WorkerPoolOutput) ToWorkerPoolOutputWithContext(ctx context.Context) WorkerPoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkerPoolOutput{})
}