// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a network endpoint group in the specified project using the parameters that are included in the request.
 */
export class RegionNetworkEndpointGroup extends pulumi.CustomResource {
    /**
     * Get an existing RegionNetworkEndpointGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegionNetworkEndpointGroup {
        return new RegionNetworkEndpointGroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:RegionNetworkEndpointGroup';

    /**
     * Returns true if the given object is an instance of RegionNetworkEndpointGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionNetworkEndpointGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionNetworkEndpointGroup.__pulumiType;
    }

    /**
     * Metadata defined as annotations on the network endpoint group.
     */
    public readonly annotations!: pulumi.Output<{[key: string]: string}>;
    /**
     * Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
     */
    public readonly appEngine!: pulumi.Output<outputs.compute.alpha.NetworkEndpointGroupAppEngineResponse>;
    /**
     * Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
     */
    public readonly cloudFunction!: pulumi.Output<outputs.compute.alpha.NetworkEndpointGroupCloudFunctionResponse>;
    /**
     * Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
     */
    public readonly cloudRun!: pulumi.Output<outputs.compute.alpha.NetworkEndpointGroupCloudRunResponse>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * The default port used if the port number is not specified in the network endpoint.
     */
    public readonly defaultPort!: pulumi.Output<number>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Type of the resource. Always compute#networkEndpointGroup for network endpoint group.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * This field is only valid when the network endpoint group is used for load balancing. [Deprecated] This field is deprecated.
     *
     * @deprecated This field is only valid when the network endpoint group is used for load balancing. [Deprecated] This field is deprecated.
     */
    public readonly loadBalancer!: pulumi.Output<outputs.compute.alpha.NetworkEndpointGroupLbNetworkEndpointGroupResponse>;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP, GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, SERVERLESS, PRIVATE_SERVICE_CONNECT.
     */
    public readonly networkEndpointType!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    public readonly pscData!: pulumi.Output<outputs.compute.alpha.NetworkEndpointGroupPscDataResponse>;
    /**
     * The target service url used to set up private service connection to a Google API or a PSC Producer Service Attachment. An example value is: "asia-northeast3-cloudkms.googleapis.com"
     */
    public readonly pscTargetService!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    /**
     * An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
     */
    public readonly requestId!: pulumi.Output<string | undefined>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    public /*out*/ readonly selfLinkWithId!: pulumi.Output<string>;
    /**
     * Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine, cloudFunction or serverlessDeployment may be set.
     */
    public readonly serverlessDeployment!: pulumi.Output<outputs.compute.alpha.NetworkEndpointGroupServerlessDeploymentResponse>;
    /**
     * [Output only] Number of network endpoints in the network endpoint group.
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * Optional URL of the subnetwork to which all network endpoints in the NEG belong.
     */
    public readonly subnetwork!: pulumi.Output<string>;
    /**
     * Specify the type of this network endpoint group. Only LOAD_BALANCING is valid for now.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The URL of the zone where the network endpoint group is located.
     */
    public /*out*/ readonly zone!: pulumi.Output<string>;

    /**
     * Create a RegionNetworkEndpointGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionNetworkEndpointGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["annotations"] = args ? args.annotations : undefined;
            resourceInputs["appEngine"] = args ? args.appEngine : undefined;
            resourceInputs["cloudFunction"] = args ? args.cloudFunction : undefined;
            resourceInputs["cloudRun"] = args ? args.cloudRun : undefined;
            resourceInputs["defaultPort"] = args ? args.defaultPort : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["loadBalancer"] = args ? args.loadBalancer : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["networkEndpointType"] = args ? args.networkEndpointType : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["pscData"] = args ? args.pscData : undefined;
            resourceInputs["pscTargetService"] = args ? args.pscTargetService : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["serverlessDeployment"] = args ? args.serverlessDeployment : undefined;
            resourceInputs["subnetwork"] = args ? args.subnetwork : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["zone"] = undefined /*out*/;
        } else {
            resourceInputs["annotations"] = undefined /*out*/;
            resourceInputs["appEngine"] = undefined /*out*/;
            resourceInputs["cloudFunction"] = undefined /*out*/;
            resourceInputs["cloudRun"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["defaultPort"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["loadBalancer"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["network"] = undefined /*out*/;
            resourceInputs["networkEndpointType"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["pscData"] = undefined /*out*/;
            resourceInputs["pscTargetService"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
            resourceInputs["serverlessDeployment"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["subnetwork"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["zone"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["project", "region"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(RegionNetworkEndpointGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegionNetworkEndpointGroup resource.
 */
export interface RegionNetworkEndpointGroupArgs {
    /**
     * Metadata defined as annotations on the network endpoint group.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
     */
    appEngine?: pulumi.Input<inputs.compute.alpha.NetworkEndpointGroupAppEngineArgs>;
    /**
     * Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
     */
    cloudFunction?: pulumi.Input<inputs.compute.alpha.NetworkEndpointGroupCloudFunctionArgs>;
    /**
     * Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine or cloudFunction may be set.
     */
    cloudRun?: pulumi.Input<inputs.compute.alpha.NetworkEndpointGroupCloudRunArgs>;
    /**
     * The default port used if the port number is not specified in the network endpoint.
     */
    defaultPort?: pulumi.Input<number>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * This field is only valid when the network endpoint group is used for load balancing. [Deprecated] This field is deprecated.
     *
     * @deprecated This field is only valid when the network endpoint group is used for load balancing. [Deprecated] This field is deprecated.
     */
    loadBalancer?: pulumi.Input<inputs.compute.alpha.NetworkEndpointGroupLbNetworkEndpointGroupArgs>;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * The URL of the network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified.
     */
    network?: pulumi.Input<string>;
    /**
     * Type of network endpoints in this network endpoint group. Can be one of GCE_VM_IP, GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_FQDN_PORT, INTERNET_IP_PORT, SERVERLESS, PRIVATE_SERVICE_CONNECT.
     */
    networkEndpointType?: pulumi.Input<enums.compute.alpha.RegionNetworkEndpointGroupNetworkEndpointType>;
    project?: pulumi.Input<string>;
    pscData?: pulumi.Input<inputs.compute.alpha.NetworkEndpointGroupPscDataArgs>;
    /**
     * The target service url used to set up private service connection to a Google API or a PSC Producer Service Attachment. An example value is: "asia-northeast3-cloudkms.googleapis.com"
     */
    pscTargetService?: pulumi.Input<string>;
    region: pulumi.Input<string>;
    /**
     * An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    /**
     * Only valid when networkEndpointType is "SERVERLESS". Only one of cloudRun, appEngine, cloudFunction or serverlessDeployment may be set.
     */
    serverlessDeployment?: pulumi.Input<inputs.compute.alpha.NetworkEndpointGroupServerlessDeploymentArgs>;
    /**
     * Optional URL of the subnetwork to which all network endpoints in the NEG belong.
     */
    subnetwork?: pulumi.Input<string>;
    /**
     * Specify the type of this network endpoint group. Only LOAD_BALANCING is valid for now.
     */
    type?: pulumi.Input<enums.compute.alpha.RegionNetworkEndpointGroupType>;
}
