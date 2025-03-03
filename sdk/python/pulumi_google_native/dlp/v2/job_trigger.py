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
from ._enums import *
from ._inputs import *

__all__ = ['JobTriggerArgs', 'JobTrigger']

@pulumi.input_type
class JobTriggerArgs:
    def __init__(__self__, *,
                 status: pulumi.Input['JobTriggerStatus'],
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 inspect_job: Optional[pulumi.Input['GooglePrivacyDlpV2InspectJobConfigArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 trigger_id: Optional[pulumi.Input[str]] = None,
                 triggers: Optional[pulumi.Input[Sequence[pulumi.Input['GooglePrivacyDlpV2TriggerArgs']]]] = None):
        """
        The set of arguments for constructing a JobTrigger resource.
        :param pulumi.Input['JobTriggerStatus'] status: A status for this trigger.
        :param pulumi.Input[str] description: User provided description (max 256 chars)
        :param pulumi.Input[str] display_name: Display name (max 100 chars)
        :param pulumi.Input['GooglePrivacyDlpV2InspectJobConfigArgs'] inspect_job: For inspect jobs, a snapshot of the configuration.
        :param pulumi.Input[str] location: Deprecated. This field has no effect.
        :param pulumi.Input[str] name: Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
        :param pulumi.Input[str] trigger_id: The trigger id can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
        :param pulumi.Input[Sequence[pulumi.Input['GooglePrivacyDlpV2TriggerArgs']]] triggers: A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
        """
        pulumi.set(__self__, "status", status)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if inspect_job is not None:
            pulumi.set(__self__, "inspect_job", inspect_job)
        if location is not None:
            warnings.warn("""Deprecated. This field has no effect.""", DeprecationWarning)
            pulumi.log.warn("""location is deprecated: Deprecated. This field has no effect.""")
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if trigger_id is not None:
            pulumi.set(__self__, "trigger_id", trigger_id)
        if triggers is not None:
            pulumi.set(__self__, "triggers", triggers)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input['JobTriggerStatus']:
        """
        A status for this trigger.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input['JobTriggerStatus']):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        User provided description (max 256 chars)
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name (max 100 chars)
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="inspectJob")
    def inspect_job(self) -> Optional[pulumi.Input['GooglePrivacyDlpV2InspectJobConfigArgs']]:
        """
        For inspect jobs, a snapshot of the configuration.
        """
        return pulumi.get(self, "inspect_job")

    @inspect_job.setter
    def inspect_job(self, value: Optional[pulumi.Input['GooglePrivacyDlpV2InspectJobConfigArgs']]):
        pulumi.set(self, "inspect_job", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Deprecated. This field has no effect.
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="triggerId")
    def trigger_id(self) -> Optional[pulumi.Input[str]]:
        """
        The trigger id can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
        """
        return pulumi.get(self, "trigger_id")

    @trigger_id.setter
    def trigger_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trigger_id", value)

    @property
    @pulumi.getter
    def triggers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GooglePrivacyDlpV2TriggerArgs']]]]:
        """
        A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
        """
        return pulumi.get(self, "triggers")

    @triggers.setter
    def triggers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GooglePrivacyDlpV2TriggerArgs']]]]):
        pulumi.set(self, "triggers", value)


class JobTrigger(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 inspect_job: Optional[pulumi.Input[pulumi.InputType['GooglePrivacyDlpV2InspectJobConfigArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input['JobTriggerStatus']] = None,
                 trigger_id: Optional[pulumi.Input[str]] = None,
                 triggers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GooglePrivacyDlpV2TriggerArgs']]]]] = None,
                 __props__=None):
        """
        Creates a job trigger to run DLP actions such as scanning storage for sensitive information on a set schedule. See https://cloud.google.com/dlp/docs/creating-job-triggers to learn more.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: User provided description (max 256 chars)
        :param pulumi.Input[str] display_name: Display name (max 100 chars)
        :param pulumi.Input[pulumi.InputType['GooglePrivacyDlpV2InspectJobConfigArgs']] inspect_job: For inspect jobs, a snapshot of the configuration.
        :param pulumi.Input[str] location: Deprecated. This field has no effect.
        :param pulumi.Input[str] name: Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
        :param pulumi.Input['JobTriggerStatus'] status: A status for this trigger.
        :param pulumi.Input[str] trigger_id: The trigger id can contain uppercase and lowercase letters, numbers, and hyphens; that is, it must match the regular expression: `[a-zA-Z\\d-_]+`. The maximum length is 100 characters. Can be empty to allow the system to generate one.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GooglePrivacyDlpV2TriggerArgs']]]] triggers: A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: JobTriggerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a job trigger to run DLP actions such as scanning storage for sensitive information on a set schedule. See https://cloud.google.com/dlp/docs/creating-job-triggers to learn more.

        :param str resource_name: The name of the resource.
        :param JobTriggerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(JobTriggerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 inspect_job: Optional[pulumi.Input[pulumi.InputType['GooglePrivacyDlpV2InspectJobConfigArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input['JobTriggerStatus']] = None,
                 trigger_id: Optional[pulumi.Input[str]] = None,
                 triggers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GooglePrivacyDlpV2TriggerArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = JobTriggerArgs.__new__(JobTriggerArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["inspect_job"] = inspect_job
            if location is not None and not opts.urn:
                warnings.warn("""Deprecated. This field has no effect.""", DeprecationWarning)
                pulumi.log.warn("""location is deprecated: Deprecated. This field has no effect.""")
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            if status is None and not opts.urn:
                raise TypeError("Missing required property 'status'")
            __props__.__dict__["status"] = status
            __props__.__dict__["trigger_id"] = trigger_id
            __props__.__dict__["triggers"] = triggers
            __props__.__dict__["create_time"] = None
            __props__.__dict__["errors"] = None
            __props__.__dict__["last_run_time"] = None
            __props__.__dict__["update_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["location", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(JobTrigger, __self__).__init__(
            'google-native:dlp/v2:JobTrigger',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'JobTrigger':
        """
        Get an existing JobTrigger resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = JobTriggerArgs.__new__(JobTriggerArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["errors"] = None
        __props__.__dict__["inspect_job"] = None
        __props__.__dict__["last_run_time"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["triggers"] = None
        __props__.__dict__["update_time"] = None
        return JobTrigger(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The creation timestamp of a triggeredJob.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        User provided description (max 256 chars)
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Display name (max 100 chars)
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def errors(self) -> pulumi.Output[Sequence['outputs.GooglePrivacyDlpV2ErrorResponse']]:
        """
        A stream of errors encountered when the trigger was activated. Repeated errors may result in the JobTrigger automatically being paused. Will return the last 100 errors. Whenever the JobTrigger is modified this list will be cleared.
        """
        return pulumi.get(self, "errors")

    @property
    @pulumi.getter(name="inspectJob")
    def inspect_job(self) -> pulumi.Output['outputs.GooglePrivacyDlpV2InspectJobConfigResponse']:
        """
        For inspect jobs, a snapshot of the configuration.
        """
        return pulumi.get(self, "inspect_job")

    @property
    @pulumi.getter(name="lastRunTime")
    def last_run_time(self) -> pulumi.Output[str]:
        """
        The timestamp of the last time this trigger executed.
        """
        return pulumi.get(self, "last_run_time")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        A status for this trigger.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def triggers(self) -> pulumi.Output[Sequence['outputs.GooglePrivacyDlpV2TriggerResponse']]:
        """
        A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
        """
        return pulumi.get(self, "triggers")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The last update timestamp of a triggeredJob.
        """
        return pulumi.get(self, "update_time")

