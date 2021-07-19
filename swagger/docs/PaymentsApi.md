# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPayment**](PaymentsApi.md#CancelPayment) | **Post** /v2/payments/{payment_id}/cancel | CancelPayment
[**CancelPaymentByIdempotencyKey**](PaymentsApi.md#CancelPaymentByIdempotencyKey) | **Post** /v2/payments/cancel | CancelPaymentByIdempotencyKey
[**CompletePayment**](PaymentsApi.md#CompletePayment) | **Post** /v2/payments/{payment_id}/complete | CompletePayment
[**CreatePayment**](PaymentsApi.md#CreatePayment) | **Post** /v2/payments | CreatePayment
[**GetPayment**](PaymentsApi.md#GetPayment) | **Get** /v2/payments/{payment_id} | GetPayment
[**ListPayments**](PaymentsApi.md#ListPayments) | **Get** /v2/payments | ListPayments
[**UpdatePayment**](PaymentsApi.md#UpdatePayment) | **Put** /v2/payments/{payment_id} | UpdatePayment

# **CancelPayment**
> CancelPaymentResponse CancelPayment(ctx, paymentId)
CancelPayment

Cancels (voids) a payment. You can use this endpoint to cancel a payment with  the APPROVED `status`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **paymentId** | **string**| The ID of the payment to cancel. | 

### Return type

[**CancelPaymentResponse**](CancelPaymentResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelPaymentByIdempotencyKey**
> CancelPaymentByIdempotencyKeyResponse CancelPaymentByIdempotencyKey(ctx, body)
CancelPaymentByIdempotencyKey

Cancels (voids) a payment identified by the idempotency key that is specified in the request.  Use this method when the status of a `CreatePayment` request is unknown (for example, after you send a `CreatePayment` request, a network error occurs and you do not get a response). In this case, you can direct Square to cancel the payment using this endpoint. In the request, you provide the same idempotency key that you provided in your `CreatePayment` request that you want to cancel. After canceling the payment, you can submit your `CreatePayment` request again.  Note that if no payment with the specified idempotency key is found, no action is taken and the endpoint returns successfully.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CancelPaymentByIdempotencyKeyRequest**](CancelPaymentByIdempotencyKeyRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CancelPaymentByIdempotencyKeyResponse**](CancelPaymentByIdempotencyKeyResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CompletePayment**
> CompletePaymentResponse CompletePayment(ctx, body, paymentId)
CompletePayment

Completes (captures) a payment. By default, payments are set to complete immediately after they are created.  You can use this endpoint to complete a payment with the APPROVED `status`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CompletePaymentRequest**](CompletePaymentRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **paymentId** | **string**| The unique ID identifying the payment to be completed. | 

### Return type

[**CompletePaymentResponse**](CompletePaymentResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePayment**
> CreatePaymentResponse CreatePayment(ctx, body)
CreatePayment

Creates a payment using the provided source. You can use this endpoint  to charge a card (credit/debit card or     Square gift card) or record a payment that the seller received outside of Square  (cash payment from a buyer or a payment that an external entity  processed on behalf of the seller).  The endpoint creates a  `Payment` object and returns it in the response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreatePaymentRequest**](CreatePaymentRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreatePaymentResponse**](CreatePaymentResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPayment**
> GetPaymentResponse GetPayment(ctx, paymentId)
GetPayment

Retrieves details for a specific payment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **paymentId** | **string**| A unique ID for the desired payment. | 

### Return type

[**GetPaymentResponse**](GetPaymentResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPayments**
> ListPaymentsResponse ListPayments(ctx, optional)
ListPayments

Retrieves a list of payments taken by the account making the request.  Results are eventually consistent, and new payments or changes to payments might take several seconds to appear.  The maximum results per page is 100.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentsApiListPaymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentsApiListPaymentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **beginTime** | **optional.String**| The timestamp for the beginning of the reporting period, in RFC 3339 format. Inclusive. Default: The current time minus one year. | 
 **endTime** | **optional.String**| The timestamp for the end of the reporting period, in RFC 3339 format.  Default: The current time. | 
 **sortOrder** | **optional.String**| The order in which results are listed: - &#x60;ASC&#x60; - Oldest to newest. - &#x60;DESC&#x60; - Newest to oldest (default). | 
 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query.  For more information, see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination). | 
 **locationId** | **optional.String**| Limit results to the location supplied. By default, results are returned for the default (main) location associated with the seller. | 
 **total** | **optional.Int64**| The exact amount in the &#x60;total_money&#x60; for a payment. | 
 **last4** | **optional.String**| The last four digits of a payment card. | 
 **cardBrand** | **optional.String**| The brand of the payment card (for example, VISA). | 
 **limit** | **optional.Int32**| The maximum number of results to be returned in a single page. It is possible to receive fewer results than the specified limit on a given page.  The default value of 100 is also the maximum allowed value. If the provided value is  greater than 100, it is ignored and the default value is used instead.  Default: &#x60;100&#x60; | 

### Return type

[**ListPaymentsResponse**](ListPaymentsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePayment**
> UpdatePaymentResponse UpdatePayment(ctx, body, paymentId)
UpdatePayment

Updates a payment with the APPROVED status. You can update the `amount_money` and `tip_money` using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatePaymentRequest**](UpdatePaymentRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **paymentId** | **string**| The ID of the payment to update. | 

### Return type

[**UpdatePaymentResponse**](UpdatePaymentResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

