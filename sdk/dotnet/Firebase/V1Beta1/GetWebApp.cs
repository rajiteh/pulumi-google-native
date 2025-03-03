// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Firebase.V1Beta1
{
    public static class GetWebApp
    {
        /// <summary>
        /// Gets the specified WebApp.
        /// </summary>
        public static Task<GetWebAppResult> InvokeAsync(GetWebAppArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWebAppResult>("google-native:firebase/v1beta1:getWebApp", args ?? new GetWebAppArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the specified WebApp.
        /// </summary>
        public static Output<GetWebAppResult> Invoke(GetWebAppInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWebAppResult>("google-native:firebase/v1beta1:getWebApp", args ?? new GetWebAppInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWebAppArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("webAppId", required: true)]
        public string WebAppId { get; set; } = null!;

        public GetWebAppArgs()
        {
        }
        public static new GetWebAppArgs Empty => new GetWebAppArgs();
    }

    public sealed class GetWebAppInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("webAppId", required: true)]
        public Input<string> WebAppId { get; set; } = null!;

        public GetWebAppInvokeArgs()
        {
        }
        public static new GetWebAppInvokeArgs Empty => new GetWebAppInvokeArgs();
    }


    [OutputType]
    public sealed class GetWebAppResult
    {
        /// <summary>
        /// The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the `WebApp`. Be aware that this value is the UID of the API key, _not_ the [`keyString`](https://cloud.google.com/api-keys/docs/reference/rest/v2/projects.locations.keys#Key.FIELDS.key_string) of the API key. The `keyString` is the value that can be found in the App's [configuration artifact](../../rest/v1beta1/projects.webApps/getConfig). If `api_key_id` is not set in requests to [`webApps.Create`](../../rest/v1beta1/projects.webApps/create), then Firebase automatically associates an `api_key_id` with the `WebApp`. This auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned. In patch requests, `api_key_id` cannot be set to an empty value, and the new UID must have no restrictions or only have restrictions that are valid for the associated `WebApp`. We recommend using the [Google Cloud Console](https://console.cloud.google.com/apis/credentials) to manage API keys.
        /// </summary>
        public readonly string ApiKeyId;
        /// <summary>
        /// Immutable. The globally unique, Firebase-assigned identifier for the `WebApp`. This identifier should be treated as an opaque token, as the data format is not specified.
        /// </summary>
        public readonly string AppId;
        /// <summary>
        /// The URLs where the `WebApp` is hosted.
        /// </summary>
        public readonly ImmutableArray<string> AppUrls;
        /// <summary>
        /// The user-assigned display name for the `WebApp`.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// This checksum is computed by the server based on the value of other fields, and it may be sent with update requests to ensure the client has an up-to-date value before proceeding. Learn more about `etag` in Google's [AIP-154 standard](https://google.aip.dev/154#declarative-friendly-resources). This etag is strongly validated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Timestamp of when the App will be considered expired and cannot be undeleted. This value is only provided if the App is in the `DELETED` state.
        /// </summary>
        public readonly string ExpireTime;
        /// <summary>
        /// The resource name of the WebApp, in the format: projects/PROJECT_IDENTIFIER /webApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.webApps#WebApp.FIELDS.app_id)).
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Immutable. A user-assigned unique identifier of the parent FirebaseProject for the `WebApp`.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// The lifecycle state of the App.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Immutable. A unique, Firebase-assigned identifier for the `WebApp`. This identifier is only used to populate the `namespace` value for the `WebApp`. For most use cases, use `appId` to identify or reference the App. The `webId` value is only unique within a `FirebaseProject` and its associated Apps.
        /// </summary>
        public readonly string WebId;

        [OutputConstructor]
        private GetWebAppResult(
            string apiKeyId,

            string appId,

            ImmutableArray<string> appUrls,

            string displayName,

            string etag,

            string expireTime,

            string name,

            string project,

            string state,

            string webId)
        {
            ApiKeyId = apiKeyId;
            AppId = appId;
            AppUrls = appUrls;
            DisplayName = displayName;
            Etag = etag;
            ExpireTime = expireTime;
            Name = name;
            Project = project;
            State = state;
            WebId = webId;
        }
    }
}
