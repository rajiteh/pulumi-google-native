# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_google_native.artifactregistry.v1 as __v1
    v1 = __v1
    import pulumi_google_native.artifactregistry.v1beta1 as __v1beta1
    v1beta1 = __v1beta1
    import pulumi_google_native.artifactregistry.v1beta2 as __v1beta2
    v1beta2 = __v1beta2
else:
    v1 = _utilities.lazy_import('pulumi_google_native.artifactregistry.v1')
    v1beta1 = _utilities.lazy_import('pulumi_google_native.artifactregistry.v1beta1')
    v1beta2 = _utilities.lazy_import('pulumi_google_native.artifactregistry.v1beta2')

