# ListWebhookSubscriptionsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** | A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for your original query.  For more information, see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination). | [optional] [default to null]
**IncludeDisabled** | **bool** | Includes disabled [Subscription](entity:WebhookSubscription)s. By default, all enabled [Subscription](entity:WebhookSubscription)s are returned. | [optional] [default to null]
**SortOrder** | [***SortOrder**](SortOrder.md) |  | [optional] [default to null]
**Limit** | **int32** | The maximum number of results to be returned in a single page. It is possible to receive fewer results than the specified limit on a given page. The default value of 100 is also the maximum allowed value. If the provided value is greater than 100, it is ignored and the default value is used instead.  Default: 100 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

