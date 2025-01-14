// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetInstanceArgs, GetInstanceResult, GetInstanceOutputArgs } from "./getInstance";
export const getInstance: typeof import("./getInstance").getInstance = null as any;
export const getInstanceOutput: typeof import("./getInstance").getInstanceOutput = null as any;
utilities.lazyLoad(exports, ["getInstance","getInstanceOutput"], () => require("./getInstance"));

export { GetNfsShareArgs, GetNfsShareResult, GetNfsShareOutputArgs } from "./getNfsShare";
export const getNfsShare: typeof import("./getNfsShare").getNfsShare = null as any;
export const getNfsShareOutput: typeof import("./getNfsShare").getNfsShareOutput = null as any;
utilities.lazyLoad(exports, ["getNfsShare","getNfsShareOutput"], () => require("./getNfsShare"));

export { GetProvisioningConfigArgs, GetProvisioningConfigResult, GetProvisioningConfigOutputArgs } from "./getProvisioningConfig";
export const getProvisioningConfig: typeof import("./getProvisioningConfig").getProvisioningConfig = null as any;
export const getProvisioningConfigOutput: typeof import("./getProvisioningConfig").getProvisioningConfigOutput = null as any;
utilities.lazyLoad(exports, ["getProvisioningConfig","getProvisioningConfigOutput"], () => require("./getProvisioningConfig"));

export { GetSnapshotArgs, GetSnapshotResult, GetSnapshotOutputArgs } from "./getSnapshot";
export const getSnapshot: typeof import("./getSnapshot").getSnapshot = null as any;
export const getSnapshotOutput: typeof import("./getSnapshot").getSnapshotOutput = null as any;
utilities.lazyLoad(exports, ["getSnapshot","getSnapshotOutput"], () => require("./getSnapshot"));

export { InstanceArgs } from "./instance";
export type Instance = import("./instance").Instance;
export const Instance: typeof import("./instance").Instance = null as any;
utilities.lazyLoad(exports, ["Instance"], () => require("./instance"));

export { NfsShareArgs } from "./nfsShare";
export type NfsShare = import("./nfsShare").NfsShare;
export const NfsShare: typeof import("./nfsShare").NfsShare = null as any;
utilities.lazyLoad(exports, ["NfsShare"], () => require("./nfsShare"));

export { ProvisioningConfigArgs } from "./provisioningConfig";
export type ProvisioningConfig = import("./provisioningConfig").ProvisioningConfig;
export const ProvisioningConfig: typeof import("./provisioningConfig").ProvisioningConfig = null as any;
utilities.lazyLoad(exports, ["ProvisioningConfig"], () => require("./provisioningConfig"));

export { SnapshotArgs } from "./snapshot";
export type Snapshot = import("./snapshot").Snapshot;
export const Snapshot: typeof import("./snapshot").Snapshot = null as any;
utilities.lazyLoad(exports, ["Snapshot"], () => require("./snapshot"));


// Export enums:
export * from "../../types/enums/baremetalsolution/v2";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:baremetalsolution/v2:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "google-native:baremetalsolution/v2:NfsShare":
                return new NfsShare(name, <any>undefined, { urn })
            case "google-native:baremetalsolution/v2:ProvisioningConfig":
                return new ProvisioningConfig(name, <any>undefined, { urn })
            case "google-native:baremetalsolution/v2:Snapshot":
                return new Snapshot(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "baremetalsolution/v2", _module)
