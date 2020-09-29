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

# **CancelPayment**
> CancelPaymentResponse CancelPayment(ctx, paymentId)
CancelPayment

Cancels (voids) a payment. If you set `autocomplete` to false when creating a payment,  you can cancel the payment using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **paymentId** | **string**| &#x60;payment_id&#x60; identifying the payment to be canceled. | 

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

Cancels (voids) a payment identified by the idempotency key that is specified in the request.  Use this method when status of a CreatePayment request is unknown. For example, after you send a CreatePayment request a network error occurs and you don't get a response. In this case, you can direct Square to cancel the payment using this endpoint. In the request, you provide the same idempotency key that you provided in your CreatePayment request you want  to cancel. After cancelling the payment, you can submit your CreatePayment request again.  Note that if no payment with the specified idempotency key is found, no action is taken, the end point returns successfully.

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
> CompletePaymentResponse CompletePayment(ctx, paymentId)
CompletePayment

Completes (captures) a payment.  By default, payments are set to complete immediately after they are created.  If you set autocomplete to false when creating a payment, you can complete (capture)  the payment using this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **paymentId** | **string**| Unique ID identifying the payment to be completed. | 

### Return type

[**CompletePaymentResponse**](CompletePaymentResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePayment**
> CreatePaymentResponse CreatePayment(ctx, body)
CreatePayment

Charges a payment source, for example, a card  represented by customer's card on file or a card nonce. In addition  to the payment source, the request must also include the  amount to accept for the payment.  There are several optional parameters that you can include in the request.  For example, tip money, whether to autocomplete the payment, or a reference ID to correlate this payment with another system.   The `PAYMENTS_WRITE_ADDITIONAL_RECIPIENTS` OAuth permission is required to enable application fees.

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

Retrieves details for a specific Payment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **paymentId** | **string**| Unique ID for the desired &#x60;Payment&#x60;. | 

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

Retrieves a list of payments taken by the account making the request.  Max results per page: 100

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentsApiListPaymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentsApiListPaymentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **beginTime** | **optional.String**| Timestamp for the beginning of the reporting period, in RFC 3339 format. Inclusive. Default: The current time minus one year. | 
 **endTime** | **optional.String**| Timestamp for the end of the requested reporting period, in RFC 3339 format.  Default: The current time. | 
 **sortOrder** | **optional.String**| The order in which results are listed. - &#x60;ASC&#x60; - oldest to newest - &#x60;DESC&#x60; - newest to oldest (default). | 
 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for the original query.  See [Pagination](https://developer.squareup.com/docs/basics/api101/pagination) for more information. | 
 **locationId** | **optional.String**| Limit results to the location supplied. By default, results are returned for the default (main) location associated with the merchant. | 
 **total** | **optional.Int64**| The exact amount in the total_money for a &#x60;Payment&#x60;. | 
 **last4** | **optional.String**| The last 4 digits of &#x60;Payment&#x60; card. | 
 **cardBrand** | **optional.String**| The brand of &#x60;Payment&#x60; card. For example, &#x60;VISA&#x60; | 
 **limit** | **optional.Int32**| Maximum number of results to be returned in a single page. It is possible to receive fewer results than the specified limit on a given page.  If the supplied value is greater than 100, at most 100 results will be returned.  Default: &#x60;100&#x60; | 

### Return type

[**ListPaymentsResponse**](ListPaymentsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

