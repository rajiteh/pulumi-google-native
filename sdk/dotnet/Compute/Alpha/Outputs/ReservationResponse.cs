// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    /// <summary>
    /// Represents a reservation resource. A reservation ensures that capacity is held in a specific zone even if the reserved VMs are not running. For more information, read Reserving zonal resources.
    /// </summary>
    [OutputType]
    public sealed class ReservationResponse
    {
        /// <summary>
        /// Reservation for aggregated resources, providing shape flexibility.
        /// </summary>
        public readonly Outputs.AllocationAggregateReservationResponse AggregateReservation;
        /// <summary>
        /// Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
        /// </summary>
        public readonly string Commitment;
        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Type of the resource. Always compute#reservations for reservations.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource policies to be added to this reservation. The key is defined by user, and the value is resource policy url. This is to define placement policy with reservation.
        /// </summary>
        public readonly ImmutableDictionary<string, string> ResourcePolicies;
        /// <summary>
        /// Status information for Reservation resource.
        /// </summary>
        public readonly Outputs.AllocationResourceStatusResponse ResourceStatus;
        /// <summary>
        /// Reserved for future use.
        /// </summary>
        public readonly bool SatisfiesPzs;
        /// <summary>
        /// Server-defined fully-qualified URL for this resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// Server-defined URL for this resource with the resource id.
        /// </summary>
        public readonly string SelfLinkWithId;
        /// <summary>
        /// Specify share-settings to create a shared reservation. This property is optional. For more information about the syntax and options for this field and its subfields, see the guide for creating a shared reservation.
        /// </summary>
        public readonly Outputs.ShareSettingsResponse ShareSettings;
        /// <summary>
        /// Reservation for instances with specific machine shapes.
        /// </summary>
        public readonly Outputs.AllocationSpecificSKUReservationResponse SpecificReservation;
        /// <summary>
        /// Indicates whether the reservation can be consumed by VMs with affinity for "any" reservation. If the field is set, then only VMs that target the reservation by name can consume from this reservation.
        /// </summary>
        public readonly bool SpecificReservationRequired;
        /// <summary>
        /// The status of the reservation.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Zone in which the reservation resides. A zone must be provided if the reservation is created within a commitment.
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private ReservationResponse(
            Outputs.AllocationAggregateReservationResponse aggregateReservation,

            string commitment,

            string creationTimestamp,

            string description,

            string kind,

            string name,

            ImmutableDictionary<string, string> resourcePolicies,

            Outputs.AllocationResourceStatusResponse resourceStatus,

            bool satisfiesPzs,

            string selfLink,

            string selfLinkWithId,

            Outputs.ShareSettingsResponse shareSettings,

            Outputs.AllocationSpecificSKUReservationResponse specificReservation,

            bool specificReservationRequired,

            string status,

            string zone)
        {
            AggregateReservation = aggregateReservation;
            Commitment = commitment;
            CreationTimestamp = creationTimestamp;
            Description = description;
            Kind = kind;
            Name = name;
            ResourcePolicies = resourcePolicies;
            ResourceStatus = resourceStatus;
            SatisfiesPzs = satisfiesPzs;
            SelfLink = selfLink;
            SelfLinkWithId = selfLinkWithId;
            ShareSettings = shareSettings;
            SpecificReservation = specificReservation;
            SpecificReservationRequired = specificReservationRequired;
            Status = status;
            Zone = zone;
        }
    }
}
