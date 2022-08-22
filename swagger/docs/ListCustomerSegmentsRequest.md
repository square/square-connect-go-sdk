# ListCustomerSegmentsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** | A pagination cursor returned by previous calls to &#x60;ListCustomerSegments&#x60;. This cursor is used to retrieve the next set of query results.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | [optional] [default to null]
**Limit** | **int32** | The maximum number of results to return in a single page. This limit is advisory. The response might contain more or fewer results. If the specified limit is less than 1 or greater than 50, Square returns a &#x60;400 VALUE_TOO_LOW&#x60; or &#x60;400 VALUE_TOO_HIGH&#x60; error. The default value is 50.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

