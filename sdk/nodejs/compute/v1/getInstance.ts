// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns the specified Instance resource.
 */
export function getInstance(args: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/v1:getInstance", {
        "instance": args.instance,
        "project": args.project,
        "zone": args.zone,
    }, opts);
}

export interface GetInstanceArgs {
    instance: string;
    project?: string;
    zone: string;
}

export interface GetInstanceResult {
    /**
     * Controls for advanced machine-related behavior features.
     */
    readonly advancedMachineFeatures: outputs.compute.v1.AdvancedMachineFeaturesResponse;
    /**
     * Allows this instance to send and receive packets with non-matching destination or source IPs. This is required if you plan to use this instance to forward routes. For more information, see Enabling IP Forwarding .
     */
    readonly canIpForward: boolean;
    readonly confidentialInstanceConfig: outputs.compute.v1.ConfidentialInstanceConfigResponse;
    /**
     * The CPU platform used by this instance.
     */
    readonly cpuPlatform: string;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * Whether the resource should be protected against deletion.
     */
    readonly deletionProtection: boolean;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Array of disks associated with this instance. Persistent disks must be created before you can assign them.
     */
    readonly disks: outputs.compute.v1.AttachedDiskResponse[];
    /**
     * Enables display device for the instance.
     */
    readonly displayDevice: outputs.compute.v1.DisplayDeviceResponse;
    /**
     * Specifies a fingerprint for this resource, which is essentially a hash of the instance's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update the instance. You must always provide an up-to-date fingerprint hash in order to update the instance. To see the latest fingerprint, make get() request to the instance.
     */
    readonly fingerprint: string;
    /**
     * A list of the type and count of accelerator cards attached to the instance.
     */
    readonly guestAccelerators: outputs.compute.v1.AcceleratorConfigResponse[];
    /**
     * Specifies the hostname of the instance. The specified hostname must be RFC1035 compliant. If hostname is not specified, the default hostname is [INSTANCE_NAME].c.[PROJECT_ID].internal when using the global DNS, and [INSTANCE_NAME].[ZONE].c.[PROJECT_ID].internal when using zonal DNS.
     */
    readonly hostname: string;
    /**
     * KeyRevocationActionType of the instance. Supported options are "STOP" and "NONE". The default value is "NONE" if it is not specified.
     */
    readonly keyRevocationActionType: string;
    /**
     * Type of the resource. Always compute#instance for instances.
     */
    readonly kind: string;
    /**
     * A fingerprint for this request, which is essentially a hash of the label's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels. To see the latest fingerprint, make get() request to the instance.
     */
    readonly labelFingerprint: string;
    /**
     * Labels to apply to this instance. These can be later modified by the setLabels method.
     */
    readonly labels: {[key: string]: string};
    /**
     * Last start timestamp in RFC3339 text format.
     */
    readonly lastStartTimestamp: string;
    /**
     * Last stop timestamp in RFC3339 text format.
     */
    readonly lastStopTimestamp: string;
    /**
     * Last suspended timestamp in RFC3339 text format.
     */
    readonly lastSuspendedTimestamp: string;
    /**
     * Full or partial URL of the machine type resource to use for this instance, in the format: zones/zone/machineTypes/machine-type. This is provided by the client when the instance is created. For example, the following is a valid partial url to a predefined machine type: zones/us-central1-f/machineTypes/n1-standard-1 To create a custom machine type, provide a URL to a machine type in the following format, where CPUS is 1 or an even number up to 32 (2, 4, 6, ... 24, etc), and MEMORY is the total memory for this instance. Memory must be a multiple of 256 MB and must be supplied in MB (e.g. 5 GB of memory is 5120 MB): zones/zone/machineTypes/custom-CPUS-MEMORY For example: zones/us-central1-f/machineTypes/custom-4-5120 For a full list of restrictions, read the Specifications for custom machine types.
     */
    readonly machineType: string;
    /**
     * The metadata key/value pairs assigned to this instance. This includes custom metadata and predefined keys.
     */
    readonly metadata: outputs.compute.v1.MetadataResponse;
    /**
     * Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as minCpuPlatform: "Intel Haswell" or minCpuPlatform: "Intel Sandy Bridge".
     */
    readonly minCpuPlatform: string;
    /**
     * The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * An array of network configurations for this instance. These specify how interfaces are configured to interact with other network services, such as connecting to the internet. Multiple interfaces are supported per instance.
     */
    readonly networkInterfaces: outputs.compute.v1.NetworkInterfaceResponse[];
    readonly networkPerformanceConfig: outputs.compute.v1.NetworkPerformanceConfigResponse;
    /**
     * Input only. [Input Only] Additional params passed with the request, but not persisted as part of resource payload.
     */
    readonly params: outputs.compute.v1.InstanceParamsResponse;
    /**
     * The private IPv6 google access type for the VM. If not specified, use INHERIT_FROM_SUBNETWORK as default.
     */
    readonly privateIpv6GoogleAccess: string;
    /**
     * Specifies the reservations that this instance can consume from.
     */
    readonly reservationAffinity: outputs.compute.v1.ReservationAffinityResponse;
    /**
     * Resource policies applied to this instance.
     */
    readonly resourcePolicies: string[];
    /**
     * Specifies values set for instance attributes as compared to the values requested by user in the corresponding input only field.
     */
    readonly resourceStatus: outputs.compute.v1.ResourceStatusResponse;
    /**
     * Reserved for future use.
     */
    readonly satisfiesPzs: boolean;
    /**
     * Sets the scheduling options for this instance.
     */
    readonly scheduling: outputs.compute.v1.SchedulingResponse;
    /**
     * Server-defined URL for this resource.
     */
    readonly selfLink: string;
    /**
     * A list of service accounts, with their specified scopes, authorized for this instance. Only one service account per VM instance is supported. Service accounts generate access tokens that can be accessed through the metadata server and used to authenticate applications on the instance. See Service Accounts for more information.
     */
    readonly serviceAccounts: outputs.compute.v1.ServiceAccountResponse[];
    readonly shieldedInstanceConfig: outputs.compute.v1.ShieldedInstanceConfigResponse;
    readonly shieldedInstanceIntegrityPolicy: outputs.compute.v1.ShieldedInstanceIntegrityPolicyResponse;
    /**
     * Source machine image
     */
    readonly sourceMachineImage: string;
    /**
     * Source machine image encryption key when creating an instance from a machine image.
     */
    readonly sourceMachineImageEncryptionKey: outputs.compute.v1.CustomerEncryptionKeyResponse;
    /**
     * Whether a VM has been restricted for start because Compute Engine has detected suspicious activity.
     */
    readonly startRestricted: boolean;
    /**
     * The status of the instance. One of the following values: PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED. For more information about the status of the instance, see Instance life cycle.
     */
    readonly status: string;
    /**
     * An optional, human-readable explanation of the status.
     */
    readonly statusMessage: string;
    /**
     * Tags to apply to this instance. Tags are used to identify valid sources or targets for network firewalls and are specified by the client during instance creation. The tags can be later modified by the setTags method. Each tag within the list must comply with RFC1035. Multiple tags can be specified via the 'tags.items' field.
     */
    readonly tags: outputs.compute.v1.TagsResponse;
    /**
     * URL of the zone where the instance resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly zone: string;
}
/**
 * Returns the specified Instance resource.
 */
export function getInstanceOutput(args: GetInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceResult> {
    return pulumi.output(args).apply((a: any) => getInstance(a, opts))
}

export interface GetInstanceOutputArgs {
    instance: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    zone: pulumi.Input<string>;
}
