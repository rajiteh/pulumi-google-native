// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Repository in a given project and location.
// Auto-naming is currently not supported for this resource.
type Repository struct {
	pulumi.CustomResourceState

	// Optional. If set, configures this repository to be linked to a Git remote.
	GitRemoteSettings GitRemoteSettingsResponseOutput `pulumi:"gitRemoteSettings"`
	Location          pulumi.StringOutput             `pulumi:"location"`
	// The repository's name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. The name of the Secret Manager secret version to be used to interpolate variables into the .npmrc file for package installation operations. Must be in the format `projects/*/secrets/*/versions/*`. The file itself must be in a JSON format.
	NpmrcEnvironmentVariablesSecretVersion pulumi.StringOutput `pulumi:"npmrcEnvironmentVariablesSecretVersion"`
	Project                                pulumi.StringOutput `pulumi:"project"`
	// Required. The ID to use for the repository, which will become the final component of the repository's resource name.
	RepositoryId pulumi.StringOutput `pulumi:"repositoryId"`
	// Optional. If set, fields of `workspace_compilation_overrides` override the default compilation settings that are specified in dataform.json when creating workspace-scoped compilation results. See documentation for `WorkspaceCompilationOverrides` for more information.
	WorkspaceCompilationOverrides WorkspaceCompilationOverridesResponseOutput `pulumi:"workspaceCompilationOverrides"`
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOption) (*Repository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RepositoryId == nil {
		return nil, errors.New("invalid value for required argument 'RepositoryId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
		"repositoryId",
	})
	opts = append(opts, replaceOnChanges)
	var resource Repository
	err := ctx.RegisterResource("google-native:dataform/v1beta1:Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepository gets an existing Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryState, opts ...pulumi.ResourceOption) (*Repository, error) {
	var resource Repository
	err := ctx.ReadResource("google-native:dataform/v1beta1:Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Repository resources.
type repositoryState struct {
}

type RepositoryState struct {
}

func (RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryState)(nil)).Elem()
}

type repositoryArgs struct {
	// Optional. If set, configures this repository to be linked to a Git remote.
	GitRemoteSettings *GitRemoteSettings `pulumi:"gitRemoteSettings"`
	Location          *string            `pulumi:"location"`
	// Optional. The name of the Secret Manager secret version to be used to interpolate variables into the .npmrc file for package installation operations. Must be in the format `projects/*/secrets/*/versions/*`. The file itself must be in a JSON format.
	NpmrcEnvironmentVariablesSecretVersion *string `pulumi:"npmrcEnvironmentVariablesSecretVersion"`
	Project                                *string `pulumi:"project"`
	// Required. The ID to use for the repository, which will become the final component of the repository's resource name.
	RepositoryId string `pulumi:"repositoryId"`
	// Optional. If set, fields of `workspace_compilation_overrides` override the default compilation settings that are specified in dataform.json when creating workspace-scoped compilation results. See documentation for `WorkspaceCompilationOverrides` for more information.
	WorkspaceCompilationOverrides *WorkspaceCompilationOverrides `pulumi:"workspaceCompilationOverrides"`
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	// Optional. If set, configures this repository to be linked to a Git remote.
	GitRemoteSettings GitRemoteSettingsPtrInput
	Location          pulumi.StringPtrInput
	// Optional. The name of the Secret Manager secret version to be used to interpolate variables into the .npmrc file for package installation operations. Must be in the format `projects/*/secrets/*/versions/*`. The file itself must be in a JSON format.
	NpmrcEnvironmentVariablesSecretVersion pulumi.StringPtrInput
	Project                                pulumi.StringPtrInput
	// Required. The ID to use for the repository, which will become the final component of the repository's resource name.
	RepositoryId pulumi.StringInput
	// Optional. If set, fields of `workspace_compilation_overrides` override the default compilation settings that are specified in dataform.json when creating workspace-scoped compilation results. See documentation for `WorkspaceCompilationOverrides` for more information.
	WorkspaceCompilationOverrides WorkspaceCompilationOverridesPtrInput
}

func (RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryArgs)(nil)).Elem()
}

type RepositoryInput interface {
	pulumi.Input

	ToRepositoryOutput() RepositoryOutput
	ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput
}

func (*Repository) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil)).Elem()
}

func (i *Repository) ToRepositoryOutput() RepositoryOutput {
	return i.ToRepositoryOutputWithContext(context.Background())
}

func (i *Repository) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryOutput)
}

type RepositoryOutput struct{ *pulumi.OutputState }

func (RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil)).Elem()
}

func (o RepositoryOutput) ToRepositoryOutput() RepositoryOutput {
	return o
}

func (o RepositoryOutput) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return o
}

// Optional. If set, configures this repository to be linked to a Git remote.
func (o RepositoryOutput) GitRemoteSettings() GitRemoteSettingsResponseOutput {
	return o.ApplyT(func(v *Repository) GitRemoteSettingsResponseOutput { return v.GitRemoteSettings }).(GitRemoteSettingsResponseOutput)
}

func (o RepositoryOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The repository's name.
func (o RepositoryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional. The name of the Secret Manager secret version to be used to interpolate variables into the .npmrc file for package installation operations. Must be in the format `projects/*/secrets/*/versions/*`. The file itself must be in a JSON format.
func (o RepositoryOutput) NpmrcEnvironmentVariablesSecretVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.NpmrcEnvironmentVariablesSecretVersion }).(pulumi.StringOutput)
}

func (o RepositoryOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Required. The ID to use for the repository, which will become the final component of the repository's resource name.
func (o RepositoryOutput) RepositoryId() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.RepositoryId }).(pulumi.StringOutput)
}

// Optional. If set, fields of `workspace_compilation_overrides` override the default compilation settings that are specified in dataform.json when creating workspace-scoped compilation results. See documentation for `WorkspaceCompilationOverrides` for more information.
func (o RepositoryOutput) WorkspaceCompilationOverrides() WorkspaceCompilationOverridesResponseOutput {
	return o.ApplyT(func(v *Repository) WorkspaceCompilationOverridesResponseOutput {
		return v.WorkspaceCompilationOverrides
	}).(WorkspaceCompilationOverridesResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryInput)(nil)).Elem(), &Repository{})
	pulumi.RegisterOutputType(RepositoryOutput{})
}
