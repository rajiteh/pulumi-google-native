# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['KeyArgs', 'Key']

@pulumi.input_type
class KeyArgs:
    def __init__(__self__, *,
                 project: pulumi.Input[str],
                 android_settings: Optional[pulumi.Input['GoogleCloudRecaptchaenterpriseV1AndroidKeySettingsArgs']] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 ios_settings: Optional[pulumi.Input['GoogleCloudRecaptchaenterpriseV1IOSKeySettingsArgs']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 testing_options: Optional[pulumi.Input['GoogleCloudRecaptchaenterpriseV1TestingOptionsArgs']] = None,
                 web_settings: Optional[pulumi.Input['GoogleCloudRecaptchaenterpriseV1WebKeySettingsArgs']] = None):
        """
        The set of arguments for constructing a Key resource.
        :param pulumi.Input['GoogleCloudRecaptchaenterpriseV1AndroidKeySettingsArgs'] android_settings: Settings for keys that can be used by Android apps.
        :param pulumi.Input[str] create_time: The timestamp corresponding to the creation of this Key.
        :param pulumi.Input[str] display_name: Human-readable display name of this key. Modifiable by user.
        :param pulumi.Input['GoogleCloudRecaptchaenterpriseV1IOSKeySettingsArgs'] ios_settings: Settings for keys that can be used by iOS apps.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: See Creating and managing labels.
        :param pulumi.Input[str] name: The resource name for the Key in the format "projects/{project}/keys/{key}".
        :param pulumi.Input['GoogleCloudRecaptchaenterpriseV1TestingOptionsArgs'] testing_options: Options for user acceptance testing.
        :param pulumi.Input['GoogleCloudRecaptchaenterpriseV1WebKeySettingsArgs'] web_settings: Settings for keys that can be used by websites.
        """
        pulumi.set(__self__, "project", project)
        if android_settings is not None:
            pulumi.set(__self__, "android_settings", android_settings)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if ios_settings is not None:
            pulumi.set(__self__, "ios_settings", ios_settings)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if testing_options is not None:
            pulumi.set(__self__, "testing_options", testing_options)
        if web_settings is not None:
            pulumi.set(__self__, "web_settings", web_settings)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="androidSettings")
    def android_settings(self) -> Optional[pulumi.Input['GoogleCloudRecaptchaenterpriseV1AndroidKeySettingsArgs']]:
        """
        Settings for keys that can be used by Android apps.
        """
        return pulumi.get(self, "android_settings")

    @android_settings.setter
    def android_settings(self, value: Optional[pulumi.Input['GoogleCloudRecaptchaenterpriseV1AndroidKeySettingsArgs']]):
        pulumi.set(self, "android_settings", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The timestamp corresponding to the creation of this Key.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Human-readable display name of this key. Modifiable by user.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="iosSettings")
    def ios_settings(self) -> Optional[pulumi.Input['GoogleCloudRecaptchaenterpriseV1IOSKeySettingsArgs']]:
        """
        Settings for keys that can be used by iOS apps.
        """
        return pulumi.get(self, "ios_settings")

    @ios_settings.setter
    def ios_settings(self, value: Optional[pulumi.Input['GoogleCloudRecaptchaenterpriseV1IOSKeySettingsArgs']]):
        pulumi.set(self, "ios_settings", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        See Creating and managing labels.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name for the Key in the format "projects/{project}/keys/{key}".
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="testingOptions")
    def testing_options(self) -> Optional[pulumi.Input['GoogleCloudRecaptchaenterpriseV1TestingOptionsArgs']]:
        """
        Options for user acceptance testing.
        """
        return pulumi.get(self, "testing_options")

    @testing_options.setter
    def testing_options(self, value: Optional[pulumi.Input['GoogleCloudRecaptchaenterpriseV1TestingOptionsArgs']]):
        pulumi.set(self, "testing_options", value)

    @property
    @pulumi.getter(name="webSettings")
    def web_settings(self) -> Optional[pulumi.Input['GoogleCloudRecaptchaenterpriseV1WebKeySettingsArgs']]:
        """
        Settings for keys that can be used by websites.
        """
        return pulumi.get(self, "web_settings")

    @web_settings.setter
    def web_settings(self, value: Optional[pulumi.Input['GoogleCloudRecaptchaenterpriseV1WebKeySettingsArgs']]):
        pulumi.set(self, "web_settings", value)


class Key(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 android_settings: Optional[pulumi.Input[pulumi.InputType['GoogleCloudRecaptchaenterpriseV1AndroidKeySettingsArgs']]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 ios_settings: Optional[pulumi.Input[pulumi.InputType['GoogleCloudRecaptchaenterpriseV1IOSKeySettingsArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 testing_options: Optional[pulumi.Input[pulumi.InputType['GoogleCloudRecaptchaenterpriseV1TestingOptionsArgs']]] = None,
                 web_settings: Optional[pulumi.Input[pulumi.InputType['GoogleCloudRecaptchaenterpriseV1WebKeySettingsArgs']]] = None,
                 __props__=None):
        """
        Creates a new reCAPTCHA Enterprise key.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GoogleCloudRecaptchaenterpriseV1AndroidKeySettingsArgs']] android_settings: Settings for keys that can be used by Android apps.
        :param pulumi.Input[str] create_time: The timestamp corresponding to the creation of this Key.
        :param pulumi.Input[str] display_name: Human-readable display name of this key. Modifiable by user.
        :param pulumi.Input[pulumi.InputType['GoogleCloudRecaptchaenterpriseV1IOSKeySettingsArgs']] ios_settings: Settings for keys that can be used by iOS apps.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: See Creating and managing labels.
        :param pulumi.Input[str] name: The resource name for the Key in the format "projects/{project}/keys/{key}".
        :param pulumi.Input[pulumi.InputType['GoogleCloudRecaptchaenterpriseV1TestingOptionsArgs']] testing_options: Options for user acceptance testing.
        :param pulumi.Input[pulumi.InputType['GoogleCloudRecaptchaenterpriseV1WebKeySettingsArgs']] web_settings: Settings for keys that can be used by websites.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new reCAPTCHA Enterprise key.

        :param str resource_name: The name of the resource.
        :param KeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 android_settings: Optional[pulumi.Input[pulumi.InputType['GoogleCloudRecaptchaenterpriseV1AndroidKeySettingsArgs']]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 ios_settings: Optional[pulumi.Input[pulumi.InputType['GoogleCloudRecaptchaenterpriseV1IOSKeySettingsArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 testing_options: Optional[pulumi.Input[pulumi.InputType['GoogleCloudRecaptchaenterpriseV1TestingOptionsArgs']]] = None,
                 web_settings: Optional[pulumi.Input[pulumi.InputType['GoogleCloudRecaptchaenterpriseV1WebKeySettingsArgs']]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KeyArgs.__new__(KeyArgs)

            __props__.__dict__["android_settings"] = android_settings
            __props__.__dict__["create_time"] = create_time
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["ios_settings"] = ios_settings
            __props__.__dict__["labels"] = labels
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["testing_options"] = testing_options
            __props__.__dict__["web_settings"] = web_settings
        super(Key, __self__).__init__(
            'google-native:recaptchaenterprise/v1:Key',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Key':
        """
        Get an existing Key resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = KeyArgs.__new__(KeyArgs)

        __props__.__dict__["android_settings"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["ios_settings"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["testing_options"] = None
        __props__.__dict__["web_settings"] = None
        return Key(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="androidSettings")
    def android_settings(self) -> pulumi.Output['outputs.GoogleCloudRecaptchaenterpriseV1AndroidKeySettingsResponse']:
        """
        Settings for keys that can be used by Android apps.
        """
        return pulumi.get(self, "android_settings")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The timestamp corresponding to the creation of this Key.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Human-readable display name of this key. Modifiable by user.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="iosSettings")
    def ios_settings(self) -> pulumi.Output['outputs.GoogleCloudRecaptchaenterpriseV1IOSKeySettingsResponse']:
        """
        Settings for keys that can be used by iOS apps.
        """
        return pulumi.get(self, "ios_settings")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        See Creating and managing labels.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name for the Key in the format "projects/{project}/keys/{key}".
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="testingOptions")
    def testing_options(self) -> pulumi.Output['outputs.GoogleCloudRecaptchaenterpriseV1TestingOptionsResponse']:
        """
        Options for user acceptance testing.
        """
        return pulumi.get(self, "testing_options")

    @property
    @pulumi.getter(name="webSettings")
    def web_settings(self) -> pulumi.Output['outputs.GoogleCloudRecaptchaenterpriseV1WebKeySettingsResponse']:
        """
        Settings for keys that can be used by websites.
        """
        return pulumi.get(self, "web_settings")
