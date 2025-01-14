// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { ControlArgs } from "./control";
export type Control = import("./control").Control;
export const Control: typeof import("./control").Control = null as any;
utilities.lazyLoad(exports, ["Control"], () => require("./control"));

export { GetControlArgs, GetControlResult, GetControlOutputArgs } from "./getControl";
export const getControl: typeof import("./getControl").getControl = null as any;
export const getControlOutput: typeof import("./getControl").getControlOutput = null as any;
utilities.lazyLoad(exports, ["getControl","getControlOutput"], () => require("./getControl"));

export { GetModelArgs, GetModelResult, GetModelOutputArgs } from "./getModel";
export const getModel: typeof import("./getModel").getModel = null as any;
export const getModelOutput: typeof import("./getModel").getModelOutput = null as any;
utilities.lazyLoad(exports, ["getModel","getModelOutput"], () => require("./getModel"));

export { GetProductArgs, GetProductResult, GetProductOutputArgs } from "./getProduct";
export const getProduct: typeof import("./getProduct").getProduct = null as any;
export const getProductOutput: typeof import("./getProduct").getProductOutput = null as any;
utilities.lazyLoad(exports, ["getProduct","getProductOutput"], () => require("./getProduct"));

export { GetServingConfigArgs, GetServingConfigResult, GetServingConfigOutputArgs } from "./getServingConfig";
export const getServingConfig: typeof import("./getServingConfig").getServingConfig = null as any;
export const getServingConfigOutput: typeof import("./getServingConfig").getServingConfigOutput = null as any;
utilities.lazyLoad(exports, ["getServingConfig","getServingConfigOutput"], () => require("./getServingConfig"));

export { ModelArgs } from "./model";
export type Model = import("./model").Model;
export const Model: typeof import("./model").Model = null as any;
utilities.lazyLoad(exports, ["Model"], () => require("./model"));

export { ProductArgs } from "./product";
export type Product = import("./product").Product;
export const Product: typeof import("./product").Product = null as any;
utilities.lazyLoad(exports, ["Product"], () => require("./product"));

export { ServingConfigArgs } from "./servingConfig";
export type ServingConfig = import("./servingConfig").ServingConfig;
export const ServingConfig: typeof import("./servingConfig").ServingConfig = null as any;
utilities.lazyLoad(exports, ["ServingConfig"], () => require("./servingConfig"));


// Export enums:
export * from "../../types/enums/retail/v2";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:retail/v2:Control":
                return new Control(name, <any>undefined, { urn })
            case "google-native:retail/v2:Model":
                return new Model(name, <any>undefined, { urn })
            case "google-native:retail/v2:Product":
                return new Product(name, <any>undefined, { urn })
            case "google-native:retail/v2:ServingConfig":
                return new ServingConfig(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "retail/v2", _module)
