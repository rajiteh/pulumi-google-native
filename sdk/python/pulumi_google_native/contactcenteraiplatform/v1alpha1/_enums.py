# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'InstanceConfigInstanceSize',
]


class InstanceConfigInstanceSize(str, Enum):
    """
    The instance size of this the instance configuration.
    """
    INSTANCE_SIZE_UNSPECIFIED = "INSTANCE_SIZE_UNSPECIFIED"
    """
    The default value. This value is used if the state is omitted.
    """
    STANDARD_SMALL = "STANDARD_SMALL"
    """
    Instance Size STANDARD_SMALL.
    """
    STANDARD_MEDIUM = "STANDARD_MEDIUM"
    """
    Instance Size STANDARD_MEDIUM.
    """
    STANDARD_LARGE = "STANDARD_LARGE"
    """
    Instance Size STANDARD_LARGE.
    """
    STANDARD_XLARGE = "STANDARD_XLARGE"
    """
    Instance Size STANDARD_XLARGE.
    """
    STANDARD2XLARGE = "STANDARD_2XLARGE"
    """
    Instance Size STANDARD_2XLARGE.
    """
    STANDARD3XLARGE = "STANDARD_3XLARGE"
    """
    Instance Size STANDARD_3XLARGE.
    """
