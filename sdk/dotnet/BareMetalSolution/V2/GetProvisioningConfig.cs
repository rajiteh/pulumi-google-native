// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BareMetalSolution.V2
{
    public static class GetProvisioningConfig
    {
        /// <summary>
        /// Get ProvisioningConfig by name.
        /// </summary>
        public static Task<GetProvisioningConfigResult> InvokeAsync(GetProvisioningConfigArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetProvisioningConfigResult>("google-native:baremetalsolution/v2:getProvisioningConfig", args ?? new GetProvisioningConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Get ProvisioningConfig by name.
        /// </summary>
        public static Output<GetProvisioningConfigResult> Invoke(GetProvisioningConfigInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetProvisioningConfigResult>("google-native:baremetalsolution/v2:getProvisioningConfig", args ?? new GetProvisioningConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProvisioningConfigArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("provisioningConfigId", required: true)]
        public string ProvisioningConfigId { get; set; } = null!;

        public GetProvisioningConfigArgs()
        {
        }
        public static new GetProvisioningConfigArgs Empty => new GetProvisioningConfigArgs();
    }

    public sealed class GetProvisioningConfigInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("provisioningConfigId", required: true)]
        public Input<string> ProvisioningConfigId { get; set; } = null!;

        public GetProvisioningConfigInvokeArgs()
        {
        }
        public static new GetProvisioningConfigInvokeArgs Empty => new GetProvisioningConfigInvokeArgs();
    }


    [OutputType]
    public sealed class GetProvisioningConfigResult
    {
        /// <summary>
        /// URI to Cloud Console UI view of this provisioning config.
        /// </summary>
        public readonly string CloudConsoleUri;
        /// <summary>
        /// Optional. The user-defined identifier of the provisioning config.
        /// </summary>
        public readonly string CustomId;
        /// <summary>
        /// Email provided to send a confirmation with provisioning config to. Deprecated in favour of email field in request messages.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// A service account to enable customers to access instance credentials upon handover.
        /// </summary>
        public readonly string HandoverServiceAccount;
        /// <summary>
        /// Instances to be created.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceConfigResponse> Instances;
        /// <summary>
        /// Optional. Location name of this ProvisioningConfig. It is optional only for Intake UI transition period.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The system-generated name of the provisioning config. This follows the UUID format.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Networks to be created.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkConfigResponse> Networks;
        /// <summary>
        /// State of ProvisioningConfig.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Optional status messages associated with the FAILED state.
        /// </summary>
        public readonly string StatusMessage;
        /// <summary>
        /// A generated ticket id to track provisioning request.
        /// </summary>
        public readonly string TicketId;
        /// <summary>
        /// Last update timestamp.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// Volumes to be created.
        /// </summary>
        public readonly ImmutableArray<Outputs.VolumeConfigResponse> Volumes;
        /// <summary>
        /// If true, VPC SC is enabled for the cluster.
        /// </summary>
        public readonly bool VpcScEnabled;

        [OutputConstructor]
        private GetProvisioningConfigResult(
            string cloudConsoleUri,

            string customId,

            string email,

            string handoverServiceAccount,

            ImmutableArray<Outputs.InstanceConfigResponse> instances,

            string location,

            string name,

            ImmutableArray<Outputs.NetworkConfigResponse> networks,

            string state,

            string statusMessage,

            string ticketId,

            string updateTime,

            ImmutableArray<Outputs.VolumeConfigResponse> volumes,

            bool vpcScEnabled)
        {
            CloudConsoleUri = cloudConsoleUri;
            CustomId = customId;
            Email = email;
            HandoverServiceAccount = handoverServiceAccount;
            Instances = instances;
            Location = location;
            Name = name;
            Networks = networks;
            State = state;
            StatusMessage = statusMessage;
            TicketId = ticketId;
            UpdateTime = updateTime;
            Volumes = volumes;
            VpcScEnabled = vpcScEnabled;
        }
    }
}
