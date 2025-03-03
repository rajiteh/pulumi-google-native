// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const DockerRepositoryPublicRepository = {
    /**
     * Unspecified repository.
     */
    PublicRepositoryUnspecified: "PUBLIC_REPOSITORY_UNSPECIFIED",
    /**
     * Docker Hub.
     */
    DockerHub: "DOCKER_HUB",
} as const;

/**
 * One of the publicly available Docker repositories supported by Artifact Registry.
 */
export type DockerRepositoryPublicRepository = (typeof DockerRepositoryPublicRepository)[keyof typeof DockerRepositoryPublicRepository];

export const MavenRepositoryConfigVersionPolicy = {
    /**
     * VERSION_POLICY_UNSPECIFIED - the version policy is not defined. When the version policy is not defined, no validation is performed for the versions.
     */
    VersionPolicyUnspecified: "VERSION_POLICY_UNSPECIFIED",
    /**
     * RELEASE - repository will accept only Release versions.
     */
    Release: "RELEASE",
    /**
     * SNAPSHOT - repository will accept only Snapshot versions.
     */
    Snapshot: "SNAPSHOT",
} as const;

/**
 * Version policy defines the versions that the registry will accept.
 */
export type MavenRepositoryConfigVersionPolicy = (typeof MavenRepositoryConfigVersionPolicy)[keyof typeof MavenRepositoryConfigVersionPolicy];

export const MavenRepositoryPublicRepository = {
    /**
     * Unspecified repository.
     */
    PublicRepositoryUnspecified: "PUBLIC_REPOSITORY_UNSPECIFIED",
    /**
     * Maven Central.
     */
    MavenCentral: "MAVEN_CENTRAL",
} as const;

/**
 * One of the publicly available Maven repositories supported by Artifact Registry.
 */
export type MavenRepositoryPublicRepository = (typeof MavenRepositoryPublicRepository)[keyof typeof MavenRepositoryPublicRepository];

export const NpmRepositoryPublicRepository = {
    /**
     * Unspecified repository.
     */
    PublicRepositoryUnspecified: "PUBLIC_REPOSITORY_UNSPECIFIED",
    /**
     * npmjs.
     */
    Npmjs: "NPMJS",
} as const;

/**
 * One of the publicly available Npm repositories supported by Artifact Registry.
 */
export type NpmRepositoryPublicRepository = (typeof NpmRepositoryPublicRepository)[keyof typeof NpmRepositoryPublicRepository];

export const PythonRepositoryPublicRepository = {
    /**
     * Unspecified repository.
     */
    PublicRepositoryUnspecified: "PUBLIC_REPOSITORY_UNSPECIFIED",
    /**
     * PyPI.
     */
    Pypi: "PYPI",
} as const;

/**
 * One of the publicly available Python repositories supported by Artifact Registry.
 */
export type PythonRepositoryPublicRepository = (typeof PythonRepositoryPublicRepository)[keyof typeof PythonRepositoryPublicRepository];

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
    /**
     * Kubeflow Pipelines package format.
     */
    Kfp: "KFP",
} as const;

/**
 * The format of packages that are stored in the repository.
 */
export type RepositoryFormat = (typeof RepositoryFormat)[keyof typeof RepositoryFormat];

export const RepositoryMode = {
    /**
     * Unspecified mode.
     */
    ModeUnspecified: "MODE_UNSPECIFIED",
    /**
     * A standard repository storing artifacts.
     */
    StandardRepository: "STANDARD_REPOSITORY",
    /**
     * A virtual repository to serve artifacts from one or more sources.
     */
    VirtualRepository: "VIRTUAL_REPOSITORY",
    /**
     * A remote repository to serve artifacts from a remote source.
     */
    RemoteRepository: "REMOTE_REPOSITORY",
} as const;

/**
 * The mode of the repository.
 */
export type RepositoryMode = (typeof RepositoryMode)[keyof typeof RepositoryMode];
