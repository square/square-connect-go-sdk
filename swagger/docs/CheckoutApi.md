# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCheckout**](CheckoutApi.md#CreateCheckout) | **Post** /v2/locations/{location_id}/checkouts | CreateCheckout
[**CreatePaymentLink**](CheckoutApi.md#CreatePaymentLink) | **Post** /v2/online-checkout/payment-links | CreatePaymentLink
[**DeletePaymentLink**](CheckoutApi.md#DeletePaymentLink) | **Delete** /v2/online-checkout/payment-links/{id} | DeletePaymentLink
[**ListPaymentLinks**](CheckoutApi.md#ListPaymentLinks) | **Get** /v2/online-checkout/payment-links | ListPaymentLinks
[**RetrievePaymentLink**](CheckoutApi.md#RetrievePaymentLink) | **Get** /v2/online-checkout/payment-links/{id} | RetrievePaymentLink
[**UpdatePaymentLink**](CheckoutApi.md#UpdatePaymentLink) | **Put** /v2/online-checkout/payment-links/{id} | UpdatePaymentLink

# **CreateCheckout**
> CreateCheckoutResponse CreateCheckout(ctx, body, locationId)
CreateCheckout

Links a `checkoutId` to a `checkout_page_url` that customers are directed to in order to provide their payment information using a payment processing workflow hosted on connect.squareup.com.    NOTE: The Checkout API has been updated with new features.  For more information, see [Checkout API highlights](https://developer.squareup.com/docs/checkout-api#checkout-api-highlights). We recommend that you use the new [CreatePaymentLink](api-endpoint:Checkout-CreatePaymentLink)  endpoint in place of this previously released endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateCheckoutRequest**](CreateCheckoutRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the business location to associate the checkout with. | 

### Return type

[**CreateCheckoutResponse**](CreateCheckoutResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePaymentLink**
> CreatePaymentLinkResponse CreatePaymentLink(ctx, body)
CreatePaymentLink

Creates a Square-hosted checkout page. Applications can share the resulting payment link with their buyer to pay for goods and services.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreatePaymentLinkRequest**](CreatePaymentLinkRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreatePaymentLinkResponse**](CreatePaymentLinkResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePaymentLink**
> DeletePaymentLinkResponse DeletePaymentLink(ctx, id)
DeletePaymentLink

Deletes a payment link.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the payment link to delete. | 

### Return type

[**DeletePaymentLinkResponse**](DeletePaymentLinkResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPaymentLinks**
> ListPaymentLinksResponse ListPaymentLinks(ctx, optional)
ListPaymentLinks

Lists all payment links.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CheckoutApiListPaymentLinksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CheckoutApiListPaymentLinksOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query. If a cursor is not provided, the endpoint returns the first page of the results. For more  information, see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination). | 
 **limit** | **optional.Int32**| A limit on the number of results to return per page. The limit is advisory and the implementation might return more or less results. If the supplied limit is negative, zero, or greater than the maximum limit of 1000, it is ignored.  Default value: &#x60;100&#x60; | 

### Return type

[**ListPaymentLinksResponse**](ListPaymentLinksResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrievePaymentLink**
> RetrievePaymentLinkResponse RetrievePaymentLink(ctx, id)
RetrievePaymentLink

Retrieves a payment link.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of link to retrieve. | 

### Return type

[**RetrievePaymentLinkResponse**](RetrievePaymentLinkResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePaymentLink**
> UpdatePaymentLinkResponse UpdatePaymentLink(ctx, body, id)
UpdatePaymentLink

Updates a payment link. You can update the `payment_link` fields such as `description`, `checkout_options`, and  `pre_populated_data`. You cannot update other fields such as the `order_id`, `version`, `URL`, or `timestamp` field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatePaymentLinkRequest**](UpdatePaymentLinkRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **id** | **string**| The ID of the payment link to update. | 

### Return type

[**UpdatePaymentLinkResponse**](UpdatePaymentLinkResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

