# SearchAvailabilityFilter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartAtRange** | [***TimeRange**](TimeRange.md) |  | [default to null]
**LocationId** | **string** | The query expression to search for buyer-accessible availabilities with their location IDs matching the specified location ID. This query expression cannot be set if &#x60;booking_id&#x60; is set. | [optional] [default to null]
**SegmentFilters** | [**[]SegmentFilter**](SegmentFilter.md) | The query expression to search for buyer-accessible availabilities matching the specified list of segment filters. If the size of the &#x60;segment_filters&#x60; list is &#x60;n&#x60;, the search returns availabilities with &#x60;n&#x60; segments per availability.  This query expression cannot be set if &#x60;booking_id&#x60; is set. | [optional] [default to null]
**BookingId** | **string** | The query expression to search for buyer-accessible availabilities for an existing booking by matching the specified &#x60;booking_id&#x60; value. This is commonly used to reschedule an appointment. If this expression is set, the &#x60;location_id&#x60; and &#x60;segment_filters&#x60; expressions cannot be set. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

