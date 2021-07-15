# SearchCustomersRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** | Include the pagination cursor in subsequent calls to this endpoint to retrieve the next set of results associated with the original query.  For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | [optional] [default to null]
**Limit** | **int64** | A limit on the number of results to be returned in a single page. The limit is advisory. The implementation might return more or fewer results. If the supplied limit is negative, zero, or higher than the maximum limit of 100, it is ignored. | [optional] [default to null]
**Query** | [***CustomerQuery**](CustomerQuery.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

