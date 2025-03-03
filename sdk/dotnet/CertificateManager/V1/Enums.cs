// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.GoogleNative.CertificateManager.V1
{
    /// <summary>
    /// Required. The key algorithm to use when generating the private key.
    /// </summary>
    [EnumType]
    public readonly struct CertificateIssuanceConfigKeyAlgorithm : IEquatable<CertificateIssuanceConfigKeyAlgorithm>
    {
        private readonly string _value;

        private CertificateIssuanceConfigKeyAlgorithm(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Unspecified key algorithm.
        /// </summary>
        public static CertificateIssuanceConfigKeyAlgorithm KeyAlgorithmUnspecified { get; } = new CertificateIssuanceConfigKeyAlgorithm("KEY_ALGORITHM_UNSPECIFIED");
        /// <summary>
        /// Specifies RSA with a 2048-bit modulus.
        /// </summary>
        public static CertificateIssuanceConfigKeyAlgorithm Rsa2048 { get; } = new CertificateIssuanceConfigKeyAlgorithm("RSA_2048");
        /// <summary>
        /// Specifies ECDSA with curve P256.
        /// </summary>
        public static CertificateIssuanceConfigKeyAlgorithm EcdsaP256 { get; } = new CertificateIssuanceConfigKeyAlgorithm("ECDSA_P256");

        public static bool operator ==(CertificateIssuanceConfigKeyAlgorithm left, CertificateIssuanceConfigKeyAlgorithm right) => left.Equals(right);
        public static bool operator !=(CertificateIssuanceConfigKeyAlgorithm left, CertificateIssuanceConfigKeyAlgorithm right) => !left.Equals(right);

        public static explicit operator string(CertificateIssuanceConfigKeyAlgorithm value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CertificateIssuanceConfigKeyAlgorithm other && Equals(other);
        public bool Equals(CertificateIssuanceConfigKeyAlgorithm other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// A predefined matcher for particular cases, other than SNI selection.
    /// </summary>
    [EnumType]
    public readonly struct CertificateMapEntryMatcher : IEquatable<CertificateMapEntryMatcher>
    {
        private readonly string _value;

        private CertificateMapEntryMatcher(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// A matcher has't been recognized.
        /// </summary>
        public static CertificateMapEntryMatcher MatcherUnspecified { get; } = new CertificateMapEntryMatcher("MATCHER_UNSPECIFIED");
        /// <summary>
        /// A primary certificate that is served when SNI wasn't specified in the request or SNI couldn't be found in the map.
        /// </summary>
        public static CertificateMapEntryMatcher Primary { get; } = new CertificateMapEntryMatcher("PRIMARY");

        public static bool operator ==(CertificateMapEntryMatcher left, CertificateMapEntryMatcher right) => left.Equals(right);
        public static bool operator !=(CertificateMapEntryMatcher left, CertificateMapEntryMatcher right) => !left.Equals(right);

        public static explicit operator string(CertificateMapEntryMatcher value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CertificateMapEntryMatcher other && Equals(other);
        public bool Equals(CertificateMapEntryMatcher other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Immutable. The scope of the certificate.
    /// </summary>
    [EnumType]
    public readonly struct CertificateScope : IEquatable<CertificateScope>
    {
        private readonly string _value;

        private CertificateScope(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        /// <summary>
        /// Certificates with default scope are served from core Google data centers. If unsure, choose this option.
        /// </summary>
        public static CertificateScope Default { get; } = new CertificateScope("DEFAULT");
        /// <summary>
        /// Certificates with scope EDGE_CACHE are special-purposed certificates, served from non-core Google data centers.
        /// </summary>
        public static CertificateScope EdgeCache { get; } = new CertificateScope("EDGE_CACHE");

        public static bool operator ==(CertificateScope left, CertificateScope right) => left.Equals(right);
        public static bool operator !=(CertificateScope left, CertificateScope right) => !left.Equals(right);

        public static explicit operator string(CertificateScope value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CertificateScope other && Equals(other);
        public bool Equals(CertificateScope other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
