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
    'GetRoutineResult',
    'AwaitableGetRoutineResult',
    'get_routine',
    'get_routine_output',
]

@pulumi.output_type
class GetRoutineResult:
    def __init__(__self__, arguments=None, creation_time=None, definition_body=None, description=None, determinism_level=None, etag=None, imported_libraries=None, language=None, last_modified_time=None, remote_function_options=None, return_table_type=None, return_type=None, routine_reference=None, routine_type=None, spark_options=None, strict_mode=None):
        if arguments and not isinstance(arguments, list):
            raise TypeError("Expected argument 'arguments' to be a list")
        pulumi.set(__self__, "arguments", arguments)
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if definition_body and not isinstance(definition_body, str):
            raise TypeError("Expected argument 'definition_body' to be a str")
        pulumi.set(__self__, "definition_body", definition_body)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if determinism_level and not isinstance(determinism_level, str):
            raise TypeError("Expected argument 'determinism_level' to be a str")
        pulumi.set(__self__, "determinism_level", determinism_level)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if imported_libraries and not isinstance(imported_libraries, list):
            raise TypeError("Expected argument 'imported_libraries' to be a list")
        pulumi.set(__self__, "imported_libraries", imported_libraries)
        if language and not isinstance(language, str):
            raise TypeError("Expected argument 'language' to be a str")
        pulumi.set(__self__, "language", language)
        if last_modified_time and not isinstance(last_modified_time, str):
            raise TypeError("Expected argument 'last_modified_time' to be a str")
        pulumi.set(__self__, "last_modified_time", last_modified_time)
        if remote_function_options and not isinstance(remote_function_options, dict):
            raise TypeError("Expected argument 'remote_function_options' to be a dict")
        pulumi.set(__self__, "remote_function_options", remote_function_options)
        if return_table_type and not isinstance(return_table_type, dict):
            raise TypeError("Expected argument 'return_table_type' to be a dict")
        pulumi.set(__self__, "return_table_type", return_table_type)
        if return_type and not isinstance(return_type, dict):
            raise TypeError("Expected argument 'return_type' to be a dict")
        pulumi.set(__self__, "return_type", return_type)
        if routine_reference and not isinstance(routine_reference, dict):
            raise TypeError("Expected argument 'routine_reference' to be a dict")
        pulumi.set(__self__, "routine_reference", routine_reference)
        if routine_type and not isinstance(routine_type, str):
            raise TypeError("Expected argument 'routine_type' to be a str")
        pulumi.set(__self__, "routine_type", routine_type)
        if spark_options and not isinstance(spark_options, dict):
            raise TypeError("Expected argument 'spark_options' to be a dict")
        pulumi.set(__self__, "spark_options", spark_options)
        if strict_mode and not isinstance(strict_mode, bool):
            raise TypeError("Expected argument 'strict_mode' to be a bool")
        pulumi.set(__self__, "strict_mode", strict_mode)

    @property
    @pulumi.getter
    def arguments(self) -> Sequence['outputs.ArgumentResponse']:
        """
        Optional.
        """
        return pulumi.get(self, "arguments")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        """
        The time when this routine was created, in milliseconds since the epoch.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="definitionBody")
    def definition_body(self) -> str:
        """
        The body of the routine. For functions, this is the expression in the AS clause. If language=SQL, it is the substring inside (but excluding) the parentheses. For example, for the function created with the following statement: `CREATE FUNCTION JoinLines(x string, y string) as (concat(x, "\\n", y))` The definition_body is `concat(x, "\\n", y)` (\\n is not replaced with linebreak). If language=JAVASCRIPT, it is the evaluated string in the AS clause. For example, for the function created with the following statement: `CREATE FUNCTION f() RETURNS STRING LANGUAGE js AS 'return "\\n";\\n'` The definition_body is `return "\\n";\\n` Note that both \\n are replaced with linebreaks.
        """
        return pulumi.get(self, "definition_body")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. The description of the routine, if defined.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="determinismLevel")
    def determinism_level(self) -> str:
        """
        Optional. The determinism level of the JavaScript UDF, if defined.
        """
        return pulumi.get(self, "determinism_level")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A hash of this resource.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="importedLibraries")
    def imported_libraries(self) -> Sequence[str]:
        """
        Optional. If language = "JAVASCRIPT", this field stores the path of the imported JAVASCRIPT libraries.
        """
        return pulumi.get(self, "imported_libraries")

    @property
    @pulumi.getter
    def language(self) -> str:
        """
        Optional. Defaults to "SQL" if remote_function_options field is absent, not set otherwise.
        """
        return pulumi.get(self, "language")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> str:
        """
        The time when this routine was last modified, in milliseconds since the epoch.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter(name="remoteFunctionOptions")
    def remote_function_options(self) -> 'outputs.RemoteFunctionOptionsResponse':
        """
        Optional. Remote function specific options.
        """
        return pulumi.get(self, "remote_function_options")

    @property
    @pulumi.getter(name="returnTableType")
    def return_table_type(self) -> 'outputs.StandardSqlTableTypeResponse':
        """
        Optional. Can be set only if routine_type = "TABLE_VALUED_FUNCTION". If absent, the return table type is inferred from definition_body at query time in each query that references this routine. If present, then the columns in the evaluated table result will be cast to match the column types specified in return table type, at query time.
        """
        return pulumi.get(self, "return_table_type")

    @property
    @pulumi.getter(name="returnType")
    def return_type(self) -> 'outputs.StandardSqlDataTypeResponse':
        """
        Optional if language = "SQL"; required otherwise. Cannot be set if routine_type = "TABLE_VALUED_FUNCTION". If absent, the return type is inferred from definition_body at query time in each query that references this routine. If present, then the evaluated result will be cast to the specified returned type at query time. For example, for the functions created with the following statements: * `CREATE FUNCTION Add(x FLOAT64, y FLOAT64) RETURNS FLOAT64 AS (x + y);` * `CREATE FUNCTION Increment(x FLOAT64) AS (Add(x, 1));` * `CREATE FUNCTION Decrement(x FLOAT64) RETURNS FLOAT64 AS (Add(x, -1));` The return_type is `{type_kind: "FLOAT64"}` for `Add` and `Decrement`, and is absent for `Increment` (inferred as FLOAT64 at query time). Suppose the function `Add` is replaced by `CREATE OR REPLACE FUNCTION Add(x INT64, y INT64) AS (x + y);` Then the inferred return type of `Increment` is automatically changed to INT64 at query time, while the return type of `Decrement` remains FLOAT64.
        """
        return pulumi.get(self, "return_type")

    @property
    @pulumi.getter(name="routineReference")
    def routine_reference(self) -> 'outputs.RoutineReferenceResponse':
        """
        Reference describing the ID of this routine.
        """
        return pulumi.get(self, "routine_reference")

    @property
    @pulumi.getter(name="routineType")
    def routine_type(self) -> str:
        """
        The type of routine.
        """
        return pulumi.get(self, "routine_type")

    @property
    @pulumi.getter(name="sparkOptions")
    def spark_options(self) -> 'outputs.SparkOptionsResponse':
        """
        Optional. Spark specific options.
        """
        return pulumi.get(self, "spark_options")

    @property
    @pulumi.getter(name="strictMode")
    def strict_mode(self) -> bool:
        """
        Optional. Can be set for procedures only. If true (default), the definition body will be validated in the creation and the updates of the procedure. For procedures with an argument of ANY TYPE, the definition body validtion is not supported at creation/update time, and thus this field must be set to false explicitly.
        """
        return pulumi.get(self, "strict_mode")


class AwaitableGetRoutineResult(GetRoutineResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRoutineResult(
            arguments=self.arguments,
            creation_time=self.creation_time,
            definition_body=self.definition_body,
            description=self.description,
            determinism_level=self.determinism_level,
            etag=self.etag,
            imported_libraries=self.imported_libraries,
            language=self.language,
            last_modified_time=self.last_modified_time,
            remote_function_options=self.remote_function_options,
            return_table_type=self.return_table_type,
            return_type=self.return_type,
            routine_reference=self.routine_reference,
            routine_type=self.routine_type,
            spark_options=self.spark_options,
            strict_mode=self.strict_mode)


def get_routine(dataset_id: Optional[str] = None,
                project: Optional[str] = None,
                read_mask: Optional[str] = None,
                routine_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRoutineResult:
    """
    Gets the specified routine resource by routine ID.
    """
    __args__ = dict()
    __args__['datasetId'] = dataset_id
    __args__['project'] = project
    __args__['readMask'] = read_mask
    __args__['routineId'] = routine_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:bigquery/v2:getRoutine', __args__, opts=opts, typ=GetRoutineResult).value

    return AwaitableGetRoutineResult(
        arguments=__ret__.arguments,
        creation_time=__ret__.creation_time,
        definition_body=__ret__.definition_body,
        description=__ret__.description,
        determinism_level=__ret__.determinism_level,
        etag=__ret__.etag,
        imported_libraries=__ret__.imported_libraries,
        language=__ret__.language,
        last_modified_time=__ret__.last_modified_time,
        remote_function_options=__ret__.remote_function_options,
        return_table_type=__ret__.return_table_type,
        return_type=__ret__.return_type,
        routine_reference=__ret__.routine_reference,
        routine_type=__ret__.routine_type,
        spark_options=__ret__.spark_options,
        strict_mode=__ret__.strict_mode)


@_utilities.lift_output_func(get_routine)
def get_routine_output(dataset_id: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       read_mask: Optional[pulumi.Input[Optional[str]]] = None,
                       routine_id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRoutineResult]:
    """
    Gets the specified routine resource by routine ID.
    """
    ...
