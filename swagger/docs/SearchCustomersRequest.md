# SearchCustomersRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** | Include the pagination cursor in subsequent calls to this endpoint to retrieve the next set of results associated with the original query.  For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | [optional] [default to null]
**Limit** | **int64** | The maximum number of results to return in a single page. This limit is advisory. The response might contain more or fewer results. The limit is ignored if it is less than the minimum or greater than the maximum value. The default value is 100.  For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | [optional] [default to null]
**Query** | [***CustomerQuery**](CustomerQuery.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

