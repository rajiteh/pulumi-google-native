// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns the specified network.
 */
export function getNetwork(args: GetNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/beta:getNetwork", {
        "network": args.network,
        "project": args.project,
    }, opts);
}

export interface GetNetworkArgs {
    network: string;
    project?: string;
}

export interface GetNetworkResult {
    /**
     * Must be set to create a VPC network. If not set, a legacy network is created. When set to true, the VPC network is created in auto mode. When set to false, the VPC network is created in custom mode. An auto mode VPC network starts with one subnet per region. Each subnet has a predetermined range as described in Auto mode VPC network IP ranges. For custom mode VPC networks, you can add subnets using the subnetworks insert method.
     */
    readonly autoCreateSubnetworks: boolean;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this field when you create the resource.
     */
    readonly description: string;
    /**
     * Enable ULA internal ipv6 on this network. Enabling this feature will assign a /48 from google defined ULA prefix fd20::/20. .
     */
    readonly enableUlaInternalIpv6: boolean;
    /**
     * URL of the firewall policy the network is associated with.
     */
    readonly firewallPolicy: string;
    /**
     * The gateway address for default routing out of the network, selected by Google Cloud.
     */
    readonly gatewayIPv4: string;
    /**
     * When enabling ula internal ipv6, caller optionally can specify the /48 range they want from the google defined ULA prefix fd20::/20. The input must be a valid /48 ULA IPv6 address and must be within the fd20::/20. Operation will fail if the speficied /48 is already in used by another resource. If the field is not speficied, then a /48 range will be randomly allocated from fd20::/20 and returned via this field. .
     */
    readonly internalIpv6Range: string;
    /**
     * Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.
     *
     * @deprecated Deprecated in favor of subnet mode networks. The range of internal addresses that are legal on this network. This range is a CIDR specification, for example: 192.168.0.0/16. Provided by the client when the network is created.
     */
    readonly ipv4Range: string;
    /**
     * Type of the resource. Always compute#network for networks.
     */
    readonly kind: string;
    /**
     * Maximum Transmission Unit in bytes. The minimum value for this field is 1300 and the maximum value is 8896. The suggested value is 1500, which is the default MTU used on the Internet, or 8896 if you want to use Jumbo frames. If unspecified, the value defaults to 1460.
     */
    readonly mtu: number;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
     */
    readonly name: string;
    /**
     * The network firewall policy enforcement order. Can be either AFTER_CLASSIC_FIREWALL or BEFORE_CLASSIC_FIREWALL. Defaults to AFTER_CLASSIC_FIREWALL if the field is not specified.
     */
    readonly networkFirewallPolicyEnforcementOrder: string;
    /**
     * A list of network peerings for the resource.
     */
    readonly peerings: outputs.compute.beta.NetworkPeeringResponse[];
    /**
     * The network-level routing configuration for this network. Used by Cloud Router to determine what type of network-wide routing behavior to enforce.
     */
    readonly routingConfig: outputs.compute.beta.NetworkRoutingConfigResponse;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    readonly selfLinkWithId: string;
    /**
     * Server-defined fully-qualified URLs for all subnetworks in this VPC network.
     */
    readonly subnetworks: string[];
}
/**
 * Returns the specified network.
 */
export function getNetworkOutput(args: GetNetworkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetworkResult> {
    return pulumi.output(args).apply((a: any) => getNetwork(a, opts))
}

export interface GetNetworkOutputArgs {
    network: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
