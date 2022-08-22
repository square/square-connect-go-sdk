# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhookSubscription**](WebhookSubscriptionsApi.md#CreateWebhookSubscription) | **Post** /v2/webhooks/subscriptions | CreateWebhookSubscription
[**DeleteWebhookSubscription**](WebhookSubscriptionsApi.md#DeleteWebhookSubscription) | **Delete** /v2/webhooks/subscriptions/{subscription_id} | DeleteWebhookSubscription
[**ListWebhookEventTypes**](WebhookSubscriptionsApi.md#ListWebhookEventTypes) | **Get** /v2/webhooks/event-types | ListWebhookEventTypes
[**ListWebhookSubscriptions**](WebhookSubscriptionsApi.md#ListWebhookSubscriptions) | **Get** /v2/webhooks/subscriptions | ListWebhookSubscriptions
[**RetrieveWebhookSubscription**](WebhookSubscriptionsApi.md#RetrieveWebhookSubscription) | **Get** /v2/webhooks/subscriptions/{subscription_id} | RetrieveWebhookSubscription
[**TestWebhookSubscription**](WebhookSubscriptionsApi.md#TestWebhookSubscription) | **Post** /v2/webhooks/subscriptions/{subscription_id}/test | TestWebhookSubscription
[**UpdateWebhookSubscription**](WebhookSubscriptionsApi.md#UpdateWebhookSubscription) | **Put** /v2/webhooks/subscriptions/{subscription_id} | UpdateWebhookSubscription
[**UpdateWebhookSubscriptionSignatureKey**](WebhookSubscriptionsApi.md#UpdateWebhookSubscriptionSignatureKey) | **Post** /v2/webhooks/subscriptions/{subscription_id}/signature-key | UpdateWebhookSubscriptionSignatureKey

# **CreateWebhookSubscription**
> CreateWebhookSubscriptionResponse CreateWebhookSubscription(ctx, body)
CreateWebhookSubscription

Creates a webhook subscription.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateWebhookSubscriptionRequest**](CreateWebhookSubscriptionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateWebhookSubscriptionResponse**](CreateWebhookSubscriptionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWebhookSubscription**
> DeleteWebhookSubscriptionResponse DeleteWebhookSubscription(ctx, subscriptionId)
DeleteWebhookSubscription

Deletes a webhook subscription.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**| [REQUIRED] The ID of the [Subscription](entity:WebhookSubscription) to delete. | 

### Return type

[**DeleteWebhookSubscriptionResponse**](DeleteWebhookSubscriptionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWebhookEventTypes**
> ListWebhookEventTypesResponse ListWebhookEventTypes(ctx, optional)
ListWebhookEventTypes

Lists all webhook event types that can be subscribed to.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhookSubscriptionsApiListWebhookEventTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookSubscriptionsApiListWebhookEventTypesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The API version for which to list event types. Setting this field overrides the default version used by the application. | 

### Return type

[**ListWebhookEventTypesResponse**](ListWebhookEventTypesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWebhookSubscriptions**
> ListWebhookSubscriptionsResponse ListWebhookSubscriptions(ctx, optional)
ListWebhookSubscriptions

Lists all webhook subscriptions owned by your application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhookSubscriptionsApiListWebhookSubscriptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhookSubscriptionsApiListWebhookSubscriptionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for your original query.  For more information, see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination). | 
 **includeDisabled** | **optional.Bool**| Includes disabled [Subscription](entity:WebhookSubscription)s. By default, all enabled [Subscription](entity:WebhookSubscription)s are returned. | [default to false]
 **sortOrder** | [**optional.Interface of SortOrder**](.md)| Sorts the returned list by when the [Subscription](entity:WebhookSubscription) was created with the specified order. This field defaults to ASC. | 
 **limit** | **optional.Int32**| The maximum number of results to be returned in a single page. It is possible to receive fewer results than the specified limit on a given page. The default value of 100 is also the maximum allowed value. If the provided value is greater than 100, it is ignored and the default value is used instead.  Default: 100 | 

### Return type

[**ListWebhookSubscriptionsResponse**](ListWebhookSubscriptionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveWebhookSubscription**
> RetrieveWebhookSubscriptionResponse RetrieveWebhookSubscription(ctx, subscriptionId)
RetrieveWebhookSubscription

Retrieves a webhook subscription identified by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**| [REQUIRED] The ID of the [Subscription](entity:WebhookSubscription) to retrieve. | 

### Return type

[**RetrieveWebhookSubscriptionResponse**](RetrieveWebhookSubscriptionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestWebhookSubscription**
> TestWebhookSubscriptionResponse TestWebhookSubscription(ctx, body, subscriptionId)
TestWebhookSubscription

Tests a webhook subscription by sending a test event to the notification URL.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TestWebhookSubscriptionRequest**](TestWebhookSubscriptionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **subscriptionId** | **string**| [REQUIRED] The ID of the [Subscription](entity:WebhookSubscription) to test. | 

### Return type

[**TestWebhookSubscriptionResponse**](TestWebhookSubscriptionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWebhookSubscription**
> UpdateWebhookSubscriptionResponse UpdateWebhookSubscription(ctx, body, subscriptionId)
UpdateWebhookSubscription

Updates a webhook subscription.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateWebhookSubscriptionRequest**](UpdateWebhookSubscriptionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **subscriptionId** | **string**| [REQUIRED] The ID of the [Subscription](entity:WebhookSubscription) to update. | 

### Return type

[**UpdateWebhookSubscriptionResponse**](UpdateWebhookSubscriptionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWebhookSubscriptionSignatureKey**
> UpdateWebhookSubscriptionSignatureKeyResponse UpdateWebhookSubscriptionSignatureKey(ctx, body, subscriptionId)
UpdateWebhookSubscriptionSignatureKey

Updates a webhook subscription by replacing the existing signature key with a new one.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateWebhookSubscriptionSignatureKeyRequest**](UpdateWebhookSubscriptionSignatureKeyRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **subscriptionId** | **string**| [REQUIRED] The ID of the [Subscription](entity:WebhookSubscription) to update. | 

### Return type

[**UpdateWebhookSubscriptionSignatureKeyResponse**](UpdateWebhookSubscriptionSignatureKeyResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

