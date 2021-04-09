// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GcpNative.AppEngine.V1Alpha.Outputs
{

    [OutputType]
    public sealed class SslSettingsResponse
    {
        /// <summary>
        /// ID of the AuthorizedCertificate resource configuring SSL for the application. Clearing this field will remove SSL support.By default, a managed certificate is automatically created for every domain mapping. To omit SSL support or to configure SSL manually, specify no_managed_certificate on a CREATE or UPDATE request. You must be authorized to administer the AuthorizedCertificate resource to manually map it to a DomainMapping resource. Example: 12345.
        /// </summary>
        public readonly string CertificateId;
        /// <summary>
        /// Whether the mapped certificate is an App Engine managed certificate. Managed certificates are created by default with a domain mapping. To opt out, specify no_managed_certificate on a CREATE or UPDATE request.@OutputOnly
        /// </summary>
        public readonly bool IsManagedCertificate;

        [OutputConstructor]
        private SslSettingsResponse(
            string certificateId,

            bool isManagedCertificate)
        {
            CertificateId = certificateId;
            IsManagedCertificate = isManagedCertificate;
        }
    }
}