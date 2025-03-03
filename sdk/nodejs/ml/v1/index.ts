// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetJobArgs, GetJobResult, GetJobOutputArgs } from "./getJob";
export const getJob: typeof import("./getJob").getJob = null as any;
export const getJobOutput: typeof import("./getJob").getJobOutput = null as any;
utilities.lazyLoad(exports, ["getJob","getJobOutput"], () => require("./getJob"));

export { GetJobIamPolicyArgs, GetJobIamPolicyResult, GetJobIamPolicyOutputArgs } from "./getJobIamPolicy";
export const getJobIamPolicy: typeof import("./getJobIamPolicy").getJobIamPolicy = null as any;
export const getJobIamPolicyOutput: typeof import("./getJobIamPolicy").getJobIamPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getJobIamPolicy","getJobIamPolicyOutput"], () => require("./getJobIamPolicy"));

export { GetModelArgs, GetModelResult, GetModelOutputArgs } from "./getModel";
export const getModel: typeof import("./getModel").getModel = null as any;
export const getModelOutput: typeof import("./getModel").getModelOutput = null as any;
utilities.lazyLoad(exports, ["getModel","getModelOutput"], () => require("./getModel"));

export { GetModelIamPolicyArgs, GetModelIamPolicyResult, GetModelIamPolicyOutputArgs } from "./getModelIamPolicy";
export const getModelIamPolicy: typeof import("./getModelIamPolicy").getModelIamPolicy = null as any;
export const getModelIamPolicyOutput: typeof import("./getModelIamPolicy").getModelIamPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getModelIamPolicy","getModelIamPolicyOutput"], () => require("./getModelIamPolicy"));

export { GetStudyArgs, GetStudyResult, GetStudyOutputArgs } from "./getStudy";
export const getStudy: typeof import("./getStudy").getStudy = null as any;
export const getStudyOutput: typeof import("./getStudy").getStudyOutput = null as any;
utilities.lazyLoad(exports, ["getStudy","getStudyOutput"], () => require("./getStudy"));

export { GetTrialArgs, GetTrialResult, GetTrialOutputArgs } from "./getTrial";
export const getTrial: typeof import("./getTrial").getTrial = null as any;
export const getTrialOutput: typeof import("./getTrial").getTrialOutput = null as any;
utilities.lazyLoad(exports, ["getTrial","getTrialOutput"], () => require("./getTrial"));

export { GetVersionArgs, GetVersionResult, GetVersionOutputArgs } from "./getVersion";
export const getVersion: typeof import("./getVersion").getVersion = null as any;
export const getVersionOutput: typeof import("./getVersion").getVersionOutput = null as any;
utilities.lazyLoad(exports, ["getVersion","getVersionOutput"], () => require("./getVersion"));

export { JobArgs } from "./job";
export type Job = import("./job").Job;
export const Job: typeof import("./job").Job = null as any;
utilities.lazyLoad(exports, ["Job"], () => require("./job"));

export { JobIamBindingArgs } from "./jobIamBinding";
export type JobIamBinding = import("./jobIamBinding").JobIamBinding;
export const JobIamBinding: typeof import("./jobIamBinding").JobIamBinding = null as any;
utilities.lazyLoad(exports, ["JobIamBinding"], () => require("./jobIamBinding"));

export { JobIamMemberArgs } from "./jobIamMember";
export type JobIamMember = import("./jobIamMember").JobIamMember;
export const JobIamMember: typeof import("./jobIamMember").JobIamMember = null as any;
utilities.lazyLoad(exports, ["JobIamMember"], () => require("./jobIamMember"));

export { JobIamPolicyArgs } from "./jobIamPolicy";
export type JobIamPolicy = import("./jobIamPolicy").JobIamPolicy;
export const JobIamPolicy: typeof import("./jobIamPolicy").JobIamPolicy = null as any;
utilities.lazyLoad(exports, ["JobIamPolicy"], () => require("./jobIamPolicy"));

export { ModelArgs } from "./model";
export type Model = import("./model").Model;
export const Model: typeof import("./model").Model = null as any;
utilities.lazyLoad(exports, ["Model"], () => require("./model"));

export { ModelIamBindingArgs } from "./modelIamBinding";
export type ModelIamBinding = import("./modelIamBinding").ModelIamBinding;
export const ModelIamBinding: typeof import("./modelIamBinding").ModelIamBinding = null as any;
utilities.lazyLoad(exports, ["ModelIamBinding"], () => require("./modelIamBinding"));

export { ModelIamMemberArgs } from "./modelIamMember";
export type ModelIamMember = import("./modelIamMember").ModelIamMember;
export const ModelIamMember: typeof import("./modelIamMember").ModelIamMember = null as any;
utilities.lazyLoad(exports, ["ModelIamMember"], () => require("./modelIamMember"));

export { ModelIamPolicyArgs } from "./modelIamPolicy";
export type ModelIamPolicy = import("./modelIamPolicy").ModelIamPolicy;
export const ModelIamPolicy: typeof import("./modelIamPolicy").ModelIamPolicy = null as any;
utilities.lazyLoad(exports, ["ModelIamPolicy"], () => require("./modelIamPolicy"));

export { StudyArgs } from "./study";
export type Study = import("./study").Study;
export const Study: typeof import("./study").Study = null as any;
utilities.lazyLoad(exports, ["Study"], () => require("./study"));

export { TrialArgs } from "./trial";
export type Trial = import("./trial").Trial;
export const Trial: typeof import("./trial").Trial = null as any;
utilities.lazyLoad(exports, ["Trial"], () => require("./trial"));

export { VersionArgs } from "./version";
export type Version = import("./version").Version;
export const Version: typeof import("./version").Version = null as any;
utilities.lazyLoad(exports, ["Version"], () => require("./version"));


// Export enums:
export * from "../../types/enums/ml/v1";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:ml/v1:Job":
                return new Job(name, <any>undefined, { urn })
            case "google-native:ml/v1:JobIamBinding":
                return new JobIamBinding(name, <any>undefined, { urn })
            case "google-native:ml/v1:JobIamMember":
                return new JobIamMember(name, <any>undefined, { urn })
            case "google-native:ml/v1:JobIamPolicy":
                return new JobIamPolicy(name, <any>undefined, { urn })
            case "google-native:ml/v1:Model":
                return new Model(name, <any>undefined, { urn })
            case "google-native:ml/v1:ModelIamBinding":
                return new ModelIamBinding(name, <any>undefined, { urn })
            case "google-native:ml/v1:ModelIamMember":
                return new ModelIamMember(name, <any>undefined, { urn })
            case "google-native:ml/v1:ModelIamPolicy":
                return new ModelIamPolicy(name, <any>undefined, { urn })
            case "google-native:ml/v1:Study":
                return new Study(name, <any>undefined, { urn })
            case "google-native:ml/v1:Trial":
                return new Trial(name, <any>undefined, { urn })
            case "google-native:ml/v1:Version":
                return new Version(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "ml/v1", _module)
