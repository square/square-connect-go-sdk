# ListInvoicesResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoices** | [**[]Invoice**](Invoice.md) | The invoices retrieved. | [optional] [default to null]
**Cursor** | **string** | When a response is truncated, it includes a cursor that you can use in a  subsequent request to fetch the next set of invoices. If empty, this is the final  response.  For more information, see [Pagination](https://developer.squareup.com/docs/docs/working-with-apis/pagination). | [optional] [default to null]
**Errors** | [**[]ModelError**](Error.md) | Information about errors encountered during the request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

