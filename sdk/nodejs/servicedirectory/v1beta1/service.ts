// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a service, and returns the new service.
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:servicedirectory/v1beta1:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * The timestamp when the service was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Endpoints associated with this service. Returned on LookupService.ResolveService. Control plane clients should use RegistrationService.ListEndpoints.
     */
    public /*out*/ readonly endpoints!: pulumi.Output<outputs.servicedirectory.v1beta1.EndpointResponse[]>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Optional. Metadata for the service. This data can be consumed by service clients. Restrictions: * The entire metadata dictionary may contain up to 2000 characters, spread accoss all key-value pairs. Metadata that goes beyond this limit are rejected * Valid metadata keys have two segments: an optional prefix and name, separated by a slash (/). The name segment is required and must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. The prefix is optional. If specified, the prefix must be a DNS subdomain: a series of DNS labels separated by dots (.), not longer than 253 characters in total, followed by a slash (/). Metadata that fails to meet these requirements are rejected Note: This field is equivalent to the `annotations` field in the v1 API. They have the same syntax and read/write to the same location in Service Directory.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string}>;
    /**
     * Immutable. The resource name for the service in the format `projects/*&#47;locations/*&#47;namespaces/*&#47;services/*`.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly namespaceId!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Required. The Resource ID must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly serviceId!: pulumi.Output<string>;
    /**
     * A globally unique identifier (in UUID4 format) for this service.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * The timestamp when the service was last updated. Note: endpoints being created/deleted/updated within the service are not considered service updates for the purpose of this timestamp.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.namespaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'namespaceId'");
            }
            if ((!args || args.serviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceId'");
            }
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespaceId"] = args ? args.namespaceId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["endpoints"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["endpoints"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["metadata"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["namespaceId"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["serviceId"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "namespaceId", "project", "serviceId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Service.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    location?: pulumi.Input<string>;
    /**
     * Optional. Metadata for the service. This data can be consumed by service clients. Restrictions: * The entire metadata dictionary may contain up to 2000 characters, spread accoss all key-value pairs. Metadata that goes beyond this limit are rejected * Valid metadata keys have two segments: an optional prefix and name, separated by a slash (/). The name segment is required and must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between. The prefix is optional. If specified, the prefix must be a DNS subdomain: a series of DNS labels separated by dots (.), not longer than 253 characters in total, followed by a slash (/). Metadata that fails to meet these requirements are rejected Note: This field is equivalent to the `annotations` field in the v1 API. They have the same syntax and read/write to the same location in Service Directory.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Immutable. The resource name for the service in the format `projects/*&#47;locations/*&#47;namespaces/*&#47;services/*`.
     */
    name?: pulumi.Input<string>;
    namespaceId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Required. The Resource ID must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z](?:[-a-z0-9]{0,61}[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    serviceId: pulumi.Input<string>;
}
