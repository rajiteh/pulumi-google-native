// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1
{
    /// <summary>
    /// Creates an App Engine application for a Google Cloud Platform project. Required fields: id - The ID of the target Cloud Platform project. location - The region (https://cloud.google.com/appengine/docs/locations) where you want the App Engine application located.For more information about App Engine applications, see Managing Projects, Applications, and Billing (https://cloud.google.com/appengine/docs/standard/python/console/).
    /// Auto-naming is currently not supported for this resource.
    /// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
    /// on Google Cloud even though it will be deleted from Pulumi state.
    /// </summary>
    [GoogleNativeResourceType("google-native:appengine/v1:Application")]
    public partial class Application : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
        /// </summary>
        [Output("authDomain")]
        public Output<string> AuthDomain { get; private set; } = null!;

        /// <summary>
        /// Google Cloud Storage bucket that can be used for storing files associated with this application. This bucket is associated with the application and can be used by the gcloud deployment commands.
        /// </summary>
        [Output("codeBucket")]
        public Output<string> CodeBucket { get; private set; } = null!;

        /// <summary>
        /// The type of the Cloud Firestore or Cloud Datastore database associated with this application.
        /// </summary>
        [Output("databaseType")]
        public Output<string> DatabaseType { get; private set; } = null!;

        /// <summary>
        /// Google Cloud Storage bucket that can be used by this application to store content.
        /// </summary>
        [Output("defaultBucket")]
        public Output<string> DefaultBucket { get; private set; } = null!;

        /// <summary>
        /// Cookie expiration policy for this application.
        /// </summary>
        [Output("defaultCookieExpiration")]
        public Output<string> DefaultCookieExpiration { get; private set; } = null!;

        /// <summary>
        /// Hostname used to reach this application, as resolved by App Engine.
        /// </summary>
        [Output("defaultHostname")]
        public Output<string> DefaultHostname { get; private set; } = null!;

        /// <summary>
        /// HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
        /// </summary>
        [Output("dispatchRules")]
        public Output<ImmutableArray<Outputs.UrlDispatchRuleResponse>> DispatchRules { get; private set; } = null!;

        /// <summary>
        /// The feature specific settings to be used in the application.
        /// </summary>
        [Output("featureSettings")]
        public Output<Outputs.FeatureSettingsResponse> FeatureSettings { get; private set; } = null!;

        /// <summary>
        /// The Google Container Registry domain used for storing managed build docker images for this application.
        /// </summary>
        [Output("gcrDomain")]
        public Output<string> GcrDomain { get; private set; } = null!;

        [Output("iap")]
        public Output<Outputs.IdentityAwareProxyResponse> Iap { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Full path to the Application resource in the API. Example: apps/myapp.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The service account associated with the application. This is the app-level default identity. If no identity provided during create version, Admin API will fallback to this one.
        /// </summary>
        [Output("serviceAccount")]
        public Output<string> ServiceAccount { get; private set; } = null!;

        /// <summary>
        /// Serving status of this application.
        /// </summary>
        [Output("servingStatus")]
        public Output<string> ServingStatus { get; private set; } = null!;


        /// <summary>
        /// Create a Application resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Application(string name, ApplicationArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:appengine/v1:Application", name, args ?? new ApplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Application(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:appengine/v1:Application", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "location",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Application resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Application Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Application(name, id, options);
        }
    }

    public sealed class ApplicationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
        /// </summary>
        [Input("authDomain")]
        public Input<string>? AuthDomain { get; set; }

        /// <summary>
        /// The type of the Cloud Firestore or Cloud Datastore database associated with this application.
        /// </summary>
        [Input("databaseType")]
        public Input<Pulumi.GoogleNative.AppEngine.V1.ApplicationDatabaseType>? DatabaseType { get; set; }

        /// <summary>
        /// Cookie expiration policy for this application.
        /// </summary>
        [Input("defaultCookieExpiration")]
        public Input<string>? DefaultCookieExpiration { get; set; }

        [Input("dispatchRules")]
        private InputList<Inputs.UrlDispatchRuleArgs>? _dispatchRules;

        /// <summary>
        /// HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
        /// </summary>
        public InputList<Inputs.UrlDispatchRuleArgs> DispatchRules
        {
            get => _dispatchRules ?? (_dispatchRules = new InputList<Inputs.UrlDispatchRuleArgs>());
            set => _dispatchRules = value;
        }

        /// <summary>
        /// The feature specific settings to be used in the application.
        /// </summary>
        [Input("featureSettings")]
        public Input<Inputs.FeatureSettingsArgs>? FeatureSettings { get; set; }

        [Input("iap")]
        public Input<Inputs.IdentityAwareProxyArgs>? Iap { get; set; }

        /// <summary>
        /// Identifier of the Application resource. This identifier is equivalent to the project ID of the Google Cloud Platform project where you want to deploy your application. Example: myapp.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Location from which this application runs. Application instances run out of the data centers in the specified location, which is also where all of the application's end user content is stored.Defaults to us-central.View the list of supported locations (https://cloud.google.com/appengine/docs/locations).
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The service account associated with the application. This is the app-level default identity. If no identity provided during create version, Admin API will fallback to this one.
        /// </summary>
        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        /// <summary>
        /// Serving status of this application.
        /// </summary>
        [Input("servingStatus")]
        public Input<Pulumi.GoogleNative.AppEngine.V1.ApplicationServingStatus>? ServingStatus { get; set; }

        public ApplicationArgs()
        {
        }
        public static new ApplicationArgs Empty => new ApplicationArgs();
    }
}
