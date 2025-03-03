// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    /// <summary>
    /// Creates an API product in an organization. You create API products after you have proxied backend services using API proxies. An API product is a collection of API resources combined with quota settings and metadata that you can use to deliver customized and productized API bundles to your developer community. This metadata can include: - Scope - Environments - API proxies - Extensible profile API products enable you repackage APIs on the fly, without having to do any additional coding or configuration. Apigee recommends that you start with a simple API product including only required elements. You then provision credentials to apps to enable them to start testing your APIs. After you have authentication and authorization working against a simple API product, you can iterate to create finer-grained API products, defining different sets of API resources for each API product. **WARNING:** - If you don't specify an API proxy in the request body, *any* app associated with the product can make calls to *any* API in your entire organization. - If you don't specify an environment in the request body, the product allows access to all environments. For more information, see What is an API product?
    /// </summary>
    [GoogleNativeResourceType("google-native:apigee/v1:ApiProduct")]
    public partial class ApiProduct : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Comma-separated list of API resources to be bundled in the API product. By default, the resource paths are mapped from the `proxy.pathsuffix` variable. The proxy path suffix is defined as the URI fragment following the ProxyEndpoint base path. For example, if the `apiResources` element is defined to be `/forecastrss` and the base path defined for the API proxy is `/weather`, then only requests to `/weather/forecastrss` are permitted by the API product. You can select a specific path, or you can select all subpaths with the following wildcard: - `/**`: Indicates that all sub-URIs are included. - `/*` : Indicates that only URIs one level down are included. By default, / supports the same resources as /** as well as the base path defined by the API proxy. For example, if the base path of the API proxy is `/v1/weatherapikey`, then the API product supports requests to `/v1/weatherapikey` and to any sub-URIs, such as `/v1/weatherapikey/forecastrss`, `/v1/weatherapikey/region/CA`, and so on. For more information, see Managing API products.
        /// </summary>
        [Output("apiResources")]
        public Output<ImmutableArray<string>> ApiResources { get; private set; } = null!;

        /// <summary>
        /// Flag that specifies how API keys are approved to access the APIs defined by the API product. If set to `manual`, the consumer key is generated and returned in "pending" state. In this case, the API keys won't work until they have been explicitly approved. If set to `auto`, the consumer key is generated and returned in "approved" state and can be used immediately. **Note:** Typically, `auto` is used to provide access to free or trial API products that provide limited quota or capabilities.
        /// </summary>
        [Output("approvalType")]
        public Output<string> ApprovalType { get; private set; } = null!;

        /// <summary>
        /// Array of attributes that may be used to extend the default API product profile with customer-specific metadata. You can specify a maximum of 18 attributes. Use this property to specify the access level of the API product as either `public`, `private`, or `internal`. Only products marked `public` are available to developers in the Apigee developer portal. For example, you can set a product to `internal` while it is in development and then change access to `public` when it is ready to release on the portal. API products marked as `private` do not appear on the portal, but can be accessed by external developers.
        /// </summary>
        [Output("attributes")]
        public Output<ImmutableArray<Outputs.GoogleCloudApigeeV1AttributeResponse>> Attributes { get; private set; } = null!;

        /// <summary>
        /// Response only. Creation time of this environment as milliseconds since epoch.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Description of the API product. Include key information about the API product that is not captured by other fields.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Name displayed in the UI or developer portal to developers registering for API access.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Comma-separated list of environment names to which the API product is bound. Requests to environments that are not listed are rejected. By specifying one or more environments, you can bind the resources listed in the API product to a specific environment, preventing developers from accessing those resources through API proxies deployed in another environment. This setting is used, for example, to prevent resources associated with API proxies in `prod` from being accessed by API proxies deployed in `test`.
        /// </summary>
        [Output("environments")]
        public Output<ImmutableArray<string>> Environments { get; private set; } = null!;

        /// <summary>
        /// Configuration used to group Apigee proxies or remote services with graphQL operation name, graphQL operation type and quotas. This grouping allows us to precisely set quota for a particular combination of graphQL name and operation type for a particular proxy request. If graphQL name is not set, this would imply quota will be applied on all graphQL requests matching the operation type.
        /// </summary>
        [Output("graphqlOperationGroup")]
        public Output<Outputs.GoogleCloudApigeeV1GraphQLOperationGroupResponse> GraphqlOperationGroup { get; private set; } = null!;

        /// <summary>
        /// Response only. Modified time of this environment as milliseconds since epoch.
        /// </summary>
        [Output("lastModifiedAt")]
        public Output<string> LastModifiedAt { get; private set; } = null!;

        /// <summary>
        /// Internal name of the API product. Characters you can use in the name are restricted to: `A-Z0-9._\-$ %`. **Note:** The internal name cannot be edited when updating the API product.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Configuration used to group Apigee proxies or remote services with resources, method types, and quotas. The resource refers to the resource URI (excluding the base path). With this grouping, the API product creator is able to fine-tune and give precise control over which REST methods have access to specific resources and how many calls can be made (using the `quota` setting). **Note:** The `api_resources` setting cannot be specified for both the API product and operation group; otherwise the call will fail.
        /// </summary>
        [Output("operationGroup")]
        public Output<Outputs.GoogleCloudApigeeV1OperationGroupResponse> OperationGroup { get; private set; } = null!;

        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// Comma-separated list of API proxy names to which this API product is bound. By specifying API proxies, you can associate resources in the API product with specific API proxies, preventing developers from accessing those resources through other API proxies. Apigee rejects requests to API proxies that are not listed. **Note:** The API proxy names must already exist in the specified environment as they will be validated upon creation.
        /// </summary>
        [Output("proxies")]
        public Output<ImmutableArray<string>> Proxies { get; private set; } = null!;

        /// <summary>
        /// Number of request messages permitted per app by this API product for the specified `quotaInterval` and `quotaTimeUnit`. For example, a `quota` of 50, for a `quotaInterval` of 12 and a `quotaTimeUnit` of hours means 50 requests are allowed every 12 hours.
        /// </summary>
        [Output("quota")]
        public Output<string> Quota { get; private set; } = null!;

        /// <summary>
        /// Scope of the quota decides how the quota counter gets applied and evaluate for quota violation. If the Scope is set as PROXY, then all the operations defined for the APIproduct that are associated with the same proxy will share the same quota counter set at the APIproduct level, making it a global counter at a proxy level. If the Scope is set as OPERATION, then each operations get the counter set at the API product dedicated, making it a local counter. Note that, the QuotaCounterScope applies only when an operation does not have dedicated quota set for itself.
        /// </summary>
        [Output("quotaCounterScope")]
        public Output<string> QuotaCounterScope { get; private set; } = null!;

        /// <summary>
        /// Time interval over which the number of request messages is calculated.
        /// </summary>
        [Output("quotaInterval")]
        public Output<string> QuotaInterval { get; private set; } = null!;

        /// <summary>
        /// Time unit defined for the `quotaInterval`. Valid values include `minute`, `hour`, `day`, or `month`.
        /// </summary>
        [Output("quotaTimeUnit")]
        public Output<string> QuotaTimeUnit { get; private set; } = null!;

        /// <summary>
        /// Comma-separated list of OAuth scopes that are validated at runtime. Apigee validates that the scopes in any access token presented match the scopes defined in the OAuth policy associated with the API product.
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;


        /// <summary>
        /// Create a ApiProduct resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApiProduct(string name, ApiProductArgs args, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:ApiProduct", name, args ?? new ApiProductArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApiProduct(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:ApiProduct", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "organizationId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ApiProduct resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApiProduct Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ApiProduct(name, id, options);
        }
    }

    public sealed class ApiProductArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiResources")]
        private InputList<string>? _apiResources;

        /// <summary>
        /// Comma-separated list of API resources to be bundled in the API product. By default, the resource paths are mapped from the `proxy.pathsuffix` variable. The proxy path suffix is defined as the URI fragment following the ProxyEndpoint base path. For example, if the `apiResources` element is defined to be `/forecastrss` and the base path defined for the API proxy is `/weather`, then only requests to `/weather/forecastrss` are permitted by the API product. You can select a specific path, or you can select all subpaths with the following wildcard: - `/**`: Indicates that all sub-URIs are included. - `/*` : Indicates that only URIs one level down are included. By default, / supports the same resources as /** as well as the base path defined by the API proxy. For example, if the base path of the API proxy is `/v1/weatherapikey`, then the API product supports requests to `/v1/weatherapikey` and to any sub-URIs, such as `/v1/weatherapikey/forecastrss`, `/v1/weatherapikey/region/CA`, and so on. For more information, see Managing API products.
        /// </summary>
        public InputList<string> ApiResources
        {
            get => _apiResources ?? (_apiResources = new InputList<string>());
            set => _apiResources = value;
        }

        /// <summary>
        /// Flag that specifies how API keys are approved to access the APIs defined by the API product. If set to `manual`, the consumer key is generated and returned in "pending" state. In this case, the API keys won't work until they have been explicitly approved. If set to `auto`, the consumer key is generated and returned in "approved" state and can be used immediately. **Note:** Typically, `auto` is used to provide access to free or trial API products that provide limited quota or capabilities.
        /// </summary>
        [Input("approvalType")]
        public Input<string>? ApprovalType { get; set; }

        [Input("attributes")]
        private InputList<Inputs.GoogleCloudApigeeV1AttributeArgs>? _attributes;

        /// <summary>
        /// Array of attributes that may be used to extend the default API product profile with customer-specific metadata. You can specify a maximum of 18 attributes. Use this property to specify the access level of the API product as either `public`, `private`, or `internal`. Only products marked `public` are available to developers in the Apigee developer portal. For example, you can set a product to `internal` while it is in development and then change access to `public` when it is ready to release on the portal. API products marked as `private` do not appear on the portal, but can be accessed by external developers.
        /// </summary>
        public InputList<Inputs.GoogleCloudApigeeV1AttributeArgs> Attributes
        {
            get => _attributes ?? (_attributes = new InputList<Inputs.GoogleCloudApigeeV1AttributeArgs>());
            set => _attributes = value;
        }

        /// <summary>
        /// Response only. Creation time of this environment as milliseconds since epoch.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Description of the API product. Include key information about the API product that is not captured by other fields.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name displayed in the UI or developer portal to developers registering for API access.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("environments")]
        private InputList<string>? _environments;

        /// <summary>
        /// Comma-separated list of environment names to which the API product is bound. Requests to environments that are not listed are rejected. By specifying one or more environments, you can bind the resources listed in the API product to a specific environment, preventing developers from accessing those resources through API proxies deployed in another environment. This setting is used, for example, to prevent resources associated with API proxies in `prod` from being accessed by API proxies deployed in `test`.
        /// </summary>
        public InputList<string> Environments
        {
            get => _environments ?? (_environments = new InputList<string>());
            set => _environments = value;
        }

        /// <summary>
        /// Configuration used to group Apigee proxies or remote services with graphQL operation name, graphQL operation type and quotas. This grouping allows us to precisely set quota for a particular combination of graphQL name and operation type for a particular proxy request. If graphQL name is not set, this would imply quota will be applied on all graphQL requests matching the operation type.
        /// </summary>
        [Input("graphqlOperationGroup")]
        public Input<Inputs.GoogleCloudApigeeV1GraphQLOperationGroupArgs>? GraphqlOperationGroup { get; set; }

        /// <summary>
        /// Response only. Modified time of this environment as milliseconds since epoch.
        /// </summary>
        [Input("lastModifiedAt")]
        public Input<string>? LastModifiedAt { get; set; }

        /// <summary>
        /// Internal name of the API product. Characters you can use in the name are restricted to: `A-Z0-9._\-$ %`. **Note:** The internal name cannot be edited when updating the API product.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configuration used to group Apigee proxies or remote services with resources, method types, and quotas. The resource refers to the resource URI (excluding the base path). With this grouping, the API product creator is able to fine-tune and give precise control over which REST methods have access to specific resources and how many calls can be made (using the `quota` setting). **Note:** The `api_resources` setting cannot be specified for both the API product and operation group; otherwise the call will fail.
        /// </summary>
        [Input("operationGroup")]
        public Input<Inputs.GoogleCloudApigeeV1OperationGroupArgs>? OperationGroup { get; set; }

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        [Input("proxies")]
        private InputList<string>? _proxies;

        /// <summary>
        /// Comma-separated list of API proxy names to which this API product is bound. By specifying API proxies, you can associate resources in the API product with specific API proxies, preventing developers from accessing those resources through other API proxies. Apigee rejects requests to API proxies that are not listed. **Note:** The API proxy names must already exist in the specified environment as they will be validated upon creation.
        /// </summary>
        public InputList<string> Proxies
        {
            get => _proxies ?? (_proxies = new InputList<string>());
            set => _proxies = value;
        }

        /// <summary>
        /// Number of request messages permitted per app by this API product for the specified `quotaInterval` and `quotaTimeUnit`. For example, a `quota` of 50, for a `quotaInterval` of 12 and a `quotaTimeUnit` of hours means 50 requests are allowed every 12 hours.
        /// </summary>
        [Input("quota")]
        public Input<string>? Quota { get; set; }

        /// <summary>
        /// Scope of the quota decides how the quota counter gets applied and evaluate for quota violation. If the Scope is set as PROXY, then all the operations defined for the APIproduct that are associated with the same proxy will share the same quota counter set at the APIproduct level, making it a global counter at a proxy level. If the Scope is set as OPERATION, then each operations get the counter set at the API product dedicated, making it a local counter. Note that, the QuotaCounterScope applies only when an operation does not have dedicated quota set for itself.
        /// </summary>
        [Input("quotaCounterScope")]
        public Input<Pulumi.GoogleNative.Apigee.V1.ApiProductQuotaCounterScope>? QuotaCounterScope { get; set; }

        /// <summary>
        /// Time interval over which the number of request messages is calculated.
        /// </summary>
        [Input("quotaInterval")]
        public Input<string>? QuotaInterval { get; set; }

        /// <summary>
        /// Time unit defined for the `quotaInterval`. Valid values include `minute`, `hour`, `day`, or `month`.
        /// </summary>
        [Input("quotaTimeUnit")]
        public Input<string>? QuotaTimeUnit { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// Comma-separated list of OAuth scopes that are validated at runtime. Apigee validates that the scopes in any access token presented match the scopes defined in the OAuth policy associated with the API product.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        public ApiProductArgs()
        {
        }
        public static new ApiProductArgs Empty => new ApiProductArgs();
    }
}
