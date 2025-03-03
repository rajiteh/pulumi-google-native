// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns the specified PacketMirroring resource.
 */
export function getPacketMirroring(args: GetPacketMirroringArgs, opts?: pulumi.InvokeOptions): Promise<GetPacketMirroringResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/alpha:getPacketMirroring", {
        "packetMirroring": args.packetMirroring,
        "project": args.project,
        "region": args.region,
    }, opts);
}

export interface GetPacketMirroringArgs {
    packetMirroring: string;
    project?: string;
    region: string;
}

export interface GetPacketMirroringResult {
    /**
     * The Forwarding Rule resource of type loadBalancingScheme=INTERNAL that will be used as collector for mirrored traffic. The specified forwarding rule must have isMirroringCollector set to true.
     */
    readonly collectorIlb: outputs.compute.alpha.PacketMirroringForwardingRuleInfoResponse;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Indicates whether or not this packet mirroring takes effect. If set to FALSE, this packet mirroring policy will not be enforced on the network. The default is TRUE.
     */
    readonly enable: string;
    /**
     * Filter for mirrored traffic. If unspecified, all traffic is mirrored.
     */
    readonly filter: outputs.compute.alpha.PacketMirroringFilterResponse;
    /**
     * Type of the resource. Always compute#packetMirroring for packet mirrorings.
     */
    readonly kind: string;
    /**
     * PacketMirroring mirroredResourceInfos. MirroredResourceInfo specifies a set of mirrored VM instances, subnetworks and/or tags for which traffic from/to all VM instances will be mirrored.
     */
    readonly mirroredResources: outputs.compute.alpha.PacketMirroringMirroredResourceInfoResponse;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a NIC in the given network. All mirrored subnetworks should belong to the given network.
     */
    readonly network: outputs.compute.alpha.PacketMirroringNetworkInfoResponse;
    /**
     * The priority of applying this configuration. Priority is used to break ties in cases where there is more than one matching rule. In the case of two rules that apply for a given Instance, the one with the lowest-numbered priority value wins. Default value is 1000. Valid range is 0 through 65535.
     */
    readonly priority: number;
    /**
     * URI of the region where the packetMirroring resides.
     */
    readonly region: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    readonly selfLinkWithId: string;
}
/**
 * Returns the specified PacketMirroring resource.
 */
export function getPacketMirroringOutput(args: GetPacketMirroringOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPacketMirroringResult> {
    return pulumi.output(args).apply((a: any) => getPacketMirroring(a, opts))
}

export interface GetPacketMirroringOutputArgs {
    packetMirroring: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
}
