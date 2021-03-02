# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelSubscription**](SubscriptionsApi.md#CancelSubscription) | **Post** /v2/subscriptions/{subscription_id}/cancel | CancelSubscription
[**CreateSubscription**](SubscriptionsApi.md#CreateSubscription) | **Post** /v2/subscriptions | CreateSubscription
[**ListSubscriptionEvents**](SubscriptionsApi.md#ListSubscriptionEvents) | **Get** /v2/subscriptions/{subscription_id}/events | ListSubscriptionEvents
[**RetrieveSubscription**](SubscriptionsApi.md#RetrieveSubscription) | **Get** /v2/subscriptions/{subscription_id} | RetrieveSubscription
[**SearchSubscriptions**](SubscriptionsApi.md#SearchSubscriptions) | **Post** /v2/subscriptions/search | SearchSubscriptions
[**UpdateSubscription**](SubscriptionsApi.md#UpdateSubscription) | **Put** /v2/subscriptions/{subscription_id} | UpdateSubscription

# **CancelSubscription**
> CancelSubscriptionResponse CancelSubscription(ctx, subscriptionId)
CancelSubscription

Sets the `canceled_date` field to the end of the active billing period. After this date, the status changes from ACTIVE to CANCELED.

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

Creates a subscription for a customer to a subscription plan.  If you provide a card on file in the request, Square charges the card for  the subscription. Otherwise, Square bills an invoice to the customer's email  address. The subscription starts immediately, unless the request includes  the optional `start_date`. Each individual subscription is associated with a particular location.

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

 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for the original query.  For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | 
 **limit** | **optional.Int32**| The upper limit on the number of subscription events to return  in the response.   Default: &#x60;200&#x60; | 

### Return type

[**ListSubscriptionEventsResponse**](ListSubscriptionEventsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveSubscription**
> RetrieveSubscriptionResponse RetrieveSubscription(ctx, subscriptionId)
RetrieveSubscription

Retrieves a subscription.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionId** | **string**| The ID of the subscription to retrieve. | 

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

Searches for subscriptions.  Results are ordered chronologically by subscription creation date. If the request specifies more than one location ID,  the endpoint orders the result  by location ID, and then by creation date within each location. If no locations are given in the query, all locations are searched.  You can also optionally specify `customer_ids` to search by customer.  If left unset, all customers  associated with the specified locations are returned.  If the request specifies customer IDs, the endpoint orders results  first by location, within location by customer ID, and within  customer by subscription creation date.  For more information, see  [Retrieve subscriptions](https://developer.squareup.com/docs/subscriptions-api/overview#retrieve-subscriptions).

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

# **UpdateSubscription**
> UpdateSubscriptionResponse UpdateSubscription(ctx, body, subscriptionId)
UpdateSubscription

Updates a subscription. You can set, modify, and clear the  `subscription` field values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSubscriptionRequest**](UpdateSubscriptionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **subscriptionId** | **string**| The ID for the subscription to update. | 

### Return type

[**UpdateSubscriptionResponse**](UpdateSubscriptionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

