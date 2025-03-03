// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta
{
    /// <summary>
    /// Create an OS Config guest policy.
    /// </summary>
    [GoogleNativeResourceType("google-native:osconfig/v1beta:GuestPolicy")]
    public partial class GuestPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the VM instances that are assigned to this policy. This allows you to target sets or groups of VM instances by different parameters such as labels, names, OS, or zones. If left empty, all VM instances underneath this policy are targeted. At the same level in the resource hierarchy (that is within a project), the service prevents the creation of multiple policies that conflict with each other. For more information, see how the service [handles assignment conflicts](/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
        /// </summary>
        [Output("assignment")]
        public Output<Outputs.AssignmentResponse> Assignment { get; private set; } = null!;

        /// <summary>
        /// Time this guest policy was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Description of the guest policy. Length of the description is limited to 1024 characters.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The etag for this guest policy. If this is provided on update, it must match the server's etag.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Required. The logical name of the guest policy in the project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
        /// </summary>
        [Output("guestPolicyId")]
        public Output<string> GuestPolicyId { get; private set; } = null!;

        /// <summary>
        /// Unique name of the resource in this project using one of the following forms: `projects/{project_number}/guestPolicies/{guest_policy_id}`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of package repositories to configure on the VM instance. This is done before any other configs are applied so they can use these repos. Package repositories are only configured if the corresponding package manager(s) are available.
        /// </summary>
        [Output("packageRepositories")]
        public Output<ImmutableArray<Outputs.PackageRepositoryResponse>> PackageRepositories { get; private set; } = null!;

        /// <summary>
        /// The software packages to be managed by this policy.
        /// </summary>
        [Output("packages")]
        public Output<ImmutableArray<Outputs.PackageResponse>> Packages { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// A list of Recipes to install on the VM instance.
        /// </summary>
        [Output("recipes")]
        public Output<ImmutableArray<Outputs.SoftwareRecipeResponse>> Recipes { get; private set; } = null!;

        /// <summary>
        /// Last time this guest policy was updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a GuestPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GuestPolicy(string name, GuestPolicyArgs args, CustomResourceOptions? options = null)
            : base("google-native:osconfig/v1beta:GuestPolicy", name, args ?? new GuestPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GuestPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:osconfig/v1beta:GuestPolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "guestPolicyId",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GuestPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GuestPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GuestPolicy(name, id, options);
        }
    }

    public sealed class GuestPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the VM instances that are assigned to this policy. This allows you to target sets or groups of VM instances by different parameters such as labels, names, OS, or zones. If left empty, all VM instances underneath this policy are targeted. At the same level in the resource hierarchy (that is within a project), the service prevents the creation of multiple policies that conflict with each other. For more information, see how the service [handles assignment conflicts](/compute/docs/os-config-management/create-guest-policy#handle-conflicts).
        /// </summary>
        [Input("assignment", required: true)]
        public Input<Inputs.AssignmentArgs> Assignment { get; set; } = null!;

        /// <summary>
        /// Description of the guest policy. Length of the description is limited to 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The etag for this guest policy. If this is provided on update, it must match the server's etag.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Required. The logical name of the guest policy in the project with the following restrictions: * Must contain only lowercase letters, numbers, and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the project.
        /// </summary>
        [Input("guestPolicyId", required: true)]
        public Input<string> GuestPolicyId { get; set; } = null!;

        /// <summary>
        /// Unique name of the resource in this project using one of the following forms: `projects/{project_number}/guestPolicies/{guest_policy_id}`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("packageRepositories")]
        private InputList<Inputs.PackageRepositoryArgs>? _packageRepositories;

        /// <summary>
        /// A list of package repositories to configure on the VM instance. This is done before any other configs are applied so they can use these repos. Package repositories are only configured if the corresponding package manager(s) are available.
        /// </summary>
        public InputList<Inputs.PackageRepositoryArgs> PackageRepositories
        {
            get => _packageRepositories ?? (_packageRepositories = new InputList<Inputs.PackageRepositoryArgs>());
            set => _packageRepositories = value;
        }

        [Input("packages")]
        private InputList<Inputs.PackageArgs>? _packages;

        /// <summary>
        /// The software packages to be managed by this policy.
        /// </summary>
        public InputList<Inputs.PackageArgs> Packages
        {
            get => _packages ?? (_packages = new InputList<Inputs.PackageArgs>());
            set => _packages = value;
        }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("recipes")]
        private InputList<Inputs.SoftwareRecipeArgs>? _recipes;

        /// <summary>
        /// A list of Recipes to install on the VM instance.
        /// </summary>
        public InputList<Inputs.SoftwareRecipeArgs> Recipes
        {
            get => _recipes ?? (_recipes = new InputList<Inputs.SoftwareRecipeArgs>());
            set => _recipes = value;
        }

        public GuestPolicyArgs()
        {
        }
        public static new GuestPolicyArgs Empty => new GuestPolicyArgs();
    }
}
