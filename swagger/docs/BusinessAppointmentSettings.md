# BusinessAppointmentSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationTypes** | [**[]BusinessAppointmentSettingsBookingLocationType**](BusinessAppointmentSettingsBookingLocationType.md) | Types of the location allowed for bookings. See [BusinessAppointmentSettingsBookingLocationType](#type-businessappointmentsettingsbookinglocationtype) for possible values | [optional] [default to null]
**AlignmentTime** | [***BusinessAppointmentSettingsAlignmentTime**](BusinessAppointmentSettingsAlignmentTime.md) |  | [optional] [default to null]
**MinBookingLeadTimeSeconds** | **int32** | The minimum lead time in seconds before a service can be booked. Bookings must be created at least this far ahead of the booking&#x27;s starting time. | [optional] [default to null]
**MaxBookingLeadTimeSeconds** | **int32** | The maximum lead time in seconds before a service can be booked. Bookings must be created at most this far ahead of the booking&#x27;s starting time. | [optional] [default to null]
**AnyTeamMemberBookingEnabled** | **bool** | Indicates whether a customer can choose from all available time slots and have a staff member assigned automatically (&#x60;true&#x60;) or not (&#x60;false&#x60;). | [optional] [default to null]
**MultipleServiceBookingEnabled** | **bool** | Indicates whether a customer can book multiple services in a single online booking. | [optional] [default to null]
**MaxAppointmentsPerDayLimitType** | [***BusinessAppointmentSettingsMaxAppointmentsPerDayLimitType**](BusinessAppointmentSettingsMaxAppointmentsPerDayLimitType.md) |  | [optional] [default to null]
**MaxAppointmentsPerDayLimit** | **int32** | The maximum number of daily appointments per team member or per location. | [optional] [default to null]
**CancellationWindowSeconds** | **int32** | The cut-off time in seconds for allowing clients to cancel or reschedule an appointment. | [optional] [default to null]
**CancellationFeeMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**CancellationPolicy** | [***BusinessAppointmentSettingsCancellationPolicy**](BusinessAppointmentSettingsCancellationPolicy.md) |  | [optional] [default to null]
**CancellationPolicyText** | **string** | The free-form text of the seller&#x27;s cancellation policy. | [optional] [default to null]
**SkipBookingFlowStaffSelection** | **bool** | Indicates whether customers has an assigned staff member (&#x60;true&#x60;) or can select s staff member of their choice (&#x60;false&#x60;). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

