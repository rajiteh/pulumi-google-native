// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const RepositoryFormat = {
    /**
     * Unspecified package format.
     */
    FormatUnspecified: "FORMAT_UNSPECIFIED",
    /**
     * Docker package format.
     */
    Docker: "DOCKER",
    /**
     * Maven package format.
     */
    Maven: "MAVEN",
    /**
     * NPM package format.
     */
    Npm: "NPM",
    /**
     * APT package format.
     */
    Apt: "APT",
    /**
     * YUM package format.
     */
    Yum: "YUM",
    /**
     * GooGet package format.
     */
    Googet: "GOOGET",
    /**
     * Python package format.
     */
    Python: "PYTHON",
} as const;

/**
 * The format of packages that are stored in the repository.
 */
export type RepositoryFormat = (typeof RepositoryFormat)[keyof typeof RepositoryFormat];
