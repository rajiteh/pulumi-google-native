// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details for a specific cluster.
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("google-native:container/v1beta1:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	ClusterId string `pulumi:"clusterId"`
	Location  string `pulumi:"location"`
	Project   string `pulumi:"project"`
}

type LookupClusterResult struct {
	// Configurations for the various addons available to run in the cluster.
	AddonsConfig AddonsConfigResponse `pulumi:"addonsConfig"`
	// Configuration controlling RBAC group membership information.
	AuthenticatorGroupsConfig AuthenticatorGroupsConfigResponse `pulumi:"authenticatorGroupsConfig"`
	// Autopilot configuration for the cluster.
	Autopilot AutopilotResponse `pulumi:"autopilot"`
	// Cluster-level autoscaling configuration.
	Autoscaling ClusterAutoscalingResponse `pulumi:"autoscaling"`
	// Configuration for Binary Authorization.
	BinaryAuthorization BinaryAuthorizationResponse `pulumi:"binaryAuthorization"`
	// The IP address range of the container pods in this cluster, in [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) notation (e.g. `10.96.0.0/14`). Leave blank to have one automatically chosen or specify a `/14` block in `10.0.0.0/8`.
	ClusterIpv4Cidr string `pulumi:"clusterIpv4Cidr"`
	// Telemetry integration for the cluster.
	ClusterTelemetry ClusterTelemetryResponse `pulumi:"clusterTelemetry"`
	// Which conditions caused the current cluster state.
	Conditions []StatusConditionResponse `pulumi:"conditions"`
	// Configuration of Confidential Nodes
	ConfidentialNodes ConfidentialNodesResponse `pulumi:"confidentialNodes"`
	// [Output only] The time the cluster was created, in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreateTime string `pulumi:"createTime"`
	// [Output only] The current software version of the master endpoint.
	CurrentMasterVersion string `pulumi:"currentMasterVersion"`
	// [Output only] Deprecated, use [NodePool.version](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1beta1/projects.locations.clusters.nodePools) instead. The current version of the node software components. If they are currently at multiple versions because they're in the process of being upgraded, this reflects the minimum version of all nodes.
	CurrentNodeVersion string `pulumi:"currentNodeVersion"`
	// Configuration of etcd encryption.
	DatabaseEncryption DatabaseEncryptionResponse `pulumi:"databaseEncryption"`
	// The default constraint on the maximum number of pods that can be run simultaneously on a node in the node pool of this cluster. Only honored if cluster created with IP Alias support.
	DefaultMaxPodsConstraint MaxPodsConstraintResponse `pulumi:"defaultMaxPodsConstraint"`
	// An optional description of this cluster.
	Description string `pulumi:"description"`
	// Kubernetes alpha features are enabled on this cluster. This includes alpha API groups (e.g. v1beta1) and features that may not be production ready in the kubernetes version of the master and nodes. The cluster has no SLA for uptime and master/node upgrades are disabled. Alpha enabled clusters are automatically deleted thirty days after creation.
	EnableKubernetesAlpha bool `pulumi:"enableKubernetesAlpha"`
	// [Output only] The IP address of this cluster's master endpoint. The endpoint can be accessed from the internet at `https://username:password@endpoint/`. See the `masterAuth` property of this resource for username and password information.
	Endpoint string `pulumi:"endpoint"`
	// [Output only] The time the cluster will be automatically deleted in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	ExpireTime string `pulumi:"expireTime"`
	// The initial Kubernetes version for this cluster. Valid versions are those found in validMasterVersions returned by getServerConfig. The version can be upgraded over time; such upgrades are reflected in currentMasterVersion and currentNodeVersion. Users may specify either explicit versions offered by Kubernetes Engine or version aliases, which have the following behavior: - "latest": picks the highest valid Kubernetes version - "1.X": picks the highest valid patch+gke.N patch in the 1.X version - "1.X.Y": picks the highest valid gke.N patch in the 1.X.Y version - "1.X.Y-gke.N": picks an explicit Kubernetes version - "","-": picks the default Kubernetes version
	InitialClusterVersion string `pulumi:"initialClusterVersion"`
	// Configuration for cluster IP allocation.
	IpAllocationPolicy IPAllocationPolicyResponse `pulumi:"ipAllocationPolicy"`
	// The fingerprint of the set of labels for this cluster.
	LabelFingerprint string `pulumi:"labelFingerprint"`
	// Configuration for the legacy ABAC authorization mode.
	LegacyAbac LegacyAbacResponse `pulumi:"legacyAbac"`
	// [Output only] The name of the Google Compute Engine [zone](https://cloud.google.com/compute/docs/regions-zones/regions-zones#available) or [region](https://cloud.google.com/compute/docs/regions-zones/regions-zones#available) in which the cluster resides.
	Location string `pulumi:"location"`
	// The list of Google Compute Engine [zones](https://cloud.google.com/compute/docs/zones#available) in which the cluster's nodes should be located. This field provides a default value if [NodePool.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters.nodePools#NodePool.FIELDS.locations) are not specified during node pool creation. Warning: changing cluster locations will update the [NodePool.Locations](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters.nodePools#NodePool.FIELDS.locations) of all node pools and will result in nodes being added and/or removed.
	Locations []string `pulumi:"locations"`
	// The logging service the cluster should use to write logs. Currently available options: * `logging.googleapis.com/kubernetes` - The Cloud Logging service with a Kubernetes-native resource model * `logging.googleapis.com` - The legacy Cloud Logging service (no longer available as of GKE 1.15). * `none` - no logs will be exported from the cluster. If left as an empty string,`logging.googleapis.com/kubernetes` will be used for GKE 1.14+ or `logging.googleapis.com` for earlier versions.
	LoggingService string `pulumi:"loggingService"`
	// Configure the maintenance policy for this cluster.
	MaintenancePolicy MaintenancePolicyResponse `pulumi:"maintenancePolicy"`
	// Configuration for master components.
	Master MasterResponse `pulumi:"master"`
	// The authentication information for accessing the master endpoint. If unspecified, the defaults are used: For clusters before v1.12, if master_auth is unspecified, `username` will be set to "admin", a random password will be generated, and a client certificate will be issued.
	MasterAuth MasterAuthResponse `pulumi:"masterAuth"`
	// The configuration options for master authorized networks feature.
	MasterAuthorizedNetworksConfig MasterAuthorizedNetworksConfigResponse `pulumi:"masterAuthorizedNetworksConfig"`
	// The monitoring service the cluster should use to write metrics. Currently available options: * "monitoring.googleapis.com/kubernetes" - The Cloud Monitoring service with a Kubernetes-native resource model * `monitoring.googleapis.com` - The legacy Cloud Monitoring service (no longer available as of GKE 1.15). * `none` - No metrics will be exported from the cluster. If left as an empty string,`monitoring.googleapis.com/kubernetes` will be used for GKE 1.14+ or `monitoring.googleapis.com` for earlier versions.
	MonitoringService string `pulumi:"monitoringService"`
	// The name of this cluster. The name must be unique within this project and location (e.g. zone or region), and can be up to 40 characters with the following restrictions: * Lowercase letters, numbers, and hyphens only. * Must start with a letter. * Must end with a number or a letter.
	Name string `pulumi:"name"`
	// The name of the Google Compute Engine [network](https://cloud.google.com/compute/docs/networks-and-firewalls#networks) to which the cluster is connected. If left unspecified, the `default` network will be used. On output this shows the network ID instead of the name.
	Network string `pulumi:"network"`
	// Configuration for cluster networking.
	NetworkConfig NetworkConfigResponse `pulumi:"networkConfig"`
	// Configuration options for the NetworkPolicy feature.
	NetworkPolicy NetworkPolicyResponse `pulumi:"networkPolicy"`
	// [Output only] The size of the address space on each node for hosting containers. This is provisioned from within the `container_ipv4_cidr` range. This field will only be set when cluster is in route-based network mode.
	NodeIpv4CidrSize int `pulumi:"nodeIpv4CidrSize"`
	// The node pools associated with this cluster. This field should not be set if "node_config" or "initial_node_count" are specified.
	NodePools []NodePoolResponse `pulumi:"nodePools"`
	// Notification configuration of the cluster.
	NotificationConfig NotificationConfigResponse `pulumi:"notificationConfig"`
	// Configuration for the PodSecurityPolicy feature.
	PodSecurityPolicyConfig PodSecurityPolicyConfigResponse `pulumi:"podSecurityPolicyConfig"`
	// Configuration for private cluster.
	PrivateClusterConfig PrivateClusterConfigResponse `pulumi:"privateClusterConfig"`
	// Release channel configuration.
	ReleaseChannel ReleaseChannelResponse `pulumi:"releaseChannel"`
	// The resource labels for the cluster to use to annotate any related Google Compute Engine resources.
	ResourceLabels map[string]string `pulumi:"resourceLabels"`
	// Configuration for exporting resource usages. Resource usage export is disabled when this config unspecified.
	ResourceUsageExportConfig ResourceUsageExportConfigResponse `pulumi:"resourceUsageExportConfig"`
	// [Output only] Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// [Output only] The IP address range of the Kubernetes services in this cluster, in [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) notation (e.g. `1.2.3.4/29`). Service addresses are typically put in the last `/16` from the container CIDR.
	ServicesIpv4Cidr string `pulumi:"servicesIpv4Cidr"`
	// Shielded Nodes configuration.
	ShieldedNodes ShieldedNodesResponse `pulumi:"shieldedNodes"`
	// [Output only] The current status of this cluster.
	Status string `pulumi:"status"`
	// The name of the Google Compute Engine [subnetwork](https://cloud.google.com/compute/docs/subnetworks) to which the cluster is connected. On output this shows the subnetwork ID instead of the name.
	Subnetwork string `pulumi:"subnetwork"`
	// Configuration for Cloud TPU support;
	TpuConfig TpuConfigResponse `pulumi:"tpuConfig"`
	// [Output only] The IP address range of the Cloud TPUs in this cluster, in [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) notation (e.g. `1.2.3.4/29`).
	TpuIpv4CidrBlock string `pulumi:"tpuIpv4CidrBlock"`
	// Cluster-level Vertical Pod Autoscaling configuration.
	VerticalPodAutoscaling VerticalPodAutoscalingResponse `pulumi:"verticalPodAutoscaling"`
	// Configuration for issuance of mTLS keys and certificates to Kubernetes pods.
	WorkloadCertificates WorkloadCertificatesResponse `pulumi:"workloadCertificates"`
	// Configuration for the use of Kubernetes Service Accounts in GCP IAM policies.
	WorkloadIdentityConfig WorkloadIdentityConfigResponse `pulumi:"workloadIdentityConfig"`
}