// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.WorkflowExecutions.V1Beta
{
    /// <summary>
    /// The call logging level associated to this execution.
    /// </summary>
    [EnumType]
    public readonly struct ExecutionCallLogLevel : IEquatable<ExecutionCallLogLevel>
    {
        private readonly string _value;

        private ExecutionCallLogLevel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// No call logging level specified.
        /// </summary>
        public static ExecutionCallLogLevel CallLogLevelUnspecified { get; } = new ExecutionCallLogLevel("CALL_LOG_LEVEL_UNSPECIFIED");
        /// <summary>
        /// Log all call steps within workflows, all call returns, and all exceptions raised.
        /// </summary>
        public static ExecutionCallLogLevel LogAllCalls { get; } = new ExecutionCallLogLevel("LOG_ALL_CALLS");
        /// <summary>
        /// Log only exceptions that are raised from call steps within workflows.
        /// </summary>
        public static ExecutionCallLogLevel LogErrorsOnly { get; } = new ExecutionCallLogLevel("LOG_ERRORS_ONLY");

        public static bool operator ==(ExecutionCallLogLevel left, ExecutionCallLogLevel right) => left.Equals(right);
        public static bool operator !=(ExecutionCallLogLevel left, ExecutionCallLogLevel right) => !left.Equals(right);

        public static explicit operator string(ExecutionCallLogLevel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ExecutionCallLogLevel other && Equals(other);
        public bool Equals(ExecutionCallLogLevel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
