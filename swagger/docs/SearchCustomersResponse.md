# SearchCustomersResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ModelError**](Error.md) | Any errors that occurred during the request. | [optional] [default to null]
**Customers** | [**[]Customer**](Customer.md) | An array of &#x60;Customer&#x60; objects that match a query. | [optional] [default to null]
**Cursor** | **string** | A pagination cursor that can be used during subsequent calls to &#x60;SearchCustomers&#x60; to retrieve the next set of results associated with the original query. Pagination cursors are only present when a request succeeds and additional results are available.  For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

