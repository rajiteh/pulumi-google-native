// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AuditLogConfigLogType = {
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
export type AuditLogConfigLogType = (typeof AuditLogConfigLogType)[keyof typeof AuditLogConfigLogType];

export const ExecutionConfigUsagesItem = {
    /**
     * Default value. This value is unused.
     */
    ExecutionEnvironmentUsageUnspecified: "EXECUTION_ENVIRONMENT_USAGE_UNSPECIFIED",
    /**
     * Use for rendering.
     */
    Render: "RENDER",
    /**
     * Use for deploying and deployment hooks.
     */
    Deploy: "DEPLOY",
    /**
     * Use for deployment verification.
     */
    Verify: "VERIFY",
} as const;

export type ExecutionConfigUsagesItem = (typeof ExecutionConfigUsagesItem)[keyof typeof ExecutionConfigUsagesItem];
