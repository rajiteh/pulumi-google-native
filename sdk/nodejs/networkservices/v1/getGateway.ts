// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets details of a single Gateway.
 */
export function getGateway(args: GetGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:networkservices/v1:getGateway", {
        "gatewayId": args.gatewayId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetGatewayArgs {
    gatewayId: string;
    location: string;
    project?: string;
}

export interface GetGatewayResult {
    /**
     * Optional. Zero or one IPv4-address on which the Gateway will receive the traffic. When no address is provided, an IP from the subnetwork is allocated This field only applies to gateways of type 'SECURE_WEB_GATEWAY'. Gateways of type 'OPEN_MESH' listen on 0.0.0.0.
     */
    readonly addresses: string[];
    /**
     * Optional. A fully-qualified Certificates URL reference. The proxy presents a Certificate (selected based on SNI) when establishing a TLS connection. This feature only applies to gateways of type 'SECURE_WEB_GATEWAY'.
     */
    readonly certificateUrls: string[];
    /**
     * The timestamp when the resource was created.
     */
    readonly createTime: string;
    /**
     * Optional. A free-text description of the resource. Max length 1024 characters.
     */
    readonly description: string;
    /**
     * Optional. A fully-qualified GatewaySecurityPolicy URL reference. Defines how a server should apply security policy to inbound (VM to Proxy) initiated connections. For example: `projects/*&#47;locations/*&#47;gatewaySecurityPolicies/swg-policy`. This policy is specific to gateways of type 'SECURE_WEB_GATEWAY'.
     */
    readonly gatewaySecurityPolicy: string;
    /**
     * Optional. Set of label tags associated with the Gateway resource.
     */
    readonly labels: {[key: string]: string};
    /**
     * Name of the Gateway resource. It matches pattern `projects/*&#47;locations/*&#47;gateways/`.
     */
    readonly name: string;
    /**
     * Optional. The relative resource name identifying the VPC network that is using this configuration. For example: `projects/*&#47;global/networks/network-1`. Currently, this field is specific to gateways of type 'SECURE_WEB_GATEWAY'.
     */
    readonly network: string;
    /**
     * One or more port numbers (1-65535), on which the Gateway will receive traffic. The proxy binds to the specified ports. Gateways of type 'SECURE_WEB_GATEWAY' are limited to 1 port. Gateways of type 'OPEN_MESH' listen on 0.0.0.0 and support multiple ports.
     */
    readonly ports: number[];
    /**
     * Optional. Scope determines how configuration across multiple Gateway instances are merged. The configuration for multiple Gateway instances with the same scope will be merged as presented as a single coniguration to the proxy/load balancer. Max length 64 characters. Scope should start with a letter and can only have letters, numbers, hyphens.
     */
    readonly scope: string;
    /**
     * Server-defined URL of this resource
     */
    readonly selfLink: string;
    /**
     * Optional. A fully-qualified ServerTLSPolicy URL reference. Specifies how TLS traffic is terminated. If empty, TLS termination is disabled.
     */
    readonly serverTlsPolicy: string;
    /**
     * Optional. The relative resource name identifying the subnetwork in which this SWG is allocated. For example: `projects/*&#47;regions/us-central1/subnetworks/network-1` Currently, this field is specific to gateways of type 'SECURE_WEB_GATEWAY".
     */
    readonly subnetwork: string;
    /**
     * Immutable. The type of the customer managed gateway. This field is required. If unspecified, an error is returned.
     */
    readonly type: string;
    /**
     * The timestamp when the resource was updated.
     */
    readonly updateTime: string;
}
/**
 * Gets details of a single Gateway.
 */
export function getGatewayOutput(args: GetGatewayOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGatewayResult> {
    return pulumi.output(args).apply((a: any) => getGateway(a, opts))
}

export interface GetGatewayOutputArgs {
    gatewayId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
