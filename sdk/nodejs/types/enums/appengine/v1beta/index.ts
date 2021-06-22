// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ApiConfigHandlerAuthFailAction = {
    /**
     * Not specified. AUTH_FAIL_ACTION_REDIRECT is assumed.
     */
    AuthFailActionUnspecified: "AUTH_FAIL_ACTION_UNSPECIFIED",
    /**
     * Redirects user to "accounts.google.com". The user is redirected back to the application URL after signing in or creating an account.
     */
    AuthFailActionRedirect: "AUTH_FAIL_ACTION_REDIRECT",
    /**
     * Rejects request with a 401 HTTP status code and an error message.
     */
    AuthFailActionUnauthorized: "AUTH_FAIL_ACTION_UNAUTHORIZED",
} as const;

/**
 * Action to take when users access resources that require authentication. Defaults to redirect.
 */
export type ApiConfigHandlerAuthFailAction = (typeof ApiConfigHandlerAuthFailAction)[keyof typeof ApiConfigHandlerAuthFailAction];

export const ApiConfigHandlerLogin = {
    /**
     * Not specified. LOGIN_OPTIONAL is assumed.
     */
    LoginUnspecified: "LOGIN_UNSPECIFIED",
    /**
     * Does not require that the user is signed in.
     */
    LoginOptional: "LOGIN_OPTIONAL",
    /**
     * If the user is not signed in, the auth_fail_action is taken. In addition, if the user is not an administrator for the application, they are given an error message regardless of auth_fail_action. If the user is an administrator, the handler proceeds.
     */
    LoginAdmin: "LOGIN_ADMIN",
    /**
     * If the user has signed in, the handler proceeds normally. Otherwise, the auth_fail_action is taken.
     */
    LoginRequired: "LOGIN_REQUIRED",
} as const;

/**
 * Level of login required to access this resource. Defaults to optional.
 */
export type ApiConfigHandlerLogin = (typeof ApiConfigHandlerLogin)[keyof typeof ApiConfigHandlerLogin];

export const ApiConfigHandlerSecurityLevel = {
    /**
     * Not specified.
     */
    SecureUnspecified: "SECURE_UNSPECIFIED",
    /**
     * Both HTTP and HTTPS requests with URLs that match the handler succeed without redirects. The application can examine the request to determine which protocol was used, and respond accordingly.
     */
    SecureDefault: "SECURE_DEFAULT",
    /**
     * Requests for a URL that match this handler that use HTTPS are automatically redirected to the HTTP equivalent URL.
     */
    SecureNever: "SECURE_NEVER",
    /**
     * Both HTTP and HTTPS requests with URLs that match the handler succeed without redirects. The application can examine the request to determine which protocol was used and respond accordingly.
     */
    SecureOptional: "SECURE_OPTIONAL",
    /**
     * Requests for a URL that match this handler that do not use HTTPS are automatically redirected to the HTTPS URL with the same path. Query parameters are reserved for the redirect.
     */
    SecureAlways: "SECURE_ALWAYS",
} as const;

/**
 * Security (HTTPS) enforcement for this URL.
 */
export type ApiConfigHandlerSecurityLevel = (typeof ApiConfigHandlerSecurityLevel)[keyof typeof ApiConfigHandlerSecurityLevel];

export const AppDatabaseType = {
    /**
     * Database type is unspecified.
     */
    DatabaseTypeUnspecified: "DATABASE_TYPE_UNSPECIFIED",
    /**
     * Cloud Datastore
     */
    CloudDatastore: "CLOUD_DATASTORE",
    /**
     * Cloud Firestore Native
     */
    CloudFirestore: "CLOUD_FIRESTORE",
    /**
     * Cloud Firestore in Datastore Mode
     */
    CloudDatastoreCompatibility: "CLOUD_DATASTORE_COMPATIBILITY",
} as const;

/**
 * The type of the Cloud Firestore or Cloud Datastore database associated with this application.
 */
export type AppDatabaseType = (typeof AppDatabaseType)[keyof typeof AppDatabaseType];

export const AppServingStatus = {
    /**
     * Serving status is unspecified.
     */
    Unspecified: "UNSPECIFIED",
    /**
     * Application is serving.
     */
    Serving: "SERVING",
    /**
     * Application has been disabled by the user.
     */
    UserDisabled: "USER_DISABLED",
    /**
     * Application has been disabled by the system.
     */
    SystemDisabled: "SYSTEM_DISABLED",
} as const;

/**
 * Serving status of this application.
 */
export type AppServingStatus = (typeof AppServingStatus)[keyof typeof AppServingStatus];

export const EndpointsApiServiceRolloutStrategy = {
    /**
     * Not specified. Defaults to FIXED.
     */
    UnspecifiedRolloutStrategy: "UNSPECIFIED_ROLLOUT_STRATEGY",
    /**
     * Endpoints service configuration ID will be fixed to the configuration ID specified by config_id.
     */
    Fixed: "FIXED",
    /**
     * Endpoints service configuration ID will be updated with each rollout.
     */
    Managed: "MANAGED",
} as const;

/**
 * Endpoints rollout strategy. If FIXED, config_id must be specified. If MANAGED, config_id must be omitted.
 */
export type EndpointsApiServiceRolloutStrategy = (typeof EndpointsApiServiceRolloutStrategy)[keyof typeof EndpointsApiServiceRolloutStrategy];

export const ErrorHandlerErrorCode = {
    /**
     * Not specified. ERROR_CODE_DEFAULT is assumed.
     */
    ErrorCodeUnspecified: "ERROR_CODE_UNSPECIFIED",
    /**
     * All other error types.
     */
    ErrorCodeDefault: "ERROR_CODE_DEFAULT",
    /**
     * Application has exceeded a resource quota.
     */
    ErrorCodeOverQuota: "ERROR_CODE_OVER_QUOTA",
    /**
     * Client blocked by the application's Denial of Service protection configuration.
     */
    ErrorCodeDosApiDenial: "ERROR_CODE_DOS_API_DENIAL",
    /**
     * Deadline reached before the application responds.
     */
    ErrorCodeTimeout: "ERROR_CODE_TIMEOUT",
} as const;

/**
 * Error condition this handler applies to.
 */
export type ErrorHandlerErrorCode = (typeof ErrorHandlerErrorCode)[keyof typeof ErrorHandlerErrorCode];

export const IngressRuleAction = {
    UnspecifiedAction: "UNSPECIFIED_ACTION",
    /**
     * Matching requests are allowed.
     */
    Allow: "ALLOW",
    /**
     * Matching requests are denied.
     */
    Deny: "DENY",
} as const;

/**
 * The action to take on matched requests.
 */
export type IngressRuleAction = (typeof IngressRuleAction)[keyof typeof IngressRuleAction];

export const ManagedCertificateStatus = {
    ManagementStatusUnspecified: "MANAGEMENT_STATUS_UNSPECIFIED",
    /**
     * Certificate was successfully obtained and inserted into the serving system.
     */
    Ok: "OK",
    /**
     * Certificate is under active attempts to acquire or renew.
     */
    Pending: "PENDING",
    /**
     * Most recent renewal failed due to an invalid DNS setup and will be retried. Renewal attempts will continue to fail until the certificate domain's DNS configuration is fixed. The last successfully provisioned certificate may still be serving.
     */
    FailedRetryingNotVisible: "FAILED_RETRYING_NOT_VISIBLE",
    /**
     * All renewal attempts have been exhausted, likely due to an invalid DNS setup.
     */
    FailedPermanent: "FAILED_PERMANENT",
    /**
     * Most recent renewal failed due to an explicit CAA record that does not include one of the in-use CAs (Google CA and Let's Encrypt). Renewals will continue to fail until the CAA is reconfigured. The last successfully provisioned certificate may still be serving.
     */
    FailedRetryingCaaForbidden: "FAILED_RETRYING_CAA_FORBIDDEN",
    /**
     * Most recent renewal failed due to a CAA retrieval failure. This means that the domain's DNS provider does not properly handle CAA records, failing requests for CAA records when no CAA records are defined. Renewals will continue to fail until the DNS provider is changed or a CAA record is added for the given domain. The last successfully provisioned certificate may still be serving.
     */
    FailedRetryingCaaChecking: "FAILED_RETRYING_CAA_CHECKING",
} as const;

/**
 * Status of certificate management. Refers to the most recent certificate acquisition or renewal attempt.@OutputOnly
 */
export type ManagedCertificateStatus = (typeof ManagedCertificateStatus)[keyof typeof ManagedCertificateStatus];

export const ResourceRecordType = {
    /**
     * An A resource record. Data is an IPv4 address.
     */
    A: "A",
    /**
     * An AAAA resource record. Data is an IPv6 address.
     */
    Aaaa: "AAAA",
    /**
     * A CNAME resource record. Data is a domain name to be aliased.
     */
    Cname: "CNAME",
} as const;

/**
 * Resource record type. Example: AAAA.
 */
export type ResourceRecordType = (typeof ResourceRecordType)[keyof typeof ResourceRecordType];

export const SslSettingsSslManagementType = {
    /**
     * SSL support for this domain is configured automatically. The mapped SSL certificate will be automatically renewed.
     */
    Automatic: "AUTOMATIC",
    /**
     * SSL support for this domain is configured manually by the user. Either the domain has no SSL support or a user-obtained SSL certificate has been explictly mapped to this domain.
     */
    Manual: "MANUAL",
} as const;

/**
 * SSL management type for this domain. If AUTOMATIC, a managed certificate is automatically provisioned. If MANUAL, certificate_id must be manually specified in order to configure SSL for this domain.
 */
export type SslSettingsSslManagementType = (typeof SslSettingsSslManagementType)[keyof typeof SslSettingsSslManagementType];

export const UrlMapAuthFailAction = {
    /**
     * Not specified. AUTH_FAIL_ACTION_REDIRECT is assumed.
     */
    AuthFailActionUnspecified: "AUTH_FAIL_ACTION_UNSPECIFIED",
    /**
     * Redirects user to "accounts.google.com". The user is redirected back to the application URL after signing in or creating an account.
     */
    AuthFailActionRedirect: "AUTH_FAIL_ACTION_REDIRECT",
    /**
     * Rejects request with a 401 HTTP status code and an error message.
     */
    AuthFailActionUnauthorized: "AUTH_FAIL_ACTION_UNAUTHORIZED",
} as const;

/**
 * Action to take when users access resources that require authentication. Defaults to redirect.
 */
export type UrlMapAuthFailAction = (typeof UrlMapAuthFailAction)[keyof typeof UrlMapAuthFailAction];

export const UrlMapLogin = {
    /**
     * Not specified. LOGIN_OPTIONAL is assumed.
     */
    LoginUnspecified: "LOGIN_UNSPECIFIED",
    /**
     * Does not require that the user is signed in.
     */
    LoginOptional: "LOGIN_OPTIONAL",
    /**
     * If the user is not signed in, the auth_fail_action is taken. In addition, if the user is not an administrator for the application, they are given an error message regardless of auth_fail_action. If the user is an administrator, the handler proceeds.
     */
    LoginAdmin: "LOGIN_ADMIN",
    /**
     * If the user has signed in, the handler proceeds normally. Otherwise, the auth_fail_action is taken.
     */
    LoginRequired: "LOGIN_REQUIRED",
} as const;

/**
 * Level of login required to access this resource. Not supported for Node.js in the App Engine standard environment.
 */
export type UrlMapLogin = (typeof UrlMapLogin)[keyof typeof UrlMapLogin];

export const UrlMapRedirectHttpResponseCode = {
    /**
     * Not specified. 302 is assumed.
     */
    RedirectHttpResponseCodeUnspecified: "REDIRECT_HTTP_RESPONSE_CODE_UNSPECIFIED",
    /**
     * 301 Moved Permanently code.
     */
    RedirectHttpResponseCode301: "REDIRECT_HTTP_RESPONSE_CODE_301",
    /**
     * 302 Moved Temporarily code.
     */
    RedirectHttpResponseCode302: "REDIRECT_HTTP_RESPONSE_CODE_302",
    /**
     * 303 See Other code.
     */
    RedirectHttpResponseCode303: "REDIRECT_HTTP_RESPONSE_CODE_303",
    /**
     * 307 Temporary Redirect code.
     */
    RedirectHttpResponseCode307: "REDIRECT_HTTP_RESPONSE_CODE_307",
} as const;

/**
 * 30x code to use when performing redirects for the secure field. Defaults to 302.
 */
export type UrlMapRedirectHttpResponseCode = (typeof UrlMapRedirectHttpResponseCode)[keyof typeof UrlMapRedirectHttpResponseCode];

export const UrlMapSecurityLevel = {
    /**
     * Not specified.
     */
    SecureUnspecified: "SECURE_UNSPECIFIED",
    /**
     * Both HTTP and HTTPS requests with URLs that match the handler succeed without redirects. The application can examine the request to determine which protocol was used, and respond accordingly.
     */
    SecureDefault: "SECURE_DEFAULT",
    /**
     * Requests for a URL that match this handler that use HTTPS are automatically redirected to the HTTP equivalent URL.
     */
    SecureNever: "SECURE_NEVER",
    /**
     * Both HTTP and HTTPS requests with URLs that match the handler succeed without redirects. The application can examine the request to determine which protocol was used and respond accordingly.
     */
    SecureOptional: "SECURE_OPTIONAL",
    /**
     * Requests for a URL that match this handler that do not use HTTPS are automatically redirected to the HTTPS URL with the same path. Query parameters are reserved for the redirect.
     */
    SecureAlways: "SECURE_ALWAYS",
} as const;

/**
 * Security (HTTPS) enforcement for this URL.
 */
export type UrlMapSecurityLevel = (typeof UrlMapSecurityLevel)[keyof typeof UrlMapSecurityLevel];

export const VersionInboundServicesItem = {
    /**
     * Not specified.
     */
    InboundServiceUnspecified: "INBOUND_SERVICE_UNSPECIFIED",
    /**
     * Allows an application to receive mail.
     */
    InboundServiceMail: "INBOUND_SERVICE_MAIL",
    /**
     * Allows an application to receive email-bound notifications.
     */
    InboundServiceMailBounce: "INBOUND_SERVICE_MAIL_BOUNCE",
    /**
     * Allows an application to receive error stanzas.
     */
    InboundServiceXmppError: "INBOUND_SERVICE_XMPP_ERROR",
    /**
     * Allows an application to receive instant messages.
     */
    InboundServiceXmppMessage: "INBOUND_SERVICE_XMPP_MESSAGE",
    /**
     * Allows an application to receive user subscription POSTs.
     */
    InboundServiceXmppSubscribe: "INBOUND_SERVICE_XMPP_SUBSCRIBE",
    /**
     * Allows an application to receive a user's chat presence.
     */
    InboundServiceXmppPresence: "INBOUND_SERVICE_XMPP_PRESENCE",
    /**
     * Registers an application for notifications when a client connects or disconnects from a channel.
     */
    InboundServiceChannelPresence: "INBOUND_SERVICE_CHANNEL_PRESENCE",
    /**
     * Enables warmup requests.
     */
    InboundServiceWarmup: "INBOUND_SERVICE_WARMUP",
} as const;

export type VersionInboundServicesItem = (typeof VersionInboundServicesItem)[keyof typeof VersionInboundServicesItem];

export const VersionServingStatus = {
    /**
     * Not specified.
     */
    ServingStatusUnspecified: "SERVING_STATUS_UNSPECIFIED",
    /**
     * Currently serving. Instances are created according to the scaling settings of the version.
     */
    Serving: "SERVING",
    /**
     * Disabled. No instances will be created and the scaling settings are ignored until the state of the version changes to SERVING.
     */
    Stopped: "STOPPED",
} as const;

/**
 * Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.SERVING_STATUS_UNSPECIFIED is an invalid value. Defaults to SERVING.
 */
export type VersionServingStatus = (typeof VersionServingStatus)[keyof typeof VersionServingStatus];