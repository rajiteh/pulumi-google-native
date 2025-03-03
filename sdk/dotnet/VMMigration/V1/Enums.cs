// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.VMMigration.V1
{
    /// <summary>
    /// The disk type to use in the VM.
    /// </summary>
    [EnumType]
    public readonly struct ComputeEngineTargetDefaultsDiskType : IEquatable<ComputeEngineTargetDefaultsDiskType>
    {
        private readonly string _value;

        private ComputeEngineTargetDefaultsDiskType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// An unspecified disk type. Will be used as STANDARD.
        /// </summary>
        public static ComputeEngineTargetDefaultsDiskType ComputeEngineDiskTypeUnspecified { get; } = new ComputeEngineTargetDefaultsDiskType("COMPUTE_ENGINE_DISK_TYPE_UNSPECIFIED");
        /// <summary>
        /// A Standard disk type.
        /// </summary>
        public static ComputeEngineTargetDefaultsDiskType ComputeEngineDiskTypeStandard { get; } = new ComputeEngineTargetDefaultsDiskType("COMPUTE_ENGINE_DISK_TYPE_STANDARD");
        /// <summary>
        /// SSD hard disk type.
        /// </summary>
        public static ComputeEngineTargetDefaultsDiskType ComputeEngineDiskTypeSsd { get; } = new ComputeEngineTargetDefaultsDiskType("COMPUTE_ENGINE_DISK_TYPE_SSD");
        /// <summary>
        /// An alternative to SSD persistent disks that balance performance and cost.
        /// </summary>
        public static ComputeEngineTargetDefaultsDiskType ComputeEngineDiskTypeBalanced { get; } = new ComputeEngineTargetDefaultsDiskType("COMPUTE_ENGINE_DISK_TYPE_BALANCED");

        public static bool operator ==(ComputeEngineTargetDefaultsDiskType left, ComputeEngineTargetDefaultsDiskType right) => left.Equals(right);
        public static bool operator !=(ComputeEngineTargetDefaultsDiskType left, ComputeEngineTargetDefaultsDiskType right) => !left.Equals(right);

        public static explicit operator string(ComputeEngineTargetDefaultsDiskType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComputeEngineTargetDefaultsDiskType other && Equals(other);
        public bool Equals(ComputeEngineTargetDefaultsDiskType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The license type to use in OS adaptation.
    /// </summary>
    [EnumType]
    public readonly struct ComputeEngineTargetDefaultsLicenseType : IEquatable<ComputeEngineTargetDefaultsLicenseType>
    {
        private readonly string _value;

        private ComputeEngineTargetDefaultsLicenseType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The license type is the default for the OS.
        /// </summary>
        public static ComputeEngineTargetDefaultsLicenseType ComputeEngineLicenseTypeDefault { get; } = new ComputeEngineTargetDefaultsLicenseType("COMPUTE_ENGINE_LICENSE_TYPE_DEFAULT");
        /// <summary>
        /// The license type is Pay As You Go license type.
        /// </summary>
        public static ComputeEngineTargetDefaultsLicenseType ComputeEngineLicenseTypePayg { get; } = new ComputeEngineTargetDefaultsLicenseType("COMPUTE_ENGINE_LICENSE_TYPE_PAYG");
        /// <summary>
        /// The license type is Bring Your Own License type.
        /// </summary>
        public static ComputeEngineTargetDefaultsLicenseType ComputeEngineLicenseTypeByol { get; } = new ComputeEngineTargetDefaultsLicenseType("COMPUTE_ENGINE_LICENSE_TYPE_BYOL");

        public static bool operator ==(ComputeEngineTargetDefaultsLicenseType left, ComputeEngineTargetDefaultsLicenseType right) => left.Equals(right);
        public static bool operator !=(ComputeEngineTargetDefaultsLicenseType left, ComputeEngineTargetDefaultsLicenseType right) => !left.Equals(right);

        public static explicit operator string(ComputeEngineTargetDefaultsLicenseType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComputeEngineTargetDefaultsLicenseType other && Equals(other);
        public bool Equals(ComputeEngineTargetDefaultsLicenseType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// How the instance should behave when the host machine undergoes maintenance that may temporarily impact instance performance.
    /// </summary>
    [EnumType]
    public readonly struct ComputeSchedulingOnHostMaintenance : IEquatable<ComputeSchedulingOnHostMaintenance>
    {
        private readonly string _value;

        private ComputeSchedulingOnHostMaintenance(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// An unknown, unexpected behavior.
        /// </summary>
        public static ComputeSchedulingOnHostMaintenance OnHostMaintenanceUnspecified { get; } = new ComputeSchedulingOnHostMaintenance("ON_HOST_MAINTENANCE_UNSPECIFIED");
        /// <summary>
        /// Terminate the instance when the host machine undergoes maintenance.
        /// </summary>
        public static ComputeSchedulingOnHostMaintenance Terminate { get; } = new ComputeSchedulingOnHostMaintenance("TERMINATE");
        /// <summary>
        /// Migrate the instance when the host machine undergoes maintenance.
        /// </summary>
        public static ComputeSchedulingOnHostMaintenance Migrate { get; } = new ComputeSchedulingOnHostMaintenance("MIGRATE");

        public static bool operator ==(ComputeSchedulingOnHostMaintenance left, ComputeSchedulingOnHostMaintenance right) => left.Equals(right);
        public static bool operator !=(ComputeSchedulingOnHostMaintenance left, ComputeSchedulingOnHostMaintenance right) => !left.Equals(right);

        public static explicit operator string(ComputeSchedulingOnHostMaintenance value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComputeSchedulingOnHostMaintenance other && Equals(other);
        public bool Equals(ComputeSchedulingOnHostMaintenance other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Whether the Instance should be automatically restarted whenever it is terminated by Compute Engine (not terminated by user). This configuration is identical to `automaticRestart` field in Compute Engine create instance under scheduling. It was changed to an enum (instead of a boolean) to match the default value in Compute Engine which is automatic restart.
    /// </summary>
    [EnumType]
    public readonly struct ComputeSchedulingRestartType : IEquatable<ComputeSchedulingRestartType>
    {
        private readonly string _value;

        private ComputeSchedulingRestartType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unspecified behavior. This will use the default.
        /// </summary>
        public static ComputeSchedulingRestartType RestartTypeUnspecified { get; } = new ComputeSchedulingRestartType("RESTART_TYPE_UNSPECIFIED");
        /// <summary>
        /// The Instance should be automatically restarted whenever it is terminated by Compute Engine.
        /// </summary>
        public static ComputeSchedulingRestartType AutomaticRestart { get; } = new ComputeSchedulingRestartType("AUTOMATIC_RESTART");
        /// <summary>
        /// The Instance isn't automatically restarted whenever it is terminated by Compute Engine.
        /// </summary>
        public static ComputeSchedulingRestartType NoAutomaticRestart { get; } = new ComputeSchedulingRestartType("NO_AUTOMATIC_RESTART");

        public static bool operator ==(ComputeSchedulingRestartType left, ComputeSchedulingRestartType right) => left.Equals(right);
        public static bool operator !=(ComputeSchedulingRestartType left, ComputeSchedulingRestartType right) => !left.Equals(right);

        public static explicit operator string(ComputeSchedulingRestartType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComputeSchedulingRestartType other && Equals(other);
        public bool Equals(ComputeSchedulingRestartType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The operator to use for the node resources specified in the `values` parameter.
    /// </summary>
    [EnumType]
    public readonly struct SchedulingNodeAffinityOperator : IEquatable<SchedulingNodeAffinityOperator>
    {
        private readonly string _value;

        private SchedulingNodeAffinityOperator(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// An unknown, unexpected behavior.
        /// </summary>
        public static SchedulingNodeAffinityOperator OperatorUnspecified { get; } = new SchedulingNodeAffinityOperator("OPERATOR_UNSPECIFIED");
        /// <summary>
        /// The node resource group should be in these resources affinity.
        /// </summary>
        public static SchedulingNodeAffinityOperator In { get; } = new SchedulingNodeAffinityOperator("IN");
        /// <summary>
        /// The node resource group should not be in these resources affinity.
        /// </summary>
        public static SchedulingNodeAffinityOperator NotIn { get; } = new SchedulingNodeAffinityOperator("NOT_IN");

        public static bool operator ==(SchedulingNodeAffinityOperator left, SchedulingNodeAffinityOperator right) => left.Equals(right);
        public static bool operator !=(SchedulingNodeAffinityOperator left, SchedulingNodeAffinityOperator right) => !left.Equals(right);

        public static explicit operator string(SchedulingNodeAffinityOperator value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SchedulingNodeAffinityOperator other && Equals(other);
        public bool Equals(SchedulingNodeAffinityOperator other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Time frame of the report.
    /// </summary>
    [EnumType]
    public readonly struct UtilizationReportTimeFrame : IEquatable<UtilizationReportTimeFrame>
    {
        private readonly string _value;

        private UtilizationReportTimeFrame(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// The time frame was not specified and will default to WEEK.
        /// </summary>
        public static UtilizationReportTimeFrame TimeFrameUnspecified { get; } = new UtilizationReportTimeFrame("TIME_FRAME_UNSPECIFIED");
        /// <summary>
        /// One week.
        /// </summary>
        public static UtilizationReportTimeFrame Week { get; } = new UtilizationReportTimeFrame("WEEK");
        /// <summary>
        /// One month.
        /// </summary>
        public static UtilizationReportTimeFrame Month { get; } = new UtilizationReportTimeFrame("MONTH");
        /// <summary>
        /// One year.
        /// </summary>
        public static UtilizationReportTimeFrame Year { get; } = new UtilizationReportTimeFrame("YEAR");

        public static bool operator ==(UtilizationReportTimeFrame left, UtilizationReportTimeFrame right) => left.Equals(right);
        public static bool operator !=(UtilizationReportTimeFrame left, UtilizationReportTimeFrame right) => !left.Equals(right);

        public static explicit operator string(UtilizationReportTimeFrame value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is UtilizationReportTimeFrame other && Equals(other);
        public bool Equals(UtilizationReportTimeFrame other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// The power state of the VM at the moment list was taken.
    /// </summary>
    [EnumType]
    public readonly struct VmwareVmDetailsPowerState : IEquatable<VmwareVmDetailsPowerState>
    {
        private readonly string _value;

        private VmwareVmDetailsPowerState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Power state is not specified.
        /// </summary>
        public static VmwareVmDetailsPowerState PowerStateUnspecified { get; } = new VmwareVmDetailsPowerState("POWER_STATE_UNSPECIFIED");
        /// <summary>
        /// The VM is turned ON.
        /// </summary>
        public static VmwareVmDetailsPowerState On { get; } = new VmwareVmDetailsPowerState("ON");
        /// <summary>
        /// The VM is turned OFF.
        /// </summary>
        public static VmwareVmDetailsPowerState Off { get; } = new VmwareVmDetailsPowerState("OFF");
        /// <summary>
        /// The VM is suspended. This is similar to hibernation or sleep mode.
        /// </summary>
        public static VmwareVmDetailsPowerState Suspended { get; } = new VmwareVmDetailsPowerState("SUSPENDED");

        public static bool operator ==(VmwareVmDetailsPowerState left, VmwareVmDetailsPowerState right) => left.Equals(right);
        public static bool operator !=(VmwareVmDetailsPowerState left, VmwareVmDetailsPowerState right) => !left.Equals(right);

        public static explicit operator string(VmwareVmDetailsPowerState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is VmwareVmDetailsPowerState other && Equals(other);
        public bool Equals(VmwareVmDetailsPowerState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
