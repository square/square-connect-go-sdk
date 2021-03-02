# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelInvoice**](InvoicesApi.md#CancelInvoice) | **Post** /v2/invoices/{invoice_id}/cancel | CancelInvoice
[**CreateInvoice**](InvoicesApi.md#CreateInvoice) | **Post** /v2/invoices | CreateInvoice
[**DeleteInvoice**](InvoicesApi.md#DeleteInvoice) | **Delete** /v2/invoices/{invoice_id} | DeleteInvoice
[**GetInvoice**](InvoicesApi.md#GetInvoice) | **Get** /v2/invoices/{invoice_id} | GetInvoice
[**ListInvoices**](InvoicesApi.md#ListInvoices) | **Get** /v2/invoices | ListInvoices
[**PublishInvoice**](InvoicesApi.md#PublishInvoice) | **Post** /v2/invoices/{invoice_id}/publish | PublishInvoice
[**SearchInvoices**](InvoicesApi.md#SearchInvoices) | **Post** /v2/invoices/search | SearchInvoices
[**UpdateInvoice**](InvoicesApi.md#UpdateInvoice) | **Put** /v2/invoices/{invoice_id} | UpdateInvoice

# **CancelInvoice**
> CancelInvoiceResponse CancelInvoice(ctx, body, invoiceId)
CancelInvoice

Cancels an invoice. The seller cannot collect payments for  the canceled invoice.  You cannot cancel an invoice in the `DRAFT` state or in a terminal state: `PAID`, `REFUNDED`, `CANCELED`, or `FAILED`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CancelInvoiceRequest**](CancelInvoiceRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **invoiceId** | **string**| The ID of the [invoice](#type-invoice) to cancel. | 

### Return type

[**CancelInvoiceResponse**](CancelInvoiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateInvoice**
> CreateInvoiceResponse CreateInvoice(ctx, body)
CreateInvoice

Creates a draft [invoice](#type-invoice)  for an order created using the Orders API.  A draft invoice remains in your account and no action is taken.  You must publish the invoice before Square can process it (send it to the customer's email address or charge the customer’s card on file).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateInvoiceRequest**](CreateInvoiceRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateInvoiceResponse**](CreateInvoiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInvoice**
> DeleteInvoiceResponse DeleteInvoice(ctx, invoiceId, optional)
DeleteInvoice

Deletes the specified invoice. When an invoice is deleted, the  associated Order status changes to CANCELED. You can only delete a draft  invoice (you cannot delete a published invoice, including one that is scheduled for processing).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceId** | **string**| The ID of the invoice to delete. | 
 **optional** | ***InvoicesApiDeleteInvoiceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoicesApiDeleteInvoiceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.Int32**| The version of the [invoice](#type-invoice) to delete. If you do not know the version, you can call [GetInvoice](#endpoint-Invoices-GetInvoice) or  [ListInvoices](#endpoint-Invoices-ListInvoices). | 

### Return type

[**DeleteInvoiceResponse**](DeleteInvoiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvoice**
> GetInvoiceResponse GetInvoice(ctx, invoiceId)
GetInvoice

Retrieves an invoice by invoice ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **invoiceId** | **string**| The id of the invoice to retrieve. | 

### Return type

[**GetInvoiceResponse**](GetInvoiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInvoices**
> ListInvoicesResponse ListInvoices(ctx, locationId, optional)
ListInvoices

Returns a list of invoices for a given location. The response  is paginated. If truncated, the response includes a `cursor` that you     use in a subsequent request to fetch the next set of invoices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location for which to list invoices. | 
 **optional** | ***InvoicesApiListInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InvoicesApiListInvoicesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint.  Provide this cursor to retrieve the next set of results for your original query.  For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | 
 **limit** | **optional.Int32**| The maximum number of invoices to return (200 is the maximum &#x60;limit&#x60;).  If not provided, the server  uses a default limit of 100 invoices. | 

### Return type

[**ListInvoicesResponse**](ListInvoicesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublishInvoice**
> PublishInvoiceResponse PublishInvoice(ctx, body, invoiceId)
PublishInvoice

Publishes the specified draft invoice.   After an invoice is published, Square  follows up based on the invoice configuration. For example, Square  sends the invoice to the customer's email address, charges the customer's card on file, or does  nothing. Square also makes the invoice available on a Square-hosted invoice page.   The invoice `status` also changes from `DRAFT` to a status  based on the invoice configuration. For example, the status changes to `UNPAID` if  Square emails the invoice or `PARTIALLY_PAID` if Square charge a card on file for a portion of the  invoice amount).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PublishInvoiceRequest**](PublishInvoiceRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **invoiceId** | **string**| The id of the invoice to publish. | 

### Return type

[**PublishInvoiceResponse**](PublishInvoiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchInvoices**
> SearchInvoicesResponse SearchInvoices(ctx, body)
SearchInvoices

Searches for invoices from a location specified in  the filter. You can optionally specify customers in the filter for whom to  retrieve invoices. In the current implementation, you can only specify one location and  optionally one customer.  The response is paginated. If truncated, the response includes a `cursor`  that you use in a subsequent request to fetch the next set of invoices.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchInvoicesRequest**](SearchInvoicesRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchInvoicesResponse**](SearchInvoicesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInvoice**
> UpdateInvoiceResponse UpdateInvoice(ctx, body, invoiceId)
UpdateInvoice

Updates an invoice by modifying fields, clearing fields, or both. For most updates, you can use a sparse  `Invoice` object to add fields or change values, and use the `fields_to_clear` field to specify fields to clear.  However, some restrictions apply. For example, you cannot change the `order_id` or `location_id` field, and you  must provide the complete `custom_fields` list to update a custom field. Published invoices have additional restrictions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateInvoiceRequest**](UpdateInvoiceRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **invoiceId** | **string**| The ID of the invoice to update. | 

### Return type

[**UpdateInvoiceResponse**](UpdateInvoiceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

