// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.IAM.V1
{
    /// <summary>
    /// Creates a new custom Role.
    /// </summary>
    [GoogleNativeResourceType("google-native:iam/v1:Role")]
    public partial class Role : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The current deleted state of the role. This field is read only. It will be ignored in calls to CreateRole and UpdateRole.
        /// </summary>
        [Output("deleted")]
        public Output<bool> Deleted { get; private set; } = null!;

        /// <summary>
        /// Optional. A human-readable description for the role.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Used to perform a consistent read-modify-write.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The names of the permissions this role grants when bound in an IAM policy.
        /// </summary>
        [Output("includedPermissions")]
        public Output<ImmutableArray<string>> IncludedPermissions { get; private set; } = null!;

        /// <summary>
        /// The name of the role. When Role is used in CreateRole, the role name must not be set. When Role is used in output and other input such as UpdateRole, the role name is the complete path, e.g., roles/logging.viewer for predefined roles and organizations/{ORGANIZATION_ID}/roles/logging.viewer for custom roles.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The current launch stage of the role. If the `ALPHA` launch stage has been selected for a role, the `stage` field will not be included in the returned definition for the role.
        /// </summary>
        [Output("stage")]
        public Output<string> Stage { get; private set; } = null!;

        /// <summary>
        /// Optional. A human-readable title for the role. Typically this is limited to 100 UTF-8 bytes.
        /// </summary>
        [Output("title")]
        public Output<string> Title { get; private set; } = null!;


        /// <summary>
        /// Create a Role resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Role(string name, RoleArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:iam/v1:Role", name, args ?? new RoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Role(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:iam/v1:Role", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Role resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Role Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Role(name, id, options);
        }
    }

    public sealed class RoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The current deleted state of the role. This field is read only. It will be ignored in calls to CreateRole and UpdateRole.
        /// </summary>
        [Input("deleted")]
        public Input<bool>? Deleted { get; set; }

        /// <summary>
        /// Optional. A human-readable description for the role.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Used to perform a consistent read-modify-write.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("includedPermissions")]
        private InputList<string>? _includedPermissions;

        /// <summary>
        /// The names of the permissions this role grants when bound in an IAM policy.
        /// </summary>
        public InputList<string> IncludedPermissions
        {
            get => _includedPermissions ?? (_includedPermissions = new InputList<string>());
            set => _includedPermissions = value;
        }

        /// <summary>
        /// The name of the role. When Role is used in CreateRole, the role name must not be set. When Role is used in output and other input such as UpdateRole, the role name is the complete path, e.g., roles/logging.viewer for predefined roles and organizations/{ORGANIZATION_ID}/roles/logging.viewer for custom roles.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role ID to use for this role. A role ID may contain alphanumeric characters, underscores (`_`), and periods (`.`). It must contain a minimum of 3 characters and a maximum of 64 characters.
        /// </summary>
        [Input("roleId")]
        public Input<string>? RoleId { get; set; }

        /// <summary>
        /// The current launch stage of the role. If the `ALPHA` launch stage has been selected for a role, the `stage` field will not be included in the returned definition for the role.
        /// </summary>
        [Input("stage")]
        public Input<Pulumi.GoogleNative.IAM.V1.RoleStage>? Stage { get; set; }

        /// <summary>
        /// Optional. A human-readable title for the role. Typically this is limited to 100 UTF-8 bytes.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public RoleArgs()
        {
        }
        public static new RoleArgs Empty => new RoleArgs();
    }
}
