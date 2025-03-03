// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Returns a specified artifact.
 */
export function getDeploymentArtifact(args: GetDeploymentArtifactArgs, opts?: pulumi.InvokeOptions): Promise<GetDeploymentArtifactResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:apigeeregistry/v1:getDeploymentArtifact", {
        "apiId": args.apiId,
        "artifactId": args.artifactId,
        "deploymentId": args.deploymentId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetDeploymentArtifactArgs {
    apiId: string;
    artifactId: string;
    deploymentId: string;
    location: string;
    project?: string;
}

export interface GetDeploymentArtifactResult {
    /**
     * Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
     */
    readonly annotations: {[key: string]: string};
    /**
     * Input only. The contents of the artifact. Provided by API callers when artifacts are created or replaced. To access the contents of an artifact, use GetArtifactContents.
     */
    readonly contents: string;
    /**
     * Creation timestamp.
     */
    readonly createTime: string;
    /**
     * A SHA-256 hash of the artifact's contents. If the artifact is gzipped, this is the hash of the uncompressed artifact.
     */
    readonly hash: string;
    /**
     * Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with "registry.googleapis.com/" and cannot be changed.
     */
    readonly labels: {[key: string]: string};
    /**
     * A content type specifier for the artifact. Content type specifiers are Media Types (https://en.wikipedia.org/wiki/Media_type) with a possible "schema" parameter that specifies a schema for the stored information. Content types can specify compression. Currently only GZip compression is supported (indicated with "+gzip").
     */
    readonly mimeType: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The size of the artifact in bytes. If the artifact is gzipped, this is the size of the uncompressed artifact.
     */
    readonly sizeBytes: number;
    /**
     * Last update timestamp.
     */
    readonly updateTime: string;
}
/**
 * Returns a specified artifact.
 */
export function getDeploymentArtifactOutput(args: GetDeploymentArtifactOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDeploymentArtifactResult> {
    return pulumi.output(args).apply((a: any) => getDeploymentArtifact(a, opts))
}

export interface GetDeploymentArtifactOutputArgs {
    apiId: pulumi.Input<string>;
    artifactId: pulumi.Input<string>;
    deploymentId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
