# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_google_native.integrations.v1alpha as __v1alpha
    v1alpha = __v1alpha
else:
    v1alpha = _utilities.lazy_import('pulumi_google_native.integrations.v1alpha')

