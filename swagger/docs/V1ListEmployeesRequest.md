# V1ListEmployeesRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | [***SortOrder**](SortOrder.md) |  | [optional] [default to null]
**BeginUpdatedAt** | **string** | If filtering results by their updated_at field, the beginning of the requested reporting period, in ISO 8601 format | [optional] [default to null]
**EndUpdatedAt** | **string** | If filtering results by there updated_at field, the end of the requested reporting period, in ISO 8601 format. | [optional] [default to null]
**BeginCreatedAt** | **string** | If filtering results by their created_at field, the beginning of the requested reporting period, in ISO 8601 format. | [optional] [default to null]
**EndCreatedAt** | **string** | If filtering results by their created_at field, the end of the requested reporting period, in ISO 8601 format. | [optional] [default to null]
**Status** | [***V1ListEmployeesRequestStatus**](V1ListEmployeesRequestStatus.md) |  | [optional] [default to null]
**ExternalId** | **string** | If provided, the endpoint returns only employee entities with the specified external_id. | [optional] [default to null]
**Limit** | **int32** | The maximum integer number of employee entities to return in a single response. Default 100, maximum 200. | [optional] [default to null]
**BatchToken** | **string** | A pagination cursor to retrieve the next set of results for your original query to the endpoint. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

