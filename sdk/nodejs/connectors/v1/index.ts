// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { ConnectionArgs } from "./connection";
export type Connection = import("./connection").Connection;
export const Connection: typeof import("./connection").Connection = null as any;
utilities.lazyLoad(exports, ["Connection"], () => require("./connection"));

export { ConnectionIamBindingArgs } from "./connectionIamBinding";
export type ConnectionIamBinding = import("./connectionIamBinding").ConnectionIamBinding;
export const ConnectionIamBinding: typeof import("./connectionIamBinding").ConnectionIamBinding = null as any;
utilities.lazyLoad(exports, ["ConnectionIamBinding"], () => require("./connectionIamBinding"));

export { ConnectionIamMemberArgs } from "./connectionIamMember";
export type ConnectionIamMember = import("./connectionIamMember").ConnectionIamMember;
export const ConnectionIamMember: typeof import("./connectionIamMember").ConnectionIamMember = null as any;
utilities.lazyLoad(exports, ["ConnectionIamMember"], () => require("./connectionIamMember"));

export { ConnectionIamPolicyArgs } from "./connectionIamPolicy";
export type ConnectionIamPolicy = import("./connectionIamPolicy").ConnectionIamPolicy;
export const ConnectionIamPolicy: typeof import("./connectionIamPolicy").ConnectionIamPolicy = null as any;
utilities.lazyLoad(exports, ["ConnectionIamPolicy"], () => require("./connectionIamPolicy"));

export { EndpointAttachmentArgs } from "./endpointAttachment";
export type EndpointAttachment = import("./endpointAttachment").EndpointAttachment;
export const EndpointAttachment: typeof import("./endpointAttachment").EndpointAttachment = null as any;
utilities.lazyLoad(exports, ["EndpointAttachment"], () => require("./endpointAttachment"));

export { GetConnectionArgs, GetConnectionResult, GetConnectionOutputArgs } from "./getConnection";
export const getConnection: typeof import("./getConnection").getConnection = null as any;
export const getConnectionOutput: typeof import("./getConnection").getConnectionOutput = null as any;
utilities.lazyLoad(exports, ["getConnection","getConnectionOutput"], () => require("./getConnection"));

export { GetConnectionIamPolicyArgs, GetConnectionIamPolicyResult, GetConnectionIamPolicyOutputArgs } from "./getConnectionIamPolicy";
export const getConnectionIamPolicy: typeof import("./getConnectionIamPolicy").getConnectionIamPolicy = null as any;
export const getConnectionIamPolicyOutput: typeof import("./getConnectionIamPolicy").getConnectionIamPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getConnectionIamPolicy","getConnectionIamPolicyOutput"], () => require("./getConnectionIamPolicy"));

export { GetEndpointAttachmentArgs, GetEndpointAttachmentResult, GetEndpointAttachmentOutputArgs } from "./getEndpointAttachment";
export const getEndpointAttachment: typeof import("./getEndpointAttachment").getEndpointAttachment = null as any;
export const getEndpointAttachmentOutput: typeof import("./getEndpointAttachment").getEndpointAttachmentOutput = null as any;
utilities.lazyLoad(exports, ["getEndpointAttachment","getEndpointAttachmentOutput"], () => require("./getEndpointAttachment"));

export { GetManagedZoneArgs, GetManagedZoneResult, GetManagedZoneOutputArgs } from "./getManagedZone";
export const getManagedZone: typeof import("./getManagedZone").getManagedZone = null as any;
export const getManagedZoneOutput: typeof import("./getManagedZone").getManagedZoneOutput = null as any;
utilities.lazyLoad(exports, ["getManagedZone","getManagedZoneOutput"], () => require("./getManagedZone"));

export { GetProviderIamPolicyArgs, GetProviderIamPolicyResult, GetProviderIamPolicyOutputArgs } from "./getProviderIamPolicy";
export const getProviderIamPolicy: typeof import("./getProviderIamPolicy").getProviderIamPolicy = null as any;
export const getProviderIamPolicyOutput: typeof import("./getProviderIamPolicy").getProviderIamPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getProviderIamPolicy","getProviderIamPolicyOutput"], () => require("./getProviderIamPolicy"));

export { ManagedZoneArgs } from "./managedZone";
export type ManagedZone = import("./managedZone").ManagedZone;
export const ManagedZone: typeof import("./managedZone").ManagedZone = null as any;
utilities.lazyLoad(exports, ["ManagedZone"], () => require("./managedZone"));

export { ProviderIamBindingArgs } from "./providerIamBinding";
export type ProviderIamBinding = import("./providerIamBinding").ProviderIamBinding;
export const ProviderIamBinding: typeof import("./providerIamBinding").ProviderIamBinding = null as any;
utilities.lazyLoad(exports, ["ProviderIamBinding"], () => require("./providerIamBinding"));

export { ProviderIamMemberArgs } from "./providerIamMember";
export type ProviderIamMember = import("./providerIamMember").ProviderIamMember;
export const ProviderIamMember: typeof import("./providerIamMember").ProviderIamMember = null as any;
utilities.lazyLoad(exports, ["ProviderIamMember"], () => require("./providerIamMember"));

export { ProviderIamPolicyArgs } from "./providerIamPolicy";
export type ProviderIamPolicy = import("./providerIamPolicy").ProviderIamPolicy;
export const ProviderIamPolicy: typeof import("./providerIamPolicy").ProviderIamPolicy = null as any;
utilities.lazyLoad(exports, ["ProviderIamPolicy"], () => require("./providerIamPolicy"));


// Export enums:
export * from "../../types/enums/connectors/v1";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:connectors/v1:Connection":
                return new Connection(name, <any>undefined, { urn })
            case "google-native:connectors/v1:ConnectionIamBinding":
                return new ConnectionIamBinding(name, <any>undefined, { urn })
            case "google-native:connectors/v1:ConnectionIamMember":
                return new ConnectionIamMember(name, <any>undefined, { urn })
            case "google-native:connectors/v1:ConnectionIamPolicy":
                return new ConnectionIamPolicy(name, <any>undefined, { urn })
            case "google-native:connectors/v1:EndpointAttachment":
                return new EndpointAttachment(name, <any>undefined, { urn })
            case "google-native:connectors/v1:ManagedZone":
                return new ManagedZone(name, <any>undefined, { urn })
            case "google-native:connectors/v1:ProviderIamBinding":
                return new ProviderIamBinding(name, <any>undefined, { urn })
            case "google-native:connectors/v1:ProviderIamMember":
                return new ProviderIamMember(name, <any>undefined, { urn })
            case "google-native:connectors/v1:ProviderIamPolicy":
                return new ProviderIamPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "connectors/v1", _module)
