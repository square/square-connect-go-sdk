# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSnippet**](SnippetsApi.md#DeleteSnippet) | **Delete** /v2/sites/{site_id}/snippet | DeleteSnippet
[**RetrieveSnippet**](SnippetsApi.md#RetrieveSnippet) | **Get** /v2/sites/{site_id}/snippet | RetrieveSnippet
[**UpsertSnippet**](SnippetsApi.md#UpsertSnippet) | **Post** /v2/sites/{site_id}/snippet | UpsertSnippet

# **DeleteSnippet**
> DeleteSnippetResponse DeleteSnippet(ctx, siteId)
DeleteSnippet

Removes your snippet from a Square Online site.  You can call [ListSites](api-endpoint:Sites-ListSites) to get the IDs of the sites that belong to a seller.   __Note:__ Square Online APIs are publicly available as part of an early access program. For more information, see [Early access program for Square Online APIs](https://developer.squareup.com/docs/online-api#early-access-program-for-square-online-apis).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**| The ID of the site that contains the snippet. | 

### Return type

[**DeleteSnippetResponse**](DeleteSnippetResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveSnippet**
> RetrieveSnippetResponse RetrieveSnippet(ctx, siteId)
RetrieveSnippet

Retrieves your snippet from a Square Online site. A site can contain snippets from multiple snippet applications, but you can retrieve only the snippet that was added by your application.  You can call [ListSites](api-endpoint:Sites-ListSites) to get the IDs of the sites that belong to a seller.   __Note:__ Square Online APIs are publicly available as part of an early access program. For more information, see [Early access program for Square Online APIs](https://developer.squareup.com/docs/online-api#early-access-program-for-square-online-apis).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **siteId** | **string**| The ID of the site that contains the snippet. | 

### Return type

[**RetrieveSnippetResponse**](RetrieveSnippetResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpsertSnippet**
> UpsertSnippetResponse UpsertSnippet(ctx, body, siteId)
UpsertSnippet

Adds a snippet to a Square Online site or updates the existing snippet on the site.  The snippet code is appended to the end of the `head` element on every page of the site, except checkout pages. A snippet application can add one snippet to a given site.   You can call [ListSites](api-endpoint:Sites-ListSites) to get the IDs of the sites that belong to a seller.   __Note:__ Square Online APIs are publicly available as part of an early access program. For more information, see [Early access program for Square Online APIs](https://developer.squareup.com/docs/online-api#early-access-program-for-square-online-apis).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpsertSnippetRequest**](UpsertSnippetRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **siteId** | **string**| The ID of the site where you want to add or update the snippet. | 

### Return type

[**UpsertSnippetResponse**](UpsertSnippetResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

