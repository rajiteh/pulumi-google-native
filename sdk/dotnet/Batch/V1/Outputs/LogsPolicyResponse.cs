// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Batch.V1.Outputs
{

    /// <summary>
    /// LogsPolicy describes how outputs from a Job's Tasks (stdout/stderr) will be preserved.
    /// </summary>
    [OutputType]
    public sealed class LogsPolicyResponse
    {
        /// <summary>
        /// Where logs should be saved.
        /// </summary>
        public readonly string Destination;
        /// <summary>
        /// The path to which logs are saved when the destination = PATH. This can be a local file path on the VM, or under the mount point of a Persistent Disk or Filestore, or a Cloud Storage path.
        /// </summary>
        public readonly string LogsPath;

        [OutputConstructor]
        private LogsPolicyResponse(
            string destination,

            string logsPath)
        {
            Destination = destination;
            LogsPath = logsPath;
        }
    }
}
