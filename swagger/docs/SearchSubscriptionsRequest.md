# SearchSubscriptionsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** | When the total number of resulting subscriptions exceeds the limit of a paged response,  specify the cursor returned from a preceding response here to fetch the next set of results. If the cursor is unset, the response contains the last page of the results.  For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | [optional] [default to null]
**Limit** | **int32** | The upper limit on the number of subscriptions to return in a paged response. | [optional] [default to null]
**Query** | [***SearchSubscriptionsQuery**](SearchSubscriptionsQuery.md) |  | [optional] [default to null]
**Include** | **[]string** | An option to include related information in the response.   The supported values are:   - &#x60;actions&#x60;: to include scheduled actions on the targeted subscriptions. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

