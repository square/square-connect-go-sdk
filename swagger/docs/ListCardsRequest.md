# ListCardsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** | A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for your original query.  See [Pagination](https://developer.squareup.com/docs/basics/api101/pagination) for more information. | [optional] [default to null]
**CustomerId** | **string** | Limit results to cards associated with the customer supplied. By default, all cards owned by the merchant are returned. | [optional] [default to null]
**IncludeDisabled** | **bool** | Includes disabled cards. By default, all enabled cards owned by the merchant are returned. | [optional] [default to null]
**ReferenceId** | **string** | Limit results to cards associated with the reference_id supplied. | [optional] [default to null]
**SortOrder** | [***SortOrder**](SortOrder.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

