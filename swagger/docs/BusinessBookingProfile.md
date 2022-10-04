# BusinessBookingProfile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SellerId** | **string** | The ID of the seller, obtainable using the Merchants API. | [optional] [default to null]
**CreatedAt** | **string** | The RFC 3339 timestamp specifying the booking&#x27;s creation time. | [optional] [default to null]
**BookingEnabled** | **bool** | Indicates whether the seller is open for booking. | [optional] [default to null]
**CustomerTimezoneChoice** | [***BusinessBookingProfileCustomerTimezoneChoice**](BusinessBookingProfileCustomerTimezoneChoice.md) |  | [optional] [default to null]
**BookingPolicy** | [***BusinessBookingProfileBookingPolicy**](BusinessBookingProfileBookingPolicy.md) |  | [optional] [default to null]
**AllowUserCancel** | **bool** | Indicates whether customers can cancel or reschedule their own bookings (&#x60;true&#x60;) or not (&#x60;false&#x60;). | [optional] [default to null]
**BusinessAppointmentSettings** | [***BusinessAppointmentSettings**](BusinessAppointmentSettings.md) |  | [optional] [default to null]
**SupportSellerLevelWrites** | **bool** | Indicates whether the seller&#x27;s subscription to Square Appointments supports creating, updating or canceling an appointment through the API (&#x60;true&#x60;) or not (&#x60;false&#x60;) using seller permission. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

