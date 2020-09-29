# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListWebhooks**](V1WebhooksApi.md#ListWebhooks) | **Get** /v1/{location_id}/webhooks | ListWebhooks
[**UpdateWebhooks**](V1WebhooksApi.md#UpdateWebhooks) | **Put** /v1/{location_id}/webhooks | UpdateWebhooks

# **ListWebhooks**
> []V1WebhooksEvents ListWebhooks(ctx, locationId)
ListWebhooks

Lists which types of events trigger webhook notifications for a particular location. See the [V1 Webhooks API guide](https://developer.squareup.com/docs/webhooks-api/what-it-does-v1) for more information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to list webhook notification types for. | 

### Return type

[**[]V1WebhooksEvents**](V1WebhooksEvents.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWebhooks**
> []V1WebhooksEvents UpdateWebhooks(ctx, locationId)
UpdateWebhooks

Changes the webhook event subscriptions for a location. See the [V1 Webhooks API guide](https://developer.squareup.com/docs/webhooks-api/what-it-does-v1) for more information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to list webhook notification types for. | 

### Return type

[**[]V1WebhooksEvents**](V1WebhooksEvents.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

