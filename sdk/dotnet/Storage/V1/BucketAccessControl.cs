// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Storage.V1
{
    /// <summary>
    /// Creates a new ACL entry on the specified bucket.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:storage/v1:BucketAccessControl")]
    public partial class BucketAccessControl : global::Pulumi.CustomResource
    {
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// The domain associated with the entity, if any.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// The email address associated with the entity, if any.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// The entity holding the permission, in one of the following forms: 
        /// - user-userId 
        /// - user-email 
        /// - group-groupId 
        /// - group-email 
        /// - domain-domain 
        /// - project-team-projectId 
        /// - allUsers 
        /// - allAuthenticatedUsers Examples: 
        /// - The user liz@example.com would be user-liz@example.com. 
        /// - The group example@googlegroups.com would be group-example@googlegroups.com. 
        /// - To refer to all members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
        /// </summary>
        [Output("entity")]
        public Output<string> Entity { get; private set; } = null!;

        /// <summary>
        /// The ID for the entity, if any.
        /// </summary>
        [Output("entityId")]
        public Output<string> EntityId { get; private set; } = null!;

        /// <summary>
        /// HTTP 1.1 Entity tag for the access-control entry.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The kind of item this is. For bucket access control entries, this is always storage#bucketAccessControl.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The project team associated with the entity, if any.
        /// </summary>
        [Output("projectTeam")]
        public Output<Outputs.BucketAccessControlProjectTeamResponse> ProjectTeam { get; private set; } = null!;

        /// <summary>
        /// The access permission for the entity.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// The link to this access-control entry.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// The project to be billed for this request. Required for Requester Pays buckets.
        /// </summary>
        [Output("userProject")]
        public Output<string?> UserProject { get; private set; } = null!;


        /// <summary>
        /// Create a BucketAccessControl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BucketAccessControl(string name, BucketAccessControlArgs args, CustomResourceOptions? options = null)
            : base("google-native:storage/v1:BucketAccessControl", name, args ?? new BucketAccessControlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BucketAccessControl(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:storage/v1:BucketAccessControl", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "bucket",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BucketAccessControl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BucketAccessControl Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BucketAccessControl(name, id, options);
        }
    }

    public sealed class BucketAccessControlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// The domain associated with the entity, if any.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The email address associated with the entity, if any.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// The entity holding the permission, in one of the following forms: 
        /// - user-userId 
        /// - user-email 
        /// - group-groupId 
        /// - group-email 
        /// - domain-domain 
        /// - project-team-projectId 
        /// - allUsers 
        /// - allAuthenticatedUsers Examples: 
        /// - The user liz@example.com would be user-liz@example.com. 
        /// - The group example@googlegroups.com would be group-example@googlegroups.com. 
        /// - To refer to all members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
        /// </summary>
        [Input("entity")]
        public Input<string>? Entity { get; set; }

        /// <summary>
        /// The ID for the entity, if any.
        /// </summary>
        [Input("entityId")]
        public Input<string>? EntityId { get; set; }

        /// <summary>
        /// HTTP 1.1 Entity tag for the access-control entry.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The ID of the access-control entry.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The kind of item this is. For bucket access control entries, this is always storage#bucketAccessControl.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// The project team associated with the entity, if any.
        /// </summary>
        [Input("projectTeam")]
        public Input<Inputs.BucketAccessControlProjectTeamArgs>? ProjectTeam { get; set; }

        /// <summary>
        /// The access permission for the entity.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The link to this access-control entry.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// The project to be billed for this request. Required for Requester Pays buckets.
        /// </summary>
        [Input("userProject")]
        public Input<string>? UserProject { get; set; }

        public BucketAccessControlArgs()
        {
        }
        public static new BucketAccessControlArgs Empty => new BucketAccessControlArgs();
    }
}
