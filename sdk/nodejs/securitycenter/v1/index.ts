// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./finding";
export * from "./notificationConfig";
export * from "./policy";
export * from "./source";

// Import resources to register:
import { Finding } from "./finding";
import { NotificationConfig } from "./notificationConfig";
import { Policy } from "./policy";
import { Source } from "./source";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-cloud:securitycenter/v1:Finding":
                return new Finding(name, <any>undefined, { urn })
            case "google-cloud:securitycenter/v1:NotificationConfig":
                return new NotificationConfig(name, <any>undefined, { urn })
            case "google-cloud:securitycenter/v1:Policy":
                return new Policy(name, <any>undefined, { urn })
            case "google-cloud:securitycenter/v1:Source":
                return new Source(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-cloud", "securitycenter/v1", _module)