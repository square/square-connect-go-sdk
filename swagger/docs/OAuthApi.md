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

As part of a URL sent to a seller to authorize permissions for  the developer, `Authorize` displays an authorization page and a  list of requested permissions. This is not a callable API endpoint.  The completed URL looks similar to the following example: https://connect.squareup.com/oauth2/authorize?client_id={YOUR_APP_ID}&scope=CUSTOMERS_WRITE+CUSTOMERS_READ&session=False&state=82201dd8d83d23cc8a48caf52b  The seller can approve or deny the permissions. If approved,` Authorize`  returns an `AuthorizeResponse` that is sent to the redirect URL and includes  a state string and an authorization code. The code is used in the `ObtainToken`  call to obtain an access token and a refresh token that the developer uses  to manage resources on behalf of the seller.  __Important:__ The `AuthorizeResponse` is sent to the redirect URL that you set on  the OAuth page of your application in the Developer Dashboard.  If an error occurs or the seller denies the request, `Authorize` returns an  error response that includes `error` and `error_description` values. If the  error is due to the seller denying the request, the error value is `access_denied`  and the `error_description` is `user_denied`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clientId** | **string**| The Square-issued ID for your application, available from  the OAuth page for your application on the Developer Dashboard. | 
 **optional** | ***OAuthApiAuthorizeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OAuthApiAuthorizeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | [**optional.Interface of OAuthPermission**](.md)| A space-separated list of the permissions that the application is requesting. Default: \&quot;&#x60;MERCHANT_PROFILE_READ PAYMENTS_READ SETTLEMENTS_READ BANK_ACCOUNTS_READ&#x60;\&quot; | 
 **locale** | **optional.String**| The locale to present the permission request form in. Square detects the appropriate locale automatically. Only provide this value if the application can definitively determine the preferred locale.  Currently supported values: &#x60;en-IE&#x60;, &#x60;en-US&#x60;, &#x60;en-CA&#x60;, &#x60;es-US&#x60;, &#x60;fr-CA&#x60;, and &#x60;ja-JP&#x60;. | 
 **session** | **optional.Bool**| If &#x60;false&#x60;, the user must log in to their Square account to view the Permission Request form, even if they already have a valid user session. This value has no effect in Sandbox. Default: &#x60;true&#x60; | [default to false]
 **state** | **optional.String**| When provided, &#x60;state&#x60; is passed to the configured redirect URL after the Permission Request form is submitted. You can include &#x60;state&#x60; and verify its value to help protect against cross-site request forgery. | 
 **codeChallenge** | **optional.String**| When provided, the oauth flow will use PKCE to authorize. The &#x60;code_challenge&#x60; will be associated with the authorization_code and a &#x60;code_verifier&#x60; will need to passed in to obtain the access token. | 

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

Returns an OAuth access token and a refresh token unless the  `short_lived` parameter is set to `true`, in which case the endpoint  returns only an access token.  The `grant_type` parameter specifies the type of OAuth request. If  `grant_type` is `authorization_code`, you must include the authorization  code you received when a seller granted you authorization. If `grant_type`  is `refresh_token`, you must provide a valid refresh token. If you are using  an old version of the Square APIs (prior to March 13, 2019), `grant_type`  can be `migration_token` and you must provide a valid migration token.  You can use the `scopes` parameter to limit the set of permissions granted  to the access token and refresh token. You can use the `short_lived` parameter  to create an access token that expires in 24 hours.  __Note:__ OAuth tokens should be encrypted and stored on a secure server.  Application clients should never interact directly with OAuth tokens.

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

`RenewToken` is deprecated. For information about refreshing OAuth access tokens, see [Migrate from Renew to Refresh OAuth Tokens](https://developer.squareup.com/docs/oauth-api/migrate-to-refresh-tokens).  Renews an OAuth access token before it expires.  OAuth access tokens besides your application's personal access token expire after 30 days. You can also renew expired tokens within 15 days of their expiration. You cannot renew an access token that has been expired for more than 15 days. Instead, the associated user must recomplete the OAuth flow from the beginning.  __Important:__ The `Authorization` header for this endpoint must have the following format:  ``` Authorization: Client APPLICATION_SECRET ```  Replace `APPLICATION_SECRET` with the application secret on the Credentials page in the [Developer Dashboard](https://developer.squareup.com/apps).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RenewTokenRequest**](RenewTokenRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **clientId** | **string**| Your application ID, which is available in the OAuth page in the [Developer Dashboard](https://developer.squareup.com/apps). | 

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

Revokes an access token generated with the OAuth flow.  If an account has more than one OAuth access token for your application, this endpoint revokes all of them, regardless of which token you specify. When an OAuth access token is revoked, all of the active subscriptions associated with that OAuth token are canceled immediately.  __Important:__ The `Authorization` header for this endpoint must have the following format:  ``` Authorization: Client APPLICATION_SECRET ```  Replace `APPLICATION_SECRET` with the application secret on the OAuth page for your application on the Developer Dashboard.

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

