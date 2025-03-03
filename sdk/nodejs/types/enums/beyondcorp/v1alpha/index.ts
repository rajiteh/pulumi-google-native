// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AppConnectionType = {
    /**
     * Default value. This value is unused.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * TCP Proxy based BeyondCorp AppConnection. API will default to this if unset.
     */
    TcpProxy: "TCP_PROXY",
} as const;

/**
 * Required. The type of network connectivity used by the AppConnection.
 */
export type AppConnectionType = (typeof AppConnectionType)[keyof typeof AppConnectionType];

export const AppGatewayHostType = {
    /**
     * Default value. This value is unused.
     */
    HostTypeUnspecified: "HOST_TYPE_UNSPECIFIED",
    /**
     * AppGateway hosted in a GCP regional managed instance group.
     */
    GcpRegionalMig: "GCP_REGIONAL_MIG",
} as const;

/**
 * Required. The type of hosting used by the AppGateway.
 */
export type AppGatewayHostType = (typeof AppGatewayHostType)[keyof typeof AppGatewayHostType];

export const AppGatewayType = {
    /**
     * Default value. This value is unused.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * TCP Proxy based BeyondCorp Connection. API will default to this if unset.
     */
    TcpProxy: "TCP_PROXY",
} as const;

/**
 * Required. The type of network connectivity used by the AppGateway.
 */
export type AppGatewayType = (typeof AppGatewayType)[keyof typeof AppGatewayType];

export const ConfigTransportProtocol = {
    /**
     * Default value. This value is unused.
     */
    TransportProtocolUnspecified: "TRANSPORT_PROTOCOL_UNSPECIFIED",
    /**
     * TCP protocol.
     */
    Tcp: "TCP",
} as const;

/**
 * Required. Immutable. The transport protocol used between the client and the server.
 */
export type ConfigTransportProtocol = (typeof ConfigTransportProtocol)[keyof typeof ConfigTransportProtocol];

export const ConnectionType = {
    /**
     * Default value. This value is unused.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * TCP Proxy based BeyondCorp Connection. API will default to this if unset.
     */
    TcpProxy: "TCP_PROXY",
} as const;

/**
 * Required. The type of network connectivity used by the connection.
 */
export type ConnectionType = (typeof ConnectionType)[keyof typeof ConnectionType];

export const GatewayType = {
    /**
     * Default value. This value is unused.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * Gateway hosted in a GCP regional managed instance group.
     */
    GcpRegionalMig: "GCP_REGIONAL_MIG",
} as const;

/**
 * Required. The type of hosting used by the gateway.
 */
export type GatewayType = (typeof GatewayType)[keyof typeof GatewayType];

export const GoogleCloudBeyondcorpAppconnectionsV1alphaAppConnectionGatewayType = {
    /**
     * Default value. This value is unused.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * Gateway hosted in a GCP regional managed instance group.
     */
    GcpRegionalMig: "GCP_REGIONAL_MIG",
} as const;

/**
 * Required. The type of hosting used by the gateway.
 */
export type GoogleCloudBeyondcorpAppconnectionsV1alphaAppConnectionGatewayType = (typeof GoogleCloudBeyondcorpAppconnectionsV1alphaAppConnectionGatewayType)[keyof typeof GoogleCloudBeyondcorpAppconnectionsV1alphaAppConnectionGatewayType];

export const GoogleCloudBeyondcorpAppconnectorsV1alphaResourceInfoStatus = {
    /**
     * Health status is unknown: not initialized or failed to retrieve.
     */
    HealthStatusUnspecified: "HEALTH_STATUS_UNSPECIFIED",
    /**
     * The resource is healthy.
     */
    Healthy: "HEALTHY",
    /**
     * The resource is unhealthy.
     */
    Unhealthy: "UNHEALTHY",
    /**
     * The resource is unresponsive.
     */
    Unresponsive: "UNRESPONSIVE",
    /**
     * Some sub-resources are UNHEALTHY.
     */
    Degraded: "DEGRADED",
} as const;

/**
 * Overall health status. Overall status is derived based on the status of each sub level resources.
 */
export type GoogleCloudBeyondcorpAppconnectorsV1alphaResourceInfoStatus = (typeof GoogleCloudBeyondcorpAppconnectorsV1alphaResourceInfoStatus)[keyof typeof GoogleCloudBeyondcorpAppconnectorsV1alphaResourceInfoStatus];

export const GoogleIamV1AuditLogConfigLogType = {
    /**
     * Default case. Should never be this.
     */
    LogTypeUnspecified: "LOG_TYPE_UNSPECIFIED",
    /**
     * Admin reads. Example: CloudIAM getIamPolicy
     */
    AdminRead: "ADMIN_READ",
    /**
     * Data writes. Example: CloudSQL Users create
     */
    DataWrite: "DATA_WRITE",
    /**
     * Data reads. Example: CloudSQL Users list
     */
    DataRead: "DATA_READ",
} as const;

/**
 * The log type that this config enables.
 */
export type GoogleIamV1AuditLogConfigLogType = (typeof GoogleIamV1AuditLogConfigLogType)[keyof typeof GoogleIamV1AuditLogConfigLogType];

export const ResourceInfoStatus = {
    /**
     * Health status is unknown: not initialized or failed to retrieve.
     */
    HealthStatusUnspecified: "HEALTH_STATUS_UNSPECIFIED",
    /**
     * The resource is healthy.
     */
    Healthy: "HEALTHY",
    /**
     * The resource is unhealthy.
     */
    Unhealthy: "UNHEALTHY",
    /**
     * The resource is unresponsive.
     */
    Unresponsive: "UNRESPONSIVE",
    /**
     * Some sub-resources are UNHEALTHY.
     */
    Degraded: "DEGRADED",
} as const;

/**
 * Overall health status. Overall status is derived based on the status of each sub level resources.
 */
export type ResourceInfoStatus = (typeof ResourceInfoStatus)[keyof typeof ResourceInfoStatus];

export const SubscriptionSku = {
    /**
     * Default value. This value is unused.
     */
    SkuUnspecified: "SKU_UNSPECIFIED",
    /**
     * Represents BeyondCorp Standard SKU.
     */
    BceStandardSku: "BCE_STANDARD_SKU",
} as const;

/**
 * Required. SKU of subscription.
 */
export type SubscriptionSku = (typeof SubscriptionSku)[keyof typeof SubscriptionSku];

export const SubscriptionType = {
    /**
     * Default value. This value is unused.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * Represents a trial subscription.
     */
    Trial: "TRIAL",
    /**
     * Represents a paid subscription.
     */
    Paid: "PAID",
    /**
     * Reresents an allowlisted subscription.
     */
    Allowlist: "ALLOWLIST",
} as const;

/**
 * Required. Type of subscription.
 */
export type SubscriptionType = (typeof SubscriptionType)[keyof typeof SubscriptionType];
