# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = ['DeviceArgs', 'Device']

@pulumi.input_type
class DeviceArgs:
    def __init__(__self__, *,
                 asset_tag: Optional[pulumi.Input[str]] = None,
                 customer: Optional[pulumi.Input[str]] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 last_sync_time: Optional[pulumi.Input[str]] = None,
                 serial_number: Optional[pulumi.Input[str]] = None,
                 wifi_mac_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Device resource.
        :param pulumi.Input[str] asset_tag: Asset tag of the device.
        :param pulumi.Input[str] customer: Optional. [Resource name](https://cloud.google.com/apis/design/resource_names) of the customer. If you're using this API for your own organization, use `customers/my_customer` If you're using this API to manage another organization, use `customers/{customer}`, where customer is the customer to whom the device belongs.
        :param pulumi.Input[str] device_id: Unique identifier for the device.
        :param pulumi.Input[str] last_sync_time: Most recent time when device synced with this service.
        :param pulumi.Input[str] serial_number: Serial Number of device. Example: HT82V1A01076.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] wifi_mac_addresses: WiFi MAC addresses of device.
        """
        if asset_tag is not None:
            pulumi.set(__self__, "asset_tag", asset_tag)
        if customer is not None:
            pulumi.set(__self__, "customer", customer)
        if device_id is not None:
            pulumi.set(__self__, "device_id", device_id)
        if last_sync_time is not None:
            pulumi.set(__self__, "last_sync_time", last_sync_time)
        if serial_number is not None:
            pulumi.set(__self__, "serial_number", serial_number)
        if wifi_mac_addresses is not None:
            pulumi.set(__self__, "wifi_mac_addresses", wifi_mac_addresses)

    @property
    @pulumi.getter(name="assetTag")
    def asset_tag(self) -> Optional[pulumi.Input[str]]:
        """
        Asset tag of the device.
        """
        return pulumi.get(self, "asset_tag")

    @asset_tag.setter
    def asset_tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "asset_tag", value)

    @property
    @pulumi.getter
    def customer(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. [Resource name](https://cloud.google.com/apis/design/resource_names) of the customer. If you're using this API for your own organization, use `customers/my_customer` If you're using this API to manage another organization, use `customers/{customer}`, where customer is the customer to whom the device belongs.
        """
        return pulumi.get(self, "customer")

    @customer.setter
    def customer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "customer", value)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier for the device.
        """
        return pulumi.get(self, "device_id")

    @device_id.setter
    def device_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_id", value)

    @property
    @pulumi.getter(name="lastSyncTime")
    def last_sync_time(self) -> Optional[pulumi.Input[str]]:
        """
        Most recent time when device synced with this service.
        """
        return pulumi.get(self, "last_sync_time")

    @last_sync_time.setter
    def last_sync_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_sync_time", value)

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> Optional[pulumi.Input[str]]:
        """
        Serial Number of device. Example: HT82V1A01076.
        """
        return pulumi.get(self, "serial_number")

    @serial_number.setter
    def serial_number(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "serial_number", value)

    @property
    @pulumi.getter(name="wifiMacAddresses")
    def wifi_mac_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        WiFi MAC addresses of device.
        """
        return pulumi.get(self, "wifi_mac_addresses")

    @wifi_mac_addresses.setter
    def wifi_mac_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "wifi_mac_addresses", value)


class Device(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asset_tag: Optional[pulumi.Input[str]] = None,
                 customer: Optional[pulumi.Input[str]] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 last_sync_time: Optional[pulumi.Input[str]] = None,
                 serial_number: Optional[pulumi.Input[str]] = None,
                 wifi_mac_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Creates a device. Only company-owned device may be created. **Note**: This method is available only to customers who have one of the following SKUs: Enterprise Standard, Enterprise Plus, Enterprise for Education, and Cloud Identity Premium
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] asset_tag: Asset tag of the device.
        :param pulumi.Input[str] customer: Optional. [Resource name](https://cloud.google.com/apis/design/resource_names) of the customer. If you're using this API for your own organization, use `customers/my_customer` If you're using this API to manage another organization, use `customers/{customer}`, where customer is the customer to whom the device belongs.
        :param pulumi.Input[str] device_id: Unique identifier for the device.
        :param pulumi.Input[str] last_sync_time: Most recent time when device synced with this service.
        :param pulumi.Input[str] serial_number: Serial Number of device. Example: HT82V1A01076.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] wifi_mac_addresses: WiFi MAC addresses of device.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DeviceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a device. Only company-owned device may be created. **Note**: This method is available only to customers who have one of the following SKUs: Enterprise Standard, Enterprise Plus, Enterprise for Education, and Cloud Identity Premium
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param DeviceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeviceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asset_tag: Optional[pulumi.Input[str]] = None,
                 customer: Optional[pulumi.Input[str]] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 last_sync_time: Optional[pulumi.Input[str]] = None,
                 serial_number: Optional[pulumi.Input[str]] = None,
                 wifi_mac_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeviceArgs.__new__(DeviceArgs)

            __props__.__dict__["asset_tag"] = asset_tag
            __props__.__dict__["customer"] = customer
            __props__.__dict__["device_id"] = device_id
            __props__.__dict__["last_sync_time"] = last_sync_time
            __props__.__dict__["serial_number"] = serial_number
            __props__.__dict__["wifi_mac_addresses"] = wifi_mac_addresses
            __props__.__dict__["android_specific_attributes"] = None
            __props__.__dict__["baseband_version"] = None
            __props__.__dict__["bootloader_version"] = None
            __props__.__dict__["brand"] = None
            __props__.__dict__["build_number"] = None
            __props__.__dict__["compromised_state"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["device_type"] = None
            __props__.__dict__["enabled_developer_options"] = None
            __props__.__dict__["enabled_usb_debugging"] = None
            __props__.__dict__["encryption_state"] = None
            __props__.__dict__["imei"] = None
            __props__.__dict__["kernel_version"] = None
            __props__.__dict__["management_state"] = None
            __props__.__dict__["manufacturer"] = None
            __props__.__dict__["meid"] = None
            __props__.__dict__["model"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["network_operator"] = None
            __props__.__dict__["os_version"] = None
            __props__.__dict__["other_accounts"] = None
            __props__.__dict__["owner_type"] = None
            __props__.__dict__["release_version"] = None
            __props__.__dict__["security_patch_time"] = None
        super(Device, __self__).__init__(
            'google-native:cloudidentity/v1:Device',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Device':
        """
        Get an existing Device resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DeviceArgs.__new__(DeviceArgs)

        __props__.__dict__["android_specific_attributes"] = None
        __props__.__dict__["asset_tag"] = None
        __props__.__dict__["baseband_version"] = None
        __props__.__dict__["bootloader_version"] = None
        __props__.__dict__["brand"] = None
        __props__.__dict__["build_number"] = None
        __props__.__dict__["compromised_state"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["customer"] = None
        __props__.__dict__["device_id"] = None
        __props__.__dict__["device_type"] = None
        __props__.__dict__["enabled_developer_options"] = None
        __props__.__dict__["enabled_usb_debugging"] = None
        __props__.__dict__["encryption_state"] = None
        __props__.__dict__["imei"] = None
        __props__.__dict__["kernel_version"] = None
        __props__.__dict__["last_sync_time"] = None
        __props__.__dict__["management_state"] = None
        __props__.__dict__["manufacturer"] = None
        __props__.__dict__["meid"] = None
        __props__.__dict__["model"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["network_operator"] = None
        __props__.__dict__["os_version"] = None
        __props__.__dict__["other_accounts"] = None
        __props__.__dict__["owner_type"] = None
        __props__.__dict__["release_version"] = None
        __props__.__dict__["security_patch_time"] = None
        __props__.__dict__["serial_number"] = None
        __props__.__dict__["wifi_mac_addresses"] = None
        return Device(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="androidSpecificAttributes")
    def android_specific_attributes(self) -> pulumi.Output['outputs.GoogleAppsCloudidentityDevicesV1AndroidAttributesResponse']:
        """
        Attributes specific to Android devices.
        """
        return pulumi.get(self, "android_specific_attributes")

    @property
    @pulumi.getter(name="assetTag")
    def asset_tag(self) -> pulumi.Output[str]:
        """
        Asset tag of the device.
        """
        return pulumi.get(self, "asset_tag")

    @property
    @pulumi.getter(name="basebandVersion")
    def baseband_version(self) -> pulumi.Output[str]:
        """
        Baseband version of the device.
        """
        return pulumi.get(self, "baseband_version")

    @property
    @pulumi.getter(name="bootloaderVersion")
    def bootloader_version(self) -> pulumi.Output[str]:
        """
        Device bootloader version. Example: 0.6.7.
        """
        return pulumi.get(self, "bootloader_version")

    @property
    @pulumi.getter
    def brand(self) -> pulumi.Output[str]:
        """
        Device brand. Example: Samsung.
        """
        return pulumi.get(self, "brand")

    @property
    @pulumi.getter(name="buildNumber")
    def build_number(self) -> pulumi.Output[str]:
        """
        Build number of the device.
        """
        return pulumi.get(self, "build_number")

    @property
    @pulumi.getter(name="compromisedState")
    def compromised_state(self) -> pulumi.Output[str]:
        """
        Represents whether the Device is compromised.
        """
        return pulumi.get(self, "compromised_state")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        When the Company-Owned device was imported. This field is empty for BYOD devices.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def customer(self) -> pulumi.Output[Optional[str]]:
        """
        Optional. [Resource name](https://cloud.google.com/apis/design/resource_names) of the customer. If you're using this API for your own organization, use `customers/my_customer` If you're using this API to manage another organization, use `customers/{customer}`, where customer is the customer to whom the device belongs.
        """
        return pulumi.get(self, "customer")

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> pulumi.Output[str]:
        """
        Unique identifier for the device.
        """
        return pulumi.get(self, "device_id")

    @property
    @pulumi.getter(name="deviceType")
    def device_type(self) -> pulumi.Output[str]:
        """
        Type of device.
        """
        return pulumi.get(self, "device_type")

    @property
    @pulumi.getter(name="enabledDeveloperOptions")
    def enabled_developer_options(self) -> pulumi.Output[bool]:
        """
        Whether developer options is enabled on device.
        """
        return pulumi.get(self, "enabled_developer_options")

    @property
    @pulumi.getter(name="enabledUsbDebugging")
    def enabled_usb_debugging(self) -> pulumi.Output[bool]:
        """
        Whether USB debugging is enabled on device.
        """
        return pulumi.get(self, "enabled_usb_debugging")

    @property
    @pulumi.getter(name="encryptionState")
    def encryption_state(self) -> pulumi.Output[str]:
        """
        Device encryption state.
        """
        return pulumi.get(self, "encryption_state")

    @property
    @pulumi.getter
    def imei(self) -> pulumi.Output[str]:
        """
        IMEI number of device if GSM device; empty otherwise.
        """
        return pulumi.get(self, "imei")

    @property
    @pulumi.getter(name="kernelVersion")
    def kernel_version(self) -> pulumi.Output[str]:
        """
        Kernel version of the device.
        """
        return pulumi.get(self, "kernel_version")

    @property
    @pulumi.getter(name="lastSyncTime")
    def last_sync_time(self) -> pulumi.Output[str]:
        """
        Most recent time when device synced with this service.
        """
        return pulumi.get(self, "last_sync_time")

    @property
    @pulumi.getter(name="managementState")
    def management_state(self) -> pulumi.Output[str]:
        """
        Management state of the device
        """
        return pulumi.get(self, "management_state")

    @property
    @pulumi.getter
    def manufacturer(self) -> pulumi.Output[str]:
        """
        Device manufacturer. Example: Motorola.
        """
        return pulumi.get(self, "manufacturer")

    @property
    @pulumi.getter
    def meid(self) -> pulumi.Output[str]:
        """
        MEID number of device if CDMA device; empty otherwise.
        """
        return pulumi.get(self, "meid")

    @property
    @pulumi.getter
    def model(self) -> pulumi.Output[str]:
        """
        Model name of device. Example: Pixel 3.
        """
        return pulumi.get(self, "model")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        [Resource name](https://cloud.google.com/apis/design/resource_names) of the Device in format: `devices/{device}`, where device is the unique id assigned to the Device.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkOperator")
    def network_operator(self) -> pulumi.Output[str]:
        """
        Mobile or network operator of device, if available.
        """
        return pulumi.get(self, "network_operator")

    @property
    @pulumi.getter(name="osVersion")
    def os_version(self) -> pulumi.Output[str]:
        """
        OS version of the device. Example: Android 8.1.0.
        """
        return pulumi.get(self, "os_version")

    @property
    @pulumi.getter(name="otherAccounts")
    def other_accounts(self) -> pulumi.Output[Sequence[str]]:
        """
        Domain name for Google accounts on device. Type for other accounts on device. On Android, will only be populated if |ownership_privilege| is |PROFILE_OWNER| or |DEVICE_OWNER|. Does not include the account signed in to the device policy app if that account's domain has only one account. Examples: "com.example", "xyz.com".
        """
        return pulumi.get(self, "other_accounts")

    @property
    @pulumi.getter(name="ownerType")
    def owner_type(self) -> pulumi.Output[str]:
        """
        Whether the device is owned by the company or an individual
        """
        return pulumi.get(self, "owner_type")

    @property
    @pulumi.getter(name="releaseVersion")
    def release_version(self) -> pulumi.Output[str]:
        """
        OS release version. Example: 6.0.
        """
        return pulumi.get(self, "release_version")

    @property
    @pulumi.getter(name="securityPatchTime")
    def security_patch_time(self) -> pulumi.Output[str]:
        """
        OS security patch update time on device.
        """
        return pulumi.get(self, "security_patch_time")

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> pulumi.Output[str]:
        """
        Serial Number of device. Example: HT82V1A01076.
        """
        return pulumi.get(self, "serial_number")

    @property
    @pulumi.getter(name="wifiMacAddresses")
    def wifi_mac_addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        WiFi MAC addresses of device.
        """
        return pulumi.get(self, "wifi_mac_addresses")

