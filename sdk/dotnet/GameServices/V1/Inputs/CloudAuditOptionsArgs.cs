// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GameServices.V1.Inputs
{

    /// <summary>
    /// Write a Cloud Audit log
    /// </summary>
    public sealed class CloudAuditOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Information used by the Cloud Audit Logging pipeline.
        /// </summary>
        [Input("authorizationLoggingOptions")]
        public Input<Inputs.AuthorizationLoggingOptionsArgs>? AuthorizationLoggingOptions { get; set; }

        /// <summary>
        /// The log_name to populate in the Cloud Audit Record.
        /// </summary>
        [Input("logName")]
        public Input<Pulumi.GoogleNative.GameServices.V1.CloudAuditOptionsLogName>? LogName { get; set; }

        public CloudAuditOptionsArgs()
        {
        }
        public static new CloudAuditOptionsArgs Empty => new CloudAuditOptionsArgs();
    }
}
