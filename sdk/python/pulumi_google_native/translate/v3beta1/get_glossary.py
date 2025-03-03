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

__all__ = [
    'GetGlossaryResult',
    'AwaitableGetGlossaryResult',
    'get_glossary',
    'get_glossary_output',
]

@pulumi.output_type
class GetGlossaryResult:
    def __init__(__self__, end_time=None, entry_count=None, input_config=None, language_codes_set=None, language_pair=None, name=None, submit_time=None):
        if end_time and not isinstance(end_time, str):
            raise TypeError("Expected argument 'end_time' to be a str")
        pulumi.set(__self__, "end_time", end_time)
        if entry_count and not isinstance(entry_count, int):
            raise TypeError("Expected argument 'entry_count' to be a int")
        pulumi.set(__self__, "entry_count", entry_count)
        if input_config and not isinstance(input_config, dict):
            raise TypeError("Expected argument 'input_config' to be a dict")
        pulumi.set(__self__, "input_config", input_config)
        if language_codes_set and not isinstance(language_codes_set, dict):
            raise TypeError("Expected argument 'language_codes_set' to be a dict")
        pulumi.set(__self__, "language_codes_set", language_codes_set)
        if language_pair and not isinstance(language_pair, dict):
            raise TypeError("Expected argument 'language_pair' to be a dict")
        pulumi.set(__self__, "language_pair", language_pair)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if submit_time and not isinstance(submit_time, str):
            raise TypeError("Expected argument 'submit_time' to be a str")
        pulumi.set(__self__, "submit_time", submit_time)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> str:
        """
        When the glossary creation was finished.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter(name="entryCount")
    def entry_count(self) -> int:
        """
        The number of entries defined in the glossary.
        """
        return pulumi.get(self, "entry_count")

    @property
    @pulumi.getter(name="inputConfig")
    def input_config(self) -> 'outputs.GlossaryInputConfigResponse':
        """
        Provides examples to build the glossary from. Total glossary must not exceed 10M Unicode codepoints.
        """
        return pulumi.get(self, "input_config")

    @property
    @pulumi.getter(name="languageCodesSet")
    def language_codes_set(self) -> 'outputs.LanguageCodesSetResponse':
        """
        Used with equivalent term set glossaries.
        """
        return pulumi.get(self, "language_codes_set")

    @property
    @pulumi.getter(name="languagePair")
    def language_pair(self) -> 'outputs.LanguageCodePairResponse':
        """
        Used with unidirectional glossaries.
        """
        return pulumi.get(self, "language_pair")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the glossary. Glossary names have the form `projects/{project-number-or-id}/locations/{location-id}/glossaries/{glossary-id}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="submitTime")
    def submit_time(self) -> str:
        """
        When CreateGlossary was called.
        """
        return pulumi.get(self, "submit_time")


class AwaitableGetGlossaryResult(GetGlossaryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGlossaryResult(
            end_time=self.end_time,
            entry_count=self.entry_count,
            input_config=self.input_config,
            language_codes_set=self.language_codes_set,
            language_pair=self.language_pair,
            name=self.name,
            submit_time=self.submit_time)


def get_glossary(glossary_id: Optional[str] = None,
                 location: Optional[str] = None,
                 project: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGlossaryResult:
    """
    Gets a glossary. Returns NOT_FOUND, if the glossary doesn't exist.
    """
    __args__ = dict()
    __args__['glossaryId'] = glossary_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:translate/v3beta1:getGlossary', __args__, opts=opts, typ=GetGlossaryResult).value

    return AwaitableGetGlossaryResult(
        end_time=__ret__.end_time,
        entry_count=__ret__.entry_count,
        input_config=__ret__.input_config,
        language_codes_set=__ret__.language_codes_set,
        language_pair=__ret__.language_pair,
        name=__ret__.name,
        submit_time=__ret__.submit_time)


@_utilities.lift_output_func(get_glossary)
def get_glossary_output(glossary_id: Optional[pulumi.Input[str]] = None,
                        location: Optional[pulumi.Input[str]] = None,
                        project: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGlossaryResult]:
    """
    Gets a glossary. Returns NOT_FOUND, if the glossary doesn't exist.
    """
    ...
