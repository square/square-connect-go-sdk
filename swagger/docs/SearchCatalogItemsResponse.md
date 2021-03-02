# SearchCatalogItemsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ModelError**](Error.md) | Any errors that occurred during the request. | [optional] [default to null]
**Items** | [**[]CatalogObject**](CatalogObject.md) | Returned items matching the specified query expressions. | [optional] [default to null]
**Cursor** | **string** | Pagination token used in the next request to return more of the search result. | [optional] [default to null]
**MatchedVariationIds** | **[]string** | Ids of returned item variations matching the specified query expression. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

