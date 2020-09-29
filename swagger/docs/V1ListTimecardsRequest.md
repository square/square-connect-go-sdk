# V1ListTimecardsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | [***SortOrder**](SortOrder.md) |  | [optional] [default to null]
**EmployeeId** | **string** | If provided, the endpoint returns only timecards for the employee with the specified ID. | [optional] [default to null]
**BeginClockinTime** | **string** | If filtering results by their clockin_time field, the beginning of the requested reporting period, in ISO 8601 format. | [optional] [default to null]
**EndClockinTime** | **string** | If filtering results by their clockin_time field, the end of the requested reporting period, in ISO 8601 format. | [optional] [default to null]
**BeginClockoutTime** | **string** | If filtering results by their clockout_time field, the beginning of the requested reporting period, in ISO 8601 format. | [optional] [default to null]
**EndClockoutTime** | **string** | If filtering results by their clockout_time field, the end of the requested reporting period, in ISO 8601 format. | [optional] [default to null]
**BeginUpdatedAt** | **string** | If filtering results by their updated_at field, the beginning of the requested reporting period, in ISO 8601 format. | [optional] [default to null]
**EndUpdatedAt** | **string** | If filtering results by their updated_at field, the end of the requested reporting period, in ISO 8601 format. | [optional] [default to null]
**Deleted** | **bool** | If true, only deleted timecards are returned. If false, only valid timecards are returned.If you don&#x27;t provide this parameter, both valid and deleted timecards are returned. | [optional] [default to null]
**Limit** | **int32** | The maximum integer number of employee entities to return in a single response. Default 100, maximum 200. | [optional] [default to null]
**BatchToken** | **string** | A pagination cursor to retrieve the next set of results for your original query to the endpoint. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

