// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./analysis";
export * from "./conversation";
export * from "./getAnalysis";
export * from "./getConversation";
export * from "./getIssueModel";
export * from "./getPhraseMatcher";
export * from "./issueModel";
export * from "./phraseMatcher";

// Export enums:
export * from "../../types/enums/contactcenterinsights/v1";

// Import resources to register:
import { Analysis } from "./analysis";
import { Conversation } from "./conversation";
import { IssueModel } from "./issueModel";
import { PhraseMatcher } from "./phraseMatcher";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:contactcenterinsights/v1:Analysis":
                return new Analysis(name, <any>undefined, { urn })
            case "google-native:contactcenterinsights/v1:Conversation":
                return new Conversation(name, <any>undefined, { urn })
            case "google-native:contactcenterinsights/v1:IssueModel":
                return new IssueModel(name, <any>undefined, { urn })
            case "google-native:contactcenterinsights/v1:PhraseMatcher":
                return new PhraseMatcher(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "contactcenterinsights/v1", _module)