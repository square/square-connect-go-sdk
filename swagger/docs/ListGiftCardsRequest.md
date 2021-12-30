# ListGiftCardsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | If a [type](entity:GiftCardType) is provided, the endpoint returns gift cards of the specified type. Otherwise, the endpoint returns gift cards of all types. | [optional] [default to null]
**State** | **string** | If a [state](entity:GiftCardStatus) is provided, the endpoint returns the gift cards in the specified state. Otherwise, the endpoint returns the gift cards of all states. | [optional] [default to null]
**Limit** | **int32** | If a limit is provided, the endpoint returns only the specified number of results per page. The maximum value is 50. The default value is 30. For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | [optional] [default to null]
**Cursor** | **string** | A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query. If a cursor is not provided, the endpoint returns the first page of the results.  For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | [optional] [default to null]
**CustomerId** | **string** | If a customer ID is provided, the endpoint returns only the gift cards linked to the specified customer. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

