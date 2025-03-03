// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudIdentity.V1
{
    /// <summary>
    /// Creates a device. Only company-owned device may be created. **Note**: This method is available only to customers who have one of the following SKUs: Enterprise Standard, Enterprise Plus, Enterprise for Education, and Cloud Identity Premium
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:cloudidentity/v1:Device")]
    public partial class Device : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Attributes specific to Android devices.
        /// </summary>
        [Output("androidSpecificAttributes")]
        public Output<Outputs.GoogleAppsCloudidentityDevicesV1AndroidAttributesResponse> AndroidSpecificAttributes { get; private set; } = null!;

        /// <summary>
        /// Asset tag of the device.
        /// </summary>
        [Output("assetTag")]
        public Output<string> AssetTag { get; private set; } = null!;

        /// <summary>
        /// Baseband version of the device.
        /// </summary>
        [Output("basebandVersion")]
        public Output<string> BasebandVersion { get; private set; } = null!;

        /// <summary>
        /// Device bootloader version. Example: 0.6.7.
        /// </summary>
        [Output("bootloaderVersion")]
        public Output<string> BootloaderVersion { get; private set; } = null!;

        /// <summary>
        /// Device brand. Example: Samsung.
        /// </summary>
        [Output("brand")]
        public Output<string> Brand { get; private set; } = null!;

        /// <summary>
        /// Build number of the device.
        /// </summary>
        [Output("buildNumber")]
        public Output<string> BuildNumber { get; private set; } = null!;

        /// <summary>
        /// Represents whether the Device is compromised.
        /// </summary>
        [Output("compromisedState")]
        public Output<string> CompromisedState { get; private set; } = null!;

        /// <summary>
        /// When the Company-Owned device was imported. This field is empty for BYOD devices.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. [Resource name](https://cloud.google.com/apis/design/resource_names) of the customer. If you're using this API for your own organization, use `customers/my_customer` If you're using this API to manage another organization, use `customers/{customer}`, where customer is the customer to whom the device belongs.
        /// </summary>
        [Output("customer")]
        public Output<string?> Customer { get; private set; } = null!;

        /// <summary>
        /// Unique identifier for the device.
        /// </summary>
        [Output("deviceId")]
        public Output<string> DeviceId { get; private set; } = null!;

        /// <summary>
        /// Type of device.
        /// </summary>
        [Output("deviceType")]
        public Output<string> DeviceType { get; private set; } = null!;

        /// <summary>
        /// Whether developer options is enabled on device.
        /// </summary>
        [Output("enabledDeveloperOptions")]
        public Output<bool> EnabledDeveloperOptions { get; private set; } = null!;

        /// <summary>
        /// Whether USB debugging is enabled on device.
        /// </summary>
        [Output("enabledUsbDebugging")]
        public Output<bool> EnabledUsbDebugging { get; private set; } = null!;

        /// <summary>
        /// Device encryption state.
        /// </summary>
        [Output("encryptionState")]
        public Output<string> EncryptionState { get; private set; } = null!;

        /// <summary>
        /// IMEI number of device if GSM device; empty otherwise.
        /// </summary>
        [Output("imei")]
        public Output<string> Imei { get; private set; } = null!;

        /// <summary>
        /// Kernel version of the device.
        /// </summary>
        [Output("kernelVersion")]
        public Output<string> KernelVersion { get; private set; } = null!;

        /// <summary>
        /// Most recent time when device synced with this service.
        /// </summary>
        [Output("lastSyncTime")]
        public Output<string> LastSyncTime { get; private set; } = null!;

        /// <summary>
        /// Management state of the device
        /// </summary>
        [Output("managementState")]
        public Output<string> ManagementState { get; private set; } = null!;

        /// <summary>
        /// Device manufacturer. Example: Motorola.
        /// </summary>
        [Output("manufacturer")]
        public Output<string> Manufacturer { get; private set; } = null!;

        /// <summary>
        /// MEID number of device if CDMA device; empty otherwise.
        /// </summary>
        [Output("meid")]
        public Output<string> Meid { get; private set; } = null!;

        /// <summary>
        /// Model name of device. Example: Pixel 3.
        /// </summary>
        [Output("model")]
        public Output<string> Model { get; private set; } = null!;

        /// <summary>
        /// [Resource name](https://cloud.google.com/apis/design/resource_names) of the Device in format: `devices/{device}`, where device is the unique id assigned to the Device.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Mobile or network operator of device, if available.
        /// </summary>
        [Output("networkOperator")]
        public Output<string> NetworkOperator { get; private set; } = null!;

        /// <summary>
        /// OS version of the device. Example: Android 8.1.0.
        /// </summary>
        [Output("osVersion")]
        public Output<string> OsVersion { get; private set; } = null!;

        /// <summary>
        /// Domain name for Google accounts on device. Type for other accounts on device. On Android, will only be populated if |ownership_privilege| is |PROFILE_OWNER| or |DEVICE_OWNER|. Does not include the account signed in to the device policy app if that account's domain has only one account. Examples: "com.example", "xyz.com".
        /// </summary>
        [Output("otherAccounts")]
        public Output<ImmutableArray<string>> OtherAccounts { get; private set; } = null!;

        /// <summary>
        /// Whether the device is owned by the company or an individual
        /// </summary>
        [Output("ownerType")]
        public Output<string> OwnerType { get; private set; } = null!;

        /// <summary>
        /// OS release version. Example: 6.0.
        /// </summary>
        [Output("releaseVersion")]
        public Output<string> ReleaseVersion { get; private set; } = null!;

        /// <summary>
        /// OS security patch update time on device.
        /// </summary>
        [Output("securityPatchTime")]
        public Output<string> SecurityPatchTime { get; private set; } = null!;

        /// <summary>
        /// Serial Number of device. Example: HT82V1A01076.
        /// </summary>
        [Output("serialNumber")]
        public Output<string> SerialNumber { get; private set; } = null!;

        /// <summary>
        /// WiFi MAC addresses of device.
        /// </summary>
        [Output("wifiMacAddresses")]
        public Output<ImmutableArray<string>> WifiMacAddresses { get; private set; } = null!;


        /// <summary>
        /// Create a Device resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Device(string name, DeviceArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:cloudidentity/v1:Device", name, args ?? new DeviceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Device(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:cloudidentity/v1:Device", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Device resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Device Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Device(name, id, options);
        }
    }

    public sealed class DeviceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Asset tag of the device.
        /// </summary>
        [Input("assetTag")]
        public Input<string>? AssetTag { get; set; }

        /// <summary>
        /// Optional. [Resource name](https://cloud.google.com/apis/design/resource_names) of the customer. If you're using this API for your own organization, use `customers/my_customer` If you're using this API to manage another organization, use `customers/{customer}`, where customer is the customer to whom the device belongs.
        /// </summary>
        [Input("customer")]
        public Input<string>? Customer { get; set; }

        /// <summary>
        /// Unique identifier for the device.
        /// </summary>
        [Input("deviceId")]
        public Input<string>? DeviceId { get; set; }

        /// <summary>
        /// Most recent time when device synced with this service.
        /// </summary>
        [Input("lastSyncTime")]
        public Input<string>? LastSyncTime { get; set; }

        /// <summary>
        /// Serial Number of device. Example: HT82V1A01076.
        /// </summary>
        [Input("serialNumber")]
        public Input<string>? SerialNumber { get; set; }

        [Input("wifiMacAddresses")]
        private InputList<string>? _wifiMacAddresses;

        /// <summary>
        /// WiFi MAC addresses of device.
        /// </summary>
        public InputList<string> WifiMacAddresses
        {
            get => _wifiMacAddresses ?? (_wifiMacAddresses = new InputList<string>());
            set => _wifiMacAddresses = value;
        }

        public DeviceArgs()
        {
        }
        public static new DeviceArgs Empty => new DeviceArgs();
    }
}
