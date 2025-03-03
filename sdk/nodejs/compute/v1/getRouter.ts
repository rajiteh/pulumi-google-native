// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns the specified Router resource.
 */
export function getRouter(args: GetRouterArgs, opts?: pulumi.InvokeOptions): Promise<GetRouterResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/v1:getRouter", {
        "project": args.project,
        "region": args.region,
        "router": args.router,
    }, opts);
}

export interface GetRouterArgs {
    project?: string;
    region: string;
    router: string;
}

export interface GetRouterResult {
    /**
     * BGP information specific to this router.
     */
    readonly bgp: outputs.compute.v1.RouterBgpResponse;
    /**
     * BGP information that must be configured into the routing stack to establish BGP peering. This information must specify the peer ASN and either the interface name, IP address, or peer IP address. Please refer to RFC4273.
     */
    readonly bgpPeers: outputs.compute.v1.RouterBgpPeerResponse[];
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Indicates if a router is dedicated for use with encrypted VLAN attachments (interconnectAttachments).
     */
    readonly encryptedInterconnectRouter: boolean;
    /**
     * Router interfaces. Each interface requires either one linked resource, (for example, linkedVpnTunnel), or IP address and IP address range (for example, ipRange), or both.
     */
    readonly interfaces: outputs.compute.v1.RouterInterfaceResponse[];
    /**
     * Type of resource. Always compute#router for routers.
     */
    readonly kind: string;
    /**
     * Keys used for MD5 authentication.
     */
    readonly md5AuthenticationKeys: outputs.compute.v1.RouterMd5AuthenticationKeyResponse[];
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * A list of NAT services created in this router.
     */
    readonly nats: outputs.compute.v1.RouterNatResponse[];
    /**
     * URI of the network to which this router belongs.
     */
    readonly network: string;
    /**
     * URI of the region where the router resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly region: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
}
/**
 * Returns the specified Router resource.
 */
export function getRouterOutput(args: GetRouterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouterResult> {
    return pulumi.output(args).apply((a: any) => getRouter(a, opts))
}

export interface GetRouterOutputArgs {
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
    router: pulumi.Input<string>;
}
