# Booking

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique ID of this object representing a booking. | [optional] [default to null]
**Version** | **int32** | The revision number for the booking used for optimistic concurrency. | [optional] [default to null]
**Status** | [***BookingStatus**](BookingStatus.md) |  | [optional] [default to null]
**CreatedAt** | **string** | The timestamp specifying the creation time of this booking. | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp specifying the most recent update time of this booking. | [optional] [default to null]
**StartAt** | **string** | The timestamp specifying the starting time of this booking. | [optional] [default to null]
**LocationId** | **string** | The ID of the [Location](#type-location) object representing the location where the booked service is provided. | [optional] [default to null]
**CustomerId** | **string** | The ID of the [Customer](#type-Customer) object representing the customer attending this booking | [optional] [default to null]
**CustomerNote** | **string** | The free-text field for the customer to supply notes about the booking. For example, the note can be preferences that cannot be expressed by supported attributes of a relevant [CatalogObject](#type-CatalogObject) instance. | [optional] [default to null]
**SellerNote** | **string** | The free-text field for the seller to supply notes about the booking. For example, the note can be preferences that cannot be expressed by supported attributes of a specific [CatalogObject](#type-CatalogObject) instance. This field should not be visible to customers. | [optional] [default to null]
**AppointmentSegments** | [**[]AppointmentSegment**](AppointmentSegment.md) | A list of appointment segments for this booking. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

