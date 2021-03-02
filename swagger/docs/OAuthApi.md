# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Authorize**](OAuthApi.md#Authorize) | **Get** /oauth2/authorize | Authorize
[**ObtainToken**](OAuthApi.md#ObtainToken) | **Post** /oauth2/token | ObtainToken
[**RenewToken**](OAuthApi.md#RenewToken) | **Post** /oauth2/clients/{client_id}/access-token/renew | RenewToken
[**RevokeToken**](OAuthApi.md#RevokeToken) | **Post** /oauth2/revoke | RevokeToken

# **Authorize**
> AuthorizeResponse Authorize(ctx, clientId, optional)
Authorize

Presents a Permission Request form that returns an access code to be exchanged during the OAuth flow for a valid OAuth access token. To send users to the Permission Request form and start the OAuth flow, configure a link with the desired permissions that directs users to the OAuth Authorization endpoint.  In the event of an error, Authorize returns an error response (`error` and `error_description`). If the failure is a result of the user denying the request, the value is `access_denied` with a description of `user_denied`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**| The Square-issued ID of the application requesting permissions. | 
 **optional** | ***OAuthApiAuthorizeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OAuthApiAuthorizeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | [**optional.Interface of OAuthPermission**](.md)| __OPTIONAL__  A space-separated list of the permissions the application is requesting. Default: \&quot;&#x60;MERCHANT_PROFILE_READ PAYMENTS_READ SETTLEMENTS_READ BANK_ACCOUNTS_READ&#x60;\&quot; | 
 **locale** | **optional.String**| __OPTIONAL__  The locale to present the permission request form in. Square detects the appropriate locale automatically. Only provide this value if the application can definitively determine the preferred locale.  Currently supported values: &#x60;en-US&#x60;, &#x60;en-CA&#x60;, &#x60;es-US&#x60;, &#x60;fr-CA&#x60;, &#x60;ja-JP&#x60;. | 
 **session** | **optional.Bool**| If &#x60;false&#x60;, the user must log in to their Square account to view the Permission Request form, even if they already have a valid user session. Default: &#x60;true&#x60; | [default to false]
 **state** | **optional.String**| __OPTIONAL__  When provided, &#x60;state&#x60; is passed along to the configured Redirect URL after the Permission Request form is submitted. You can include state and verify its value to help protect against cross-site request forgery. | 

### Return type

[**AuthorizeResponse**](AuthorizeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ObtainToken**
> ObtainTokenResponse ObtainToken(ctx, body)
ObtainToken

Returns an OAuth access token.  The endpoint supports distinct methods of obtaining OAuth access tokens. Applications specify a method by adding the `grant_type` parameter in the request and also provide relevant information.  __Note:__ Regardless of the method application specified, the endpoint always returns two items; an OAuth access token and a refresh token in the response.  __OAuth tokens should only live on secure servers. Application clients should never interact directly with OAuth tokens__.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ObtainTokenRequest**](ObtainTokenRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**ObtainTokenResponse**](ObtainTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenewToken**
> RenewTokenResponse RenewToken(ctx, body, clientId)
RenewToken

`RenewToken` is deprecated. For information about refreshing OAuth access tokens, see [Renew OAuth Token](https://developer.squareup.com/docs/oauth-api/cookbook/renew-oauth-tokens).   Renews an OAuth access token before it expires.  OAuth access tokens besides your application's personal access token expire after __30 days__. You can also renew expired tokens within __15 days__ of their expiration. You cannot renew an access token that has been expired for more than 15 days. Instead, the associated user must re-complete the OAuth flow from the beginning.  __Important:__ The `Authorization` header for this endpoint must have the following format:  ``` Authorization: Client APPLICATION_SECRET ```  Replace `APPLICATION_SECRET` with the application secret on the Credentials page in the [application dashboard](https://connect.squareup.com/apps).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RenewTokenRequest**](RenewTokenRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **clientId** | **string**| Your application ID, available from the [application dashboard](https://connect.squareup.com/apps). | 

### Return type

[**RenewTokenResponse**](RenewTokenResponse.md)

### Authorization

[oauth2ClientSecret](../README.md#oauth2ClientSecret)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeToken**
> RevokeTokenResponse RevokeToken(ctx, body)
RevokeToken

Revokes an access token generated with the OAuth flow.  If an account has more than one OAuth access token for your application, this endpoint revokes all of them, regardless of which token you specify. When an OAuth access token is revoked, all of the active subscriptions associated with that OAuth token are canceled immediately.  __Important:__ The `Authorization` header for this endpoint must have the following format:  ``` Authorization: Client APPLICATION_SECRET ```  Replace `APPLICATION_SECRET` with the application secret on the Credentials page in the [Developer Dashboard](https://developer.squareup.com/apps).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RevokeTokenRequest**](RevokeTokenRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**RevokeTokenResponse**](RevokeTokenResponse.md)

### Authorization

[oauth2ClientSecret](../README.md#oauth2ClientSecret)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

