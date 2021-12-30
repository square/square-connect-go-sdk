# ListGiftCardActivitiesRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GiftCardId** | **string** | If a gift card ID is provided, the endpoint returns activities related  to the specified gift card. Otherwise, the endpoint returns all gift card activities for  the seller. | [optional] [default to null]
**Type_** | **string** | If a [type](entity:GiftCardActivityType) is provided, the endpoint returns gift card activities of the specified type.  Otherwise, the endpoint returns all types of gift card activities. | [optional] [default to null]
**LocationId** | **string** | If a location ID is provided, the endpoint returns gift card activities for the specified location.  Otherwise, the endpoint returns gift card activities for all locations. | [optional] [default to null]
**BeginTime** | **string** | The timestamp for the beginning of the reporting period, in RFC 3339 format. This start time is inclusive. The default value is the current time minus one year. | [optional] [default to null]
**EndTime** | **string** | The timestamp for the end of the reporting period, in RFC 3339 format. This end time is inclusive. The default value is the current time. | [optional] [default to null]
**Limit** | **int32** | If a limit is provided, the endpoint returns the specified number  of results (or fewer) per page. The maximum value is 100. The default value is 50. For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | [optional] [default to null]
**Cursor** | **string** | A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query. If a cursor is not provided, the endpoint returns the first page of the results. For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | [optional] [default to null]
**SortOrder** | **string** | The order in which the endpoint returns the activities, based on &#x60;created_at&#x60;. - &#x60;ASC&#x60; - Oldest to newest. - &#x60;DESC&#x60; - Newest to oldest (default). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

