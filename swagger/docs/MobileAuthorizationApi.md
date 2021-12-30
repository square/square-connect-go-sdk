# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMobileAuthorizationCode**](MobileAuthorizationApi.md#CreateMobileAuthorizationCode) | **Post** /mobile/authorization-code | CreateMobileAuthorizationCode

# **CreateMobileAuthorizationCode**
> CreateMobileAuthorizationCodeResponse CreateMobileAuthorizationCode(ctx, body)
CreateMobileAuthorizationCode

Generates code to authorize a mobile application to connect to a Square card reader.  Authorization codes are one-time-use codes and expire 60 minutes after being issued.  __Important:__ The `Authorization` header you provide to this endpoint must have the following format:  ``` Authorization: Bearer ACCESS_TOKEN ```  Replace `ACCESS_TOKEN` with a [valid production authorization credential](https://developer.squareup.com/docs/build-basics/access-tokens).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateMobileAuthorizationCodeRequest**](CreateMobileAuthorizationCodeRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateMobileAuthorizationCodeResponse**](CreateMobileAuthorizationCodeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

