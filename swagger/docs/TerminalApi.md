# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelTerminalAction**](TerminalApi.md#CancelTerminalAction) | **Post** /v2/terminals/actions/{action_id}/cancel | CancelTerminalAction
[**CancelTerminalCheckout**](TerminalApi.md#CancelTerminalCheckout) | **Post** /v2/terminals/checkouts/{checkout_id}/cancel | CancelTerminalCheckout
[**CancelTerminalRefund**](TerminalApi.md#CancelTerminalRefund) | **Post** /v2/terminals/refunds/{terminal_refund_id}/cancel | CancelTerminalRefund
[**CreateTerminalAction**](TerminalApi.md#CreateTerminalAction) | **Post** /v2/terminals/actions | CreateTerminalAction
[**CreateTerminalCheckout**](TerminalApi.md#CreateTerminalCheckout) | **Post** /v2/terminals/checkouts | CreateTerminalCheckout
[**CreateTerminalRefund**](TerminalApi.md#CreateTerminalRefund) | **Post** /v2/terminals/refunds | CreateTerminalRefund
[**GetTerminalAction**](TerminalApi.md#GetTerminalAction) | **Get** /v2/terminals/actions/{action_id} | GetTerminalAction
[**GetTerminalCheckout**](TerminalApi.md#GetTerminalCheckout) | **Get** /v2/terminals/checkouts/{checkout_id} | GetTerminalCheckout
[**GetTerminalRefund**](TerminalApi.md#GetTerminalRefund) | **Get** /v2/terminals/refunds/{terminal_refund_id} | GetTerminalRefund
[**SearchTerminalActions**](TerminalApi.md#SearchTerminalActions) | **Post** /v2/terminals/actions/search | SearchTerminalActions
[**SearchTerminalCheckouts**](TerminalApi.md#SearchTerminalCheckouts) | **Post** /v2/terminals/checkouts/search | SearchTerminalCheckouts
[**SearchTerminalRefunds**](TerminalApi.md#SearchTerminalRefunds) | **Post** /v2/terminals/refunds/search | SearchTerminalRefunds

# **CancelTerminalAction**
> CancelTerminalActionResponse CancelTerminalAction(ctx, body, actionId)
CancelTerminalAction

Cancels a Terminal action request if the status of the request permits it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CancelTerminalActionRequest**](CancelTerminalActionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **actionId** | **string**| Unique ID for the desired &#x60;TerminalAction&#x60; | 

### Return type

[**CancelTerminalActionResponse**](CancelTerminalActionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **CreateTerminalAction**
> CreateTerminalActionResponse CreateTerminalAction(ctx, body)
CreateTerminalAction

Creates a Terminal action request and sends it to the specified device.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateTerminalActionRequest**](CreateTerminalActionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateTerminalActionResponse**](CreateTerminalActionResponse.md)

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

Creates a request to refund an Interac payment completed on a Square Terminal. Refunds for Interac payments on a Square Terminal are supported only for Interac debit cards in Canada. Other refunds for Terminal payments should use the Refunds API. For more information, see [Refunds API](api:Refunds).

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

# **GetTerminalAction**
> GetTerminalActionResponse GetTerminalAction(ctx, actionId)
GetTerminalAction

Retrieves a Terminal action request by `action_id`. Terminal action requests are available for 30 days.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **actionId** | **string**| Unique ID for the desired &#x60;TerminalAction&#x60; | 

### Return type

[**GetTerminalActionResponse**](GetTerminalActionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTerminalCheckout**
> GetTerminalCheckoutResponse GetTerminalCheckout(ctx, checkoutId)
GetTerminalCheckout

Retrieves a Terminal checkout request by `checkout_id`. Terminal checkout requests are available for 30 days.

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

Retrieves an Interac Terminal refund object by ID. Terminal refund objects are available for 30 days.

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

# **SearchTerminalActions**
> SearchTerminalActionsResponse SearchTerminalActions(ctx, body)
SearchTerminalActions

Retrieves a filtered list of Terminal action requests created by the account making the request. Terminal action requests are available for 30 days.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchTerminalActionsRequest**](SearchTerminalActionsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchTerminalActionsResponse**](SearchTerminalActionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTerminalCheckouts**
> SearchTerminalCheckoutsResponse SearchTerminalCheckouts(ctx, body)
SearchTerminalCheckouts

Returns a filtered list of Terminal checkout requests created by the application making the request. Only Terminal checkout requests created for the merchant scoped to the OAuth token are returned. Terminal checkout requests are available for 30 days.

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

Retrieves a filtered list of Interac Terminal refund requests created by the seller making the request. Terminal refund requests are available for 30 days.

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

