// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// The share setting for reservations and sole tenancy node groups.
    /// </summary>
    public sealed class ShareSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("folderMap")]
        private InputMap<string>? _folderMap;

        /// <summary>
        /// A map of folder id and folder config to specify consumer projects for this shared-reservation. This is only valid when share_type's value is DIRECT_PROJECTS_UNDER_SPECIFIC_FOLDERS. Folder id should be a string of number, and without "folders/" prefix.
        /// </summary>
        public InputMap<string> FolderMap
        {
            get => _folderMap ?? (_folderMap = new InputMap<string>());
            set => _folderMap = value;
        }

        [Input("projectMap")]
        private InputMap<string>? _projectMap;

        /// <summary>
        /// A map of project id and project config. This is only valid when share_type's value is SPECIFIC_PROJECTS.
        /// </summary>
        public InputMap<string> ProjectMap
        {
            get => _projectMap ?? (_projectMap = new InputMap<string>());
            set => _projectMap = value;
        }

        [Input("projects")]
        private InputList<string>? _projects;

        /// <summary>
        /// A List of Project names to specify consumer projects for this shared-reservation. This is only valid when share_type's value is SPECIFIC_PROJECTS.
        /// </summary>
        public InputList<string> Projects
        {
            get => _projects ?? (_projects = new InputList<string>());
            set => _projects = value;
        }

        /// <summary>
        /// Type of sharing for this shared-reservation
        /// </summary>
        [Input("shareType")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.ShareSettingsShareType>? ShareType { get; set; }

        public ShareSettingsArgs()
        {
        }
        public static new ShareSettingsArgs Empty => new ShareSettingsArgs();
    }
}
