# ListWebhookSubscriptionsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ModelError**](Error.md) | Information on errors encountered during the request. | [optional] [default to null]
**Subscriptions** | [**[]WebhookSubscription**](WebhookSubscription.md) | The requested list of [Subscription](entity:WebhookSubscription)s. | [optional] [default to null]
**Cursor** | **string** | The pagination cursor to be used in a subsequent request. If empty, this is the final response.  For more information, see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

