# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ListWebhooks**](V1WebhooksApi.md#V1ListWebhooks) | **Get** /v1/{location_id}/webhooks | V1ListWebhooks
[**V1UpdateWebhooks**](V1WebhooksApi.md#V1UpdateWebhooks) | **Put** /v1/{location_id}/webhooks | V1UpdateWebhooks

# **V1ListWebhooks**
> []V1WebhooksEvents V1ListWebhooks(ctx, locationId)
V1ListWebhooks

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

# **V1UpdateWebhooks**
> []V1WebhooksEvents V1UpdateWebhooks(ctx, locationId)
V1UpdateWebhooks

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

