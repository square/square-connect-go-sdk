# SearchSubscriptionsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ModelError**](Error.md) | Errors encountered during the request. | [optional] [default to null]
**Subscriptions** | [**[]Subscription**](Subscription.md) | The subscriptions matching the specified query expressions. | [optional] [default to null]
**Cursor** | **string** | When the total number of resulting subscription exceeds the limit of a paged response,  the response includes a cursor for you to use in a subsequent request to fetch the next set of results. If the cursor is unset, the response contains the last page of the results.  For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

