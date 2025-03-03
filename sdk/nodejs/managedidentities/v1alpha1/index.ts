// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { BackupArgs } from "./backup";
export type Backup = import("./backup").Backup;
export const Backup: typeof import("./backup").Backup = null as any;
utilities.lazyLoad(exports, ["Backup"], () => require("./backup"));

export { DomainArgs } from "./domain";
export type Domain = import("./domain").Domain;
export const Domain: typeof import("./domain").Domain = null as any;
utilities.lazyLoad(exports, ["Domain"], () => require("./domain"));

export { DomainBackupIamBindingArgs } from "./domainBackupIamBinding";
export type DomainBackupIamBinding = import("./domainBackupIamBinding").DomainBackupIamBinding;
export const DomainBackupIamBinding: typeof import("./domainBackupIamBinding").DomainBackupIamBinding = null as any;
utilities.lazyLoad(exports, ["DomainBackupIamBinding"], () => require("./domainBackupIamBinding"));

export { DomainBackupIamMemberArgs } from "./domainBackupIamMember";
export type DomainBackupIamMember = import("./domainBackupIamMember").DomainBackupIamMember;
export const DomainBackupIamMember: typeof import("./domainBackupIamMember").DomainBackupIamMember = null as any;
utilities.lazyLoad(exports, ["DomainBackupIamMember"], () => require("./domainBackupIamMember"));

export { DomainBackupIamPolicyArgs } from "./domainBackupIamPolicy";
export type DomainBackupIamPolicy = import("./domainBackupIamPolicy").DomainBackupIamPolicy;
export const DomainBackupIamPolicy: typeof import("./domainBackupIamPolicy").DomainBackupIamPolicy = null as any;
utilities.lazyLoad(exports, ["DomainBackupIamPolicy"], () => require("./domainBackupIamPolicy"));

export { DomainIamBindingArgs } from "./domainIamBinding";
export type DomainIamBinding = import("./domainIamBinding").DomainIamBinding;
export const DomainIamBinding: typeof import("./domainIamBinding").DomainIamBinding = null as any;
utilities.lazyLoad(exports, ["DomainIamBinding"], () => require("./domainIamBinding"));

export { DomainIamMemberArgs } from "./domainIamMember";
export type DomainIamMember = import("./domainIamMember").DomainIamMember;
export const DomainIamMember: typeof import("./domainIamMember").DomainIamMember = null as any;
utilities.lazyLoad(exports, ["DomainIamMember"], () => require("./domainIamMember"));

export { DomainIamPolicyArgs } from "./domainIamPolicy";
export type DomainIamPolicy = import("./domainIamPolicy").DomainIamPolicy;
export const DomainIamPolicy: typeof import("./domainIamPolicy").DomainIamPolicy = null as any;
utilities.lazyLoad(exports, ["DomainIamPolicy"], () => require("./domainIamPolicy"));

export { GetBackupArgs, GetBackupResult, GetBackupOutputArgs } from "./getBackup";
export const getBackup: typeof import("./getBackup").getBackup = null as any;
export const getBackupOutput: typeof import("./getBackup").getBackupOutput = null as any;
utilities.lazyLoad(exports, ["getBackup","getBackupOutput"], () => require("./getBackup"));

export { GetDomainArgs, GetDomainResult, GetDomainOutputArgs } from "./getDomain";
export const getDomain: typeof import("./getDomain").getDomain = null as any;
export const getDomainOutput: typeof import("./getDomain").getDomainOutput = null as any;
utilities.lazyLoad(exports, ["getDomain","getDomainOutput"], () => require("./getDomain"));

export { GetDomainBackupIamPolicyArgs, GetDomainBackupIamPolicyResult, GetDomainBackupIamPolicyOutputArgs } from "./getDomainBackupIamPolicy";
export const getDomainBackupIamPolicy: typeof import("./getDomainBackupIamPolicy").getDomainBackupIamPolicy = null as any;
export const getDomainBackupIamPolicyOutput: typeof import("./getDomainBackupIamPolicy").getDomainBackupIamPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getDomainBackupIamPolicy","getDomainBackupIamPolicyOutput"], () => require("./getDomainBackupIamPolicy"));

export { GetDomainIamPolicyArgs, GetDomainIamPolicyResult, GetDomainIamPolicyOutputArgs } from "./getDomainIamPolicy";
export const getDomainIamPolicy: typeof import("./getDomainIamPolicy").getDomainIamPolicy = null as any;
export const getDomainIamPolicyOutput: typeof import("./getDomainIamPolicy").getDomainIamPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getDomainIamPolicy","getDomainIamPolicyOutput"], () => require("./getDomainIamPolicy"));

export { GetPeeringArgs, GetPeeringResult, GetPeeringOutputArgs } from "./getPeering";
export const getPeering: typeof import("./getPeering").getPeering = null as any;
export const getPeeringOutput: typeof import("./getPeering").getPeeringOutput = null as any;
utilities.lazyLoad(exports, ["getPeering","getPeeringOutput"], () => require("./getPeering"));

export { GetPeeringIamPolicyArgs, GetPeeringIamPolicyResult, GetPeeringIamPolicyOutputArgs } from "./getPeeringIamPolicy";
export const getPeeringIamPolicy: typeof import("./getPeeringIamPolicy").getPeeringIamPolicy = null as any;
export const getPeeringIamPolicyOutput: typeof import("./getPeeringIamPolicy").getPeeringIamPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getPeeringIamPolicy","getPeeringIamPolicyOutput"], () => require("./getPeeringIamPolicy"));

export { PeeringArgs } from "./peering";
export type Peering = import("./peering").Peering;
export const Peering: typeof import("./peering").Peering = null as any;
utilities.lazyLoad(exports, ["Peering"], () => require("./peering"));

export { PeeringIamBindingArgs } from "./peeringIamBinding";
export type PeeringIamBinding = import("./peeringIamBinding").PeeringIamBinding;
export const PeeringIamBinding: typeof import("./peeringIamBinding").PeeringIamBinding = null as any;
utilities.lazyLoad(exports, ["PeeringIamBinding"], () => require("./peeringIamBinding"));

export { PeeringIamMemberArgs } from "./peeringIamMember";
export type PeeringIamMember = import("./peeringIamMember").PeeringIamMember;
export const PeeringIamMember: typeof import("./peeringIamMember").PeeringIamMember = null as any;
utilities.lazyLoad(exports, ["PeeringIamMember"], () => require("./peeringIamMember"));

export { PeeringIamPolicyArgs } from "./peeringIamPolicy";
export type PeeringIamPolicy = import("./peeringIamPolicy").PeeringIamPolicy;
export const PeeringIamPolicy: typeof import("./peeringIamPolicy").PeeringIamPolicy = null as any;
utilities.lazyLoad(exports, ["PeeringIamPolicy"], () => require("./peeringIamPolicy"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:managedidentities/v1alpha1:Backup":
                return new Backup(name, <any>undefined, { urn })
            case "google-native:managedidentities/v1alpha1:Domain":
                return new Domain(name, <any>undefined, { urn })
            case "google-native:managedidentities/v1alpha1:DomainBackupIamBinding":
                return new DomainBackupIamBinding(name, <any>undefined, { urn })
            case "google-native:managedidentities/v1alpha1:DomainBackupIamMember":
                return new DomainBackupIamMember(name, <any>undefined, { urn })
            case "google-native:managedidentities/v1alpha1:DomainBackupIamPolicy":
                return new DomainBackupIamPolicy(name, <any>undefined, { urn })
            case "google-native:managedidentities/v1alpha1:DomainIamBinding":
                return new DomainIamBinding(name, <any>undefined, { urn })
            case "google-native:managedidentities/v1alpha1:DomainIamMember":
                return new DomainIamMember(name, <any>undefined, { urn })
            case "google-native:managedidentities/v1alpha1:DomainIamPolicy":
                return new DomainIamPolicy(name, <any>undefined, { urn })
            case "google-native:managedidentities/v1alpha1:Peering":
                return new Peering(name, <any>undefined, { urn })
            case "google-native:managedidentities/v1alpha1:PeeringIamBinding":
                return new PeeringIamBinding(name, <any>undefined, { urn })
            case "google-native:managedidentities/v1alpha1:PeeringIamMember":
                return new PeeringIamMember(name, <any>undefined, { urn })
            case "google-native:managedidentities/v1alpha1:PeeringIamPolicy":
                return new PeeringIamPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "managedidentities/v1alpha1", _module)
