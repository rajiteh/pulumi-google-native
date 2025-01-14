// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.DeploymentManager.V2Beta
{
    /// <summary>
    /// The log type that this config enables.
    /// </summary>
    [EnumType]
    public readonly struct AuditLogConfigLogType : IEquatable<AuditLogConfigLogType>
    {
        private readonly string _value;

        private AuditLogConfigLogType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Default case. Should never be this.
        /// </summary>
        public static AuditLogConfigLogType LogTypeUnspecified { get; } = new AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED");
        /// <summary>
        /// Admin reads. Example: CloudIAM getIamPolicy
        /// </summary>
        public static AuditLogConfigLogType AdminRead { get; } = new AuditLogConfigLogType("ADMIN_READ");
        /// <summary>
        /// Data writes. Example: CloudSQL Users create
        /// </summary>
        public static AuditLogConfigLogType DataWrite { get; } = new AuditLogConfigLogType("DATA_WRITE");
        /// <summary>
        /// Data reads. Example: CloudSQL Users list
        /// </summary>
        public static AuditLogConfigLogType DataRead { get; } = new AuditLogConfigLogType("DATA_READ");

        public static bool operator ==(AuditLogConfigLogType left, AuditLogConfigLogType right) => left.Equals(right);
        public static bool operator !=(AuditLogConfigLogType left, AuditLogConfigLogType right) => !left.Equals(right);

        public static explicit operator string(AuditLogConfigLogType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AuditLogConfigLogType other && Equals(other);
        public bool Equals(AuditLogConfigLogType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    [EnumType]
    public readonly struct CompositeTypeStatus : IEquatable<CompositeTypeStatus>
    {
        private readonly string _value;

        private CompositeTypeStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CompositeTypeStatus UnknownStatus { get; } = new CompositeTypeStatus("UNKNOWN_STATUS");
        public static CompositeTypeStatus Deprecated { get; } = new CompositeTypeStatus("DEPRECATED");
        public static CompositeTypeStatus Experimental { get; } = new CompositeTypeStatus("EXPERIMENTAL");
        public static CompositeTypeStatus Supported { get; } = new CompositeTypeStatus("SUPPORTED");

        public static bool operator ==(CompositeTypeStatus left, CompositeTypeStatus right) => left.Equals(right);
        public static bool operator !=(CompositeTypeStatus left, CompositeTypeStatus right) => !left.Equals(right);

        public static explicit operator string(CompositeTypeStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CompositeTypeStatus other && Equals(other);
        public bool Equals(CompositeTypeStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Level to record this diagnostic.
    /// </summary>
    [EnumType]
    public readonly struct DiagnosticLevel : IEquatable<DiagnosticLevel>
    {
        private readonly string _value;

        private DiagnosticLevel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DiagnosticLevel Unknown { get; } = new DiagnosticLevel("UNKNOWN");
        /// <summary>
        /// If level is informational, it only gets displayed as part of the resource.
        /// </summary>
        public static DiagnosticLevel Information { get; } = new DiagnosticLevel("INFORMATION");
        /// <summary>
        /// If level is warning, will end up in the resource as a warning.
        /// </summary>
        public static DiagnosticLevel Warning { get; } = new DiagnosticLevel("WARNING");
        /// <summary>
        /// If level is error, it will indicate an error occurred after finishCondition is set, and this field will populate resource errors and operation errors.
        /// </summary>
        public static DiagnosticLevel Error { get; } = new DiagnosticLevel("ERROR");

        public static bool operator ==(DiagnosticLevel left, DiagnosticLevel right) => left.Equals(right);
        public static bool operator !=(DiagnosticLevel left, DiagnosticLevel right) => !left.Equals(right);

        public static explicit operator string(DiagnosticLevel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DiagnosticLevel other && Equals(other);
        public bool Equals(DiagnosticLevel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The location where this mapping applies.
    /// </summary>
    [EnumType]
    public readonly struct InputMappingLocation : IEquatable<InputMappingLocation>
    {
        private readonly string _value;

        private InputMappingLocation(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static InputMappingLocation Unknown { get; } = new InputMappingLocation("UNKNOWN");
        public static InputMappingLocation Path { get; } = new InputMappingLocation("PATH");
        public static InputMappingLocation Query { get; } = new InputMappingLocation("QUERY");
        public static InputMappingLocation Body { get; } = new InputMappingLocation("BODY");
        public static InputMappingLocation Header { get; } = new InputMappingLocation("HEADER");

        public static bool operator ==(InputMappingLocation left, InputMappingLocation right) => left.Equals(right);
        public static bool operator !=(InputMappingLocation left, InputMappingLocation right) => !left.Equals(right);

        public static explicit operator string(InputMappingLocation value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InputMappingLocation other && Equals(other);
        public bool Equals(InputMappingLocation other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Which interpreter (python or jinja) should be used during expansion.
    /// </summary>
    [EnumType]
    public readonly struct TemplateContentsInterpreter : IEquatable<TemplateContentsInterpreter>
    {
        private readonly string _value;

        private TemplateContentsInterpreter(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static TemplateContentsInterpreter UnknownInterpreter { get; } = new TemplateContentsInterpreter("UNKNOWN_INTERPRETER");
        public static TemplateContentsInterpreter Python { get; } = new TemplateContentsInterpreter("PYTHON");
        public static TemplateContentsInterpreter Jinja { get; } = new TemplateContentsInterpreter("JINJA");

        public static bool operator ==(TemplateContentsInterpreter left, TemplateContentsInterpreter right) => left.Equals(right);
        public static bool operator !=(TemplateContentsInterpreter left, TemplateContentsInterpreter right) => !left.Equals(right);

        public static explicit operator string(TemplateContentsInterpreter value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is TemplateContentsInterpreter other && Equals(other);
        public bool Equals(TemplateContentsInterpreter other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Customize how deployment manager will validate the resource against schema errors.
    /// </summary>
    [EnumType]
    public readonly struct ValidationOptionsSchemaValidation : IEquatable<ValidationOptionsSchemaValidation>
    {
        private readonly string _value;

        private ValidationOptionsSchemaValidation(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ValidationOptionsSchemaValidation Unknown { get; } = new ValidationOptionsSchemaValidation("UNKNOWN");
        /// <summary>
        /// Ignore schema failures.
        /// </summary>
        public static ValidationOptionsSchemaValidation Ignore { get; } = new ValidationOptionsSchemaValidation("IGNORE");
        /// <summary>
        /// Ignore schema failures but display them as warnings.
        /// </summary>
        public static ValidationOptionsSchemaValidation IgnoreWithWarnings { get; } = new ValidationOptionsSchemaValidation("IGNORE_WITH_WARNINGS");
        /// <summary>
        /// Fail the resource if the schema is not valid, this is the default behavior.
        /// </summary>
        public static ValidationOptionsSchemaValidation Fail { get; } = new ValidationOptionsSchemaValidation("FAIL");

        public static bool operator ==(ValidationOptionsSchemaValidation left, ValidationOptionsSchemaValidation right) => left.Equals(right);
        public static bool operator !=(ValidationOptionsSchemaValidation left, ValidationOptionsSchemaValidation right) => !left.Equals(right);

        public static explicit operator string(ValidationOptionsSchemaValidation value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ValidationOptionsSchemaValidation other && Equals(other);
        public bool Equals(ValidationOptionsSchemaValidation other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specify what to do with extra properties when executing a request.
    /// </summary>
    [EnumType]
    public readonly struct ValidationOptionsUndeclaredProperties : IEquatable<ValidationOptionsUndeclaredProperties>
    {
        private readonly string _value;

        private ValidationOptionsUndeclaredProperties(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ValidationOptionsUndeclaredProperties Unknown { get; } = new ValidationOptionsUndeclaredProperties("UNKNOWN");
        /// <summary>
        /// Always include even if not present on discovery doc.
        /// </summary>
        public static ValidationOptionsUndeclaredProperties Include { get; } = new ValidationOptionsUndeclaredProperties("INCLUDE");
        /// <summary>
        /// Always ignore if not present on discovery doc.
        /// </summary>
        public static ValidationOptionsUndeclaredProperties Ignore { get; } = new ValidationOptionsUndeclaredProperties("IGNORE");
        /// <summary>
        /// Include on request, but emit a warning.
        /// </summary>
        public static ValidationOptionsUndeclaredProperties IncludeWithWarnings { get; } = new ValidationOptionsUndeclaredProperties("INCLUDE_WITH_WARNINGS");
        /// <summary>
        /// Ignore properties, but emit a warning.
        /// </summary>
        public static ValidationOptionsUndeclaredProperties IgnoreWithWarnings { get; } = new ValidationOptionsUndeclaredProperties("IGNORE_WITH_WARNINGS");
        /// <summary>
        /// Always fail if undeclared properties are present.
        /// </summary>
        public static ValidationOptionsUndeclaredProperties Fail { get; } = new ValidationOptionsUndeclaredProperties("FAIL");

        public static bool operator ==(ValidationOptionsUndeclaredProperties left, ValidationOptionsUndeclaredProperties right) => left.Equals(right);
        public static bool operator !=(ValidationOptionsUndeclaredProperties left, ValidationOptionsUndeclaredProperties right) => !left.Equals(right);

        public static explicit operator string(ValidationOptionsUndeclaredProperties value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ValidationOptionsUndeclaredProperties other && Equals(other);
        public bool Equals(ValidationOptionsUndeclaredProperties other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
