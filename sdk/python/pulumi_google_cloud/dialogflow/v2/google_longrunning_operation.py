# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ._inputs import *

__all__ = ['GoogleLongrunningOperation']


class GoogleLongrunningOperation(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content_uri: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enable_auto_reload: Optional[pulumi.Input[bool]] = None,
                 knowledge_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 latest_reload_status: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowV2DocumentReloadStatusArgs']]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 mime_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent: Optional[pulumi.Input[str]] = None,
                 raw_content: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a new document. Operation

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content_uri: The URI where the file content is located. For documents stored in Google Cloud Storage, these URIs must have the form `gs:///`. NOTE: External URLs must correspond to public webpages, i.e., they must be indexed by Google Search. In particular, URLs for showing documents in Google Cloud Storage (i.e. the URL in your browser) are not supported. Instead use the `gs://` format URI described above.
        :param pulumi.Input[str] display_name: Required. The display name of the document. The name must be 1024 bytes or less; otherwise, the creation request fails.
        :param pulumi.Input[bool] enable_auto_reload: Optional. If true, we try to automatically reload the document every day (at a time picked by the system). If false or unspecified, we don't try to automatically reload the document. Currently you can only enable automatic reload for documents sourced from a public url, see `source` field for the source types. Reload status can be tracked in `latest_reload_status`. If a reload fails, we will keep the document unchanged. If a reload fails with internal errors, the system will try to reload the document on the next day. If a reload fails with non-retriable errors (e.g. PERMISION_DENIED), the system will not try to reload the document anymore. You need to manually reload the document successfully by calling `ReloadDocument` and clear the errors.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] knowledge_types: Required. The knowledge type of document content.
        :param pulumi.Input[pulumi.InputType['GoogleCloudDialogflowV2DocumentReloadStatusArgs']] latest_reload_status: Output only. The time and status of the latest reload. This reload may have been triggered automatically or manually and may not have succeeded.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Optional. Metadata for the document. The metadata supports arbitrary key-value pairs. Suggested use cases include storing a document's title, an external URL distinct from the document's content_uri, etc. The max size of a `key` or a `value` of the metadata is 1024 bytes.
        :param pulumi.Input[str] mime_type: Required. The MIME type of this document.
        :param pulumi.Input[str] name: Optional. The document resource name. The name must be empty when creating a document. Format: `projects//locations//knowledgeBases//documents/`.
        :param pulumi.Input[str] parent: Required. The knowledge base to create a document for. Format: `projects//locations//knowledgeBases/`.
        :param pulumi.Input[str] raw_content: The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['content_uri'] = content_uri
            __props__['display_name'] = display_name
            __props__['enable_auto_reload'] = enable_auto_reload
            __props__['knowledge_types'] = knowledge_types
            __props__['latest_reload_status'] = latest_reload_status
            __props__['metadata'] = metadata
            __props__['mime_type'] = mime_type
            __props__['name'] = name
            if parent is None and not opts.urn:
                raise TypeError("Missing required property 'parent'")
            __props__['parent'] = parent
            __props__['raw_content'] = raw_content
        super(GoogleLongrunningOperation, __self__).__init__(
            'google-cloud:dialogflow/v2:GoogleLongrunningOperation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'GoogleLongrunningOperation':
        """
        Get an existing GoogleLongrunningOperation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return GoogleLongrunningOperation(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
