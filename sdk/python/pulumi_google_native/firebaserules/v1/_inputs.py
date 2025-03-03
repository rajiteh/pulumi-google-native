# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'FileArgs',
    'SourceArgs',
]

@pulumi.input_type
class FileArgs:
    def __init__(__self__, *,
                 content: pulumi.Input[str],
                 name: pulumi.Input[str],
                 fingerprint: Optional[pulumi.Input[str]] = None):
        """
        `File` containing source content.
        :param pulumi.Input[str] content: Textual Content.
        :param pulumi.Input[str] name: File name.
        :param pulumi.Input[str] fingerprint: Fingerprint (e.g. github sha) associated with the `File`.
        """
        pulumi.set(__self__, "content", content)
        pulumi.set(__self__, "name", name)
        if fingerprint is not None:
            pulumi.set(__self__, "fingerprint", fingerprint)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Input[str]:
        """
        Textual Content.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: pulumi.Input[str]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        File name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def fingerprint(self) -> Optional[pulumi.Input[str]]:
        """
        Fingerprint (e.g. github sha) associated with the `File`.
        """
        return pulumi.get(self, "fingerprint")

    @fingerprint.setter
    def fingerprint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fingerprint", value)


@pulumi.input_type
class SourceArgs:
    def __init__(__self__, *,
                 files: pulumi.Input[Sequence[pulumi.Input['FileArgs']]]):
        """
        `Source` is one or more `File` messages comprising a logical set of rules.
        :param pulumi.Input[Sequence[pulumi.Input['FileArgs']]] files: `File` set constituting the `Source` bundle.
        """
        pulumi.set(__self__, "files", files)

    @property
    @pulumi.getter
    def files(self) -> pulumi.Input[Sequence[pulumi.Input['FileArgs']]]:
        """
        `File` set constituting the `Source` bundle.
        """
        return pulumi.get(self, "files")

    @files.setter
    def files(self, value: pulumi.Input[Sequence[pulumi.Input['FileArgs']]]):
        pulumi.set(self, "files", value)


