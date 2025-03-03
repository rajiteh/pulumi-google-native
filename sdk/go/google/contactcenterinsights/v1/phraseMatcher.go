// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a phrase matcher.
type PhraseMatcher struct {
	pulumi.CustomResourceState

	// The most recent time at which the activation status was updated.
	ActivationUpdateTime pulumi.StringOutput `pulumi:"activationUpdateTime"`
	// Applies the phrase matcher only when it is active.
	Active pulumi.BoolOutput `pulumi:"active"`
	// The human-readable name of the phrase matcher.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	Location    pulumi.StringOutput `pulumi:"location"`
	// The resource name of the phrase matcher. Format: projects/{project}/locations/{location}/phraseMatchers/{phrase_matcher}
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of phase match rule groups that are included in this matcher.
	PhraseMatchRuleGroups GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupResponseArrayOutput `pulumi:"phraseMatchRuleGroups"`
	Project               pulumi.StringOutput                                                       `pulumi:"project"`
	// The timestamp of when the revision was created. It is also the create time when a new matcher is added.
	RevisionCreateTime pulumi.StringOutput `pulumi:"revisionCreateTime"`
	// Immutable. The revision ID of the phrase matcher. A new revision is committed whenever the matcher is changed, except when it is activated or deactivated. A server generated random ID will be used. Example: locations/global/phraseMatchers/my-first-matcher@1234567
	RevisionId pulumi.StringOutput `pulumi:"revisionId"`
	// The role whose utterances the phrase matcher should be matched against. If the role is ROLE_UNSPECIFIED it will be matched against any utterances in the transcript.
	RoleMatch pulumi.StringOutput `pulumi:"roleMatch"`
	// The type of this phrase matcher.
	Type pulumi.StringOutput `pulumi:"type"`
	// The most recent time at which the phrase matcher was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// The customized version tag to use for the phrase matcher. If not specified, it will default to `revision_id`.
	VersionTag pulumi.StringOutput `pulumi:"versionTag"`
}

// NewPhraseMatcher registers a new resource with the given unique name, arguments, and options.
func NewPhraseMatcher(ctx *pulumi.Context,
	name string, args *PhraseMatcherArgs, opts ...pulumi.ResourceOption) (*PhraseMatcher, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource PhraseMatcher
	err := ctx.RegisterResource("google-native:contactcenterinsights/v1:PhraseMatcher", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPhraseMatcher gets an existing PhraseMatcher resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPhraseMatcher(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PhraseMatcherState, opts ...pulumi.ResourceOption) (*PhraseMatcher, error) {
	var resource PhraseMatcher
	err := ctx.ReadResource("google-native:contactcenterinsights/v1:PhraseMatcher", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PhraseMatcher resources.
type phraseMatcherState struct {
}

type PhraseMatcherState struct {
}

func (PhraseMatcherState) ElementType() reflect.Type {
	return reflect.TypeOf((*phraseMatcherState)(nil)).Elem()
}

type phraseMatcherArgs struct {
	// Applies the phrase matcher only when it is active.
	Active *bool `pulumi:"active"`
	// The human-readable name of the phrase matcher.
	DisplayName *string `pulumi:"displayName"`
	Location    *string `pulumi:"location"`
	// The resource name of the phrase matcher. Format: projects/{project}/locations/{location}/phraseMatchers/{phrase_matcher}
	Name *string `pulumi:"name"`
	// A list of phase match rule groups that are included in this matcher.
	PhraseMatchRuleGroups []GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroup `pulumi:"phraseMatchRuleGroups"`
	Project               *string                                                  `pulumi:"project"`
	// The role whose utterances the phrase matcher should be matched against. If the role is ROLE_UNSPECIFIED it will be matched against any utterances in the transcript.
	RoleMatch *PhraseMatcherRoleMatch `pulumi:"roleMatch"`
	// The type of this phrase matcher.
	Type PhraseMatcherType `pulumi:"type"`
	// The customized version tag to use for the phrase matcher. If not specified, it will default to `revision_id`.
	VersionTag *string `pulumi:"versionTag"`
}

// The set of arguments for constructing a PhraseMatcher resource.
type PhraseMatcherArgs struct {
	// Applies the phrase matcher only when it is active.
	Active pulumi.BoolPtrInput
	// The human-readable name of the phrase matcher.
	DisplayName pulumi.StringPtrInput
	Location    pulumi.StringPtrInput
	// The resource name of the phrase matcher. Format: projects/{project}/locations/{location}/phraseMatchers/{phrase_matcher}
	Name pulumi.StringPtrInput
	// A list of phase match rule groups that are included in this matcher.
	PhraseMatchRuleGroups GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupArrayInput
	Project               pulumi.StringPtrInput
	// The role whose utterances the phrase matcher should be matched against. If the role is ROLE_UNSPECIFIED it will be matched against any utterances in the transcript.
	RoleMatch PhraseMatcherRoleMatchPtrInput
	// The type of this phrase matcher.
	Type PhraseMatcherTypeInput
	// The customized version tag to use for the phrase matcher. If not specified, it will default to `revision_id`.
	VersionTag pulumi.StringPtrInput
}

func (PhraseMatcherArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*phraseMatcherArgs)(nil)).Elem()
}

type PhraseMatcherInput interface {
	pulumi.Input

	ToPhraseMatcherOutput() PhraseMatcherOutput
	ToPhraseMatcherOutputWithContext(ctx context.Context) PhraseMatcherOutput
}

func (*PhraseMatcher) ElementType() reflect.Type {
	return reflect.TypeOf((**PhraseMatcher)(nil)).Elem()
}

func (i *PhraseMatcher) ToPhraseMatcherOutput() PhraseMatcherOutput {
	return i.ToPhraseMatcherOutputWithContext(context.Background())
}

func (i *PhraseMatcher) ToPhraseMatcherOutputWithContext(ctx context.Context) PhraseMatcherOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PhraseMatcherOutput)
}

type PhraseMatcherOutput struct{ *pulumi.OutputState }

func (PhraseMatcherOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PhraseMatcher)(nil)).Elem()
}

func (o PhraseMatcherOutput) ToPhraseMatcherOutput() PhraseMatcherOutput {
	return o
}

func (o PhraseMatcherOutput) ToPhraseMatcherOutputWithContext(ctx context.Context) PhraseMatcherOutput {
	return o
}

// The most recent time at which the activation status was updated.
func (o PhraseMatcherOutput) ActivationUpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *PhraseMatcher) pulumi.StringOutput { return v.ActivationUpdateTime }).(pulumi.StringOutput)
}

// Applies the phrase matcher only when it is active.
func (o PhraseMatcherOutput) Active() pulumi.BoolOutput {
	return o.ApplyT(func(v *PhraseMatcher) pulumi.BoolOutput { return v.Active }).(pulumi.BoolOutput)
}

// The human-readable name of the phrase matcher.
func (o PhraseMatcherOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *PhraseMatcher) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o PhraseMatcherOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PhraseMatcher) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the phrase matcher. Format: projects/{project}/locations/{location}/phraseMatchers/{phrase_matcher}
func (o PhraseMatcherOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PhraseMatcher) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A list of phase match rule groups that are included in this matcher.
func (o PhraseMatcherOutput) PhraseMatchRuleGroups() GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupResponseArrayOutput {
	return o.ApplyT(func(v *PhraseMatcher) GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupResponseArrayOutput {
		return v.PhraseMatchRuleGroups
	}).(GoogleCloudContactcenterinsightsV1PhraseMatchRuleGroupResponseArrayOutput)
}

func (o PhraseMatcherOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *PhraseMatcher) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The timestamp of when the revision was created. It is also the create time when a new matcher is added.
func (o PhraseMatcherOutput) RevisionCreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *PhraseMatcher) pulumi.StringOutput { return v.RevisionCreateTime }).(pulumi.StringOutput)
}

// Immutable. The revision ID of the phrase matcher. A new revision is committed whenever the matcher is changed, except when it is activated or deactivated. A server generated random ID will be used. Example: locations/global/phraseMatchers/my-first-matcher@1234567
func (o PhraseMatcherOutput) RevisionId() pulumi.StringOutput {
	return o.ApplyT(func(v *PhraseMatcher) pulumi.StringOutput { return v.RevisionId }).(pulumi.StringOutput)
}

// The role whose utterances the phrase matcher should be matched against. If the role is ROLE_UNSPECIFIED it will be matched against any utterances in the transcript.
func (o PhraseMatcherOutput) RoleMatch() pulumi.StringOutput {
	return o.ApplyT(func(v *PhraseMatcher) pulumi.StringOutput { return v.RoleMatch }).(pulumi.StringOutput)
}

// The type of this phrase matcher.
func (o PhraseMatcherOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PhraseMatcher) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The most recent time at which the phrase matcher was updated.
func (o PhraseMatcherOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *PhraseMatcher) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// The customized version tag to use for the phrase matcher. If not specified, it will default to `revision_id`.
func (o PhraseMatcherOutput) VersionTag() pulumi.StringOutput {
	return o.ApplyT(func(v *PhraseMatcher) pulumi.StringOutput { return v.VersionTag }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PhraseMatcherInput)(nil)).Elem(), &PhraseMatcher{})
	pulumi.RegisterOutputType(PhraseMatcherOutput{})
}
