# V1ListSettlementsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | [***SortOrder**](SortOrder.md) |  | [optional] [default to null]
**BeginTime** | **string** | The beginning of the requested reporting period, in ISO 8601 format. If this value is before January 1, 2013 (2013-01-01T00:00:00Z), this endpoint returns an error. Default value: The current time minus one year. | [optional] [default to null]
**EndTime** | **string** | The end of the requested reporting period, in ISO 8601 format. If this value is more than one year greater than begin_time, this endpoint returns an error. Default value: The current time. | [optional] [default to null]
**Limit** | **int32** | The maximum number of settlements to return in a single response. This value cannot exceed 200. | [optional] [default to null]
**Status** | [***V1ListSettlementsRequestStatus**](V1ListSettlementsRequestStatus.md) |  | [optional] [default to null]
**BatchToken** | **string** | A pagination cursor to retrieve the next set of results for your original query to the endpoint. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

