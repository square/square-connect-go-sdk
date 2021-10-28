# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CaptureTransaction**](TransactionsApi.md#CaptureTransaction) | **Post** /v2/locations/{location_id}/transactions/{transaction_id}/capture | CaptureTransaction
[**ListTransactions**](TransactionsApi.md#ListTransactions) | **Get** /v2/locations/{location_id}/transactions | ListTransactions
[**RetrieveTransaction**](TransactionsApi.md#RetrieveTransaction) | **Get** /v2/locations/{location_id}/transactions/{transaction_id} | RetrieveTransaction
[**VoidTransaction**](TransactionsApi.md#VoidTransaction) | **Post** /v2/locations/{location_id}/transactions/{transaction_id}/void | VoidTransaction

# **CaptureTransaction**
> CaptureTransactionResponse CaptureTransaction(ctx, locationId, transactionId)
CaptureTransaction

Captures a transaction that was created with the [Charge](api-endpoint:Transactions-Charge) endpoint with a `delay_capture` value of `true`.   See [Delayed capture transactions](https://developer.squareup.com/docs/payments/transactions/overview#delayed-capture) for more information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**|  | 
  **transactionId** | **string**|  | 

### Return type

[**CaptureTransactionResponse**](CaptureTransactionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTransactions**
> ListTransactionsResponse ListTransactions(ctx, locationId, optional)
ListTransactions

Lists transactions for a particular location.  Transactions include payment information from sales and exchanges and refund information from returns and exchanges.  Max results per [page](https://developer.squareup.com/docs/working-with-apis/pagination): 50

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to list transactions for. | 
 **optional** | ***TransactionsApiListTransactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransactionsApiListTransactionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **beginTime** | **optional.String**| The beginning of the requested reporting period, in RFC 3339 format.  See [Date ranges](https://developer.squareup.com/docs/build-basics/working-with-dates) for details on date inclusivity/exclusivity.  Default value: The current time minus one year. | 
 **endTime** | **optional.String**| The end of the requested reporting period, in RFC 3339 format.  See [Date ranges](https://developer.squareup.com/docs/build-basics/working-with-dates) for details on date inclusivity/exclusivity.  Default value: The current time. | 
 **sortOrder** | [**optional.Interface of SortOrder**](.md)| The order in which results are listed in the response (&#x60;ASC&#x60; for oldest first, &#x60;DESC&#x60; for newest first).  Default value: &#x60;DESC&#x60; | 
 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for your original query.  See [Paginating results](https://developer.squareup.com/docs/working-with-apis/pagination) for more information. | 

### Return type

[**ListTransactionsResponse**](ListTransactionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveTransaction**
> RetrieveTransactionResponse RetrieveTransaction(ctx, locationId, transactionId)
RetrieveTransaction

Retrieves details for a single transaction.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the transaction&#x27;s associated location. | 
  **transactionId** | **string**| The ID of the transaction to retrieve. | 

### Return type

[**RetrieveTransactionResponse**](RetrieveTransactionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VoidTransaction**
> VoidTransactionResponse VoidTransaction(ctx, locationId, transactionId)
VoidTransaction

Cancels a transaction that was created with the [Charge](api-endpoint:Transactions-Charge) endpoint with a `delay_capture` value of `true`.   See [Delayed capture transactions](https://developer.squareup.com/docs/payments/transactions/overview#delayed-capture) for more information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**|  | 
  **transactionId** | **string**|  | 

### Return type

[**VoidTransactionResponse**](VoidTransactionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

