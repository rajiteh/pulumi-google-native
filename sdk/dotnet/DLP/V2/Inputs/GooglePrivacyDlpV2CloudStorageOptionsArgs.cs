// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Inputs
{

    /// <summary>
    /// Options defining a file or a set of files within a Cloud Storage bucket.
    /// </summary>
    public sealed class GooglePrivacyDlpV2CloudStorageOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Max number of bytes to scan from a file. If a scanned file's size is bigger than this value then the rest of the bytes are omitted. Only one of bytes_limit_per_file and bytes_limit_per_file_percent can be specified. Cannot be set if de-identification is requested.
        /// </summary>
        [Input("bytesLimitPerFile")]
        public Input<string>? BytesLimitPerFile { get; set; }

        /// <summary>
        /// Max percentage of bytes to scan from a file. The rest are omitted. The number of bytes scanned is rounded down. Must be between 0 and 100, inclusively. Both 0 and 100 means no limit. Defaults to 0. Only one of bytes_limit_per_file and bytes_limit_per_file_percent can be specified. Cannot be set if de-identification is requested.
        /// </summary>
        [Input("bytesLimitPerFilePercent")]
        public Input<int>? BytesLimitPerFilePercent { get; set; }

        /// <summary>
        /// The set of one or more files to scan.
        /// </summary>
        [Input("fileSet")]
        public Input<Inputs.GooglePrivacyDlpV2FileSetArgs>? FileSet { get; set; }

        [Input("fileTypes")]
        private InputList<Pulumi.GoogleNative.DLP.V2.GooglePrivacyDlpV2CloudStorageOptionsFileTypesItem>? _fileTypes;

        /// <summary>
        /// List of file type groups to include in the scan. If empty, all files are scanned and available data format processors are applied. In addition, the binary content of the selected files is always scanned as well. Images are scanned only as binary if the specified region does not support image inspection and no file_types were specified. Image inspection is restricted to 'global', 'us', 'asia', and 'europe'.
        /// </summary>
        public InputList<Pulumi.GoogleNative.DLP.V2.GooglePrivacyDlpV2CloudStorageOptionsFileTypesItem> FileTypes
        {
            get => _fileTypes ?? (_fileTypes = new InputList<Pulumi.GoogleNative.DLP.V2.GooglePrivacyDlpV2CloudStorageOptionsFileTypesItem>());
            set => _fileTypes = value;
        }

        /// <summary>
        /// Limits the number of files to scan to this percentage of the input FileSet. Number of files scanned is rounded down. Must be between 0 and 100, inclusively. Both 0 and 100 means no limit. Defaults to 0.
        /// </summary>
        [Input("filesLimitPercent")]
        public Input<int>? FilesLimitPercent { get; set; }

        [Input("sampleMethod")]
        public Input<Pulumi.GoogleNative.DLP.V2.GooglePrivacyDlpV2CloudStorageOptionsSampleMethod>? SampleMethod { get; set; }

        public GooglePrivacyDlpV2CloudStorageOptionsArgs()
        {
        }
        public static new GooglePrivacyDlpV2CloudStorageOptionsArgs Empty => new GooglePrivacyDlpV2CloudStorageOptionsArgs();
    }
}
