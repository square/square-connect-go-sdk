# SearchInvoicesRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | [***InvoiceQuery**](InvoiceQuery.md) |  | [default to null]
**Limit** | **int32** | The maximum number of invoices to return (200 is the maximum &#x60;limit&#x60;).  If not provided, the server  uses a default limit of 100 invoices. | [optional] [default to null]
**Cursor** | **string** | A pagination cursor returned by a previous call to this endpoint.  Provide this cursor to retrieve the next set of results for your original query.  For more information, see [Pagination](https://developer.squareup.com/docs/docs/working-with-apis/pagination). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

