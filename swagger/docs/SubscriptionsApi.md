# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelSubscription**](SubscriptionsApi.md#CancelSubscription) | **Post** /v2/subscriptions/{subscription_id}/cancel | CancelSubscription
[**CreateSubscription**](SubscriptionsApi.md#CreateSubscription) | **Post** /v2/subscriptions | CreateSubscription
[**DeleteSubscriptionAction**](SubscriptionsApi.md#DeleteSubscriptionAction) | **Delete** /v2/subscriptions/{subscription_id}/actions/{action_id} | DeleteSubscriptionAction
[**ListSubscriptionEvents**](SubscriptionsApi.md#ListSubscriptionEvents) | **Get** /v2/subscriptions/{subscription_id}/events | ListSubscriptionEvents
[**PauseSubscription**](SubscriptionsApi.md#PauseSubscription) | **Post** /v2/subscriptions/{subscription_id}/pause | PauseSubscription
[**ResumeSubscription**](SubscriptionsApi.md#ResumeSubscription) | **Post** /v2/subscriptions/{subscription_id}/resume | ResumeSubscription
[**RetrieveSubscription**](SubscriptionsApi.md#RetrieveSubscription) | **Get** /v2/subscriptions/{subscription_id} | RetrieveSubscription
[**SearchSubscriptions**](SubscriptionsApi.md#SearchSubscriptions) | **Post** /v2/subscriptions/search | SearchSubscriptions
[**SwapPlan**](SubscriptionsApi.md#SwapPlan) | **Post** /v2/subscriptions/{subscription_id}/swap-plan | SwapPlan
[**UpdateSubscription**](SubscriptionsApi.md#UpdateSubscription) | **Put** /v2/subscriptions/{subscription_id} | UpdateSubscription

# **CancelSubscription**
> CancelSubscriptionResponse CancelSubscription(ctx, subscriptionId)
CancelSubscription

Schedules a `CANCEL` action to cancel an active subscription  by setting the `canceled_date` field to the end of the active billing period  and changing the subscription status from ACTIVE to CANCELED after this date.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**| The ID of the subscription to cancel. | 

### Return type

[**CancelSubscriptionResponse**](CancelSubscriptionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSubscription**
> CreateSubscriptionResponse CreateSubscription(ctx, body)
CreateSubscription

Creates a subscription to a subscription plan by a customer.  If you provide a card on file in the request, Square charges the card for the subscription. Otherwise, Square bills an invoice to the customer's email address. The subscription starts immediately, unless the request includes the optional `start_date`. Each individual subscription is associated with a particular location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSubscriptionRequest**](CreateSubscriptionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateSubscriptionResponse**](CreateSubscriptionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSubscriptionAction**
> DeleteSubscriptionActionResponse DeleteSubscriptionAction(ctx, subscriptionId, actionId)
DeleteSubscriptionAction

Deletes a scheduled action for a subscription.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**| The ID of the subscription the targeted action is to act upon. | 
  **actionId** | **string**| The ID of the targeted action to be deleted. | 

### Return type

[**DeleteSubscriptionActionResponse**](DeleteSubscriptionActionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSubscriptionEvents**
> ListSubscriptionEventsResponse ListSubscriptionEvents(ctx, subscriptionId, optional)
ListSubscriptionEvents

Lists all events for a specific subscription. In the current implementation, only `START_SUBSCRIPTION` and `STOP_SUBSCRIPTION` (when the subscription was canceled) events are returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**| The ID of the subscription to retrieve the events for. | 
 **optional** | ***SubscriptionsApiListSubscriptionEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubscriptionsApiListSubscriptionEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| When the total number of resulting subscription events exceeds the limit of a paged response,  specify the cursor returned from a preceding response here to fetch the next set of results. If the cursor is unset, the response contains the last page of the results.  For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | 
 **limit** | **optional.Int32**| The upper limit on the number of subscription events to return in a paged response. | 

### Return type

[**ListSubscriptionEventsResponse**](ListSubscriptionEventsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PauseSubscription**
> PauseSubscriptionResponse PauseSubscription(ctx, body, subscriptionId)
PauseSubscription

Schedules a `PAUSE` action to pause an active subscription.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PauseSubscriptionRequest**](PauseSubscriptionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **subscriptionId** | **string**| The ID of the subscription to pause. | 

### Return type

[**PauseSubscriptionResponse**](PauseSubscriptionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResumeSubscription**
> ResumeSubscriptionResponse ResumeSubscription(ctx, body, subscriptionId)
ResumeSubscription

Schedules a `RESUME` action to resume a paused or a deactivated subscription.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResumeSubscriptionRequest**](ResumeSubscriptionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **subscriptionId** | **string**| The ID of the subscription to resume. | 

### Return type

[**ResumeSubscriptionResponse**](ResumeSubscriptionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveSubscription**
> RetrieveSubscriptionResponse RetrieveSubscription(ctx, subscriptionId, optional)
RetrieveSubscription

Retrieves a subscription.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**| The ID of the subscription to retrieve. | 
 **optional** | ***SubscriptionsApiRetrieveSubscriptionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubscriptionsApiRetrieveSubscriptionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **optional.String**| A query parameter to specify related information to be included in the response.   The supported query parameter values are:   - &#x60;actions&#x60;: to include scheduled actions on the targeted subscription. | 

### Return type

[**RetrieveSubscriptionResponse**](RetrieveSubscriptionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSubscriptions**
> SearchSubscriptionsResponse SearchSubscriptions(ctx, body)
SearchSubscriptions

Searches for subscriptions.  Results are ordered chronologically by subscription creation date. If the request specifies more than one location ID, the endpoint orders the result by location ID, and then by creation date within each location. If no locations are given in the query, all locations are searched.  You can also optionally specify `customer_ids` to search by customer. If left unset, all customers associated with the specified locations are returned. If the request specifies customer IDs, the endpoint orders results first by location, within location by customer ID, and within customer by subscription creation date.  For more information, see [Retrieve subscriptions](https://developer.squareup.com/docs/subscriptions-api/overview#retrieve-subscriptions).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchSubscriptionsRequest**](SearchSubscriptionsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchSubscriptionsResponse**](SearchSubscriptionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwapPlan**
> SwapPlanResponse SwapPlan(ctx, body, subscriptionId)
SwapPlan

Schedules a `SWAP_PLAN` action to swap a subscription plan in an existing subscription.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SwapPlanRequest**](SwapPlanRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **subscriptionId** | **string**| The ID of the subscription to swap the subscription plan for. | 

### Return type

[**SwapPlanResponse**](SwapPlanResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSubscription**
> UpdateSubscriptionResponse UpdateSubscription(ctx, body, subscriptionId)
UpdateSubscription

Updates a subscription. You can set, modify, and clear the `subscription` field values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSubscriptionRequest**](UpdateSubscriptionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **subscriptionId** | **string**| The ID of the subscription to update. | 

### Return type

[**UpdateSubscriptionResponse**](UpdateSubscriptionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

