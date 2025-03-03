// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new note.
// Auto-naming is currently not supported for this resource.
type Note struct {
	pulumi.CustomResourceState

	// A note describing an attestation role.
	AttestationAuthority AuthorityResponseOutput `pulumi:"attestationAuthority"`
	// A note describing a base image.
	BaseImage BasisResponseOutput `pulumi:"baseImage"`
	// A note describing build provenance for a verifiable build.
	Build BuildResponseOutput `pulumi:"build"`
	// The time this note was created. This field can be used as a filter in list requests.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A note describing something that can be deployed.
	Deployable DeployableResponseOutput `pulumi:"deployable"`
	// A note describing the initial analysis of a resource.
	Discovery DiscoveryResponseOutput `pulumi:"discovery"`
	// Time of expiration for this note. Empty if note does not expire.
	ExpirationTime pulumi.StringOutput `pulumi:"expirationTime"`
	// A note describing an in-toto link.
	Intoto InTotoResponseOutput `pulumi:"intoto"`
	// The type of analysis. This field can be used as a filter in list requests.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A detailed description of this note.
	LongDescription pulumi.StringOutput `pulumi:"longDescription"`
	// The name of the note in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Required. The ID to use for this note.
	NoteId pulumi.StringOutput `pulumi:"noteId"`
	// A note describing a package hosted by various package managers.
	Package PackageResponseOutput `pulumi:"package"`
	Project pulumi.StringOutput   `pulumi:"project"`
	// Other notes related to this note.
	RelatedNoteNames pulumi.StringArrayOutput `pulumi:"relatedNoteNames"`
	// URLs associated with this note.
	RelatedUrl RelatedUrlResponseArrayOutput `pulumi:"relatedUrl"`
	// A note describing a software bill of materials.
	Sbom DocumentNoteResponseOutput `pulumi:"sbom"`
	// A note describing an SBOM reference.
	SbomReference SBOMReferenceNoteResponseOutput `pulumi:"sbomReference"`
	// A one sentence description of this note.
	ShortDescription pulumi.StringOutput `pulumi:"shortDescription"`
	// A note describing an SPDX File.
	SpdxFile FileNoteResponseOutput `pulumi:"spdxFile"`
	// A note describing an SPDX Package.
	SpdxPackage PackageInfoNoteResponseOutput `pulumi:"spdxPackage"`
	// A note describing an SPDX File.
	SpdxRelationship RelationshipNoteResponseOutput `pulumi:"spdxRelationship"`
	// The time this note was last updated. This field can be used as a filter in list requests.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// A note describing a package vulnerability.
	Vulnerability VulnerabilityResponseOutput `pulumi:"vulnerability"`
	// A note describing a vulnerability assessment.
	VulnerabilityAssessment VulnerabilityAssessmentNoteResponseOutput `pulumi:"vulnerabilityAssessment"`
}

// NewNote registers a new resource with the given unique name, arguments, and options.
func NewNote(ctx *pulumi.Context,
	name string, args *NoteArgs, opts ...pulumi.ResourceOption) (*Note, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NoteId == nil {
		return nil, errors.New("invalid value for required argument 'NoteId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"noteId",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource Note
	err := ctx.RegisterResource("google-native:containeranalysis/v1beta1:Note", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNote gets an existing Note resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNote(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NoteState, opts ...pulumi.ResourceOption) (*Note, error) {
	var resource Note
	err := ctx.ReadResource("google-native:containeranalysis/v1beta1:Note", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Note resources.
type noteState struct {
}

type NoteState struct {
}

func (NoteState) ElementType() reflect.Type {
	return reflect.TypeOf((*noteState)(nil)).Elem()
}

type noteArgs struct {
	// A note describing an attestation role.
	AttestationAuthority *Authority `pulumi:"attestationAuthority"`
	// A note describing a base image.
	BaseImage *Basis `pulumi:"baseImage"`
	// A note describing build provenance for a verifiable build.
	Build *Build `pulumi:"build"`
	// A note describing something that can be deployed.
	Deployable *Deployable `pulumi:"deployable"`
	// A note describing the initial analysis of a resource.
	Discovery *Discovery `pulumi:"discovery"`
	// Time of expiration for this note. Empty if note does not expire.
	ExpirationTime *string `pulumi:"expirationTime"`
	// A note describing an in-toto link.
	Intoto *InToto `pulumi:"intoto"`
	// A detailed description of this note.
	LongDescription *string `pulumi:"longDescription"`
	// Required. The ID to use for this note.
	NoteId string `pulumi:"noteId"`
	// A note describing a package hosted by various package managers.
	Package *Package `pulumi:"package"`
	Project *string  `pulumi:"project"`
	// Other notes related to this note.
	RelatedNoteNames []string `pulumi:"relatedNoteNames"`
	// URLs associated with this note.
	RelatedUrl []RelatedUrl `pulumi:"relatedUrl"`
	// A note describing a software bill of materials.
	Sbom *DocumentNote `pulumi:"sbom"`
	// A note describing an SBOM reference.
	SbomReference *SBOMReferenceNote `pulumi:"sbomReference"`
	// A one sentence description of this note.
	ShortDescription *string `pulumi:"shortDescription"`
	// A note describing an SPDX File.
	SpdxFile *FileNote `pulumi:"spdxFile"`
	// A note describing an SPDX Package.
	SpdxPackage *PackageInfoNote `pulumi:"spdxPackage"`
	// A note describing an SPDX File.
	SpdxRelationship *RelationshipNote `pulumi:"spdxRelationship"`
	// A note describing a package vulnerability.
	Vulnerability *Vulnerability `pulumi:"vulnerability"`
	// A note describing a vulnerability assessment.
	VulnerabilityAssessment *VulnerabilityAssessmentNote `pulumi:"vulnerabilityAssessment"`
}

// The set of arguments for constructing a Note resource.
type NoteArgs struct {
	// A note describing an attestation role.
	AttestationAuthority AuthorityPtrInput
	// A note describing a base image.
	BaseImage BasisPtrInput
	// A note describing build provenance for a verifiable build.
	Build BuildPtrInput
	// A note describing something that can be deployed.
	Deployable DeployablePtrInput
	// A note describing the initial analysis of a resource.
	Discovery DiscoveryPtrInput
	// Time of expiration for this note. Empty if note does not expire.
	ExpirationTime pulumi.StringPtrInput
	// A note describing an in-toto link.
	Intoto InTotoPtrInput
	// A detailed description of this note.
	LongDescription pulumi.StringPtrInput
	// Required. The ID to use for this note.
	NoteId pulumi.StringInput
	// A note describing a package hosted by various package managers.
	Package PackagePtrInput
	Project pulumi.StringPtrInput
	// Other notes related to this note.
	RelatedNoteNames pulumi.StringArrayInput
	// URLs associated with this note.
	RelatedUrl RelatedUrlArrayInput
	// A note describing a software bill of materials.
	Sbom DocumentNotePtrInput
	// A note describing an SBOM reference.
	SbomReference SBOMReferenceNotePtrInput
	// A one sentence description of this note.
	ShortDescription pulumi.StringPtrInput
	// A note describing an SPDX File.
	SpdxFile FileNotePtrInput
	// A note describing an SPDX Package.
	SpdxPackage PackageInfoNotePtrInput
	// A note describing an SPDX File.
	SpdxRelationship RelationshipNotePtrInput
	// A note describing a package vulnerability.
	Vulnerability VulnerabilityPtrInput
	// A note describing a vulnerability assessment.
	VulnerabilityAssessment VulnerabilityAssessmentNotePtrInput
}

func (NoteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*noteArgs)(nil)).Elem()
}

type NoteInput interface {
	pulumi.Input

	ToNoteOutput() NoteOutput
	ToNoteOutputWithContext(ctx context.Context) NoteOutput
}

func (*Note) ElementType() reflect.Type {
	return reflect.TypeOf((**Note)(nil)).Elem()
}

func (i *Note) ToNoteOutput() NoteOutput {
	return i.ToNoteOutputWithContext(context.Background())
}

func (i *Note) ToNoteOutputWithContext(ctx context.Context) NoteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoteOutput)
}

type NoteOutput struct{ *pulumi.OutputState }

func (NoteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Note)(nil)).Elem()
}

func (o NoteOutput) ToNoteOutput() NoteOutput {
	return o
}

func (o NoteOutput) ToNoteOutputWithContext(ctx context.Context) NoteOutput {
	return o
}

// A note describing an attestation role.
func (o NoteOutput) AttestationAuthority() AuthorityResponseOutput {
	return o.ApplyT(func(v *Note) AuthorityResponseOutput { return v.AttestationAuthority }).(AuthorityResponseOutput)
}

// A note describing a base image.
func (o NoteOutput) BaseImage() BasisResponseOutput {
	return o.ApplyT(func(v *Note) BasisResponseOutput { return v.BaseImage }).(BasisResponseOutput)
}

// A note describing build provenance for a verifiable build.
func (o NoteOutput) Build() BuildResponseOutput {
	return o.ApplyT(func(v *Note) BuildResponseOutput { return v.Build }).(BuildResponseOutput)
}

// The time this note was created. This field can be used as a filter in list requests.
func (o NoteOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Note) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A note describing something that can be deployed.
func (o NoteOutput) Deployable() DeployableResponseOutput {
	return o.ApplyT(func(v *Note) DeployableResponseOutput { return v.Deployable }).(DeployableResponseOutput)
}

// A note describing the initial analysis of a resource.
func (o NoteOutput) Discovery() DiscoveryResponseOutput {
	return o.ApplyT(func(v *Note) DiscoveryResponseOutput { return v.Discovery }).(DiscoveryResponseOutput)
}

// Time of expiration for this note. Empty if note does not expire.
func (o NoteOutput) ExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Note) pulumi.StringOutput { return v.ExpirationTime }).(pulumi.StringOutput)
}

// A note describing an in-toto link.
func (o NoteOutput) Intoto() InTotoResponseOutput {
	return o.ApplyT(func(v *Note) InTotoResponseOutput { return v.Intoto }).(InTotoResponseOutput)
}

// The type of analysis. This field can be used as a filter in list requests.
func (o NoteOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Note) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// A detailed description of this note.
func (o NoteOutput) LongDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *Note) pulumi.StringOutput { return v.LongDescription }).(pulumi.StringOutput)
}

// The name of the note in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`.
func (o NoteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Note) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Required. The ID to use for this note.
func (o NoteOutput) NoteId() pulumi.StringOutput {
	return o.ApplyT(func(v *Note) pulumi.StringOutput { return v.NoteId }).(pulumi.StringOutput)
}

// A note describing a package hosted by various package managers.
func (o NoteOutput) Package() PackageResponseOutput {
	return o.ApplyT(func(v *Note) PackageResponseOutput { return v.Package }).(PackageResponseOutput)
}

func (o NoteOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Note) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Other notes related to this note.
func (o NoteOutput) RelatedNoteNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Note) pulumi.StringArrayOutput { return v.RelatedNoteNames }).(pulumi.StringArrayOutput)
}

// URLs associated with this note.
func (o NoteOutput) RelatedUrl() RelatedUrlResponseArrayOutput {
	return o.ApplyT(func(v *Note) RelatedUrlResponseArrayOutput { return v.RelatedUrl }).(RelatedUrlResponseArrayOutput)
}

// A note describing a software bill of materials.
func (o NoteOutput) Sbom() DocumentNoteResponseOutput {
	return o.ApplyT(func(v *Note) DocumentNoteResponseOutput { return v.Sbom }).(DocumentNoteResponseOutput)
}

// A note describing an SBOM reference.
func (o NoteOutput) SbomReference() SBOMReferenceNoteResponseOutput {
	return o.ApplyT(func(v *Note) SBOMReferenceNoteResponseOutput { return v.SbomReference }).(SBOMReferenceNoteResponseOutput)
}

// A one sentence description of this note.
func (o NoteOutput) ShortDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *Note) pulumi.StringOutput { return v.ShortDescription }).(pulumi.StringOutput)
}

// A note describing an SPDX File.
func (o NoteOutput) SpdxFile() FileNoteResponseOutput {
	return o.ApplyT(func(v *Note) FileNoteResponseOutput { return v.SpdxFile }).(FileNoteResponseOutput)
}

// A note describing an SPDX Package.
func (o NoteOutput) SpdxPackage() PackageInfoNoteResponseOutput {
	return o.ApplyT(func(v *Note) PackageInfoNoteResponseOutput { return v.SpdxPackage }).(PackageInfoNoteResponseOutput)
}

// A note describing an SPDX File.
func (o NoteOutput) SpdxRelationship() RelationshipNoteResponseOutput {
	return o.ApplyT(func(v *Note) RelationshipNoteResponseOutput { return v.SpdxRelationship }).(RelationshipNoteResponseOutput)
}

// The time this note was last updated. This field can be used as a filter in list requests.
func (o NoteOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Note) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// A note describing a package vulnerability.
func (o NoteOutput) Vulnerability() VulnerabilityResponseOutput {
	return o.ApplyT(func(v *Note) VulnerabilityResponseOutput { return v.Vulnerability }).(VulnerabilityResponseOutput)
}

// A note describing a vulnerability assessment.
func (o NoteOutput) VulnerabilityAssessment() VulnerabilityAssessmentNoteResponseOutput {
	return o.ApplyT(func(v *Note) VulnerabilityAssessmentNoteResponseOutput { return v.VulnerabilityAssessment }).(VulnerabilityAssessmentNoteResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NoteInput)(nil)).Elem(), &Note{})
	pulumi.RegisterOutputType(NoteOutput{})
}
