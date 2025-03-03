// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns the specified regional BackendService resource.
 */
export function getRegionBackendService(args: GetRegionBackendServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionBackendServiceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/beta:getRegionBackendService", {
        "backendService": args.backendService,
        "project": args.project,
        "region": args.region,
    }, opts);
}

export interface GetRegionBackendServiceArgs {
    backendService: string;
    project?: string;
    region: string;
}

export interface GetRegionBackendServiceResult {
    /**
     * Lifetime of cookies in seconds. This setting is applicable to external and internal HTTP(S) load balancers and Traffic Director and requires GENERATED_COOKIE or HTTP_COOKIE session affinity. If set to 0, the cookie is non-persistent and lasts only until the end of the browser session (or equivalent). The maximum allowed value is two weeks (1,209,600). Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
     */
    readonly affinityCookieTtlSec: number;
    /**
     * The list of backends that serve this BackendService.
     */
    readonly backends: outputs.compute.beta.BackendResponse[];
    /**
     * Cloud CDN configuration for this BackendService. Only available for specified load balancer types.
     */
    readonly cdnPolicy: outputs.compute.beta.BackendServiceCdnPolicyResponse;
    readonly circuitBreakers: outputs.compute.beta.CircuitBreakersResponse;
    /**
     * Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
     */
    readonly compressionMode: string;
    readonly connectionDraining: outputs.compute.beta.ConnectionDrainingResponse;
    /**
     * Connection Tracking configuration for this BackendService. Connection tracking policy settings are only available for Network Load Balancing and Internal TCP/UDP Load Balancing.
     */
    readonly connectionTrackingPolicy: outputs.compute.beta.BackendServiceConnectionTrackingPolicyResponse;
    /**
     * Consistent Hash-based load balancing can be used to provide soft session affinity based on HTTP headers, cookies or other properties. This load balancing policy is applicable only for HTTP connections. The affinity to a particular destination host will be lost when one or more hosts are added/removed from the destination service. This field specifies parameters that control consistent hashing. This field is only applicable when localityLbPolicy is set to MAGLEV or RING_HASH. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. 
     */
    readonly consistentHash: outputs.compute.beta.ConsistentHashLoadBalancerSettingsResponse;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * Headers that the load balancer adds to proxied requests. See [Creating custom headers](https://cloud.google.com/load-balancing/docs/custom-headers).
     */
    readonly customRequestHeaders: string[];
    /**
     * Headers that the load balancer adds to proxied responses. See [Creating custom headers](https://cloud.google.com/load-balancing/docs/custom-headers).
     */
    readonly customResponseHeaders: string[];
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * The resource URL for the edge security policy associated with this backend service.
     */
    readonly edgeSecurityPolicy: string;
    /**
     * If true, enables Cloud CDN for the backend service of an external HTTP(S) load balancer.
     */
    readonly enableCDN: boolean;
    /**
     * Requires at least one backend instance group to be defined as a backup (failover) backend. For load balancers that have configurable failover: [Internal TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/internal/failover-overview) and [external TCP/UDP Load Balancing](https://cloud.google.com/load-balancing/docs/network/networklb-failover-overview).
     */
    readonly failoverPolicy: outputs.compute.beta.BackendServiceFailoverPolicyResponse;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a BackendService. An up-to-date fingerprint must be provided in order to update the BackendService, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a BackendService.
     */
    readonly fingerprint: string;
    /**
     * The list of URLs to the healthChecks, httpHealthChecks (legacy), or httpsHealthChecks (legacy) resource for health checking this backend service. Not all backend services support legacy health checks. See Load balancer guide. Currently, at most one health check can be specified for each backend service. Backend services with instance group or zonal NEG backends must have a health check. Backend services with internet or serverless NEG backends must not have a health check.
     */
    readonly healthChecks: string[];
    /**
     * The configurations for Identity-Aware Proxy on this resource. Not available for Internal TCP/UDP Load Balancing and Network Load Balancing.
     */
    readonly iap: outputs.compute.beta.BackendServiceIAPResponse;
    /**
     * Type of resource. Always compute#backendService for backend services.
     */
    readonly kind: string;
    /**
     * Specifies the load balancer type. A backend service created for one type of load balancer cannot be used with another. For more information, refer to Choosing a load balancer.
     */
    readonly loadBalancingScheme: string;
    /**
     * A list of locality load-balancing policies to be used in order of preference. When you use localityLbPolicies, you must set at least one value for either the localityLbPolicies[].policy or the localityLbPolicies[].customPolicy field. localityLbPolicies overrides any value set in the localityLbPolicy field. For an example of how to use this field, see Define a list of preferred policies. Caution: This field and its children are intended for use in a service mesh that includes gRPC clients only. Envoy proxies can't use backend services that have this configuration.
     */
    readonly localityLbPolicies: outputs.compute.beta.BackendServiceLocalityLoadBalancingPolicyConfigResponse[];
    /**
     * The load balancing algorithm used within the scope of the locality. The possible values are: - ROUND_ROBIN: This is a simple policy in which each healthy backend is selected in round robin order. This is the default. - LEAST_REQUEST: An O(1) algorithm which selects two random healthy hosts and picks the host which has fewer active requests. - RING_HASH: The ring/modulo hash load balancer implements consistent hashing to backends. The algorithm has the property that the addition/removal of a host from a set of N hosts only affects 1/N of the requests. - RANDOM: The load balancer selects a random healthy host. - ORIGINAL_DESTINATION: Backend host is selected based on the client connection metadata, i.e., connections are opened to the same address as the destination address of the incoming connection before the connection was redirected to the load balancer. - MAGLEV: used as a drop in replacement for the ring hash load balancer. Maglev is not as stable as ring hash but has faster table lookup build times and host selection times. For more information about Maglev, see https://ai.google/research/pubs/pub44824 This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. If sessionAffinity is not NONE, and this field is not set to MAGLEV or RING_HASH, session affinity settings will not take effect. Only ROUND_ROBIN and RING_HASH are supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true.
     */
    readonly localityLbPolicy: string;
    /**
     * This field denotes the logging options for the load balancer traffic served by this backend service. If logging is enabled, logs will be exported to Stackdriver.
     */
    readonly logConfig: outputs.compute.beta.BackendServiceLogConfigResponse;
    /**
     * Specifies the default maximum duration (timeout) for streams to this service. Duration is computed from the beginning of the stream until the response has been completely processed, including all retries. A stream that does not complete in this duration is closed. If not specified, there will be no timeout limit, i.e. the maximum duration is infinite. This value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service. This field is only allowed when the loadBalancingScheme of the backend service is INTERNAL_SELF_MANAGED.
     */
    readonly maxStreamDuration: outputs.compute.beta.DurationResponse;
    /**
     * Deployment metadata associated with the resource to be set by a GKE hub controller and read by the backend RCTH
     */
    readonly metadatas: {[key: string]: string};
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * The URL of the network to which this backend service belongs. This field can only be specified when the load balancing scheme is set to INTERNAL.
     */
    readonly network: string;
    /**
     * Settings controlling the eviction of unhealthy hosts from the load balancing pool for the backend service. If not set, this feature is considered disabled. This field is applicable to either: - A regional backend service with the service_protocol set to HTTP, HTTPS, HTTP2, or GRPC, and load_balancing_scheme set to INTERNAL_MANAGED. - A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED. 
     */
    readonly outlierDetection: outputs.compute.beta.OutlierDetectionResponse;
    /**
     * Deprecated in favor of portName. The TCP port to connect on the backend. The default value is 80. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port.
     *
     * @deprecated Deprecated in favor of portName. The TCP port to connect on the backend. The default value is 80. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port.
     */
    readonly port: number;
    /**
     * A named port on a backend instance group representing the port for communication to the backend VMs in that group. The named port must be [defined on each backend instance group](https://cloud.google.com/load-balancing/docs/backend-service#named_ports). This parameter has no meaning if the backends are NEGs. For Internal TCP/UDP Load Balancing and Network Load Balancing, omit port_name.
     */
    readonly portName: string;
    /**
     * The protocol this BackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, TCP, SSL, UDP or GRPC. depending on the chosen load balancer or Traffic Director configuration. Refer to the documentation for the load balancers or for Traffic Director for more information. Must be set to GRPC when the backend service is referenced by a URL map that is bound to target gRPC proxy.
     */
    readonly protocol: string;
    /**
     * URL of the region where the regional backend service resides. This field is not applicable to global backend services. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly region: string;
    /**
     * The resource URL for the security policy associated with this backend service.
     */
    readonly securityPolicy: string;
    /**
     * This field specifies the security settings that apply to this backend service. This field is applicable to a global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
     */
    readonly securitySettings: outputs.compute.beta.SecuritySettingsResponse;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * URLs of networkservices.ServiceBinding resources. Can only be set if load balancing scheme is INTERNAL_SELF_MANAGED. If set, lists of backends and health checks must be both empty.
     */
    readonly serviceBindings: string[];
    /**
     * Type of session affinity to use. The default is NONE. Only NONE and HEADER_FIELD are supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. For more details, see: [Session Affinity](https://cloud.google.com/load-balancing/docs/backend-service#session_affinity).
     */
    readonly sessionAffinity: string;
    readonly subsetting: outputs.compute.beta.SubsettingResponse;
    /**
     * The backend service timeout has a different meaning depending on the type of load balancer. For more information see, Backend service settings. The default is 30 seconds. The full range of timeout values allowed goes from 1 through 2,147,483,647 seconds. This value can be overridden in the PathMatcher configuration of the UrlMap that references this backend service. Not supported when the backend service is referenced by a URL map that is bound to target gRPC proxy that has validateForProxyless field set to true. Instead, use maxStreamDuration.
     */
    readonly timeoutSec: number;
}
/**
 * Returns the specified regional BackendService resource.
 */
export function getRegionBackendServiceOutput(args: GetRegionBackendServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegionBackendServiceResult> {
    return pulumi.output(args).apply((a: any) => getRegionBackendService(a, opts))
}

export interface GetRegionBackendServiceOutputArgs {
    backendService: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
}
