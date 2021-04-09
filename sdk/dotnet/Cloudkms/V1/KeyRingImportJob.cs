// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.Cloudkms.V1
{
    /// <summary>
    /// Create a new ImportJob within a KeyRing. ImportJob.import_method is required.
    /// </summary>
    [GcpNativeResourceType("gcp-native:cloudkms/v1:KeyRingImportJob")]
    public partial class KeyRingImportJob : Pulumi.CustomResource
    {
        /// <summary>
        /// Statement that was generated and signed by the key creator (for example, an HSM) at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google. Only present if the chosen ImportMethod is one with a protection level of HSM.
        /// </summary>
        [Output("attestation")]
        public Output<Outputs.KeyOperationAttestationResponse> Attestation { get; private set; } = null!;

        /// <summary>
        /// The time at which this ImportJob was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The time this ImportJob expired. Only present if state is EXPIRED.
        /// </summary>
        [Output("expireEventTime")]
        public Output<string> ExpireEventTime { get; private set; } = null!;

        /// <summary>
        /// The time at which this ImportJob is scheduled for expiration and can no longer be used to import key material.
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// The time this ImportJob's key material was generated.
        /// </summary>
        [Output("generateTime")]
        public Output<string> GenerateTime { get; private set; } = null!;

        /// <summary>
        /// Required. Immutable. The wrapping method to be used for incoming key material.
        /// </summary>
        [Output("importMethod")]
        public Output<string> ImportMethod { get; private set; } = null!;

        /// <summary>
        /// The resource name for this ImportJob in the format `projects/*/locations/*/keyRings/*/importJobs/*`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Required. Immutable. The protection level of the ImportJob. This must match the protection_level of the version_template on the CryptoKey you attempt to import into.
        /// </summary>
        [Output("protectionLevel")]
        public Output<string> ProtectionLevel { get; private set; } = null!;

        /// <summary>
        /// The public key with which to wrap key material prior to import. Only returned if state is ACTIVE.
        /// </summary>
        [Output("publicKey")]
        public Output<Outputs.WrappingPublicKeyResponse> PublicKey { get; private set; } = null!;

        /// <summary>
        /// The current state of the ImportJob, indicating if it can be used.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;


        /// <summary>
        /// Create a KeyRingImportJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KeyRingImportJob(string name, KeyRingImportJobArgs args, CustomResourceOptions? options = null)
            : base("gcp-native:cloudkms/v1:KeyRingImportJob", name, args ?? new KeyRingImportJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KeyRingImportJob(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("gcp-native:cloudkms/v1:KeyRingImportJob", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing KeyRingImportJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KeyRingImportJob Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new KeyRingImportJob(name, id, options);
        }
    }

    public sealed class KeyRingImportJobArgs : Pulumi.ResourceArgs
    {
        [Input("importJobsId", required: true)]
        public Input<string> ImportJobsId { get; set; } = null!;

        /// <summary>
        /// Required. Immutable. The wrapping method to be used for incoming key material.
        /// </summary>
        [Input("importMethod")]
        public Input<string>? ImportMethod { get; set; }

        [Input("keyRingsId", required: true)]
        public Input<string> KeyRingsId { get; set; } = null!;

        [Input("locationsId", required: true)]
        public Input<string> LocationsId { get; set; } = null!;

        [Input("projectsId", required: true)]
        public Input<string> ProjectsId { get; set; } = null!;

        /// <summary>
        /// Required. Immutable. The protection level of the ImportJob. This must match the protection_level of the version_template on the CryptoKey you attempt to import into.
        /// </summary>
        [Input("protectionLevel")]
        public Input<string>? ProtectionLevel { get; set; }

        public KeyRingImportJobArgs()
        {
        }
    }
}