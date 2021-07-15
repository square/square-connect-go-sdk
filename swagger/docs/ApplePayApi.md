# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegisterDomain**](ApplePayApi.md#RegisterDomain) | **Post** /v2/apple-pay/domains | RegisterDomain

# **RegisterDomain**
> RegisterDomainResponse RegisterDomain(ctx, body)
RegisterDomain

Activates a domain for use with Apple Pay on the Web and Square. A validation is performed on this domain by Apple to ensure that it is properly set up as an Apple Pay enabled domain.  This endpoint provides an easy way for platform developers to bulk activate Apple Pay on the Web with Square for merchants using their platform.  To learn more about Web Apple Pay, see [Add the Apple Pay on the Web Button](https://developer.squareup.com/docs/payment-form/add-digital-wallets/apple-pay).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RegisterDomainRequest**](RegisterDomainRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**RegisterDomainResponse**](RegisterDomainResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

