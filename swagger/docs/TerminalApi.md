# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelTerminalCheckout**](TerminalApi.md#CancelTerminalCheckout) | **Post** /v2/terminals/checkouts/{checkout_id}/cancel | CancelTerminalCheckout
[**CancelTerminalRefund**](TerminalApi.md#CancelTerminalRefund) | **Post** /v2/terminals/refunds/{terminal_refund_id}/cancel | CancelTerminalRefund
[**CreateTerminalCheckout**](TerminalApi.md#CreateTerminalCheckout) | **Post** /v2/terminals/checkouts | CreateTerminalCheckout
[**CreateTerminalRefund**](TerminalApi.md#CreateTerminalRefund) | **Post** /v2/terminals/refunds | CreateTerminalRefund
[**GetTerminalCheckout**](TerminalApi.md#GetTerminalCheckout) | **Get** /v2/terminals/checkouts/{checkout_id} | GetTerminalCheckout
[**GetTerminalRefund**](TerminalApi.md#GetTerminalRefund) | **Get** /v2/terminals/refunds/{terminal_refund_id} | GetTerminalRefund
[**SearchTerminalCheckouts**](TerminalApi.md#SearchTerminalCheckouts) | **Post** /v2/terminals/checkouts/search | SearchTerminalCheckouts
[**SearchTerminalRefunds**](TerminalApi.md#SearchTerminalRefunds) | **Post** /v2/terminals/refunds/search | SearchTerminalRefunds

# **CancelTerminalCheckout**
> CancelTerminalCheckoutResponse CancelTerminalCheckout(ctx, body, checkoutId)
CancelTerminalCheckout

Cancels a Terminal checkout request if the status of the request permits it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CancelTerminalCheckoutRequest**](CancelTerminalCheckoutRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **checkoutId** | **string**| The unique ID for the desired &#x60;TerminalCheckout&#x60;. | 

### Return type

[**CancelTerminalCheckoutResponse**](CancelTerminalCheckoutResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelTerminalRefund**
> CancelTerminalRefundResponse CancelTerminalRefund(ctx, body, terminalRefundId)
CancelTerminalRefund

Cancels an Interac Terminal refund request by refund request ID if the status of the request permits it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CancelTerminalRefundRequest**](CancelTerminalRefundRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **terminalRefundId** | **string**| The unique ID for the desired &#x60;TerminalRefund&#x60;. | 

### Return type

[**CancelTerminalRefundResponse**](CancelTerminalRefundResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTerminalCheckout**
> CreateTerminalCheckoutResponse CreateTerminalCheckout(ctx, body)
CreateTerminalCheckout

Creates a Terminal checkout request and sends it to the specified device to take a payment for the requested amount.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateTerminalCheckoutRequest**](CreateTerminalCheckoutRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateTerminalCheckoutResponse**](CreateTerminalCheckoutResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTerminalRefund**
> CreateTerminalRefundResponse CreateTerminalRefund(ctx, body)
CreateTerminalRefund

Creates a request to refund an Interac payment completed on a Square Terminal.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateTerminalRefundRequest**](CreateTerminalRefundRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateTerminalRefundResponse**](CreateTerminalRefundResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTerminalCheckout**
> GetTerminalCheckoutResponse GetTerminalCheckout(ctx, checkoutId)
GetTerminalCheckout

Retrieves a Terminal checkout request by `checkout_id`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **checkoutId** | **string**| The unique ID for the desired &#x60;TerminalCheckout&#x60;. | 

### Return type

[**GetTerminalCheckoutResponse**](GetTerminalCheckoutResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTerminalRefund**
> GetTerminalRefundResponse GetTerminalRefund(ctx, terminalRefundId)
GetTerminalRefund

Retrieves an Interac Terminal refund object by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **terminalRefundId** | **string**| The unique ID for the desired &#x60;TerminalRefund&#x60;. | 

### Return type

[**GetTerminalRefundResponse**](GetTerminalRefundResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTerminalCheckouts**
> SearchTerminalCheckoutsResponse SearchTerminalCheckouts(ctx, body)
SearchTerminalCheckouts

Retrieves a filtered list of Terminal checkout requests created by the account making the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchTerminalCheckoutsRequest**](SearchTerminalCheckoutsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchTerminalCheckoutsResponse**](SearchTerminalCheckoutsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTerminalRefunds**
> SearchTerminalRefundsResponse SearchTerminalRefunds(ctx, body)
SearchTerminalRefunds

Retrieves a filtered list of Interac Terminal refund requests created by the seller making the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchTerminalRefundsRequest**](SearchTerminalRefundsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchTerminalRefundsResponse**](SearchTerminalRefundsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

