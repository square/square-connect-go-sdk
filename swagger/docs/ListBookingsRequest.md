# ListBookingsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | The maximum number of results per page to return in a paged response. | [optional] [default to null]
**Cursor** | **string** | The pagination cursor from the preceding response to return the next page of the results. Do not set this when retrieving the first page of the results. | [optional] [default to null]
**TeamMemberId** | **string** | The team member for whom to retrieve bookings. If this is not set, bookings of all members are retrieved. | [optional] [default to null]
**LocationId** | **string** | The location for which to retrieve bookings. If this is not set, all locations&#x27; bookings are retrieved. | [optional] [default to null]
**StartAtMin** | **string** | The RFC 3339 timestamp specifying the earliest of the start time. If this is not set, the current time is used. | [optional] [default to null]
**StartAtMax** | **string** | The RFC 3339 timestamp specifying the latest of the start time. If this is not set, the time of 31 days after &#x60;start_at_min&#x60; is used. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

