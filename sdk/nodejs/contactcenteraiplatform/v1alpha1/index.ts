// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { ContactCenterArgs } from "./contactCenter";
export type ContactCenter = import("./contactCenter").ContactCenter;
export const ContactCenter: typeof import("./contactCenter").ContactCenter = null as any;
utilities.lazyLoad(exports, ["ContactCenter"], () => require("./contactCenter"));

export { GetContactCenterArgs, GetContactCenterResult, GetContactCenterOutputArgs } from "./getContactCenter";
export const getContactCenter: typeof import("./getContactCenter").getContactCenter = null as any;
export const getContactCenterOutput: typeof import("./getContactCenter").getContactCenterOutput = null as any;
utilities.lazyLoad(exports, ["getContactCenter","getContactCenterOutput"], () => require("./getContactCenter"));


// Export enums:
export * from "../../types/enums/contactcenteraiplatform/v1alpha1";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:contactcenteraiplatform/v1alpha1:ContactCenter":
                return new ContactCenter(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "contactcenteraiplatform/v1alpha1", _module)