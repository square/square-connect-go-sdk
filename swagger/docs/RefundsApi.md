# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPaymentRefund**](RefundsApi.md#GetPaymentRefund) | **Get** /v2/refunds/{refund_id} | GetPaymentRefund
[**ListPaymentRefunds**](RefundsApi.md#ListPaymentRefunds) | **Get** /v2/refunds | ListPaymentRefunds
[**RefundPayment**](RefundsApi.md#RefundPayment) | **Post** /v2/refunds | RefundPayment

# **GetPaymentRefund**
> GetPaymentRefundResponse GetPaymentRefund(ctx, refundId)
GetPaymentRefund

Retrieves a specific refund using the `refund_id`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **refundId** | **string**| The unique ID for the desired &#x60;PaymentRefund&#x60;. | 

### Return type

[**GetPaymentRefundResponse**](GetPaymentRefundResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPaymentRefunds**
> ListPaymentRefundsResponse ListPaymentRefunds(ctx, optional)
ListPaymentRefunds

Retrieves a list of refunds for the account making the request.  The maximum results per page is 100.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RefundsApiListPaymentRefundsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RefundsApiListPaymentRefundsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **beginTime** | **optional.String**| The timestamp for the beginning of the requested reporting period, in RFC 3339 format.  Default: The current time minus one year. | 
 **endTime** | **optional.String**| The timestamp for the end of the requested reporting period, in RFC 3339 format.  Default: The current time. | 
 **sortOrder** | **optional.String**| The order in which results are listed: - &#x60;ASC&#x60; - Oldest to newest. - &#x60;DESC&#x60; - Newest to oldest (default). | 
 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query.  For more information, see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination). | 
 **locationId** | **optional.String**| Limit results to the location supplied. By default, results are returned for all locations associated with the seller. | 
 **status** | **optional.String**| If provided, only refunds with the given status are returned. For a list of refund status values, see [PaymentRefund](#type-paymentrefund).  Default: If omitted, refunds are returned regardless of their status. | 
 **sourceType** | **optional.String**| If provided, only refunds with the given source type are returned. - &#x60;CARD&#x60; - List refunds only for payments where &#x60;CARD&#x60; was specified as the payment source.  Default: If omitted, refunds are returned regardless of the source type. | 
 **limit** | **optional.Int32**| The maximum number of results to be returned in a single page.  It is possible to receive fewer results than the specified limit on a given page.  If the supplied value is greater than 100, no more than 100 results are returned.  Default: 100 | 

### Return type

[**ListPaymentRefundsResponse**](ListPaymentRefundsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RefundPayment**
> RefundPaymentResponse RefundPayment(ctx, body)
RefundPayment

Refunds a payment. You can refund the entire payment amount or a  portion of it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RefundPaymentRequest**](RefundPaymentRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**RefundPaymentResponse**](RefundPaymentResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

