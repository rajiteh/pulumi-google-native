// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.Testing.V1.Outputs
{

    [OutputType]
    public sealed class IosDeviceFileResponse
    {
        /// <summary>
        /// The bundle id of the app where this file lives. iOS apps sandbox their own filesystem, so app files must specify which app installed on the device.
        /// </summary>
        public readonly string BundleId;
        /// <summary>
        /// The source file
        /// </summary>
        public readonly Outputs.FileReferenceResponse Content;
        /// <summary>
        /// Location of the file on the device, inside the app's sandboxed filesystem
        /// </summary>
        public readonly string DevicePath;

        [OutputConstructor]
        private IosDeviceFileResponse(
            string bundleId,

            Outputs.FileReferenceResponse content,

            string devicePath)
        {
            BundleId = bundleId;
            Content = content;
            DevicePath = devicePath;
        }
    }
}