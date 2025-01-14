// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Outputs
{

    /// <summary>
    /// Configuration detail of coordinate, it used for UI
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudIntegrationsV1alphaCoordinateResponse
    {
        /// <summary>
        /// X axis of the coordinate
        /// </summary>
        public readonly int X;
        /// <summary>
        /// Y axis of the coordinate
        /// </summary>
        public readonly int Y;

        [OutputConstructor]
        private GoogleCloudIntegrationsV1alphaCoordinateResponse(
            int x,

            int y)
        {
            X = x;
            Y = y;
        }
    }
}
