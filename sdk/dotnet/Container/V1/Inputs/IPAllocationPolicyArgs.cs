// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Inputs
{

    /// <summary>
    /// Configuration for controlling how IPs are allocated in the cluster.
    /// </summary>
    public sealed class IPAllocationPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This field is deprecated, use cluster_ipv4_cidr_block.
        /// </summary>
        [Input("clusterIpv4Cidr")]
        public Input<string>? ClusterIpv4Cidr { get; set; }

        /// <summary>
        /// The IP address range for the cluster pod IPs. If this field is set, then `cluster.cluster_ipv4_cidr` must be left blank. This field is only applicable when `use_ip_aliases` is true. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. `/14`) to have a range chosen with a specific netmask. Set to a [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) notation (e.g. `10.96.0.0/14`) from the RFC-1918 private networks (e.g. `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`) to pick a specific range to use.
        /// </summary>
        [Input("clusterIpv4CidrBlock")]
        public Input<string>? ClusterIpv4CidrBlock { get; set; }

        /// <summary>
        /// The name of the secondary range to be used for the cluster CIDR block. The secondary range will be used for pod IP addresses. This must be an existing secondary range associated with the cluster subnetwork. This field is only applicable with use_ip_aliases is true and create_subnetwork is false.
        /// </summary>
        [Input("clusterSecondaryRangeName")]
        public Input<string>? ClusterSecondaryRangeName { get; set; }

        /// <summary>
        /// Whether a new subnetwork will be created automatically for the cluster. This field is only applicable when `use_ip_aliases` is true.
        /// </summary>
        [Input("createSubnetwork")]
        public Input<bool>? CreateSubnetwork { get; set; }

        /// <summary>
        /// The ipv6 access type (internal or external) when create_subnetwork is true
        /// </summary>
        [Input("ipv6AccessType")]
        public Input<Pulumi.GoogleNative.Container.V1.IPAllocationPolicyIpv6AccessType>? Ipv6AccessType { get; set; }

        /// <summary>
        /// This field is deprecated, use node_ipv4_cidr_block.
        /// </summary>
        [Input("nodeIpv4Cidr")]
        public Input<string>? NodeIpv4Cidr { get; set; }

        /// <summary>
        /// The IP address range of the instance IPs in this cluster. This is applicable only if `create_subnetwork` is true. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. `/14`) to have a range chosen with a specific netmask. Set to a [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) notation (e.g. `10.96.0.0/14`) from the RFC-1918 private networks (e.g. `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`) to pick a specific range to use.
        /// </summary>
        [Input("nodeIpv4CidrBlock")]
        public Input<string>? NodeIpv4CidrBlock { get; set; }

        /// <summary>
        /// [PRIVATE FIELD] Pod CIDR size overprovisioning config for the cluster. Pod CIDR size per node depends on max_pods_per_node. By default, the value of max_pods_per_node is doubled and then rounded off to next power of 2 to get the size of pod CIDR block per node. Example: max_pods_per_node of 30 would result in 64 IPs (/26). This config can disable the doubling of IPs (we still round off to next power of 2) Example: max_pods_per_node of 30 will result in 32 IPs (/27) when overprovisioning is disabled.
        /// </summary>
        [Input("podCidrOverprovisionConfig")]
        public Input<Inputs.PodCIDROverprovisionConfigArgs>? PodCidrOverprovisionConfig { get; set; }

        /// <summary>
        /// This field is deprecated, use services_ipv4_cidr_block.
        /// </summary>
        [Input("servicesIpv4Cidr")]
        public Input<string>? ServicesIpv4Cidr { get; set; }

        /// <summary>
        /// The IP address range of the services IPs in this cluster. If blank, a range will be automatically chosen with the default size. This field is only applicable when `use_ip_aliases` is true. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. `/14`) to have a range chosen with a specific netmask. Set to a [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) notation (e.g. `10.96.0.0/14`) from the RFC-1918 private networks (e.g. `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`) to pick a specific range to use.
        /// </summary>
        [Input("servicesIpv4CidrBlock")]
        public Input<string>? ServicesIpv4CidrBlock { get; set; }

        /// <summary>
        /// The name of the secondary range to be used as for the services CIDR block. The secondary range will be used for service ClusterIPs. This must be an existing secondary range associated with the cluster subnetwork. This field is only applicable with use_ip_aliases is true and create_subnetwork is false.
        /// </summary>
        [Input("servicesSecondaryRangeName")]
        public Input<string>? ServicesSecondaryRangeName { get; set; }

        /// <summary>
        /// The IP stack type of the cluster
        /// </summary>
        [Input("stackType")]
        public Input<Pulumi.GoogleNative.Container.V1.IPAllocationPolicyStackType>? StackType { get; set; }

        /// <summary>
        /// A custom subnetwork name to be used if `create_subnetwork` is true. If this field is empty, then an automatic name will be chosen for the new subnetwork.
        /// </summary>
        [Input("subnetworkName")]
        public Input<string>? SubnetworkName { get; set; }

        /// <summary>
        /// The IP address range of the Cloud TPUs in this cluster. If unspecified, a range will be automatically chosen with the default size. This field is only applicable when `use_ip_aliases` is true. If unspecified, the range will use the default size. Set to /netmask (e.g. `/14`) to have a range chosen with a specific netmask. Set to a [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) notation (e.g. `10.96.0.0/14`) from the RFC-1918 private networks (e.g. `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`) to pick a specific range to use.
        /// </summary>
        [Input("tpuIpv4CidrBlock")]
        public Input<string>? TpuIpv4CidrBlock { get; set; }

        /// <summary>
        /// Whether alias IPs will be used for pod IPs in the cluster. This is used in conjunction with use_routes. It cannot be true if use_routes is true. If both use_ip_aliases and use_routes are false, then the server picks the default IP allocation mode
        /// </summary>
        [Input("useIpAliases")]
        public Input<bool>? UseIpAliases { get; set; }

        /// <summary>
        /// Whether routes will be used for pod IPs in the cluster. This is used in conjunction with use_ip_aliases. It cannot be true if use_ip_aliases is true. If both use_ip_aliases and use_routes are false, then the server picks the default IP allocation mode
        /// </summary>
        [Input("useRoutes")]
        public Input<bool>? UseRoutes { get; set; }

        public IPAllocationPolicyArgs()
        {
        }
        public static new IPAllocationPolicyArgs Empty => new IPAllocationPolicyArgs();
    }
}
