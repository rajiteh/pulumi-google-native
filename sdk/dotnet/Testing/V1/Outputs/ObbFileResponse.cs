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
    public sealed class ObbFileResponse
    {
        /// <summary>
        /// Required. Opaque Binary Blob (OBB) file(s) to install on the device.
        /// </summary>
        public readonly Outputs.FileReferenceResponse Obb;
        /// <summary>
        /// Required. OBB file name which must conform to the format as specified by Android e.g. [main|patch].0300110.com.example.android.obb which will be installed into \/Android/obb/\/ on the device.
        /// </summary>
        public readonly string ObbFileName;

        [OutputConstructor]
        private ObbFileResponse(
            Outputs.FileReferenceResponse obb,

            string obbFileName)
        {
            Obb = obb;
            ObbFileName = obbFileName;
        }
    }
}