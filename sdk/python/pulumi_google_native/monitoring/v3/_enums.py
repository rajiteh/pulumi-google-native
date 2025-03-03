# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AggregationCrossSeriesReducer',
    'AggregationPerSeriesAligner',
    'AlertPolicyCombiner',
    'ContentMatcherMatcher',
    'HttpCheckContentType',
    'HttpCheckRequestMethod',
    'InternalCheckerState',
    'JsonPathMatcherJsonMatcher',
    'LabelDescriptorValueType',
    'MetricDescriptorLaunchStage',
    'MetricDescriptorMetadataLaunchStage',
    'MetricDescriptorMetricKind',
    'MetricDescriptorValueType',
    'MetricThresholdComparison',
    'MetricThresholdEvaluationMissingData',
    'MonitoringQueryLanguageConditionEvaluationMissingData',
    'NotificationChannelVerificationStatus',
    'ResourceGroupResourceType',
    'ResponseStatusCodeStatusClass',
    'ServiceLevelObjectiveCalendarPeriod',
    'UptimeCheckConfigCheckerType',
    'UptimeCheckConfigSelectedRegionsItem',
]


class AggregationCrossSeriesReducer(str, Enum):
    """
    The reduction operation to be used to combine time series into a single time series, where the value of each data point in the resulting series is a function of all the already aligned values in the input time series.Not all reducer operations can be applied to all time series. The valid choices depend on the metric_kind and the value_type of the original time series. Reduction can yield a time series with a different metric_kind or value_type than the input time series.Time series data must first be aligned (see per_series_aligner) in order to perform cross-time series reduction. If cross_series_reducer is specified, then per_series_aligner must be specified, and must not be ALIGN_NONE. An alignment_period must also be specified; otherwise, an error is returned.
    """
    REDUCE_NONE = "REDUCE_NONE"
    """
    No cross-time series reduction. The output of the Aligner is returned.
    """
    REDUCE_MEAN = "REDUCE_MEAN"
    """
    Reduce by computing the mean value across time series for each alignment period. This reducer is valid for DELTA and GAUGE metrics with numeric or distribution values. The value_type of the output is DOUBLE.
    """
    REDUCE_MIN = "REDUCE_MIN"
    """
    Reduce by computing the minimum value across time series for each alignment period. This reducer is valid for DELTA and GAUGE metrics with numeric values. The value_type of the output is the same as the value_type of the input.
    """
    REDUCE_MAX = "REDUCE_MAX"
    """
    Reduce by computing the maximum value across time series for each alignment period. This reducer is valid for DELTA and GAUGE metrics with numeric values. The value_type of the output is the same as the value_type of the input.
    """
    REDUCE_SUM = "REDUCE_SUM"
    """
    Reduce by computing the sum across time series for each alignment period. This reducer is valid for DELTA and GAUGE metrics with numeric and distribution values. The value_type of the output is the same as the value_type of the input.
    """
    REDUCE_STDDEV = "REDUCE_STDDEV"
    """
    Reduce by computing the standard deviation across time series for each alignment period. This reducer is valid for DELTA and GAUGE metrics with numeric or distribution values. The value_type of the output is DOUBLE.
    """
    REDUCE_COUNT = "REDUCE_COUNT"
    """
    Reduce by computing the number of data points across time series for each alignment period. This reducer is valid for DELTA and GAUGE metrics of numeric, Boolean, distribution, and string value_type. The value_type of the output is INT64.
    """
    REDUCE_COUNT_TRUE = "REDUCE_COUNT_TRUE"
    """
    Reduce by computing the number of True-valued data points across time series for each alignment period. This reducer is valid for DELTA and GAUGE metrics of Boolean value_type. The value_type of the output is INT64.
    """
    REDUCE_COUNT_FALSE = "REDUCE_COUNT_FALSE"
    """
    Reduce by computing the number of False-valued data points across time series for each alignment period. This reducer is valid for DELTA and GAUGE metrics of Boolean value_type. The value_type of the output is INT64.
    """
    REDUCE_FRACTION_TRUE = "REDUCE_FRACTION_TRUE"
    """
    Reduce by computing the ratio of the number of True-valued data points to the total number of data points for each alignment period. This reducer is valid for DELTA and GAUGE metrics of Boolean value_type. The output value is in the range 0.0, 1.0 and has value_type DOUBLE.
    """
    REDUCE_PERCENTILE99 = "REDUCE_PERCENTILE_99"
    """
    Reduce by computing the 99th percentile (https://en.wikipedia.org/wiki/Percentile) of data points across time series for each alignment period. This reducer is valid for GAUGE and DELTA metrics of numeric and distribution type. The value of the output is DOUBLE.
    """
    REDUCE_PERCENTILE95 = "REDUCE_PERCENTILE_95"
    """
    Reduce by computing the 95th percentile (https://en.wikipedia.org/wiki/Percentile) of data points across time series for each alignment period. This reducer is valid for GAUGE and DELTA metrics of numeric and distribution type. The value of the output is DOUBLE.
    """
    REDUCE_PERCENTILE50 = "REDUCE_PERCENTILE_50"
    """
    Reduce by computing the 50th percentile (https://en.wikipedia.org/wiki/Percentile) of data points across time series for each alignment period. This reducer is valid for GAUGE and DELTA metrics of numeric and distribution type. The value of the output is DOUBLE.
    """
    REDUCE_PERCENTILE05 = "REDUCE_PERCENTILE_05"
    """
    Reduce by computing the 5th percentile (https://en.wikipedia.org/wiki/Percentile) of data points across time series for each alignment period. This reducer is valid for GAUGE and DELTA metrics of numeric and distribution type. The value of the output is DOUBLE.
    """


class AggregationPerSeriesAligner(str, Enum):
    """
    An Aligner describes how to bring the data points in a single time series into temporal alignment. Except for ALIGN_NONE, all alignments cause all the data points in an alignment_period to be mathematically grouped together, resulting in a single data point for each alignment_period with end timestamp at the end of the period.Not all alignment operations may be applied to all time series. The valid choices depend on the metric_kind and value_type of the original time series. Alignment can change the metric_kind or the value_type of the time series.Time series data must be aligned in order to perform cross-time series reduction. If cross_series_reducer is specified, then per_series_aligner must be specified and not equal to ALIGN_NONE and alignment_period must be specified; otherwise, an error is returned.
    """
    ALIGN_NONE = "ALIGN_NONE"
    """
    No alignment. Raw data is returned. Not valid if cross-series reduction is requested. The value_type of the result is the same as the value_type of the input.
    """
    ALIGN_DELTA = "ALIGN_DELTA"
    """
    Align and convert to DELTA. The output is delta = y1 - y0.This alignment is valid for CUMULATIVE and DELTA metrics. If the selected alignment period results in periods with no data, then the aligned value for such a period is created by interpolation. The value_type of the aligned result is the same as the value_type of the input.
    """
    ALIGN_RATE = "ALIGN_RATE"
    """
    Align and convert to a rate. The result is computed as rate = (y1 - y0)/(t1 - t0), or "delta over time". Think of this aligner as providing the slope of the line that passes through the value at the start and at the end of the alignment_period.This aligner is valid for CUMULATIVE and DELTA metrics with numeric values. If the selected alignment period results in periods with no data, then the aligned value for such a period is created by interpolation. The output is a GAUGE metric with value_type DOUBLE.If, by "rate", you mean "percentage change", see the ALIGN_PERCENT_CHANGE aligner instead.
    """
    ALIGN_INTERPOLATE = "ALIGN_INTERPOLATE"
    """
    Align by interpolating between adjacent points around the alignment period boundary. This aligner is valid for GAUGE metrics with numeric values. The value_type of the aligned result is the same as the value_type of the input.
    """
    ALIGN_NEXT_OLDER = "ALIGN_NEXT_OLDER"
    """
    Align by moving the most recent data point before the end of the alignment period to the boundary at the end of the alignment period. This aligner is valid for GAUGE metrics. The value_type of the aligned result is the same as the value_type of the input.
    """
    ALIGN_MIN = "ALIGN_MIN"
    """
    Align the time series by returning the minimum value in each alignment period. This aligner is valid for GAUGE and DELTA metrics with numeric values. The value_type of the aligned result is the same as the value_type of the input.
    """
    ALIGN_MAX = "ALIGN_MAX"
    """
    Align the time series by returning the maximum value in each alignment period. This aligner is valid for GAUGE and DELTA metrics with numeric values. The value_type of the aligned result is the same as the value_type of the input.
    """
    ALIGN_MEAN = "ALIGN_MEAN"
    """
    Align the time series by returning the mean value in each alignment period. This aligner is valid for GAUGE and DELTA metrics with numeric values. The value_type of the aligned result is DOUBLE.
    """
    ALIGN_COUNT = "ALIGN_COUNT"
    """
    Align the time series by returning the number of values in each alignment period. This aligner is valid for GAUGE and DELTA metrics with numeric or Boolean values. The value_type of the aligned result is INT64.
    """
    ALIGN_SUM = "ALIGN_SUM"
    """
    Align the time series by returning the sum of the values in each alignment period. This aligner is valid for GAUGE and DELTA metrics with numeric and distribution values. The value_type of the aligned result is the same as the value_type of the input.
    """
    ALIGN_STDDEV = "ALIGN_STDDEV"
    """
    Align the time series by returning the standard deviation of the values in each alignment period. This aligner is valid for GAUGE and DELTA metrics with numeric values. The value_type of the output is DOUBLE.
    """
    ALIGN_COUNT_TRUE = "ALIGN_COUNT_TRUE"
    """
    Align the time series by returning the number of True values in each alignment period. This aligner is valid for GAUGE metrics with Boolean values. The value_type of the output is INT64.
    """
    ALIGN_COUNT_FALSE = "ALIGN_COUNT_FALSE"
    """
    Align the time series by returning the number of False values in each alignment period. This aligner is valid for GAUGE metrics with Boolean values. The value_type of the output is INT64.
    """
    ALIGN_FRACTION_TRUE = "ALIGN_FRACTION_TRUE"
    """
    Align the time series by returning the ratio of the number of True values to the total number of values in each alignment period. This aligner is valid for GAUGE metrics with Boolean values. The output value is in the range 0.0, 1.0 and has value_type DOUBLE.
    """
    ALIGN_PERCENTILE99 = "ALIGN_PERCENTILE_99"
    """
    Align the time series by using percentile aggregation (https://en.wikipedia.org/wiki/Percentile). The resulting data point in each alignment period is the 99th percentile of all data points in the period. This aligner is valid for GAUGE and DELTA metrics with distribution values. The output is a GAUGE metric with value_type DOUBLE.
    """
    ALIGN_PERCENTILE95 = "ALIGN_PERCENTILE_95"
    """
    Align the time series by using percentile aggregation (https://en.wikipedia.org/wiki/Percentile). The resulting data point in each alignment period is the 95th percentile of all data points in the period. This aligner is valid for GAUGE and DELTA metrics with distribution values. The output is a GAUGE metric with value_type DOUBLE.
    """
    ALIGN_PERCENTILE50 = "ALIGN_PERCENTILE_50"
    """
    Align the time series by using percentile aggregation (https://en.wikipedia.org/wiki/Percentile). The resulting data point in each alignment period is the 50th percentile of all data points in the period. This aligner is valid for GAUGE and DELTA metrics with distribution values. The output is a GAUGE metric with value_type DOUBLE.
    """
    ALIGN_PERCENTILE05 = "ALIGN_PERCENTILE_05"
    """
    Align the time series by using percentile aggregation (https://en.wikipedia.org/wiki/Percentile). The resulting data point in each alignment period is the 5th percentile of all data points in the period. This aligner is valid for GAUGE and DELTA metrics with distribution values. The output is a GAUGE metric with value_type DOUBLE.
    """
    ALIGN_PERCENT_CHANGE = "ALIGN_PERCENT_CHANGE"
    """
    Align and convert to a percentage change. This aligner is valid for GAUGE and DELTA metrics with numeric values. This alignment returns ((current - previous)/previous) * 100, where the value of previous is determined based on the alignment_period.If the values of current and previous are both 0, then the returned value is 0. If only previous is 0, the returned value is infinity.A 10-minute moving mean is computed at each point of the alignment period prior to the above calculation to smooth the metric and prevent false positives from very short-lived spikes. The moving mean is only applicable for data whose values are >= 0. Any values < 0 are treated as a missing datapoint, and are ignored. While DELTA metrics are accepted by this alignment, special care should be taken that the values for the metric will always be positive. The output is a GAUGE metric with value_type DOUBLE.
    """


class AlertPolicyCombiner(str, Enum):
    """
    How to combine the results of multiple conditions to determine if an incident should be opened. If condition_time_series_query_language is present, this must be COMBINE_UNSPECIFIED.
    """
    COMBINE_UNSPECIFIED = "COMBINE_UNSPECIFIED"
    """
    An unspecified combiner.
    """
    AND_ = "AND"
    """
    Combine conditions using the logical AND operator. An incident is created only if all the conditions are met simultaneously. This combiner is satisfied if all conditions are met, even if they are met on completely different resources.
    """
    OR_ = "OR"
    """
    Combine conditions using the logical OR operator. An incident is created if any of the listed conditions is met.
    """
    AND_WITH_MATCHING_RESOURCE = "AND_WITH_MATCHING_RESOURCE"
    """
    Combine conditions using logical AND operator, but unlike the regular AND option, an incident is created only if all conditions are met simultaneously on at least one resource.
    """


class ContentMatcherMatcher(str, Enum):
    """
    The type of content matcher that will be applied to the server output, compared to the content string when the check is run.
    """
    CONTENT_MATCHER_OPTION_UNSPECIFIED = "CONTENT_MATCHER_OPTION_UNSPECIFIED"
    """
    No content matcher type specified (maintained for backward compatibility, but deprecated for future use). Treated as CONTAINS_STRING.
    """
    CONTAINS_STRING = "CONTAINS_STRING"
    """
    Selects substring matching. The match succeeds if the output contains the content string. This is the default value for checks without a matcher option, or where the value of matcher is CONTENT_MATCHER_OPTION_UNSPECIFIED.
    """
    NOT_CONTAINS_STRING = "NOT_CONTAINS_STRING"
    """
    Selects negation of substring matching. The match succeeds if the output does NOT contain the content string.
    """
    MATCHES_REGEX = "MATCHES_REGEX"
    """
    Selects regular-expression matching. The match succeeds if the output matches the regular expression specified in the content string. Regex matching is only supported for HTTP/HTTPS checks.
    """
    NOT_MATCHES_REGEX = "NOT_MATCHES_REGEX"
    """
    Selects negation of regular-expression matching. The match succeeds if the output does NOT match the regular expression specified in the content string. Regex matching is only supported for HTTP/HTTPS checks.
    """
    MATCHES_JSON_PATH = "MATCHES_JSON_PATH"
    """
    Selects JSONPath matching. See JsonPathMatcher for details on when the match succeeds. JSONPath matching is only supported for HTTP/HTTPS checks.
    """
    NOT_MATCHES_JSON_PATH = "NOT_MATCHES_JSON_PATH"
    """
    Selects JSONPath matching. See JsonPathMatcher for details on when the match succeeds. Succeeds when output does NOT match as specified. JSONPath is only supported for HTTP/HTTPS checks.
    """


class HttpCheckContentType(str, Enum):
    """
    The content type header to use for the check. The following configurations result in errors: 1. Content type is specified in both the headers field and the content_type field. 2. Request method is GET and content_type is not TYPE_UNSPECIFIED 3. Request method is POST and content_type is TYPE_UNSPECIFIED. 4. Request method is POST and a "Content-Type" header is provided via headers field. The content_type field should be used instead.
    """
    TYPE_UNSPECIFIED = "TYPE_UNSPECIFIED"
    """
    No content type specified.
    """
    URL_ENCODED = "URL_ENCODED"
    """
    body is in URL-encoded form. Equivalent to setting the Content-Type to application/x-www-form-urlencoded in the HTTP request.
    """
    USER_PROVIDED = "USER_PROVIDED"
    """
    body is in custom_content_type form. Equivalent to setting the Content-Type to the contents of custom_content_type in the HTTP request.
    """


class HttpCheckRequestMethod(str, Enum):
    """
    The HTTP request method to use for the check. If set to METHOD_UNSPECIFIED then request_method defaults to GET.
    """
    METHOD_UNSPECIFIED = "METHOD_UNSPECIFIED"
    """
    No request method specified.
    """
    GET = "GET"
    """
    GET request.
    """
    POST = "POST"
    """
    POST request.
    """


class InternalCheckerState(str, Enum):
    """
    The current operational state of the internal checker.
    """
    UNSPECIFIED = "UNSPECIFIED"
    """
    An internal checker should never be in the unspecified state.
    """
    CREATING = "CREATING"
    """
    The checker is being created, provisioned, and configured. A checker in this state can be returned by ListInternalCheckers or GetInternalChecker, as well as by examining the long running Operation (https://cloud.google.com/apis/design/design_patterns#long_running_operations) that created it.
    """
    RUNNING = "RUNNING"
    """
    The checker is running and available for use. A checker in this state can be returned by ListInternalCheckers or GetInternalChecker as well as by examining the long running Operation (https://cloud.google.com/apis/design/design_patterns#long_running_operations) that created it. If a checker is being torn down, it is neither visible nor usable, so there is no "deleting" or "down" state.
    """


class JsonPathMatcherJsonMatcher(str, Enum):
    """
    The type of JSONPath match that will be applied to the JSON output (ContentMatcher.content)
    """
    JSON_PATH_MATCHER_OPTION_UNSPECIFIED = "JSON_PATH_MATCHER_OPTION_UNSPECIFIED"
    """
    No JSONPath matcher type specified (not valid).
    """
    EXACT_MATCH = "EXACT_MATCH"
    """
    Selects 'exact string' matching. The match succeeds if the content at the json_path within the output is exactly the same as the content string.
    """
    REGEX_MATCH = "REGEX_MATCH"
    """
    Selects regular-expression matching. The match succeeds if the content at the json_path within the output matches the regular expression specified in the content string.
    """


class LabelDescriptorValueType(str, Enum):
    """
    The type of data that can be assigned to the label.
    """
    STRING = "STRING"
    """
    A variable-length string, not to exceed 1,024 characters. This is the default value type.
    """
    BOOL = "BOOL"
    """
    Boolean; true or false.
    """
    INT64 = "INT64"
    """
    A 64-bit signed integer.
    """


class MetricDescriptorLaunchStage(str, Enum):
    """
    Optional. The launch stage of the metric definition.
    """
    LAUNCH_STAGE_UNSPECIFIED = "LAUNCH_STAGE_UNSPECIFIED"
    """
    Do not use this default value.
    """
    UNIMPLEMENTED = "UNIMPLEMENTED"
    """
    The feature is not yet implemented. Users can not use it.
    """
    PRELAUNCH = "PRELAUNCH"
    """
    Prelaunch features are hidden from users and are only visible internally.
    """
    EARLY_ACCESS = "EARLY_ACCESS"
    """
    Early Access features are limited to a closed group of testers. To use these features, you must sign up in advance and sign a Trusted Tester agreement (which includes confidentiality provisions). These features may be unstable, changed in backward-incompatible ways, and are not guaranteed to be released.
    """
    ALPHA = "ALPHA"
    """
    Alpha is a limited availability test for releases before they are cleared for widespread use. By Alpha, all significant design issues are resolved and we are in the process of verifying functionality. Alpha customers need to apply for access, agree to applicable terms, and have their projects allowlisted. Alpha releases don't have to be feature complete, no SLAs are provided, and there are no technical support obligations, but they will be far enough along that customers can actually use them in test environments or for limited-use tests -- just like they would in normal production cases.
    """
    BETA = "BETA"
    """
    Beta is the point at which we are ready to open a release for any customer to use. There are no SLA or technical support obligations in a Beta release. Products will be complete from a feature perspective, but may have some open outstanding issues. Beta releases are suitable for limited production use cases.
    """
    GA = "GA"
    """
    GA features are open to all developers and are considered stable and fully qualified for production use.
    """
    DEPRECATED = "DEPRECATED"
    """
    Deprecated features are scheduled to be shut down and removed. For more information, see the "Deprecation Policy" section of our Terms of Service (https://cloud.google.com/terms/) and the Google Cloud Platform Subject to the Deprecation Policy (https://cloud.google.com/terms/deprecation) documentation.
    """


class MetricDescriptorMetadataLaunchStage(str, Enum):
    """
    Deprecated. Must use the MetricDescriptor.launch_stage instead.
    """
    LAUNCH_STAGE_UNSPECIFIED = "LAUNCH_STAGE_UNSPECIFIED"
    """
    Do not use this default value.
    """
    UNIMPLEMENTED = "UNIMPLEMENTED"
    """
    The feature is not yet implemented. Users can not use it.
    """
    PRELAUNCH = "PRELAUNCH"
    """
    Prelaunch features are hidden from users and are only visible internally.
    """
    EARLY_ACCESS = "EARLY_ACCESS"
    """
    Early Access features are limited to a closed group of testers. To use these features, you must sign up in advance and sign a Trusted Tester agreement (which includes confidentiality provisions). These features may be unstable, changed in backward-incompatible ways, and are not guaranteed to be released.
    """
    ALPHA = "ALPHA"
    """
    Alpha is a limited availability test for releases before they are cleared for widespread use. By Alpha, all significant design issues are resolved and we are in the process of verifying functionality. Alpha customers need to apply for access, agree to applicable terms, and have their projects allowlisted. Alpha releases don't have to be feature complete, no SLAs are provided, and there are no technical support obligations, but they will be far enough along that customers can actually use them in test environments or for limited-use tests -- just like they would in normal production cases.
    """
    BETA = "BETA"
    """
    Beta is the point at which we are ready to open a release for any customer to use. There are no SLA or technical support obligations in a Beta release. Products will be complete from a feature perspective, but may have some open outstanding issues. Beta releases are suitable for limited production use cases.
    """
    GA = "GA"
    """
    GA features are open to all developers and are considered stable and fully qualified for production use.
    """
    DEPRECATED = "DEPRECATED"
    """
    Deprecated features are scheduled to be shut down and removed. For more information, see the "Deprecation Policy" section of our Terms of Service (https://cloud.google.com/terms/) and the Google Cloud Platform Subject to the Deprecation Policy (https://cloud.google.com/terms/deprecation) documentation.
    """


class MetricDescriptorMetricKind(str, Enum):
    """
    Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metric_kind and value_type might not be supported.
    """
    METRIC_KIND_UNSPECIFIED = "METRIC_KIND_UNSPECIFIED"
    """
    Do not use this default value.
    """
    GAUGE = "GAUGE"
    """
    An instantaneous measurement of a value.
    """
    DELTA = "DELTA"
    """
    The change in a value during a time interval.
    """
    CUMULATIVE = "CUMULATIVE"
    """
    A value accumulated over a time interval. Cumulative measurements in a time series should have the same start time and increasing end times, until an event resets the cumulative value to zero and sets a new start time for the following points.
    """


class MetricDescriptorValueType(str, Enum):
    """
    Whether the measurement is an integer, a floating-point number, etc. Some combinations of metric_kind and value_type might not be supported.
    """
    VALUE_TYPE_UNSPECIFIED = "VALUE_TYPE_UNSPECIFIED"
    """
    Do not use this default value.
    """
    BOOL = "BOOL"
    """
    The value is a boolean. This value type can be used only if the metric kind is GAUGE.
    """
    INT64 = "INT64"
    """
    The value is a signed 64-bit integer.
    """
    DOUBLE = "DOUBLE"
    """
    The value is a double precision floating point number.
    """
    STRING = "STRING"
    """
    The value is a text string. This value type can be used only if the metric kind is GAUGE.
    """
    DISTRIBUTION = "DISTRIBUTION"
    """
    The value is a Distribution.
    """
    MONEY = "MONEY"
    """
    The value is money.
    """


class MetricThresholdComparison(str, Enum):
    """
    The comparison to apply between the time series (indicated by filter and aggregation) and the threshold (indicated by threshold_value). The comparison is applied on each time series, with the time series on the left-hand side and the threshold on the right-hand side.Only COMPARISON_LT and COMPARISON_GT are supported currently.
    """
    COMPARISON_UNSPECIFIED = "COMPARISON_UNSPECIFIED"
    """
    No ordering relationship is specified.
    """
    COMPARISON_GT = "COMPARISON_GT"
    """
    True if the left argument is greater than the right argument.
    """
    COMPARISON_GE = "COMPARISON_GE"
    """
    True if the left argument is greater than or equal to the right argument.
    """
    COMPARISON_LT = "COMPARISON_LT"
    """
    True if the left argument is less than the right argument.
    """
    COMPARISON_LE = "COMPARISON_LE"
    """
    True if the left argument is less than or equal to the right argument.
    """
    COMPARISON_EQ = "COMPARISON_EQ"
    """
    True if the left argument is equal to the right argument.
    """
    COMPARISON_NE = "COMPARISON_NE"
    """
    True if the left argument is not equal to the right argument.
    """


class MetricThresholdEvaluationMissingData(str, Enum):
    """
    A condition control that determines how metric-threshold conditions are evaluated when data stops arriving.
    """
    EVALUATION_MISSING_DATA_UNSPECIFIED = "EVALUATION_MISSING_DATA_UNSPECIFIED"
    """
    An unspecified evaluation missing data option. Equivalent to EVALUATION_MISSING_DATA_NO_OP.
    """
    EVALUATION_MISSING_DATA_INACTIVE = "EVALUATION_MISSING_DATA_INACTIVE"
    """
    If there is no data to evaluate the condition, then evaluate the condition as false.
    """
    EVALUATION_MISSING_DATA_ACTIVE = "EVALUATION_MISSING_DATA_ACTIVE"
    """
    If there is no data to evaluate the condition, then evaluate the condition as true.
    """
    EVALUATION_MISSING_DATA_NO_OP = "EVALUATION_MISSING_DATA_NO_OP"
    """
    Do not evaluate the condition to any value if there is no data.
    """


class MonitoringQueryLanguageConditionEvaluationMissingData(str, Enum):
    """
    A condition control that determines how metric-threshold conditions are evaluated when data stops arriving.
    """
    EVALUATION_MISSING_DATA_UNSPECIFIED = "EVALUATION_MISSING_DATA_UNSPECIFIED"
    """
    An unspecified evaluation missing data option. Equivalent to EVALUATION_MISSING_DATA_NO_OP.
    """
    EVALUATION_MISSING_DATA_INACTIVE = "EVALUATION_MISSING_DATA_INACTIVE"
    """
    If there is no data to evaluate the condition, then evaluate the condition as false.
    """
    EVALUATION_MISSING_DATA_ACTIVE = "EVALUATION_MISSING_DATA_ACTIVE"
    """
    If there is no data to evaluate the condition, then evaluate the condition as true.
    """
    EVALUATION_MISSING_DATA_NO_OP = "EVALUATION_MISSING_DATA_NO_OP"
    """
    Do not evaluate the condition to any value if there is no data.
    """


class NotificationChannelVerificationStatus(str, Enum):
    """
    Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require verification or that this specific channel has been exempted from verification because it was created prior to verification being required for channels of this type.This field cannot be modified using a standard UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.
    """
    VERIFICATION_STATUS_UNSPECIFIED = "VERIFICATION_STATUS_UNSPECIFIED"
    """
    Sentinel value used to indicate that the state is unknown, omitted, or is not applicable (as in the case of channels that neither support nor require verification in order to function).
    """
    UNVERIFIED = "UNVERIFIED"
    """
    The channel has yet to be verified and requires verification to function. Note that this state also applies to the case where the verification process has been initiated by sending a verification code but where the verification code has not been submitted to complete the process.
    """
    VERIFIED = "VERIFIED"
    """
    It has been proven that notifications can be received on this notification channel and that someone on the project has access to messages that are delivered to that channel.
    """


class ResourceGroupResourceType(str, Enum):
    """
    The resource type of the group members.
    """
    RESOURCE_TYPE_UNSPECIFIED = "RESOURCE_TYPE_UNSPECIFIED"
    """
    Default value (not valid).
    """
    INSTANCE = "INSTANCE"
    """
    A group of instances from Google Cloud Platform (GCP) or Amazon Web Services (AWS).
    """
    AWS_ELB_LOAD_BALANCER = "AWS_ELB_LOAD_BALANCER"
    """
    A group of Amazon ELB load balancers.
    """


class ResponseStatusCodeStatusClass(str, Enum):
    """
    A class of status codes to accept.
    """
    STATUS_CLASS_UNSPECIFIED = "STATUS_CLASS_UNSPECIFIED"
    """
    Default value that matches no status codes.
    """
    STATUS_CLASS1XX = "STATUS_CLASS_1XX"
    """
    The class of status codes between 100 and 199.
    """
    STATUS_CLASS2XX = "STATUS_CLASS_2XX"
    """
    The class of status codes between 200 and 299.
    """
    STATUS_CLASS3XX = "STATUS_CLASS_3XX"
    """
    The class of status codes between 300 and 399.
    """
    STATUS_CLASS4XX = "STATUS_CLASS_4XX"
    """
    The class of status codes between 400 and 499.
    """
    STATUS_CLASS5XX = "STATUS_CLASS_5XX"
    """
    The class of status codes between 500 and 599.
    """
    STATUS_CLASS_ANY = "STATUS_CLASS_ANY"
    """
    The class of all status codes.
    """


class ServiceLevelObjectiveCalendarPeriod(str, Enum):
    """
    A calendar period, semantically "since the start of the current ". At this time, only DAY, WEEK, FORTNIGHT, and MONTH are supported.
    """
    CALENDAR_PERIOD_UNSPECIFIED = "CALENDAR_PERIOD_UNSPECIFIED"
    """
    Undefined period, raises an error.
    """
    DAY = "DAY"
    """
    A day.
    """
    WEEK = "WEEK"
    """
    A week. Weeks begin on Monday, following ISO 8601 (https://en.wikipedia.org/wiki/ISO_week_date).
    """
    FORTNIGHT = "FORTNIGHT"
    """
    A fortnight. The first calendar fortnight of the year begins at the start of week 1 according to ISO 8601 (https://en.wikipedia.org/wiki/ISO_week_date).
    """
    MONTH = "MONTH"
    """
    A month.
    """
    QUARTER = "QUARTER"
    """
    A quarter. Quarters start on dates 1-Jan, 1-Apr, 1-Jul, and 1-Oct of each year.
    """
    HALF = "HALF"
    """
    A half-year. Half-years start on dates 1-Jan and 1-Jul.
    """
    YEAR = "YEAR"
    """
    A year.
    """


class UptimeCheckConfigCheckerType(str, Enum):
    """
    The type of checkers to use to execute the Uptime check.
    """
    CHECKER_TYPE_UNSPECIFIED = "CHECKER_TYPE_UNSPECIFIED"
    """
    The default checker type. Currently converted to STATIC_IP_CHECKERS on creation, the default conversion behavior may change in the future.
    """
    STATIC_IP_CHECKERS = "STATIC_IP_CHECKERS"
    """
    STATIC_IP_CHECKERS are used for uptime checks that perform egress across the public internet. STATIC_IP_CHECKERS use the static IP addresses returned by ListUptimeCheckIps.
    """
    VPC_CHECKERS = "VPC_CHECKERS"
    """
    VPC_CHECKERS are used for uptime checks that perform egress using Service Directory and private network access. When using VPC_CHECKERS, the monitored resource type must be servicedirectory_service.
    """


class UptimeCheckConfigSelectedRegionsItem(str, Enum):
    REGION_UNSPECIFIED = "REGION_UNSPECIFIED"
    """
    Default value if no region is specified. Will result in Uptime checks running from all regions.
    """
    USA = "USA"
    """
    Allows checks to run from locations within the United States of America.
    """
    EUROPE = "EUROPE"
    """
    Allows checks to run from locations within the continent of Europe.
    """
    SOUTH_AMERICA = "SOUTH_AMERICA"
    """
    Allows checks to run from locations within the continent of South America.
    """
    ASIA_PACIFIC = "ASIA_PACIFIC"
    """
    Allows checks to run from locations within the Asia Pacific area (ex: Singapore).
    """
    USA_OREGON = "USA_OREGON"
    """
    Allows checks to run from locations within the western United States of America
    """
    USA_IOWA = "USA_IOWA"
    """
    Allows checks to run from locations within the central United States of America
    """
    USA_VIRGINIA = "USA_VIRGINIA"
    """
    Allows checks to run from locations within the eastern United States of America
    """
