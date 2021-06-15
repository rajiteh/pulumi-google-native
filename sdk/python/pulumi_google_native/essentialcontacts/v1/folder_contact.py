# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['FolderContactArgs', 'FolderContact']

@pulumi.input_type
class FolderContactArgs:
    def __init__(__self__, *,
                 folder_id: pulumi.Input[str],
                 email: Optional[pulumi.Input[str]] = None,
                 language_tag: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_category_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 validate_time: Optional[pulumi.Input[str]] = None,
                 validation_state: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FolderContact resource.
        :param pulumi.Input[str] email: Required. The email address to send notifications to. This does not need to be a Google account.
        :param pulumi.Input[str] language_tag: The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
        :param pulumi.Input[str] name: The identifier for the contact. Format: {resource_type}/{resource_id}/contacts/{contact_id}
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_category_subscriptions: The categories of notifications that the contact will receive communications for.
        :param pulumi.Input[str] validate_time: The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
        :param pulumi.Input[str] validation_state: The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
        """
        pulumi.set(__self__, "folder_id", folder_id)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if language_tag is not None:
            pulumi.set(__self__, "language_tag", language_tag)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if notification_category_subscriptions is not None:
            pulumi.set(__self__, "notification_category_subscriptions", notification_category_subscriptions)
        if validate_time is not None:
            pulumi.set(__self__, "validate_time", validate_time)
        if validation_state is not None:
            pulumi.set(__self__, "validation_state", validation_state)

    @property
    @pulumi.getter(name="folderId")
    def folder_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "folder_id")

    @folder_id.setter
    def folder_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "folder_id", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The email address to send notifications to. This does not need to be a Google account.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="languageTag")
    def language_tag(self) -> Optional[pulumi.Input[str]]:
        """
        The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
        """
        return pulumi.get(self, "language_tag")

    @language_tag.setter
    def language_tag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "language_tag", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier for the contact. Format: {resource_type}/{resource_id}/contacts/{contact_id}
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notificationCategorySubscriptions")
    def notification_category_subscriptions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The categories of notifications that the contact will receive communications for.
        """
        return pulumi.get(self, "notification_category_subscriptions")

    @notification_category_subscriptions.setter
    def notification_category_subscriptions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "notification_category_subscriptions", value)

    @property
    @pulumi.getter(name="validateTime")
    def validate_time(self) -> Optional[pulumi.Input[str]]:
        """
        The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
        """
        return pulumi.get(self, "validate_time")

    @validate_time.setter
    def validate_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "validate_time", value)

    @property
    @pulumi.getter(name="validationState")
    def validation_state(self) -> Optional[pulumi.Input[str]]:
        """
        The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
        """
        return pulumi.get(self, "validation_state")

    @validation_state.setter
    def validation_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "validation_state", value)


class FolderContact(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 language_tag: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_category_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 validate_time: Optional[pulumi.Input[str]] = None,
                 validation_state: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Adds a new contact for a resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] email: Required. The email address to send notifications to. This does not need to be a Google account.
        :param pulumi.Input[str] language_tag: The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
        :param pulumi.Input[str] name: The identifier for the contact. Format: {resource_type}/{resource_id}/contacts/{contact_id}
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notification_category_subscriptions: The categories of notifications that the contact will receive communications for.
        :param pulumi.Input[str] validate_time: The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
        :param pulumi.Input[str] validation_state: The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FolderContactArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Adds a new contact for a resource.

        :param str resource_name: The name of the resource.
        :param FolderContactArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FolderContactArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 folder_id: Optional[pulumi.Input[str]] = None,
                 language_tag: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_category_subscriptions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 validate_time: Optional[pulumi.Input[str]] = None,
                 validation_state: Optional[pulumi.Input[str]] = None,
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
            __props__ = FolderContactArgs.__new__(FolderContactArgs)

            __props__.__dict__["email"] = email
            if folder_id is None and not opts.urn:
                raise TypeError("Missing required property 'folder_id'")
            __props__.__dict__["folder_id"] = folder_id
            __props__.__dict__["language_tag"] = language_tag
            __props__.__dict__["name"] = name
            __props__.__dict__["notification_category_subscriptions"] = notification_category_subscriptions
            __props__.__dict__["validate_time"] = validate_time
            __props__.__dict__["validation_state"] = validation_state
        super(FolderContact, __self__).__init__(
            'google-native:essentialcontacts/v1:FolderContact',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'FolderContact':
        """
        Get an existing FolderContact resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FolderContactArgs.__new__(FolderContactArgs)

        __props__.__dict__["email"] = None
        __props__.__dict__["language_tag"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["notification_category_subscriptions"] = None
        __props__.__dict__["validate_time"] = None
        __props__.__dict__["validation_state"] = None
        return FolderContact(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[str]:
        """
        Required. The email address to send notifications to. This does not need to be a Google account.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="languageTag")
    def language_tag(self) -> pulumi.Output[str]:
        """
        The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
        """
        return pulumi.get(self, "language_tag")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The identifier for the contact. Format: {resource_type}/{resource_id}/contacts/{contact_id}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationCategorySubscriptions")
    def notification_category_subscriptions(self) -> pulumi.Output[Sequence[str]]:
        """
        The categories of notifications that the contact will receive communications for.
        """
        return pulumi.get(self, "notification_category_subscriptions")

    @property
    @pulumi.getter(name="validateTime")
    def validate_time(self) -> pulumi.Output[str]:
        """
        The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
        """
        return pulumi.get(self, "validate_time")

    @property
    @pulumi.getter(name="validationState")
    def validation_state(self) -> pulumi.Output[str]:
        """
        The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
        """
        return pulumi.get(self, "validation_state")
