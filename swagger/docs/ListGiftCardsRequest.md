# ListGiftCardsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | If a type is provided, gift cards of this type are returned  (see [GiftCardType](entity:GiftCardType)). If no type is provided, it returns gift cards of all types. | [optional] [default to null]
**State** | **string** | If the state is provided, it returns the gift cards in the specified state  (see [GiftCardStatus](entity:GiftCardStatus)). Otherwise, it returns the gift cards of all states. | [optional] [default to null]
**Limit** | **int32** | If a value is provided, it returns only that number of results per page. The maximum number of results allowed per page is 50. The default value is 30. | [optional] [default to null]
**Cursor** | **string** | A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query. If a cursor is not provided, it returns the first page of the results.  For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | [optional] [default to null]
**CustomerId** | **string** | If a value is provided, returns only the gift cards linked to the specified customer | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

