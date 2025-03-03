// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets a device registry configuration.
 */
export function getRegistry(args: GetRegistryArgs, opts?: pulumi.InvokeOptions): Promise<GetRegistryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:cloudiot/v1:getRegistry", {
        "location": args.location,
        "project": args.project,
        "registryId": args.registryId,
    }, opts);
}

export interface GetRegistryArgs {
    location: string;
    project?: string;
    registryId: string;
}

export interface GetRegistryResult {
    /**
     * The credentials used to verify the device credentials. No more than 10 credentials can be bound to a single registry at a time. The verification process occurs at the time of device creation or update. If this field is empty, no verification is performed. Otherwise, the credentials of a newly created device or added credentials of an updated device should be signed with one of these registry credentials. Note, however, that existing devices will never be affected by modifications to this list of credentials: after a device has been successfully created in a registry, it should be able to connect even if its registry credentials are revoked, deleted, or modified.
     */
    readonly credentials: outputs.cloudiot.v1.RegistryCredentialResponse[];
    /**
     * The configuration for notification of telemetry events received from the device. All telemetry events that were successfully published by the device and acknowledged by Cloud IoT Core are guaranteed to be delivered to Cloud Pub/Sub. If multiple configurations match a message, only the first matching configuration is used. If you try to publish a device telemetry event using MQTT without specifying a Cloud Pub/Sub topic for the device's registry, the connection closes automatically. If you try to do so using an HTTP connection, an error is returned. Up to 10 configurations may be provided.
     */
    readonly eventNotificationConfigs: outputs.cloudiot.v1.EventNotificationConfigResponse[];
    /**
     * The DeviceService (HTTP) configuration for this device registry.
     */
    readonly httpConfig: outputs.cloudiot.v1.HttpConfigResponse;
    /**
     * **Beta Feature** The default logging verbosity for activity from devices in this registry. The verbosity level can be overridden by Device.log_level.
     */
    readonly logLevel: string;
    /**
     * The MQTT configuration for this device registry.
     */
    readonly mqttConfig: outputs.cloudiot.v1.MqttConfigResponse;
    /**
     * The resource path name. For example, `projects/example-project/locations/us-central1/registries/my-registry`.
     */
    readonly name: string;
    /**
     * The configuration for notification of new states received from the device. State updates are guaranteed to be stored in the state history, but notifications to Cloud Pub/Sub are not guaranteed. For example, if permissions are misconfigured or the specified topic doesn't exist, no notification will be published but the state will still be stored in Cloud IoT Core.
     */
    readonly stateNotificationConfig: outputs.cloudiot.v1.StateNotificationConfigResponse;
}
/**
 * Gets a device registry configuration.
 */
export function getRegistryOutput(args: GetRegistryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegistryResult> {
    return pulumi.output(args).apply((a: any) => getRegistry(a, opts))
}

export interface GetRegistryOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    registryId: pulumi.Input<string>;
}
